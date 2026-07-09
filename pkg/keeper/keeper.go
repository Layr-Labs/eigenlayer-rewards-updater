package keeper

import (
	"context"
	"fmt"
	"math/big"

	"github.com/Layr-Labs/eigenlayer-rewards-updater/internal/metrics"
	"go.uber.org/zap"
	ddTracer "gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

// MaxUint256 (2^256 - 1) is passed as pressButton's `length` to process every
// pending distribution in one call — the same value the previous keeper used.
var MaxUint256 = new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 256), big.NewInt(1))

type Keeper struct {
	emissions EmissionsController
	logger    *zap.Logger
}

func NewKeeper(emissions EmissionsController, logger *zap.Logger) *Keeper {
	return &Keeper{emissions: emissions, logger: logger}
}

// Result describes what a Press did, for the caller/logs.
type Result struct {
	Pressed bool
	TxHash  string
}

// Press presses the emissions button when it is pressable, and is otherwise a
// healthy no-op. It never submits when isButtonPressable() is false, so re-runs
// (and the once-a-week cron firing against an already-processed epoch) don't
// revert with AllDistributionsProcessed — they exit 0 as "nothing to do".
func (k *Keeper) Press(ctx context.Context) (*Result, error) {
	span, ctx := ddTracer.StartSpanFromContext(ctx, "keeper::Press")
	defer span.Finish()

	pressable, err := k.emissions.IsButtonPressable(ctx)
	if err != nil {
		metrics.GetStatsdClient().Incr(metrics.Counter_KeeperFails, nil, 1)
		return nil, fmt.Errorf("failed to read isButtonPressable: %w", err)
	}

	if !pressable {
		k.logger.Sugar().Infow("Emissions button is not pressable; nothing to do (healthy no-op)")
		metrics.GetStatsdClient().Incr(metrics.Counter_KeeperNoPress, nil, 1)
		return &Result{Pressed: false}, nil
	}

	// Best-effort context for the logs; failures here must not block the press.
	if epoch, err := k.emissions.CurrentEpoch(ctx); err == nil {
		if pending, err := k.emissions.TotalProcessableDistributions(ctx); err == nil {
			k.logger.Sugar().Infow("Emissions button is pressable; pressing",
				zap.String("current_epoch", epoch.String()),
				zap.String("processable_distributions", pending.String()),
			)
		}
	}

	txHash, err := k.emissions.PressButton(ctx, MaxUint256)
	if err != nil {
		metrics.GetStatsdClient().Incr(metrics.Counter_KeeperFails, nil, 1)
		k.logger.Sugar().Errorw("Failed to press emissions button", zap.Error(err))
		return nil, err
	}

	metrics.GetStatsdClient().Incr(metrics.Counter_KeeperPressSuccess, nil, 1)
	k.logger.Sugar().Infow("Pressed emissions button", zap.String("tx_hash", txHash))
	return &Result{Pressed: true, TxHash: txHash}, nil
}
