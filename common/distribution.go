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

// Set sets the value for a given address.
func (d *Distribution) Set(address gethcommon.Address, amount *big.Int) {
	d.data[address] = amount
}

// Get gets the value for a given address.
func (d *Distribution) Get(address gethcommon.Address) *big.Int {
	return d.data[address]
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

func NewRandomDistribution(numAddrs int) *Distribution {
	d := NewDistribution()
	for i := 0; i < numAddrs; i++ {
		d.Set(gethcommon.BigToAddress(big.NewInt(int64(i))), big.NewInt(int64(i)))
	}
	return d
}

// Merklizes the distribution and returns the merkle root.
func (d *Distribution) Merklize(token gethcommon.Address) ([32]byte, error) {
	// todo: parallelize this
	leafs := make([][]byte, len(d.data))
	for address, amount := range d.data {
		leafs = append(leafs, encodeLeaf(address, token, amount))
	}

	return Merklize(leafs)
}

// encodeLeaf encodes an address and an amount into a leaf.
func encodeLeaf(address, token gethcommon.Address, amount *big.Int) []byte {
	// todo: handle this better
	amountU256, _ := uint256.FromBig(amount)
	amountBytes := amountU256.Bytes32()
	// (address || token || amount)
	return append(append(address.Bytes(), token.Bytes()...), amountBytes[:]...)
}
