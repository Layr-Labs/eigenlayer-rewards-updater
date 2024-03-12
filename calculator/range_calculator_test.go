package calculator

import (
	"context"
	"math/big"
	"testing"

	"github.com/Layr-Labs/eigenlayer-payment-updater/calculator/mocks"
	"github.com/Layr-Labs/eigenlayer-payment-updater/calculator/utils"
	"github.com/Layr-Labs/eigenlayer-payment-updater/common"
	"github.com/Layr-Labs/eigenlayer-payment-updater/common/distribution"
	commonmocks "github.com/Layr-Labs/eigenlayer-payment-updater/common/services/mocks"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	EIGENDA_ADDRESS = gethcommon.HexToAddress("0x9FcE30E01a740660189bD8CbEaA48Abd36040010")
	STETH_ADDRESS   = gethcommon.HexToAddress("0xae7ab96520DE3A18E5e111B5EaAb095312D7fE84")
	WETH_ADDRESS    = gethcommon.HexToAddress("0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2")
)

func TestRangePaymentCalculator(t *testing.T) {

	calculationIntervalSeconds := int64(100)
	startTimestamp := big.NewInt(200)

	// testRangePayments := []*contractIPaymentCoordinator.IPaymentCoordinatorRangePayment{
	// 	{
	// 		Avs:                 EIGENDA_ADDRESS,
	// 		Strategy:            BEACON_CHAIN_ETH_STRATEGY_ADDRESS,
	// 		Token:               STETH_ADDRESS,
	// 		Amount:              big.NewInt(1000000000000),
	// 		StartRangeTimestamp: big.NewInt(200),
	// 		EndRangeTimestamp:   big.NewInt(700),
	// 	},
	// 	{
	// 		Avs:                 EIGENDA_ADDRESS,
	// 		Strategy:            BEACON_CHAIN_ETH_STRATEGY_ADDRESS,
	// 		Token:               WETH_ADDRESS,
	// 		Amount:              big.NewInt(2000000000000),
	// 		StartRangeTimestamp: big.NewInt(450),
	// 		EndRangeTimestamp:   big.NewInt(700),
	// 	},
	// }

	t.Run("test GetPaymentsCalculatedUntilTimestamp with no range payments", func(t *testing.T) {
		mockPaymentsDataService := &commonmocks.PaymentsDataService{}
		mockOperatorSetDataService := &mocks.OperatorSetDataService{}

		elpc := NewRangePaymentCalculator(calculationIntervalSeconds, mockPaymentsDataService, mockOperatorSetDataService)

		mockPaymentsDataService.On("GetRangePaymentsWithOverlappingRange", mock.AnythingOfType("*big.Int"), mock.AnythingOfType("*big.Int"), mock.AnythingOfType("*big.Int"), mock.AnythingOfType("*big.Int")).Return(nil, pgx.ErrNoRows)

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
}

func TestCalculateDistributionToOperatorForInterval(t *testing.T) {
	tinyPaymentToDistributePerInterval := big.NewInt(50)
	normalPaymentToDistributePerInterval := big.NewInt(1000000000000)

	testStrategies1 := []gethcommon.Address{utils.TEST_STRATEGY_ADDRESS_1}
	testMultipliers1 := []*big.Int{big.NewInt(1)}

	testStrategies2 := []gethcommon.Address{utils.TEST_STRATEGY_ADDRESS_1, utils.TEST_STRATEGY_ADDRESS_2}
	testMultipliers2 := []*big.Int{big.NewInt(1), big.NewInt(2)}

	t.Run("test CalculateDistributionToOperatorForInterval for single operator operatorSet", func(t *testing.T) {
		operatorSet := &common.OperatorSet{
			Operators: []*common.Operator{utils.GetSelfDelegatedOperator()},
		}
		operatorSet.FillTotals()

		totalStakeWeight := operatorSet.TotalStakedWeight(testStrategies1, testMultipliers1)

		diffDistribution := distribution.NewDistribution()
		diffDistribution = CalculateDistributionToOperatorForInterval(context.Background(), diffDistribution, 0, operatorSet, totalStakeWeight, testStrategies1, testMultipliers1, STETH_ADDRESS, normalPaymentToDistributePerInterval)

		amount, _ := diffDistribution.Get(operatorSet.Operators[0].Recipient, STETH_ADDRESS)
		assert.Equal(t, normalPaymentToDistributePerInterval, amount)
	})

	t.Run("test CalculateDistributionToOperatorForInterval for single operator with 2 outside stakers", func(t *testing.T) {
		operatorSet := &common.OperatorSet{
			Operators: []*common.Operator{utils.GetOperatorWith2OutsideStakers()},
		}
		operatorSet.FillTotals()

		totalStakeWeight := operatorSet.TotalStakedWeight(testStrategies1, testMultipliers1)

		diffDistribution := distribution.NewDistribution()

		diffDistribution = CalculateDistributionToOperatorForInterval(context.Background(), diffDistribution, 0, operatorSet, totalStakeWeight, testStrategies1, testMultipliers1, STETH_ADDRESS, normalPaymentToDistributePerInterval)

		amount, _ := diffDistribution.Get(operatorSet.Operators[0].Recipient, STETH_ADDRESS)
		assert.Equal(t, big.NewInt(100000000000), amount)
		amount, _ = diffDistribution.Get(operatorSet.Operators[0].Stakers[1].Recipient, STETH_ADDRESS)
		assert.Equal(t, big.NewInt(450000000000), amount)
		amount, _ = diffDistribution.Get(operatorSet.Operators[0].Stakers[2].Recipient, STETH_ADDRESS)
		assert.Equal(t, big.NewInt(450000000000), amount)
	})

	t.Run("test CalculateDistributionToOperatorForInterval for single operator with 1 outside staker", func(t *testing.T) {
		operatorSet := &common.OperatorSet{
			Operators: []*common.Operator{utils.GetOperatorWith1OutsideStaker()},
		}
		operatorSet.FillTotals()

		totalStakeWeight := operatorSet.TotalStakedWeight(testStrategies1, testMultipliers1)

		diffDistribution := distribution.NewDistribution()

		diffDistribution = CalculateDistributionToOperatorForInterval(context.Background(), diffDistribution, 0, operatorSet, totalStakeWeight, testStrategies1, testMultipliers1, STETH_ADDRESS, normalPaymentToDistributePerInterval)

		amount, _ := diffDistribution.Get(operatorSet.Operators[0].Recipient, STETH_ADDRESS)
		assert.Equal(t, big.NewInt(550000000000), amount)
		amount, _ = diffDistribution.Get(operatorSet.Operators[0].Stakers[1].Recipient, STETH_ADDRESS)
		assert.Equal(t, big.NewInt(450000000000), amount)
	})

	t.Run("test CalculateDistributionToOperatorForInterval for three operators equal split", func(t *testing.T) {
		operatorSet := &common.OperatorSet{
			Operators: []*common.Operator{utils.GetSelfDelegatedOperator(), utils.GetOperatorWith2OutsideStakers(), utils.GetOperatorWith1OutsideStaker()},
		}
		operatorSet.FillTotals()

		totalStakeWeight := operatorSet.TotalStakedWeight(testStrategies1, testMultipliers1)

		diffDistribution := distribution.NewDistribution()

		diffDistribution = CalculateDistributionToOperatorForInterval(context.Background(), diffDistribution, 0, operatorSet, totalStakeWeight, testStrategies1, testMultipliers1, STETH_ADDRESS, normalPaymentToDistributePerInterval)
		diffDistribution = CalculateDistributionToOperatorForInterval(context.Background(), diffDistribution, 1, operatorSet, totalStakeWeight, testStrategies1, testMultipliers1, STETH_ADDRESS, normalPaymentToDistributePerInterval)
		diffDistribution = CalculateDistributionToOperatorForInterval(context.Background(), diffDistribution, 2, operatorSet, totalStakeWeight, testStrategies1, testMultipliers1, STETH_ADDRESS, normalPaymentToDistributePerInterval)

		amount, _ := diffDistribution.Get(operatorSet.Operators[0].Recipient, STETH_ADDRESS)
		assert.Equal(t, big.NewInt(333333333333), amount)
		amount, _ = diffDistribution.Get(operatorSet.Operators[1].Recipient, STETH_ADDRESS)
		assert.Equal(t, big.NewInt(33333333333), amount)
		amount, _ = diffDistribution.Get(operatorSet.Operators[1].Stakers[1].Recipient, STETH_ADDRESS)
		assert.Equal(t, big.NewInt(150000000000), amount)
		amount, _ = diffDistribution.Get(operatorSet.Operators[1].Stakers[2].Recipient, STETH_ADDRESS)
		assert.Equal(t, big.NewInt(150000000000), amount)
		amount, _ = diffDistribution.Get(operatorSet.Operators[2].Recipient, STETH_ADDRESS)
		assert.Equal(t, big.NewInt(183333333333), amount)
		amount, _ = diffDistribution.Get(operatorSet.Operators[2].Stakers[1].Recipient, STETH_ADDRESS)
		assert.Equal(t, big.NewInt(150000000000), amount)
	})

	t.Run("test CalculateDistributionToOperatorForInterval for three operators 1/10 3/10 6/10", func(t *testing.T) {
		operatorSet := get3Opereator136Split()
		totalStakeWeight := operatorSet.TotalStakedWeight(testStrategies1, testMultipliers1)

		diffDistribution := distribution.NewDistribution()

		diffDistribution = CalculateDistributionToOperatorForInterval(context.Background(), diffDistribution, 0, operatorSet, totalStakeWeight, testStrategies1, testMultipliers1, STETH_ADDRESS, normalPaymentToDistributePerInterval)
		diffDistribution = CalculateDistributionToOperatorForInterval(context.Background(), diffDistribution, 1, operatorSet, totalStakeWeight, testStrategies1, testMultipliers1, STETH_ADDRESS, normalPaymentToDistributePerInterval)
		diffDistribution = CalculateDistributionToOperatorForInterval(context.Background(), diffDistribution, 2, operatorSet, totalStakeWeight, testStrategies1, testMultipliers1, STETH_ADDRESS, normalPaymentToDistributePerInterval)

		amount, _ := diffDistribution.Get(operatorSet.Operators[0].Recipient, STETH_ADDRESS)
		assert.Equal(t, big.NewInt(100000000000), amount)
		amount, _ = diffDistribution.Get(operatorSet.Operators[1].Recipient, STETH_ADDRESS)
		assert.Equal(t, big.NewInt(30000000000), amount)
		amount, _ = diffDistribution.Get(operatorSet.Operators[1].Stakers[1].Recipient, STETH_ADDRESS)
		assert.Equal(t, big.NewInt(180000000000), amount)
		amount, _ = diffDistribution.Get(operatorSet.Operators[1].Stakers[2].Recipient, STETH_ADDRESS)
		assert.Equal(t, big.NewInt(90000000000), amount)
		amount, _ = diffDistribution.Get(operatorSet.Operators[2].Recipient, STETH_ADDRESS)
		assert.Equal(t, big.NewInt(420000000000), amount)
		amount, _ = diffDistribution.Get(operatorSet.Operators[2].Stakers[1].Recipient, STETH_ADDRESS)
		assert.Equal(t, big.NewInt(180000000000), amount)
	})

	t.Run("test CalculateDistributionToOperatorForInterval for two operators with tiny payment", func(t *testing.T) {
		operatorSet := &common.OperatorSet{
			Operators: []*common.Operator{utils.GetSelfDelegatedOperator(), utils.GetOperatorWith2OutsideStakers()},
		}
		operatorSet.ModifyShares(utils.TEST_OPERATOR_ADDRESS_2, utils.TEST_STAKER_ADDRESS_1, utils.TEST_STRATEGY_ADDRESS_1, big.NewInt(1e17))
		totalStakeWeight := operatorSet.TotalStakedWeight(testStrategies1, testMultipliers1)

		diffDistribution := distribution.NewDistribution()

		diffDistribution = CalculateDistributionToOperatorForInterval(context.Background(), diffDistribution, 0, operatorSet, totalStakeWeight, testStrategies1, testMultipliers1, STETH_ADDRESS, tinyPaymentToDistributePerInterval)
		diffDistribution = CalculateDistributionToOperatorForInterval(context.Background(), diffDistribution, 1, operatorSet, totalStakeWeight, testStrategies1, testMultipliers1, STETH_ADDRESS, tinyPaymentToDistributePerInterval)

		amount, _ := diffDistribution.Get(operatorSet.Operators[0].Recipient, STETH_ADDRESS)
		assert.Equal(t, big.NewInt(31), amount)
		amount, _ = diffDistribution.Get(operatorSet.Operators[1].Recipient, STETH_ADDRESS)
		assert.Equal(t, big.NewInt(1), amount)
		amount, _ = diffDistribution.Get(operatorSet.Operators[1].Stakers[1].Recipient, STETH_ADDRESS)
		assert.Equal(t, big.NewInt(2), amount)
		amount, _ = diffDistribution.Get(operatorSet.Operators[1].Stakers[2].Recipient, STETH_ADDRESS)
		assert.Equal(t, big.NewInt(14), amount)
	})

	t.Run("test CalculateDistributionToOperatorForInterval for two operators with tiny payment round down to 0", func(t *testing.T) {
		operatorSet := &common.OperatorSet{
			Operators: []*common.Operator{utils.GetSelfDelegatedOperator(), utils.GetOperatorWith2OutsideStakers()},
		}
		operatorSet.ModifyShares(utils.TEST_OPERATOR_ADDRESS_2, utils.TEST_STAKER_ADDRESS_1, utils.TEST_STRATEGY_ADDRESS_1, big.NewInt(1e17))
		operatorSet.ModifyShares(utils.TEST_OPERATOR_ADDRESS_2, utils.TEST_STAKER_ADDRESS_2, utils.TEST_STRATEGY_ADDRESS_1, big.NewInt(1e17))
		totalStakeWeight := operatorSet.TotalStakedWeight(testStrategies1, testMultipliers1)

		diffDistribution := distribution.NewDistribution()

		diffDistribution = CalculateDistributionToOperatorForInterval(context.Background(), diffDistribution, 0, operatorSet, totalStakeWeight, testStrategies1, testMultipliers1, STETH_ADDRESS, tinyPaymentToDistributePerInterval)
		diffDistribution = CalculateDistributionToOperatorForInterval(context.Background(), diffDistribution, 1, operatorSet, totalStakeWeight, testStrategies1, testMultipliers1, STETH_ADDRESS, tinyPaymentToDistributePerInterval)

		amount, _ := diffDistribution.Get(operatorSet.Operators[0].Recipient, STETH_ADDRESS)
		assert.Equal(t, big.NewInt(41), amount)
		amount, _ = diffDistribution.Get(operatorSet.Operators[1].Recipient, STETH_ADDRESS)
		assert.Equal(t, big.NewInt(0), amount)
		amount, _ = diffDistribution.Get(operatorSet.Operators[1].Stakers[1].Recipient, STETH_ADDRESS)
		assert.Equal(t, big.NewInt(4), amount)
		amount, _ = diffDistribution.Get(operatorSet.Operators[1].Stakers[2].Recipient, STETH_ADDRESS)
		assert.Equal(t, big.NewInt(4), amount)
	})

	t.Run("test CalculateDistributionToOperatorForInterval for three operators 1/10 3/10 6/10 with randomized Recipients", func(t *testing.T) {
		operatorSet := get3Opereator136Split()
		operatorSet.RandomizeRecipients()
		totalStakeWeight := operatorSet.TotalStakedWeight(testStrategies1, testMultipliers1)

		diffDistribution := distribution.NewDistribution()

		diffDistribution = CalculateDistributionToOperatorForInterval(context.Background(), diffDistribution, 0, operatorSet, totalStakeWeight, testStrategies1, testMultipliers1, STETH_ADDRESS, normalPaymentToDistributePerInterval)
		diffDistribution = CalculateDistributionToOperatorForInterval(context.Background(), diffDistribution, 1, operatorSet, totalStakeWeight, testStrategies1, testMultipliers1, STETH_ADDRESS, normalPaymentToDistributePerInterval)
		diffDistribution = CalculateDistributionToOperatorForInterval(context.Background(), diffDistribution, 2, operatorSet, totalStakeWeight, testStrategies1, testMultipliers1, STETH_ADDRESS, normalPaymentToDistributePerInterval)

		amount, _ := diffDistribution.Get(operatorSet.Operators[0].Recipient, STETH_ADDRESS)
		assert.Equal(t, big.NewInt(100000000000), amount)
		amount, _ = diffDistribution.Get(operatorSet.Operators[1].Recipient, STETH_ADDRESS)
		assert.Equal(t, big.NewInt(30000000000), amount)
		amount, _ = diffDistribution.Get(operatorSet.Operators[1].Stakers[1].Recipient, STETH_ADDRESS)
		assert.Equal(t, big.NewInt(180000000000), amount)
		amount, _ = diffDistribution.Get(operatorSet.Operators[1].Stakers[2].Recipient, STETH_ADDRESS)
		assert.Equal(t, big.NewInt(90000000000), amount)
		amount, _ = diffDistribution.Get(operatorSet.Operators[2].Recipient, STETH_ADDRESS)
		assert.Equal(t, big.NewInt(420000000000), amount)
		amount, _ = diffDistribution.Get(operatorSet.Operators[2].Stakers[1].Recipient, STETH_ADDRESS)
		assert.Equal(t, big.NewInt(180000000000), amount)
	})

	t.Run("test CalculateDistributionToOperatorForInterval for three operators 1/10 3/10 6/10 with nonzero initial amounts", func(t *testing.T) {
		operatorSet := get3Opereator136Split()
		totalStakeWeight := operatorSet.TotalStakedWeight([]gethcommon.Address{utils.TEST_STRATEGY_ADDRESS_1}, []*big.Int{big.NewInt(1)})

		diffDistribution := distribution.NewDistribution()
		diffDistribution.Set(operatorSet.Operators[0].Recipient, STETH_ADDRESS, big.NewInt(1))
		diffDistribution.Set(operatorSet.Operators[1].Recipient, STETH_ADDRESS, big.NewInt(2))
		diffDistribution.Set(operatorSet.Operators[1].Stakers[1].Recipient, STETH_ADDRESS, big.NewInt(3))
		diffDistribution.Set(operatorSet.Operators[1].Stakers[2].Recipient, STETH_ADDRESS, big.NewInt(4))
		diffDistribution.Set(operatorSet.Operators[2].Recipient, STETH_ADDRESS, big.NewInt(5))
		diffDistribution.Set(operatorSet.Operators[2].Stakers[1].Recipient, STETH_ADDRESS, big.NewInt(6))

		diffDistribution = CalculateDistributionToOperatorForInterval(context.Background(), diffDistribution, 0, operatorSet, totalStakeWeight, testStrategies1, testMultipliers1, STETH_ADDRESS, normalPaymentToDistributePerInterval)
		diffDistribution = CalculateDistributionToOperatorForInterval(context.Background(), diffDistribution, 1, operatorSet, totalStakeWeight, testStrategies1, testMultipliers1, STETH_ADDRESS, normalPaymentToDistributePerInterval)
		diffDistribution = CalculateDistributionToOperatorForInterval(context.Background(), diffDistribution, 2, operatorSet, totalStakeWeight, testStrategies1, testMultipliers1, STETH_ADDRESS, normalPaymentToDistributePerInterval)

		amount, _ := diffDistribution.Get(operatorSet.Operators[0].Recipient, STETH_ADDRESS)
		assert.Equal(t, big.NewInt(100000000001), amount)
		amount, _ = diffDistribution.Get(operatorSet.Operators[1].Recipient, STETH_ADDRESS)
		assert.Equal(t, big.NewInt(30000000002), amount)
		amount, _ = diffDistribution.Get(operatorSet.Operators[1].Stakers[1].Recipient, STETH_ADDRESS)
		assert.Equal(t, big.NewInt(180000000003), amount)
		amount, _ = diffDistribution.Get(operatorSet.Operators[1].Stakers[2].Recipient, STETH_ADDRESS)
		assert.Equal(t, big.NewInt(90000000004), amount)
		amount, _ = diffDistribution.Get(operatorSet.Operators[2].Recipient, STETH_ADDRESS)
		assert.Equal(t, big.NewInt(420000000005), amount)
		amount, _ = diffDistribution.Get(operatorSet.Operators[2].Stakers[1].Recipient, STETH_ADDRESS)
		assert.Equal(t, big.NewInt(180000000006), amount)
	})

	t.Run("test CalculateDistributionToOperatorForInterval for three operator for two strategies and multipliers", func(t *testing.T) {
		operatorSet := &common.OperatorSet{Operators: []*common.Operator{utils.GetSelfDelegatedOperator(), utils.GetOperatorWith2OutsideStakers(), utils.GetOperatorWith1OutsideStaker()}}
		operatorSet.FillTotals()

		totalStakeWeight := operatorSet.TotalStakedWeight(testStrategies2, testMultipliers2)

		diffDistribution := distribution.NewDistribution()
		diffDistribution = CalculateDistributionToOperatorForInterval(context.Background(), diffDistribution, 0, operatorSet, totalStakeWeight, testStrategies2, testMultipliers2, STETH_ADDRESS, normalPaymentToDistributePerInterval)
		diffDistribution = CalculateDistributionToOperatorForInterval(context.Background(), diffDistribution, 1, operatorSet, totalStakeWeight, testStrategies2, testMultipliers2, STETH_ADDRESS, normalPaymentToDistributePerInterval)
		diffDistribution = CalculateDistributionToOperatorForInterval(context.Background(), diffDistribution, 2, operatorSet, totalStakeWeight, testStrategies2, testMultipliers2, STETH_ADDRESS, normalPaymentToDistributePerInterval)

		amount, _ := diffDistribution.Get(operatorSet.Operators[0].Recipient, STETH_ADDRESS)
		assert.Equal(t, big.NewInt(405405405405), amount)
		amount, _ = diffDistribution.Get(operatorSet.Operators[1].Recipient, STETH_ADDRESS)
		assert.Equal(t, big.NewInt(40540540540), amount)
		amount, _ = diffDistribution.Get(operatorSet.Operators[1].Stakers[1].Recipient, STETH_ADDRESS)
		assert.Equal(t, big.NewInt(109459459459), amount)
		amount, _ = diffDistribution.Get(operatorSet.Operators[1].Stakers[2].Recipient, STETH_ADDRESS)
		assert.Equal(t, big.NewInt(255405405405), amount)
		amount, _ = diffDistribution.Get(operatorSet.Operators[2].Recipient, STETH_ADDRESS)
		assert.Equal(t, big.NewInt(128378378377), amount)
		amount, _ = diffDistribution.Get(operatorSet.Operators[2].Stakers[1].Recipient, STETH_ADDRESS)
		assert.Equal(t, big.NewInt(60810810811), amount)
	})
}

func get3Opereator136Split() *common.OperatorSet {
	operatorSet := &common.OperatorSet{
		Operators: []*common.Operator{utils.GetSelfDelegatedOperator(), utils.GetOperatorWith2OutsideStakers(), utils.GetOperatorWith1OutsideStaker()},
	}
	operatorSet.ModifyShares(utils.TEST_OPERATOR_ADDRESS_1, utils.TEST_OPERATOR_ADDRESS_1, utils.TEST_STRATEGY_ADDRESS_1, big.NewInt(1e17))

	operatorSet.ModifyShares(utils.TEST_OPERATOR_ADDRESS_2, utils.TEST_STAKER_ADDRESS_1, utils.TEST_STRATEGY_ADDRESS_1, big.NewInt(2e17))
	operatorSet.ModifyShares(utils.TEST_OPERATOR_ADDRESS_2, utils.TEST_STAKER_ADDRESS_2, utils.TEST_STRATEGY_ADDRESS_1, big.NewInt(1e17))

	operatorSet.ModifyShares(utils.TEST_OPERATOR_ADDRESS_3, utils.TEST_OPERATOR_ADDRESS_3, utils.TEST_STRATEGY_ADDRESS_1, big.NewInt(4e17))
	operatorSet.ModifyShares(utils.TEST_OPERATOR_ADDRESS_3, utils.TEST_STAKER_ADDRESS_3, utils.TEST_STRATEGY_ADDRESS_1, big.NewInt(2e17))

	return operatorSet
}
