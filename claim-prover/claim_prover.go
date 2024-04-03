package claimprover

import (
	"context"
	"fmt"
	"sync"
	"time"

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

	distributionDataService services.DistributionDataService
	distribution            *distribution.Distribution

	accountTree *merkletree.MerkleTree
	tokenTrees  map[gethcommon.Address]*merkletree.MerkleTree

	mu sync.RWMutex
}

func NewClaimProver(updateIntervalSeconds int64, distributionDataService services.DistributionDataService) *ClaimProver {
	claimProver := &ClaimProver{
		updateInterval:          time.Second * time.Duration(updateIntervalSeconds),
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

	cp.mu.Unlock()

	return nil
}

func (cp *ClaimProver) GetProof(recipient gethcommon.Address, tokens []gethcommon.Address) (*ClaimProofs, error) {
	// aquire read lock
	cp.mu.RLock()

	accountIndex, found := cp.distribution.GetAccountIndex(recipient)
	if !found {
		return nil, fmt.Errorf("account index not found for recipient %s", recipient.Hex())
	}

	// get the token proofs
	tokenProofs := make([]*merkletree.Proof, len(tokens))
	for i, token := range tokens {
		tokenIndex, found := cp.distribution.GetTokenIndex(recipient, token)
		if !found {
			return nil, fmt.Errorf("token index not found for token %s and recipient %s", token.Hex(), recipient.Hex())
		}

		tokenProof, err := cp.tokenTrees[recipient].GenerateProofWithIndex(tokenIndex, 0)
		if err != nil {
			return nil, err
		}
		tokenProofs[i] = tokenProof
	}

	// get the account proof
	accountProof, err := cp.accountTree.GenerateProofWithIndex(accountIndex, 0)
	if err != nil {
		log.Error().Msgf("failed to generate account proof: %s", err)
	}

	cp.mu.RUnlock()

	return &ClaimProofs{
		AccountProof: accountProof,
		TokenProofs:  tokenProofs,
	}, nil
}
