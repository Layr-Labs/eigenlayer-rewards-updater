package main

import (
	"os"

	"database/sql"

	gethcommon "github.com/ethereum/go-ethereum/common"

	"github.com/Layr-Labs/eigenlayer-payment-updater/common"
	"github.com/Layr-Labs/eigenlayer-payment-updater/common/services"
	"github.com/Layr-Labs/eigenlayer-payment-updater/updater"
	"github.com/ethereum/go-ethereum/ethclient"
	drv "github.com/uber/athenadriver/go"
)

const paymentCoordinatorAddress = "0x56c119bD92Af45eb74443ab14D4e93B7f5C67896"

const region = "us-east-1"
const outputBucket = "s3://payment-poc-mock/query-results/"

func main() {
	ethClient, err := ethclient.Dial(os.Getenv("RPC_URL"))
	if err != nil {
		panic(err)
	}

	chainClient, err := common.NewChainClient(ethClient, os.Getenv("PRIVATE_KEY"))
	if err != nil {
		panic(err)
	}

	transactor, err := services.NewTransactor(chainClient, gethcommon.HexToAddress(paymentCoordinatorAddress))
	if err != nil {
		panic(err)
	}

	// Step 1. Set AWS Credential in Driver Config.
	conf, _ := drv.NewDefaultConfig(outputBucket, region, os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"))
	// Step 2. Open Connection.
	db, _ := sql.Open(drv.DriverName, conf.Stringify())
	defer db.Close()

	dds := services.NewDistributionDataService(db, transactor)

	u, err := updater.NewUpdater(1000, transactor, dds)
	if err != nil {
		panic(err)
	}

	u.Start()
}
