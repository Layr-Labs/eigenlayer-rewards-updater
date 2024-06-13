package httpProofDataFetcher

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Layr-Labs/eigenlayer-rewards-proofs/pkg/distribution"
	"github.com/Layr-Labs/eigenlayer-rewards-proofs/pkg/utils"
	"github.com/Layr-Labs/eigenlayer-rewards-updater/pkg/proofDataFetcher"
	"go.uber.org/zap"
	ddTracer "gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
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

func (h *HttpProofDataFetcher) FetchClaimAmountsForDate(ctx context.Context, date string) (*proofDataFetcher.RewardProofData, error) {
	span, ctx := ddTracer.StartSpanFromContext(ctx, "httpProofDataFetcher::FetchClaimAmountsForDate")
	defer span.Finish()

	h.logger.Sugar().Debug(fmt.Sprintf("Fetching claim amounts for date '%s'", date), zap.String("date", date))
	fullUrl := h.buildClaimAmountsUrl(date)

	rawBody, err := h.handleRequest(ctx, fullUrl)
	if err != nil {
		return nil, err
	}

	return h.ProcessClaimAmountsFromRawBody(ctx, rawBody)
}

func (h *HttpProofDataFetcher) ProcessClaimAmountsFromRawBody(ctx context.Context, rawBody []byte) (*proofDataFetcher.RewardProofData, error) {
	span, ctx := ddTracer.StartSpanFromContext(ctx, "httpProofDataFetcher::ProcessClaimAmountsFromRawBody")
	defer span.Finish()

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

	proof := &proofDataFetcher.RewardProofData{
		Distribution: distro,
		AccountTree:  accountTree,
		TokenTree:    tokenTree,
		Hash:         utils.ConvertBytesToString(accountTree.Root()),
	}

	return proof, nil
}

func (h *HttpProofDataFetcher) FetchRecentSnapshotList(ctx context.Context) ([]*proofDataFetcher.Snapshot, error) {
	span, ctx := ddTracer.StartSpanFromContext(ctx, "httpProofDataFetcher::FetchRecentSnapshotList")
	defer span.Finish()

	fullUrl := h.buildRecentSnapshotsUrl()

	rawBody, err := h.handleRequest(ctx, fullUrl)
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

func (h *HttpProofDataFetcher) FetchLatestSnapshot(ctx context.Context) (*proofDataFetcher.Snapshot, error) {
	span, ctx := ddTracer.StartSpanFromContext(ctx, "httpProofDataFetcher::FetchLatestSnapshot")
	defer span.Finish()

	snapshots, err := h.FetchRecentSnapshotList(ctx)

	if err != nil {
		return nil, err
	}
	if len(snapshots) == 0 {
		return nil, fmt.Errorf("no snapshots found")
	}
	return snapshots[0], nil
}

func (h *HttpProofDataFetcher) FetchPostedRewards(ctx context.Context) ([]*proofDataFetcher.SubmittedRewardRoot, error) {
	span, ctx := ddTracer.StartSpanFromContext(ctx, "httpProofDataFetcher::FetchPostedRewards")
	defer span.Finish()

	fullUrl := h.buildPostedRewardsUrl()

	rawBody, err := h.handleRequest(ctx, fullUrl)
	if err != nil {
		h.logger.Sugar().Error("Failed to fetch posted rewards", zap.Error(err))
		return nil, err
	}

	rewards := make([]*proofDataFetcher.SubmittedRewardRoot, 0)
	if err := json.Unmarshal(rawBody, &rewards); err != nil {
		h.logger.Sugar().Error("Failed to unmarshal rewards", zap.Error(err))
		return nil, err
	}
	return rewards, nil
}

func (h *HttpProofDataFetcher) handleRequest(ctx context.Context, fullUrl string) ([]byte, error) {
	span, ctx := ddTracer.StartSpanFromContext(ctx, "httpProofDataFetcher::handleRequest")
	defer span.Finish()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fullUrl, nil)
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
		h.logger.Error("Failed to read response body",
			zap.String("url", fullUrl),
			zap.Error(err),
		)
	}

	if res.StatusCode >= 400 {
		errMsg := fmt.Sprintf("Received error code '%d'", res.StatusCode)
		h.logger.Sugar().Error(errMsg,
			zap.String("url", fullUrl),
			zap.String("body", string(rawBody)),
		)
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

func (h *HttpProofDataFetcher) buildPostedRewardsUrl() string {
	// <baseurl>/<env>/<network>/submitted-payments.json
	return fmt.Sprintf("%s/%s/%s/submitted-payments.json",
		h.BaseUrl,
		h.Environment,
		h.Network,
	)
}
