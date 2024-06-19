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
	SubmitRewardClaim(ctx context.Context, claim rewardsCoordinator.IRewardsCoordinatorRewardsMerkleClaim, earnerAddress gethcommon.Address) error
	GetRootByIndex(index uint64) (*rewardsCoordinator.IRewardsCoordinatorDistributionRoot, error)
	GetCurrentRoot() (*rewardsCoordinator.IRewardsCoordinatorDistributionRoot, error)
}

type TransactorImpl struct {
	ChainClient           *chainClient.ChainClient
	CoordinatorCaller     *rewardsCoordinator.IRewardsCoordinatorCaller
	CoordinatorTransactor *rewardsCoordinator.IRewardsCoordinatorTransactor
	RawContract           *bind.BoundContract
}

func getRawRewardsCoordinator(address gethcommon.Address, chainClient *chainClient.ChainClient) (*bind.BoundContract, error) {
	parsed, err := rewardsCoordinator.IRewardsCoordinatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	boundContract := bind.NewBoundContract(address, *parsed, chainClient, nil, nil)

	return boundContract, nil
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

	boundContract, err := getRawRewardsCoordinator(coordinatorAddress, chainClient)
	if err != nil {
		return nil, err
	}

	return &TransactorImpl{
		ChainClient:           chainClient,
		CoordinatorCaller:     coordinatorCaller,
		CoordinatorTransactor: coordinatorTransactor,
		RawContract:           boundContract,
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
		return fmt.Errorf("Rewards coordinator, failed to submit root: %+v - %+v", err, tx)
	}

	receipt, err := t.ChainClient.EstimateGasPriceAndLimitAndSendTx(ctx, tx, "submitRoot")
	if err != nil {
		return fmt.Errorf("Failed to estimate gas: %+v\n", err)
	}

	if receipt.Status != 1 {
		return chainClient.ErrTransactionFailed
	}

	return nil
}

func (t *TransactorImpl) SubmitRewardClaim(ctx context.Context, claim rewardsCoordinator.IRewardsCoordinatorRewardsMerkleClaim, earnerAddress gethcommon.Address) error {
	tx, err := t.CoordinatorTransactor.ProcessClaim(t.ChainClient.NoSendTransactOpts, claim, earnerAddress)
	if err != nil {
		return fmt.Errorf("Rewards coordinator, failed to submit reward claim: %+v - %+v", err, tx)
	}

	receipt, err := t.ChainClient.EstimateGasPriceAndLimitAndSendTx(ctx, tx, "submitRewardClaim")
	if err != nil {
		return fmt.Errorf("Failed to estimate gas: %+v\n", err)
	}

	if receipt.Status != 1 {
		return chainClient.ErrTransactionFailed
	}

	return nil
}

func (t *TransactorImpl) GetRootByIndex(index uint64) (*rewardsCoordinator.IRewardsCoordinatorDistributionRoot, error) {
	root, err := t.CoordinatorCaller.GetDistributionRootAtIndex(&bind.CallOpts{}, big.NewInt(int64(index)))
	return &root, err
}

func (t *TransactorImpl) GetCurrentRoot() (*rewardsCoordinator.IRewardsCoordinatorDistributionRoot, error) {
	root, err := t.CoordinatorCaller.GetCurrentDistributionRoot(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	return &root, nil
}
