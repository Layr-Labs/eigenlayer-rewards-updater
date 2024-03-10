package utils

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Layr-Labs/eigenlayer-payment-updater/common"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
)

// Utils for unit and integration tests

var ()

type TestPGConnection struct {
	dbpool *pgxpool.Pool
}

func NewTestPGConnection(dbpool *pgxpool.Pool) *TestPGConnection {
	return &TestPGConnection{
		dbpool: dbpool,
	}
}

func (conn *TestPGConnection) ExecSQL(sql string, arguments ...any) {
	_, err := conn.dbpool.Exec(context.Background(), sql, arguments...)
	if err != nil {
		fmt.Printf("Exec failed: %s\n", sql)
		panic(err)
	}
}

func (conn *TestPGConnection) CleanSubgraphDeployment() {
	conn.ExecSQL(`DROP TABLE satsuma.subgraph_schema`)
	conn.ExecSQL(`DROP SCHEMA satsuma`)
}

func (conn *TestPGConnection) CreateSubgraphDeployments() {
	conn.ExecSQL(`CREATE SCHEMA IF NOT EXISTS satsuma`)
	// 	satsuma_subgraph_name | character varying |           | not null |
	//  entity_schema         | character varying |           | not null |
	conn.ExecSQL(`CREATE TABLE IF NOT EXISTS satsuma.subgraph_schema(satsuma_subgraph_name text, entity_schema text)`)
	conn.ExecSQL(`INSERT INTO satsuma.subgraph_schema VALUES ($1, $2)`, TEST_SUBGRAPH_CLAIMING_MANAGER, "sgd34")
	conn.ExecSQL(`INSERT INTO satsuma.subgraph_schema VALUES ($1, $2)`, TEST_SUBGRAPH_PAYMENT_COORDINATOR, "sgd34")
	conn.ExecSQL(`INSERT INTO satsuma.subgraph_schema VALUES ($1, $2)`, TEST_SUBGRAPH_DELEGATION_MANAGER, "sgd34")
	conn.ExecSQL(`INSERT INTO satsuma.subgraph_schema VALUES ($1, $2)`, TEST_SUBGRAPH_DELEGATION_SHARE_TRACKER, "sgd34")

	conn.ExecSQL(`CREATE SCHEMA IF NOT EXISTS sgd34`)
}

func InitializePGDocker() (*dockertest.Pool, *dockertest.Resource, *pgxpool.Pool) {
	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not construct pool: %s", err)
	}

	err = pool.Client.Ping()
	if err != nil {
		log.Fatalf("Could not connect to Docker: %s", err)
	}

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "14", // we're using Aurora 14
		Env: []string{
			"POSTGRES_PASSWORD=secret",
			"POSTGRES_USER=user_name",
			"POSTGRES_DB=dbname",
			"listen_addresses = '*'",
		},
	}, func(config *docker.HostConfig) {
		// set AutoRemove to true so that stopped container goes away by itself
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	hostAndPort := resource.GetHostPort("5432/tcp")
	databaseUrl := fmt.Sprintf("postgres://user_name:secret@%s/dbname?sslmode=disable", hostAndPort)
	log.Println("Connecting to database on url: ", databaseUrl)

	_ = resource.Expire(120) // Tell docker to hard kill the container in 120 seconds

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	pool.MaxWait = 120 * time.Second
	dbPool := common.CreateConnectionOrDie(databaseUrl)

	if err = pool.Retry(
		func() error {
			return dbPool.Ping(context.Background())
		}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}
	return pool, resource, dbPool
}
