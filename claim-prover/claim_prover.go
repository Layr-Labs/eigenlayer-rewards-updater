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

	paymentsDataService services.PaymentsDataService

	distributionDataService services.DistributionDataService
	root                    [32]byte
	distribution            *distribution.Distribution

	accountTree *merkletree.MerkleTree
	tokenTrees  map[gethcommon.Address]*merkletree.MerkleTree

	mu sync.RWMutex
}

func NewClaimProver(updateIntervalSeconds int64, paymentsDataService services.PaymentsDataService, distributionDataService services.DistributionDataService) *ClaimProver {
	claimProver := &ClaimProver{
		updateInterval:          time.Second * time.Duration(updateIntervalSeconds),
		paymentsDataService:     paymentsDataService,
		distributionDataService: distributionDataService,
		mu:                      sync.RWMutex{},
	}

	// start the update
	claimProver.start()

	return claimProver
}

func (cp *ClaimProver) start() {
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
	// get the latest root submission
	root, _, err := cp.paymentsDataService.GetLatestRootSubmission(ctx)
	if err != nil {
		return err
	}

	// if the root is the same as the current root, do nothing
	if root == cp.root {
		return nil
	}

	// aquire write lock
	cp.mu.Lock()

	// get the distribution for the root
	cp.distribution, err = cp.distributionDataService.GetDistribution(root)
	if err != nil {
		return err
	}

	// generate the trees
	cp.accountTree, cp.tokenTrees, err = cp.distribution.Merklize()
	if err != nil {
		return err
	}
	// set the new root
	cp.root = root

	cp.mu.Unlock()

	return nil
}

func (cp *ClaimProver) GetProof(recipient gethcommon.Address, tokens []gethcommon.Address) (*ClaimProofs, error) {
	// aquire read lock
	cp.mu.RLock()

	// get the account root
	tokenTree, found := cp.tokenTrees[recipient]
	if !found {
		return nil, fmt.Errorf("recipient not found %s", recipient.Hex())
	}
	accountRoot := tokenTree.Root()

	// get the token proofs
	tokenProofs := make([]*merkletree.Proof, len(tokens))
	for i, token := range tokens {
		amount, found := cp.distribution.Get(recipient, token)
		if !found {
			return nil, fmt.Errorf("token %s for recipient %s", token.Hex(), recipient.Hex())
		}

		tokenProof, err := cp.tokenTrees[token].GenerateProof(distribution.EncodeTokenLeaf(token, amount), 0)
		if err != nil {
			return nil, err
		}
		tokenProofs[i] = tokenProof
	}

	// get the account proof
	accountProof, err := cp.accountTree.GenerateProof(distribution.EncodeAccountLeaf(recipient, accountRoot), 0)
	if err != nil {
		log.Error().Msgf("failed to generate account proof: %s", err)
	}

	cp.mu.RUnlock()

	return &ClaimProofs{
		AccountProof: accountProof,
		TokenProofs:  tokenProofs,
	}, nil
}
