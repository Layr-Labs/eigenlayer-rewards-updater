package calculator

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"
	"time"

	contractIClaimingManager "github.com/Layr-Labs/eigenlayer-payment-updater/bindings/IClaimingManager"
	"github.com/Layr-Labs/eigenlayer-payment-updater/common"
	"github.com/Layr-Labs/eigenlayer-payment-updater/common/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
	"github.com/shopspring/decimal"
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
	// GetWeightedOperatorSetAtTimestamp returns the operator set and their weights according to strategies at the given timestamp
	GetWeightedOperatorSetAtTimestamp(ctx context.Context, timestamp *big.Int, avs gethcommon.Address, strategies []gethcommon.Address) (*common.OperatorSet, error)
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

func (s *OperatorSetDataServiceImpl) GetWeightedOperatorSetAtTimestamp(ctx context.Context, timestamp *big.Int, avs gethcommon.Address, strategies []gethcommon.Address) (*common.OperatorSet, error) {
	log.Info().Msgf("getting operator set for avs %s for strategies %s at timestamp %d", avs.Hex(), strategies[:], timestamp)

	operatorSet := &common.OperatorSet{}
	operatorSet.TotalStakedShares = make(map[gethcommon.Address]*big.Int)

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
	operatorAddresses, err := s.GetOperatorAddressesForAVSAtTimestamp(timestamp, avs)
	if err != nil {
		return nil, err
	}

	log.Info().Msgf("found %d operators in %s", len(operatorAddresses), time.Since(start))

	operatorSet.Operators = make([]*common.Operator, len(operatorAddresses))

	// todo: parallelize this
	// loop thru each operator and get their staker sets
	for i, operatorAddress := range operatorAddresses {
		operator := &common.Operator{}
		operator.Address = operatorAddress
		operator.Commission = globalCommission
		operator.DelegatedShares = make(map[gethcommon.Address]*big.Int)

		start = time.Now()

		// get the stakers of the operator
		err := s.GetStakeWeightsAtTimestamp(operator, timestamp, strategies)
		if err != nil {
			return nil, err
		}

		log.Info().Msgf("found %d stakers for operator %s in %s", len(operator.Stakers), operatorAddress.Hex(), time.Since(start))
		start = time.Now()

		addresses := make([]gethcommon.Address, len(operator.Stakers)+1)
		for i, staker := range operator.Stakers {
			addresses[i] = staker.Address
		}
		addresses[len(addresses)-1] = operatorAddress

		// get the recipients of each staker and the operator
		recipients, err := s.GetRecipientsAtTimestamp(timestamp, addresses)
		if err != nil {
			return nil, err
		}

		log.Info().Msgf("got recipients of %d stakers in %s", len(addresses), time.Since(start))

		// loop thru each staker and set their recipient
		for j, staker := range operator.Stakers {
			operator.Stakers[j].Recipient = recipients[staker.Address]
		}
		operator.Recipient = recipients[operatorAddress]

		// add the operator to the operator set
		operatorSet.Operators[i] = operator

		// add the operator's total delegated strategy shares to the operator set's total staked strategy shares
		for strategy, shares := range operator.DelegatedShares {
			operatorSet.TotalStakedShares[strategy].Add(operatorSet.TotalStakedShares[strategy], shares)
		}
	}

	return operatorSet, nil
}

func (s *OperatorSetDataServiceImpl) GetOperatorAddressesForAVSAtTimestamp(timestamp *big.Int, avs gethcommon.Address) ([]gethcommon.Address, error) {
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

func (s *OperatorSetDataServiceImpl) GetStakeWeightsAtTimestamp(operator *common.Operator, timestamp *big.Int, strategies []gethcommon.Address) error {
	// get the schema id for the claiming manager subgraph
	schemaID, err := s.schemaService.GetSubgraphSchema(context.Background(), utils.SUBGRAPH_DELEGATION_SHARE_TRACKER)
	if err != nil {
		return err
	}

	// format the query
	formattedQuery := fmt.Sprintf(stakerSetSharesAtTimestampQuery, schemaID, timestamp, toSQLAddreses(strategies), toSQLAddress(operator.Address))

	log.Info().Msgf("executing query: %s", formattedQuery)

	// query the database
	rows, err := s.dbpool.Query(context.Background(), formattedQuery)
	if err != nil {
		return err
	}
	defer rows.Close()

	prevStaker := &common.Staker{}

	// create a list of stakers
	for rows.Next() {
		var (
			stakerBytes   []byte
			strategyBytes []byte
			sharesDecimal decimal.Decimal
		)

		err := rows.Scan(
			&stakerBytes,
			&strategyBytes,
			&sharesDecimal,
		)
		if err != nil {
			return err
		}

		stakerAddress := gethcommon.HexToAddress(hex.EncodeToString(stakerBytes))
		strategyAddress := gethcommon.HexToAddress(hex.EncodeToString(strategyBytes))
		shares := sharesDecimal.BigInt()

		log.Info().Msgf("staker %s has %d shares of strategy %s ", stakerAddress.Hex(), shares, strategyAddress.Hex())

		if stakerAddress.Cmp(prevStaker.Address) == 0 {
			prevStaker.Shares[strategyAddress] = shares
		} else {
			if prevStaker.Address.Cmp(gethcommon.Address{}) != 0 {
				operator.Stakers = append(operator.Stakers, prevStaker)
			}
			prevStaker = &common.Staker{
				Address: stakerAddress,
				Shares:  make(map[gethcommon.Address]*big.Int),
			}
			prevStaker.Shares[strategyAddress] = shares
		}

		if _, found := operator.DelegatedShares[strategyAddress]; !found {
			operator.DelegatedShares[strategyAddress] = big.NewInt(0)
		}
		// add the staker's shares to the operator's total delegated strategy shares
		operator.DelegatedShares[strategyAddress].Add(operator.DelegatedShares[strategyAddress], shares)
	}

	if prevStaker.Address.Cmp(gethcommon.Address{}) != 0 {
		operator.Stakers = append(operator.Stakers, prevStaker)
	}

	return nil
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
