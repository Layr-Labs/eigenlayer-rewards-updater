package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/Layr-Labs/eigenlayer-payment-proofs/pkg/claimgen"
	"github.com/Layr-Labs/eigenlayer-payment-updater/internal/logger"
	"github.com/Layr-Labs/eigenlayer-payment-updater/pkg/chainClient"
	"github.com/Layr-Labs/eigenlayer-payment-updater/pkg/config"
	"github.com/Layr-Labs/eigenlayer-payment-updater/pkg/proofDataFetcher/httpProofDataFetcher"
	"github.com/Layr-Labs/eigenlayer-payment-updater/pkg/services"
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

func runClaimgen(
	cfg *config.ClaimConfig,
	l *zap.Logger,
) (
	*claimgen.IPaymentCoordinatorPaymentMerkleClaimStrings,
	error,
) {
	ethClient, err := ethclient.Dial(cfg.RPCUrl)
	if err != nil {
		fmt.Println("Failed to create new eth client")
		l.Sugar().Errorf("Failed to create new eth client", zap.Error(err))
		return nil, err
	}

	chainClient, err := chainClient.NewChainClient(ethClient, cfg.PrivateKey)
	if err != nil {
		l.Sugar().Errorf("Failed to create new chain client with private key", zap.Error(err))
		return nil, err
	}

	e, _ := config.StringEnvironmentFromEnum(cfg.Environment)
	dataFetcher := httpProofDataFetcher.NewHttpProofDataFetcher(cfg.ProofStoreBaseUrl, e, cfg.Network, http.DefaultClient, l)

	claimDate := cfg.ClaimTimestamp

	if cfg.ClaimTimestamp == "latest" {
		l.Sugar().Info("Generating claim based on latest submitted payment")
		transactor, err := services.NewTransactor(chainClient, gethcommon.HexToAddress(cfg.PaymentCoordinatorAddress))
		if err != nil {
			l.Sugar().Errorf("Failed to initialize transactor", zap.Error(err))
			return nil, err
		}

		latestSubmittedTimestamp, err := transactor.CurrPaymentCalculationEndTimestamp()
		claimDate = time.Unix(int64(latestSubmittedTimestamp), 0).UTC().Format(time.DateOnly)
	}

	proofData, err := dataFetcher.FetchClaimAmountsForDate(claimDate)
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
	fmt.Printf("Tokens: %+v\n", tokenAddresses)

	accounts, claim, err := cg.GenerateClaimProofForEarner(
		gethcommon.HexToAddress(cfg.EarnerAddress),
		tokenAddresses,
		0,
	)
	if err != nil {
		return nil, err
	}

	solidity := claimgen.FormatProofForSolidity(accounts.Root(), claim)

	// transactor.GetPaymentCoordinator().ProcessClaim(&bind.TransactOpts{
	// 	From:   gethcommon.HexToAddress(cfg.EarnerAddress),
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
	rootCmd.AddCommand(claimCmd)

	claimCmd.Flags().String("environment", "dev", "The environment to use")
	claimCmd.Flags().String("network", "localnet", "Which network to use")
	claimCmd.Flags().String("rpc-url", "", "https://ethereum-holesky-rpc.publicnode.com")
	claimCmd.Flags().String("private-key", "", "An ethereum private key")
	claimCmd.Flags().String("payment-coordinator-address", "0x56c119bD92Af45eb74443ab14D4e93B7f5C67896", "Ethereum address of the payment coordinator contract")
	claimCmd.Flags().String("output", "", "File to write output json to")
	claimCmd.Flags().String("earner-address", "", "Address of the earner")
	claimCmd.Flags().StringSlice("tokens", []string{}, "List of token addresses")
	claimCmd.Flags().String("proof-store-base-url", "", "HTTP base url where data is stored")
	claimCmd.Flags().String("claim-timestamp", "", "YYYY-MM-DD - Timestamp of the payment root to claim against")

	claimCmd.Flags().VisitAll(func(f *pflag.Flag) {
		if err := viper.BindPFlag(config.KebabToSnakeCase(f.Name), f); err != nil {
			fmt.Printf("Failed to bind flag '%s' - %+v\n", f.Name, err)
		}
	})

}
