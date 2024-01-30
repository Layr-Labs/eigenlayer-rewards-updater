package main

import (
	"os"
	"time"

	"github.com/Layr-Labs/eigenlayer-payment-updater/calculator"
	"github.com/Layr-Labs/eigenlayer-payment-updater/common"
	"github.com/Layr-Labs/eigenlayer-payment-updater/updater"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

func main() {

	elpc := calculator.NewMockCalculator()

	rpcClient, err := rpc.Dial("https://goerli.infura.io/v3/9aa3d95b3bc440fa88ea12eaa4456161")
	if err != nil {
		panic(err)
	}

	ethClient := ethclient.NewClient(rpcClient)

	privateKeyString := os.Getenv("PRIVATE_KEY")

	chainClient, err := common.NewChainClient(ethClient, privateKeyString)
	if err != nil {
		panic(err)
	}

	claimingManagerAddress := gethcommon.HexToAddress("0x44F49aC9B4CB1D0CC891Bfb2C0Cc5dbc34BA7181")

	elpu, err := updater.NewUpdater(time.Second*10, elpc, chainClient, claimingManagerAddress)
	if err != nil {
		panic(err)
	}

	elpu.Start()
}
