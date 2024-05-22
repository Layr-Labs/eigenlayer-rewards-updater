package ledger

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/signer/core"
	"github.com/ethereum/go-ethereum/signer/storage"
	"math/big"
	"regexp"
)

type LedgerSigner struct {
	address common.Address
}

func NewLedgerSigner(address common.Address) (*LedgerSigner, error) {
	return &LedgerSigner{
		address: address,
	}, nil
}

func (ls *LedgerSigner) GetTransactOpts(chainId *big.Int) (*bind.TransactOpts, error) {
	return &bind.TransactOpts{
		From: ls.address,
		Signer: func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
			if address != ls.address {
				return nil, bind.ErrNotAuthorized
			}
			return ls.SignTx(tx, chainId)
		},
		Context: context.Background(),
	}, nil
}

func (ls *LedgerSigner) GetAddress() common.Address {
	return ls.address
}

func (ls *LedgerSigner) SignTx(tx *types.Transaction, chainId *big.Int) (*types.Transaction, error) {
	var (
		ui                        = core.NewCommandlineUI()
		pwStorage storage.Storage = &storage.NoStorage{}
	)
	am := core.StartClefAccountManager("", false, false, "")
	defer am.Close()

	signerApi := core.NewSignerAPI(am, 17000, false, ui, nil, false, pwStorage)

	ui.RegisterUIServer(core.NewUIServerAPI(signerApi))

	if len(am.Wallets()) == 0 {
		return nil, fmt.Errorf("LedgerSigner: no wallets found")
	}

	wallet := am.Wallets()[0]

	status, err := wallet.Status()

	statusRegex := regexp.MustCompile(`online$`)
	if !statusRegex.MatchString(status) || err != nil {
		return nil, fmt.Errorf("LedgerSigner: wallet is locked")
	}

	account, err := wallet.Derive(accounts.DefaultBaseDerivationPath, true)

	if err != nil {
		return nil, fmt.Errorf("LedgerSigner: failed to derive account")
	}

	if account.Address.String() != ls.address.String() {
		return nil, fmt.Errorf("LedgerSigner: address from wallet does not match expected signer address")
	}

	return wallet.SignTx(account, tx, chainId)
}
