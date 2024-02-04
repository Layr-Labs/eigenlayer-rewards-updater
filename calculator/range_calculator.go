package calculator

import (
	"context"
	"fmt"
	"math/big"

	"github.com/Layr-Labs/eigenlayer-payment-updater/common"
	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog/log"
)

type RangePaymentCalculator struct {
	intervalSecondsLength *big.Int
	dataService           PaymentCalculatorDataService
}

func NewRangePaymentCalculator(intervalSecondsLength *big.Int, dataService PaymentCalculatorDataService) PaymentCalculator {
	return &RangePaymentCalculator{
		intervalSecondsLength: intervalSecondsLength,
		dataService:           dataService,
	}
}

func (c *RangePaymentCalculator) CalculateDistributionUntilTimestamp(ctx context.Context, endTimestamp *big.Int) (*big.Int, *common.Distribution, error) {
	startTimestamp, err := c.dataService.GetPaymentsCalculatedUntilTimestamp(ctx)
	if err != nil && err != pgx.ErrNoRows {
		return nil, nil, err
	}
	if err == pgx.ErrNoRows {
		// TODO: this correctly
		// set timestamp to 1 interval behind end timestamp
		startTimestamp = new(big.Int).Sub(endTimestamp, c.intervalSecondsLength)
		startTimestamp.Sub(startTimestamp, new(big.Int).Mod(startTimestamp, c.intervalSecondsLength))
	}

	// make sure the start timestamp is rounded to the nearest interval granularity
	if new(big.Int).Mod(startTimestamp, c.intervalSecondsLength).Cmp(big.NewInt(0)) != 0 {
		return nil, nil, fmt.Errorf("start timestamp must be rounded to the nearest interval granularity")
	}

	// round the end timestamp to the nearest interval granularity. the start is assumed to be rounded already
	endTimestamp.Sub(endTimestamp, new(big.Int).Mod(endTimestamp, c.intervalSecondsLength))

	// make sure the end timestamp is after the start timestamp
	if endTimestamp.Cmp(startTimestamp) <= 0 {
		return nil, nil, fmt.Errorf("end timestamp must be after start timestamp")
	}

	// todo remove
	// clamp the end timestamp to 2 intervals ahead of the start timestamp
	if endTimestamp.Cmp(new(big.Int).Add(startTimestamp, new(big.Int).Mul(c.intervalSecondsLength, big.NewInt(1)))) >= 0 {
		endTimestamp = new(big.Int).Add(startTimestamp, new(big.Int).Mul(c.intervalSecondsLength, big.NewInt(1)))
	}

	log.Info().Msgf("calculating distributions from %d to %d", startTimestamp, endTimestamp)

	// get distribution at the start timestamp
	distribution, err := c.dataService.GetDistributionAtTimestamp(startTimestamp)
	if err != nil {
		return nil, nil, err
	}

	numIntervals := new(big.Int).Div(new(big.Int).Sub(endTimestamp, startTimestamp), c.intervalSecondsLength).Int64()

	// get all range payments that overlap with the given interval
	rangePayments, err := c.dataService.GetRangePaymentsWithOverlappingRange(startTimestamp, endTimestamp)
	if err != nil && err != pgx.ErrNoRows {
		return nil, nil, err
	}
	if err == pgx.ErrNoRows {
		log.Info().Msg("no range payments found")
		return endTimestamp, distribution, nil
	}

	log.Info().Msgf("found %d range payments", len(rangePayments))

	// loop through all range payments
	for _, rangePayment := range rangePayments {
		// calculate the payment per second
		paymentPerSecond := new(big.Int).Div(rangePayment.Amount, new(big.Int).Sub(rangePayment.EndRangeTimestamp, rangePayment.StartRangeTimestamp))

		// loop through all intervals
		for i := int64(0); i < numIntervals; i++ {
			// calculate the start and end of the interval
			intervalStart := new(big.Int).Add(startTimestamp, new(big.Int).Mul(c.intervalSecondsLength, big.NewInt(i)))
			intervalEnd := new(big.Int).Add(startTimestamp, new(big.Int).Mul(c.intervalSecondsLength, big.NewInt(i+1)))

			// calculate overlap between the interval and the range payment
			overlapStart := max(intervalStart, rangePayment.StartRangeTimestamp)
			overlapEnd := min(intervalEnd, rangePayment.EndRangeTimestamp)

			// calculate the payment to distribute
			paymentToDistribute := new(big.Int).Mul(paymentPerSecond, new(big.Int).Sub(overlapEnd, overlapStart))

			// get the operator set at the interval start
			operatorSet, err := c.dataService.GetOperatorSetForStrategyAtTimestamp(intervalStart, rangePayment.Avs, rangePayment.Strategy)
			if err != nil {
				return nil, nil, err
			}

			// if the operator set has no staked strategy shares, skip
			if operatorSet.TotalStakedStrategyShares.Cmp(big.NewInt(0)) == 0 {
				continue
			}

			// loop through all operators
			for _, operator := range operatorSet.Operators {
				// totalPaymentToOperatorAndStakers = paymentToDistribute * operatorDelegatedStrategyShares / totalStrategyShares
				totalPaymentToOperatorAndStakers := new(big.Int).Div(new(big.Int).Mul(paymentToDistribute, operator.TotalDelegatedStrategyShares), operatorSet.TotalStakedStrategyShares)
				log.Info().Msgf("total payment to operator and stakers: %s", totalPaymentToOperatorAndStakers)

				// if the operator has no delegated strategy shares, skip
				if operator.TotalDelegatedStrategyShares.Cmp(big.NewInt(0)) == 0 {
					continue
				}

				// increment token balance according to the operator's commission
				operatorAmt := distribution.Get(operator.Address, rangePayment.Token)

				// operatorBalance += totalPaymentToOperatorAndStakers * operatorCommissions / 10000
				distribution.Set(operator.Address, rangePayment.Token, operatorAmt.Add(operatorAmt, new(big.Int).Div(new(big.Int).Mul(totalPaymentToOperatorAndStakers, operator.Commission), BIPS_DENOMINATOR)))

				// loop through all stakers
				for _, staker := range operator.Stakers {
					// increment token balance according to the staker's proportion of the strategy shares
					stakerAmt := distribution.Get(staker.Address, rangePayment.Token)

					// stakerBalance += totalPaymentToOperatorAndStakers * (1 - operatorCommissions) * stakerShares / 10000 / operatorDelegatedStrategyShares
					distribution.Set(staker.Address, rangePayment.Token, stakerAmt.Add(stakerAmt, new(big.Int).Div(new(big.Int).Div(new(big.Int).Mul(new(big.Int).Mul(totalPaymentToOperatorAndStakers, new(big.Int).Sub(BIPS_DENOMINATOR, operator.Commission)), staker.StrategyShares), BIPS_DENOMINATOR), operator.TotalDelegatedStrategyShares)))
				}
			}
		}
	}

	// set the distributions at the end timestamp
	err = c.dataService.SetDistributionAtTimestamp(endTimestamp, distribution)
	if err != nil {
		return nil, nil, err
	}

	return endTimestamp, distribution, nil
}

func max(a, b *big.Int) *big.Int {
	if a.Cmp(b) == 1 {
		return a
	}
	return b
}

func min(a, b *big.Int) *big.Int {
	if a.Cmp(b) == -1 {
		return a
	}
	return b
}
