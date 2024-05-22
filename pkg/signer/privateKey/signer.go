package privateKey

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
)

type PrivateKeySigner struct {
	privateKey *ecdsa.PrivateKey
	address    common.Address
}

func NewPrivateKeySigner(privateKeyString string) (*PrivateKeySigner, error) {
	privateKey, err := crypto.HexToECDSA(privateKeyString)
	if err != nil {
		return nil, fmt.Errorf("cannot parse private key: %w", err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("cannot get publicKeyECDSA")
	}
	accountAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	return &PrivateKeySigner{
		privateKey: privateKey,
		address:    accountAddress,
	}, nil
}

func (pk *PrivateKeySigner) GetTransactOpts(chainId *big.Int) (*bind.TransactOpts, error) {
	opts, err := bind.NewKeyedTransactorWithChainID(pk.privateKey, chainId)
	if err != nil {
		return nil, fmt.Errorf("EstimateGasPriceAndLimitAndSendTx: cannot create transactOpts: %w", err)
	}
	return opts, nil
}

func (pk *PrivateKeySigner) GetAddress() common.Address {
	return pk.address
}
