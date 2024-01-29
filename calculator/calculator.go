package calculator

import (
	"context"
	"math/big"

	"github.com/Layr-Labs/eigenlayer-payment-updater/common"
	gethcommon "github.com/ethereum/go-ethereum/common"
)

var BIPS_DENOMINATOR = big.NewInt(10000)

type PaymentCalculator interface {
	// CalculateDistributionsOverInterval returns the distributions of given tokens until the given end timestamp
	CalculateDistributionsUntilTimestamp(ctx context.Context, endTimestamp *big.Int) (*big.Int, map[gethcommon.Address]*common.Distribution, error)
}
