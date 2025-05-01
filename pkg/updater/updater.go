package updater

import (
	"context"
	"fmt"
	"github.com/Layr-Labs/eigenlayer-rewards-proofs/pkg/utils"
	"github.com/Layr-Labs/eigenlayer-rewards-updater/internal/metrics"
	"github.com/Layr-Labs/eigenlayer-rewards-updater/pkg/services"
	"github.com/Layr-Labs/eigenlayer-rewards-updater/pkg/sidecar"
	rewardsV1 "github.com/Layr-Labs/protocol-apis/gen/protos/eigenlayer/sidecar/v1/rewards"
	"go.uber.org/zap"
	ddTracer "gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
	"time"
)

type Updater struct {
	transactor    services.Transactor
	logger        *zap.Logger
	sidecarClient *sidecar.SidecarClient
}

func NewUpdater(
	transactor services.Transactor,
	sc *sidecar.SidecarClient,
	logger *zap.Logger,
) (*Updater, error) {
	return &Updater{
		transactor:    transactor,
		logger:        logger,
		sidecarClient: sc,
	}, nil
}

type UpdatedRoot struct {
	SnapshotDate string
	Root         string
}

// Update fetches the most recent snapshot and the most recent submitted timestamp from the chain.
func (u *Updater) Update(ctx context.Context) (*UpdatedRoot, error) {
	span, ctx := ddTracer.StartSpanFromContext(ctx, "updater::Update")
	defer span.Finish()

	u.logger.Sugar().Infow("Generating a new rewards snapshot (this may take a while, please wait)")
	res, err := u.sidecarClient.Rewards.GenerateRewards(ctx, &rewardsV1.GenerateRewardsRequest{
		RespondWithRewardsData: false,
		WaitForComplete:        true,
		CutoffDate:             "latest",
	})
	if err != nil {
		return nil, fmt.Errorf("failed to generate rewards: %w", err)
	}

	u.logger.Sugar().Infow("Generating a rewards root",
		zap.String("cutoffDate", res.CutoffDate),
	)
	rootRes, err := u.sidecarClient.Rewards.GenerateRewardsRoot(ctx, &rewardsV1.GenerateRewardsRootRequest{
		CutoffDate: res.CutoffDate,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to generate rewards root: %w", err)
	}

	u.logger.Sugar().Debugw("Rewards snapshot generated",
		zap.String("rewardsCalculationEndDate", rootRes.RewardsCalcEndDate),
		zap.String("root", rootRes.RewardsRoot),
	)

	rewardsCalcEnd, err := time.Parse(time.DateOnly, rootRes.RewardsCalcEndDate)
	if err != nil {
		return nil, fmt.Errorf("failed to parse snapshot date: %w", err)
	}

	rootBytes, err := utils.ConvertStringToBytes(rootRes.RewardsRoot)
	if err != nil {
		return nil, fmt.Errorf("failed to convert root to bytes: %w", err)
	}

	// send the merkle root to the smart contract
	u.logger.Sugar().Infow("updating rewards", zap.String("new_root", rootRes.RewardsRoot))

	u.logger.Sugar().Infow("Calculated timestamp",
		zap.Int64("calculated_until_timestamp", rewardsCalcEnd.Unix()),
		zap.String("calculated_until_date", rewardsCalcEnd.Format(time.DateOnly)),
	)
	if err := u.transactor.SubmitRoot(ctx, [32]byte(rootBytes), uint32(rewardsCalcEnd.Unix())); err != nil {
		metrics.GetStatsdClient().Incr(metrics.Counter_UpdateFails, nil, 1)
		metrics.IncCounterUpdateRun(metrics.CounterUpdateRunsFailed)
		u.logger.Sugar().Errorw("Failed to submit root", zap.Error(err))
		return nil, err
	} else {
		metrics.GetStatsdClient().Incr(metrics.Counter_UpdateSuccess, nil, 1)
		metrics.IncCounterUpdateRun(metrics.CounterUpdateRunsSuccess)
	}

	return &UpdatedRoot{
		SnapshotDate: rootRes.RewardsCalcEndDate,
		Root:         rootRes.RewardsRoot,
	}, nil
}
