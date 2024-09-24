package cmd

import (
	"context"
	"fmt"
	"github.com/Layr-Labs/eigenlayer-rewards-updater/internal/logger"
	"github.com/Layr-Labs/eigenlayer-rewards-updater/pkg/chainClient"
	"github.com/Layr-Labs/eigenlayer-rewards-updater/pkg/config"
	"github.com/Layr-Labs/eigenlayer-rewards-updater/pkg/disabler"
	"github.com/Layr-Labs/eigenlayer-rewards-updater/pkg/services"
	"github.com/Layr-Labs/eigenlayer-rewards-updater/pkg/tracer"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	ddTracer "gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
	"log"
)

func runDisableRoot(ctx context.Context, cfg *config.DisableRootConfig, logger *zap.Logger) error {
	span, ctx := ddTracer.StartSpanFromContext(ctx, "runDisableRoot")
	defer span.Finish()

	rootIndex := cfg.RootIndex
	if rootIndex < 0 {
		return fmt.Errorf("root index must be greater than or equal to 0")
	}

	logger.Sugar().Infow("Disabling root", zap.Uint32("rootIndex", cfg.RootIndex))

	ethClient, err := ethclient.Dial(cfg.RPCUrl)
	if err != nil {
		logger.Sugar().Errorf("Failed to create new eth client", zap.Error(err))
		return err
	}

	cc, err := chainClient.NewChainClient(ctx, ethClient, cfg.PrivateKey)
	if err != nil {
		logger.Sugar().Errorf("Failed to create new chain client with private key", zap.Error(err))
		return err
	}

	transactor, err := services.NewTransactor(cc, gethcommon.HexToAddress(cfg.RewardsCoordinatorAddress))
	if err != nil {
		logger.Sugar().Errorf("Failed to initialize transactor", zap.Error(err))
		return err
	}

	d, err := disabler.NewDisabler(transactor, logger)
	if err != nil {
		logger.Sugar().Errorf("Failed to create updater", zap.Error(err))
		return err
	}

	err = d.DisableRoot(cfg.RootIndex)
	if err != nil {
		logger.Sugar().Infow("Failed to disable root", zap.Error(err))
		return nil
	}
	logger.Sugar().Infow("Root disabled", zap.Uint32("rootIndex", cfg.RootIndex))
	return nil
}

// distribution represents the updater command
var disableRootCmd = &cobra.Command{
	Use:   "disable-root",
	Short: "Disable root",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.NewDisableRootConfig()

		tracer.StartTracer(cfg.EnableTracing)
		defer ddTracer.Stop()

		span, ctx := ddTracer.StartSpanFromContext(context.Background(), "cmd::updater")
		defer span.Finish()

		logger, err := logger.NewLogger(&logger.LoggerConfig{
			Debug: cfg.Debug,
		})
		if err != nil {
			log.Fatalln(err)
		}
		defer logger.Sync()

		err = runDisableRoot(ctx, cfg, logger)
		if err != nil {
			logger.Sugar().Error(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(disableRootCmd)

	disableRootCmd.Flags().Int32("root-index", -1, "Index of the root to disable")

	disableRootCmd.Flags().VisitAll(func(f *pflag.Flag) {
		if err := viper.BindPFlag(config.KebabToSnakeCase(f.Name), f); err != nil {
			fmt.Printf("Failed to bind flag '%s' - %+v\n", f.Name, err)
		}
		viper.BindEnv(f.Name)
	})

}
