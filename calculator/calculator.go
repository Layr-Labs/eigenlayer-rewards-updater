package calculator

import (
	"math/big"

	"github.com/Layr-Labs/eigenlayer-payment-updater/common"
	gethcommon "github.com/ethereum/go-ethereum/common"
)

type PaymentCalculator interface {
	// CalculateDistribution returns the distribution of given tokens at a given timestamp. 1 distribution per token
	CalculateDistributions(events []*common.PaymentEvent) (map[gethcommon.Address]*common.Distribution, error)
}

type PaymentCalculatorImpl struct {
	dataService PaymentCalculatorDataService
}

func NewPaymentCalculator(dataService PaymentCalculatorDataService) *PaymentCalculatorImpl {
	return &PaymentCalculatorImpl{}
}

func (p *PaymentCalculatorImpl) CalculateDistributions(events []*common.PaymentEvent) (map[gethcommon.Address]*common.Distribution, error) {
	distributions := make(map[gethcommon.Address]*common.Distribution)

	var err error
	for _, event := range events {
		if _, ok := distributions[event.Token]; !ok {
			distributions[event.Token] = common.NewDistribution()
		}

		// get the distribution of ETH at the time of this event
		var total *big.Int
		var distribution *common.Distribution
		if event.Token == common.EIGEN_TOKEN_ADDRESS {
			total, distribution, err = p.dataService.GetClampedIntegratedETHDistributionAtTimestamp(event.SnapshotTimestamp)
			if err != nil {
				return nil, err
			}
		} else {
			total, distribution, err = p.dataService.GetIntegratedETHDistributionAtTimestamp(event.SnapshotTimestamp)
			if err != nil {
				return nil, err
			}
		}

		// calculate the distribution for this event
		distribution.MulDiv(event.Amount, total)

		// add this distribution to the map
		distributions[event.Token].Add(distribution)
	}

	return distributions, nil
}
