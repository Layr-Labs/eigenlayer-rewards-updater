package calculator

import (
	"math/big"
	"testing"
	"time"

	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/stretchr/testify/assert"
)

func TestOperatorSetDataService(t *testing.T) {
	testBlockNumber := big.NewInt(10102668)

	STETH_STRATEGY_ADDRESS := gethcommon.HexToAddress("0xB613E78E2068d7489bb66419fB1cfa11275d14da")

	firstClaimerSetTimestamp := big.NewInt(1706728896)
	secondClaimerSetTimestamp := big.NewInt(1706728956)
	thirdClaimerSetTimestamp := big.NewInt(1706732424)

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

	rpcClient, err := rpc.Dial("https://goerli.infura.io/v3/9aa3d95b3bc440fa88ea12eaa4456161")
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
		_, err := osds.GetBlockNumberAtTimestamp(big.NewInt(time.Now().Unix() - 2*24*60*60))
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("GetBlockNumberAtTimestamp took %s", time.Since(start))

		t.Fail()
	})

	t.Run("test GetClaimersAtTimestamp", func(t *testing.T) {
		createClaimerSetTable()

		claimers, err := osds.GetClaimersAtTimestamp(firstClaimerSetTimestamp, testingAccounts)
		if err != nil {
			t.Fatal(err)
		}

		// only the first test account has a claimer that is different from itself
		for _, account := range testingAccounts {
			if account == testingAccounts[0] {
				assert.Equal(t, testingAccounts[1], claimers[account])
			} else {
				assert.Equal(t, account, claimers[account])
			}
		}

		claimers, err = osds.GetClaimersAtTimestamp(secondClaimerSetTimestamp, testingAccounts)
		if err != nil {
			t.Fatal(err)
		}

		// the first test account has a different claimer from the first set
		for _, account := range testingAccounts {
			if account == testingAccounts[0] {
				assert.Equal(t, testingAccounts[2], claimers[account])
			} else {
				assert.Equal(t, account, claimers[account])
			}
		}

		claimers, err = osds.GetClaimersAtTimestamp(thirdClaimerSetTimestamp, testingAccounts)
		if err != nil {
			t.Fatal(err)
		}

		// the first and fourth test accounts have different claimers from the first set
		for _, account := range testingAccounts {
			if account == testingAccounts[0] {
				assert.Equal(t, testingAccounts[2], claimers[account])
			} else if account == testingAccounts[3] {
				assert.Equal(t, testingAccounts[4], claimers[account])
			} else {
				assert.Equal(t, account, claimers[account])
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

	t.Cleanup(func() {
		conn.ExecSQL(`
			DROP TABLE IF EXISTS sgd34.claimer_set;
		`)
	})

}

func createClaimerSetTable() {
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
			decode('1234', 'hex'),
			decode('27977e6E4426A525d055A587d2a0537b4cb376eA', 'hex'),
			decode('20392d0d40Bdb3Bb2727aA9e34b3A631c8C7bE8F', 'hex'),
			10559231,
			1706728896,
			decode('1234567890123456789012345678901234567890123456789012345678901234', 'hex')
		);
	`)

	conn.ExecSQL(`
		INSERT INTO sgd34.claimer_set VALUES (
			decode('5678', 'hex'),
			decode('27977e6E4426A525d055A587d2a0537b4cb376eA', 'hex'),
			decode('6dF1eB642bF863E3A0547Bf347844BE1725cB678', 'hex'),
			10559232,
			1706728956,
			decode('5678901234567890123456789012345678901234567890123456789012345678', 'hex')
		);
	`)

	conn.ExecSQL(`
		INSERT INTO sgd34.claimer_set VALUES (
			decode('9101', 'hex'),
			decode('bCAc81D98ad3b9cAA48db35d20eDe91D2C59a0e1', 'hex'),
			decode('81dB2Cf17E7E6E3f4AA66D450E647a69E8CB2487', 'hex'),
			10559233,
			1706732424,
			decode('6678901234567890123456789012345678901234567890123456789012345678', 'hex')
		);
	`)
}
