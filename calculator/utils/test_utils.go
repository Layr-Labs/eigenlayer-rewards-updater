package utils

import (
	"encoding/json"
	"math/big"

	"github.com/Layr-Labs/eigenlayer-payment-updater/common"
	gethcommon "github.com/ethereum/go-ethereum/common"
)

var (
	TEST_OPERATOR_ADDRESS_1 = gethcommon.HexToAddress("0x0000111111111111111111111111111111111111")
	TEST_OPERATOR_ADDRESS_2 = gethcommon.HexToAddress("0x0000222222222222222222222222222222222222")
	TEST_OPERATOR_ADDRESS_3 = gethcommon.HexToAddress("0x0000333333333333333333333333333333333333")

	TEST_STAKER_ADDRESS_1 = gethcommon.HexToAddress("0x1111111111111111111111111111111111111111")
	TEST_STAKER_ADDRESS_2 = gethcommon.HexToAddress("0x1111222222222222222222222222222222222222")
	TEST_STAKER_ADDRESS_3 = gethcommon.HexToAddress("0x1111333333333333333333333333333333333333")

	TEST_STRATEGY_ADDRESS_1 = gethcommon.HexToAddress("0x1234567890987654321234567890987654321234")
	TEST_STRATEGY_ADDRESS_2 = gethcommon.HexToAddress("0x0987654321234567890987654321234567890987")

	// a self delegated operator with no other stakers
	SelfDelegatedOperator = &common.Operator{
		Earner: common.Earner{
			Recipient: TEST_OPERATOR_ADDRESS_1,
		},
		Address:    TEST_OPERATOR_ADDRESS_1,
		Commission: big.NewInt(1000),
		DelegatedShares: map[gethcommon.Address]*big.Int{
			TEST_STRATEGY_ADDRESS_1: big.NewInt(1e18),
		},
		Stakers: []*common.Staker{
			{
				Earner: common.Earner{
					Recipient: TEST_OPERATOR_ADDRESS_1,
				},
				Address: TEST_OPERATOR_ADDRESS_1,
				Shares: map[gethcommon.Address]*big.Int{
					TEST_STRATEGY_ADDRESS_1: big.NewInt(1e18),
				},
			},
		},
	}

	// an operator with 2 stakers
	OperatorWith2OutsideStakers = &common.Operator{
		Earner: common.Earner{
			Recipient: TEST_OPERATOR_ADDRESS_2,
		},
		Address:    TEST_OPERATOR_ADDRESS_2,
		Commission: big.NewInt(1000),
		DelegatedShares: map[gethcommon.Address]*big.Int{
			TEST_STRATEGY_ADDRESS_1: big.NewInt(1e18),
		},
		Stakers: []*common.Staker{
			{
				Earner: common.Earner{
					Recipient: TEST_OPERATOR_ADDRESS_1,
				},
				Address: TEST_OPERATOR_ADDRESS_1,
				Shares: map[gethcommon.Address]*big.Int{
					TEST_STRATEGY_ADDRESS_1: big.NewInt(0),
				},
			},
			{
				Earner: common.Earner{
					Recipient: TEST_STAKER_ADDRESS_1,
				},
				Address: TEST_STAKER_ADDRESS_1,
				Shares: map[gethcommon.Address]*big.Int{
					TEST_STRATEGY_ADDRESS_1: big.NewInt(5e17),
				},
			},
			{
				Earner: common.Earner{
					Recipient: TEST_STAKER_ADDRESS_2,
				},
				Address: TEST_STAKER_ADDRESS_2,
				Shares: map[gethcommon.Address]*big.Int{
					TEST_STRATEGY_ADDRESS_1: big.NewInt(5e17),
				},
			},
		},
	}

	OperatorWith1OutsideStaker = &common.Operator{
		Earner: common.Earner{
			Recipient: TEST_OPERATOR_ADDRESS_3,
		},
		Address:    TEST_OPERATOR_ADDRESS_3,
		Commission: big.NewInt(1000),
		DelegatedShares: map[gethcommon.Address]*big.Int{
			TEST_STRATEGY_ADDRESS_1: big.NewInt(1e18),
		},
		Stakers: []*common.Staker{
			{
				Earner: common.Earner{
					Recipient: TEST_OPERATOR_ADDRESS_3,
				},
				Address: TEST_OPERATOR_ADDRESS_3,
				Shares: map[gethcommon.Address]*big.Int{
					TEST_STRATEGY_ADDRESS_1: big.NewInt(5e17),
				},
			},
			{
				Earner: common.Earner{
					Recipient: TEST_STAKER_ADDRESS_3,
				},
				Address: TEST_STAKER_ADDRESS_3,
				Shares: map[gethcommon.Address]*big.Int{
					TEST_STRATEGY_ADDRESS_1: big.NewInt(5e17),
				},
			},
		},
	}
)

func DeepCopy(src, dst interface{}) error {
	bytes, err := json.Marshal(src)
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, dst)
}

func GetSelfDelegatedOperator() *common.Operator {
	operator := &common.Operator{}
	DeepCopy(SelfDelegatedOperator, operator)
	return operator
}

func GetOperatorWith2OutsideStakers() *common.Operator {
	operator := &common.Operator{}
	DeepCopy(OperatorWith2OutsideStakers, operator)
	return operator
}

func GetOperatorWith1OutsideStaker() *common.Operator {
	operator := &common.Operator{}
	DeepCopy(OperatorWith1OutsideStaker, operator)
	return operator
}
