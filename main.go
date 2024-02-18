package main

import (
	"math/big"
	"os"
	"time"

	"github.com/Layr-Labs/eigenlayer-payment-updater/calculator"
	"github.com/Layr-Labs/eigenlayer-payment-updater/common"
	"github.com/Layr-Labs/eigenlayer-payment-updater/updater"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

const (
	rpcUrl = "https://goerli.infura.io/v3/9aa3d95b3bc440fa88ea12eaa4456161"

	DB_USER = "eigenlabs_team"
	DB_HOST = "eigenlabs-graph-node-production-3.cg7azkhq5rv5.us-east-1.rds.amazonaws.com"
	DB_PORT = "5432"
	DB_NAME = "graph_node_eigenlabs_3"

	GOERLI_ENV = "testnet-goerli"
)

func main() {
	rpcClient, err := rpc.Dial(rpcUrl)
	if err != nil {
		panic(err)
	}

	ethClient := ethclient.NewClient(rpcClient)

	privateKeyString := os.Getenv("PRIVATE_KEY")

	chainClient, err := common.NewChainClient(ethClient, privateKeyString)
	if err != nil {
		panic(err)
	}

	connString := common.CreateConnectionString(
		DB_USER,
		os.Getenv("DB_PASSWORD"),
		DB_HOST,
		DB_PORT,
		DB_NAME,
	)
	dbpool := common.CreateConnectionOrDie(connString)
	defer dbpool.Close()

	schemaService := common.NewSubgraphSchemaService(GOERLI_ENV, dbpool)

	pds := calculator.NewPaymentsDataService(
		dbpool,
		schemaService,
	)
	osds := calculator.NewOperatorSetDataService(
		dbpool,
		schemaService,
		ethClient,
	)
	dds := calculator.NewDistributionDataServiceImpl()

	intervalSecondsLength := big.NewInt(10)

	elpc := calculator.NewRangePaymentCalculator(intervalSecondsLength, pds, osds, dds)

	elpu, err := updater.NewUpdater(time.Second*100, elpc, chainClient, calculator.CLAIMING_MANAGER_ADDRESS)
	if err != nil {
		panic(err)
	}

	elpu.Start()
}
