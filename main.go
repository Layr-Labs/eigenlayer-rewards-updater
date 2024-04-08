package main

// import (
// 	"os"

// 	"github.com/Layr-Labs/eigenlayer-payment-updater/common"
// 	"github.com/Layr-Labs/eigenlayer-payment-updater/common/services"
// 	"github.com/Layr-Labs/eigenlayer-payment-updater/updater"
// 	"github.com/ethereum/go-ethereum/ethclient"
// 	"github.com/ethereum/go-ethereum/rpc"

// 	gethcommon "github.com/ethereum/go-ethereum/common"
// )

// const (
// 	rpcUrl = "https://goerli.infura.io/v3/9aa3d95b3bc440fa88ea12eaa4456161"

// 	DB_USER = "eigenlabs_team"
// 	DB_HOST = "eigenlabs-graph-node-production-3.cg7azkhq5rv5.us-east-1.rds.amazonaws.com"
// 	DB_PORT = "5432"
// 	DB_NAME = "graph_node_eigenlabs_3"

// 	GOERLI_ENV = "testnet-goerli"
// )

// var claimingManagerAddress = gethcommon.HexToAddress("0x7b3f8f4b8e3b7f4f19e7e3f0b7b6e3f1f1b2f1b")

// func main() {
// 	rpcClient, err := rpc.Dial(rpcUrl)
// 	if err != nil {
// 		panic(err)
// 	}

// 	ethClient := ethclient.NewClient(rpcClient)

// 	privateKeyString := os.Getenv("PRIVATE_KEY")

// 	chainClient, err := common.NewChainClient(ethClient, privateKeyString)
// 	if err != nil {
// 		panic(err)
// 	}

// 	connString := common.CreateConnectionString(
// 		DB_USER,
// 		os.Getenv("DB_PASSWORD"),
// 		DB_HOST,
// 		DB_PORT,
// 		DB_NAME,
// 	)
// 	dbpool := common.CreateConnectionOrDie(connString)
// 	defer dbpool.Close()

// 	dds := services.NewDistributionDataServiceImpl(dbpool)

// 	updateIntervalSeconds := 100

// 	// elpu, err := updater.NewUpdater(updateIntervalSeconds, dds, chainClient, claimingManagerAddress)
// 	if err != nil {
// 		panic(err)
// 	}

// 	elpu.Start()
// }
