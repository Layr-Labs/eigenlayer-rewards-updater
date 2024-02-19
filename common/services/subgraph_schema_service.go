package services

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
)

type SubgraphProvider int

const (
	DefaultSubgraphProvider SubgraphProvider = iota
	SatsumaSubgraphProvider
	UnknownSubgraphProvider
)

const (
	satsumaSchemaIDQuery = `
		SELECT entity_schema
		FROM satsuma.subgraph_schema
		WHERE satsuma_subgraph_name = $1
	`
)

var subgraphNameSuffix = map[string]string{
	"test":             "-test",
	"dev-goerli":       "-dev-goerli",
	"testnet-goerli":   "-goerli",
	"testnet-holesky":  "-holesky",
	"mainnet-ethereum": "-mainnet",
}

type SubgraphSchemaService struct {
	env    string
	dbpool *pgxpool.Pool
}

func NewSubgraphSchemaService(env string, dbpool *pgxpool.Pool) *SubgraphSchemaService {
	return &SubgraphSchemaService{
		env:    env,
		dbpool: dbpool,
	}
}

func (s *SubgraphSchemaService) GetSubgraphSchema(ctx context.Context, subgraphName string) (string, error) {
	var (
		query      = satsumaSchemaIDQuery
		schemaName string
	)

	subgraphName += subgraphNameSuffix[s.env]

	err := s.dbpool.QueryRow(ctx, query, subgraphName).Scan(&schemaName)
	if err != nil {
		log.Err(err).Msg("error while getting schema id for subgraph: " + subgraphName)
		return "", err
	}

	return schemaName, nil
}
