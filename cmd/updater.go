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
	"github.com/Layr-Labs/eigenlayer-rewards-updater/pkg/updater"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	ddTracer "gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
	"log"
)

func runUpdater(ctx context.Context, cfg *config.UpdaterConfig, logger *zap.Logger) error {
	span, ctx := ddTracer.StartSpanFromContext(ctx, "runUpdater")
	defer span.Finish()

	ethClient, err := ethclient.Dial(cfg.RPCUrl)
	if err != nil {
		logger.Sugar().Errorf("Failed to create new eth client", zap.Error(err))
		return err
	}

	cc, err := chainClient.NewChainClient(ctx, ethClient, cfg.PrivateKey)
	if err != nil {
		logger.Sugar().Errorf("Failed to create new chain client with private key", zap.Error(err))
		return err
	}

	sidecarClient, err := sidecar.NewSidecarClient(cfg.SidecarRpcUrl, cfg.SidecarInsecureRpc)
	if err != nil {
		logger.Sugar().Errorf("Failed to create sidecar client", zap.Error(err))
		return err
	}

	transactor, err := services.NewTransactor(cc, gethcommon.HexToAddress(cfg.RewardsCoordinatorAddress))
	if err != nil {
		logger.Sugar().Errorf("Failed to initialize transactor", zap.Error(err))
		return err
	}

	u, err := updater.NewUpdater(transactor, sidecarClient, logger)
	if err != nil {
		logger.Sugar().Errorf("Failed to create updater", zap.Error(err))
		return err
	}

	_, err = u.Update(ctx)
	if err != nil {
		logger.Sugar().Infow("Failed to update", zap.Error(err))
		return nil
	}
	logger.Sugar().Infow("Update successful")
	return nil
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
		if err != nil {
			logger.Sugar().Error(err)
		}
		if err := s.Close(); err != nil {
			logger.Sugar().Errorw("Failed to close statsd client", zap.Error(err))
		}
		// Push metrics to pushgateway at the end of the run
		if err := metrics.PushToPushgateway(); err != nil {
			logger.Sugar().Errorw("Failed to push metrics to pushgateway", zap.Error(err))
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
