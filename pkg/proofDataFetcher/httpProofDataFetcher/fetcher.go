package httpProofDataFetcher

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Layr-Labs/eigenlayer-payment-proofs/pkg/distribution"
	"github.com/Layr-Labs/eigenlayer-payment-updater/pkg/proofDataFetcher"
	"github.com/Layr-Labs/eigenlayer-payment-updater/pkg/utils"
	"go.uber.org/zap"
	"io"
	"net/http"
	"strings"
)

type HttpProofDataFetcher struct {
	logger      *zap.Logger
	Client      proofDataFetcher.HTTPClient
	BaseUrl     string
	Environment string
	Network     string
}

func NewHttpProofDataFetcher(
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

func (h *HttpProofDataFetcher) FetchClaimAmountsForDate(date string) (*proofDataFetcher.PaymentProofData, error) {
	fullUrl := h.buildClaimAmountsUrl(date)

	rawBody, err := h.handleRequest(fullUrl)
	if err != nil {
		return nil, err
	}

	return h.ProcessClaimAmountsFromRawBody(rawBody)
}

func (h *HttpProofDataFetcher) ProcessClaimAmountsFromRawBody(rawBody []byte) (*proofDataFetcher.PaymentProofData, error) {
	strLines := strings.Split(string(rawBody), "\n")
	distro := distribution.NewDistribution()
	lines := []*distribution.EarnerLine{}
	for _, line := range strLines {
		if line == "" {
			continue
		}
		earner := &distribution.EarnerLine{}
		if err := json.Unmarshal([]byte(line), earner); err != nil {
			h.logger.Sugar().Errorf("Failed to unmarshal line: %s", line)
			return nil, err
		}
		lines = append(lines, earner)
	}
	if err := distro.LoadLines(lines); err != nil {
		h.logger.Sugar().Errorf("Failed to load lines: %s\n", err)
		return nil, err
	}

	accountTree, tokenTree, err := distro.Merklize()
	if err != nil {
		return nil, err
	}

	proof := &proofDataFetcher.PaymentProofData{
		Distribution: distro,
		AccountTree:  accountTree,
		TokenTree:    tokenTree,
		Hash:         utils.ConvertBytesToString(accountTree.Root()),
	}

	return proof, nil
}

func (h *HttpProofDataFetcher) FetchRecentSnapshotList() ([]*proofDataFetcher.Snapshot, error) {
	fullUrl := h.buildRecentSnapshotsUrl()

	rawBody, err := h.handleRequest(fullUrl)
	if err != nil {
		return nil, err
	}

	snapshots := make([]*proofDataFetcher.Snapshot, 0)
	if err := json.Unmarshal(rawBody, &snapshots); err != nil {
		h.logger.Sugar().Error("Failed to unmarshal snapshots", zap.Error(err))
		return nil, err
	}
	return snapshots, nil
}

func (h *HttpProofDataFetcher) FetchLatestSnapshot() (*proofDataFetcher.Snapshot, error) {
	snapshots, err := h.FetchRecentSnapshotList()
	if err != nil {
		return nil, err
	}
	if len(snapshots) == 0 {
		return nil, fmt.Errorf("no snapshots found")
	}
	return snapshots[0], nil
}

func (h *HttpProofDataFetcher) handleRequest(fullUrl string) ([]byte, error) {
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

	return rawBody, nil
}

func (h *HttpProofDataFetcher) buildRecentSnapshotsUrl() string {
	// <baseurl>/<env>/<network>/recent-snapshots.json
	return fmt.Sprintf("%s/%s/%s/recent-snapshots.json",
		h.BaseUrl,
		h.Environment,
		h.Network,
	)
}

func (h *HttpProofDataFetcher) buildClaimAmountsUrl(snapshotDate string) string {
	// <baseurl>/<env>/<network>/<snapshot_date>/claim-amounts.json
	return fmt.Sprintf("%s/%s/%s/%s/claim-amounts.json",
		h.BaseUrl,
		h.Environment,
		h.Network,
		snapshotDate,
	)
}
