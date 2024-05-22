package signer

import (
	"fmt"
	"github.com/Layr-Labs/eigenlayer-payment-updater/pkg/config"
	"github.com/Layr-Labs/eigenlayer-payment-updater/pkg/signer/ledger"
	"github.com/Layr-Labs/eigenlayer-payment-updater/pkg/signer/privateKey"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type Signer interface {
	GetTransactOpts(chainId *big.Int) (*bind.TransactOpts, error)
	GetAddress() common.Address
}

func GetSignerForBackend(cfg *config.UpdaterConfig) (Signer, error) {
	switch cfg.SigningBackend {
	case config.SigningBackendKind_PrivateKey:
		return privateKey.NewPrivateKeySigner(cfg.PrivateKey)
	case config.SigningBackendKind_Ledger:
		return ledger.NewLedgerSigner(cfg.LedgerAddress)
	}
	return nil, fmt.Errorf("unsupported signing backend: %s", cfg.SigningBackend)
}
