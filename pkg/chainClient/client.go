package chainClient

import (
	"context"
	"errors"
	"fmt"
	"math/big"

	"github.com/Layr-Labs/eigenlayer-rewards-updater/pkg/txsigner"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rs/zerolog/log"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

var (
	FallbackGasTipCap    = big.NewInt(15000000000)
	ErrTransactionFailed = errors.New("ErrTransactionFailed")
)

type ChainClient struct {
	*ethclient.Client
	signer             txsigner.ITransactionSigner
	chainID            *big.Int
	AccountAddress     common.Address
	NoSendTransactOpts *bind.TransactOpts
	Contracts          map[common.Address]*bind.BoundContract
}

// NewChainClient builds a ChainClient that signs with a local private key.
// An empty privateKeyString yields a read-only client (no signer / no
// NoSendTransactOpts), preserving the previous behavior.
func NewChainClient(ctx context.Context, ethClient *ethclient.Client, privateKeyString string) (*ChainClient, error) {
	span, ctx := tracer.StartSpanFromContext(ctx, "chainClient::NewChainClient")
	defer span.Finish()

	if len(privateKeyString) == 0 {
		return newChainClient(ctx, ethClient, nil)
	}

	signer, err := txsigner.NewPrivateKeySigner(privateKeyString)
	if err != nil {
		return nil, fmt.Errorf("NewClient: cannot parse private key: %w", err)
	}
	return newChainClient(ctx, ethClient, signer)
}

// NewChainClientWithSigner builds a ChainClient backed by an arbitrary signer
// (e.g. AWS KMS), so the signing key never has to live on the host.
func NewChainClientWithSigner(ctx context.Context, ethClient *ethclient.Client, signer txsigner.ITransactionSigner) (*ChainClient, error) {
	span, ctx := tracer.StartSpanFromContext(ctx, "chainClient::NewChainClientWithSigner")
	defer span.Finish()

	if signer == nil {
		return nil, errors.New("NewChainClientWithSigner: signer must not be nil")
	}
	return newChainClient(ctx, ethClient, signer)
}

func newChainClient(ctx context.Context, ethClient *ethclient.Client, signer txsigner.ITransactionSigner) (*ChainClient, error) {
	c := &ChainClient{
		Client:    ethClient,
		signer:    signer,
		Contracts: make(map[common.Address]*bind.BoundContract),
	}

	if signer == nil {
		return c, nil
	}

	chainID, err := ethClient.ChainID(ctx)
	if err != nil {
		return nil, fmt.Errorf("NewClient: cannot get chainId: %w", err)
	}
	c.chainID = chainID
	c.AccountAddress = signer.GetAddress()

	// generate and memoize NoSendTransactOpts
	c.NoSendTransactOpts = &bind.TransactOpts{
		From:    c.AccountAddress,
		Signer:  signer.SignerFn(chainID),
		Context: ctx,
		NoSend:  true,
	}

	return c, nil
}

func (c *ChainClient) GetCurrentBlockNumber(ctx context.Context) (uint32, error) {
	bn, err := c.Client.BlockNumber(ctx)
	return uint32(bn), err
}

func (c *ChainClient) GetAccountAddress() common.Address {
	return c.AccountAddress
}

func (c *ChainClient) GetNoSendTransactOpts() *bind.TransactOpts {
	return c.NoSendTransactOpts
}

// EstimateGasPriceAndLimitAndSendTx sends and returns an otherwise identical txn
// to the one provided but with updated gas prices sampled from the existing network
// conditions and an accurate gasLimit
//
// Note: tx must be a to a contract, not an EOA
//
// Slightly modified from: https://github.com/ethereum-optimism/optimism/blob/ec266098641820c50c39c31048aa4e953bece464/batch-submitter/drivers/sequencer/driver.go#L314
func (c *ChainClient) EstimateGasPriceAndLimitAndSendTx(
	ctx context.Context,
	tx *types.Transaction,
	tag string,
) (*types.Receipt, error) {
	span, ctx := tracer.StartSpanFromContext(ctx, "chainClient::EstimateGasPriceAndLimitAndSendTx")
	defer span.Finish()

	gasTipCap, err := c.SuggestGasTipCap(ctx)
	if err != nil {
		// If the transaction failed because the backend does not support
		// eth_maxPriorityFeePerGas, fallback to using the default constant.
		// Currently Alchemy is the only backend provider that exposes this
		// method, so in the event their API is unreachable we can fallback to a
		// degraded mode of operation. This also applies to our test
		// environments, as hardhat doesn't support the query either.
		log.Debug().Msgf("EstimateGasPriceAndLimitAndSendTx: cannot get gasTipCap: %v", err)
		gasTipCap = FallbackGasTipCap
	}

	header, err := c.HeaderByNumber(ctx, nil)
	if err != nil {
		return nil, err
	}
	// get header basefee * 3/2
	overestimatedBasefee := new(big.Int).Div(new(big.Int).Mul(header.BaseFee, big.NewInt(3)), big.NewInt(2))

	gasFeeCap := new(big.Int).Add(overestimatedBasefee, gasTipCap)

	// The estimated gas limits performed by RawTransact fail semi-regularly
	// with out of gas exceptions. To remedy this we extract the internal calls
	// to perform gas price/gas limit estimation here and add a buffer to
	// account for any network variability.
	gasLimit, err := c.Client.EstimateGas(ctx, ethereum.CallMsg{
		From:      c.AccountAddress,
		To:        tx.To(),
		GasTipCap: gasTipCap,
		GasFeeCap: gasFeeCap,
		Value:     nil,
		Data:      tx.Data(),
	})

	if err != nil {
		return nil, err
	}

	if c.signer == nil {
		return nil, errors.New("EstimateGasPriceAndLimitAndSendTx: chain client has no signer")
	}
	opts := &bind.TransactOpts{
		From:      c.AccountAddress,
		Signer:    c.signer.SignerFn(c.chainID),
		Context:   ctx,
		Nonce:     new(big.Int).SetUint64(tx.Nonce()),
		GasTipCap: gasTipCap,
		GasFeeCap: gasFeeCap,
		GasLimit:  addGasBuffer(gasLimit),
	}

	contract := c.Contracts[*tx.To()]
	// if the contract has not been cached
	if contract == nil {
		// create a dummy bound contract tied to the `to` address of the transaction
		contract = bind.NewBoundContract(*tx.To(), abi.ABI{}, c.Client, c.Client, c.Client)
		// cache the contract for later use
		c.Contracts[*tx.To()] = contract
	}

	log.Info().Msgf("EstimateGasPriceAndLimitAndSendTx: sending txn (%s) with gasTipCap=%v gasFeeCap=%v gasLimit=%v", tag, gasTipCap, gasFeeCap, opts.GasLimit)

	tx, err = contract.RawTransact(opts, tx.Data())
	if err != nil {
		return nil, fmt.Errorf("EstimateGasPriceAndLimitAndSendTx: failed to send txn (%s): %w", tag, err)
	}

	log.Info().Msgf("EstimateGasPriceAndLimitAndSendTx: sent txn (%s) with hash=%s", tag, tx.Hash().Hex())

	receipt, err := c.EnsureTransactionEvaled(ctx, tx, tag)
	if err != nil {
		return nil, err
	}

	return receipt, err
}

func (c *ChainClient) EnsureTransactionEvaled(ctx context.Context, tx *types.Transaction, tag string) (*types.Receipt, error) {
	span, ctx := tracer.StartSpanFromContext(ctx, "chainClient::EnsureTransactionEvaled")
	defer span.Finish()

	log.Info().Msgf("EnsureTransactionEvaled entered")

	receipt, err := bind.WaitMined(ctx, c.Client, tx)
	if err != nil {
		return nil, fmt.Errorf("EnsureTransactionEvaled: failed to wait for transaction (%s) to mine: %w", tag, err)
	}
	if receipt.Status != 1 {
		log.Debug().Msgf("EnsureTransactionEvaled: transaction (%s) failed: %v", tag, receipt)
		return nil, ErrTransactionFailed
	}
	log.Debug().Msgf("EnsureTransactionEvaled: transaction (%s) succeeded: %v", tag, receipt.TxHash.Hex())
	return receipt, nil
}

func addGasBuffer(gasLimit uint64) uint64 {
	return 6 * gasLimit / 5 // add 20% buffer to gas limit
}
