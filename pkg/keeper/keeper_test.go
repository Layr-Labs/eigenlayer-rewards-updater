package keeper

import (
	"context"
	"errors"
	"math/big"
	"testing"

	"github.com/Layr-Labs/eigenlayer-rewards-updater/internal/logger"
	"github.com/Layr-Labs/eigenlayer-rewards-updater/internal/metrics"
	"github.com/stretchr/testify/assert"
)

type fakeEmissions struct {
	pressable    bool
	pressableErr error
	pressErr     error
	pressCalls   int
	lastLength   *big.Int
}

func (f *fakeEmissions) IsButtonPressable(ctx context.Context) (bool, error) {
	return f.pressable, f.pressableErr
}
func (f *fakeEmissions) TotalProcessableDistributions(ctx context.Context) (*big.Int, error) {
	return big.NewInt(2), nil
}
func (f *fakeEmissions) CurrentEpoch(ctx context.Context) (*big.Int, error) {
	return big.NewInt(16), nil
}
func (f *fakeEmissions) PressButton(ctx context.Context, length *big.Int) (string, error) {
	f.pressCalls++
	f.lastLength = length
	if f.pressErr != nil {
		return "", f.pressErr
	}
	return "0xdeadbeef", nil
}

func newTestKeeper(e EmissionsController) *Keeper {
	_, _ = metrics.InitStatsdClient("", false)
	l, _ := logger.NewLogger(&logger.LoggerConfig{Debug: true})
	return NewKeeper(e, l)
}

// When the button is pressable, Press submits pressButton(MaxUint256) once.
func TestPress_Pressable(t *testing.T) {
	e := &fakeEmissions{pressable: true}
	res, err := newTestKeeper(e).Press(context.Background())
	assert.Nil(t, err)
	assert.True(t, res.Pressed)
	assert.Equal(t, "0xdeadbeef", res.TxHash)
	assert.Equal(t, 1, e.pressCalls)
	assert.Equal(t, 0, MaxUint256.Cmp(e.lastLength)) // pressed with type(uint256).max
}

// When the button is not pressable, Press is a healthy no-op: no submit, no error.
func TestPress_NotPressable(t *testing.T) {
	e := &fakeEmissions{pressable: false}
	res, err := newTestKeeper(e).Press(context.Background())
	assert.Nil(t, err)
	assert.False(t, res.Pressed)
	assert.Equal(t, 0, e.pressCalls)
}

// A failed press surfaces the error (so the process exits non-zero).
func TestPress_PressError(t *testing.T) {
	e := &fakeEmissions{pressable: true, pressErr: errors.New("boom")}
	_, err := newTestKeeper(e).Press(context.Background())
	assert.NotNil(t, err)
	assert.Equal(t, 1, e.pressCalls)
}

// A failed pressability check is an error, not a silent no-op.
func TestPress_PressableCheckError(t *testing.T) {
	e := &fakeEmissions{pressableErr: errors.New("rpc down")}
	_, err := newTestKeeper(e).Press(context.Background())
	assert.NotNil(t, err)
	assert.Equal(t, 0, e.pressCalls)
}
