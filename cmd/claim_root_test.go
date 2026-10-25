package cmd

import (
	"math/big"
	"testing"
	"time"

	rewardsCoordinator "github.com/Layr-Labs/eigenlayer-contracts/pkg/bindings/IRewardsCoordinator"
	"github.com/Layr-Labs/eigenlayer-rewards-updater/mocks"
)

func TestLatestClaimableRootSkipsPendingAndDisabledRoots(t *testing.T) {
	now := time.Unix(1792500000, 0) // 2026-10-25, between the activation of roots 119 and 120
	transactor := &mocks.Transactor{}
	transactor.On("GetNumberOfPublishedRoots").Return(big.NewInt(3), nil)
	// index 2: newest root, still inside its activation delay
	transactor.On("GetRootByIndex", uint64(2)).Return(&rewardsCoordinator.IRewardsCoordinatorDistributionRoot{
		RewardsCalculationEndTimestamp: 1792281600, ActivatedAt: 1793113295,
	}, nil)
	// index 1: disabled
	transactor.On("GetRootByIndex", uint64(1)).Return(&rewardsCoordinator.IRewardsCoordinatorDistributionRoot{
		RewardsCalculationEndTimestamp: 1791676800, ActivatedAt: 1792000000, Disabled: true,
	}, nil)
	// index 0: active
	transactor.On("GetRootByIndex", uint64(0)).Return(&rewardsCoordinator.IRewardsCoordinatorDistributionRoot{
		RewardsCalculationEndTimestamp: 1791072000, ActivatedAt: 1791900000,
	}, nil)

	index, date, err := latestClaimableRoot(transactor, now)
	if err != nil {
		t.Fatal(err)
	}
	if index != 0 || date != "2026-10-04" {
		t.Fatalf("got index %d date %s, want index 0 date 2026-10-04", index, date)
	}
}
