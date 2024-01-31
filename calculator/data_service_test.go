package calculator

import (
	"context"
	"math/big"
	"os"
	"testing"

	"github.com/Layr-Labs/eigenlayer-payment-updater/common"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/joho/godotenv"
)

func TestPaymentCalculatorDataService(t *testing.T) {
	err := godotenv.Load("../.env") // Replace with your file path
	if err != nil {
		t.Fatal("Error loading .env file", err)
	}

	testBlockNumber := big.NewInt(10102668)

	STETH_STRATEGY_ADDRESS := gethcommon.HexToAddress("0xB613E78E2068d7489bb66419fB1cfa11275d14da")

	stakers := []gethcommon.Address{
		gethcommon.HexToAddress("0xb1bd1266ec811048161424f534e74c76c48e6ce2"), gethcommon.HexToAddress("0xc473d412dc52e349862209924c8981b2ee420768"), gethcommon.HexToAddress("0x5152bee7840e3a6261034e7fecaf8ffbff5cb6ee"), gethcommon.HexToAddress("0x67185a8067dc178daff0571b4835d52bcfe0de4c"), gethcommon.HexToAddress("0xbfc9ca1c434ab19e5f75acd2d603dc0621ef64e2"), gethcommon.HexToAddress("0x6095577a3504f6622cc6aab8d550f86d3f348140"), gethcommon.HexToAddress("0x2bff1268c6d9228a417c03296fb40458b5bb4223"), gethcommon.HexToAddress("0xd1c2c0b70afb8f3c207642e7c9ee3c36375a72b1"), gethcommon.HexToAddress("0xb59f4da97be6ea64cd176d48f9142f90730a7b86"), gethcommon.HexToAddress("0x1b09ccb3c1bce8d5316f9fb1d77816561062cbe7"), gethcommon.HexToAddress("0x2515926e3a14ea1f3e7fd70b35cb8d6214398fb0"), gethcommon.HexToAddress("0x3064455ad5cc2b376f0f3695188977d183d6b27d"), gethcommon.HexToAddress("0x720226044dafcce013c13ea0dcac76bc873a5848"), gethcommon.HexToAddress("0x6fd76a04b40073b18d46ceb8a23da1359931afd0"), gethcommon.HexToAddress("0x7e8570cb64185ad764c1689ba56c20c5072f3a76"), gethcommon.HexToAddress("0xf7cf1a5062035e1fcade4c1841e3ed7f587cd5b6"), gethcommon.HexToAddress("0x0ff65f3c24c1410c34ccef7b888d19736a036665"), gethcommon.HexToAddress("0x338cba3b652821dbf07eb89b072fb1e16f0e5c0b"), gethcommon.HexToAddress("0xfcb0eb7f31da2d7f31f684e706100ccfa4224786"), gethcommon.HexToAddress("0xd81008065ec031e540b251e9afa5a3f246e1c6fb"), gethcommon.HexToAddress("0x424a01febefe5404b3ddd92d3583acfd229e4cb9"), gethcommon.HexToAddress("0xcce47af2034b3248027578429be36eab7e0b7631"), gethcommon.HexToAddress("0x6d66645c5a3686d774c028f17237c7f4cef3eda0"), gethcommon.HexToAddress("0x3426648b4cccb33a7df81d69993b71ef5e7207d7"), gethcommon.HexToAddress("0xa3242c63f875e5442b80eb0d31873201cf3923f7"), gethcommon.HexToAddress("0x0eff12a3469b90bc6bec4a42bc826d8052adc1f5"), gethcommon.HexToAddress("0x4aa57f4785919a5a1a28749662bdf7c92a33f9d0"), gethcommon.HexToAddress("0x685974f3f0e12dd361159a5a8a4d238dd6d1820c"), gethcommon.HexToAddress("0x993f3a844fbfdb1c32bc2a30ebdedaee74b78df8"), gethcommon.HexToAddress("0xf7703717e3167215dea7d17ce2091adb6c7e7f11"), gethcommon.HexToAddress("0xeb0413c61adfd8c5295b1b359a865006af87647b"), gethcommon.HexToAddress("0xcb2494591781c50ee74420483b45c6b02ac6ec06"), gethcommon.HexToAddress("0xa9d23500890e8485e03046202caff1182de16666"), gethcommon.HexToAddress("0x5fb09ab5a6870de3c51556a4a67c79c4bb6472ca"), gethcommon.HexToAddress("0x06fe365fa22f4dfcc19b1449d569963d2568c788"), gethcommon.HexToAddress("0x9ddcc06a8897f5179c72f790f2e1c23c0247ca41"), gethcommon.HexToAddress("0x76a48ce589e8b53916e299c43e3741473c42ca6c"), gethcommon.HexToAddress("0x8e3a4d4d0afe7ad55a1935f3f69e2b666eeb09bb"), gethcommon.HexToAddress("0xe07867de534301bc3743ec4c8773ce8e6554689e"), gethcommon.HexToAddress("0x687b2df0a4fffd019cd02f8a1c873f848d4a1c54"), gethcommon.HexToAddress("0xe82bcae45bc947620274a576f3a5f96cf425e01c"), gethcommon.HexToAddress("0xc8088abd2fdaf4819230eb0fda2d9766fdf9f409"), gethcommon.HexToAddress("0x8c51ad423fe00e9cfabbbcf41b914b19056754af"), gethcommon.HexToAddress("0x0375901b7b96014f78fda77836054f1da9df939d"), gethcommon.HexToAddress("0xcb83da237a650123c7c4d492bc1fdb280bc36669"), gethcommon.HexToAddress("0xf41754def634ef1526e33dad5f17ef3d247999e1"), gethcommon.HexToAddress("0xb07b3de709c6162061cbd47e7181269618170d86"), gethcommon.HexToAddress("0xaf96d790363818ea200e9b4c4a01b571d73c67f4"), gethcommon.HexToAddress("0xabc4c47b1d8d6d3e1286df824c4bf845e9f36274"), gethcommon.HexToAddress("0xc819812695160126aecd624b446851c97e00ec52"), gethcommon.HexToAddress("0xdda6bc823b22e23b647822e21ba301bfe8f6bc85"), gethcommon.HexToAddress("0x8eb737f45fd2592aa62363e9493cde645cc2835b"), gethcommon.HexToAddress("0xdf43250d84e80948b5ba91e84f73d92f30e237e9"), gethcommon.HexToAddress("0xb345e12df498b540276517e4495f72339abf85ff"), gethcommon.HexToAddress("0x5ec82d33eadae84ec6ad8e757e458e73f569de53"),
	}

	rpcClient, err := rpc.Dial("https://goerli.infura.io/v3/9aa3d95b3bc440fa88ea12eaa4456161")
	if err != nil {
		panic(err)
	}

	ethClient := ethclient.NewClient(rpcClient)

	connString := common.CreateConnectionString(
		"eigenlabs_team",
		os.Getenv("DB_PASSWORD"),
		"eigenlabs-graph-node-production-3.cg7azkhq5rv5.us-east-1.rds.amazonaws.com",
		"5432",
		"graph_node_eigenlabs_3",
	)
	dbpool := common.MustCreateConnection(connString)
	defer dbpool.Close()
	schemaService := common.NewSubgraphSchemaService(dbpool)

	subgraphProvider, err := common.ToSubgraphProvider("satsuma")
	if err != nil {
		panic(err)
	}

	elpds := NewPaymentCalculatorDataServiceImpl(dbpool, schemaService, subgraphProvider, ethClient)

	t.Run("test GetPaymentsCalculatedUntilTimestamp", func(t *testing.T) {
		paymentsCalculatedUntilTimestamp, err := elpds.GetPaymentsCalculatedUntilTimestamp(context.Background())
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("payments calculated until timestamp: %v", paymentsCalculatedUntilTimestamp)
		t.Fail()
	})

	t.Run("test GetSharesOfStakersAtBlockNumber for beacon chain eth", func(t *testing.T) {
		strategyShares, err := elpds.GetSharesOfStakersAtBlockNumber(testBlockNumber, BEACON_CHAIN_ETH_STRATEGY_ADDRESS, stakers)
		if err != nil {
			t.Fatal(err)
		}

		for _, staker := range stakers {
			t.Logf("strategy shares for staker %s: %v", staker.Hex(), strategyShares[staker])
		}
		t.Fail()
	})

	t.Run("test GetSharesOfStakersAtBlockNumber for steth", func(t *testing.T) {
		strategyShares, err := elpds.GetSharesOfStakersAtBlockNumber(testBlockNumber, STETH_STRATEGY_ADDRESS, stakers)
		if err != nil {
			t.Fatal(err)
		}

		for _, staker := range stakers {
			t.Logf("strategy shares for staker %s: %v", staker.Hex(), strategyShares[staker])
		}
		t.Fail()
	})

}
