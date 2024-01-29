package common

import (
	"context"
	"fmt"

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
	DefaultSubgraphProviderName = "default"
	SatsumaSubgraphProviderName = "satsuma"

	defaultSchemaIDQuery = `
		SELECT schema_name
		FROM info.subgraph_info
		WHERE status = 'current' AND name = $1
	`

	satsumaSchemaIDQuery = `
		SELECT entity_schema
		FROM satsuma.subgraph_schema
		WHERE satsuma_subgraph_name = $1
	`
)

type SubgraphSchemaService struct {
	dbpool *pgxpool.Pool
}

func NewSubgraphSchemaService(dbpool *pgxpool.Pool) *SubgraphSchemaService {
	return &SubgraphSchemaService{
		dbpool: dbpool,
	}
}

func (s *SubgraphSchemaService) GetSubgraphSchema(ctx context.Context, subgraphName string, provider SubgraphProvider) (string, error) {
	var (
		query      string
		schemaName string
	)
	switch provider {
	case DefaultSubgraphProvider:
		query = defaultSchemaIDQuery
	case SatsumaSubgraphProvider:
		query = satsumaSchemaIDQuery
	default:
		return "", fmt.Errorf("invalid subgraph provider: %d", provider)
	}

	err := s.dbpool.QueryRow(ctx, query, subgraphName).Scan(&schemaName)
	if err != nil {
		log.Err(err).Msg("error while getting schema id for subgraph: " + subgraphName)
		return "", err
	}

	return schemaName, nil
}

func ToSubgraphProvider(provider string) (SubgraphProvider, error) {
	switch provider {
	case DefaultSubgraphProviderName:
		return DefaultSubgraphProvider, nil
	case SatsumaSubgraphProviderName:
		return SatsumaSubgraphProvider, nil
	default:
		return UnknownSubgraphProvider, fmt.Errorf("invalid subgraph provider: %s", provider)
	}
}
