package claimprover

import (
	"github.com/Layr-Labs/eigenlayer-payment-updater/common/distribution"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/wealdtech/go-merkletree/v2"
)

type ClaimProver struct {
	accountTree             *merkletree.MerkleTree
	tokenTrees              map[gethcommon.Address]*merkletree.MerkleTree
	distribution            *distribution.Distribution
	distributionDataService distribution.DistributionDataService
}

func NewClaimProver(distributionDataService distribution.DistributionDataService) *ClaimProver {
	// load latest root from chain and trees into mem

	return &ClaimProver{
		distributionDataService: distributionDataService,
	}
}
