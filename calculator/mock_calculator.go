package calculator

import (
	"context"
	"math/big"

	"github.com/Layr-Labs/eigenlayer-payment-updater/common"
	gethcommon "github.com/ethereum/go-ethereum/common"
)

type MockCalculator struct {
}

func NewMockCalculator() PaymentCalculator {
	return &MockCalculator{}
}

func (c *MockCalculator) CalculateDistributionsUntilTimestamp(ctx context.Context, endTimestamp *big.Int) (*big.Int, map[gethcommon.Address]*common.Distribution, error) {
	distributions := make(map[gethcommon.Address]*common.Distribution)
	token1 := gethcommon.HexToAddress("0x0000000000000000000000000000000000000000")
	token2 := gethcommon.HexToAddress("0x0000000000000000000000000000000000000001")
	token3 := gethcommon.HexToAddress("0x0000000000000000000000000000000000000002")

	distributions[token1] = common.NewRandomDistribution(100)
	distributions[token2] = common.NewRandomDistribution(200)
	distributions[token3] = common.NewRandomDistribution(300)

	return endTimestamp, distributions, nil
}
