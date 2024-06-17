package validator

import (
	"context"
	"fmt"
	"github.com/Layr-Labs/eigenlayer-rewards-updater/pkg/proofDataFetcher"
	"github.com/Layr-Labs/eigenlayer-rewards-updater/pkg/services"
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

func (v *Validator) ValidatePostedRoot(ctx context.Context) error {
	span, ctx := ddTracer.StartSpanFromContext(ctx, "validator::ValidatePostedRoot")
	defer span.Finish()

	retrievedRoot, err := v.transactor.GetCurrentRoot()
	fmt.Printf("Retrieved root: %+v\n", retrievedRoot)

	// Grab latest submitted timestamp directly from chain
	latestSubmittedTimestamp, err := v.transactor.CurrRewardsCalculationEndTimestamp()
	if err != nil {
		v.logger.Sugar().Error("Failed to get latest submitted timestamp", zap.Error(err))
		return err
	}
	lst := time.Unix(int64(latestSubmittedTimestamp), 0).UTC().Format(time.DateOnly)

	// Get the data for the latest snapshot and load it into a distribution instance
	rewardsProofData, err := v.proofDataFetcher.FetchClaimAmountsForDate(ctx, lst)
	if err != nil {
		v.logger.Sugar().Error(fmt.Sprintf("Failed to fetch claim amounts for date '%s'", lst), zap.Error(err))
		return err
	}

	root := rewardsProofData.AccountTree.Root()

	if err != nil {
		v.logger.Sugar().Error("Failed to get root by index", zap.Error(err))
		return err
	}

	fmt.Printf("Retrieved root: %+v\n", retrievedRoot)
	fmt.Printf("calculated root: %+v\n", root)
	return nil
}
