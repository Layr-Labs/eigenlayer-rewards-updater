package common

import (
	"math/big"

	gethcommon "github.com/ethereum/go-ethereum/common"
)

type PaymentEvent struct {
	Token             gethcommon.Address
	Amount            *big.Int
	SnapshotTimestamp uint64
}
