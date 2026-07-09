package txsigner

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

// PrivateKeySigner signs transactions with a local secp256k1 private key.
// Intended for local development and testnets; prefer AWSKMSSigner in production.
type PrivateKeySigner struct {
	privateKey *ecdsa.PrivateKey
	address    common.Address
}

// NewPrivateKeySigner builds a signer from a hex-encoded private key (no 0x prefix).
func NewPrivateKeySigner(privateKeyHex string) (*PrivateKeySigner, error) {
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %w", err)
	}
	return &PrivateKeySigner{
		privateKey: privateKey,
		address:    crypto.PubkeyToAddress(privateKey.PublicKey),
	}, nil
}

func (s *PrivateKeySigner) GetAddress() common.Address {
	return s.address
}

func (s *PrivateKeySigner) SignerFn(chainID *big.Int) bind.SignerFn {
	opts, err := bind.NewKeyedTransactorWithChainID(s.privateKey, chainID)
	if err != nil {
		// NewKeyedTransactorWithChainID only errors on a nil chainID; surface
		// it through the SignerFn so the caller sees it at signing time.
		return func(common.Address, *types.Transaction) (*types.Transaction, error) {
			return nil, fmt.Errorf("failed to build keyed transactor: %w", err)
		}
	}
	return opts.Signer
}
