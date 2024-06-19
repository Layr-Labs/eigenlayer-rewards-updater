package validator

import (
	"context"
	"fmt"
	"github.com/Layr-Labs/eigenlayer-rewards-updater/pkg/proofDataFetcher"
	"github.com/Layr-Labs/eigenlayer-rewards-updater/pkg/services"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/google/go-cmp/cmp"
	"go.uber.org/zap"
	ddTracer "gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
	"time"
)

type Validator struct {
	transactor       services.Transactor
	logger           *zap.Logger
	proofDataFetcher proofDataFetcher.ProofDataFetcher
}

func NewValidator(
	transactor services.Transactor,
	fetcher proofDataFetcher.ProofDataFetcher,
	logger *zap.Logger,
) *Validator {
	return &Validator{
		transactor:       transactor,
		logger:           logger,
		proofDataFetcher: fetcher,
	}
}

func (v *Validator) ValidatePostedRoot(ctx context.Context) (string, bool, error) {
	span, ctx := ddTracer.StartSpanFromContext(ctx, "validator::ValidatePostedRoot")
	defer span.Finish()

	retrievedRoot, err := v.transactor.GetCurrentRoot()

	lst := time.Unix(int64(retrievedRoot.RewardsCalculationEndTimestamp), 0).UTC().Format(time.DateOnly)

	v.logger.Sugar().Info(fmt.Sprintf("Retrieved root has timestamp of '%v'", lst))

	// Get the data for the latest snapshot and load it into a distribution instance
	rewardsProofData, err := v.proofDataFetcher.FetchClaimAmountsForDate(ctx, lst)
	if err != nil {
		v.logger.Sugar().Error(fmt.Sprintf("Failed to fetch claim amounts for date '%s'", lst), zap.Error(err))
		return lst, false, err
	}

	root := rewardsProofData.AccountTree.Root()

	if err != nil {
		v.logger.Sugar().Error("Failed to get root by index", zap.Error(err))
		return lst, false, err
	}

	postedRoot := hexutil.Encode(retrievedRoot.Root[:])
	computedRoot := hexutil.Encode(root[:])

	if !cmp.Equal(postedRoot, computedRoot) {
		v.logger.Sugar().Error("Roots do not match",
			zap.String("postedRoot", postedRoot),
			zap.String("computedRoot", computedRoot),
		)
		return lst, false, nil
	}
	v.logger.Sugar().Info("Roots match")
	return lst, true, nil
}
