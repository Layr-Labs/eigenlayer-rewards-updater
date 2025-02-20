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
}

func (m *mockRewardsClient) GenerateRewards(ctx context.Context, req *v1.GenerateRewardsRequest, opts ...grpc.CallOption) (*v1.GenerateRewardsResponse, error) {
	return &v1.GenerateRewardsResponse{CutoffDate: "2024-10-31"}, nil
}

func (m *mockRewardsClient) GenerateRewardsRoot(ctx context.Context, req *v1.GenerateRewardsRootRequest, opts ...grpc.CallOption) (*v1.GenerateRewardsRootResponse, error) {
	return &v1.GenerateRewardsRootResponse{
		RewardsRoot:        "0xb4a614cc0bf38dff74822a0744aab5b8897a6868c3b612980436be219a25be21",
		RewardsCalcEndDate: "2024-10-31",
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
	expectedSnapshotDate := "2024-10-31"

	expectedSnapshotDateTime, _ := time.Parse(time.DateOnly, expectedSnapshotDate)

	expectedRootBytes, _ := utils.ConvertStringToBytes(expectedRoot)

	mockTransactor.On("SubmitRoot", mock.Anything, [32]byte(expectedRootBytes), uint32(expectedSnapshotDateTime.Unix())).Return(nil)

	updatedRoot, err := updater.Update(ctx)
	assert.Nil(t, err)
	assert.Equal(t, expectedRoot, updatedRoot.Root)
	assert.Equal(t, "2024-10-31", updatedRoot.SnapshotDate)
}
