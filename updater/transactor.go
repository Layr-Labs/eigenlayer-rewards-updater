package updater

import (
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
)

type UpdaterTransactor interface {
	SubmitRoot(paymentsCalculatedUntilTimestamp *big.Int, root [32]byte) error
}

type UpdaterTransactorImpl struct {
	chainClient *ChainClient
}

func NewUpdaterTransactor(ethClient *ethclient.Client, privateKeyString string) (UpdaterTransactor, error) {
	chainClient, err := NewChainClient(ethClient, privateKeyString)
	if err != nil {
		return nil, err
	}

	return &UpdaterTransactorImpl{
		chainClient: chainClient,
	}, nil
}

func (u *UpdaterTransactorImpl) SubmitRoot(paymentsCalculatedUntilTimestamp *big.Int, root [32]byte) error {
	return nil
}
