package services

import (
	"context"
	"github.com/Layr-Labs/eigenlayer-payment-proofs/pkg/paymentCoordinator"
	"github.com/Layr-Labs/eigenlayer-payment-updater/pkg"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethcommon "github.com/ethereum/go-ethereum/common"
)

const FINALIZATION_DEPTH = 100

type Transactor interface {
	CurrPaymentCalculationEndTimestamp() (uint64, error)
	GetRootIndex(root [32]byte) (uint32, error)
	SubmitRoot(ctx context.Context, root [32]byte, paymentsCalculatedUntilTimestamp *big.Int) error
}

type TransactorImpl struct {
	ChainClient        *pkg.ChainClient
	PaymentCoordinator *paymentCoordinator.ContractIPaymentCoordinator
}

func NewTransactor(chainClient *pkg.ChainClient, paymentCoordinatorAddress gethcommon.Address) (Transactor, error) {
	paymentCoordinatorContract, err := paymentCoordinator.NewContractIPaymentCoordinator(paymentCoordinatorAddress, chainClient.Client)
	if err != nil {
		return nil, err
	}

	return &TransactorImpl{
		ChainClient:        chainClient,
		PaymentCoordinator: paymentCoordinatorContract,
	}, nil
}

func (t *TransactorImpl) CurrPaymentCalculationEndTimestamp() (uint64, error) {
	return t.PaymentCoordinator.CurrPaymentCalculationEndTimestamp(&bind.CallOpts{})
}

func (s *TransactorImpl) GetRootIndex(root [32]byte) (uint32, error) {
	return s.PaymentCoordinator.GetRootIndexFromHash(&bind.CallOpts{}, root)
}

func (t *TransactorImpl) SubmitRoot(ctx context.Context, root [32]byte, paymentsCalculatedUntilTimestamp *big.Int) error {
	// todo: params
	tx, err := t.PaymentCoordinator.SubmitRoot(t.ChainClient.NoSendTransactOpts, root, paymentsCalculatedUntilTimestamp.Uint64())
	if err != nil {
		return err
	}

	receipt, err := t.ChainClient.EstimateGasPriceAndLimitAndSendTx(ctx, tx, "submitRoot")
	if err != nil {
		return err
	}

	if receipt.Status != 1 {
		return pkg.ErrTransactionFailed
	}

	return nil
}
