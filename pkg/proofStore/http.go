package proofStore

import (
	"encoding/json"
	"fmt"
	"github.com/Layr-Labs/eigenlayer-payment-proofs/pkg/distribution"
	"github.com/Layr-Labs/eigenlayer-payment-updater/pkg/utils"
	"go.uber.org/zap"
	"io"
	"net/http"
	"strings"
	"time"
)

type HTTPProofStore struct {
	BaseUrl     string
	Environment string
	Network     string

	RecentPayments []SubmittedPayment
	ActivePayment  *SubmittedPayment
	Snapshots      []Snapshot
	Proofs         map[string]*PaymentProofData
	logger         *zap.Logger
}

func NewHTTPProofStore(baseUrl string, env string, network string, l *zap.Logger) *HTTPProofStore {
	return &HTTPProofStore{
		BaseUrl:     baseUrl,
		Environment: env,
		Network:     network,
		Proofs:      make(map[string]*PaymentProofData),
		logger:      l,
	}
}

func (h *HTTPProofStore) GetProofForActivePayment() (*PaymentProofData, error) {
	var proof *PaymentProofData
	var err error

	if h.ActivePayment == nil {
		err := h.refreshSnapshotData()
		if err != nil {
			return nil, err
		}

		// If activePayment is still nil, then we have no active payment
		if h.ActivePayment == nil {
			return nil, fmt.Errorf("no active payment found")
		}
	}
	proof, ok := h.Proofs[h.ActivePayment.GetPaymentDate()]
	if !ok {
		h.logger.Sugar().Infof("No proof found for date %s. Fetching proof", h.ActivePayment.GetPaymentDate())
		proof, err = h.fetchClaimAmountsForDate(h.ActivePayment)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch claim amounts for date %s: %s", h.ActivePayment.GetPaymentDate(), err)
		}
	}

	return proof, nil
}

// refreshSnapshotData:
//   - fetches the list of submitted payments from the S3 bucket
//   - For the active payment, fetches the claim amounts
func (h *HTTPProofStore) refreshSnapshotData() error {
	err := h.fetchSubmittedPayments()
	if err != nil {
		return err
	}

	if h.ActivePayment == nil {
		return fmt.Errorf("no active payment found")
	}

	h.logger.Sugar().Debug("Fetching active payment")
	_, err = h.fetchClaimAmountsForDate(h.ActivePayment)
	if err != nil {
		return err
	}

	return nil
}

// fetchSubmittedPayments: fetches the list of on-chain submitted payments from the S3 bucket
func (h *HTTPProofStore) fetchSubmittedPayments() error {
	// <baseurl>/<env>/<network>/submitted-payments.json
	u := fmt.Sprintf("%s/%s/%s/submitted-payments.json",
		h.BaseUrl,
		h.Environment,
		h.Network,
	)
	h.logger.Sugar().Debugf("Getting submitted payments from %s", u)
	resp, err := http.Get(u)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	rawBody, err := io.ReadAll(resp.Body)

	if err != nil {
		return fmt.Errorf("failed to read response body: %s", err)
	}

	if resp.StatusCode >= 400 {
		return fmt.Errorf("failed to fetch submitted-payments. got status code %d", resp.StatusCode)
	}

	submittedPayments := make([]SubmittedPayment, 0)
	err = json.Unmarshal(rawBody, &submittedPayments)

	h.RecentPayments = submittedPayments

	if len(submittedPayments) == 0 {
		return nil
	}

	// If the latest payment has reached its activation date, it is the active payment and can be claimed against
	if submittedPayments[0].ActivatedAt.Before(time.Now()) {
		h.ActivePayment = &submittedPayments[0]
		return nil
	} else if len(submittedPayments) > 1 {
		// Otherwise, the active payment is the one before the latest payment
		h.ActivePayment = &submittedPayments[1]
	}
	// Else, we'll leave the active payment as nil
	return err
}

// fetchClaimAmountsForDate: fetches the claim amounts for a given payment date
func (h *HTTPProofStore) fetchClaimAmountsForDate(submittedPayment *SubmittedPayment) (*PaymentProofData, error) {
	paymentDate := submittedPayment.GetPaymentDate()

	// <baseurl>/<env>/<network>/<snapshot_date>/claim_amounts.json
	u := fmt.Sprintf("%s/%s/%s/%s/claim-amounts.json",
		h.BaseUrl,
		h.Environment,
		h.Network,
		paymentDate,
	)
	h.logger.Sugar().Debugf("Getting claim amounts for %s from %s", paymentDate, u)
	resp, err := http.Get(u)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	rawBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("failed to fetch snapshot for %s: %s. Got status code %d", paymentDate, string(rawBody), resp.StatusCode)
	}

	proof, err := h.processSnapshotFromRawBody(rawBody)
	if err != nil {
		h.logger.Sugar().Errorf("Failed to process snapshot from raw body: %s", err)
		return nil, err
	}
	h.Proofs[paymentDate] = proof
	h.logger.Sugar().Infof("Added proof %s for date %s to cache", proof.Hash, paymentDate)
	return proof, nil
}

func (h *HTTPProofStore) processSnapshotFromRawBody(rawBody []byte) (*PaymentProofData, error) {
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

	proof := &PaymentProofData{
		Distribution: distro,
		AccountTree:  accountTree,
		TokenTree:    tokenTree,
		Hash:         utils.ConvertBytesToString(accountTree.Root()),
	}

	return proof, nil
}
