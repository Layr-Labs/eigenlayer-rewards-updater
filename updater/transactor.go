package updater

import (
	"context"
	"math/big"

	contractIClaimingManager "github.com/Layr-Labs/eigenlayer-payment-updater/bindings/IClaimingManager"
	"github.com/Layr-Labs/eigenlayer-payment-updater/common"
	gethcommon "github.com/ethereum/go-ethereum/common"
)

type UpdaterTransactor interface {
	SubmitRoot(ctx context.Context, root [32]byte, paymentsCalculatedUntilTimestamp *big.Int) error
}

type UpdaterTransactorImpl struct {
	ChainClient     *common.ChainClient
	ClaimingManager *contractIClaimingManager.ContractIClaimingManager
}

func NewUpdaterTransactor(chainClient *common.ChainClient, claimingManagerAddress gethcommon.Address) (UpdaterTransactor, error) {
	claimingManager, err := contractIClaimingManager.NewContractIClaimingManager(claimingManagerAddress, chainClient.Client)
	if err != nil {
		return nil, err
	}

	return &UpdaterTransactorImpl{
		ChainClient:     chainClient,
		ClaimingManager: claimingManager,
	}, nil
}

func (t *UpdaterTransactorImpl) SubmitRoot(ctx context.Context, root [32]byte, paymentsCalculatedUntilTimestamp *big.Int) error {
	tx, err := t.ClaimingManager.SubmitRoot(t.ChainClient.NoSendTransactOpts, root, uint32(paymentsCalculatedUntilTimestamp.Uint64()))
	if err != nil {
		return err
	}

	receipt, err := t.ChainClient.EstimateGasPriceAndLimitAndSendTx(ctx, tx, "submitRoot")
	if err != nil {
		return err
	}

	if receipt.Status != 1 {
		return common.ErrTransactionFailed
	}

	return nil
}
