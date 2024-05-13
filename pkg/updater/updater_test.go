package updater_test

import (
	"context"
	"fmt"
	"github.com/Layr-Labs/eigenlayer-payment-updater/internal/logger"
	"github.com/Layr-Labs/eigenlayer-payment-updater/internal/testData"
	"github.com/Layr-Labs/eigenlayer-payment-updater/mocks"
	"github.com/Layr-Labs/eigenlayer-payment-updater/pkg/proofDataFetcher/httpProofDataFetcher"
	"github.com/Layr-Labs/eigenlayer-payment-updater/pkg/updater"
	"io"
	"net/http"
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var testTimestamp int32 = 1712127631

type mockHttpClient struct {
	mockDo func(r *http.Request) *http.Response
}

func (m *mockHttpClient) Do(req *http.Request) (*http.Response, error) {
	fmt.Printf("DO request url: %s\n", req.URL.String())
	return m.mockDo(req), nil
}

func TestUpdaterUpdate(t *testing.T) {
	env := "preprod"
	network := "holesky"
	baseUrl := "https://eigenpayments-dev.s3.us-east-2.amazonaws.com"

	mockClient := &mockHttpClient{
		mockDo: func(r *http.Request) *http.Response {
			recentSnapshotsRegex := regexp.MustCompile(`\/recent-snapshots\.json$`)
			claimAmountsRegex := regexp.MustCompile(`\/(\d{4}-\d{2}-\d{2})\/claim-amounts\.json$`)

			fmt.Printf("request url: %s\n", r.URL.String())
			if recentSnapshotsRegex.Match([]byte(r.URL.String())) {
				fmt.Printf("Matched recent snapshots: %s\n", r.URL.String())
				return &http.Response{
					StatusCode: 200,
					Body:       io.NopCloser(strings.NewReader(testData.GetFullSnapshotDatesList())),
				}
			} else if claimAmountsRegex.Match([]byte(r.URL.String())) {
				fmt.Printf("Matched claim amounts: %s\n", r.URL.String())
				return &http.Response{
					StatusCode: 200,
					Body:       io.NopCloser(strings.NewReader(testData.GetFullTestEarnerLines())),
				}
			}

			return &http.Response{StatusCode: 400}
		},
	}
	// 2024-05-06
	currentPaymentCalcEndTimestamp := uint32(1714953600)
	expectedPaymentTimestamp := time.Unix(int64(1715040000), 0).UTC()

	logger, _ := logger.NewLogger(&logger.LoggerConfig{Debug: true})

	mockTransactor := &mocks.Transactor{}
	mockDistributionDataService := &mocks.DistributionDataService{}

	fetcher := httpProofDataFetcher.NewHttpProofDataFetcher(baseUrl, env, network, mockClient, logger)

	updater, err := updater.NewUpdater(mockTransactor, mockDistributionDataService, fetcher, logger)
	assert.Nil(t, err)

	// setup data
	processedData, _ := fetcher.ProcessClaimAmountsFromRawBody([]byte(testData.GetFullTestEarnerLines()))

	rootBytes := processedData.AccountTree.Root()

	var root [32]byte
	copy(root[:], rootBytes)

	mockTransactor.On("CurrPaymentCalculationEndTimestamp").Return(currentPaymentCalcEndTimestamp, nil)
	mockTransactor.On("SubmitRoot", mock.Anything, root, uint32(expectedPaymentTimestamp.Unix())).Return(nil)

	accountTree, err := updater.Update(context.Background())
	assert.Nil(t, err)
	assert.Equal(t, rootBytes, accountTree.Root())
}
