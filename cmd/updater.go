package cmd

import (
	"context"
	"fmt"
	"github.com/Layr-Labs/eigenlayer-payment-updater/internal/logger"
	"github.com/Layr-Labs/eigenlayer-payment-updater/pkg/chainClient"
	"github.com/Layr-Labs/eigenlayer-payment-updater/pkg/config"
	"github.com/Layr-Labs/eigenlayer-payment-updater/pkg/proofDataFetcher/httpProofDataFetcher"
	"github.com/Layr-Labs/eigenlayer-payment-updater/pkg/services"
	"github.com/Layr-Labs/eigenlayer-payment-updater/pkg/updater"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"log"
	"net/http"
)

func runUpdater(cfg *config.UpdaterConfig, logger *zap.Logger) error {
	ctx := context.Background()

	ethClient, err := ethclient.Dial(cfg.RPCUrl)
	if err != nil {
		fmt.Println("Failed to create new eth client")
		logger.Sugar().Errorf("Failed to create new eth client", zap.Error(err))
		return err
	}

	chainClient, err := chainClient.NewChainClient(ethClient, cfg.PrivateKey)
	if err != nil {
		logger.Sugar().Errorf("Failed to create new chain client with private key", zap.Error(err))
		return err
	}

	e, _ := config.StringEnvironmentFromEnum(cfg.Environment)
	dataFetcher := httpProofDataFetcher.NewHttpProofDataFetcher(cfg.ProofStoreBaseUrl, e, cfg.Network, http.DefaultClient, logger)

	transactor, err := services.NewTransactor(chainClient, gethcommon.HexToAddress(cfg.PaymentCoordinatorAddress))
	if err != nil {
		logger.Sugar().Errorf("Failed to initialize transactor", zap.Error(err))
		return err
	}

	u, err := updater.NewUpdater(transactor, dataFetcher, logger)
	if err != nil {
		logger.Sugar().Errorf("Failed to create updater", zap.Error(err))
		return err
	}

	tree, err := u.Update(ctx)
	fmt.Printf("tree: %+v\n", tree)
	if err != nil {
		logger.Sugar().Info("Failed to update", zap.Error(err))
		return nil
	}
	logger.Sugar().Info("Update successful")
	return nil
}

// distribution represents the updater command
var updaterCmd = &cobra.Command{
	Use:   "updater",
	Short: "Generate and update payments merkle tree",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.NewUpdaterConfig()
		logger, err := logger.NewLogger(&logger.LoggerConfig{
			Debug: cfg.Debug,
		})
		if err != nil {
			log.Fatalln(err)
		}
		defer logger.Sync()
		logger.Sugar().Debug(cfg)

		err = runUpdater(cfg, logger)
		if err != nil {
			logger.Sugar().Error(err)
		}
	},
}

func init() {
	fmt.Println("Updater init")
	rootCmd.AddCommand(updaterCmd)

	updaterCmd.Flags().String("environment", "dev", "The environment to use")
	updaterCmd.Flags().String("network", "localnet", "Which network to use")
	updaterCmd.Flags().String("rpc-url", "", "https://ethereum-holesky-rpc.publicnode.com")
	updaterCmd.Flags().String("private-key", "", "An ethereum private key")
	updaterCmd.Flags().String("aws-access-key-id", "", "AWS access key ID")
	updaterCmd.Flags().String("aws-secret-access-key", "", "AWS secret access key")
	updaterCmd.Flags().String("aws-region", "us-east-1", "us-east-1")
	updaterCmd.Flags().String("s3-output-bucket", "", "s3://<bucket name and path>")
	updaterCmd.Flags().String("payment-coordinator-address", "0x56c119bD92Af45eb74443ab14D4e93B7f5C67896", "Ethereum address of the payment coordinator contract")
	updaterCmd.Flags().String("proof-store-base-url", "", "HTTP base url where data is stored")

	updaterCmd.Flags().VisitAll(func(f *pflag.Flag) {
		if err := viper.BindPFlag(config.KebabToSnakeCase(f.Name), f); err != nil {
			fmt.Printf("Failed to bind flag '%s' - %+v\n", f.Name, err)
		}
	})

}
