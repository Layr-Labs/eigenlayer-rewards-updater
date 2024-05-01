package cmd

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Layr-Labs/eigenlayer-payment-proofs/pkg/distribution"
	"github.com/Layr-Labs/eigenlayer-payment-updater/internal/logger"
	"github.com/Layr-Labs/eigenlayer-payment-updater/pkg"
	"github.com/Layr-Labs/eigenlayer-payment-updater/pkg/config"
	"github.com/Layr-Labs/eigenlayer-payment-updater/pkg/services"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	drv "github.com/uber/athenadriver/go"
	"github.com/wealdtech/go-merkletree/v2"
	"go.uber.org/zap"
	"log"
	"os"
	"time"
)

type Result struct {
	distro      *distribution.Distribution
	timestamp   int64
	accountTree *merkletree.MerkleTree
	tokenTree   map[gethcommon.Address]*merkletree.MerkleTree
}

func run(
	config *config.DistributionConfig,
	logger *zap.Logger,
) (
	result *Result,
	err error,
) {
	ctx := context.Background()
	result = &Result{}

	ethClient, err := ethclient.Dial(config.RPCUrl)
	if err != nil {
		fmt.Println("Failed to create new eth client")
		logger.Sugar().Errorf("Failed to create new eth client", zap.Error(err))
		return nil, err
	}

	chainClient, err := pkg.NewChainClient(ethClient, config.PrivateKey)
	if err != nil {
		logger.Sugar().Errorf("Failed to create new chain client with private key", zap.Error(err))
		return nil, err
	}

	transactor, err := services.NewTransactor(chainClient, gethcommon.HexToAddress(config.PaymentCoordinatorAddress))
	if err != nil {
		logger.Sugar().Errorf("Failed to initialize transactor", zap.Error(err))
		return nil, err
	}

	// Step 1. Set AWS Credential in Driver Config.
	conf, err := drv.NewDefaultConfig(config.S3OutputBucket, config.AWSRegion, config.AWSAccessKeyId, config.AWSSecretAccessKey)
	if err != nil {
		logger.Sugar().Errorf("Failed to create athena driver config", zap.Error(err))
		return nil, err
	}

	// Step 2. Open Connection.
	db, err := sql.Open(drv.DriverName, conf.Stringify())
	if err != nil {
		logger.Sugar().Errorf("Failed to open database connection", zap.Error(err))
		return nil, err
	}
	defer db.Close()

	envNetwork, err := config.GetEnvNetwork()
	if err != nil {
		logger.Sugar().Errorf("Failed to get EnvNetwork", zap.Error(err))
		return nil, err
	}
	dds := services.NewDistributionDataService(db, transactor, &services.DistributionDataServiceConfig{
		EnvNetwork: envNetwork,
		Logger:     logger,
	})

	distro, ts, err := dds.GetLatestSubmittedDistribution(ctx)
	if err != nil {
		return nil, err
	}
	result.distro = distro
	result.timestamp = ts

	accountTree, tokenTree, err := distro.Merklize()
	if err != nil {
		return result, err
	}

	result.accountTree = accountTree
	result.tokenTree = tokenTree

	return result, nil

}

// distribution represents the updater command
var distributionCmd = &cobra.Command{
	Use:   "distribution",
	Short: "Access distribution data",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.NewDistributionConfig()
		logger, err := logger.NewLogger(&logger.LoggerConfig{
			Debug: cfg.Debug,
		})
		if err != nil {
			log.Fatalln(err)
		}
		defer logger.Sync()
		logger.Sugar().Debug(cfg)

		res, err := run(cfg, logger)

		if err != nil {
			logger.Sugar().Error(err)
		}

		logger.Sugar().Infof("Got distribution at timestamp '%s'", time.Unix(res.timestamp, 0).String())

		jsonTree, err := res.distro.MarshalJSON()

		if err != nil {
			logger.Sugar().Fatal("Failed to unmarshal tree", zap.Error(err))
		}

		if cfg.Output != "" {
			path := fmt.Sprintf("%s/%d.json", cfg.Output, res.timestamp)
			err := os.WriteFile(path, jsonTree, 0755)
			if err != nil {
				logger.Sugar().Fatal("Failed to write to output file", zap.Error(err))
			}
		} else {
			fmt.Printf("distribution: %+v\n", string(jsonTree))
		}
	},
}

func init() {
	fmt.Println("Updater init")
	rootCmd.AddCommand(distributionCmd)

	distributionCmd.Flags().String("environment", "dev", "The environment to use")
	distributionCmd.Flags().String("network", "localnet", "Which network to use")
	distributionCmd.Flags().String("rpc-url", "", "https://ethereum-holesky-rpc.publicnode.com")
	distributionCmd.Flags().String("private-key", "", "An ethereum private key")
	distributionCmd.Flags().String("aws-access-key-id", "", "AWS access key ID")
	distributionCmd.Flags().String("aws-secret-access-key", "", "AWS secret access key")
	distributionCmd.Flags().String("aws-region", "us-east-1", "us-east-1")
	distributionCmd.Flags().String("s3-output-bucket", "", "s3://<bucket name and path>")
	distributionCmd.Flags().String("payment-coordinator-address", "0x56c119bD92Af45eb74443ab14D4e93B7f5C67896", "Ethereum address of the payment coordinator contract")
	distributionCmd.Flags().String("output", "", "File to write output json to")

	distributionCmd.Flags().VisitAll(func(f *pflag.Flag) {
		if err := viper.BindPFlag(config.KebabToSnakeCase(f.Name), f); err != nil {
			fmt.Printf("Failed to bind flag '%s' - %+v\n", f.Name, err)
		}
	})

}
