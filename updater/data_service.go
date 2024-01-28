package updater

import (
	"github.com/Layr-Labs/eigenlayer-payment-updater/common"
	gethcommon "github.com/ethereum/go-ethereum/common"
)

type PaymentDataService interface {
	// GetEventsSinceLastUpdate returns all events that have occurred since the last update, along with the timestamp of the last event taken into account
	GetEventsSinceLastUpdate() (uint64, []*common.PaymentEvent, error)

	// GetDistributionOfTokensAtTimestamp returns the distribution of given tokens at a given timestamp
	GetDistributionsOfTokensAtTimestamp(timestamp uint64, tokens []gethcommon.Address) (map[gethcommon.Address]*common.Distribution, error)
}
