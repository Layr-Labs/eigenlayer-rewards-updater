package distribution

import (
	"fmt"
	"math/big"

	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/holiman/uint256"
	"github.com/wealdtech/go-merkletree/v2"
	"github.com/wealdtech/go-merkletree/v2/keccak256"
	orderedmap "github.com/wk8/go-ordered-map/v2"
)

// Used for marshalling and unmarshalling big integers.
type BigInt struct {
	*big.Int
}

func (b BigInt) MarshalJSON() ([]byte, error) {
	return []byte(b.String()), nil
}

func (b *BigInt) UnmarshalJSON(p []byte) error {
	if string(p) == "null" {
		return nil
	}
	var z big.Int
	_, ok := z.SetString(string(p), 10)
	if !ok {
		return fmt.Errorf("not a valid big integer: %s", p)
	}
	b.Int = &z
	return nil
}

type Distribution struct {
	accountIndices map[gethcommon.Address]uint64                        // used for optimizing proving
	tokenIndices   map[gethcommon.Address]map[gethcommon.Address]uint64 // used for optimizing proving
	data           *orderedmap.OrderedMap[gethcommon.Address, *orderedmap.OrderedMap[gethcommon.Address, *BigInt]]
}

func NewDistribution() *Distribution {
	data := orderedmap.New[gethcommon.Address, *orderedmap.OrderedMap[gethcommon.Address, *BigInt]]()
	return &Distribution{
		data: data,
	}
}

// Set sets the value for a given address.
func (d *Distribution) Set(address, token gethcommon.Address, amount *big.Int) {
	allocatedTokens, found := d.data.Get(address)
	if !found {
		allocatedTokens = orderedmap.New[gethcommon.Address, *BigInt]()
		d.data.Set(address, allocatedTokens)
	}
	allocatedTokens.Set(token, &BigInt{Int: amount})
}

// Get gets the value for a given address and whether it was in the distribution
func (d *Distribution) Get(address, token gethcommon.Address) (*big.Int, bool) {
	allocatedTokens, found := d.data.Get(address)
	if !found {
		return big.NewInt(0), false
	}
	amount, found := allocatedTokens.Get(token)
	if !found {
		return big.NewInt(0), false
	}
	return amount.Int, true
}

// Sets the index of the account in the distribution
func (d *Distribution) SetAccountIndex(address gethcommon.Address, index uint64) {
	if d.accountIndices == nil {
		d.accountIndices = make(map[gethcommon.Address]uint64)
	}

	d.accountIndices[address] = index
}

// Gets the index of the account in the distribution
// Note that the indices must be set before calling this function
func (d *Distribution) GetAccountIndex(address gethcommon.Address) (uint64, bool) {
	index, found := d.accountIndices[address]
	return index, found
}

// Sets the index of the token for a certain account in the distribution
func (d *Distribution) SetTokenIndex(address, token gethcommon.Address, index uint64) {
	if d.tokenIndices == nil {
		d.tokenIndices = make(map[gethcommon.Address]map[gethcommon.Address]uint64)
	}

	indices, found := d.tokenIndices[address]
	if !found {
		indices = make(map[gethcommon.Address]uint64)
		d.tokenIndices[address] = indices
	}

	indices[token] = index
}

// Gets the index of the token for a certain account in the distribution
// Note that the indices must be set before calling this function
func (d *Distribution) GetTokenIndex(address, token gethcommon.Address) (uint64, bool) {
	indices, found := d.tokenIndices[address]
	if !found {
		return 0, false
	}

	index, found := indices[token]
	return index, found
}

// Add adds the other distribution to this distribution.
// assumes other is non nil
func (d *Distribution) Add(other *Distribution) {
	for accountPair := other.data.Oldest(); accountPair != nil; accountPair = accountPair.Next() {
		address := accountPair.Key
		for tokenPair := accountPair.Value.Oldest(); tokenPair != nil; tokenPair = tokenPair.Next() {
			token := tokenPair.Key
			amount := tokenPair.Value
			currentAmount, _ := d.Get(address, token)
			d.Set(address, token, currentAmount.Add(currentAmount, amount.Int))
		}
	}
}

func (d *Distribution) GetNumLeaves() int {
	numLeaves := 0
	for accountPair := d.data.Oldest(); accountPair != nil; accountPair = accountPair.Next() {
		numLeaves += accountPair.Value.Len()
	}
	return numLeaves
}

func (d *Distribution) MarshalJSON() ([]byte, error) {
	return d.data.MarshalJSON()
}

func (d *Distribution) UnmarshalJSON(p []byte) error {
	data := orderedmap.New[gethcommon.Address, *orderedmap.OrderedMap[gethcommon.Address, *BigInt]]()
	err := data.UnmarshalJSON(p)
	if err != nil {
		return err
	}
	d.data = data
	return nil
}

// Merklizes the distribution and returns the account tree and the token trees.
// TODO: Do we need to have an option to merklize without all returning all the token trees and data?
func (d *Distribution) Merklize() (*merkletree.MerkleTree, map[gethcommon.Address]*merkletree.MerkleTree, error) {
	tokenTrees := make(map[gethcommon.Address]*merkletree.MerkleTree, d.data.Len())

	// todo: parallelize this
	accountIndex := uint64(0)
	accountLeafs := make([][]byte, d.data.Len())
	for accountPair := d.data.Oldest(); accountPair != nil; accountPair = accountPair.Next() {
		address := accountPair.Key
		d.SetAccountIndex(address, accountIndex)
		// fetch the leafs for the tokens for this account
		tokenIndex := uint64(0)
		tokenLeafs := make([][]byte, accountPair.Value.Len())
		for tokenPair := accountPair.Value.Oldest(); tokenPair != nil; tokenPair = tokenPair.Next() {
			token := tokenPair.Key
			amount := tokenPair.Value
			d.SetTokenIndex(address, token, tokenIndex)
			tokenLeafs = append(tokenLeafs, EncodeTokenLeaf(token, amount.Int))
			tokenIndex++
		}

		// create a merkle tree for the tokens for this account
		tokenTree, err := merkletree.NewTree(
			merkletree.WithData(tokenLeafs),
			merkletree.WithHashType(keccak256.New()),
		)
		if err != nil {
			return nil, nil, err
		}
		tokenTrees[address] = tokenTree

		// append the root to the list of account leafs
		accountRoot := tokenTree.Root()
		accountLeafs = append(accountLeafs, EncodeAccountLeaf(address, accountRoot))
		accountIndex++
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
func EncodeAccountLeaf(account gethcommon.Address, accountRoot []byte) []byte {
	// (account || accountRoot)
	return append(account.Bytes(), accountRoot[:]...)
}

// encodeTokenLeaf encodes a token leaf for a token distribution.
func EncodeTokenLeaf(token gethcommon.Address, amount *big.Int) []byte {
	// todo: handle this better
	amountU256, _ := uint256.FromBig(amount)
	amountBytes := amountU256.Bytes32()
	// (token || amount)
	return append(token.Bytes(), amountBytes[:]...)
}
