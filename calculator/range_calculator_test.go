package calculator

import (
	"context"
	"math/big"
	"testing"

	"github.com/Layr-Labs/eigenlayer-payment-updater/calculator/mocks"
	"github.com/Layr-Labs/eigenlayer-payment-updater/common"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/mock"
)

func TestRangePaymentCalculator(t *testing.T) {

	intervalSecondsLength := big.NewInt(100)
	startTimestamp := big.NewInt(200)

	t.Run("test GetPaymentsCalculatedUntilTimestamp with no range payments", func(t *testing.T) {
		mockPaymentCalculatorDataService := &mocks.PaymentCalculatorDataService{}

		elpc := NewRangePaymentCalculator(intervalSecondsLength, mockPaymentCalculatorDataService)

		mockPaymentCalculatorDataService.On("GetPaymentsCalculatedUntilTimestamp", mock.Anything).Return(startTimestamp, nil)
		emptyDistributions := make(map[gethcommon.Address]*common.Distribution)
		mockPaymentCalculatorDataService.On("GetDistributionsAtTimestamp", mock.AnythingOfType("*big.Int")).Return(emptyDistributions, nil)
		mockPaymentCalculatorDataService.On("GetRangePaymentsWithOverlappingRange", mock.AnythingOfType("*big.Int"), mock.AnythingOfType("*big.Int")).Return(nil, pgx.ErrNoRows)

		endTimestampPassedIn := big.NewInt(300)
		endTimestamp, distributions, err := elpc.CalculateDistributionsUntilTimestamp(context.Background(), endTimestampPassedIn)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if endTimestamp.Cmp(endTimestampPassedIn) != 0 {
			t.Errorf("expected end timestamp to be %s, got %d", endTimestampPassedIn, endTimestamp)
		}

		// make sure distributions are empty
		if len(distributions) != 0 {
			t.Errorf("expected distributions to be empty, got %v", distributions)
		}
	})
}
