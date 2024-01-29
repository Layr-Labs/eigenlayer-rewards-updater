package calculator

import (
	"context"
	"fmt"
	"math/big"

	contractIPaymentCoordinator "github.com/Layr-Labs/eigenlayer-payment-updater/bindings/IPaymentCoordinator"
	"github.com/Layr-Labs/eigenlayer-payment-updater/common"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/shopspring/decimal"
)

type PaymentCalculatorDataService interface {
	// GetPaymentsCalculatedUntilTimestamp returns the timestamp until which payments have been calculated
	GetPaymentsCalculatedUntilTimestamp(ctx context.Context) (*big.Int, error)
	// GetRangePaymentsWithOverlappingRange returns all range payments that overlap with the given range
	GetRangePaymentsWithOverlappingRange(startTimestamp, endTimestamp *big.Int) ([]*contractIPaymentCoordinator.IPaymentCoordinatorRangePayment, error)
	// GetDistributionsAtTimestamp returns the distributions of all tokens at a given timestamp
	GetDistributionsAtTimestamp(timestamp *big.Int) (map[gethcommon.Address]*common.Distribution, error)
	// GetOperatorSetAtTimestamp returns the operator set at a given timestamps
	GetOperatorSetAtTimestamp(avs gethcommon.Address, timestamp *big.Int) (common.OperatorSet, error)
}

type PaymentCalculatorDataServiceImpl struct {
	dbpool                     *pgxpool.Pool
	schemaService              *common.SubgraphSchemaService
	claimingManagerSubgraph    string
	paymentCoordinatorSubgraph string
	subgraphProvider           common.SubgraphProvider
}

func NewPaymentCalculatorDataService(
	dbpool *pgxpool.Pool,
	schemaService *common.SubgraphSchemaService,
	claimingManagerSubgraph string,
	paymentCoordinatorSubgraph string,
	subgraphProvider common.SubgraphProvider,
) PaymentCalculatorDataService {
	return &PaymentCalculatorDataServiceImpl{
		dbpool:                     dbpool,
		schemaService:              schemaService,
		claimingManagerSubgraph:    claimingManagerSubgraph,
		paymentCoordinatorSubgraph: paymentCoordinatorSubgraph,
		subgraphProvider:           subgraphProvider,
	}
}

var paymentsCalculatedUntilQuery string = `
	SELECT payments_calculated_until_timestamp
	FROM %s.root_submitted
	ORDER BY payments_calculated_until_timestamp DESC
	LIMIT 1
`

func (s *PaymentCalculatorDataServiceImpl) GetPaymentsCalculatedUntilTimestamp(ctx context.Context) (*big.Int, error) {
	schemaID, err := s.schemaService.GetSubgraphSchema(ctx, s.claimingManagerSubgraph, s.subgraphProvider)
	if err != nil {
		return nil, err
	}

	formattedQuery := fmt.Sprintf(paymentsCalculatedUntilQuery, schemaID)
	row := s.dbpool.QueryRow(ctx, formattedQuery)

	var resDecimal decimal.Decimal
	err = row.Scan(
		&resDecimal,
	)
	return resDecimal.BigInt(), nil
}

var overlappingRangePaymentsQuery string = `
	SELECT *
	FROM %s.range_payment_created
	WHERE start_range_timestamp < $1 AND end_range_timestamp > $2
	LIMIT 1
`

func (s *PaymentCalculatorDataServiceImpl) GetRangePaymentsWithOverlappingRange(startTimestamp, endTimestamp *big.Int) ([]*contractIPaymentCoordinator.IPaymentCoordinatorRangePayment, error) {
	// schemaID, err := s.schemaService.GetSubgraphSchema(context.Background(), s.paymentCoordinatorSubgraph, s.subgraphProvider)
	// if err != nil {
	// 	return nil, err
	// }

	// formattedQuery := fmt.Sprintf(overlappingRangePaymentsQuery, schemaID)
	// rows, err := s.dbpool.Query(context.Background(), formattedQuery, endTimestamp, startTimestamp)
	// if err != nil {
	// 	return nil, err
	// }

	// var rangePayments []*contractIPaymentCoordinator.IPaymentCoordinatorRangePayment
	// for rows.Next() {
	// 	var (
	// 		id                  int64
	// 		amount              decimal.Decimal
	// 		token               gethcommon.Address
	// 		startRangeTimestamp decimal.Decimal
	// 		endRangeTimestamp   decimal.Decimal
	// 	)
	// 	err := rows.Scan(
	// 		&id,
	// 		&amount,
	return nil, nil
}

func (s *PaymentCalculatorDataServiceImpl) GetDistributionsAtTimestamp(timestamp *big.Int) (map[gethcommon.Address]*common.Distribution, error) {
	return nil, nil
}

func (s *PaymentCalculatorDataServiceImpl) GetOperatorSetAtTimestamp(avs gethcommon.Address, timestamp *big.Int) (common.OperatorSet, error) {
	return common.OperatorSet{}, nil
}
