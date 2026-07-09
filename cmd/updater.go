package cmd

import (
	"context"
	"fmt"
	"github.com/Layr-Labs/eigenlayer-rewards-updater/internal/logger"
	"github.com/Layr-Labs/eigenlayer-rewards-updater/internal/metrics"
	"github.com/Layr-Labs/eigenlayer-rewards-updater/pkg/chainClient"
	"github.com/Layr-Labs/eigenlayer-rewards-updater/pkg/config"
	"github.com/Layr-Labs/eigenlayer-rewards-updater/pkg/services"
	"github.com/Layr-Labs/eigenlayer-rewards-updater/pkg/sidecar"
	"github.com/Layr-Labs/eigenlayer-rewards-updater/pkg/tracer"
	"github.com/Layr-Labs/eigenlayer-rewards-updater/pkg/txsigner"
	"github.com/Layr-Labs/eigenlayer-rewards-updater/pkg/updater"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	ddTracer "gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
	"log"
	"os"
)

func runUpdater(ctx context.Context, cfg *config.UpdaterConfig, logger *zap.Logger) error {
	span, ctx := ddTracer.StartSpanFromContext(ctx, "runUpdater")
	defer span.Finish()

	if err := cfg.Validate(); err != nil {
		logger.Sugar().Errorw("Invalid updater config", zap.Error(err))
		return err
	}

	ethClient, err := ethclient.Dial(cfg.RPCUrl)
	if err != nil {
		logger.Sugar().Errorw("Failed to create new eth client", zap.Error(err))
		return err
	}

	cc, err := newChainClientForConfig(ctx, ethClient, cfg, logger)
	if err != nil {
		return err
	}
	logger.Sugar().Infow("Configured signer",
		zap.String("signer_type", cfg.SignerType),
		zap.String("signer_address", cc.AccountAddress.Hex()),
	)

	sidecarClient, err := sidecar.NewSidecarClient(cfg.SidecarRpcUrl, cfg.SidecarInsecureRpc)
	if err != nil {
		logger.Sugar().Errorw("Failed to create sidecar client", zap.Error(err))
		return err
	}

	transactor, err := services.NewTransactor(cc, gethcommon.HexToAddress(cfg.RewardsCoordinatorAddress))
	if err != nil {
		logger.Sugar().Errorw("Failed to initialize transactor", zap.Error(err))
		return err
	}

	u, err := updater.NewUpdater(transactor, sidecarClient, logger)
	if err != nil {
		logger.Sugar().Errorw("Failed to create updater", zap.Error(err))
		return err
	}

	if _, err := u.Update(ctx); err != nil {
		logger.Sugar().Errorw("Failed to update", zap.Error(err))
		return err
	}
	logger.Sugar().Infow("Update successful")
	return nil
}

// newChainClientForConfig builds a chain client whose signer is selected by
// cfg.SignerType: a local private key (default) or AWS KMS.
func newChainClientForConfig(ctx context.Context, ethClient *ethclient.Client, cfg *config.UpdaterConfig, logger *zap.Logger) (*chainClient.ChainClient, error) {
	switch cfg.SignerType {
	case config.SignerTypeAWSKMS:
		signer, err := txsigner.NewAWSKMSSigner(cfg.KMSKeyID, cfg.AWSRegion)
		if err != nil {
			logger.Sugar().Errorw("Failed to create AWS KMS signer", zap.Error(err))
			return nil, err
		}
		cc, err := chainClient.NewChainClientWithSigner(ctx, ethClient, signer)
		if err != nil {
			logger.Sugar().Errorw("Failed to create chain client with KMS signer", zap.Error(err))
			return nil, err
		}
		return cc, nil
	default: // "" or private_key
		cc, err := chainClient.NewChainClient(ctx, ethClient, cfg.PrivateKey)
		if err != nil {
			logger.Sugar().Errorw("Failed to create chain client with private key", zap.Error(err))
			return nil, err
		}
		return cc, nil
	}
}

// distribution represents the updater command
var updaterCmd = &cobra.Command{
	Use:   "updater",
	Short: "Generate and update rewards merkle tree",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.NewUpdaterConfig()

		tracer.StartTracer(cfg.EnableTracing)
		defer ddTracer.Stop()

		span, ctx := ddTracer.StartSpanFromContext(context.Background(), "cmd::updater")
		defer span.Finish()

		s, err := metrics.InitStatsdClient(cfg.DDStatsdUrl, cfg.EnableStatsd)
		if err != nil {
			log.Fatalln(err)
		}

		s.Incr(metrics.Counter_UpdateRuns, nil, 1)

		if cfg.PushgatewayEnabled {
			// Init Pushgateway prometheusPusher client
			err = metrics.InitPrometheusPusherClient(cfg.PushgatewayUrl, "updater")
			if err != nil {
				log.Fatalln(err)
			}
		}

		metrics.IncCounterUpdateRun(metrics.CounterUpdateRunsInvoked)

		logger, err := logger.NewLogger(&logger.LoggerConfig{
			Debug: cfg.Debug,
		})
		if err != nil {
			log.Fatalln(err)
		}
		defer logger.Sync()

		err = runUpdater(ctx, cfg, logger)

		// Flush metrics on ALL exit paths. log.Fatalln/os.Exit skip deferred calls,
		// so on the failure path the just-emitted update_fails would otherwise be
		// dropped from the statsd buffer before it's sent (only the early update_runs
		// would survive). Close() flushes the buffer; do it before exiting.
		if cerr := s.Close(); cerr != nil {
			logger.Sugar().Errorw("Failed to close statsd client", zap.Error(cerr))
		}
		if perr := metrics.PushToPushgateway(); perr != nil {
			logger.Sugar().Errorw("Failed to push metrics to pushgateway", zap.Error(perr))
		}
		if err != nil {
			logger.Sugar().Errorw("Update failed", zap.Error(err))
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(updaterCmd)

	updaterCmd.Flags().VisitAll(func(f *pflag.Flag) {
		if err := viper.BindPFlag(config.KebabToSnakeCase(f.Name), f); err != nil {
			fmt.Printf("Failed to bind flag '%s' - %+v\n", f.Name, err)
		}
		viper.BindEnv(f.Name)
	})

}
