package updater

import (
	"context"
	"math/big"

	contractIClaimingManager "github.com/Layr-Labs/eigenlayer-payment-updater/bindings/IClaimingManager"
	gethcommon "github.com/ethereum/go-ethereum/common"
)

type UpdaterTransactor interface {
	SubmitRoot(root [32]byte, paymentsCalculatedUntilTimestamp *big.Int) error
}

type UpdaterTransactorImpl struct {
	ChainClient     *ChainClient
	ClaimingManager *contractIClaimingManager.ContractIClaimingManager
}

func NewUpdaterTransactor(chainClient *ChainClient, claimingManagerAddress gethcommon.Address) (UpdaterTransactor, error) {
	claimingManager, err := contractIClaimingManager.NewContractIClaimingManager(claimingManagerAddress, chainClient.Client)
	if err != nil {
		return nil, err
	}

	return &UpdaterTransactorImpl{
		ChainClient:     chainClient,
		ClaimingManager: claimingManager,
	}, nil
}

func (t *UpdaterTransactorImpl) SubmitRoot(root [32]byte, paymentsCalculatedUntilTimestamp *big.Int) error {
	tx, err := t.ClaimingManager.SubmitRoot(t.ChainClient.NoSendTransactOpts, root, uint32(paymentsCalculatedUntilTimestamp.Uint64()))
	if err != nil {
		return err
	}

	ctx := context.Background()

	receipt, err := t.ChainClient.EstimateGasPriceAndLimitAndSendTx(ctx, tx, "submitRoot")
	if err != nil {
		return err
	}

	if receipt.Status != 1 {
		return ErrTransactionFailed
	}

	return nil
}
