package calculator

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"strings"
	"sync"

	contractIEigenPodManager "github.com/Layr-Labs/eigenlayer-payment-updater/bindings/IEigenPodManager"
	contractIPaymentCoordinator "github.com/Layr-Labs/eigenlayer-payment-updater/bindings/IPaymentCoordinator"
	contractIStrategyManager "github.com/Layr-Labs/eigenlayer-payment-updater/bindings/IStrategyManager"
	"github.com/Layr-Labs/eigenlayer-payment-updater/common"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
	"github.com/shopspring/decimal"
)

// todo: set this to the global default AVS
var GLOBAL_DEFAULT_AVS = gethcommon.HexToAddress("0x0000000000000000000000000000000000000000")

// multicall has the same address on all networks
var MULTICALL3_ADDRESS = gethcommon.HexToAddress("0xcA11bde05977b3631167028862bE2a173976CA11")

// strategy manager address
var STRATEGY_MANAGER_ADDRESS = gethcommon.HexToAddress("0x779d1b5315df083e3F9E94cB495983500bA8E907")

// eigen pod manager address
var EIGEN_POD_MANAGER_ADDRESS = gethcommon.HexToAddress("0xa286b84C96aF280a49Fe1F40B9627C2A2827df41")

// beacon chain eth strategy address
var BEACON_CHAIN_ETH_STRATEGY_ADDRESS = gethcommon.HexToAddress("0xbeaC0eeEeeeeEEeEeEEEEeeEEeEeeeEeeEEBEaC0")

// delegation manager address
var DELEGATION_MANAGER_ADDRESS = gethcommon.HexToAddress("0x1b7b8F6b258f95Cf9596EabB9aa18B62940Eb0a8")

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
	GetOperatorSetForStrategyAtTimestamp(timestamp *big.Int, avs gethcommon.Address, strategy gethcommon.Address) (*common.OperatorSet, error)

	GetBlockNumberAtTimestamp(timestamp *big.Int) (*big.Int, error)
}

type PaymentCalculatorDataServiceImpl struct {
	dbpool                     *pgxpool.Pool
	schemaService              *common.SubgraphSchemaService
	subgraphProvider           common.SubgraphProvider
	claimingManagerSubgraph    string
	paymentCoordinatorSubgraph string
	delegationManagerSubgraph  string
	ethClient                  *ethclient.Client
}

func NewPaymentCalculatorDataService(
	dbpool *pgxpool.Pool,
	schemaService *common.SubgraphSchemaService,
	subgraphProvider common.SubgraphProvider,
	claimingManagerSubgraph string,
	paymentCoordinatorSubgraph string,
	delegationManagerSubgraph string,
	ethClient *ethclient.Client,
) PaymentCalculatorDataService {
	return &PaymentCalculatorDataServiceImpl{
		dbpool:                     dbpool,
		schemaService:              schemaService,
		subgraphProvider:           subgraphProvider,
		claimingManagerSubgraph:    claimingManagerSubgraph,
		paymentCoordinatorSubgraph: paymentCoordinatorSubgraph,
		delegationManagerSubgraph:  delegationManagerSubgraph,
		ethClient:                  ethClient,
	}
}

func NewPaymentCalculatorDataServiceImpl(
	dbpool *pgxpool.Pool,
	schemaService *common.SubgraphSchemaService,
	subgraphProvider common.SubgraphProvider,
	claimingManagerSubgraph string,
	paymentCoordinatorSubgraph string,
	delegationManagerSubgraph string,
	ethClient *ethclient.Client,
) *PaymentCalculatorDataServiceImpl {
	return &PaymentCalculatorDataServiceImpl{
		dbpool:                     dbpool,
		schemaService:              schemaService,
		subgraphProvider:           subgraphProvider,
		claimingManagerSubgraph:    claimingManagerSubgraph,
		paymentCoordinatorSubgraph: paymentCoordinatorSubgraph,
		delegationManagerSubgraph:  delegationManagerSubgraph,
		ethClient:                  ethClient,
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

func (s *PaymentCalculatorDataServiceImpl) GetOperatorSetForStrategyAtTimestamp(timestamp *big.Int, avs gethcommon.Address, strategy gethcommon.Address) (*common.OperatorSet, error) {
	operatorSet := common.OperatorSet{}
	operatorSet.TotalStakedStrategyShares = big.NewInt(0)

	// get all operators for the given strategy at the given timestamp
	operatorAddresses, err := s.GetOperatorAddressesForAVSAtTimestamp(timestamp, avs, strategy)
	if err != nil {
		return nil, err
	}

	operatorSet.Operators = make([]common.Operator, len(operatorAddresses))

	// get the commission for each operator
	commissions, err := s.GetCommissionForAVSAtTimestamp(timestamp, avs, operatorAddresses)
	if err != nil {
		return nil, err
	}

	// loop thru each operator and get their staker sets
	for i, operatorAddress := range operatorAddresses {
		operatorSet.Operators[i].Address = operatorAddress
		operatorSet.Operators[i].Commission = commissions[operatorAddress]
		operatorSet.Operators[i].TotalDelegatedStrategyShares = big.NewInt(0)

		// get the stakers of the operator
		stakers, err := s.GetStakersDelegatedToOperatorAtTimestamp(timestamp, operatorAddress)
		if err != nil {
			return nil, err
		}

		operatorSet.Operators[i].Stakers = make([]common.Staker, len(stakers))

		// get the claimers of each staker and the operator
		claimers, err := s.GetClaimersAtTimestamp(timestamp, append(stakers, operatorAddress))
		if err != nil {
			return nil, err
		}

		// get the blocknumber of the block at the given timestamp
		blockNumber, err := s.GetBlockNumberAtTimestamp(timestamp)
		if err != nil {
			return nil, err
		}

		// get the shares of each staker
		strategyShareMap, err := s.GetSharesOfStakersAtBlockNumber(blockNumber, strategy, stakers)
		if err != nil {
			return nil, err
		}

		// loop thru each staker and get their shares
		for j, stakerAddress := range stakers {
			operatorSet.Operators[i].Stakers[j].Address = stakerAddress
			operatorSet.Operators[i].Stakers[j].Claimer = claimers[stakerAddress]
			operatorSet.Operators[i].Stakers[j].StrategyShares = strategyShareMap[stakerAddress]

			operatorSet.Operators[i].Claimer = claimers[operatorAddress]
			// add the staker's shares to the operator's total delegated strategy shares
			operatorSet.Operators[i].TotalDelegatedStrategyShares = operatorSet.Operators[i].TotalDelegatedStrategyShares.Add(operatorSet.Operators[i].TotalDelegatedStrategyShares, strategyShareMap[stakerAddress])
		}

		// add the operator's total delegated strategy shares to the operator set's total staked strategy shares
		operatorSet.TotalStakedStrategyShares = operatorSet.TotalStakedStrategyShares.Add(operatorSet.TotalStakedStrategyShares, operatorSet.Operators[i].TotalDelegatedStrategyShares)
	}

	return &operatorSet, nil
}

func (s *PaymentCalculatorDataServiceImpl) GetOperatorAddressesForAVSAtTimestamp(timestamp *big.Int, avs gethcommon.Address, strategy gethcommon.Address) ([]gethcommon.Address, error) {
	// TODO
	return nil, nil
}

func (s *PaymentCalculatorDataServiceImpl) GetCommissionForAVSAtTimestamp(timestamp *big.Int, avs gethcommon.Address, operators []gethcommon.Address) (map[gethcommon.Address]*big.Int, error) {
	commissions, err := s.GetCommissionForAVSAtTimestampWithoutGlobalDefault(timestamp, avs, operators)
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

	commissionsWithGlobalDefault, err := s.GetCommissionForAVSAtTimestampWithoutGlobalDefault(timestamp, GLOBAL_DEFAULT_AVS, operatorsWithoutSpecificCommission)
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

func (s *PaymentCalculatorDataServiceImpl) GetCommissionForAVSAtTimestampWithoutGlobalDefault(timestamp *big.Int, avs gethcommon.Address, operators []gethcommon.Address) (map[gethcommon.Address]*big.Int, error) {
	// get the schema id for the claiming manager subgraph
	schemaID, err := s.schemaService.GetSubgraphSchema(context.Background(), s.claimingManagerSubgraph, s.subgraphProvider)
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

func (s *PaymentCalculatorDataServiceImpl) GetClaimersAtTimestamp(timestamp *big.Int, accounts []gethcommon.Address) (map[gethcommon.Address]gethcommon.Address, error) {
	// get the schema id for the claiming manager subgraph
	schemaID, err := s.schemaService.GetSubgraphSchema(context.Background(), s.claimingManagerSubgraph, s.subgraphProvider)
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
WITH latest_undelegations AS (
    SELECT
        DISTINCT ON (staker) staker,
        block_timestamp AS undelegation_timestamp
    FROM
        %s.staker_undelegated
    WHERE
        encode(operator, 'hex') = $1
    ORDER BY
        staker,
        block_timestamp DESC
), latest_delegations AS (
    SELECT
        DISTINCT ON (staker) staker,
        block_timestamp AS delegation_timestamp
    FROM
        %s.staker_delegated
    WHERE
        encode(operator, 'hex') = $1
    ORDER BY
        staker,
        block_timestamp DESC
)
SELECT
    ld.staker
FROM
    latest_delegations ld
LEFT JOIN
    latest_undelegations lu ON ld.staker = lu.staker
WHERE
    lu.staker IS NULL OR ld.delegation_timestamp > lu.undelegation_timestamp;`

func (s *PaymentCalculatorDataServiceImpl) GetStakersDelegatedToOperatorAtTimestamp(timestamp *big.Int, operator gethcommon.Address) ([]gethcommon.Address, error) {
	// get the schema id for the claiming manager subgraph
	schemaID, err := s.schemaService.GetSubgraphSchema(context.Background(), s.delegationManagerSubgraph, s.subgraphProvider)
	if err != nil {
		return nil, err
	}

	// format the query
	formattedQuery := fmt.Sprintf(stakerSetAtTimestampQuery, schemaID, schemaID)

	// query the database
	rows, err := s.dbpool.Query(context.Background(), formattedQuery, toSQLAddress(operator))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// create a list of stakers
	stakers := make([]gethcommon.Address, 0)
	for rows.Next() {
		var (
			stakerBytes []byte
		)

		err := rows.Scan(
			&stakerBytes,
		)
		if err != nil {
			return nil, err
		}

		staker := gethcommon.HexToAddress(hex.EncodeToString(stakerBytes))

		stakers = append(stakers, staker)
	}

	return stakers, nil
}

func (s *PaymentCalculatorDataServiceImpl) GetBlockNumberAtTimestamp(timestamp *big.Int) (*big.Int, error) {
	head, err := s.ethClient.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	var lo, hi = big.NewInt(0), head.Number
	for lo.Cmp(hi) < 0 {
		mid := new(big.Int).Add(lo, hi)
		mid.Div(mid, big.NewInt(2))

		header, err := s.ethClient.HeaderByNumber(context.Background(), mid)
		if err != nil {
			return nil, err
		}
		log.Info().Msgf("mid: %d, header time: %d, timestamp: %d", mid, header.Time, timestamp)
		if header.Time < timestamp.Uint64() {
			lo = mid.Add(mid, big.NewInt(1))
		} else {
			hi = mid
		}
	}

	return lo.Sub(lo, big.NewInt(1)), nil
}

func (s *PaymentCalculatorDataServiceImpl) GetSharesOfStakersAtBlockNumber(blockNumber *big.Int, strategy gethcommon.Address, stakers []gethcommon.Address) (map[gethcommon.Address]*big.Int, error) {
	if strategy == BEACON_CHAIN_ETH_STRATEGY_ADDRESS {
		return s.getSharesOfBeaconChainETHStrategyForStakersAtBlockNumber(blockNumber, stakers)
	} else {
		return s.getSharesOfStrategyManagerStrategyForStakersAtBlockNumber(blockNumber, strategy, stakers)
	}
}

func (s *PaymentCalculatorDataServiceImpl) getSharesOfStrategyManagerStrategyForStakersAtBlockNumber(blockNumber *big.Int, strategy gethcommon.Address, stakers []gethcommon.Address) (map[gethcommon.Address]*big.Int, error) {
	strategyManagerContract, err := contractIStrategyManager.NewContractIStrategyManager(STRATEGY_MANAGER_ADDRESS, s.ethClient)
	if err != nil {
		return nil, err
	}

	// TODO: make this a batch call
	return fillMapFromAddressToBigIntParallel(stakers, func(staker gethcommon.Address) (*big.Int, error) {
		shares, err := strategyManagerContract.StakerStrategyShares(&bind.CallOpts{BlockNumber: blockNumber}, staker, strategy)
		if err != nil {
			return nil, err
		}

		return shares, nil
	})
}

func (s *PaymentCalculatorDataServiceImpl) getSharesOfBeaconChainETHStrategyForStakersAtBlockNumber(blockNumber *big.Int, stakers []gethcommon.Address) (map[gethcommon.Address]*big.Int, error) {
	eigenPodManagerContract, err := contractIEigenPodManager.NewContractIEigenPodManager(EIGEN_POD_MANAGER_ADDRESS, s.ethClient)
	if err != nil {
		return nil, err
	}

	// TODO: make this a batch call
	return fillMapFromAddressToBigIntParallel(stakers, func(staker gethcommon.Address) (*big.Int, error) {
		shares, err := eigenPodManagerContract.PodOwnerShares(&bind.CallOpts{BlockNumber: blockNumber}, staker)
		if err != nil {
			return nil, err
		}

		return shares, nil
	})
}

func fillMapFromAddressToBigIntParallel(addresses []gethcommon.Address, getValue func(gethcommon.Address) (*big.Int, error)) (map[gethcommon.Address]*big.Int, error) {
	resMap := make(map[gethcommon.Address]*big.Int)
	var mu sync.Mutex     // Used to safely write to the map
	var wg sync.WaitGroup // Used to wait for all goroutines to finish

	errChan := make(chan error, len(addresses)) // Channel to collect errors
	resChan := make(chan struct {
		addr  gethcommon.Address
		value *big.Int
	}, len(addresses)) // Channel to collect

	for _, addr := range addresses {
		wg.Add(1)
		go func(addr gethcommon.Address) {
			defer wg.Done()
			value, err := getValue(addr)
			if err != nil {
				errChan <- err
				return
			}
			resChan <- struct {
				addr  gethcommon.Address
				value *big.Int
			}{addr, value}
		}(addr)
	}

	go func() {
		wg.Wait()
		close(resChan)
		close(errChan)
	}()

	for res := range resChan {
		mu.Lock()
		resMap[res.addr] = res.value
		mu.Unlock()
	}

	// Check if there were any errors
	if len(errChan) > 0 {
		return nil, <-errChan // Return the first error encountered
	}

	return resMap, nil
}

func toSQLAddress(address gethcommon.Address) string {
	return strings.ToLower(address.Hex()[2:])
}
