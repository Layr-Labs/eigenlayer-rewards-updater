package updater

import (
	"context"
	"math/big"

	contractIClaimingManager "github.com/Layr-Labs/eigenlayer-payment-updater/bindings/IClaimingManager"
	"github.com/Layr-Labs/eigenlayer-payment-updater/common"
	gethcommon "github.com/ethereum/go-ethereum/common"
)

const FINALIZATION_DEPTH = 100

type Transactor interface {
	GetLatestFinalizedTimestamp(ctx context.Context) (*big.Int, error)
	SubmitRoot(ctx context.Context, root [32]byte, paymentsCalculatedUntilTimestamp *big.Int) error
}

type TransactorImpl struct {
	ChainClient     *common.ChainClient
	ClaimingManager *contractIClaimingManager.ContractIClaimingManager
}

func NewTransactor(chainClient *common.ChainClient, claimingManagerAddress gethcommon.Address) (Transactor, error) {
	claimingManager, err := contractIClaimingManager.NewContractIClaimingManager(claimingManagerAddress, chainClient.Client)
	if err != nil {
		return nil, err
	}

	return &TransactorImpl{
		ChainClient:     chainClient,
		ClaimingManager: claimingManager,
	}, nil
}

func (s *TransactorImpl) GetLatestFinalizedTimestamp(ctx context.Context) (*big.Int, error) {
	latestBlockNumber, err := s.ChainClient.GetCurrentBlockNumber(ctx)
	if err != nil {
		return nil, err
	}

	latestFinalizedBlockNumber := latestBlockNumber - FINALIZATION_DEPTH

	latestFinalizedBlock, err := s.ChainClient.HeaderByNumber(ctx, big.NewInt(int64(latestFinalizedBlockNumber)))
	if err != nil {
		return nil, err
	}

	return big.NewInt(int64(latestFinalizedBlock.Time)), nil
}

func (t *TransactorImpl) SubmitRoot(ctx context.Context, root [32]byte, paymentsCalculatedUntilTimestamp *big.Int) error {
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
