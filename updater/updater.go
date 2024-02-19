package updater

import (
	"context"
	"time"

	calculator "github.com/Layr-Labs/eigenlayer-payment-updater/calculator"
	"github.com/Layr-Labs/eigenlayer-payment-updater/common"
	"github.com/Layr-Labs/eigenlayer-payment-updater/common/distribution"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/rs/zerolog/log"
)

type Updater struct {
	UpdateInterval          time.Duration
	paymentsDataService     PaymentsDataService
	distributionDataService DistributionDataService
	calculator              calculator.PaymentCalculator
	transactor              Transactor
}

func NewUpdater(
	updateInterval time.Duration,
	paymentsDataService PaymentsDataService,
	distributionDataService DistributionDataService,
	calculator calculator.PaymentCalculator,
	chainClient *common.ChainClient,
	claimingManagerAddress gethcommon.Address,
) (*Updater, error) {
	transactor, err := NewTransactor(chainClient, claimingManagerAddress)
	if err != nil {
		log.Fatal().Msgf("failed to create transactor: %s", err)
	}

	return &Updater{
		UpdateInterval:          updateInterval,
		distributionDataService: distributionDataService,
		paymentsDataService:     paymentsDataService,
		calculator:              calculator,
		transactor:              transactor,
	}, nil
}

func (u *Updater) Start() error {
	// run a loop unning once every u.UpdateInterval that calls u.update()
	log.Info().Msg("service started")
	ctx := context.Background()

	// run the first update immediately
	if err := u.update(ctx); err != nil {
		log.Error().Msgf("failed to update: %s", err)
	}

	ticker := time.NewTicker(u.UpdateInterval)
	for range ticker.C {
		log.Info().Msg("running update")
		if err := u.update(ctx); err != nil {
			log.Error().Msgf("failed to update: %s", err)
		}
	}
	return nil
}

func (u *Updater) update(ctx context.Context) error {
	// get the interval of time that we need to update payments for
	log.Info().Msg("getting latest finalized timestamp")
	latestFinalizedTimestamp, err := u.transactor.GetLatestFinalizedTimestamp(ctx)
	if err != nil {
		return err
	}

	log.Info().Msgf("latest finalized timestamp: %d", latestFinalizedTimestamp)

	// get the time until which payments have been calculated
	log.Info().Msg("getting payments calculated until timestamp")
	paymentsCalculatedUntilTimestamp, err := u.paymentsDataService.GetPaymentsCalculatedUntilTimestamp(ctx)
	if err != nil {
		return err
	}

	log.Info().Msgf("payments calculated until timestamp: %d", paymentsCalculatedUntilTimestamp)

	// give the interval to the distribution calculator, get the map from address => token => amount
	log.Info().Msg("calculating distribution")
	newPaymentsCalculatedUntilTimestamp, diffDistribution, err := u.calculator.CalculateDistributionUntilTimestamp(ctx, paymentsCalculatedUntilTimestamp, latestFinalizedTimestamp)
	if err != nil {
		return err
	}

	// get the current distribution
	log.Info().Msg("getting current distribution")
	currentDistribution, err := u.distributionDataService.GetDistributionAtTimestamp(paymentsCalculatedUntilTimestamp)
	if err != nil {
		return err
	}

	// add the diff distribution to the current distribution
	log.Info().Msg("adding diff distribution to current distribution")
	newDistribution := currentDistribution
	newDistribution.Add(diffDistribution)

	// merklize the distribution roots
	log.Info().Msg("merklizing distribution roots")
	root, err := newDistribution.Merklize(distribution.SimpleMerklize)
	if err != nil {
		return err
	}

	// set the distribution at the timestamp
	log.Info().Msg("setting distribution")
	if err := u.distributionDataService.SetDistributionAtTimestamp(newPaymentsCalculatedUntilTimestamp, newDistribution); err != nil {
		return err
	}

	// send the merkle root to the smart contract
	log.Info().Msg("updating payments")
	if err := u.transactor.SubmitRoot(ctx, root, newPaymentsCalculatedUntilTimestamp); err != nil {
		return err
	}

	return nil
}
