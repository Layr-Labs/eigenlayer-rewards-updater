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
	conn.ExecSQL(`DROP TABLE subgraphs.subgraph_version`)
	conn.ExecSQL(`DROP TABLE subgraphs.subgraph`)
	conn.ExecSQL(`DROP SCHEMA subgraphs`)
}

func (conn *TestPGConnection) CreateSubgraphDeployments() {
	conn.ExecSQL(`CREATE SCHEMA IF NOT EXISTS subgraphs`)
	conn.ExecSQL(`CREATE TABLE IF NOT EXISTS subgraphs.subgraph_version(id text, subgraph text, vid bigint)`)
	conn.ExecSQL(`INSERT INTO subgraphs.subgraph_version VALUES ($1, $2, $3)`, "anytext", TEST_SUBGRAPH_CLAIMING_MANAGER, 34)
	conn.ExecSQL(`INSERT INTO subgraphs.subgraph_version VALUES ($1, $2, $3)`, "anytext", TEST_SUBGRAPH_PAYMENT_COORDINATOR, 34)
	conn.ExecSQL(`INSERT INTO subgraphs.subgraph_version VALUES ($1, $2, $3)`, "anytext", TEST_SUBGRAPH_DELEGATION_MANAGER, 34)

	conn.ExecSQL(`CREATE SCHEMA IF NOT EXISTS sgd34`)

	schema_name := "sgd34"
	status := "current"
	conn.ExecSQL(`CREATE SCHEMA IF NOT EXISTS info`)
	conn.ExecSQL(`CREATE TABLE IF NOT EXISTS info.subgraph_info(schema_name text, name text, status text)`)
	conn.ExecSQL(`INSERT INTO info.subgraph_info VALUES ($1, $2, $3)`, schema_name, TEST_SUBGRAPH_CLAIMING_MANAGER, status)
	conn.ExecSQL(`INSERT INTO info.subgraph_info VALUES ($1, $2, $3)`, schema_name, TEST_SUBGRAPH_PAYMENT_COORDINATOR, status)
	conn.ExecSQL(`INSERT INTO info.subgraph_info VALUES ($1, $2, $3)`, schema_name, TEST_SUBGRAPH_DELEGATION_MANAGER, status)

	conn.ExecSQL(`CREATE TABLE IF NOT EXISTS sgd34."poi2$"(block_range int4range)`)
	conn.ExecSQL(`INSERT INTO sgd34."poi2$" VALUES ($1)`, "[100000001,)")
	conn.ExecSQL(`INSERT INTO sgd34."poi2$" VALUES ($1)`, "[100000,100000001)")

	conn.ExecSQL(`CREATE TYPE registration_status AS ENUM ('REGISTERED', 'UNREGISTERED')`)
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
