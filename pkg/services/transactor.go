package services

import (
	"context"
	"fmt"
	rewardsCoordinator "github.com/Layr-Labs/eigenlayer-contracts/pkg/bindings/IRewardsCoordinator"
	"github.com/Layr-Labs/eigenlayer-rewards-updater/pkg/chainClient"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"math/big"
)

type Transactor interface {
	CurrRewardsCalculationEndTimestamp() (uint32, error)
	GetNumberOfPublishedRoots() (*big.Int, error)
	GetRootIndex(root [32]byte) (uint32, error)
	SubmitRoot(ctx context.Context, root [32]byte, rewardsUnixTimestamp uint32) error
}

type TransactorImpl struct {
	ChainClient           *chainClient.ChainClient
	CoordinatorCaller     *rewardsCoordinator.IRewardsCoordinatorCaller
	CoordinatorTransactor *rewardsCoordinator.IRewardsCoordinatorTransactor
}

func NewTransactor(chainClient *chainClient.ChainClient, coordinatorAddress gethcommon.Address) (Transactor, error) {
	coordinatorCaller, err := rewardsCoordinator.NewIRewardsCoordinatorCaller(coordinatorAddress, chainClient.Client)
	if err != nil {
		return nil, err
	}

	coordinatorTransactor, err := rewardsCoordinator.NewIRewardsCoordinatorTransactor(coordinatorAddress, chainClient.Client)
	if err != nil {
		return nil, err
	}

	return &TransactorImpl{
		ChainClient:           chainClient,
		CoordinatorCaller:     coordinatorCaller,
		CoordinatorTransactor: coordinatorTransactor,
	}, nil
}

func (t *TransactorImpl) CurrRewardsCalculationEndTimestamp() (uint32, error) {
	return t.CoordinatorCaller.CurrRewardsCalculationEndTimestamp(&bind.CallOpts{})
}

func (t *TransactorImpl) GetNumberOfPublishedRoots() (*big.Int, error) {
	return t.CoordinatorCaller.GetDistributionRootsLength(&bind.CallOpts{})
}

func (s *TransactorImpl) GetRootIndex(root [32]byte) (uint32, error) {
	return s.CoordinatorCaller.GetRootIndexFromHash(&bind.CallOpts{}, root)
}

func (t *TransactorImpl) SubmitRoot(ctx context.Context, root [32]byte, rewardsUnixTimestamp uint32) error {
	tx, err := t.CoordinatorTransactor.SubmitRoot(t.ChainClient.NoSendTransactOpts, root, rewardsUnixTimestamp)
	if err != nil {
		fmt.Printf("Rewards coordinator, failed to submit root: %+v - %+v\n", err, tx)
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
