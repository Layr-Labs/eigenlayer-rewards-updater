package transactor

import (
	"context"
	"fmt"
	paymentCoordinator "github.com/Layr-Labs/eigenlayer-contracts/pkg/bindings/IPaymentCoordinator"
	"github.com/Layr-Labs/eigenlayer-payment-updater/pkg/chainClient"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethcommon "github.com/ethereum/go-ethereum/common"
)

type Transactor interface {
	CurrPaymentCalculationEndTimestamp() (uint32, error)
	GetRootIndex(root [32]byte) (uint32, error)
	SubmitRoot(ctx context.Context, root [32]byte, paymentsUnixTimestamp uint32) error
}

type TransactorImpl struct {
	ChainClient                  *chainClient.ChainClient
	PaymentCoordinatorCaller     *paymentCoordinator.IPaymentCoordinatorCaller
	PaymentCoordinatorTransactor *paymentCoordinator.IPaymentCoordinatorTransactor
}

func NewTransactor(chainClient *chainClient.ChainClient, paymentCoordinatorAddress gethcommon.Address) (Transactor, error) {
	paymentCoordinatorCaller, err := paymentCoordinator.NewIPaymentCoordinatorCaller(paymentCoordinatorAddress, chainClient.Client)
	if err != nil {
		return nil, err
	}

	paymentCoordinatorTransactor, err := paymentCoordinator.NewIPaymentCoordinatorTransactor(paymentCoordinatorAddress, chainClient.Client)
	if err != nil {
		return nil, err
	}

	return &TransactorImpl{
		ChainClient:                  chainClient,
		PaymentCoordinatorCaller:     paymentCoordinatorCaller,
		PaymentCoordinatorTransactor: paymentCoordinatorTransactor,
	}, nil
}

func (t *TransactorImpl) CurrPaymentCalculationEndTimestamp() (uint32, error) {
	return t.PaymentCoordinatorCaller.CurrPaymentCalculationEndTimestamp(&bind.CallOpts{})
}

func (s *TransactorImpl) GetRootIndex(root [32]byte) (uint32, error) {
	return s.PaymentCoordinatorCaller.GetRootIndexFromHash(&bind.CallOpts{}, root)
}

func (t *TransactorImpl) SubmitRoot(ctx context.Context, root [32]byte, paymentsUnixTimestamp uint32) error {
	// todo: params
	tx, err := t.PaymentCoordinatorTransactor.SubmitRoot(t.ChainClient.NoSendTransactOpts, root, paymentsUnixTimestamp)
	if err != nil {
		fmt.Printf("Payment coordinator, failed to submit root: %+v - %+v\n", err, tx)
		return err
	}

	receipt, err := t.ChainClient.EstimateGasPriceAndLimitAndSendTx(ctx, tx, "submitRoot")
	if err != nil {
		fmt.Printf("Failed to estimate gas: %+v\n", err)
		return err
	}

	if receipt.Status != 1 {
		return chainClient.ErrTransactionFailed
	}

	return nil
}
