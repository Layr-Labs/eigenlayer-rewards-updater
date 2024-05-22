package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/signer/core"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"github.com/ethereum/go-ethereum/signer/storage"
	"strconv"
	"time"
)

func main() {
	var (
		ui                        = core.NewCommandlineUI()
		pwStorage storage.Storage = &storage.NoStorage{}
	)
	am := core.StartClefAccountManager("", false, false, "")
	defer am.Close()

	signerApi := core.NewSignerAPI(am, 17000, false, ui, nil, false, pwStorage)

	ui.RegisterUIServer(core.NewUIServerAPI(signerApi))

	// internalApi := core.NewUIServerAPI(signerApi)

	// accounts, err := signerApi.List(context.Background())
	// if err != nil {
	// 	panic(err)
	// }
	fmt.Printf("Accounts: %v\n", am.Wallets())

	status, err := am.Wallets()[0].Status()
	fmt.Printf("Status: %v, Error: %v\n", status, err)

	wallet := am.Wallets()[0]

	account, err := wallet.Derive(accounts.DefaultBaseDerivationPath, true)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Account: %v\n", account)

	salt := "some-random-string-or-hash-here"
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	nonceBytes := make([]byte, 32)
	n, err := rand.Read(nonceBytes)
	if n != 32 {
		panic(fmt.Errorf("nonce: n != 64 (bytes)"))
	} else if err != nil {
		panic(err)
	}
	nonce := hex.EncodeToString(nonceBytes)

	signerData := apitypes.TypedData{
		Types: apitypes.Types{
			"Challenge": []apitypes.Type{
				{Name: "address", Type: "address"},
				{Name: "nonce", Type: "string"},
				{Name: "timestamp", Type: "string"},
			},
			"EIP712Domain": []apitypes.Type{
				{Name: "name", Type: "string"},
				{Name: "chainId", Type: "uint256"},
				{Name: "version", Type: "string"},
				{Name: "salt", Type: "string"},
			},
		},
		PrimaryType: "Challenge",
		Domain: apitypes.TypedDataDomain{
			Name:    "ETHChallenger",
			Version: "1",
			Salt:    salt,
			ChainId: math.NewHexOrDecimal256(17000),
		},
		Message: apitypes.TypedDataMessage{
			"timestamp": timestamp,
			"address":   account.Address.Hex(),
			"nonce":     nonce,
		},
	}

	typedDataHash, _ := signerData.HashStruct(signerData.PrimaryType, signerData.Message)
	domainSeparator, _ := signerData.HashStruct("EIP712Domain", signerData.Domain.Map())

	rawData := []byte(fmt.Sprintf("\x19\x01%s%s", string(domainSeparator), string(typedDataHash)))
	// challengeHash := crypto.Keccak256Hash(rawData)

	signedText, err := wallet.SignData(account, accounts.MimetypeTypedData, rawData)

	fmt.Printf("Signed message: %+v - err: %+v\n", signedText, err)

	//hub, err := usbwallet.NewLedgerHub()
	//if err != nil {
	//	panic(err)
	//}
	//
	//wallets := hub.Wallets()
	//
	//fmt.Printf("wallets: %+v\n", wallets)
	//
	//for _, w := range wallets {
	//	status, err := w.Status()
	//	w.Open()
	//	fmt.Printf("wallet: %+v - %+v\n", status, err)
	//}

}
