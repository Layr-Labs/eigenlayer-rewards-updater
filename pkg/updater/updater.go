package updater

import (
	"context"
	"fmt"
	"github.com/Layr-Labs/eigenlayer-payment-updater/pkg/proofDataFetcher"
	"github.com/Layr-Labs/eigenlayer-payment-updater/pkg/services"
	"github.com/Layr-Labs/eigenlayer-payment-updater/pkg/utils"
	"github.com/wealdtech/go-merkletree/v2"
	"go.uber.org/zap"
	"math/big"
	"time"
)

type Updater struct {
	transactor              services.Transactor
	distributionDataService services.DistributionDataService
	logger                  *zap.Logger
	proofDataFetcher        proofDataFetcher.ProofDataFetcher
}

func NewUpdater(
	transactor services.Transactor,
	distributionDataService services.DistributionDataService,
	fetcher proofDataFetcher.ProofDataFetcher,
	logger *zap.Logger,
) (*Updater, error) {
	return &Updater{
		transactor:              transactor,
		distributionDataService: distributionDataService,
		logger:                  logger,
		proofDataFetcher:        fetcher,
	}, nil
}

func (u *Updater) Update(ctx context.Context) (*merkletree.MerkleTree, error) {
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

	if err := u.transactor.SubmitRoot(ctx, [32]byte(newRoot), big.NewInt(int64(calculatedUntilTimestamp))); err != nil {
		fmt.Printf("Failed to submit root: %v\n", err)
		u.logger.Sugar().Error("Failed to submit root", zap.Error(err))
		return paymentProofData.AccountTree, err
	}

	return paymentProofData.AccountTree, nil
}
