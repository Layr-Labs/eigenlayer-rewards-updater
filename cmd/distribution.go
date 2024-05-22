package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/Layr-Labs/eigenlayer-payment-updater/internal/logger"
	"github.com/Layr-Labs/eigenlayer-payment-updater/pkg/chainClient"
	"github.com/Layr-Labs/eigenlayer-payment-updater/pkg/config"
	"github.com/Layr-Labs/eigenlayer-payment-updater/pkg/proofDataFetcher/httpProofDataFetcher"
	"github.com/Layr-Labs/eigenlayer-payment-updater/pkg/signer/privateKey"
	"github.com/Layr-Labs/eigenlayer-payment-updater/pkg/transactor"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"time"
)

type Result struct {
	LatestPaymentDate      string `json:"latestPaymentDate"`
	MostRecentSnapshotDate string `json:"mostRecentSnapshotDate"`
}

func run(
	cfg *config.DistributionConfig,
	l *zap.Logger,
) (
	result *Result,
	err error,
) {
	ethClient, err := ethclient.Dial(cfg.RPCUrl)
	if err != nil {
		fmt.Println("Failed to create new eth client")
		l.Sugar().Errorf("Failed to create new eth client", zap.Error(err))
		return nil, err
	}

	signer, err := privateKey.NewPrivateKeySigner(cfg.PrivateKey)
	if err != nil {
		l.Sugar().Error("Failed to create new private key signer", zap.Error(err))
	}

	chainClient, err := chainClient.NewChainClient(ethClient, signer, l)
	if err != nil {
		l.Sugar().Errorf("Failed to create new chain client with private key", zap.Error(err))
		return nil, err
	}

	transactor, err := transactor.NewTransactor(chainClient, gethcommon.HexToAddress(cfg.PaymentCoordinatorAddress), l)
	if err != nil {
		l.Sugar().Errorf("Failed to initialize transactor", zap.Error(err))
		return nil, err
	}

	e, _ := config.StringEnvironmentFromEnum(cfg.Environment)
	dataFetcher := httpProofDataFetcher.NewHttpProofDataFetcher(cfg.ProofStoreBaseUrl, e, cfg.Network, http.DefaultClient, l)

	latestSnapshot, err := dataFetcher.FetchLatestSnapshot()
	if err != nil {
		return nil, err
	}

	l.Sugar().Debugf("latest snapshot: %s", latestSnapshot.GetDateString())

	latestSubmittedTimestamp, err := transactor.CurrPaymentCalculationEndTimestamp()
	lst := time.Unix(int64(latestSubmittedTimestamp), 0).UTC()

	l.Sugar().Debugf("latest submitted timestamp: %s", lst.Format(time.DateOnly))

	return &Result{
		LatestPaymentDate:      lst.Format(time.DateOnly),
		MostRecentSnapshotDate: latestSnapshot.GetDateString(),
	}, nil

}

// distribution represents the updater command
var distributionCmd = &cobra.Command{
	Use:   "distribution",
	Short: "Access distribution data",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.NewDistributionConfig()
		fmt.Printf("Config: %+v\n", cfg)
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

		fmt.Printf("result: %+v\n", res)

		jsonRes, err := json.MarshalIndent(res, "", "  ")

		if err != nil {
			logger.Sugar().Fatal("Failed to marshal results", zap.Error(err))
		}

		if cfg.Output != "" {
			path := fmt.Sprintf("%s/%d.json", cfg.Output, jsonRes)
			err := os.WriteFile(path, jsonRes, 0755)
			if err != nil {
				logger.Sugar().Fatal("Failed to write to output file", zap.Error(err))
			}
		} else {
			fmt.Printf("distribution: %+v\n", string(jsonRes))
		}
	},
}

func init() {
	rootCmd.AddCommand(distributionCmd)

	distributionCmd.Flags().String("environment", "dev", "The environment to use")
	distributionCmd.Flags().String("network", "localnet", "Which network to use")
	distributionCmd.Flags().String("rpc-url", "", "https://ethereum-holesky-rpc.publicnode.com")
	distributionCmd.Flags().String("private-key", "", "An ethereum private key")
	distributionCmd.Flags().String("payment-coordinator-address", "0x56c119bD92Af45eb74443ab14D4e93B7f5C67896", "Ethereum address of the payment coordinator contract")
	distributionCmd.Flags().String("output", "", "File to write output json to")
	distributionCmd.Flags().String("proof-store-base-url", "", "HTTP base url where data is stored")

	distributionCmd.Flags().VisitAll(func(f *pflag.Flag) {
		if err := viper.BindPFlag(config.KebabToSnakeCase(f.Name), f); err != nil {
			fmt.Printf("Failed to bind flag '%s' - %+v\n", f.Name, err)
		}
	})

}
