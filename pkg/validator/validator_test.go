package validator

import (
	"context"
	"errors"
	"testing"

	"github.com/Layr-Labs/eigenlayer-rewards-updater/mocks"
	"go.uber.org/zap"
)

func TestValidatePostedRoot_GetCurrentRootError(t *testing.T) {
	transactor := &mocks.Transactor{}
	transactor.On("GetCurrentRoot").Return(nil, errors.New("rpc unavailable"))
	fetcher := &mocks.ProofDataFetcher{}

	v := NewValidator(transactor, fetcher, zap.NewNop())
	_, ok, err := v.ValidatePostedRoot(context.Background())
	if err == nil || ok {
		t.Fatalf("expected an error and no match, got ok=%v err=%v", ok, err)
	}
}
