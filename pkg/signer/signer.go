package signer

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type Signer interface {
	GetTransactOpts(chainId *big.Int) (*bind.TransactOpts, error)
	GetAddress() common.Address
}
