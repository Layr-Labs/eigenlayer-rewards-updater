package distribution

import (
	"encoding/json"
	"fmt"
	"math/big"

	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/holiman/uint256"
	"github.com/wealdtech/go-merkletree/v2"
	"github.com/wealdtech/go-merkletree/v2/keccak256"
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
	// TODO: fix this
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
	// TODO: fix this
	if len(d.data) == 0 {
		d.data = make(map[gethcommon.Address]map[gethcommon.Address]*big.Int)
	}
	if len(other.data) == 0 {
		other.data = make(map[gethcommon.Address]map[gethcommon.Address]*big.Int)
	}

	for address, tokenAmts := range other.data {
		for token, amount := range tokenAmts {
			if d.data[address][token] == nil {
				// TODO: fix this
				if len(d.data[address]) == 0 {
					d.data[address] = make(map[gethcommon.Address]*big.Int)
				}
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
			if d.data[address] == nil {
				d.data[address] = make(map[gethcommon.Address]*big.Int)
			}
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

// Merklizes the distribution and returns the account tree and the token trees.
func (d *Distribution) Merklize() (*merkletree.MerkleTree, []*merkletree.MerkleTree, error) {
	tokenTrees := make([]*merkletree.MerkleTree, 0)

	// todo: parallelize this
	accountLeafs := make([][]byte, len(d.data))
	for address, tokenAmts := range d.data {
		tokenLeafs := make([][]byte, len(tokenAmts))
		for token, amount := range tokenAmts {
			tokenLeafs = append(tokenLeafs, encodeTokenLeaf(token, amount))
		}

		// create a merkle tree for the tokens for this account
		tokenTree, err := merkletree.NewTree(
			merkletree.WithData(tokenLeafs),
			merkletree.WithHashType(keccak256.New()),
		)
		if err != nil {
			return nil, nil, err
		}
		tokenTrees = append(tokenTrees, tokenTree)

		// append the root to the list of account leafs
		accountRoot := tokenTree.Root()
		accountLeafs = append(accountLeafs, encodeAccountLeaf(address, accountRoot))
	}

	accountTree, err := merkletree.NewTree(
		merkletree.WithData(accountLeafs),
		merkletree.WithHashType(keccak256.New()),
	)
	if err != nil {
		return nil, nil, err
	}

	return accountTree, tokenTrees, nil
}

// encodeAccountLeaf encodes an account leaf for a token distribution.
func encodeAccountLeaf(account gethcommon.Address, accountRoot []byte) []byte {
	// (account || accountRoot)
	return append(account.Bytes(), accountRoot[:]...)
}

// encodeTokenLeaf encodes a token leaf for a token distribution.
func encodeTokenLeaf(token gethcommon.Address, amount *big.Int) []byte {
	// todo: handle this better
	amountU256, _ := uint256.FromBig(amount)
	amountBytes := amountU256.Bytes32()
	// (token || amount)
	return append(token.Bytes(), amountBytes[:]...)
}
