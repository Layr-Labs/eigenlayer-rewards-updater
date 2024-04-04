package updater

import (
	"context"
	"math/big"
	"time"

	"github.com/Layr-Labs/eigenlayer-payment-updater/common/services"
	"github.com/rs/zerolog/log"
)

type Updater struct {
	updateInterval          time.Duration
	transactor              services.Transactor
	distributionDataService services.DistributionDataService
}

func NewUpdater(
	updateIntervalSeconds int,
	transactor services.Transactor,
	distributionDataService services.DistributionDataService,
) (*Updater, error) {
	return &Updater{
		updateInterval:          time.Second * time.Duration(updateIntervalSeconds),
		transactor:              transactor,
		distributionDataService: distributionDataService,
	}, nil
}

func (u *Updater) Start() error {
	// run a loop unning once every u.UpdateInterval that calls u.update()
	log.Info().Msg("service started")
	ctx := context.Background()

	// run the first update immediately
	if err := u.Update(ctx); err != nil {
		log.Error().Msgf("failed to update: %s", err)
	}

	ticker := time.NewTicker(u.updateInterval)
	for range ticker.C {
		log.Info().Msg("running update")
		if err := u.Update(ctx); err != nil {
			log.Error().Msgf("failed to update: %s", err)
		}
	}
	return nil
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
