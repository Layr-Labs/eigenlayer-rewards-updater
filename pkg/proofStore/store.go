package proofStore

import (
	"encoding/json"
	"github.com/Layr-Labs/eigenlayer-payment-proofs/pkg/distribution"
	"github.com/ethereum/go-ethereum/common"
	"github.com/wealdtech/go-merkletree/v2"
	"time"
)

type ProofStore interface {
	// Get the current proof
	GetProofForActivePayment() (*PaymentProofData, error)
	GetRecentSnapshots() ([]Snapshot, error)
}

type Snapshot struct {
	SnapshotDate time.Time `json:"snapshot_date"`
}

func (s *Snapshot) GetDateString() string {
	return s.SnapshotDate.UTC().Format("2006-01-02")
}

func (s *Snapshot) UnmarshalJSON(data []byte) error {
	var index interface{}
	err := json.Unmarshal(data, &index)
	if err != nil {
		return err
	}

	m := index.(map[string]interface{})
	for k, v := range m {
		switch k {
		case "snapshot_date":
			s.SnapshotDate = time.Unix(int64(v.(float64))/1000, 0).UTC()
		}
	}
	return nil
}

type PaymentProofData struct {
	Distribution *distribution.Distribution
	AccountTree  *merkletree.MerkleTree
	TokenTree    map[common.Address]*merkletree.MerkleTree
	Hash         string
}

type SubmittedPayment struct {
	RootIndex        uint64    `json:"root_index"`
	Root             string    `json:"root"`
	CalcEndTimestamp time.Time `json:"calc_end_timestamp"`
	ActivatedAt      time.Time `json:"activated_at"`
	BlockDate        time.Time `json:"block_date"`
	BlockNumber      uint64    `json:"block_number"`
}

func (s *SubmittedPayment) UnmarshalJSON(data []byte) error {
	var index interface{}
	err := json.Unmarshal(data, &index)
	if err != nil {
		return err
	}

	m := index.(map[string]interface{})
	for k, v := range m {
		switch k {
		case "calc_end_timestamp":
			s.CalcEndTimestamp = time.Unix(int64(v.(float64))/1000, 0).UTC()
		case "activated_at":
			s.ActivatedAt = time.Unix(int64(v.(float64))/1000, 0).UTC()
		case "block_date":
			s.BlockDate = time.Unix(int64(v.(float64))/1000, 0).UTC()
		}
	}
	return nil
}

func (sp *SubmittedPayment) GetPaymentDate() string {
	return formatPaymentTimeAsDateString(sp.CalcEndTimestamp)
}

func (sp *SubmittedPayment) GetActivatedAtDate() string {
	return formatPaymentTimeAsDateString(sp.ActivatedAt)
}

func formatPaymentTimeAsDateString(t time.Time) string {
	return t.UTC().Format("2006-01-02")
}
