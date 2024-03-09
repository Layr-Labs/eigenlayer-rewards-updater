package calculator

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"
	"time"

	contractIClaimingManager "github.com/Layr-Labs/eigenlayer-payment-updater/bindings/IClaimingManager"
	contractIEigenPodManager "github.com/Layr-Labs/eigenlayer-payment-updater/bindings/IEigenPodManager"
	contractIStrategyManager "github.com/Layr-Labs/eigenlayer-payment-updater/bindings/IStrategyManager"
	contractMulticall3 "github.com/Layr-Labs/eigenlayer-payment-updater/bindings/Multicall3"
	"github.com/Layr-Labs/eigenlayer-payment-updater/common"
	"github.com/Layr-Labs/eigenlayer-payment-updater/common/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
)

var SECONDS_PER_BLOCK_ESTIMATE = big.NewInt(12)

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

// claiming manager address
var CLAIMING_MANAGER_ADDRESS = gethcommon.HexToAddress("0xBF81C737bc6871f1Dfa143f0eb416C34Cb22f47d")

type OperatorSetDataService interface {
	// GetOperatorSetForStrategyAtTimestamp returns the operator set for a given strategy at a given timestamps
	GetOperatorSetForStrategyAtTimestamp(ctx context.Context, timestamp *big.Int, avs gethcommon.Address, strategy gethcommon.Address) (*common.OperatorSet, error)
}

type OperatorSetDataServiceImpl struct {
	dbpool        *pgxpool.Pool
	schemaService *common.SubgraphSchemaService
	ethClient     *ethclient.Client
}

func NewOperatorSetDataService(
	dbpool *pgxpool.Pool,
	schemaService *common.SubgraphSchemaService,
	ethClient *ethclient.Client,
) OperatorSetDataService {
	return &OperatorSetDataServiceImpl{
		dbpool:        dbpool,
		schemaService: schemaService,
		ethClient:     ethClient,
	}
}

func NewOperatorSetDataServiceImpl(
	dbpool *pgxpool.Pool,
	schemaService *common.SubgraphSchemaService,
	ethClient *ethclient.Client,
) *OperatorSetDataServiceImpl {
	return &OperatorSetDataServiceImpl{
		dbpool:        dbpool,
		schemaService: schemaService,
		ethClient:     ethClient,
	}
}

func (s *OperatorSetDataServiceImpl) GetOperatorSetForStrategyAtTimestamp(ctx context.Context, timestamp *big.Int, avs gethcommon.Address, strategy gethcommon.Address) (*common.OperatorSet, error) {
	log.Info().Msgf("getting operator set for avs %s for strategy %s at timestamp %d", avs.Hex(), strategy.Hex(), timestamp)

	operatorSet := &common.OperatorSet{}
	operatorSet.TotalStakedStrategyShares = big.NewInt(0)

	start := time.Now()

	// get the blocknumber of the block at the given timestamp
	blockNumber, err := s.GetBlockNumberAtTimestamp(ctx, timestamp)
	if err != nil {
		return nil, err
	}

	log.Info().Msgf("got block number of %d in %s", blockNumber, time.Since(start))
	start = time.Now()

	// get the global commission at the given block number
	globalCommission, err := s.GetGlobalCommissionAtBlockNumber(blockNumber)
	if err != nil {
		return nil, err
	}

	log.Info().Msgf("got global commission in %s", time.Since(start))
	start = time.Now()

	// get all operators for the given strategy at the given timestamp
	operatorAddresses, err := s.GetOperatorAddressesForAVSAtTimestamp(timestamp, avs, strategy)
	if err != nil {
		return nil, err
	}

	log.Info().Msgf("found %d operators in %s", len(operatorAddresses), time.Since(start))

	operatorSet.Operators = make([]*common.Operator, len(operatorAddresses))

	// loop thru each operator and get their staker sets
	for i, operatorAddress := range operatorAddresses {
		operatorSet.Operators[i] = &common.Operator{}
		operatorSet.Operators[i].Address = operatorAddress
		operatorSet.Operators[i].Commission = globalCommission
		operatorSet.Operators[i].TotalDelegatedStrategyShares = big.NewInt(0)

		start = time.Now()

		// get the stakers of the operator
		stakers, err := s.GetStakersDelegatedToOperatorAtTimestamp(timestamp, operatorAddress)
		if err != nil {
			return nil, err
		}

		log.Info().Msgf("found %d stakers for operator %s in %s", len(stakers), operatorAddress.Hex(), time.Since(start))
		start = time.Now()

		// get the recipients of each staker and the operator
		recipients, err := s.GetRecipientsAtTimestamp(timestamp, append(stakers, operatorAddress))
		if err != nil {
			return nil, err
		}

		log.Info().Msgf("got recipients of %d stakers in %s", len(stakers), time.Since(start))
		start = time.Now()

		// get the shares of each staker
		strategyShareMap, err := s.GetSharesOfStakersAtBlockNumber(blockNumber, strategy, stakers)
		if err != nil {
			return nil, err
		}

		log.Info().Msgf("got shares of %d stakers in %s", len(stakers), time.Since(start))

		operatorSet.Operators[i].Stakers = make([]*common.Staker, len(stakers))

		// loop thru each staker and get their shares
		for j, stakerAddress := range stakers {
			operatorSet.Operators[i].Stakers[j] = &common.Staker{}
			operatorSet.Operators[i].Stakers[j].Address = stakerAddress
			operatorSet.Operators[i].Stakers[j].Recipient = recipients[stakerAddress]
			operatorSet.Operators[i].Stakers[j].StrategyShares = strategyShareMap[stakerAddress]

			operatorSet.Operators[i].Recipient = recipients[operatorAddress]
			// add the staker's shares to the operator's total delegated strategy shares
			operatorSet.Operators[i].TotalDelegatedStrategyShares = operatorSet.Operators[i].TotalDelegatedStrategyShares.Add(operatorSet.Operators[i].TotalDelegatedStrategyShares, strategyShareMap[stakerAddress])
		}

		// add the operator's total delegated strategy shares to the operator set's total staked strategy shares
		operatorSet.TotalStakedStrategyShares = operatorSet.TotalStakedStrategyShares.Add(operatorSet.TotalStakedStrategyShares, operatorSet.Operators[i].TotalDelegatedStrategyShares)
	}

	return operatorSet, nil
}

func (s *OperatorSetDataServiceImpl) GetOperatorAddressesForAVSAtTimestamp(timestamp *big.Int, avs gethcommon.Address, strategy gethcommon.Address) ([]gethcommon.Address, error) {
	// TODO: actually query the database

	operatorAddresses := []gethcommon.Address{
		gethcommon.HexToAddress("0x9631af6a712d296fedb800132edcf08e493b12cd"),
		// gethcommon.HexToAddress("0x96091e6a492b17b389d5ae9e0662049890446568"),
		// gethcommon.HexToAddress("0x63a27fd29a5385561991108e0cefb288c627cc03"),
	}

	return operatorAddresses, nil
}

func (s *OperatorSetDataServiceImpl) GetGlobalCommissionAtBlockNumber(blockNumber *big.Int) (*big.Int, error) {
	claimingManagerContract, err := contractIClaimingManager.NewContractIClaimingManager(CLAIMING_MANAGER_ADDRESS, s.ethClient)
	if err != nil {
		return nil, err
	}

	globalCommission, err := claimingManagerContract.GlobalCommissionBips(&bind.CallOpts{BlockNumber: blockNumber})
	if err != nil {
		return nil, err
	}

	log.Info().Msgf("got global commission of %d at block number %d", globalCommission, blockNumber)

	return big.NewInt(int64(globalCommission)), nil
}

func (s *OperatorSetDataServiceImpl) GetRecipientsAtTimestamp(timestamp *big.Int, accounts []gethcommon.Address) (map[gethcommon.Address]gethcommon.Address, error) {
	// get the schema id for the claiming manager subgraph
	schemaID, err := s.schemaService.GetSubgraphSchema(context.Background(), utils.SUBGRAPH_CLAIMING_MANAGER)
	if err != nil {
		return nil, err
	}

	// format the query
	formattedQuery := fmt.Sprintf(recipientsAtTimestampQuery, schemaID, toSQLAddreses(accounts))

	// query the database
	rows, err := s.dbpool.Query(context.Background(), formattedQuery, timestamp)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// create a map of account to recipient
	recipients := make(map[gethcommon.Address]gethcommon.Address)
	for rows.Next() {
		var (
			accountBytes   []byte
			recipientBytes []byte
		)

		err := rows.Scan(
			&accountBytes,
			&recipientBytes,
		)
		if err != nil {
			return nil, err
		}

		account := gethcommon.HexToAddress(hex.EncodeToString(accountBytes))
		recipient := gethcommon.HexToAddress(hex.EncodeToString(recipientBytes))

		recipients[account] = recipient
	}

	// set the recipient of any account that doesn't have one to the account itself
	for _, account := range accounts {
		if _, ok := recipients[account]; !ok {
			recipients[account] = account
		}
	}

	return recipients, nil
}

func (s *OperatorSetDataServiceImpl) GetStakersDelegatedToOperatorAtTimestamp(timestamp *big.Int, operator gethcommon.Address) ([]gethcommon.Address, error) {
	// get the schema id for the claiming manager subgraph
	schemaID, err := s.schemaService.GetSubgraphSchema(context.Background(), utils.SUBGRAPH_DELEGATION_MANAGER)
	if err != nil {
		return nil, err
	}

	// format the query
	formattedQuery := fmt.Sprintf(stakerSetAtTimestampQuery, schemaID, schemaID)

	// query the database
	rows, err := s.dbpool.Query(context.Background(), formattedQuery, toSQLAddress(operator), timestamp)
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

func (s *OperatorSetDataServiceImpl) GetBlockNumberAtTimestamp(ctx context.Context, timestamp *big.Int) (*big.Int, error) {
	head, err := s.ethClient.HeaderByNumber(ctx, nil)
	if err != nil {
		return nil, err
	}

	headTimestamp := big.NewInt(int64(head.Time))
	if headTimestamp.Cmp(timestamp) < 0 {
		return nil, fmt.Errorf("timestamp %d is in the future", timestamp)
	}

	lowerBound, err := s.ethClient.HeaderByNumber(ctx, new(big.Int).Sub(head.Number, mul(div(headTimestamp.Sub(headTimestamp, timestamp), SECONDS_PER_BLOCK_ESTIMATE), big.NewInt(2))))
	if err != nil {
		return nil, err
	}

	// decrease the lower bound until it is less than the timestamp
	for lowerBound.Time > timestamp.Uint64() {
		lowerBound, err = s.ethClient.HeaderByNumber(ctx, lowerBound.Number.Sub(lowerBound.Number, big.NewInt(10)))
		if err != nil {
			return nil, err
		}
	}

	var lo, hi = lowerBound.Number, head.Number
	for lo.Cmp(hi) < 0 {
		mid := new(big.Int).Add(lo, hi)
		mid.Div(mid, big.NewInt(2))

		header, err := s.ethClient.HeaderByNumber(ctx, mid)
		if err != nil {
			return nil, err
		}
		if header.Time < timestamp.Uint64() {
			lo = mid.Add(mid, big.NewInt(1))
		} else {
			hi = mid
		}
	}

	return lo.Sub(lo, big.NewInt(1)), nil
}

func (s *OperatorSetDataServiceImpl) GetSharesOfStakersAtBlockNumber(blockNumber *big.Int, strategy gethcommon.Address, stakers []gethcommon.Address) (map[gethcommon.Address]*big.Int, error) {
	stakerToStrategyShares := make(map[gethcommon.Address]*big.Int)

	// get the mutlicalls for shares
	var getShareCall func(staker, strategy gethcommon.Address) (contractMulticall3.Multicall3Call, error)
	if strategy == BEACON_CHAIN_ETH_STRATEGY_ADDRESS {
		eigenPodManagerAbi, _ := contractIEigenPodManager.ContractIEigenPodManagerMetaData.GetAbi()
		getShareCall = func(staker, strategy gethcommon.Address) (contractMulticall3.Multicall3Call, error) {
			sharesCall, err := eigenPodManagerAbi.Pack("podOwnerShares", staker)
			if err != nil {
				return contractMulticall3.Multicall3Call{}, err
			}

			return contractMulticall3.Multicall3Call{
				Target:   EIGEN_POD_MANAGER_ADDRESS,
				CallData: sharesCall,
			}, nil
		}
	} else {
		strategyManagerAbi, _ := contractIStrategyManager.ContractIStrategyManagerMetaData.GetAbi()
		getShareCall = func(staker, strategy gethcommon.Address) (contractMulticall3.Multicall3Call, error) {
			sharesCall, err := strategyManagerAbi.Pack("stakerStrategyShares", staker, strategy)
			if err != nil {
				return contractMulticall3.Multicall3Call{}, err
			}

			return contractMulticall3.Multicall3Call{
				Target:   STRATEGY_MANAGER_ADDRESS,
				CallData: sharesCall,
			}, nil
		}
	}

	// get the shares of the stakers in batches
	var batchSize = 300
	for i := 0; i < len(stakers); i += batchSize {
		end := i + batchSize
		if end > len(stakers) {
			end = len(stakers)
		}

		// todo: parallelize this
		err := s.GetStrategySharesForStakerBatchAtABlockNumber(stakerToStrategyShares, getShareCall, blockNumber, strategy, stakers[i:end])
		if err != nil {
			return nil, err
		}
	}

	return stakerToStrategyShares, nil
}

func (s *OperatorSetDataServiceImpl) GetStrategySharesForStakerBatchAtABlockNumber(
	stakerToStrategyShares map[gethcommon.Address]*big.Int,
	getShareCall func(staker, strategy gethcommon.Address) (contractMulticall3.Multicall3Call, error),
	blockNumber *big.Int,
	strategy gethcommon.Address,
	stakers []gethcommon.Address,
) error {
	calls := make([]contractMulticall3.Multicall3Call, 0)
	for _, staker := range stakers {
		shareCall, err := getShareCall(staker, strategy)
		if err != nil {
			return err
		}

		calls = append(calls, shareCall)
	}

	results, err := s.aggregateMulticall(blockNumber, calls)
	if err != nil {
		return err
	}

	for i, shareBytes := range results {
		staker := stakers[i]
		stakerToStrategyShares[staker] = new(big.Int).SetBytes(shareBytes)
	}

	return nil
}

func (s *OperatorSetDataServiceImpl) aggregateMulticall(blockNumber *big.Int, calls []contractMulticall3.Multicall3Call) ([][]byte, error) {
	multicallAbi, err := contractMulticall3.ContractMulticall3MetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	multicallContract := bind.NewBoundContract(MULTICALL3_ADDRESS, *multicallAbi, s.ethClient, s.ethClient, s.ethClient)

	var res []interface{}
	err = multicallContract.Call(&bind.CallOpts{BlockNumber: blockNumber}, &res, "aggregate", calls)
	if err != nil {
		return nil, err
	}

	return res[1].([][]byte), nil
}

func toSQLAddreses(addresses []gethcommon.Address) string {
	var sqlAddresses string
	for _, address := range addresses {
		sqlAddresses += fmt.Sprintf("'%s',", toSQLAddress(address))
	}
	return strings.TrimRight(sqlAddresses, ",")
}

func toSQLAddress(address gethcommon.Address) string {
	return strings.ToLower(address.Hex()[2:])
}
