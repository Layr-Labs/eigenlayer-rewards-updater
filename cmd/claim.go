package cmd

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/Layr-Labs/eigenlayer-payment-proofs/pkg/claimgen"
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
	"go.uber.org/zap"
	"log"
	"os"
	"time"
)

func runClaimgen(
	config *config.ClaimConfig,
	logger *zap.Logger,
) (
	*claimgen.IPaymentCoordinatorPaymentMerkleClaimStrings,
	error,
) {
	ctx := context.Background()

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
	slo := drv.NewServiceLimitOverride()
	slo.SetDMLQueryTimeout(10)
	conf.SetWorkGroup(drv.NewDefaultWG("eigenLabs_workgroup", nil, nil))
	conf.SetServiceLimitOverride(*slo)

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

	distro, _, err := dds.GetLatestSubmittedDistribution(ctx)
	if err != nil {
		return nil, err
	}

	cg := claimgen.NewClaimgen(distro)

	tokenAddresses := make([]gethcommon.Address, 0)
	for _, t := range config.Tokens {
		t = t
		tokenAddresses = append(tokenAddresses, gethcommon.HexToAddress(t))
	}
	fmt.Printf("Tokens: %+v\n", tokenAddresses)

	accounts, claim, err := cg.GenerateClaimProofForEarner(
		gethcommon.HexToAddress(config.EarnerAddress),
		tokenAddresses,
		0,
	)
	if err != nil {
		return nil, err
	}

	solidity := claimgen.FormatProofForSolidity(accounts.Root(), claim)

	// transactor.GetPaymentCoordinator().ProcessClaim(&bind.TransactOpts{
	// 	From:   gethcommon.HexToAddress(config.EarnerAddress),
	// 	NoSend: true,
	// })

	return solidity, nil

}

// distribution represents the updater command
var claimCmd = &cobra.Command{
	Use:   "claim",
	Short: "Generate claim",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.NewClaimConfig()
		fmt.Printf("Config: %+v\n", cfg)
		logger, err := logger.NewLogger(&logger.LoggerConfig{
			Debug: cfg.Debug,
		})
		if err != nil {
			log.Fatalln(err)
		}
		defer logger.Sync()
		logger.Sugar().Debug(cfg)

		solidity, err := runClaimgen(cfg, logger)

		if err != nil {
			logger.Sugar().Error(err)
		}

		jsonString, err := json.MarshalIndent(solidity, "", "  ")

		if err != nil {
			logger.Sugar().Fatal("Failed to marshal solidity", zap.Error(err))
		}

		if cfg.Output != "" {
			path := fmt.Sprintf("%s/%d.json", cfg.Output, time.Now().Unix())
			err := os.WriteFile(path, jsonString, 0755)
			if err != nil {
				logger.Sugar().Fatal("Failed to write to output file", zap.Error(err))
			}
		} else {
			fmt.Printf("distribution: %+v\n", string(jsonString))
		}
	},
}

func init() {
	fmt.Println("Claim init")
	rootCmd.AddCommand(claimCmd)

	claimCmd.Flags().String("environment", "dev", "The environment to use")
	claimCmd.Flags().String("network", "localnet", "Which network to use")
	claimCmd.Flags().String("rpc-url", "", "https://ethereum-holesky-rpc.publicnode.com")
	claimCmd.Flags().String("private-key", "", "An ethereum private key")
	claimCmd.Flags().String("aws-access-key-id", "", "AWS access key ID")
	claimCmd.Flags().String("aws-secret-access-key", "", "AWS secret access key")
	claimCmd.Flags().String("aws-region", "us-east-1", "us-east-1")
	claimCmd.Flags().String("s3-output-bucket", "", "s3://<bucket name and path>")
	claimCmd.Flags().String("payment-coordinator-address", "0x56c119bD92Af45eb74443ab14D4e93B7f5C67896", "Ethereum address of the payment coordinator contract")
	claimCmd.Flags().String("output", "", "File to write output json to")
	claimCmd.Flags().String("earnerAddress", "", "Address of the earner")
	claimCmd.Flags().StringSlice("tokens", []string{}, "List of token addresses")

	claimCmd.Flags().VisitAll(func(f *pflag.Flag) {
		if err := viper.BindPFlag(config.KebabToSnakeCase(f.Name), f); err != nil {
			fmt.Printf("Failed to bind flag '%s' - %+v\n", f.Name, err)
		}
	})

}
