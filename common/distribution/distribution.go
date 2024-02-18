package distribution

import (
	"encoding/json"
	"fmt"
	"math/big"

	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/holiman/uint256"
)

var ZERO_LEAF [32]byte

type Distribution struct {
	data map[gethcommon.Address]map[gethcommon.Address]*big.Int
}

func NewDistribution() *Distribution {
	data := make(map[gethcommon.Address]map[gethcommon.Address]*big.Int)
	return &Distribution{
		data: data,
	}
}

// Set sets the value for a given address.
func (d *Distribution) Set(address, token gethcommon.Address, amount *big.Int) {
	if len(d.data) == 0 {
		d.data = make(map[gethcommon.Address]map[gethcommon.Address]*big.Int)
	}
	if len(d.data[address]) == 0 {
		d.data[address] = make(map[gethcommon.Address]*big.Int)
	}
	d.data[address][token] = amount
}

// Get gets the value for a given address.
func (d *Distribution) Get(address, token gethcommon.Address) *big.Int {
	if len(d.data) == 0 {
		return big.NewInt(0)
	}
	if len(d.data[address]) == 0 {
		return big.NewInt(0)
	}
	amount := d.data[address][token]
	if amount == nil {
		return big.NewInt(0)
	}
	return amount
}

// Add adds the other distribution to this distribution.
func (d *Distribution) Add(other *Distribution) {
	for address, tokenAmts := range other.data {
		for token, amount := range tokenAmts {
			if d.data[address][token] == nil {
				d.data[address][token] = big.NewInt(0)
			}
			d.data[address][token].Add(d.data[address][token], amount)
		}
	}
}

func (d *Distribution) MarshalJSON() ([]byte, error) {
	// dereference the big.Ints
	data := make(map[gethcommon.Address]map[gethcommon.Address]string)
	for address, tokenAmts := range d.data {
		data[address] = make(map[gethcommon.Address]string)
		for token, amt := range tokenAmts {
			data[address][token] = amt.String()
		}

	}
	return json.Marshal(data)
}

func (d *Distribution) UnmarshalJSON(data []byte) error {
	// dereference the big.Ints
	var ok bool
	var dataMap map[gethcommon.Address]map[gethcommon.Address]string
	if err := json.Unmarshal(data, &dataMap); err != nil {
		return err
	}
	d.data = make(map[gethcommon.Address]map[gethcommon.Address]*big.Int)
	for address, tokenAmts := range dataMap {
		for token, amt := range tokenAmts {
			d.data[address][token], ok = new(big.Int).SetString(amt, 10)
			if !ok {
				return fmt.Errorf("failed to parse big.Int from string: %s", amt)
			}
		}
	}
	return nil
}

func (d *Distribution) GetNumLeaves() int {
	numLeaves := 0
	for _, tokenAmts := range d.data {
		numLeaves += len(tokenAmts)
	}
	return numLeaves
}

// Merklizes the distribution and returns the merkle root.
func (d *Distribution) Merklize(merklizeFunc func([][32]byte) ([32]byte, error)) ([32]byte, error) {
	// todo: parallelize this
	accountLeafs := make([][32]byte, len(d.data))
	for address, tokenAmts := range d.data {
		tokenLeafs := make([][32]byte, len(tokenAmts))
		for token, amount := range tokenAmts {
			tokenLeafs = append(tokenLeafs, encodeLeaf(address, token, amount))
		}
		// merklize all leaves for this address
		accountRoot, err := merklizeFunc(tokenLeafs)
		if err != nil {
			return [32]byte{}, err
		}
		// append the root to the list of leafs
		accountLeafs = append(accountLeafs, accountRoot)
	}

	return merklizeFunc(accountLeafs)
}

// encodeLeaf encodes an account and an amount into a leaf.
func encodeLeaf(account, token gethcommon.Address, amount *big.Int) [32]byte {
	// todo: handle this better
	amountU256, _ := uint256.FromBig(amount)
	amountBytes := amountU256.Bytes32()
	// (account || token || amount)
	return [32]byte(crypto.Keccak256(append(append(account.Bytes(), token.Bytes()...), amountBytes[:]...)))
}
