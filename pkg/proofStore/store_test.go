package proofStore

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSnapshot_UnmarshalJSON(t *testing.T) {
	snapshotJson := `{"snapshot_date":1714953600000}`

	snapshot := &Snapshot{}
	err := snapshot.UnmarshalJSON([]byte(snapshotJson))
	assert.Nil(t, err)
	assert.Equal(t, "2024-05-06", snapshot.GetDateString())
}
