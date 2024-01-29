package updater

import (
	"math/big"
)

type PaymentDataService interface {
	GetLatestFinalizedTimestamp() (*big.Int, error)
}
