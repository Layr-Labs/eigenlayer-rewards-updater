package claimprover

import (
	"context"
	"fmt"
	"sync"
	"time"

	paymentCoordinator "github.com/Layr-Labs/eigenlayer-payment-updater/bindings/IPaymentCoordinator"
	"github.com/Layr-Labs/eigenlayer-payment-updater/common/distribution"
	"github.com/Layr-Labs/eigenlayer-payment-updater/common/services"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/rs/zerolog/log"
	"github.com/wealdtech/go-merkletree/v2"
)

type ClaimProofs struct {
	AccountProof *merkletree.Proof
	TokenProofs  []*merkletree.Proof
}

type ClaimProver struct {
	updateInterval time.Duration

	transactor services.Transactor

	distributionDataService services.DistributionDataService
	distribution            *distribution.Distribution

	rootIndex   uint32
	accountTree *merkletree.MerkleTree
	tokenTrees  map[gethcommon.Address]*merkletree.MerkleTree

	mu sync.RWMutex
}

func NewClaimProver(updateIntervalSeconds int64, transactor services.Transactor, distributionDataService services.DistributionDataService) *ClaimProver {

	claimProver := &ClaimProver{
		updateInterval:          time.Second * time.Duration(updateIntervalSeconds),
		transactor:              transactor,
		distributionDataService: distributionDataService,
		mu:                      sync.RWMutex{},
	}

	return claimProver
}

func (cp *ClaimProver) Start() {
	// run a loop unning once every u.UpdateInterval that calls u.update()
	log.Info().Msg("service started")
	ctx := context.Background()

	// run the first update immediately
	if err := cp.update(ctx); err != nil {
		log.Error().Msgf("failed to update: %s", err)
	}

	ticker := time.NewTicker(cp.updateInterval)
	for range ticker.C {
		log.Info().Msg("running update")
		if err := cp.update(ctx); err != nil {
			log.Error().Msgf("failed to update: %s", err)
		}
	}
}

func (cp *ClaimProver) update(ctx context.Context) error {
	// get latest submitted distribution
	distribution, _, err := cp.distributionDataService.GetLatestSubmittedDistribution(ctx)
	if err != nil {
		return err
	}

	// aquire write lock
	cp.mu.Lock()

	// get the distribution for the root
	cp.distribution = distribution

	// generate the trees
	cp.accountTree, cp.tokenTrees, err = cp.distribution.Merklize()
	if err != nil {
		return err
	}

	var root [32]byte
	copy(root[:], cp.accountTree.Root())

	cp.rootIndex, err = cp.transactor.GetRootIndex(root)
	if err != nil {
		return err
	}

	cp.mu.Unlock()

	return nil
}

func (cp *ClaimProver) GetProof(earner gethcommon.Address, tokens []gethcommon.Address) (*paymentCoordinator.IPaymentCoordinatorPaymentMerkleClaim, error) {
	// aquire read lock
	cp.mu.RLock()

	earnerIndex, found := cp.distribution.GetAccountIndex(earner)
	if !found {
		return nil, fmt.Errorf("account index not found for recipient %s", earner.Hex())
	}

	// get the token proofs
	tokenIndices := make([]uint32, len(tokens))
	tokenProofsBytes := make([][]byte, len(tokens))
	tokenLeaves := make([]paymentCoordinator.IPaymentCoordinatorTokenTreeMerkleLeaf, len(tokens))
	for i, token := range tokens {
		tokenIndex, found := cp.distribution.GetTokenIndex(earner, token)
		if !found {
			return nil, fmt.Errorf("token index not found for token %s and recipient %s", token.Hex(), earner.Hex())
		}
		tokenIndices[i] = uint32(tokenIndex)

		tokenProof, err := cp.tokenTrees[earner].GenerateProofWithIndex(tokenIndex, 0)
		if err != nil {
			return nil, err
		}
		tokenProofsBytes[i] = flattenHashes(tokenProof.Hashes)

		amount, found := cp.distribution.Get(earner, token)
		if !found {
			// this should never happen due to the token index check above
			return nil, fmt.Errorf("amount not found for token %s and recipient %s", token.Hex(), earner.Hex())
		}

		tokenLeaves[i] = paymentCoordinator.IPaymentCoordinatorTokenTreeMerkleLeaf{
			Token:              token,
			CumulativeEarnings: amount,
		}
	}

	var earnerRoot [32]byte
	copy(earnerRoot[:], cp.tokenTrees[earner].Root())

	// get the account proof
	earnerTreeProof, err := cp.accountTree.GenerateProofWithIndex(earnerIndex, 0)
	if err != nil {
		log.Error().Msgf("failed to generate account proof: %s", err)
	}

	earnerTreeProofBytes := flattenHashes(earnerTreeProof.Hashes)

	cp.mu.RUnlock()

	return &paymentCoordinator.IPaymentCoordinatorPaymentMerkleClaim{
		RootIndex:       cp.rootIndex,
		EarnerIndex:     uint32(earnerIndex),
		EarnerTreeProof: earnerTreeProofBytes,
		EarnerLeaf: paymentCoordinator.IPaymentCoordinatorEarnerTreeMerkleLeaf{
			Earner:          earner,
			EarnerTokenRoot: earnerRoot,
		},
		LeafIndices:     tokenIndices,
		TokenTreeProofs: tokenProofsBytes,
		TokenLeaves:     tokenLeaves,
	}, nil
}

func flattenHashes(hashes [][]byte) []byte {
	result := make([]byte, 0)
	for i := 0; i < len(hashes); i++ {
		result = append(result, hashes[i]...)
	}
	return result
}
