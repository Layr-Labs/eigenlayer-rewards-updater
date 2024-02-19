package updater

import (
	"context"
	"math/big"
	"time"

	calculator "github.com/Layr-Labs/eigenlayer-payment-updater/calculator"
	"github.com/Layr-Labs/eigenlayer-payment-updater/common"
	"github.com/Layr-Labs/eigenlayer-payment-updater/common/distribution"
	"github.com/Layr-Labs/eigenlayer-payment-updater/common/services"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog/log"
)

type Updater struct {
	UpdateInterval          time.Duration
	paymentsDataService     services.PaymentsDataService
	distributionDataService DistributionDataService
	calculator              calculator.PaymentCalculator
	transactor              Transactor
}

func NewUpdater(
	updateIntervalSeconds int,
	paymentsDataService services.PaymentsDataService,
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
		UpdateInterval:          time.Second * time.Duration(updateIntervalSeconds),
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
	latestRoot, paymentsCalculatedUntilTimestamp, fetchRootSubmissionErr := u.paymentsDataService.GetLatestRootSubmission(ctx)
	if fetchRootSubmissionErr == pgx.ErrNoRows {
		// if there are no rows, then we haven't submitted any roots yet, so we should start from 0
		paymentsCalculatedUntilTimestamp = big.NewInt(0)
	} else if fetchRootSubmissionErr != nil {
		return fetchRootSubmissionErr
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
	currentDistribution := distribution.NewDistribution()
	if paymentsCalculatedUntilTimestamp.Cmp(big.NewInt(0)) != 0 {
		currentDistribution, err = u.distributionDataService.GetDistribution(latestRoot)
		if err != nil {
			return err
		}
	}

	// add the diff distribution to the current distribution
	log.Info().Msg("adding diff distribution to current distribution")
	newDistribution := currentDistribution
	newDistribution.Add(diffDistribution)

	// merklize the distribution roots
	log.Info().Msg("merklizing distribution roots")
	newRoot, err := newDistribution.Merklize(distribution.SimpleMerklize)
	if err != nil {
		return err
	}

	// set the distribution at the timestamp
	log.Info().Msg("setting distribution")
	if err := u.distributionDataService.SetDistribution(newRoot, newDistribution); err != nil {
		return err
	}

	// send the merkle root to the smart contract
	log.Info().Msg("updating payments")
	if err := u.transactor.SubmitRoot(ctx, newRoot, newPaymentsCalculatedUntilTimestamp); err != nil {
		return err
	}

	return nil
}
