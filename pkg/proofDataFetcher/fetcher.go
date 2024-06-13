package proofDataFetcher

import (
	"context"
	"encoding/json"
	"github.com/Layr-Labs/eigenlayer-rewards-proofs/pkg/distribution"
	"github.com/ethereum/go-ethereum/common"
	"github.com/wealdtech/go-merkletree/v2"
	"net/http"
	"time"
)

type ProofDataFetcher interface {
	FetchClaimAmountsForDate(ctx context.Context, date string) (*RewardProofData, error)
	FetchRecentSnapshotList(ctx context.Context) ([]*Snapshot, error)
	FetchLatestSnapshot(ctx context.Context) (*Snapshot, error)
	FetchPostedRewards(ctx context.Context) ([]*SubmittedRewardRoot, error)
}

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
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

type RewardProofData struct {
	Distribution *distribution.Distribution
	AccountTree  *merkletree.MerkleTree
	TokenTree    map[common.Address]*merkletree.MerkleTree
	Hash         string
}

type SubmittedRewardRoot struct {
	RootIndex        uint32    `json:"root_index"`
	Root             string    `json:"root"`
	CalcEndTimestamp time.Time `json:"calc_end_timestamp"`
	ActivatedAt      time.Time `json:"activated_at"`
	BlockDate        time.Time `json:"block_date"`
	BlockNumber      uint64    `json:"block_number"`
}
