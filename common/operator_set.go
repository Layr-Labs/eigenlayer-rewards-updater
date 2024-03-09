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
	Address        gethcommon.Address
	StrategyShares *big.Int
}

type Operator struct {
	Earner
	Address                      gethcommon.Address
	Commission                   *big.Int
	TotalDelegatedStrategyShares *big.Int
	Stakers                      []*Staker
}

type OperatorSet struct {
	TotalStakedStrategyShares *big.Int
	Operators                 []*Operator
}

func (os *OperatorSet) FillTotals() {
	os.TotalStakedStrategyShares = big.NewInt(0)
	for _, operator := range os.Operators {
		operator.TotalDelegatedStrategyShares = big.NewInt(0)
		for _, staker := range operator.Stakers {
			operator.TotalDelegatedStrategyShares = new(big.Int).Add(operator.TotalDelegatedStrategyShares, staker.StrategyShares)
		}
		os.TotalStakedStrategyShares = new(big.Int).Add(os.TotalStakedStrategyShares, operator.TotalDelegatedStrategyShares)
	}
}

func (os *OperatorSet) ModifyStrategyShares(operatorAddress, stakerAddress gethcommon.Address, newStrategyShares *big.Int) {
	for _, operator := range os.Operators {
		if operator.Address == operatorAddress {
			for _, staker := range operator.Stakers {
				if staker.Address == stakerAddress {
					staker.StrategyShares = newStrategyShares
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
