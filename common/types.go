package common

import (
	"math/big"

	gethcommon "github.com/ethereum/go-ethereum/common"
)

type OperatorSet struct {
	TotalStakedStrategyShares map[gethcommon.Address]*big.Int
	Operators                 []Operator
}

type Earner struct {
	Claimer gethcommon.Address
}

type Operator struct {
	Earner
	Address                      gethcommon.Address
	Commissions                  map[gethcommon.Address]*big.Int
	TotalDelegatedStrategyShares map[gethcommon.Address]*big.Int
	Stakers                      []Staker
}

type Staker struct {
	Earner
	Address gethcommon.Address
	Shares  map[gethcommon.Address]*big.Int
}
