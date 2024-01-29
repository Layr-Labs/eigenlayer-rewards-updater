package updater

import "math/big"

type UpdaterTransactor interface {
	SubmitRoot(paymentsCalculatedUntilTimestamp *big.Int, root [32]byte) error
}
