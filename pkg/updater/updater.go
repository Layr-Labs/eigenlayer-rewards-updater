package updater

import (
	"context"
	"fmt"
	"github.com/Layr-Labs/eigenlayer-payment-updater/pkg/services"
	"github.com/Layr-Labs/eigenlayer-payment-updater/pkg/utils"
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
		u.logger.Sugar().Debug("Failed to merklize distribution roots", zap.Error(err))
		return nil, err
	}

	newRoot := accountTree.Root()

	// send the merkle root to the smart contract
	u.logger.Sugar().Info("updating payments", zap.String("new_root", utils.ConvertBytesToString(newRoot)))

	if err := u.transactor.SubmitRoot(ctx, [32]byte(newRoot), big.NewInt(int64(calculatedUntilTimestamp))); err != nil {
		fmt.Printf("Failed to submit root: %v\n", err)
		u.logger.Sugar().Error("Failed to submit root", zap.Error(err))
		return accountTree, err
	}

	return accountTree, nil
}
