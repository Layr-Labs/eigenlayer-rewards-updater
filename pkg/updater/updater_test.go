package updater_test

import (
	"context"
	"fmt"
	"github.com/Layr-Labs/eigenlayer-rewards-proofs/pkg/utils"
	"github.com/Layr-Labs/eigenlayer-rewards-updater/internal/logger"
	"github.com/Layr-Labs/eigenlayer-rewards-updater/internal/metrics"
	"github.com/Layr-Labs/eigenlayer-rewards-updater/mocks"
	"github.com/Layr-Labs/eigenlayer-rewards-updater/pkg/sidecar"
	"github.com/Layr-Labs/eigenlayer-rewards-updater/pkg/updater"
	v1 "github.com/Layr-Labs/protocol-apis/gen/protos/eigenlayer/sidecar/v1/rewards"
	"google.golang.org/grpc"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/mocktracer"
	ddTracer "gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockHttpClient struct {
	mockDo func(r *http.Request) *http.Response
}

func (m *mockHttpClient) Do(req *http.Request) (*http.Response, error) {
	return m.mockDo(req), nil
}

type mockRewardsClient struct {
	mock.Mock
	calcEndDate string // controls the computed rewardsCalcEndDate; defaults to a Sunday
}

func (m *mockRewardsClient) endDate() string {
	if m.calcEndDate != "" {
		return m.calcEndDate
	}
	return "2026-06-28" // a Sunday (weekly boundary)
}

func (m *mockRewardsClient) GenerateRewards(ctx context.Context, req *v1.GenerateRewardsRequest, opts ...grpc.CallOption) (*v1.GenerateRewardsResponse, error) {
	return &v1.GenerateRewardsResponse{CutoffDate: m.endDate()}, nil
}

func (m *mockRewardsClient) GenerateRewardsRoot(ctx context.Context, req *v1.GenerateRewardsRootRequest, opts ...grpc.CallOption) (*v1.GenerateRewardsRootResponse, error) {
	return &v1.GenerateRewardsRootResponse{
		RewardsRoot:        "0xb4a614cc0bf38dff74822a0744aab5b8897a6868c3b612980436be219a25be21",
		RewardsCalcEndDate: m.endDate(),
	}, nil
}

func TestUpdaterUpdate(t *testing.T) {
	_, err := metrics.InitStatsdClient("", false)
	fmt.Printf("err: %v\n", err)

	mt := mocktracer.Start()
	defer mt.Stop()

	span, ctx := ddTracer.StartSpanFromContext(context.Background(), "TestUpdaterUpdate")
	defer span.Finish()

	l, _ := logger.NewLogger(&logger.LoggerConfig{Debug: true})

	mockTransactor := &mocks.Transactor{}

	mockSidecarClient := &sidecar.SidecarClient{
		Rewards: &mockRewardsClient{},
	}

	updater, err := updater.NewUpdater(mockTransactor, mockSidecarClient, l)
	assert.Nil(t, err)

	expectedRoot := "0xb4a614cc0bf38dff74822a0744aab5b8897a6868c3b612980436be219a25be21"
	expectedSnapshotDate := "2026-06-28" // Sunday

	expectedSnapshotDateTime, _ := time.Parse(time.DateOnly, expectedSnapshotDate)

	expectedRootBytes, _ := utils.ConvertStringToBytes(expectedRoot)

	// on-chain root is older than the computed one -> submit proceeds
	mockTransactor.On("CurrRewardsCalculationEndTimestamp").Return(uint32(expectedSnapshotDateTime.Unix())-uint32(7*24*3600), nil)
	mockTransactor.On("SubmitRoot", mock.Anything, [32]byte(expectedRootBytes), uint32(expectedSnapshotDateTime.Unix())).Return(nil)

	updatedRoot, err := updater.Update(ctx)
	assert.Nil(t, err)
	assert.Equal(t, expectedRoot, updatedRoot.Root)
	assert.Equal(t, expectedSnapshotDate, updatedRoot.SnapshotDate)
	mockTransactor.AssertExpectations(t)
}

// When the on-chain root is already at (or beyond) the computed cutoff, the
// updater must NOT submit, and must return success (a healthy no-op).
func TestUpdaterUpdate_NoNewRoot(t *testing.T) {
	_, _ = metrics.InitStatsdClient("", false)

	mt := mocktracer.Start()
	defer mt.Stop()
	span, ctx := ddTracer.StartSpanFromContext(context.Background(), "TestUpdaterUpdate_NoNewRoot")
	defer span.Finish()

	l, _ := logger.NewLogger(&logger.LoggerConfig{Debug: true})
	mockTransactor := &mocks.Transactor{}
	mockSidecarClient := &sidecar.SidecarClient{Rewards: &mockRewardsClient{}}

	u, err := updater.NewUpdater(mockTransactor, mockSidecarClient, l)
	assert.Nil(t, err)

	snapshotDateTime, _ := time.Parse(time.DateOnly, "2026-06-28") // Sunday
	// on-chain root already at the same cutoff -> nothing to submit
	mockTransactor.On("CurrRewardsCalculationEndTimestamp").Return(uint32(snapshotDateTime.Unix()), nil)

	updatedRoot, err := u.Update(ctx)
	assert.Nil(t, err)
	assert.Equal(t, "0xb4a614cc0bf38dff74822a0744aab5b8897a6868c3b612980436be219a25be21", updatedRoot.Root)
	assert.Equal(t, "2026-06-28", updatedRoot.SnapshotDate)
	mockTransactor.AssertNotCalled(t, "SubmitRoot", mock.Anything, mock.Anything, mock.Anything)
}

// A non-Sunday (intermediate/daily) calc-end must NEVER be submitted on-chain,
// regardless of when the updater runs.
func TestUpdaterUpdate_NonSundaySkipped(t *testing.T) {
	_, _ = metrics.InitStatsdClient("", false)

	mt := mocktracer.Start()
	defer mt.Stop()
	span, ctx := ddTracer.StartSpanFromContext(context.Background(), "TestUpdaterUpdate_NonSundaySkipped")
	defer span.Finish()

	l, _ := logger.NewLogger(&logger.LoggerConfig{Debug: true})
	mockTransactor := &mocks.Transactor{}
	// 2026-06-25 is a Thursday (not a weekly boundary)
	mockSidecarClient := &sidecar.SidecarClient{Rewards: &mockRewardsClient{calcEndDate: "2026-06-25"}}

	u, err := updater.NewUpdater(mockTransactor, mockSidecarClient, l)
	assert.Nil(t, err)

	updatedRoot, err := u.Update(ctx)
	assert.Nil(t, err) // healthy no-op, not an error
	assert.Equal(t, "2026-06-25", updatedRoot.SnapshotDate)
	// guard fires before any on-chain interaction: neither read nor submit
	mockTransactor.AssertNotCalled(t, "SubmitRoot", mock.Anything, mock.Anything, mock.Anything)
	mockTransactor.AssertNotCalled(t, "CurrRewardsCalculationEndTimestamp")
}
