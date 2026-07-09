package keeper

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/Layr-Labs/eigenlayer-rewards-updater/pkg/chainClient"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethcommon "github.com/ethereum/go-ethereum/common"
)

// Minimal ABI for the bits of EmissionsController the keeper uses. The contract
// is a TransparentUpgradeableProxy (impl EmissionsController); pressButton is
// permissionless (only nonReentrant + onlyWhenNotPaused), so any funded caller
// can press it — no on-chain role is required.
const emissionsControllerABI = `[
  {"type":"function","name":"pressButton","stateMutability":"nonpayable","inputs":[{"name":"length","type":"uint256"}],"outputs":[]},
  {"type":"function","name":"isButtonPressable","stateMutability":"view","inputs":[],"outputs":[{"name":"","type":"bool"}]},
  {"type":"function","name":"getTotalProcessableDistributions","stateMutability":"view","inputs":[],"outputs":[{"name":"","type":"uint256"}]},
  {"type":"function","name":"getCurrentEpoch","stateMutability":"view","inputs":[],"outputs":[{"name":"","type":"uint256"}]}
]`

// EmissionsController is the subset of the on-chain contract the keeper needs.
// It is an interface so the keeper flow can be unit-tested without a live chain.
type EmissionsController interface {
	IsButtonPressable(ctx context.Context) (bool, error)
	TotalProcessableDistributions(ctx context.Context) (*big.Int, error)
	CurrentEpoch(ctx context.Context) (*big.Int, error)
	// PressButton submits pressButton(length) and blocks until it is mined,
	// returning the transaction hash. length caps how many distributions are
	// processed in one call; callers pass MaxUint256 to process all.
	PressButton(ctx context.Context, length *big.Int) (string, error)
}

// boundEmissionsController implements EmissionsController against a live chain
// via the shared ChainClient (which carries the KMS/private-key signer).
type boundEmissionsController struct {
	cc       *chainClient.ChainClient
	contract *bind.BoundContract
}

func NewEmissionsController(cc *chainClient.ChainClient, address gethcommon.Address) (EmissionsController, error) {
	parsed, err := abi.JSON(strings.NewReader(emissionsControllerABI))
	if err != nil {
		return nil, fmt.Errorf("failed to parse EmissionsController ABI: %w", err)
	}
	contract := bind.NewBoundContract(address, parsed, cc.Client, cc.Client, cc.Client)
	return &boundEmissionsController{cc: cc, contract: contract}, nil
}

func (e *boundEmissionsController) callBool(ctx context.Context, method string) (bool, error) {
	var out []interface{}
	if err := e.contract.Call(&bind.CallOpts{Context: ctx}, &out, method); err != nil {
		return false, err
	}
	v, ok := out[0].(bool)
	if !ok {
		return false, fmt.Errorf("%s: unexpected return type", method)
	}
	return v, nil
}

func (e *boundEmissionsController) callBig(ctx context.Context, method string) (*big.Int, error) {
	var out []interface{}
	if err := e.contract.Call(&bind.CallOpts{Context: ctx}, &out, method); err != nil {
		return nil, err
	}
	v, ok := out[0].(*big.Int)
	if !ok {
		return nil, fmt.Errorf("%s: unexpected return type", method)
	}
	return v, nil
}

func (e *boundEmissionsController) IsButtonPressable(ctx context.Context) (bool, error) {
	return e.callBool(ctx, "isButtonPressable")
}

func (e *boundEmissionsController) TotalProcessableDistributions(ctx context.Context) (*big.Int, error) {
	return e.callBig(ctx, "getTotalProcessableDistributions")
}

func (e *boundEmissionsController) CurrentEpoch(ctx context.Context) (*big.Int, error) {
	return e.callBig(ctx, "getCurrentEpoch")
}

func (e *boundEmissionsController) PressButton(ctx context.Context, length *big.Int) (string, error) {
	// Build an unsent tx (NoSend opts carry the signer), then re-price and send
	// via the shared gas-estimation path — same pattern as submitRoot.
	tx, err := e.contract.Transact(e.cc.NoSendTransactOpts, "pressButton", length)
	if err != nil {
		return "", fmt.Errorf("failed to build pressButton txn: %w", err)
	}
	receipt, err := e.cc.EstimateGasPriceAndLimitAndSendTx(ctx, tx, "pressButton")
	if err != nil {
		return "", fmt.Errorf("failed to send pressButton txn: %w", err)
	}
	if receipt.Status != 1 {
		return receipt.TxHash.Hex(), chainClient.ErrTransactionFailed
	}
	return receipt.TxHash.Hex(), nil
}
