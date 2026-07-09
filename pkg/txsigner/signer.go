// Package txsigner provides pluggable Ethereum transaction signing for the
// rewards updater. It supports a local private key (for testnets / local dev)
// and AWS KMS (recommended for production, so the key never leaves the HSM).
package txsigner

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// ITransactionSigner is implemented by anything that can sign Ethereum
// transactions for a fixed account. Implementations return a bind.SignerFn
// that go-ethereum's contract bindings invoke when broadcasting a transaction.
type ITransactionSigner interface {
	// GetAddress returns the Ethereum address that signs (the account that must
	// be authorized as the rewardsUpdater on the RewardsCoordinator).
	GetAddress() common.Address

	// SignerFn returns a go-ethereum bind.SignerFn bound to the given chain ID.
	SignerFn(chainID *big.Int) bind.SignerFn
}
