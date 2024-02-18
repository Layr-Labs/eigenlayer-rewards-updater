package calculator

import (
	"context"
	"os"
	"testing"

	"github.com/Layr-Labs/eigenlayer-payment-updater/common"
	"github.com/stretchr/testify/assert"

	"github.com/joho/godotenv"
)

func TestPaymentCalculatorDataService(t *testing.T) {
	err := godotenv.Load("../.env") // Replace with your file path
	if err != nil {
		t.Fatal("Error loading .env file", err)
	}

	connString := common.CreateConnectionString(
		"eigenlabs_team",
		os.Getenv("DB_PASSWORD"),
		"eigenlabs-graph-node-production-3.cg7azkhq5rv5.us-east-1.rds.amazonaws.com",
		"5432",
		"graph_node_eigenlabs_3",
	)
	dbpool := common.CreateConnectionOrDie(connString)
	defer dbpool.Close()
	schemaService := common.NewSubgraphSchemaService("test", dbpool)

	elpds := NewPaymentsDataServiceImpl(
		dbpool,
		schemaService,
	)

	t.Run("test GetPaymentsCalculatedUntilTimestamp", func(t *testing.T) {
		paymentsCalculatedUntilTimestamp, err := elpds.GetPaymentsCalculatedUntilTimestamp(context.Background())
		if err != nil {
			t.Fatal(err)
		}
		assert.GreaterOrEqual(t, paymentsCalculatedUntilTimestamp.Int64(), int64(1708285990))
	})

	// TODO: overlapping range payments test

}
