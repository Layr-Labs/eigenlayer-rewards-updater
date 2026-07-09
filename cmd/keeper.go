package cmd

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/Layr-Labs/eigenlayer-rewards-updater/internal/logger"
	"github.com/Layr-Labs/eigenlayer-rewards-updater/internal/metrics"
	"github.com/Layr-Labs/eigenlayer-rewards-updater/pkg/config"
	"github.com/Layr-Labs/eigenlayer-rewards-updater/pkg/keeper"
	"github.com/Layr-Labs/eigenlayer-rewards-updater/pkg/tracer"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	ddTracer "gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func runKeeper(ctx context.Context, cfg *config.UpdaterConfig, logger *zap.Logger) error {
	span, ctx := ddTracer.StartSpanFromContext(ctx, "runKeeper")
	defer span.Finish()

	if err := cfg.Validate(); err != nil {
		logger.Sugar().Errorw("Invalid keeper config", zap.Error(err))
		return err
	}
	if cfg.EmissionsControllerAddress == "" {
		return errors.New("emissions_controller_address is required for the keeper command")
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
		zap.String("emissions_controller", cfg.EmissionsControllerAddress),
	)

	emissions, err := keeper.NewEmissionsController(cc, gethcommon.HexToAddress(cfg.EmissionsControllerAddress))
	if err != nil {
		logger.Sugar().Errorw("Failed to create emissions controller client", zap.Error(err))
		return err
	}

	k := keeper.NewKeeper(emissions, logger)
	if _, err := k.Press(ctx); err != nil {
		logger.Sugar().Errorw("Failed to press emissions button", zap.Error(err))
		return err
	}
	logger.Sugar().Infow("Keeper run successful")
	return nil
}

// keeperCmd presses the EmissionsController button (weekly emissions) via KMS.
var keeperCmd = &cobra.Command{
	Use:   "keeper",
	Short: "Press the EmissionsController button to mint/process weekly emissions",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.NewUpdaterConfig()

		tracer.StartTracer(cfg.EnableTracing)
		defer ddTracer.Stop()

		span, ctx := ddTracer.StartSpanFromContext(context.Background(), "cmd::keeper")
		defer span.Finish()

		s, err := metrics.InitStatsdClient(cfg.DDStatsdUrl, cfg.EnableStatsd)
		if err != nil {
			log.Fatalln(err)
		}

		s.Incr(metrics.Counter_KeeperRuns, nil, 1)

		l, err := logger.NewLogger(&logger.LoggerConfig{
			Debug: cfg.Debug,
		})
		if err != nil {
			log.Fatalln(err)
		}
		defer l.Sync()

		err = runKeeper(ctx, cfg, l)

		// Flush metrics on ALL exit paths: os.Exit skips deferred calls, so the
		// just-emitted keeper_press_fails would otherwise be dropped before the
		// statsd buffer is sent. Close() flushes; do it before exiting.
		if cerr := s.Close(); cerr != nil {
			l.Sugar().Errorw("Failed to close statsd client", zap.Error(cerr))
		}
		if err != nil {
			l.Sugar().Errorw("Keeper run failed", zap.Error(err))
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(keeperCmd)

	keeperCmd.Flags().VisitAll(func(f *pflag.Flag) {
		if err := viper.BindPFlag(config.KebabToSnakeCase(f.Name), f); err != nil {
			fmt.Printf("Failed to bind flag '%s' - %+v\n", f.Name, err)
		}
		viper.BindEnv(f.Name)
	})
}
