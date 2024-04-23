package cmd

import (
	"context"
	"database/sql"
	"fmt"
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
)

func runUpdater(config *config.UpdaterConfig) {
	ctx := context.Background()

	ethClient, err := ethclient.Dial(config.RPCUrl)
	if err != nil {
		fmt.Println("Failed to create new eth client")
		panic(err)
	}

	chainClient, err := pkg.NewChainClient(ethClient, config.PrivateKey)
	if err != nil {
		fmt.Println("Failed to create new chain client with private key")
		panic(err)
	}

	transactor, err := services.NewTransactor(chainClient, gethcommon.HexToAddress(config.PaymentCoordinatorAddress))
	if err != nil {
		fmt.Println("Failed to initialize transactor")
		panic(err)
	}

	// Step 1. Set AWS Credential in Driver Config.
	conf, _ := drv.NewDefaultConfig(config.S3OutputBucket, config.AWSRegion, config.AWSAccessKeyId, config.AWSSecretAccessKey)
	// Step 2. Open Connection.
	db, _ := sql.Open(drv.DriverName, conf.Stringify())
	defer db.Close()

	envNetwork, err := config.GetEnvNetwork()
	if err != nil {
		panic(err)
	}
	dds := services.NewDistributionDataService(db, transactor, &services.DistributionDataServiceConfig{
		EnvNetwork: envNetwork,
	})

	u, err := updater.NewUpdater(transactor, dds)
	if err != nil {
		fmt.Println("Failed to create updater")
		panic(err)
	}

	if err := u.Update(ctx); err != nil {
		fmt.Println("Failed to update")
		panic(err)
	}
}

// updaterCmd represents the updater command
var updaterCmd = &cobra.Command{
	Use:   "updater",
	Short: "Generate and update payments merkle tree",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.NewUpdaterConfig()

		runUpdater(cfg)
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
		fmt.Printf("flag: %v\n", config.KebabToSnakeCase(f.Name))
		if err := viper.BindPFlag(config.KebabToSnakeCase(f.Name), f); err != nil {
			fmt.Printf("Failed to bind flag '%s' - %+v\n", f.Name, err)
		}
	})

}
