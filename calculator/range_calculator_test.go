package calculator

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"math/big"
	"testing"

	contractIPaymentCoordinator "github.com/Layr-Labs/eigenlayer-payment-updater/bindings/IPaymentCoordinator"
	"github.com/Layr-Labs/eigenlayer-payment-updater/calculator/mocks"
	"github.com/Layr-Labs/eigenlayer-payment-updater/common"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/mock"
)

func TestRangePaymentCalculator(t *testing.T) {
	EIGENDA_ADDRESS := gethcommon.HexToAddress("0x9FcE30E01a740660189bD8CbEaA48Abd36040010")
	STETH_ADDRESS := gethcommon.HexToAddress("0xae7ab96520DE3A18E5e111B5EaAb095312D7fE84")
	WETH_ADDRESS := gethcommon.HexToAddress("0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2")

	calculationIntervalSeconds := int64(100)
	startTimestamp := big.NewInt(200)

	testRangePayments := []*contractIPaymentCoordinator.IPaymentCoordinatorRangePayment{
		{
			Avs:                 EIGENDA_ADDRESS,
			Strategy:            BEACON_CHAIN_ETH_STRATEGY_ADDRESS,
			Token:               STETH_ADDRESS,
			Amount:              big.NewInt(1000000000000),
			StartRangeTimestamp: big.NewInt(200),
			EndRangeTimestamp:   big.NewInt(700),
		},
		{
			Avs:                 EIGENDA_ADDRESS,
			Strategy:            BEACON_CHAIN_ETH_STRATEGY_ADDRESS,
			Token:               WETH_ADDRESS,
			Amount:              big.NewInt(2000000000000),
			StartRangeTimestamp: big.NewInt(450),
			EndRangeTimestamp:   big.NewInt(700),
		},
	}

	t.Run("test GetPaymentsCalculatedUntilTimestamp with no range payments", func(t *testing.T) {
		mockPaymentCalculatorDataService := &mocks.PaymentCalculatorDataService{}
		mockOperatorSetDataService := &mocks.OperatorSetDataService{}

		elpc := NewRangePaymentCalculator(calculationIntervalSeconds, mockPaymentCalculatorDataService, mockOperatorSetDataService)

		mockPaymentCalculatorDataService.On("GetRangePaymentsWithOverlappingRange", mock.AnythingOfType("*big.Int"), mock.AnythingOfType("*big.Int")).Return(nil, pgx.ErrNoRows)

		endTimestampPassedIn := big.NewInt(300)
		endTimestamp, distribution, err := elpc.CalculateDistributionUntilTimestamp(context.Background(), startTimestamp, endTimestampPassedIn)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if endTimestamp.Cmp(endTimestampPassedIn) != 0 {
			t.Errorf("expected end timestamp to be %s, got %d", endTimestampPassedIn, endTimestamp)
		}

		// make sure distribution are empty
		if distribution.GetNumLeaves() != 0 {
			t.Errorf("expected distribution to be empty, got %v", distribution)
		}
	})

	t.Run("test GetPaymentsCalculatedUntilTimestamp with single range payment for 1 interval", func(t *testing.T) {
		mockPaymentCalculatorDataService := &mocks.PaymentCalculatorDataService{}
		mockOperatorSetDataService := &mocks.OperatorSetDataService{}

		elpc := NewRangePaymentCalculator(calculationIntervalSeconds, mockPaymentCalculatorDataService, mockOperatorSetDataService)

		mockPaymentCalculatorDataService.On("GetRangePaymentsWithOverlappingRange", mock.AnythingOfType("*big.Int"), mock.AnythingOfType("*big.Int")).Return(testRangePayments[:1], nil)

		operatorSet := &common.OperatorSet{
			Operators: []common.Operator{
				{
					Earner: common.Earner{
						Claimer: getRandomAddress(),
					},
					Address:    getRandomAddress(),
					Commission: big.NewInt(5000), // 50%
					Stakers: []common.Staker{
						{
							Earner: common.Earner{
								Claimer: getRandomAddress(),
							},
							Address:        getRandomAddress(),
							StrategyShares: big.NewInt(100),
						},
						{
							Earner: common.Earner{
								Claimer: getRandomAddress(),
							},
							Address:        getRandomAddress(),
							StrategyShares: big.NewInt(200),
						},
					},
					TotalDelegatedStrategyShares: big.NewInt(300),
				},
				{
					Earner: common.Earner{
						Claimer: getRandomAddress(),
					},
					Address:    getRandomAddress(),
					Commission: big.NewInt(1000), // 10%
					Stakers: []common.Staker{
						{
							Earner: common.Earner{
								Claimer: getRandomAddress(),
							},
							Address:        getRandomAddress(),
							StrategyShares: big.NewInt(400),
						},
					},
					TotalDelegatedStrategyShares: big.NewInt(400),
				},
			},
			TotalStakedStrategyShares: big.NewInt(700),
		}

		mockOperatorSetDataService.On("GetOperatorSetForStrategyAtTimestamp", mock.AnythingOfType("*big.Int"), testRangePayments[0].Avs, testRangePayments[0].Strategy).Return(operatorSet, nil)
		endTimestampPassedIn := big.NewInt(300)
		endTimestamp, distribution, err := elpc.CalculateDistributionUntilTimestamp(context.Background(), startTimestamp, endTimestampPassedIn)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if endTimestamp.Cmp(endTimestampPassedIn) != 0 {
			t.Errorf("expected end timestamp to be %s, got %d", endTimestampPassedIn, endTimestamp)
		}

		if distribution.GetNumLeaves() != 5 {
			t.Errorf("expected distributions to have 5 entry, got %d", distribution.GetNumLeaves())
		}

		// make sure the disitrubution is accurate according to precalculated values
		if distribution.Get(operatorSet.Operators[0].Address, testRangePayments[0].Token).Cmp(big.NewInt(42857142857)) != 0 {
			t.Errorf("expected operator balance to be 42857142857, got %s", distribution.Get(operatorSet.Operators[0].Address, testRangePayments[0].Token))
		}

		if distribution.Get(operatorSet.Operators[0].Stakers[0].Address, testRangePayments[0].Token).Cmp(big.NewInt(14285714285)) != 0 {
			t.Errorf("expected staker balance to be 14285714285, got %s", distribution.Get(operatorSet.Operators[0].Stakers[0].Address, testRangePayments[0].Token))
		}

		if distribution.Get(operatorSet.Operators[0].Stakers[1].Address, testRangePayments[0].Token).Cmp(big.NewInt(28571428571)) != 0 {
			t.Errorf("expected staker balance to be 28571428571, got %s", distribution.Get(operatorSet.Operators[0].Stakers[1].Address, testRangePayments[0].Token))
		}

		if distribution.Get(operatorSet.Operators[1].Address, testRangePayments[0].Token).Cmp(big.NewInt(11428571428)) != 0 {
			t.Errorf("expected operator balance to be 14285714285, got %s", distribution.Get(operatorSet.Operators[1].Address, testRangePayments[0].Token))
		}

		if distribution.Get(operatorSet.Operators[1].Stakers[0].Address, testRangePayments[0].Token).Cmp(big.NewInt(102857142856)) != 0 {
			t.Errorf("expected staker balance to be 102857142856, got %s", distribution.Get(operatorSet.Operators[1].Stakers[0].Address, testRangePayments[0].Token))
		}
	})
}

func fillInTotals(operatorSet common.OperatorSet) common.OperatorSet {
	operatorSet.TotalStakedStrategyShares = big.NewInt(0)
	for i := 0; i < len(operatorSet.Operators); i++ {
		operatorSet.Operators[i].TotalDelegatedStrategyShares = big.NewInt(0)
		for j := 0; j < len(operatorSet.Operators[i].Stakers); j++ {
			operatorSet.Operators[i].TotalDelegatedStrategyShares.Add(operatorSet.Operators[i].TotalDelegatedStrategyShares, operatorSet.Operators[i].Stakers[j].StrategyShares)
		}
		operatorSet.TotalStakedStrategyShares.Add(operatorSet.TotalStakedStrategyShares, operatorSet.Operators[i].TotalDelegatedStrategyShares)
	}
	return operatorSet
}

func getStakerList(num int, minStake, maxStake *big.Int) []common.Staker {
	stakers := make([]common.Staker, num)
	for i := 0; i < num; i++ {
		stakers[i] = getRandomStaker(minStake, maxStake)
	}
	return stakers
}

func getRandomStaker(minStake, maxStake *big.Int) common.Staker {
	return common.Staker{
		Earner: common.Earner{
			Claimer: getRandomAddress(),
		},
		Address:        getRandomAddress(),
		StrategyShares: getRandomBigInt(minStake, maxStake),
	}
}

func getRandomBigInt(min, max *big.Int) *big.Int {
	diff := new(big.Int).Sub(max, min)
	randBigInt, _ := rand.Int(rand.Reader, diff)
	randBigInt.Add(randBigInt, min)
	return randBigInt
}

func getRandomAddress() gethcommon.Address {
	randomHex, _ := getRandomHex(20)
	return gethcommon.HexToAddress(randomHex)
}

func getRandomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
