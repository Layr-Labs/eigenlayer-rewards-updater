package claimgen

import (
	"context"
	paymentCoordinator "github.com/Layr-Labs/eigenlayer-payment-updater/pkg/bindings/IPaymentCoordinator"
	"github.com/Layr-Labs/eigenlayer-payment-updater/pkg/distribution"
	"github.com/Layr-Labs/eigenlayer-payment-updater/pkg/utils"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/wealdtech/go-merkletree/v2"
)

type IPaymentCoordinatorEarnerTreeMerkleLeafStrings struct {
	Earner          gethcommon.Address
	EarnerTokenRoot string
}

type IPaymentCoordinatorPaymentMerkleClaimStrings struct {
	Root               string
	RootIndex          uint32
	EarnerIndex        uint32
	EarnerTreeProof    string
	EarnerLeaf         IPaymentCoordinatorEarnerTreeMerkleLeafStrings
	LeafIndices        []uint32
	TokenTreeProofs    []string
	TokenLeaves        []paymentCoordinator.IPaymentCoordinatorTokenTreeMerkleLeaf
	TokenTreeProofsNum uint32
	TokenLeavesNum     uint32
}

func GenerateClaimProofForEarner(
	ctx context.Context,
	earner gethcommon.Address,
	tokens []gethcommon.Address,
	rootIndex uint32,
	distribution *distribution.Distribution,
) (
	*merkletree.MerkleTree,
	*paymentCoordinator.IPaymentCoordinatorPaymentMerkleClaim,
	error,
) {
	accountTree, tokenTrees, err := distribution.Merklize()
	if err != nil {
		return nil, nil, err
	}

	merkleClaim, err := GetProof(
		distribution,
		rootIndex,
		accountTree,
		tokenTrees,
		earner,
		tokens,
	)

	if err != nil {
		return nil, nil, err
	}

	return accountTree, merkleClaim, err
}

func FormatProofForSolidity(accountTreeRoot []byte, proof *paymentCoordinator.IPaymentCoordinatorPaymentMerkleClaim) *IPaymentCoordinatorPaymentMerkleClaimStrings {
	return &IPaymentCoordinatorPaymentMerkleClaimStrings{
		Root:            utils.ConvertBytesToString(accountTreeRoot),
		RootIndex:       proof.RootIndex,
		EarnerIndex:     proof.EarnerIndex,
		EarnerTreeProof: utils.ConvertBytesToString(proof.EarnerTreeProof),
		EarnerLeaf: IPaymentCoordinatorEarnerTreeMerkleLeafStrings{
			Earner:          proof.EarnerLeaf.Earner,
			EarnerTokenRoot: utils.ConvertBytes32ToString(proof.EarnerLeaf.EarnerTokenRoot),
		},
		LeafIndices:        proof.TokenIndices,
		TokenTreeProofs:    utils.ConvertBytesToStrings(proof.TokenTreeProofs),
		TokenLeaves:        proof.TokenLeaves,
		TokenTreeProofsNum: uint32(len(proof.TokenTreeProofs)),
		TokenLeavesNum:     uint32(len(proof.TokenLeaves)),
	}
}
