package claimprover_test

import (
	"context"
	"testing"

	paymentCoordinator "github.com/Layr-Labs/eigenlayer-payment-updater/bindings/IPaymentCoordinator"
	claimprover "github.com/Layr-Labs/eigenlayer-payment-updater/claimgen"
	"github.com/Layr-Labs/eigenlayer-payment-updater/common/distribution"
	"github.com/Layr-Labs/eigenlayer-payment-updater/common/services/mocks"
	"github.com/Layr-Labs/eigenlayer-payment-updater/common/utils"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/wealdtech/go-merkletree/v2"
	"github.com/wealdtech/go-merkletree/v2/keccak256"
)

var testRootIndex uint32 = 4007
var testTimestamp int64 = 1712127631

var testUpdateIntervalSeconds int64 = 100

func TestClaimProverUpdate(t *testing.T) {
	_, accountTree, tokenTrees, rootBytes, cp := createUpdatableClaimProver()

	cp.Update(context.Background())

	fetchedAccountTrees, fetchedTokenTrees, err := cp.Distribution.Merklize()

	assert.Nil(t, err)

	// make sure the expected token trees are the same as thos from merlizing the cp's distribution and the CPs cached trees
	assert.Equal(t, accountTree.Root(), fetchedAccountTrees.Root())
	assert.Equal(t, len(tokenTrees), len(fetchedTokenTrees))
	assert.Equal(t, len(tokenTrees), len(cp.TokenTrees))
	for earner, tree := range tokenTrees {
		fetchedTree := fetchedTokenTrees[earner]
		assert.Equal(t, tree.Root(), fetchedTree.Root())
		assert.Equal(t, tree.Root(), cp.TokenTrees[earner].Root())
	}

	assert.Equal(t, rootBytes, cp.AccountTree.Root())
	assert.Equal(t, testRootIndex, cp.RootIndex)
}

func TestClaimProverGetProof(t *testing.T) {
	d, _, tokenTrees, rootBytes, cp := createUpdatableClaimProver()

	cp.Update(context.Background())

	claim, err := cp.GetProof(utils.TestAddresses[0], []gethcommon.Address{utils.TestTokens[0], utils.TestTokens[3]})
	assert.Nil(t, err)

	assert.Equal(t, testRootIndex, claim.RootIndex)
	verifyEarner(t, rootBytes, tokenTrees, 0, claim)
	verifyTokens(t, d, []int{0, 3}, claim)
}

func TestClaimProverGetProofDecreasingTokenOrder(t *testing.T) {
	d, _, tokenTrees, rootBytes, cp := createUpdatableClaimProver()

	cp.Update(context.Background())

	claim, err := cp.GetProof(utils.TestAddresses[2], []gethcommon.Address{utils.TestTokens[2], utils.TestTokens[0]})
	assert.Nil(t, err)

	assert.Equal(t, testRootIndex, claim.RootIndex)
	verifyEarner(t, rootBytes, tokenTrees, 2, claim)
	verifyTokens(t, d, []int{2, 0}, claim)
}

func TestClaimProverGetProofForNonExistantEarner(t *testing.T) {
	_, _, _, _, cp := createUpdatableClaimProver()

	cp.Update(context.Background())

	_, err := cp.GetProof(utils.TestTokens[0], []gethcommon.Address{utils.TestTokens[0]})
	assert.ErrorIs(t, err, claimprover.ErrEarnerIndexNotFound)
}

func TestClaimProverGetProofForNonExistantToken(t *testing.T) {
	_, _, _, _, cp := createUpdatableClaimProver()

	cp.Update(context.Background())

	_, err := cp.GetProof(utils.TestAddresses[0], []gethcommon.Address{utils.TestAddresses[0]})
	assert.ErrorIs(t, err, claimprover.ErrTokenIndexNotFound)
}

func createUpdatableClaimProver() (*distribution.Distribution, *merkletree.MerkleTree, map[gethcommon.Address]*merkletree.MerkleTree, []byte, *claimprover.ClaimProver) {
	mockTransactor := &mocks.Transactor{}
	mockDistributionDataService := &mocks.DistributionDataService{}

	d := utils.GetTestDistribution()
	accountTree, tokenTrees, _ := d.Merklize()
	rootBytes := accountTree.Root()
	var root [32]byte
	copy(root[:], rootBytes)

	cp := claimprover.NewClaimProver(testUpdateIntervalSeconds, mockTransactor, mockDistributionDataService)

	mockDistributionDataService.On("GetLatestSubmittedDistribution", mock.Anything).Return(d, testTimestamp, nil)
	mockTransactor.On("GetRootIndex", root).Return(testRootIndex, nil)

	return d, accountTree, tokenTrees, rootBytes, cp
}

func verifyEarner(t *testing.T, rootBytes []byte, tokenTrees map[gethcommon.Address]*merkletree.MerkleTree, testAddressIndex int, claim *paymentCoordinator.IPaymentCoordinatorPaymentMerkleClaim) {
	assert.Equal(t, uint32(testAddressIndex), claim.EarnerIndex)
	assert.Equal(t, utils.TestAddresses[claim.EarnerIndex], claim.EarnerLeaf.Earner)
	assert.Equal(t, tokenTrees[claim.EarnerLeaf.Earner].Root(), claim.EarnerLeaf.EarnerTokenRoot[:])

	// verify the earner proof
	verified, err := merkletree.VerifyProofUsing(
		distribution.EncodeAccountLeaf(claim.EarnerLeaf.Earner, claim.EarnerLeaf.EarnerTokenRoot[:]),
		false,
		getProofFromBytesAndIndex(claim.EarnerTreeProof, claim.EarnerIndex),
		[][]byte{rootBytes},
		keccak256.New(),
	)
	assert.Nil(t, err)
	assert.True(t, verified)
}

func verifyTokens(t *testing.T, d *distribution.Distribution, testTokenIndices []int, claim *paymentCoordinator.IPaymentCoordinatorPaymentMerkleClaim) {
	assert.Equal(t, len(testTokenIndices), len(claim.LeafIndices))
	assert.Equal(t, len(testTokenIndices), len(claim.TokenTreeProofs))
	assert.Equal(t, len(testTokenIndices), len(claim.TokenLeaves))

	for i, index := range testTokenIndices {
		// verify index and leaf
		assert.Equal(t, uint32(index), claim.LeafIndices[i])
		assert.Equal(t, utils.TestTokens[index], claim.TokenLeaves[i].Token)

		testAmount, found := d.Get(claim.EarnerLeaf.Earner, utils.TestTokens[index])
		assert.True(t, found)

		assert.Equal(t, testAmount, claim.TokenLeaves[i].CumulativeEarnings)

		// verify the token proof
		verified, err := merkletree.VerifyProofUsing(
			distribution.EncodeTokenLeaf(utils.TestTokens[index], testAmount),
			false,
			getProofFromBytesAndIndex(claim.TokenTreeProofs[i], uint32(index)),
			[][]byte{claim.EarnerLeaf.EarnerTokenRoot[:]},
			keccak256.New(),
		)
		assert.Nil(t, err)
		assert.True(t, verified)
	}
}

func getProofFromBytesAndIndex(byteArr []byte, index uint32) *merkletree.Proof {
	hashes := make([][]byte, 0)
	for i := 0; i < len(byteArr); i += 32 {
		hashes = append(hashes, byteArr[i:i+32])
	}

	return &merkletree.Proof{
		Hashes: hashes,
		Index:  uint64(index),
	}
}
