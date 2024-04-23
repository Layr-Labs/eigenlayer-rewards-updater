package cmd

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Layr-Labs/eigenlayer-payment-updater/internal/logger"
	"github.com/Layr-Labs/eigenlayer-payment-updater/pkg"
	"github.com/Layr-Labs/eigenlayer-payment-updater/pkg/config"
	"github.com/Layr-Labs/eigenlayer-payment-updater/pkg/services"
	"github.com/Layr-Labs/eigenlayer-payment-updater/pkg/updater"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	drv "github.com/uber/athenadriver/go"
	"go.uber.org/zap"
	"log"
)

func runUpdater(config *config.UpdaterConfig, logger *zap.Logger) error {
	ctx := context.Background()

	ethClient, err := ethclient.Dial(config.RPCUrl)
	if err != nil {
		fmt.Println("Failed to create new eth client")
		logger.Sugar().Errorf("Failed to create new eth client", zap.Error(err))
		return err
	}

	chainClient, err := pkg.NewChainClient(ethClient, config.PrivateKey)
	if err != nil {
		logger.Sugar().Errorf("Failed to create new chain client with private key", zap.Error(err))
		return err
	}

	transactor, err := services.NewTransactor(chainClient, gethcommon.HexToAddress(config.PaymentCoordinatorAddress))
	if err != nil {
		logger.Sugar().Errorf("Failed to initialize transactor", zap.Error(err))
		return err
	}

	// Step 1. Set AWS Credential in Driver Config.
	conf, err := drv.NewDefaultConfig(config.S3OutputBucket, config.AWSRegion, config.AWSAccessKeyId, config.AWSSecretAccessKey)
	if err != nil {
		logger.Sugar().Errorf("Failed to create athena driver config", zap.Error(err))
		return err
	}

	// Step 2. Open Connection.
	db, err := sql.Open(drv.DriverName, conf.Stringify())
	if err != nil {
		logger.Sugar().Errorf("Failed to open database connection", zap.Error(err))
		return err
	}
	defer db.Close()

	envNetwork, err := config.GetEnvNetwork()
	if err != nil {
		logger.Sugar().Errorf("Failed to get EnvNetwork", zap.Error(err))
		return err
	}
	dds := services.NewDistributionDataService(db, transactor, &services.DistributionDataServiceConfig{
		EnvNetwork: envNetwork,
		Logger:     logger,
	})

	u, err := updater.NewUpdater(transactor, dds, logger)
	if err != nil {
		logger.Sugar().Errorf("Failed to create updater", zap.Error(err))
		return err
	}

	if err := u.Update(ctx); err != nil {
		logger.Sugar().Errorf("Failed to update", zap.Error(err))
		return err
	}
	logger.Sugar().Info("Update successful")
	return nil
}

// updaterCmd represents the updater command
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
			logger.Sugar().Fatalln(err)
		}
	},
}

func init() {
	fmt.Println("Updater init")
	rootCmd.AddCommand(updaterCmd)

	updaterCmd.Flags().String("rpc-url", "", "https://ethereum-holesky-rpc.publicnode.com")
	updaterCmd.Flags().String("private-key", "", "An ethereum private key")
	updaterCmd.Flags().String("aws-access-key-id", "", "AWS access key ID")
	updaterCmd.Flags().String("aws-secret-access-key", "", "AWS secret access key")
	updaterCmd.Flags().String("aws-region", "us-east-1", "us-east-1")
	updaterCmd.Flags().String("s3-output-bucket", "", "s3://<bucket name and path>")
	updaterCmd.Flags().String("payment-coordinator-address", "0x56c119bD92Af45eb74443ab14D4e93B7f5C67896", "Ethereum address of the payment coordinator contract")

	updaterCmd.Flags().VisitAll(func(f *pflag.Flag) {
		if err := viper.BindPFlag(config.KebabToSnakeCase(f.Name), f); err != nil {
			fmt.Printf("Failed to bind flag '%s' - %+v\n", f.Name, err)
		}
	})

}
