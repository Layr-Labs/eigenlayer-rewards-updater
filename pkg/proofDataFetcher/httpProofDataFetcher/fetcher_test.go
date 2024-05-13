package httpProofDataFetcher

import (
	"github.com/Layr-Labs/eigenlayer-payment-updater/internal/logger"
	"github.com/Layr-Labs/eigenlayer-payment-updater/internal/testData"
	gethcommon "github.com/ethereum/go-ethereum/common"
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

	fetcher := NewHttpProofDataFetcher(baseUrl, env, network, mockClient, l)

	snapshots, err := fetcher.FetchRecentSnapshotList()
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

	l, _ := logger.NewLogger(&logger.LoggerConfig{Debug: true})

	fetcher := NewHttpProofDataFetcher(baseUrl, env, network, mockClient, l)

	snapshot, err := fetcher.FetchLatestSnapshot()
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

	l, _ := logger.NewLogger(&logger.LoggerConfig{Debug: true})

	fetcher := NewHttpProofDataFetcher(baseUrl, env, network, mockClient, l)

	proofData, err := fetcher.FetchClaimAmountsForDate("2024-05-07")

	assert.Nil(t, err)
	assert.NotNil(t, proofData)

	earnerAddrString := "0x2222aac0c980cc029624b7ff55b88bc6f63c538f"
	earnerAddr := gethcommon.HexToAddress(earnerAddrString)

	tokenAddrString := "0x3f1c547b21f65e10480de3ad8e19faac46c95034"
	tokenAddr := gethcommon.HexToAddress(tokenAddrString)

	amount, found := proofData.Distribution.Get(earnerAddr, tokenAddr)

	assert.True(t, found)
	assert.Equal(t, "1027015602000000000", amount.String())
}
