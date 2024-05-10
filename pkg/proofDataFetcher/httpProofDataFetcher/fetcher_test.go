package httpProofDataFetcher

import (
	"github.com/Layr-Labs/eigenlayer-payment-updater/internal/logger"
	"github.com/Layr-Labs/eigenlayer-payment-updater/internal/testData"
	"github.com/stretchr/testify/assert"
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

	l, _ := logger.NewLogger(&logger.LoggerConfig{Debug: true})

	fetcher := NewS3ProofDataFetcher(baseUrl, env, network, mockClient, l)

	snapshots, err := fetcher.FetchRecentSnapshotList()
	assert.Nil(t, err)
	assert.Len(t, snapshots, 10)
}
