package common

import (
	"encoding/json"
	"math/big"

	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/holiman/uint256"
)

var ZERO_LEAF = make([]byte, 20+32)

type Distribution struct {
	data map[gethcommon.Address]*big.Int
}

func NewDistribution() *Distribution {
	data := make(map[gethcommon.Address]*big.Int)
	return &Distribution{
		data: data,
	}
}

// Set sets the value for a given address.
func (d *Distribution) Set(address gethcommon.Address, amount *big.Int) {
	if len(d.data) == 0 {
		d.data = make(map[gethcommon.Address]*big.Int)
	}
	d.data[address] = amount
}

// Get gets the value for a given address.
func (d *Distribution) Get(address gethcommon.Address) *big.Int {
	amt := d.data[address]
	if amt == nil {
		return big.NewInt(0)
	}
	return amt
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

func (d *Distribution) MarshalJSON() ([]byte, error) {
	// dereference the big.Ints
	data := make(map[gethcommon.Address]string)
	for address, amount := range d.data {
		data[address] = amount.String()
	}
	return json.Marshal(data)
}

func (d *Distribution) UnmarshalJSON(data []byte) error {
	// dereference the big.Ints
	var dataMap map[gethcommon.Address]string
	if err := json.Unmarshal(data, &dataMap); err != nil {
		return err
	}
	d.data = make(map[gethcommon.Address]*big.Int)
	for address, amount := range dataMap {
		d.data[address] = new(big.Int)
		d.data[address].SetString(amount, 10)
	}
	return nil
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
