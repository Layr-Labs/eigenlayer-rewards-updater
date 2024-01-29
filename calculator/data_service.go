package calculator

import (
	"math/big"

	contractIPaymentCoordinator "github.com/Layr-Labs/eigenlayer-payment-updater/bindings/IPaymentCoordinator"
	"github.com/Layr-Labs/eigenlayer-payment-updater/common"
	gethcommon "github.com/ethereum/go-ethereum/common"
)

type PaymentCalculatorDataService interface {
	// GetPaymentsCalculatedUntilTimestamp returns the timestamp until which payments have been calculated
	GetPaymentsCalculatedUntilTimestamp() (*big.Int, error)
	// GetRangePaymentsWithOverlappingRange returns all range payments that overlap with the given range
	GetRangePaymentsWithOverlappingRange(startTimestamp, endTimestamp *big.Int) ([]*contractIPaymentCoordinator.IPaymentCoordinatorRangePayment, error)
	// GetDistributionsAtTimestamp returns the distributions of all tokens at a given timestamp
	GetDistributionsAtTimestamp(timestamp *big.Int) (map[gethcommon.Address]*common.Distribution, error)
	// GetOperatorSetAtTimestamp returns the operator set at a given timestamps
	GetOperatorSetAtTimestamp(avs gethcommon.Address, timestamp *big.Int) (common.OperatorSet, error)
}
