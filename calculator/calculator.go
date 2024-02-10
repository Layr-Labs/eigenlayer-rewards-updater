package calculator

import (
	"context"
	"math/big"

	"github.com/Layr-Labs/eigenlayer-payment-updater/common/distribution"
)

var BIPS_DENOMINATOR = big.NewInt(10000)

type PaymentCalculator interface {
	// CalculateDistributionUntilTimestamp returns the distribution of given tokens until the given end timestamp
	CalculateDistributionUntilTimestamp(ctx context.Context, endTimestamp *big.Int) (*big.Int, *distribution.Distribution, error)
}
