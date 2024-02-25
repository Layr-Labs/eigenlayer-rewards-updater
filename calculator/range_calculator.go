package calculator

import (
	"context"
	"fmt"
	"math/big"

	contractIPaymentCoordinator "github.com/Layr-Labs/eigenlayer-payment-updater/bindings/IPaymentCoordinator"
	"github.com/Layr-Labs/eigenlayer-payment-updater/common/distribution"
	"github.com/Layr-Labs/eigenlayer-payment-updater/common/services"
	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog/log"
)

type RangePaymentCalculator struct {
	calculationIntervalSeconds *big.Int
	paymentsDataService        services.PaymentsDataService
	operatorSetDataService     OperatorSetDataService
}

func NewRangePaymentCalculator(
	calculationIntervalSeconds int64,
	paymentsDataService services.PaymentsDataService,
	operatorSetDataService OperatorSetDataService,
) PaymentCalculator {
	return &RangePaymentCalculator{
		calculationIntervalSeconds: big.NewInt(calculationIntervalSeconds),
		paymentsDataService:        paymentsDataService,
		operatorSetDataService:     operatorSetDataService,
	}
}

func (c *RangePaymentCalculator) CalculateDistributionUntilTimestamp(ctx context.Context, startTimestamp, endTimestamp *big.Int) (*big.Int, *distribution.Distribution, error) {
	// make sure the start timestamp is rounded to the nearest interval granularity
	if new(big.Int).Mod(startTimestamp, c.calculationIntervalSeconds).Cmp(big.NewInt(0)) != 0 {
		return nil, nil, fmt.Errorf("start timestamp must be rounded to the nearest interval granularity")
	}

	// round the end timestamp to the nearest interval granularity. the start is assumed to be rounded already
	endTimestamp.Sub(endTimestamp, new(big.Int).Mod(endTimestamp, c.calculationIntervalSeconds))

	// make sure the end timestamp is after the start timestamp
	if endTimestamp.Cmp(startTimestamp) <= 0 {
		return nil, nil, fmt.Errorf("end timestamp must be after start timestamp")
	}

	// todo remove
	// clamp the end timestamp to 2 intervals ahead of the start timestamp
	if endTimestamp.Cmp(new(big.Int).Add(startTimestamp, new(big.Int).Mul(c.calculationIntervalSeconds, big.NewInt(1)))) >= 0 {
		endTimestamp = new(big.Int).Add(startTimestamp, new(big.Int).Mul(c.calculationIntervalSeconds, big.NewInt(1)))
	}

	log.Info().Msgf("calculating distributions from %d to %d", startTimestamp, endTimestamp)

	// get all range payments that overlap with the given interval that are coninuing, meaning they were created before the start timestamp
	// calculation for them has already been done up until startTimestamp
	continuingRangePayments, err := c.paymentsDataService.GetRangePaymentsWithOverlappingRange(startTimestamp, endTimestamp, big.NewInt(0), startTimestamp)
	if err != nil && err != pgx.ErrNoRows {
		return nil, nil, err
	}
	if err == pgx.ErrNoRows {
		log.Info().Msg("no range payments found")
		// should we return the distribution at the end timestamp? or just "skip" somehow?
	}

	log.Info().Msgf("found %d coninuing range payments", len(continuingRangePayments))

	// calculate the distribution over the range
	diffDistribution := distribution.NewDistribution()
	err = c.CalculateDistributionFromRangePayments(ctx, startTimestamp, endTimestamp, continuingRangePayments, diffDistribution)
	if err != nil {
		return nil, nil, err
	}

	// get all range payments that overlap with the given interval that were created after the start timestamp
	// calculation for them has not been done yet until startTimestamp
	newRangePayments, err := c.paymentsDataService.GetRangePaymentsWithOverlappingRange(startTimestamp, endTimestamp, startTimestamp, endTimestamp)
	if err != nil && err != pgx.ErrNoRows {
		return nil, nil, err
	}

	log.Info().Msgf("found %d new range payments", len(newRangePayments))

	// the range for calculaton for new range payments should be the minimum start timestamp of the new range payments
	// and the end timestamp

	startTimestampNewRangePayments := new(big.Int).Set(startTimestamp)
	for _, rangePayment := range newRangePayments {
		startTimestampNewRangePayments = min(startTimestampNewRangePayments, rangePayment.StartRangeTimestamp)
	}

	// calculate the distribution over the range
	err = c.CalculateDistributionFromRangePayments(ctx, startTimestampNewRangePayments, endTimestamp, newRangePayments, diffDistribution)
	if err != nil {
		return nil, nil, err
	}

	return endTimestamp, diffDistribution, nil
}

// CalculateDistributionFromRangePayments calculates the distribution of payments over the given range for the given range payments.
// it overwrites the diffDistribution with the new updated values
// it assumes that startTimestamp < endTimestamp and they are at an interval boundary
func (c *RangePaymentCalculator) CalculateDistributionFromRangePayments(
	ctx context.Context,
	startTimestamp, endTimestamp *big.Int,
	rangePayments []*contractIPaymentCoordinator.IPaymentCoordinatorRangePayment,
	diffDistribution *distribution.Distribution,
) error {
	numIntervals := new(big.Int).Div(new(big.Int).Sub(endTimestamp, startTimestamp), c.calculationIntervalSeconds).Int64()

	// loop through all range payments
	for _, rangePayment := range rangePayments {
		// calculate the payment to distribute over each interval
		paymentToDistribute := div(mul(rangePayment.Amount, c.calculationIntervalSeconds), new(big.Int).Sub(rangePayment.EndRangeTimestamp, rangePayment.StartRangeTimestamp))

		intervalStart := new(big.Int).Set(rangePayment.StartRangeTimestamp)
		intervalEnd := new(big.Int).Add(intervalStart, c.calculationIntervalSeconds)

		// loop through all intervals
		for i := int64(0); i < numIntervals; i++ {
			// if the interval start is at or after the end of the range payment's range, move to the next range payment
			if intervalStart.Cmp(rangePayment.EndRangeTimestamp) >= 0 {
				break
			}

			// get the operator set at the interval start
			operatorSet, err := c.operatorSetDataService.GetOperatorSetForStrategyAtTimestamp(intervalStart, rangePayment.Avs, rangePayment.Strategy)
			if err != nil {
				return err
			}

			// if the operator set has no staked strategy shares, skip
			if operatorSet.TotalStakedStrategyShares.Cmp(big.NewInt(0)) == 0 {
				continue
			}

			// loop through all operators
			for _, operator := range operatorSet.Operators {
				// totalPaymentToOperatorAndStakers = paymentToDistribute * operatorDelegatedStrategyShares / totalStrategyShares
				totalPaymentToOperatorAndStakers := div(mul(paymentToDistribute, operator.TotalDelegatedStrategyShares), operatorSet.TotalStakedStrategyShares)
				log.Info().Msgf("total payment to operator and stakers: %s", totalPaymentToOperatorAndStakers)

				// if the operator has no delegated strategy shares, skip
				if operator.TotalDelegatedStrategyShares.Cmp(big.NewInt(0)) == 0 {
					continue
				}

				// increment token balance according to the operator's commission
				operatorAmt := diffDistribution.Get(operator.Address, rangePayment.Token)

				// operatorBalance += totalPaymentToOperatorAndStakers * operatorCommissions / 10000
				diffDistribution.Set(
					operator.Address,
					rangePayment.Token,
					operatorAmt.Add(operatorAmt, div(mul(totalPaymentToOperatorAndStakers, operator.Commission), BIPS_DENOMINATOR)),
				)

				// loop through all stakers
				for _, staker := range operator.Stakers {
					// increment token balance according to the staker's proportion of the strategy shares
					stakerAmt := diffDistribution.Get(staker.Address, rangePayment.Token)

					// stakerBalance += totalPaymentToOperatorAndStakers * (10000 - operatorCommissions) * stakerShares / 10000 / operatorDelegatedStrategyShares
					diffDistribution.Set(
						staker.Address,
						rangePayment.Token,
						stakerAmt.Add(
							stakerAmt,
							div(
								mul(totalPaymentToOperatorAndStakers, new(big.Int).Sub(BIPS_DENOMINATOR, operator.Commission), staker.StrategyShares),
								BIPS_DENOMINATOR, operator.TotalDelegatedStrategyShares,
							),
						),
					)
				}

				// increment the interval start/end
				intervalStart.Add(intervalStart, c.calculationIntervalSeconds)
				intervalEnd.Add(intervalEnd, c.calculationIntervalSeconds)
			}
		}
	}

	return nil
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

// p_0 * p_0 * ...
func mul(ps ...*big.Int) *big.Int {
	res := big.NewInt(1)
	for _, p := range ps {
		res.Mul(res, p)
	}

	return res
}

// p / d_0 / d_1 / ...
func div(p *big.Int, ds ...*big.Int) *big.Int {
	res := new(big.Int).Set(p)
	for _, d := range ds {
		res.Div(res, d)
	}

	return res
}
