package httpProofDataFetcher

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Layr-Labs/eigenlayer-payment-updater/pkg/proofDataFetcher"
	"go.uber.org/zap"
	"io"
	"net/http"
)

type HttpProofDataFetcher struct {
	logger      *zap.Logger
	Client      proofDataFetcher.HTTPClient
	BaseUrl     string
	Environment string
	Network     string
}

func NewS3ProofDataFetcher(
	baseUrl string,
	environment string,
	network string,
	c proofDataFetcher.HTTPClient,
	l *zap.Logger,
) *HttpProofDataFetcher {
	return &HttpProofDataFetcher{
		logger:      l,
		Client:      c,
		BaseUrl:     baseUrl,
		Environment: environment,
		Network:     network,
	}
}

func (h *HttpProofDataFetcher) FetchProofDataForDate(date string) error {
	return nil
}

func (h *HttpProofDataFetcher) FetchRecentSnapshotList() ([]proofDataFetcher.Snapshot, error) {
	fullUrl := h.buildRecentSnapshotsUrl()

	req, err := http.NewRequest(http.MethodGet, fullUrl, nil)
	if err != nil {
		h.logger.Sugar().Error("Failed to form request", zap.Error(err))
	}

	res, err := h.Client.Do(req)
	if err != nil {
		h.logger.Sugar().Error("Request failed", zap.Error(err))
		return nil, err
	}
	defer res.Body.Close()

	rawBody, err := io.ReadAll(res.Body)
	if err != nil {
		h.logger.Error("Failed to read response body", zap.Error(err))
	}

	if res.StatusCode >= 400 {
		errMsg := fmt.Sprintf("Received error code '%d'", res.StatusCode)
		h.logger.Sugar().Error(errMsg, zap.String("body", string(rawBody)))
		return nil, errors.New(errMsg)
	}

	snapshots := make([]proofDataFetcher.Snapshot, 0)
	if err := json.Unmarshal(rawBody, &snapshots); err != nil {
		h.logger.Sugar().Error("Failed to unmarshal snapshots", zap.Error(err))
		return nil, err
	}
	return snapshots, nil
}

func (h *HttpProofDataFetcher) buildRecentSnapshotsUrl() string {
	// <baseurl>/<env>/<network>/recent-snapshots.json
	return fmt.Sprintf("%s/%s/%s/recent-snapshots.json",
		h.BaseUrl,
		h.Environment,
		h.Network,
	)
}
