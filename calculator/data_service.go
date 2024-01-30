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

// todo: set this to the global default AVS
var GLOBAL_DEFAULT_AVS = gethcommon.HexToAddress("0x0000000000000000000000000000000000000000")

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

const (
	claimingManagerSubgraph    = "claiming-manager-raw-events"
	paymentCoordinatorSubgraph = "payment-coordinator-raw-events"
)

type PaymentCalculatorDataServiceImpl struct {
	dbpool           *pgxpool.Pool
	schemaService    *common.SubgraphSchemaService
	subgraphProvider common.SubgraphProvider
}

func NewPaymentCalculatorDataService(
	dbpool *pgxpool.Pool,
	schemaService *common.SubgraphSchemaService,
	subgraphProvider common.SubgraphProvider,
) PaymentCalculatorDataService {
	return &PaymentCalculatorDataServiceImpl{
		dbpool:           dbpool,
		schemaService:    schemaService,
		subgraphProvider: subgraphProvider,
	}
}

var paymentsCalculatedUntilQuery string = `
	SELECT payments_calculated_until_timestamp
	FROM %s.root_submitted
	ORDER BY payments_calculated_until_timestamp DESC
	LIMIT 1
`

func (s *PaymentCalculatorDataServiceImpl) GetPaymentsCalculatedUntilTimestamp(ctx context.Context) (*big.Int, error) {
	schemaID, err := s.schemaService.GetSubgraphSchema(ctx, claimingManagerSubgraph, s.subgraphProvider)
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

// gets all range payments that overlap with the given range ($1, $2)
var overlappingRangePaymentsQuery string = `
	SELECT range_payment_avs, range_payment_strategy, range_payment_token, range_payment_amount, range_payment_start_range_timestamp, range_payment_end_range_timestamp
	FROM %s.range_payment_created
	WHERE range_payment_start_range_timestamp < $2 AND range_payment_end_range_timestamp > $1
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
	schemaID, err := s.schemaService.GetSubgraphSchema(context.Background(), paymentCoordinatorSubgraph, s.subgraphProvider)
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
	// if the data directory doesn't exist, create it and return empty map
	_, err := os.Stat("./data")
	if os.IsNotExist(err) {
		err = os.Mkdir("./data", 0755)
		if err != nil {
			return nil, err
		}
		return make(map[gethcommon.Address]*common.Distribution), nil
	}
	if err != nil {
		return nil, err
	}

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

func (s *PaymentCalculatorDataServiceImpl) getCommissionForAVSAtTimestamp(timestamp *big.Int, avs gethcommon.Address, operators []gethcommon.Address) (map[gethcommon.Address]*big.Int, error) {
	commissions, err := s.getCommissionForAVSAtTimestampWithoutGlobalDefault(timestamp, avs, operators)
	if err != nil {
		return nil, err
	}

	// for all operators without a specific commission, get their global default commission
	operatorsWithoutSpecificCommission := make([]gethcommon.Address, 0)
	for _, operator := range operators {
		if _, ok := commissions[operator]; !ok {
			operatorsWithoutSpecificCommission = append(operatorsWithoutSpecificCommission, operator)
		}
	}

	commissionsWithGlobalDefault, err := s.getCommissionForAVSAtTimestampWithoutGlobalDefault(timestamp, GLOBAL_DEFAULT_AVS, operatorsWithoutSpecificCommission)
	if err != nil {
		return nil, err
	}

	// merge the two maps
	for operator, commission := range commissionsWithGlobalDefault {
		commissions[operator] = commission
	}

	return commissions, nil
}

var commissionAtTimestampQuery string = `
	SELECT DISTINCT ON (operator) operator, commission_bips
	FROM %s.commission_set
	WHERE block_timestamp <= $1 AND avs = $2 AND operator in ($3)
	ORDER BY account, block_timestamp DESC`

func (s *PaymentCalculatorDataServiceImpl) getCommissionForAVSAtTimestampWithoutGlobalDefault(timestamp *big.Int, avs gethcommon.Address, operators []gethcommon.Address) (map[gethcommon.Address]*big.Int, error) {
	// get the schema id for the claiming manager subgraph
	schemaID, err := s.schemaService.GetSubgraphSchema(context.Background(), claimingManagerSubgraph, s.subgraphProvider)
	if err != nil {
		return nil, err
	}

	// format the query
	formattedQuery := fmt.Sprintf(commissionAtTimestampQuery, schemaID)

	// query the database
	rows, err := s.dbpool.Query(context.Background(), formattedQuery, timestamp, avs, operators)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// create a map of operator to commission
	commissions := make(map[gethcommon.Address]*big.Int)
	for rows.Next() {
		var (
			operatorBytes []byte
			commission    decimal.Decimal
		)

		err := rows.Scan(
			&operatorBytes,
			&commission,
		)
		if err != nil {
			return nil, err
		}

		operator := gethcommon.HexToAddress(hex.EncodeToString(operatorBytes))

		commissions[operator] = commission.BigInt()
	}

	// for all operators that don't have a commission, set it to 0
	for _, operator := range operators {
		if _, ok := commissions[operator]; !ok {
			commissions[operator] = big.NewInt(0)
		}
	}

	return commissions, nil
}

var claimersAtTimestampQuery string = `
	SELECT DISTINCT ON (account) account, claimer
	FROM %s.claimer_set
	WHERE block_timestamp <= $1 AND account in ($2)
	ORDER BY account, block_timestamp DESC`

func (s *PaymentCalculatorDataServiceImpl) getClaimersAtTimestamp(timestamp *big.Int, accounts []gethcommon.Address) (map[gethcommon.Address]gethcommon.Address, error) {
	// get the schema id for the claiming manager subgraph
	schemaID, err := s.schemaService.GetSubgraphSchema(context.Background(), claimingManagerSubgraph, s.subgraphProvider)
	if err != nil {
		return nil, err
	}

	// format the query
	formattedQuery := fmt.Sprintf(claimersAtTimestampQuery, schemaID)

	// query the database
	rows, err := s.dbpool.Query(context.Background(), formattedQuery, timestamp, accounts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// create a map of account to claimer
	claimers := make(map[gethcommon.Address]gethcommon.Address)
	for rows.Next() {
		var (
			accountBytes []byte
			claimerBytes []byte
		)

		err := rows.Scan(
			&accountBytes,
			&claimerBytes,
		)
		if err != nil {
			return nil, err
		}

		account := gethcommon.HexToAddress(hex.EncodeToString(accountBytes))
		claimer := gethcommon.HexToAddress(hex.EncodeToString(claimerBytes))

		claimers[account] = claimer
	}

	// set the claimer of any account that doesn't have one to the account itself
	for _, account := range accounts {
		if _, ok := claimers[account]; !ok {
			claimers[account] = account
		}
	}

	return claimers, nil
}

// get all stakers that have an entry in the staker_delegated table with the given operator with a block timestamp higher than the entry in the staker_undelegated table for the same staker
var stakerSetAtTimestampQuery string = `
	SELECT staker as stakerAddress 
	FROM %s.staker_delegated
	WHERE operator = $1 AND block_timestamp > (
		SELECT block_timestamp FROM %s.staker_undelegated
		WHERE operator = $1 AND staker = stakerAddress
		ORDER BY block_timestamp DESC
		LIMIT 1
	)`
