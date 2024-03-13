package calculator

import (
	"context"
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/Layr-Labs/eigenlayer-payment-updater/calculator/utils"
	"github.com/Layr-Labs/eigenlayer-payment-updater/common"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
)

func createOperatorSetDataService() *OperatorSetDataServiceImpl {
	err := godotenv.Load("../.env") // Replace with your file path
	if err != nil {
		panic(err)
	}

	rpcClient, err := rpc.Dial(os.Getenv("RPC_URL"))
	if err != nil {
		panic(err)
	}

	ethClient := ethclient.NewClient(rpcClient)

	osds := NewOperatorSetDataServiceImpl(
		dbpool,
		schemaService,
		ethClient,
	)

	return osds
}

func TestGetBlockNumberAtTimestamp(t *testing.T) {
	t.Run("test GetBlockNumberAtTimestamp", func(t *testing.T) {
		osds := createOperatorSetDataService()

		start := time.Now()
		blockNumber, err := osds.GetBlockNumberAtTimestamp(context.Background(), big.NewInt(1708285982))
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, big.NewInt(10558769), blockNumber)
		log.Info().Msgf("GetBlockNumberAtTimestamp took %s", time.Since(start))
	})
}

func TestGetRecipientsAtTimestamp(t *testing.T) {
	firstRecipientSetTimestamp := big.NewInt(1706728896)
	secondRecipientSetTimestamp := big.NewInt(1706728956)
	thirdRecipientSetTimestamp := big.NewInt(1706732424)

	testingAccounts := []gethcommon.Address{
		gethcommon.HexToAddress("0x27977e6E4426A525d055A587d2a0537b4cb376eA"),
		gethcommon.HexToAddress("0x20392d0d40Bdb3Bb2727aA9e34b3A631c8C7bE8F"),
		gethcommon.HexToAddress("0x6dF1eB642bF863E3A0547Bf347844BE1725cB678"),
		gethcommon.HexToAddress("0xbCAc81D98ad3b9cAA48db35d20eDe91D2C59a0e1"),
		gethcommon.HexToAddress("0x81dB2Cf17E7E6E3f4AA66D450E647a69E8CB2487"),
	}

	t.Run("test GetRecipientsAtTimestamp", func(t *testing.T) {
		createRecipientSetTable()

		osds := createOperatorSetDataService()
		recipients, err := osds.GetRecipientsAtTimestamp(firstRecipientSetTimestamp, testingAccounts)
		if err != nil {
			t.Fatal(err)
		}

		// only the first test account has a recipient that is different from itself
		for _, account := range testingAccounts {
			if account == testingAccounts[0] {
				assert.Equal(t, testingAccounts[1], recipients[account])
			} else {
				assert.Equal(t, account, recipients[account])
			}
		}

		recipients, err = osds.GetRecipientsAtTimestamp(secondRecipientSetTimestamp, testingAccounts)
		if err != nil {
			t.Fatal(err)
		}

		// the first test account has a different recipient from the first set
		for _, account := range testingAccounts {
			if account == testingAccounts[0] {
				assert.Equal(t, testingAccounts[2], recipients[account])
			} else {
				assert.Equal(t, account, recipients[account])
			}
		}

		recipients, err = osds.GetRecipientsAtTimestamp(thirdRecipientSetTimestamp, testingAccounts)
		if err != nil {
			t.Fatal(err)
		}

		// the first and fourth test accounts have different recipients from the first set
		for _, account := range testingAccounts {
			if account == testingAccounts[0] {
				assert.Equal(t, testingAccounts[2], recipients[account])
			} else if account == testingAccounts[3] {
				assert.Equal(t, testingAccounts[4], recipients[account])
			} else {
				assert.Equal(t, account, recipients[account])
			}
		}
	})

	t.Cleanup(func() {
		t.Cleanup(func() {
			conn.ExecSQL(`
				DROP TABLE IF EXISTS sgd34.claimer_set;
			`)
		})
	})
}

func TestGetStakerSetSharesAtTimestamp(t *testing.T) {
	timestamp1 := big.NewInt(1000000)
	timestamp2 := big.NewInt(2000000)

	t.Run("test GetStakerSetSharesAtTimestamp for 1 strategy", func(t *testing.T) {
		createDelegationStrategySharesTable()
		osds := createOperatorSetDataService()

		operator := &common.Operator{
			Address:         gethcommon.HexToAddress("0xb613e78e2068d7489bb66419fb1cfa11275d14da"),
			DelegatedShares: make(map[gethcommon.Address]*big.Int),
		}
		err := osds.GetStakeWeightsAtTimestamp(operator, timestamp1, []gethcommon.Address{utils.TEST_STRATEGY_ADDRESS_1})
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, 1, len(operator.DelegatedShares))
		assert.Equal(t, "80000000", operator.DelegatedShares[utils.TEST_STRATEGY_ADDRESS_1].String())
		assert.Equal(t, 3, len(operator.Stakers))
		assert.Equal(t, "0x5152bee7840E3A6261034e7FeCAf8FfBFf5cB6eE", operator.Stakers[0].Address.String())
		assert.Equal(t, 1, len(operator.Stakers[0].Shares))
		assert.Equal(t, "10000000", operator.Stakers[0].Shares[utils.TEST_STRATEGY_ADDRESS_1].String())
		assert.Equal(t, "0x67185a8067DC178dAFF0571b4835d52bCFE0dE4C", operator.Stakers[1].Address.String())
		assert.Equal(t, 1, len(operator.Stakers[1].Shares))
		assert.Equal(t, "30000000", operator.Stakers[1].Shares[utils.TEST_STRATEGY_ADDRESS_1].String())
		assert.Equal(t, "0xbfc9ca1c434ab19E5F75ACd2d603dc0621ef64E2", operator.Stakers[2].Address.String())
		assert.Equal(t, 1, len(operator.Stakers[2].Shares))
		assert.Equal(t, "40000000", operator.Stakers[2].Shares[utils.TEST_STRATEGY_ADDRESS_1].String())

		operator = &common.Operator{
			Address:         gethcommon.HexToAddress("0xb613e78e2068d7489bb66419fb1cfa11275d14da"),
			DelegatedShares: make(map[gethcommon.Address]*big.Int),
		}

		err = osds.GetStakeWeightsAtTimestamp(operator, timestamp2, []gethcommon.Address{utils.TEST_STRATEGY_ADDRESS_1})
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, 1, len(operator.DelegatedShares))
		assert.Equal(t, "45000000", operator.DelegatedShares[utils.TEST_STRATEGY_ADDRESS_1].String())
		assert.Equal(t, 2, len(operator.Stakers))
		assert.Equal(t, "0x67185a8067DC178dAFF0571b4835d52bCFE0dE4C", operator.Stakers[0].Address.String())
		assert.Equal(t, 1, len(operator.Stakers[0].Shares))
		assert.Equal(t, "5000000", operator.Stakers[0].Shares[utils.TEST_STRATEGY_ADDRESS_1].String())
		assert.Equal(t, "0xbfc9ca1c434ab19E5F75ACd2d603dc0621ef64E2", operator.Stakers[1].Address.String())
		assert.Equal(t, 1, len(operator.Stakers[1].Shares))
		assert.Equal(t, "40000000", operator.Stakers[1].Shares[utils.TEST_STRATEGY_ADDRESS_1].String())
	})

	t.Run("test GetStakerSetSharesAtTimestamp for 2 strategies and multipliers", func(t *testing.T) {
		createDelegationStrategySharesTable()
		osds := createOperatorSetDataService()

		operator := &common.Operator{
			Address:         gethcommon.HexToAddress("0xb613e78e2068d7489bb66419fb1cfa11275d14da"),
			DelegatedShares: make(map[gethcommon.Address]*big.Int),
		}
		err := osds.GetStakeWeightsAtTimestamp(
			operator,
			timestamp1,
			[]gethcommon.Address{utils.TEST_STRATEGY_ADDRESS_1, utils.TEST_STRATEGY_ADDRESS_2},
		)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, 2, len(operator.DelegatedShares))
		assert.Equal(t, "80000000", operator.DelegatedShares[utils.TEST_STRATEGY_ADDRESS_1].String())
		assert.Equal(t, "30000000", operator.DelegatedShares[utils.TEST_STRATEGY_ADDRESS_2].String())
		assert.Equal(t, 3, len(operator.Stakers))
		assert.Equal(t, "0x5152bee7840E3A6261034e7FeCAf8FfBFf5cB6eE", operator.Stakers[0].Address.String())
		assert.Equal(t, 2, len(operator.Stakers[0].Shares))
		assert.Equal(t, "10000000", operator.Stakers[0].Shares[utils.TEST_STRATEGY_ADDRESS_1].String())
		assert.Equal(t, "20000000", operator.Stakers[0].Shares[utils.TEST_STRATEGY_ADDRESS_2].String())
		assert.Equal(t, "0x67185a8067DC178dAFF0571b4835d52bCFE0dE4C", operator.Stakers[1].Address.String())
		assert.Equal(t, 1, len(operator.Stakers[1].Shares))
		assert.Equal(t, "30000000", operator.Stakers[1].Shares[utils.TEST_STRATEGY_ADDRESS_1].String())
		assert.Equal(t, "0xbfc9ca1c434ab19E5F75ACd2d603dc0621ef64E2", operator.Stakers[2].Address.String())
		assert.Equal(t, 2, len(operator.Stakers[2].Shares))
		assert.Equal(t, "40000000", operator.Stakers[2].Shares[utils.TEST_STRATEGY_ADDRESS_1].String())
		assert.Equal(t, "10000000", operator.Stakers[2].Shares[utils.TEST_STRATEGY_ADDRESS_2].String())

		operator = &common.Operator{
			Address:         gethcommon.HexToAddress("0xb613e78e2068d7489bb66419fb1cfa11275d14da"),
			DelegatedShares: make(map[gethcommon.Address]*big.Int),
		}

		err = osds.GetStakeWeightsAtTimestamp(
			operator,
			big.NewInt(2000000),
			[]gethcommon.Address{utils.TEST_STRATEGY_ADDRESS_1, utils.TEST_STRATEGY_ADDRESS_2},
		)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, 2, len(operator.DelegatedShares))
		assert.Equal(t, "45000000", operator.DelegatedShares[utils.TEST_STRATEGY_ADDRESS_1].String())
		assert.Equal(t, "10000000", operator.DelegatedShares[utils.TEST_STRATEGY_ADDRESS_2].String())
		assert.Equal(t, 2, len(operator.Stakers))
		assert.Equal(t, "0x67185a8067DC178dAFF0571b4835d52bCFE0dE4C", operator.Stakers[0].Address.String())
		assert.Equal(t, 1, len(operator.Stakers[0].Shares))
		assert.Equal(t, "5000000", operator.Stakers[0].Shares[utils.TEST_STRATEGY_ADDRESS_1].String())
		assert.Equal(t, "0xbfc9ca1c434ab19E5F75ACd2d603dc0621ef64E2", operator.Stakers[1].Address.String())
		assert.Equal(t, 2, len(operator.Stakers[1].Shares))
		assert.Equal(t, "40000000", operator.Stakers[1].Shares[utils.TEST_STRATEGY_ADDRESS_1].String())
		assert.Equal(t, "10000000", operator.Stakers[1].Shares[utils.TEST_STRATEGY_ADDRESS_2].String())
	})
}

func createRecipientSetTable() {
	conn.ExecSQL(`
		CREATE TABLE IF NOT EXISTS sgd34.claimer_set (
			id bytea PRIMARY KEY,
			account bytea NOT NULL,
			claimer bytea NOT NULL,
			block_number numeric NOT NULL,
			block_timestamp numeric NOT NULL,
			transaction_hash bytea NOT NULL
		);
	`)

	// insert a couple rows
	conn.ExecSQL(`
		INSERT INTO sgd34.claimer_set VALUES (
			decode('abcd', 'hex'),
			decode('27977e6E4426A525d055A587d2a0537b4cb376eA', 'hex'),
			decode('20392d0d40Bdb3Bb2727aA9e34b3A631c8C7bE8F', 'hex'),
			10559231,
			1706728896,
			decode('1234567890123456789012345678901234567890123456789012345678901234', 'hex')
		);
	`)

	conn.ExecSQL(`
		INSERT INTO sgd34.claimer_set VALUES (
			decode('efab', 'hex'),
			decode('27977e6E4426A525d055A587d2a0537b4cb376eA', 'hex'),
			decode('6dF1eB642bF863E3A0547Bf347844BE1725cB678', 'hex'),
			10559232,
			1706728956,
			decode('5678901234567890123456789012345678901234567890123456789012345678', 'hex')
		);
	`)

	conn.ExecSQL(`
		INSERT INTO sgd34.claimer_set VALUES (
			decode('cdef', 'hex'),
			decode('bCAc81D98ad3b9cAA48db35d20eDe91D2C59a0e1', 'hex'),
			decode('81dB2Cf17E7E6E3f4AA66D450E647a69E8CB2487', 'hex'),
			10559233,
			1706732424,
			decode('6678901234567890123456789012345678901234567890123456789012345678', 'hex')
		);
	`)
}

func createDelegationStrategySharesTable() {
	conn.ExecSQL(`
			DROP TABLE IF EXISTS sgd34.staker_delegation_share;
		`)

	conn.ExecSQL(`
		CREATE TABLE IF NOT EXISTS sgd34.staker_delegation_share (
			id bytea PRIMARY KEY,
			staker bytea NOT NULL,
			operator bytea NOT NULL,
			strategy bytea NOT NULL,
			shares numeric NOT NULL,
			update_block_timestamp numeric NOT NULL
		);
	`)

	// insert a couple rows
	conn.ExecSQL(`
		INSERT INTO sgd34.staker_delegation_share VALUES (
			decode('1234', 'hex'),
			decode('5152bee7840E3A6261034e7FeCAf8FfBFf5cB6eE', 'hex'),
			decode('b613e78e2068d7489bb66419fb1cfa11275d14da', 'hex'),
			decode('1234567890987654321234567890987654321234', 'hex'),
			10000000,
			1000000
		);
	`)

	conn.ExecSQL(`
		INSERT INTO sgd34.staker_delegation_share VALUES (
			decode('1235', 'hex'),
			decode('5152bee7840E3A6261034e7FeCAf8FfBFf5cB6eE', 'hex'),
			decode('b613e78e2068d7489bb66419fb1cfa11275d14da', 'hex'),
			decode('0987654321234567890987654321234567890987', 'hex'),
			20000000,
			1000000
		);
	`)

	conn.ExecSQL(`
		INSERT INTO sgd34.staker_delegation_share VALUES (
			decode('5678', 'hex'),
			decode('67185a8067DC178dAFF0571b4835d52bCFE0dE4C', 'hex'),
			decode('b613e78e2068d7489bb66419fb1cfa11275d14da', 'hex'),
			decode('1234567890987654321234567890987654321234', 'hex'),
			30000000,
			1000000
		);
	`)

	conn.ExecSQL(`
		INSERT INTO sgd34.staker_delegation_share VALUES (
			decode('9101', 'hex'),
			decode('bfc9ca1c434ab19E5F75ACd2d603dc0621ef64E2', 'hex'),
			decode('b613e78e2068d7489bb66419fb1cfa11275d14da', 'hex'),
			decode('1234567890987654321234567890987654321234', 'hex'),
			40000000,
			1000000
		);
	`)

	conn.ExecSQL(`
		INSERT INTO sgd34.staker_delegation_share VALUES (
			decode('9102', 'hex'),
			decode('bfc9ca1c434ab19E5F75ACd2d603dc0621ef64E2', 'hex'),
			decode('b613e78e2068d7489bb66419fb1cfa11275d14da', 'hex'),
			decode('0987654321234567890987654321234567890987', 'hex'),
			10000000,
			1000000
		);
	`)

	conn.ExecSQL(`
		INSERT INTO sgd34.staker_delegation_share VALUES (
			decode('1236', 'hex'),
			decode('5152bee7840E3A6261034e7FeCAf8FfBFf5cB6eE', 'hex'),
			decode('0000000000000000000000000000000000000000', 'hex'),
			decode('1234567890987654321234567890987654321234', 'hex'),
			0,
			2000000
		);
	`)

	conn.ExecSQL(`
		INSERT INTO sgd34.staker_delegation_share VALUES (
			decode('1237', 'hex'),
			decode('5152bee7840E3A6261034e7FeCAf8FfBFf5cB6eE', 'hex'),
			decode('0000000000000000000000000000000000000000', 'hex'),
			decode('0987654321234567890987654321234567890987', 'hex'),
			0,
			2000000
		);
	`)

	conn.ExecSQL(`
		INSERT INTO sgd34.staker_delegation_share VALUES (
			decode('5679', 'hex'),
			decode('67185a8067DC178dAFF0571b4835d52bCFE0dE4C', 'hex'),
			decode('b613e78e2068d7489bb66419fb1cfa11275d14da', 'hex'),
			decode('1234567890987654321234567890987654321234', 'hex'),
			5000000,
			2000000
		);
	`)

}
