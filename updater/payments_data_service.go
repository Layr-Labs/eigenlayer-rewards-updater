package updater

import (
	"context"
	"fmt"
	"math/big"

	"github.com/Layr-Labs/eigenlayer-payment-updater/common"
	"github.com/Layr-Labs/eigenlayer-payment-updater/common/utils"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/shopspring/decimal"
)

type PaymentsDataService interface {
	// GetPaymentsCalculatedUntilTimestamp returns the timestamp until which payments have been calculated
	GetPaymentsCalculatedUntilTimestamp(ctx context.Context) (*big.Int, error)
}

type PaymentsDataServiceImpl struct {
	dbpool        *pgxpool.Pool
	schemaService *common.SubgraphSchemaService
}

func NewPaymentsDataService(
	dbpool *pgxpool.Pool,
	schemaService *common.SubgraphSchemaService,
) PaymentsDataService {
	return &PaymentsDataServiceImpl{
		dbpool:        dbpool,
		schemaService: schemaService,
	}
}

func NewPaymentsDataServiceImpl(
	dbpool *pgxpool.Pool,
	schemaService *common.SubgraphSchemaService,
) *PaymentsDataServiceImpl {
	return &PaymentsDataServiceImpl{
		dbpool:        dbpool,
		schemaService: schemaService,
	}
}

func (s *PaymentsDataServiceImpl) GetPaymentsCalculatedUntilTimestamp(ctx context.Context) (*big.Int, error) {
	schemaID, err := s.schemaService.GetSubgraphSchema(ctx, utils.SUBGRAPH_CLAIMING_MANAGER)
	if err != nil {
		return nil, err
	}

	formattedQuery := fmt.Sprintf(paymentsCalculatedUntilQuery, schemaID)
	row := s.dbpool.QueryRow(ctx, formattedQuery)

	var resDecimal decimal.Decimal
	err = row.Scan(
		&resDecimal,
	)
	if err != nil {
		return nil, err
	}

	return resDecimal.BigInt(), nil
}
