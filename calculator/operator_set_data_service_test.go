package calculator

import (
	"context"
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/Layr-Labs/eigenlayer-payment-updater/common"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
)

func TestOperatorSetDataService(t *testing.T) {
	err := godotenv.Load("../.env") // Replace with your file path
	if err != nil {
		t.Fatal("Error loading .env file", err)
	}

	testBlockNumber := big.NewInt(10102668)

	STETH_STRATEGY_ADDRESS := gethcommon.HexToAddress("0xB613E78E2068d7489bb66419fB1cfa11275d14da")

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

	stETHStakers := []gethcommon.Address{
		gethcommon.HexToAddress("0x5152bee7840E3A6261034e7FeCAf8FfBFf5cB6eE"),
		gethcommon.HexToAddress("0x67185a8067DC178dAFF0571b4835d52bCFE0dE4C"),
		gethcommon.HexToAddress("0xbfc9ca1c434ab19E5F75ACd2d603dc0621ef64E2"),
	}

	beaconChainETHStakers := []gethcommon.Address{
		gethcommon.HexToAddress("0x687b2dF0a4fFFD019CD02F8a1c873f848d4A1c54"),
		gethcommon.HexToAddress("0xe82Bcae45bC947620274a576f3A5F96Cf425e01c"),
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

	t.Run("test GetBlockNumberAtTimestamp", func(t *testing.T) {
		start := time.Now()
		blockNumber, err := osds.GetBlockNumberAtTimestamp(context.Background(), big.NewInt(1708285982))
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, big.NewInt(10558769), blockNumber)
		log.Info().Msgf("GetBlockNumberAtTimestamp took %s", time.Since(start))
	})

	t.Run("test GetRecipientsAtTimestamp", func(t *testing.T) {
		createRecipientSetTable()

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

	// add back when someone gets mad
	// t.Run("test GetStakersDelegatedToOperatorAtTimestamp", func(t *testing.T) {
	// 	stakersDelegatedToOperator, err := osds.GetStakersDelegatedToOperatorAtTimestamp(testTimestamp, P2P_OPERATOR_ADDRESS)
	// 	if err != nil {
	// 		t.Fatal(err)
	// 	}

	// 	assert.Equal(t, 5671, len(stakersDelegatedToOperator))
	// })

	t.Run("test GetSharesOfStakersAtBlockNumber for steth", func(t *testing.T) {
		strategyShares, err := osds.GetSharesOfStakersAtBlockNumber(testBlockNumber, STETH_STRATEGY_ADDRESS, stETHStakers)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, 3, len(strategyShares))
		assert.Equal(t, "0", strategyShares[stETHStakers[0]].String())
		assert.Equal(t, "43069823021157214260", strategyShares[stETHStakers[1]].String())
		assert.Equal(t, "151291540795049465", strategyShares[stETHStakers[2]].String())
	})

	t.Run("test GetSharesOfStakersAtBlockNumber for beacon chain eth", func(t *testing.T) {
		strategyShares, err := osds.GetSharesOfStakersAtBlockNumber(testBlockNumber, BEACON_CHAIN_ETH_STRATEGY_ADDRESS, beaconChainETHStakers)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, 2, len(strategyShares))
		assert.Equal(t, "32000000000000000000", strategyShares[beaconChainETHStakers[0]].String())
		assert.Equal(t, "0", strategyShares[beaconChainETHStakers[1]].String())
	})

	t.Run("test GetSharesOfStakersAtBlockNumber for steth with 1400 stakers", func(t *testing.T) {
		stakers := []gethcommon.Address{}
		for i := 0; i < 901; i++ {
			stakers = append(stakers, common.GetRandomAddress())
		}

		strategyShares, err := osds.GetSharesOfStakersAtBlockNumber(testBlockNumber, STETH_STRATEGY_ADDRESS, stakers)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, 901, len(strategyShares))
	})

	t.Cleanup(func() {
		conn.ExecSQL(`
			DROP TABLE IF EXISTS sgd34.recipient_set;
		`)
	})

}

func createRecipientSetTable() {
	conn.ExecSQL(`
		CREATE TABLE IF NOT EXISTS sgd34.recipient_set (
			id bytea PRIMARY KEY,
			account bytea NOT NULL,
			recipient bytea NOT NULL,
			block_number numeric NOT NULL,
			block_timestamp numeric NOT NULL,
			transaction_hash bytea NOT NULL
		);
	`)

	// insert a couple rows
	conn.ExecSQL(`
		INSERT INTO sgd34.recipient_set VALUES (
			decode('1234', 'hex'),
			decode('27977e6E4426A525d055A587d2a0537b4cb376eA', 'hex'),
			decode('20392d0d40Bdb3Bb2727aA9e34b3A631c8C7bE8F', 'hex'),
			10559231,
			1706728896,
			decode('1234567890123456789012345678901234567890123456789012345678901234', 'hex')
		);
	`)

	conn.ExecSQL(`
		INSERT INTO sgd34.recipient_set VALUES (
			decode('5678', 'hex'),
			decode('27977e6E4426A525d055A587d2a0537b4cb376eA', 'hex'),
			decode('6dF1eB642bF863E3A0547Bf347844BE1725cB678', 'hex'),
			10559232,
			1706728956,
			decode('5678901234567890123456789012345678901234567890123456789012345678', 'hex')
		);
	`)

	conn.ExecSQL(`
		INSERT INTO sgd34.recipient_set VALUES (
			decode('9101', 'hex'),
			decode('bCAc81D98ad3b9cAA48db35d20eDe91D2C59a0e1', 'hex'),
			decode('81dB2Cf17E7E6E3f4AA66D450E647a69E8CB2487', 'hex'),
			10559233,
			1706732424,
			decode('6678901234567890123456789012345678901234567890123456789012345678', 'hex')
		);
	`)
}
