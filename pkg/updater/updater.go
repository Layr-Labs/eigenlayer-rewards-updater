package updater

import (
	"context"
	"fmt"
	"github.com/Layr-Labs/eigenlayer-rewards-proofs/pkg/utils"
	"github.com/Layr-Labs/eigenlayer-rewards-updater/internal/metrics"
	"github.com/Layr-Labs/eigenlayer-rewards-updater/pkg/proofDataFetcher"
	"github.com/Layr-Labs/eigenlayer-rewards-updater/pkg/services"
	"github.com/wealdtech/go-merkletree/v2"
	"go.uber.org/zap"
	ddTracer "gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
	"time"
)

type Updater struct {
	transactor       services.Transactor
	logger           *zap.Logger
	proofDataFetcher proofDataFetcher.ProofDataFetcher
}

func NewUpdater(
	transactor services.Transactor,
	fetcher proofDataFetcher.ProofDataFetcher,
	logger *zap.Logger,
) (*Updater, error) {
	return &Updater{
		transactor:       transactor,
		logger:           logger,
		proofDataFetcher: fetcher,
	}, nil
}

// Update fetches the most recent snapshot and the most recent submitted timestamp from the chain.
func (u *Updater) Update(ctx context.Context) (*merkletree.MerkleTree, error) {
	span, ctx := ddTracer.StartSpanFromContext(ctx, "updater::Update")
	defer span.Finish()
	//ctx = opentracing.ContextWithSpan(ctx, span)
	/*
		1. Fetch the list of most recently generated snapshots (list of timestamps)
		2. Get the timestamp of the most recently submitted on-chain rewards
		3. If the most recent snapshot is less than or equal to the latest on-chain rewards, no new rewards exists.
		4. Fetch the claim amounts generated for the new snapshot based on snapshot date
		5. Generate Merkle tree from claim amounts
		6. Submit the new Merkle root to the smart contract
	*/

	// Get the most recent snapshot timestamp
	latestSnapshot, err := u.proofDataFetcher.FetchLatestSnapshot(ctx)
	if err != nil {
		return nil, err
	}

	u.logger.Sugar().Debugf("latest snapshot: %s", latestSnapshot.GetDateString())

	// Grab latest submitted timestamp directly from chain
	latestSubmittedTimestamp, err := u.transactor.CurrRewardsCalculationEndTimestamp()
	lst := time.Unix(int64(latestSubmittedTimestamp), 0).UTC()

	u.logger.Sugar().Debugf("latest submitted timestamp: %s", lst.Format(time.DateOnly))

	// If most recent snapshot's timestamp is equal to the latest submitted timestamp, then we don't need to update
	if lst.Equal(latestSnapshot.SnapshotDate) {
		metrics.GetStatsdClient().Incr(metrics.Counter_UpdateNoUpdate, nil, 1)
		u.logger.Sugar().Info("latest snapshot is the most recent reward")
		return nil, nil
	}
	// If the most recent snapshot timestamp is less than whats already on chain, we have a problem
	if lst.After(latestSnapshot.SnapshotDate) {
		return nil, fmt.Errorf("recent snapshot occurs before latest submitted timestamp")
	}

	// Get the data for the latest snapshot and load it into a distribution instance
	rewardsProofData, err := u.proofDataFetcher.FetchClaimAmountsForDate(ctx, latestSnapshot.GetDateString())
	if err != nil {
		return nil, err
	}

	newRoot := rewardsProofData.AccountTree.Root()

	calculatedUntilTimestamp := latestSnapshot.SnapshotDate.UTC().Unix()

	// send the merkle root to the smart contract
	u.logger.Sugar().Info("updating rewards", zap.String("new_root", utils.ConvertBytesToString(newRoot)))

	// return rewardsProofData.AccountTree, nil
	u.logger.Sugar().Info("Calculated timestamp", zap.Int64("calculated_until_timestamp", calculatedUntilTimestamp))
	if err := u.transactor.SubmitRoot(ctx, [32]byte(newRoot), uint32(calculatedUntilTimestamp)); err != nil {
		metrics.GetStatsdClient().Incr(metrics.Counter_UpdateFails, nil, 1)
		u.logger.Sugar().Error("Failed to submit root", zap.Error(err))
		return rewardsProofData.AccountTree, err
	} else {
		metrics.GetStatsdClient().Incr(metrics.Counter_UpdateSuccess, nil, 1)
	}

	return rewardsProofData.AccountTree, nil
}
