package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Layr-Labs/eigenlayer-rewards-proofs/pkg/claimgen"
	"github.com/Layr-Labs/eigenlayer-rewards-updater/internal/logger"
	"github.com/Layr-Labs/eigenlayer-rewards-updater/internal/metrics"
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

func runClaimgen(
	ctx context.Context,
	cfg *config.ClaimConfig,
	l *zap.Logger,
) (
	*claimgen.IRewardsCoordinatorRewardsMerkleClaimStrings,
	error,
) {
	span, ctx := ddTracer.StartSpanFromContext(ctx, "runClaimgen")
	defer span.Finish()

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

	e, _ := config.StringEnvironmentFromEnum(cfg.Environment)
	dataFetcher := httpProofDataFetcher.NewHttpProofDataFetcher(cfg.ProofStoreBaseUrl, e, cfg.Network, http.DefaultClient, l)

	transactor, err := services.NewTransactor(chainClient, gethcommon.HexToAddress(cfg.RewardsCoordinatorAddress))
	if err != nil {
		l.Sugar().Errorf("Failed to initialize transactor", zap.Error(err))
		return nil, err
	}

	claimDate := cfg.ClaimTimestamp
	var rootIndex uint32

	if cfg.ClaimTimestamp == "latest" {
		l.Sugar().Info("Generating claim based on latest submitted reward")
		transactor, err := services.NewTransactor(chainClient, gethcommon.HexToAddress(cfg.RewardsCoordinatorAddress))
		if err != nil {
			l.Sugar().Errorf("Failed to initialize transactor", zap.Error(err))
			return nil, err
		}

		latestSubmittedTimestamp, err := transactor.CurrRewardsCalculationEndTimestamp()
		if err != nil {
			l.Sugar().Errorf("Failed to get latest submitted timestamp", zap.Error(err))
			return nil, err
		}
		claimDate = time.Unix(int64(latestSubmittedTimestamp), 0).UTC().Format(time.DateOnly)
		l.Sugar().Debug("Latest submitted timestamp", zap.String("claimDate", claimDate))

		rootCount, err := transactor.GetNumberOfPublishedRoots()
		if err != nil {
			l.Sugar().Errorf("Failed to get number of published roots", zap.Error(err))
			return nil, err
		}
		rootIndex = uint32(rootCount.Uint64() - 1)
	} else {
		return nil, fmt.Errorf("Claim timestamp must be 'latest'")
	}

	proofData, err := dataFetcher.FetchClaimAmountsForDate(ctx, claimDate)
	if err != nil {
		l.Sugar().Errorf("Failed to fetch proof data", zap.Error(err))
		return nil, err
	}

	cg := claimgen.NewClaimgen(proofData.Distribution)

	tokenAddresses := make([]gethcommon.Address, 0)
	for _, t := range cfg.Tokens {
		t = t
		tokenAddresses = append(tokenAddresses, gethcommon.HexToAddress(t))
	}

	accounts, claim, err := cg.GenerateClaimProofForEarner(
		gethcommon.HexToAddress(cfg.EarnerAddress),
		tokenAddresses,
		rootIndex,
	)
	if err != nil {
		return nil, err
	}

	solidity := claimgen.FormatProofForSolidity(accounts.Root(), claim)

	if cfg.SubmitClaim {
		metrics.GetStatsdClient().Incr(metrics.Counter_ClaimsGenerated, nil, 1)
		err := transactor.SubmitRewardClaim(ctx, *claim, gethcommon.HexToAddress(cfg.EarnerAddress))
		if err != nil {
			metrics.GetStatsdClient().Incr(metrics.Counter_ClaimsSubmittedFail, nil, 1)
			l.Sugar().Errorf("Failed to submit claim", zap.Error(err))
		} else {
			metrics.GetStatsdClient().Incr(metrics.Counter_ClaimsSubmittedSuccess, nil, 1)
		}
	}

	return solidity, nil

}

// distribution represents the updater command
var claimCmd = &cobra.Command{
	Use:   "claim",
	Short: "Generate claim",
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

		solidity, err := runClaimgen(ctx, cfg, logger)

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
			fmt.Printf("%s\n", string(jsonString))
		}
	},
}

func init() {
	rootCmd.AddCommand(claimCmd)

	claimCmd.Flags().String("environment", "dev", "The environment to use")
	claimCmd.Flags().String("network", "localnet", "Which network to use")
	claimCmd.Flags().String("rpc-url", "", "https://ethereum-holesky-rpc.publicnode.com")
	claimCmd.Flags().String("private-key", "", "An ethereum private key")
	claimCmd.Flags().String("rewards-coordinator-address", "0x56c119bD92Af45eb74443ab14D4e93B7f5C67896", "Ethereum address of the rewards coordinator contract")
	claimCmd.Flags().String("output", "", "File to write output json to")
	claimCmd.Flags().String("earner-address", "", "Address of the earner")
	claimCmd.Flags().StringSlice("tokens", []string{}, "List of token addresses")
	claimCmd.Flags().String("proof-store-base-url", "", "HTTP base url where data is stored")
	claimCmd.Flags().String("claim-timestamp", "", "YYYY-MM-DD - Timestamp of the rewards root to claim against")
	claimCmd.Flags().Bool("submit-claim", false, "Post the claim to the rewards coordinator")

	claimCmd.Flags().VisitAll(func(f *pflag.Flag) {
		if err := viper.BindPFlag(config.KebabToSnakeCase(f.Name), f); err != nil {
			fmt.Printf("Failed to bind flag '%s' - %+v\n", f.Name, err)
		}
		viper.BindEnv(f.Name)
	})

}
