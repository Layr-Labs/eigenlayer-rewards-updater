package calculator

import (
	"context"
	"os"
	"testing"

	"github.com/Layr-Labs/eigenlayer-payment-updater/common"

	"github.com/joho/godotenv"
)

func TestPaymentCalculatorDataService(t *testing.T) {
	const (
		claimingManagerSubgraph    = "claiming-manager-raw-events"
		paymentCoordinatorSubgraph = "payment-coordinator-raw-events"
	)

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
	dbpool := common.MustCreateConnection(connString)
	defer dbpool.Close()
	schemaService := common.NewSubgraphSchemaService(dbpool)

	subgraphProvider, err := common.ToSubgraphProvider("satsuma")
	if err != nil {
		panic(err)
	}

	elpds := NewPaymentsDataServiceImpl(
		dbpool,
		schemaService,
		subgraphProvider,
		claimingManagerSubgraph,
		paymentCoordinatorSubgraph,
	)

	t.Run("test GetPaymentsCalculatedUntilTimestamp", func(t *testing.T) {
		paymentsCalculatedUntilTimestamp, err := elpds.GetPaymentsCalculatedUntilTimestamp(context.Background())
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("payments calculated until timestamp: %v", paymentsCalculatedUntilTimestamp)
		t.Fail()
	})

	// TODO: overlapping range payments test

}
