package common

import (
	"crypto/rand"
	"encoding/hex"
	"math/big"

	gethcommon "github.com/ethereum/go-ethereum/common"
)

type Earner struct {
	Recipient gethcommon.Address
}

type Staker struct {
	Earner
	Address gethcommon.Address
	Weight  *big.Int
}

type Operator struct {
	Earner
	Address         gethcommon.Address
	Commission      *big.Int
	DelegatedWeight *big.Int
	Stakers         []*Staker
}

type OperatorSet struct {
	TotalStakedWeight *big.Int
	Operators         []*Operator
}

func (os *OperatorSet) FillTotals() {
	os.TotalStakedWeight = big.NewInt(0)
	for _, operator := range os.Operators {
		operator.DelegatedWeight = big.NewInt(0)
		for _, staker := range operator.Stakers {
			operator.DelegatedWeight.Add(operator.DelegatedWeight, staker.Weight)
		}
		os.TotalStakedWeight.Add(os.TotalStakedWeight, operator.DelegatedWeight)
	}
}

func (os *OperatorSet) ModifyWeight(operatorAddress, stakerAddress gethcommon.Address, newWeight *big.Int) {
	for _, operator := range os.Operators {
		if operator.Address == operatorAddress {
			for _, staker := range operator.Stakers {
				if staker.Address == stakerAddress {
					staker.Weight = newWeight
					break
				}
			}
		}
	}

	os.FillTotals()
}

func (os *OperatorSet) RandomizeRecipients() {
	for _, operator := range os.Operators {
		operator.Recipient = GetRandomAddress()
		for _, staker := range operator.Stakers {
			if staker.Address == operator.Address {
				staker.Recipient = operator.Recipient
			} else {
				staker.Recipient = GetRandomAddress()
			}
		}
	}
}

func GetRandomAddress() gethcommon.Address {
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
