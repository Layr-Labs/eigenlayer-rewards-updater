package updater

import (
	"time"

	calculator "github.com/Layr-Labs/eigenlayer-payment-updater/calculator"
	"github.com/Layr-Labs/eigenlayer-payment-updater/common"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/rs/zerolog/log"
)

type Updater struct {
	UpdateInterval time.Duration
	dataService    PaymentDataService
	calculator     calculator.PaymentCalculator
	transactor     UpdaterTransactor
}

func NewUpdater(updateInterval time.Duration, dataService PaymentDataService, calculator calculator.PaymentCalculator, transactor UpdaterTransactor) *Updater {
	return &Updater{
		UpdateInterval: updateInterval,
		dataService:    dataService,
		calculator:     calculator,
		transactor:     transactor,
	}
}

func (u *Updater) Start() error {
	// run a loop unning once every u.UpdateInterval that calls u.update()
	log.Info().Msg("service started")

	ticker := time.NewTicker(u.UpdateInterval)
	for range ticker.C {
		log.Info().Msg("running update")
		if err := u.update(); err != nil {
			log.Error().Msgf("failed to update: %s", err)
		}
	}
	return nil
}

func (u *Updater) update() error {
	// get all events since the last update
	log.Info().Msg("getting events since last update")
	paymentsUpdatedUntilTimestamp, newPaymentEvents, err := u.dataService.GetEventsSinceLastUpdate()
	if err != nil {
		return err
	}

	// feed these events into the distribution calculator, get the map from address => token => amount
	log.Info().Msg("calculating distribution")
	newDistributions, err := u.calculator.CalculateDistributions(newPaymentEvents)
	if err != nil {
		return err
	}

	// get all tokens
	tokens := make([]gethcommon.Address, len(newDistributions))
	i := 0
	for token := range newDistributions {
		tokens[i] = token
		i++
	}

	// get previous cumulative distribution of given tokens
	log.Info().Msg("getting previous distribution")
	previousDistributions, err := u.dataService.GetDistributionsOfTokensAtTimestamp(paymentsUpdatedUntilTimestamp, tokens)
	if err != nil {
		return err
	}

	// add the pending distribution to the previous distribution
	log.Info().Msg("adding pending distribution to previous distribution and merklizing")
	distributionRoots := make([][]byte, len(tokens))
	for token, distribution := range previousDistributions {
		newDistributions[token].Add(distribution)

		distributionRoot, err := newDistributions[token].Merklize()
		if err != nil {
			return err
		}

		distributionRoots = append(distributionRoots, distributionRoot[:])
	}

	// merklize the distribution roots
	log.Info().Msg("merklizing distribution roots")
	root, err := common.Merklize(distributionRoots)

	// send the merkle root to the smart contract
	log.Info().Msg("updating payments")
	if err := u.transactor.SubmitRoot(paymentsUpdatedUntilTimestamp, root); err != nil {
		return err
	}

	return nil
}
