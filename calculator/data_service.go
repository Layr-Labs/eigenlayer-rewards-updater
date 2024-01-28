package calculator

import (
	"math/big"

	"github.com/Layr-Labs/eigenlayer-payment-updater/common"
)

type PaymentCalculatorDataService interface {
	GetIntegratedETHDistributionAtTimestamp(timestamp uint64) (*big.Int, *common.Distribution, error)
	GetClampedIntegratedETHDistributionAtTimestamp(timestamp uint64) (*big.Int, *common.Distribution, error)
}
