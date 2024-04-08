package claimprover

import (
	"context"
	"errors"
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

var ErrEarnerIndexNotFound = errors.New("earner index not found")
var ErrTokenIndexNotFound = errors.New("token not found")
var ErrAmountNotFound = errors.New("amount not found")

type ClaimProofs struct {
	AccountProof *merkletree.Proof
	TokenProofs  []*merkletree.Proof
}

type ClaimProver struct {
	updateInterval time.Duration

	transactor services.Transactor

	distributionDataService services.DistributionDataService
	Distribution            *distribution.Distribution

	RootIndex   uint32
	AccountTree *merkletree.MerkleTree
	TokenTrees  map[gethcommon.Address]*merkletree.MerkleTree

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
	if err := cp.Update(ctx); err != nil {
		log.Error().Msgf("failed to update: %s", err)
	}

	ticker := time.NewTicker(cp.updateInterval)
	for range ticker.C {
		log.Info().Msg("running update")
		if err := cp.Update(ctx); err != nil {
			log.Error().Msgf("failed to update: %s", err)
		}
	}
}

func (cp *ClaimProver) Update(ctx context.Context) error {
	// get latest submitted distribution
	distribution, _, err := cp.distributionDataService.GetLatestSubmittedDistribution(ctx)
	if err != nil {
		return err
	}

	// aquire write lock
	cp.mu.Lock()

	// get the distribution for the root
	cp.Distribution = distribution

	// generate the trees
	cp.AccountTree, cp.TokenTrees, err = cp.Distribution.Merklize()
	if err != nil {
		return err
	}

	var root [32]byte
	copy(root[:], cp.AccountTree.Root())

	cp.RootIndex, err = cp.transactor.GetRootIndex(root)
	if err != nil {
		return err
	}

	cp.mu.Unlock()

	return nil
}

func (cp *ClaimProver) GetProof(earner gethcommon.Address, tokens []gethcommon.Address) (*paymentCoordinator.IPaymentCoordinatorPaymentMerkleClaim, error) {
	// aquire read lock
	cp.mu.RLock()

	earnerIndex, found := cp.Distribution.GetAccountIndex(earner)
	if !found {
		return nil, fmt.Errorf("%w for earner %s", ErrEarnerIndexNotFound, earner.Hex())
	}

	// get the token proofs
	tokenIndices := make([]uint32, len(tokens))
	tokenProofsBytes := make([][]byte, len(tokens))
	tokenLeaves := make([]paymentCoordinator.IPaymentCoordinatorTokenTreeMerkleLeaf, len(tokens))
	for i, token := range tokens {
		tokenIndex, found := cp.Distribution.GetTokenIndex(earner, token)
		if !found {
			return nil, fmt.Errorf("%w for token %s and earner %s", ErrTokenIndexNotFound, token.Hex(), earner.Hex())
		}
		tokenIndices[i] = uint32(tokenIndex)

		tokenProof, err := cp.TokenTrees[earner].GenerateProofWithIndex(tokenIndex, 0)
		if err != nil {
			return nil, err
		}
		tokenProofsBytes[i] = FlattenHashes(tokenProof.Hashes)

		amount, found := cp.Distribution.Get(earner, token)
		if !found {
			// this should never happen due to the token index check above
			return nil, fmt.Errorf("%w for token %s and earner %s", ErrAmountNotFound, token.Hex(), earner.Hex())
		}

		tokenLeaves[i] = paymentCoordinator.IPaymentCoordinatorTokenTreeMerkleLeaf{
			Token:              token,
			CumulativeEarnings: amount,
		}
	}

	var earnerRoot [32]byte
	copy(earnerRoot[:], cp.TokenTrees[earner].Root())

	// get the account proof
	earnerTreeProof, err := cp.AccountTree.GenerateProofWithIndex(earnerIndex, 0)
	if err != nil {
		return nil, err
	}

	earnerTreeProofBytes := FlattenHashes(earnerTreeProof.Hashes)

	cp.mu.RUnlock()

	return &paymentCoordinator.IPaymentCoordinatorPaymentMerkleClaim{
		RootIndex:       cp.RootIndex,
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

func FlattenHashes(hashes [][]byte) []byte {
	result := make([]byte, 0)
	for i := 0; i < len(hashes); i++ {
		result = append(result, hashes[i]...)
	}
	return result
}
