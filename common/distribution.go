package common

import (
	"math/big"

	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/holiman/uint256"
)

var ZERO_LEAF = make([]byte, 20+32)

type Distribution struct {
	data map[gethcommon.Address]*big.Int
}

func NewDistribution() *Distribution {
	return &Distribution{
		data: make(map[gethcommon.Address]*big.Int),
	}
}

// Add adds the other distribution to this distribution.
func (d *Distribution) Add(other *Distribution) {
	for address, amount := range other.data {
		if d.data[address] == nil {
			d.data[address] = big.NewInt(0)
		}
		d.data[address].Add(d.data[address], amount)
	}
}

// MulDiv multiplies the distribution by a numerator and divides by a denominator.
func (d *Distribution) MulDiv(numerator, denominator *big.Int) {
	for address, amount := range d.data {
		amount.Mul(amount, numerator)
		amount.Div(amount, denominator)
		d.data[address] = amount
	}
}

// Merklizes the distribution and returns the merkle root.
func (d *Distribution) Merklize() ([32]byte, error) {
	// todo: parallelize this
	leafs := make([][]byte, len(d.data))
	for address, amount := range d.data {
		leafs = append(leafs, encodeLeaf(address, amount))
	}

	return Merklize(leafs)
}

// encodeLeaf encodes an address and an amount into a leaf.
func encodeLeaf(address gethcommon.Address, amount *big.Int) []byte {
	// todo: handle this better
	amountU256, _ := uint256.FromBig(amount)
	amountBytes := amountU256.Bytes32()
	return append(address.Bytes(), amountBytes[:]...)
}
