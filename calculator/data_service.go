package calculator

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"os"

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
	// SetDistributionsAtTimestamp sets the distributions of all tokens at a given timestamp
	SetDistributionsAtTimestamp(timestamp *big.Int, distributions map[gethcommon.Address]*common.Distribution) error
	// GetOperatorSetForStrategyAtTimestamp returns the operator set for a given strategy at a given timestamps
	GetOperatorSetForStrategyAtTimestamp(avs gethcommon.Address, strategy gethcommon.Address, timestamp *big.Int) (common.OperatorSet, error)
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

// gets all range payments that overlap with the given range ($1, $2)
var overlappingRangePaymentsQuery string = `
	SELECT range_payment_avs, range_payment_strategy, range_payment_token, range_payment_amount, range_payment_start_range_timestamp, range_payment_end_range_timestamp
	FROM %s.range_payment_created
	WHERE start_range_timestamp < $2 AND end_range_timestamp > $1
	LIMIT 1
`

//  id                                  | bytea   |           | not null |
//  range_payment_avs                   | bytea   |           | not null |
//  range_payment_strategy              | bytea   |           | not null |
//  range_payment_token                 | bytea   |           | not null |
//  range_payment_amount                | numeric |           | not null |
//  range_payment_start_range_timestamp | numeric |           | not null |
//  range_payment_end_range_timestamp   | numeric |           | not null |
//  block_number                        | numeric |           | not null |
//  block_timestamp                     | numeric |           | not null |
//  transaction_hash                    | bytea   |           | not null |

func (s *PaymentCalculatorDataServiceImpl) GetRangePaymentsWithOverlappingRange(startTimestamp, endTimestamp *big.Int) ([]*contractIPaymentCoordinator.IPaymentCoordinatorRangePayment, error) {
	schemaID, err := s.schemaService.GetSubgraphSchema(context.Background(), s.paymentCoordinatorSubgraph, s.subgraphProvider)
	if err != nil {
		return nil, err
	}

	formattedQuery := fmt.Sprintf(overlappingRangePaymentsQuery, schemaID)
	rows, err := s.dbpool.Query(context.Background(), formattedQuery, startTimestamp, endTimestamp)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rangePayments []*contractIPaymentCoordinator.IPaymentCoordinatorRangePayment
	for rows.Next() {
		var (
			rangePaymentAVSBytes                   []byte
			rangePaymentStrategyBytes              []byte
			rangePaymentTokenBytes                 []byte
			rangePaymentAmountDecimal              decimal.Decimal
			rangePaymentStartRangeTimestampDecimal decimal.Decimal
			rangePaymentEndRangeTimestampDecimal   decimal.Decimal
		)

		err := rows.Scan(
			&rangePaymentAVSBytes,
			&rangePaymentStrategyBytes,
			&rangePaymentTokenBytes,
			&rangePaymentAmountDecimal,
			&rangePaymentStartRangeTimestampDecimal,
			&rangePaymentEndRangeTimestampDecimal,
		)
		if err != nil {
			return nil, err
		}

		rangePayment := &contractIPaymentCoordinator.IPaymentCoordinatorRangePayment{
			Avs:                 gethcommon.HexToAddress(hex.EncodeToString(rangePaymentAVSBytes)),
			Strategy:            gethcommon.HexToAddress(hex.EncodeToString(rangePaymentStrategyBytes)),
			Token:               gethcommon.HexToAddress(hex.EncodeToString(rangePaymentTokenBytes)),
			Amount:              rangePaymentAmountDecimal.BigInt(),
			StartRangeTimestamp: rangePaymentStartRangeTimestampDecimal.BigInt(),
			EndRangeTimestamp:   rangePaymentEndRangeTimestampDecimal.BigInt(),
		}

		rangePayments = append(rangePayments, rangePayment)
	}
	return rangePayments, nil
}

func (s *PaymentCalculatorDataServiceImpl) GetDistributionsAtTimestamp(timestamp *big.Int) (map[gethcommon.Address]*common.Distribution, error) {
	// read from data/distributions_{timestamp}.json
	file, err := os.ReadFile(fmt.Sprintf("data/distributions_%d.json", timestamp))
	if err != nil {
		return nil, err
	}

	// deserialize from json
	var distributions map[gethcommon.Address]*common.Distribution
	err = json.Unmarshal(file, &distributions)
	if err != nil {
		return nil, err
	}

	return distributions, nil
}

func (s *PaymentCalculatorDataServiceImpl) SetDistributionsAtTimestamp(timestamp *big.Int, distributions map[gethcommon.Address]*common.Distribution) error {
	// seralize to json and write to data/distributions_{timestamp}.json
	marshalledDistributions, err := json.Marshal(distributions)
	if err != nil {
		return err
	}

	// write to file
	err = os.WriteFile(fmt.Sprintf("data/distributions_%d.json", timestamp), marshalledDistributions, 0644)
	if err != nil {
		return err
	}

	return err
}

// type OperatorSet struct {
// 	TotalStakedStrategyShares map[gethcommon.Address]*big.Int
// 	Operators                 []Operator
// }

// type Earner struct {
// 	Claimer gethcommon.Address
// }

// type Operator struct {
// 	Earner
// 	Address                      gethcommon.Address
// 	Commissions                  map[gethcommon.Address]*big.Int
// 	TotalDelegatedStrategyShares map[gethcommon.Address]*big.Int
// 	Stakers                      []Staker
// }

// type Staker struct {
// 	Earner
// 	Address gethcommon.Address
// 	Shares  map[gethcommon.Address]*big.Int
// }

func (s *PaymentCalculatorDataServiceImpl) GetOperatorSetForStrategyAtTimestamp(avs gethcommon.Address, strategy gethcommon.Address, timestamp *big.Int) (common.OperatorSet, error) {
	return common.OperatorSet{}, nil
}

func 
