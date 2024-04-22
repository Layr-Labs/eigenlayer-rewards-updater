package distribution_test

import (
	"encoding/hex"
	"math/big"
	"testing"

	"github.com/Layr-Labs/eigenlayer-payment-updater/common/distribution"
	"github.com/Layr-Labs/eigenlayer-payment-updater/common/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func FuzzSetAndGet(f *testing.F) {
	f.Add([]byte{69}, []byte{42, 0}, uint64(69420))

	f.Fuzz(func(t *testing.T, addressBytes, tokenBytes []byte, amounUintFuzz uint64) {
		address := common.Address{}
		address.SetBytes(addressBytes)

		token := common.Address{}
		token.SetBytes(tokenBytes)

		amount := new(big.Int).SetUint64(amounUintFuzz)

		d := distribution.NewDistribution()
		err := d.Set(address, token, amount)
		assert.NoError(t, err)

		fetched, found := d.Get(address, token)
		assert.True(t, found)
		assert.Equal(t, amount, fetched)
	})
}

func TestSetNilAmount(t *testing.T) {
	d := distribution.NewDistribution()
	err := d.Set(common.Address{}, common.Address{}, nil)
	assert.NoError(t, err)

	_, found := d.Get(common.Address{}, common.Address{})
	assert.True(t, found)
}

func TestSetAddressesInNonAlphabeticalOrder(t *testing.T) {
	d := distribution.NewDistribution()

	err := d.Set(utils.TestAddresses[1], utils.TestTokens[0], big.NewInt(1))
	assert.NoError(t, err)

	err = d.Set(utils.TestAddresses[0], utils.TestTokens[0], big.NewInt(2))
	assert.ErrorIs(t, err, distribution.ErrAddressNotInOrder)

	amount1, found := d.Get(utils.TestAddresses[1], utils.TestTokens[0])
	assert.Equal(t, big.NewInt(1), amount1)
	assert.True(t, found)

	amount2, found := d.Get(utils.TestAddresses[0], utils.TestTokens[0])
	assert.Equal(t, big.NewInt(0), amount2)
	assert.False(t, found)
}

func TestSetTokensInNonAlphabeticalOrder(t *testing.T) {
	d := distribution.NewDistribution()

	err := d.Set(utils.TestAddresses[0], utils.TestTokens[1], big.NewInt(1))
	assert.NoError(t, err)

	err = d.Set(utils.TestAddresses[0], utils.TestTokens[0], big.NewInt(2))
	assert.ErrorIs(t, err, distribution.ErrTokenNotInOrder)

	amount1, found := d.Get(utils.TestAddresses[0], utils.TestTokens[1])
	assert.Equal(t, big.NewInt(1), amount1)
	assert.True(t, found)

	amount2, found := d.Get(utils.TestAddresses[0], utils.TestTokens[0])
	assert.Equal(t, big.NewInt(0), amount2)
	assert.False(t, found)
}

func TestGetUnset(t *testing.T) {
	d := distribution.NewDistribution()

	fetched, found := d.Get(utils.TestAddresses[0], utils.TestTokens[0])
	assert.Equal(t, big.NewInt(0), fetched)
	assert.False(t, found)
}

func TestEncodeAccountLeaf(t *testing.T) {
	for i := 0; i < len(utils.TestAddresses); i++ {
		testRoot, _ := hex.DecodeString(utils.TestRootsString[i])
		leaf := distribution.EncodeAccountLeaf(utils.TestAddresses[i], testRoot)
		assert.Equal(t, distribution.EARNER_LEAF_SALT[0], leaf[0], "The first byte of the leaf should be EARNER_LEAF_SALT")
		assert.Equal(t, utils.TestAddresses[i][:], leaf[1:21])
		assert.Equal(t, testRoot, leaf[21:])
	}
}

func TestEncodeTokenLeaf(t *testing.T) {
	for i := 0; i < len(utils.TestTokens); i++ {
		testAmount, _ := new(big.Int).SetString(utils.TestAmountsString[i], 10)
		leaf := distribution.EncodeTokenLeaf(utils.TestTokens[i], testAmount)
		assert.Equal(t, distribution.TOKEN_LEAF_SALT[0], leaf[0], "The first byte of the leaf should be TOKEN_LEAF_SALT")
		assert.Equal(t, utils.TestTokens[i][:], leaf[1:21])
		assert.Equal(t, utils.TestAmountsBytes32[i], hex.EncodeToString(leaf[21:]))
	}
}

func TestGetAccountIndexBeforeMerklization(t *testing.T) {
	d := utils.GetTestDistribution()

	accountIndex, found := d.GetAccountIndex(utils.TestAddresses[1])
	assert.False(t, found)
	assert.Equal(t, uint64(0), accountIndex)
}

func TestGetTokenIndexBeforeMerklization(t *testing.T) {
	d := utils.GetTestDistribution()

	tokenIndex, found := d.GetTokenIndex(utils.TestAddresses[1], utils.TestTokens[1])
	assert.False(t, found)
	assert.Equal(t, uint64(0), tokenIndex)
}

func TestMerklize(t *testing.T) {
	d := utils.GetTestDistribution()

	accountTree, tokenTrees, err := d.Merklize()
	assert.NoError(t, err)

	// check the token trees
	assert.Len(t, tokenTrees, len(utils.TestAddresses))
	for i := 0; i < len(tokenTrees); i++ {
		tokenTree, found := tokenTrees[utils.TestAddresses[i]]
		assert.True(t, found)
		assert.Len(t, tokenTree.Data, len(utils.TestTokens)-i)

		// check the data, that means the leafs are the same
		for j := 0; j < len(utils.TestTokens)-i; j++ {
			leaf := tokenTree.Data[j]
			assert.Equal(t, distribution.EncodeTokenLeaf(utils.TestTokens[j], big.NewInt(int64(j+i+1))), leaf)
		}
	}

	// check the account tree
	assert.Len(t, accountTree.Data, len(utils.TestAddresses))
	for i := 0; i < len(utils.TestAddresses); i++ {
		accountRoot := tokenTrees[utils.TestAddresses[i]].Root()
		leaf := accountTree.Data[i]
		assert.Equal(t, distribution.EncodeAccountLeaf(utils.TestAddresses[i], accountRoot), leaf)

		accountIndex, found := d.GetAccountIndex(utils.TestAddresses[i])
		assert.True(t, found)
		assert.Equal(t, uint64(i), accountIndex)

		for j := 0; j < len(utils.TestTokens)-i; j++ {
			tokenIndex, found := d.GetTokenIndex(utils.TestAddresses[i], utils.TestTokens[j])
			assert.True(t, found)
			assert.Equal(t, uint64(j), tokenIndex)
		}
	}
}
