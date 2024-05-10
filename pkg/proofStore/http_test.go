package proofStore

import (
	"github.com/Layr-Labs/eigenlayer-payment-updater/internal/logger"
	"github.com/Layr-Labs/eigenlayer-payment-updater/internal/testData"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHTTPProofStore_processSnapshotFromRawBody(t *testing.T) {
	const bucketUrl = "https://eigenpayments-dev.s3.us-east-2.amazonaws.com"
	l, _ := logger.NewLogger(&logger.LoggerConfig{Debug: true})

	testPayload := testData.GetFullTestEarnerLines()

	t.Run("takes a raw json lines payload and parses it to a distribution", func(t *testing.T) {
		store := NewHTTPProofStore(bucketUrl, "preprod", "holesky", l)

		proof, err := store.processSnapshotFromRawBody([]byte(testPayload))
		assert.Nil(t, err)
		assert.NotNil(t, proof)
	})
}
