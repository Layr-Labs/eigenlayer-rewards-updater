package httpProofDataFetcher

import (
	"context"
	"github.com/Layr-Labs/eigenlayer-rewards-updater/internal/logger"
	"github.com/Layr-Labs/eigenlayer-rewards-updater/internal/testData"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/mocktracer"
	ddTracer "gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
	"io"
	"net/http"
	"strings"
	"testing"
)

type mockHttpClient struct {
	mockDo func(r *http.Request) *http.Response
}

func (m *mockHttpClient) Do(req *http.Request) (*http.Response, error) {
	return m.mockDo(req), nil
}

func TestHttpProofDataFetcher_FetchRecentSnapshotList(t *testing.T) {
	env := "preprod"
	network := "holesky"
	baseUrl := "https://eigenpayments-dev.s3.us-east-2.amazonaws.com"

	mockClient := &mockHttpClient{
		mockDo: func(r *http.Request) *http.Response {
			return &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(strings.NewReader(testData.GetFullSnapshotDatesList())),
			}
		},
	}

	mt := mocktracer.Start()
	defer mt.Stop()

	span, ctx := ddTracer.StartSpanFromContext(context.Background(), "TestHttpProofDataFetcher_FetchRecentSnapshotList")
	defer span.Finish()

	l, _ := logger.NewLogger(&logger.LoggerConfig{Debug: true})

	fetcher := NewHttpProofDataFetcher(baseUrl, env, network, mockClient, l)

	snapshots, err := fetcher.FetchRecentSnapshotList(ctx)
	assert.Nil(t, err)
	assert.Len(t, snapshots, 10)
}

func TestHttpProofDataFetcher_FetchLatestSnapshot(t *testing.T) {
	env := "preprod"
	network := "holesky"
	baseUrl := "https://eigenpayments-dev.s3.us-east-2.amazonaws.com"

	mockClient := &mockHttpClient{
		mockDo: func(r *http.Request) *http.Response {
			return &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(strings.NewReader(testData.GetFullSnapshotDatesList())),
			}
		},
	}

	mt := mocktracer.Start()
	defer mt.Stop()

	span, ctx := ddTracer.StartSpanFromContext(context.Background(), "TestHttpProofDataFetcher_FetchLatestSnapshot")
	defer span.Finish()

	l, _ := logger.NewLogger(&logger.LoggerConfig{Debug: true})

	fetcher := NewHttpProofDataFetcher(baseUrl, env, network, mockClient, l)

	snapshot, err := fetcher.FetchLatestSnapshot(ctx)
	assert.Nil(t, err)
	assert.NotNil(t, snapshot)
}

func TestHttpProofDataFetcher_FetchClaimAmountsForDate(t *testing.T) {
	env := "preprod"
	network := "holesky"
	baseUrl := "https://eigenpayments-dev.s3.us-east-2.amazonaws.com"

	mockClient := &mockHttpClient{
		mockDo: func(r *http.Request) *http.Response {
			return &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(strings.NewReader(testData.GetFullTestEarnerLines())),
			}
		},
	}

	mt := mocktracer.Start()
	defer mt.Stop()

	span, ctx := ddTracer.StartSpanFromContext(context.Background(), "TestHttpProofDataFetcher_FetchClaimAmountsForDate")
	defer span.Finish()

	l, _ := logger.NewLogger(&logger.LoggerConfig{Debug: true})

	fetcher := NewHttpProofDataFetcher(baseUrl, env, network, mockClient, l)

	proofData, err := fetcher.FetchClaimAmountsForDate(ctx, "2024-05-07")

	assert.Nil(t, err)
	assert.NotNil(t, proofData)

	earnerAddrString := "0xd37f737629e0ddad7fc8adc7247d2e79c0296c35"
	earnerAddr := gethcommon.HexToAddress(earnerAddrString)

	tokenAddrString := "0xe1b7a1249c71b538cc183b0080ffc3efd02bffb9"
	tokenAddr := gethcommon.HexToAddress(tokenAddrString)

	amount, found := proofData.Distribution.Get(earnerAddr, tokenAddr)

	assert.True(t, found)
	assert.Equal(t, "2690822690822645700000000000", amount.String())
}
