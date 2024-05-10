package proofDataFetcher

import (
	"encoding/json"
	"net/http"
	"time"
)

type ProofDataFetcher interface {
	FetchProofDataForDate(date string) error
	FetchRecentSnapshotList() ([]Snapshot, error)
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
