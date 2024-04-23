package updater

import (
	"context"
	"github.com/Layr-Labs/eigenlayer-payment-updater/pkg/services"
	"github.com/rs/zerolog/log"
	"math/big"
)

type Updater struct {
	transactor              services.Transactor
	distributionDataService services.DistributionDataService
}

func NewUpdater(
	transactor services.Transactor,
	distributionDataService services.DistributionDataService,
) (*Updater, error) {
	return &Updater{
		transactor:              transactor,
		distributionDataService: distributionDataService,
	}, nil
}

func (u *Updater) Update(ctx context.Context) error {
	// get the current distribution
	log.Info().Msg("getting current distribution")
	distribution, calculatedUntilTimestamp, err := u.distributionDataService.GetDistributionToSubmit(ctx)
	if err != nil {
		return err
	}

	// merklize the distribution roots
	log.Info().Msg("merklizing distribution roots")
	accountTree, _, err := distribution.Merklize()
	if err != nil {
		return err
	}

	var newRoot [32]byte
	copy(newRoot[:], accountTree.Root())

	// send the merkle root to the smart contract
	log.Info().Msg("updating payments")
	if err := u.transactor.SubmitRoot(ctx, newRoot, big.NewInt(calculatedUntilTimestamp)); err != nil {
		return err
	}

	return nil
}
