package common

import (
	"math/big"

	gethcommon "github.com/ethereum/go-ethereum/common"
)

type OperatorSet struct {
	TotalStakedStrategyShares *big.Int
	Operators                 []Operator
}

type Earner struct {
	Claimer gethcommon.Address
}

type Operator struct {
	Earner
	Address                      gethcommon.Address
	Commission                   *big.Int
	TotalDelegatedStrategyShares *big.Int
	Stakers                      []Staker
}

type Staker struct {
	Earner
	Address        gethcommon.Address
	StrategyShares *big.Int
}
