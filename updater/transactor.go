package updater

import (
	"context"
	"math/big"

	contractIClaimingManager "github.com/Layr-Labs/eigenlayer-payment-updater/bindings/IClaimingManager"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type UpdaterTransactor interface {
	SubmitRoot(root [32]byte, paymentsCalculatedUntilTimestamp *big.Int) error
}

type UpdaterTransactorImpl struct {
	chainClient     *ChainClient
	claimingManager *contractIClaimingManager.ContractIClaimingManager
}

func NewUpdaterTransactor(ethClient *ethclient.Client, privateKeyString string, claimingManagerAddress gethcommon.Address) (UpdaterTransactor, error) {
	chainClient, err := NewChainClient(ethClient, privateKeyString)
	if err != nil {
		return nil, err
	}

	claimingManager, err := contractIClaimingManager.NewContractIClaimingManager(claimingManagerAddress, chainClient.Client)
	if err != nil {
		return nil, err
	}

	return &UpdaterTransactorImpl{
		chainClient:     chainClient,
		claimingManager: claimingManager,
	}, nil
}

func (t *UpdaterTransactorImpl) SubmitRoot(root [32]byte, paymentsCalculatedUntilTimestamp *big.Int) error {
	tx, err := t.claimingManager.SubmitRoot(t.chainClient.NoSendTransactOpts, root, uint32(paymentsCalculatedUntilTimestamp.Uint64()))
	if err != nil {
		return err
	}

	ctx := context.Background()

	receipt, err := t.chainClient.EstimateGasPriceAndLimitAndSendTx(ctx, tx, "submitRoot")
	if err != nil {
		return err
	}

	if receipt.Status != 1 {
		return ErrTransactionFailed
	}

	return nil
}
