package common

import (
	"context"
	"fmt"
	"net/url"
	"os"

	pgxdecimal "github.com/jackc/pgx-shopspring-decimal"
	pgxzerolog "github.com/jackc/pgx-zerolog"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/tracelog"
	"github.com/rs/zerolog/log"
)

var maxPoolSize = 10

// used this for webserver to ensure it fails fast if cannot connect to postgres
func CreateConnectionAndPingOrDie(connString string) *pgxpool.Pool {
	dbPool := CreateConnectionOrDie(connString)

	err := dbPool.Ping(context.Background())
	if err != nil {
		log.Err(err).Msg("Could not ping to:" + connString)
		panic(err)
	}

	return dbPool
}

// used this for docker integration tests, rather than CreateConnectionAndPingOrDie
// because docker postgres takes time to start up
func CreateConnectionOrDie(connString string) *pgxpool.Pool {
	dbPool, err := pgxpool.NewWithConfig(context.Background(), CreateConnectionConfig(connString))
	if err != nil {
		log.Err(err).Msg("Could not create connection to:" + connString)
		panic(err)
	}

	return dbPool
}

/* used this for webserver to ensure it fails fast if cannot connect to postgres
 * @config is a pgxpool config, you can set more options here, like max connections
 *   e.g. config.MaxConns = 20
 */
func CreateConnectionAndPingOrDieWithConfig(config *pgxpool.Config) *pgxpool.Pool {
	dbPool := CreateConnectionOrDieWithConfig(config)

	err := dbPool.Ping(context.Background())
	if err != nil {
		log.Err(err).Msg(fmt.Sprintf("Could not ping to Postgres with config: %v", config))
		panic(err)
	}

	return dbPool
}

// used this for docker integration tests, rather than CreateConnectionAndPingOrDie
// because docker postgres takes time to start up
// @config is a pgxpool config, you can set more options here, like max connections
func CreateConnectionOrDieWithConfig(config *pgxpool.Config) *pgxpool.Pool {
	dbPool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Err(err).Msg(fmt.Sprintf("Could not create connection with config: %v", config))
		panic(err)
	}

	return dbPool
}

func CreateConnectionConfig(connString string) *pgxpool.Config {
	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		panic(err)
	}
	config.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		// register decimal type
		pgxdecimal.Register(conn.TypeMap())
		return nil
	}

	if os.Getenv("ENV") == "dev" {
		logger := pgxzerolog.NewLogger(log.Logger)
		config.ConnConfig.Tracer = &tracelog.TraceLog{
			Logger:   logger,
			LogLevel: tracelog.LogLevelTrace,
		}
	}

	return config
}

func CreateConnectionString(user, password, host, port, db string) string {
	// URL-encode the user and password to ensure special characters do not break the connection string
	encodedUser := url.QueryEscape(user)
	encodedPassword := url.QueryEscape(password)

	return fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?pool_max_conns=%d",
		encodedUser,
		encodedPassword,
		host,
		port,
		db,
		maxPoolSize,
	)
}
