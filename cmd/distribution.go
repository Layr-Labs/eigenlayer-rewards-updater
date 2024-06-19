package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Layr-Labs/eigenlayer-rewards-updater/internal/logger"
	"github.com/Layr-Labs/eigenlayer-rewards-updater/pkg/chainClient"
	"github.com/Layr-Labs/eigenlayer-rewards-updater/pkg/config"
	"github.com/Layr-Labs/eigenlayer-rewards-updater/pkg/proofDataFetcher/httpProofDataFetcher"
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
	"net/http"
	"os"
	"time"
)

type Result struct {
	// TODO(seanmcgary): update json field name once pipeline is updated
	LatestRewardDate       string `json:"latestPaymentDate"`
	MostRecentSnapshotDate string `json:"mostRecentSnapshotDate"`
}

func runDistribution(
	ctx context.Context,
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

	chainClient, err := chainClient.NewChainClient(ctx, ethClient, cfg.PrivateKey)
	if err != nil {
		l.Sugar().Errorf("Failed to create new chain client with private key", zap.Error(err))
		return nil, err
	}

	transactor, err := services.NewTransactor(chainClient, gethcommon.HexToAddress(cfg.RewardsCoordinatorAddress))
	if err != nil {
		l.Sugar().Errorf("Failed to initialize transactor", zap.Error(err))
		return nil, err
	}

	e, _ := config.StringEnvironmentFromEnum(cfg.Environment)
	dataFetcher := httpProofDataFetcher.NewHttpProofDataFetcher(cfg.ProofStoreBaseUrl, e, cfg.Network, http.DefaultClient, l)

	latestSnapshot, err := dataFetcher.FetchLatestSnapshot(ctx)
	if err != nil {
		return nil, err
	}

	l.Sugar().Debugf("latest snapshot: %s", latestSnapshot.GetDateString())

	latestSubmittedTimestamp, err := transactor.CurrRewardsCalculationEndTimestamp()
	lst := time.Unix(int64(latestSubmittedTimestamp), 0).UTC()

	l.Sugar().Debugf("latest submitted timestamp: %s", lst.Format(time.DateOnly))

	return &Result{
		LatestRewardDate:       lst.Format(time.DateOnly),
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

		tracer.StartTracer(cfg.EnableTracing)
		defer ddTracer.Stop()

		span, ctx := ddTracer.StartSpanFromContext(context.Background(), "cmd::distribution")
		defer span.Finish()

		logger, err := logger.NewLogger(&logger.LoggerConfig{
			Debug: cfg.Debug,
		})
		if err != nil {
			log.Fatalln(err)
		}
		defer logger.Sync()

		res, err := runDistribution(ctx, cfg, logger)

		if err != nil {
			logger.Sugar().Error(err)
		}

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
		}
	},
}

func init() {
	rootCmd.AddCommand(distributionCmd)

	distributionCmd.Flags().String("environment", "dev", "The environment to use")
	distributionCmd.Flags().String("network", "localnet", "Which network to use")
	distributionCmd.Flags().String("rpc-url", "", "https://ethereum-holesky-rpc.publicnode.com")
	distributionCmd.Flags().String("private-key", "", "An ethereum private key")
	distributionCmd.Flags().String("rewards-coordinator-address", "0x56c119bD92Af45eb74443ab14D4e93B7f5C67896", "Ethereum address of the rewards coordinator contract")
	distributionCmd.Flags().String("output", "", "File to write output json to")
	distributionCmd.Flags().String("proof-store-base-url", "", "HTTP base url where data is stored")

	distributionCmd.Flags().VisitAll(func(f *pflag.Flag) {
		if err := viper.BindPFlag(config.KebabToSnakeCase(f.Name), f); err != nil {
			fmt.Printf("Failed to bind flag '%s' - %+v\n", f.Name, err)
		}
		viper.BindEnv(f.Name)
	})

}
