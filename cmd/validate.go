package cmd

import (
	"context"
	"fmt"
	"github.com/Layr-Labs/eigenlayer-rewards-updater/internal/logger"
	"github.com/Layr-Labs/eigenlayer-rewards-updater/internal/metrics"
	"github.com/Layr-Labs/eigenlayer-rewards-updater/pkg/chainClient"
	"github.com/Layr-Labs/eigenlayer-rewards-updater/pkg/config"
	"github.com/Layr-Labs/eigenlayer-rewards-updater/pkg/proofDataFetcher/httpProofDataFetcher"
	"github.com/Layr-Labs/eigenlayer-rewards-updater/pkg/services"
	"github.com/Layr-Labs/eigenlayer-rewards-updater/pkg/tracer"
	"github.com/Layr-Labs/eigenlayer-rewards-updater/pkg/validator"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	ddTracer "gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
	"log"
	"net/http"
)

func runValidator(
	ctx context.Context,
	cfg *config.ClaimConfig,
	l *zap.Logger,
) error {
	span, ctx := ddTracer.StartSpanFromContext(ctx, "runClaimgen")
	defer span.Finish()

	ethClient, err := ethclient.Dial(cfg.RPCUrl)
	if err != nil {
		fmt.Println("Failed to create new eth client")
		l.Sugar().Errorf("Failed to create new eth client", zap.Error(err))
		return nil
	}

	chainClient, err := chainClient.NewChainClient(ctx, ethClient, cfg.PrivateKey)
	if err != nil {
		l.Sugar().Errorf("Failed to create new chain client with private key", zap.Error(err))
		return nil
	}

	e, _ := config.StringEnvironmentFromEnum(cfg.Environment)
	dataFetcher := httpProofDataFetcher.NewHttpProofDataFetcher(cfg.ProofStoreBaseUrl, e, cfg.Network, http.DefaultClient, l)

	transactor, err := services.NewTransactor(chainClient, gethcommon.HexToAddress(cfg.RewardsCoordinatorAddress))
	if err != nil {
		l.Sugar().Errorf("Failed to initialize transactor", zap.Error(err))
		return nil
	}

	v := validator.NewValidator(transactor, dataFetcher, l)

	v.ValidatePostedRoot(ctx)

	return nil

}

// distribution represents the updater command
var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate the latest posted root",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.NewClaimConfig()

		metrics.InitStatsdClient(cfg.DDStatsdUrl, cfg.EnableStatsd)

		tracer.StartTracer(cfg.EnableTracing)
		defer ddTracer.Stop()

		span, ctx := ddTracer.StartSpanFromContext(context.Background(), "cmd::claim")
		defer span.Finish()

		logger, err := logger.NewLogger(&logger.LoggerConfig{
			Debug: cfg.Debug,
		})
		if err != nil {
			log.Fatalln(err)
		}
		defer logger.Sync()
		logger.Sugar().Debug(cfg)

		err = runValidator(ctx, cfg, logger)

		if err != nil {
			logger.Sugar().Error(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(validateCmd)

	validateCmd.Flags().String("environment", "dev", "The environment to use")
	validateCmd.Flags().String("network", "localnet", "Which network to use")
	validateCmd.Flags().String("rpc-url", "", "https://ethereum-holesky-rpc.publicnode.com")
	validateCmd.Flags().String("private-key", "", "An ethereum private key")
	validateCmd.Flags().String("rewards-coordinator-address", "0x56c119bD92Af45eb74443ab14D4e93B7f5C67896", "Ethereum address of the rewards coordinator contract")
	validateCmd.Flags().String("proof-store-base-url", "", "HTTP base url where data is stored")

	validateCmd.Flags().VisitAll(func(f *pflag.Flag) {
		if err := viper.BindPFlag(config.KebabToSnakeCase(f.Name), f); err != nil {
			fmt.Printf("Failed to bind flag '%s' - %+v\n", f.Name, err)
		}
		viper.BindEnv(f.Name)
	})

}
