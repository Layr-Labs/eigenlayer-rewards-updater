package services

import (
	"context"
	"math/big"

	paymentCoordinator "github.com/Layr-Labs/eigenlayer-payment-updater/bindings/IPaymentCoordinator"
	"github.com/Layr-Labs/eigenlayer-payment-updater/common"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethcommon "github.com/ethereum/go-ethereum/common"
)

const FINALIZATION_DEPTH = 100

type Transactor interface {
	GetRootIndex(root [32]byte) (uint32, error)
	SubmitRoot(ctx context.Context, root [32]byte, paymentsCalculatedUntilTimestamp *big.Int) error
}

type TransactorImpl struct {
	ChainClient        *common.ChainClient
	PaymentCoordinator *paymentCoordinator.ContractIPaymentCoordinator
}

func NewTransactor(chainClient *common.ChainClient, paymentCoordinatorAddress gethcommon.Address) (Transactor, error) {
	paymentCoordinatorContract, err := paymentCoordinator.NewContractIPaymentCoordinator(paymentCoordinatorAddress, chainClient.Client)
	if err != nil {
		return nil, err
	}

	return &TransactorImpl{
		ChainClient:        chainClient,
		PaymentCoordinator: paymentCoordinatorContract,
	}, nil
}

func (s *TransactorImpl) GetRootIndex(root [32]byte) (uint32, error) {
	return s.PaymentCoordinator.GetRootIndex(&bind.CallOpts{}, root)
}

func (t *TransactorImpl) SubmitRoot(ctx context.Context, root [32]byte, paymentsCalculatedUntilTimestamp *big.Int) error {
	// todo: params
	tx, err := t.PaymentCoordinator.SubmitRoot(t.ChainClient.NoSendTransactOpts, root, paymentsCalculatedUntilTimestamp.Uint64(), 0)
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
