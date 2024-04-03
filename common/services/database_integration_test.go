package services_test

import (
	"log"
	"os"
	"testing"

	"github.com/Layr-Labs/eigenlayer-payment-updater/common/utils"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ory/dockertest/v3"
)

var (
	pool     *dockertest.Pool
	resource *dockertest.Resource
	dbpool   *pgxpool.Pool
	conn     *utils.TestPGConnection
)

func TestMain(m *testing.M) {
	utils.SetTestEnv()

	pool, resource, dbpool = utils.InitializePGDocker()

	// Initialize setups
	conn = utils.NewTestPGConnection(dbpool)
	conn.CreateDB()

	//Run tests
	code := m.Run()

	conn.CleanDB()

	// You can't defer this because os.Exit doesn't care for defer
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(code)
}
