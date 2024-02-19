package services

import (
	"log"
	"os"
	"testing"

	"github.com/Layr-Labs/eigenlayer-payment-updater/common"
	"github.com/Layr-Labs/eigenlayer-payment-updater/common/utils"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ory/dockertest/v3"
)

var (
	pool          *dockertest.Pool
	resource      *dockertest.Resource
	dbpool        *pgxpool.Pool
	schemaService *common.SubgraphSchemaService
	conn          *utils.TestPGConnection
)

func TestMain(m *testing.M) {
	pool, resource, dbpool = utils.InitializePGDocker()

	// Initialize setups
	schemaService = common.NewSubgraphSchemaService("test", dbpool)
	conn = utils.NewTestPGConnection(dbpool)
	conn.CreateSubgraphDeployments()

	//Run tests
	code := m.Run()

	// Clean up setups
	conn.CleanSubgraphDeployment()

	// You can't defer this because os.Exit doesn't care for defer
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(code)
}
