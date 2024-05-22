package updater

import (
	"context"
	"fmt"
	"github.com/Layr-Labs/eigenlayer-payment-proofs/pkg/utils"
	"github.com/Layr-Labs/eigenlayer-payment-updater/pkg/proofDataFetcher"
	"github.com/Layr-Labs/eigenlayer-payment-updater/pkg/transactor"
	"github.com/wealdtech/go-merkletree/v2"
	"go.uber.org/zap"
	"time"
)

type Updater struct {
	transactor       transactor.Transactor
	logger           *zap.Logger
	proofDataFetcher proofDataFetcher.ProofDataFetcher
}

func NewUpdater(
	transactor transactor.Transactor,
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
	/*
		1. Fetch the list of most recently generated snapshots (list of timestamps)
		2. Get the timestamp of the most recently submitted on-chain payment
		3. If the most recent snapshot is less than or equal to the latest on-chain payment, no new payment exists.
		4. Fetch the claim amounts generated for the new snapshot based on snapshot date
		5. Generate Merkle tree from claim amounts
		6. Submit the new Merkle root to the smart contract
	*/

	// Get the most recent snapshot timestamp
	latestSnapshot, err := u.proofDataFetcher.FetchLatestSnapshot()
	if err != nil {
		return nil, err
	}

	u.logger.Sugar().Debugf("latest snapshot: %s", latestSnapshot.GetDateString())

	// Grab latest submitted timestamp directly from chain
	latestSubmittedTimestamp, err := u.transactor.CurrPaymentCalculationEndTimestamp()
	lst := time.Unix(int64(latestSubmittedTimestamp), 0).UTC()

	u.logger.Sugar().Debugf("latest submitted timestamp: %s", lst.Format(time.DateOnly))

	// If most recent snapshot's timestamp is equal to the latest submitted timestamp, then we don't need to update
	if lst.Equal(latestSnapshot.SnapshotDate) {
		return nil, fmt.Errorf("latest snapshot is the most recent payment")
	}
	// If the most recent snapshot timestamp is less than whats already on chain, we have a problem
	if lst.After(latestSnapshot.SnapshotDate) {
		return nil, fmt.Errorf("recent snapshot occurs before latest submitted timestamp")
	}

	// Get the data for the latest snapshot and load it into a distribution instance
	paymentProofData, err := u.proofDataFetcher.FetchClaimAmountsForDate(latestSnapshot.GetDateString())
	if err != nil {
		return nil, err
	}

	newRoot := paymentProofData.AccountTree.Root()

	calculatedUntilTimestamp := latestSnapshot.SnapshotDate.UTC().Unix()

	// send the merkle root to the smart contract
	u.logger.Sugar().Info("updating payments", zap.String("new_root", utils.ConvertBytesToString(newRoot)))

	// return paymentProofData.AccountTree, nil
	fmt.Printf("Calculated timestamp: %+v\n", calculatedUntilTimestamp)
	if err := u.transactor.SubmitRoot(ctx, [32]byte(newRoot), uint32(calculatedUntilTimestamp)); err != nil {
		u.logger.Sugar().Error("Failed to submit root", zap.Error(err))
		return paymentProofData.AccountTree, err
	}

	return paymentProofData.AccountTree, nil
}
