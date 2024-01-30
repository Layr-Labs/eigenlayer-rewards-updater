package common

import (
	"context"
	"fmt"
	"net/url"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
)

const (
	maxPoolSize = 10
)

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

func MustCreateConnection(connString string) *pgxpool.Pool {
	cfg, err := pgxpool.ParseConfig(connString)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect to DB")
	}

	dbPool, err := pgxpool.NewWithConfig(context.Background(), cfg)
	if err != nil {
		log.Err(err).Msg("Could not create connection to:" + connString)
		panic(err)
	}

	err = dbPool.Ping(context.Background())
	if err != nil {
		log.Err(err).Msg("Could not ping to:" + connString)
		panic(err)
	}

	return dbPool
}
