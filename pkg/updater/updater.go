package updater

import (
	"context"
	"github.com/Layr-Labs/eigenlayer-payment-updater/pkg/services"
	"github.com/wealdtech/go-merkletree/v2"
	"go.uber.org/zap"
	"math/big"
)

type Updater struct {
	transactor              services.Transactor
	distributionDataService services.DistributionDataService
	logger                  *zap.Logger
}

func NewUpdater(
	transactor services.Transactor,
	distributionDataService services.DistributionDataService,
	logger *zap.Logger,
) (*Updater, error) {
	return &Updater{
		transactor:              transactor,
		distributionDataService: distributionDataService,
		logger:                  logger,
	}, nil
}

func (u *Updater) Update(ctx context.Context) (*merkletree.MerkleTree, error) {
	// get the current distribution
	u.logger.Sugar().Info("getting current distribution")
	distribution, calculatedUntilTimestamp, err := u.distributionDataService.GetDistributionToSubmit(ctx)
	if err != nil {
		return nil, err
	}

	// merklize the distribution roots
	u.logger.Sugar().Info("merklizing distribution roots")
	accountTree, _, err := distribution.Merklize()
	if err != nil {
		return nil, err
	}

	var newRoot [32]byte
	copy(newRoot[:], accountTree.Root())

	// send the merkle root to the smart contract
	u.logger.Sugar().Info("updating payments")
	if err := u.transactor.SubmitRoot(ctx, newRoot, big.NewInt(calculatedUntilTimestamp)); err != nil {
		return accountTree, err
	}

	return accountTree, nil
}
