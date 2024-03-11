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
	Shares  map[gethcommon.Address]*big.Int
}

type Operator struct {
	Earner
	Address         gethcommon.Address
	Commission      *big.Int
	DelegatedShares map[gethcommon.Address]*big.Int
	Stakers         []*Staker
}

type OperatorSet struct {
	TotalStakedShares map[gethcommon.Address]*big.Int
	Operators         []*Operator
}

func Weight(strategyToShares map[gethcommon.Address]*big.Int, strategies []gethcommon.Address, multipliers []*big.Int) *big.Int {
	weight := big.NewInt(0)
	for i, strategy := range strategies {
		shares, found := strategyToShares[strategy]
		if !found {
			continue
		}

		weight.Add(weight, new(big.Int).Mul(shares, multipliers[i]))
	}

	return weight
}

func (s *Staker) Weight(strategies []gethcommon.Address, multipliers []*big.Int) *big.Int {
	return Weight(s.Shares, strategies, multipliers)
}

func (op *Operator) Weight(strategies []gethcommon.Address, multipliers []*big.Int) *big.Int {
	return Weight(op.DelegatedShares, strategies, multipliers)
}

func (os *OperatorSet) TotalStakedWeight(strategies []gethcommon.Address, multipliers []*big.Int) *big.Int {
	return Weight(os.TotalStakedShares, strategies, multipliers)
}

func (os *OperatorSet) FillTotals() {
	os.TotalStakedShares = make(map[gethcommon.Address]*big.Int)
	for _, operator := range os.Operators {
		operator.DelegatedShares = make(map[gethcommon.Address]*big.Int)
		for _, staker := range operator.Stakers {
			for strategy, shares := range staker.Shares {
				if _, found := operator.DelegatedShares[strategy]; !found {
					operator.DelegatedShares[strategy] = big.NewInt(0)
				}
				if _, found := os.TotalStakedShares[strategy]; !found {
					os.TotalStakedShares[strategy] = big.NewInt(0)
				}

				operator.DelegatedShares[strategy].Add(operator.DelegatedShares[strategy], shares)
				os.TotalStakedShares[strategy].Add(os.TotalStakedShares[strategy], shares)
			}
		}
	}
}

func (os *OperatorSet) ModifyShares(operatorAddress, stakerAddress, strategyAddress gethcommon.Address, newShares *big.Int) {
	for _, operator := range os.Operators {
		if operator.Address == operatorAddress {
			for _, staker := range operator.Stakers {
				if staker.Address == stakerAddress {
					staker.Shares[strategyAddress] = newShares
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
