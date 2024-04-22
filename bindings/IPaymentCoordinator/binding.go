// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractIPaymentCoordinator

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// IPaymentCoordinatorEarnerTreeMerkleLeaf is an auto generated low-level Go binding around an user-defined struct.
type IPaymentCoordinatorEarnerTreeMerkleLeaf struct {
	Earner          common.Address
	EarnerTokenRoot [32]byte
}

// IPaymentCoordinatorPaymentMerkleClaim is an auto generated low-level Go binding around an user-defined struct.
type IPaymentCoordinatorPaymentMerkleClaim struct {
	RootIndex       uint32
	EarnerIndex     uint32
	EarnerTreeProof []byte
	EarnerLeaf      IPaymentCoordinatorEarnerTreeMerkleLeaf
	TokenIndices    []uint32
	TokenTreeProofs [][]byte
	TokenLeaves     []IPaymentCoordinatorTokenTreeMerkleLeaf
}

// IPaymentCoordinatorRangePayment is an auto generated low-level Go binding around an user-defined struct.
type IPaymentCoordinatorRangePayment struct {
	StrategiesAndMultipliers []IPaymentCoordinatorStrategyAndMultiplier
	Token                    common.Address
	Amount                   *big.Int
	StartTimestamp           uint64
	Duration                 uint64
}

// IPaymentCoordinatorStrategyAndMultiplier is an auto generated low-level Go binding around an user-defined struct.
type IPaymentCoordinatorStrategyAndMultiplier struct {
	Strategy   common.Address
	Multiplier *big.Int
}

// IPaymentCoordinatorTokenTreeMerkleLeaf is an auto generated low-level Go binding around an user-defined struct.
type IPaymentCoordinatorTokenTreeMerkleLeaf struct {
	Token              common.Address
	CumulativeEarnings *big.Int
}

// ContractIPaymentCoordinatorMetaData contains all meta data concerning the ContractIPaymentCoordinator contract.
var ContractIPaymentCoordinatorMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"GENESIS_PAYMENT_TIMESTAMP\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_FUTURE_LENGTH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_PAYMENT_DURATION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_RETROACTIVE_LENGTH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"activationDelay\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateEarnerLeafHash\",\"inputs\":[{\"name\":\"leaf\",\"type\":\"tuple\",\"internalType\":\"structIPaymentCoordinator.EarnerTreeMerkleLeaf\",\"components\":[{\"name\":\"earner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"earnerTokenRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"calculateTokenLeafHash\",\"inputs\":[{\"name\":\"leaf\",\"type\":\"tuple\",\"internalType\":\"structIPaymentCoordinator.TokenTreeMerkleLeaf\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"cumulativeEarnings\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"calculationIntervalSeconds\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkClaim\",\"inputs\":[{\"name\":\"claim\",\"type\":\"tuple\",\"internalType\":\"structIPaymentCoordinator.PaymentMerkleClaim\",\"components\":[{\"name\":\"rootIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"earnerIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"earnerTreeProof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"earnerLeaf\",\"type\":\"tuple\",\"internalType\":\"structIPaymentCoordinator.EarnerTreeMerkleLeaf\",\"components\":[{\"name\":\"earner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"earnerTokenRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"tokenIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"tokenTreeProofs\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"tokenLeaves\",\"type\":\"tuple[]\",\"internalType\":\"structIPaymentCoordinator.TokenTreeMerkleLeaf[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"cumulativeEarnings\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claimerFor\",\"inputs\":[{\"name\":\"earner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"cumulativeClaimed\",\"inputs\":[{\"name\":\"claimer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"currPaymentCalculationEndTimestamp\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRootIndexFromHash\",\"inputs\":[{\"name\":\"rootHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"globalOperatorCommissionBips\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"payAllForRange\",\"inputs\":[{\"name\":\"rangePayment\",\"type\":\"tuple[]\",\"internalType\":\"structIPaymentCoordinator.RangePayment[]\",\"components\":[{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[]\",\"internalType\":\"structIPaymentCoordinator.StrategyAndMultiplier[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"duration\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"payForRange\",\"inputs\":[{\"name\":\"rangePayments\",\"type\":\"tuple[]\",\"internalType\":\"structIPaymentCoordinator.RangePayment[]\",\"components\":[{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[]\",\"internalType\":\"structIPaymentCoordinator.StrategyAndMultiplier[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"duration\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paymentUpdater\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"processClaim\",\"inputs\":[{\"name\":\"claim\",\"type\":\"tuple\",\"internalType\":\"structIPaymentCoordinator.PaymentMerkleClaim\",\"components\":[{\"name\":\"rootIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"earnerIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"earnerTreeProof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"earnerLeaf\",\"type\":\"tuple\",\"internalType\":\"structIPaymentCoordinator.EarnerTreeMerkleLeaf\",\"components\":[{\"name\":\"earner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"earnerTokenRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"tokenIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"tokenTreeProofs\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"tokenLeaves\",\"type\":\"tuple[]\",\"internalType\":\"structIPaymentCoordinator.TokenTreeMerkleLeaf[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"cumulativeEarnings\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setActivationDelay\",\"inputs\":[{\"name\":\"_activationDelay\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setClaimerFor\",\"inputs\":[{\"name\":\"claimer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setGlobalOperatorCommission\",\"inputs\":[{\"name\":\"_globalCommissionBips\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPaymentUpdater\",\"inputs\":[{\"name\":\"_paymentUpdater\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitRoot\",\"inputs\":[{\"name\":\"root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"paymentCalculationEndTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ActivationDelaySet\",\"inputs\":[{\"name\":\"oldActivationDelay\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"newActivationDelay\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CalculationIntervalSecondsSet\",\"inputs\":[{\"name\":\"oldCalculationIntervalSeconds\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"newCalculationIntervalSeconds\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ClaimerForSet\",\"inputs\":[{\"name\":\"earner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"oldClaimer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"claimer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DistributionRootSubmitted\",\"inputs\":[{\"name\":\"rootIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"root\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"paymentCalculationEndTimestamp\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"activatedAt\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GlobalCommissionBipsSet\",\"inputs\":[{\"name\":\"oldGlobalCommissionBips\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"newGlobalCommissionBips\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PayAllForRangeSubmitterSet\",\"inputs\":[{\"name\":\"payAllForRangeSubmitter\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"oldValue\",\"type\":\"bool\",\"indexed\":true,\"internalType\":\"bool\"},{\"name\":\"newValue\",\"type\":\"bool\",\"indexed\":true,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PaymentClaimed\",\"inputs\":[{\"name\":\"root\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"leaf\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIPaymentCoordinator.TokenTreeMerkleLeaf\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"cumulativeEarnings\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PaymentUpdaterSet\",\"inputs\":[{\"name\":\"oldPaymentUpdater\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPaymentUpdater\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RangePaymentCreated\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"paymentNonce\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"rangePaymentHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"rangePayment\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIPaymentCoordinator.RangePayment\",\"components\":[{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[]\",\"internalType\":\"structIPaymentCoordinator.StrategyAndMultiplier[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"duration\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RangePaymentForAllCreated\",\"inputs\":[{\"name\":\"submitter\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"paymentNonce\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"rangePaymentHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"rangePayment\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIPaymentCoordinator.RangePayment\",\"components\":[{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[]\",\"internalType\":\"structIPaymentCoordinator.StrategyAndMultiplier[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"duration\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"anonymous\":false}]",
}

// ContractIPaymentCoordinatorABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractIPaymentCoordinatorMetaData.ABI instead.
var ContractIPaymentCoordinatorABI = ContractIPaymentCoordinatorMetaData.ABI

// ContractIPaymentCoordinator is an auto generated Go binding around an Ethereum contract.
type ContractIPaymentCoordinator struct {
	ContractIPaymentCoordinatorCaller     // Read-only binding to the contract
	ContractIPaymentCoordinatorTransactor // Write-only binding to the contract
	ContractIPaymentCoordinatorFilterer   // Log filterer for contract events
}

// ContractIPaymentCoordinatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractIPaymentCoordinatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractIPaymentCoordinatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractIPaymentCoordinatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractIPaymentCoordinatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractIPaymentCoordinatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractIPaymentCoordinatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractIPaymentCoordinatorSession struct {
	Contract     *ContractIPaymentCoordinator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                // Call options to use throughout this session
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ContractIPaymentCoordinatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractIPaymentCoordinatorCallerSession struct {
	Contract *ContractIPaymentCoordinatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                      // Call options to use throughout this session
}

// ContractIPaymentCoordinatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractIPaymentCoordinatorTransactorSession struct {
	Contract     *ContractIPaymentCoordinatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                      // Transaction auth options to use throughout this session
}

// ContractIPaymentCoordinatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractIPaymentCoordinatorRaw struct {
	Contract *ContractIPaymentCoordinator // Generic contract binding to access the raw methods on
}

// ContractIPaymentCoordinatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractIPaymentCoordinatorCallerRaw struct {
	Contract *ContractIPaymentCoordinatorCaller // Generic read-only contract binding to access the raw methods on
}

// ContractIPaymentCoordinatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractIPaymentCoordinatorTransactorRaw struct {
	Contract *ContractIPaymentCoordinatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractIPaymentCoordinator creates a new instance of ContractIPaymentCoordinator, bound to a specific deployed contract.
func NewContractIPaymentCoordinator(address common.Address, backend bind.ContractBackend) (*ContractIPaymentCoordinator, error) {
	contract, err := bindContractIPaymentCoordinator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractIPaymentCoordinator{ContractIPaymentCoordinatorCaller: ContractIPaymentCoordinatorCaller{contract: contract}, ContractIPaymentCoordinatorTransactor: ContractIPaymentCoordinatorTransactor{contract: contract}, ContractIPaymentCoordinatorFilterer: ContractIPaymentCoordinatorFilterer{contract: contract}}, nil
}

// NewContractIPaymentCoordinatorCaller creates a new read-only instance of ContractIPaymentCoordinator, bound to a specific deployed contract.
func NewContractIPaymentCoordinatorCaller(address common.Address, caller bind.ContractCaller) (*ContractIPaymentCoordinatorCaller, error) {
	contract, err := bindContractIPaymentCoordinator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractIPaymentCoordinatorCaller{contract: contract}, nil
}

// NewContractIPaymentCoordinatorTransactor creates a new write-only instance of ContractIPaymentCoordinator, bound to a specific deployed contract.
func NewContractIPaymentCoordinatorTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractIPaymentCoordinatorTransactor, error) {
	contract, err := bindContractIPaymentCoordinator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractIPaymentCoordinatorTransactor{contract: contract}, nil
}

// NewContractIPaymentCoordinatorFilterer creates a new log filterer instance of ContractIPaymentCoordinator, bound to a specific deployed contract.
func NewContractIPaymentCoordinatorFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractIPaymentCoordinatorFilterer, error) {
	contract, err := bindContractIPaymentCoordinator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractIPaymentCoordinatorFilterer{contract: contract}, nil
}

// bindContractIPaymentCoordinator binds a generic wrapper to an already deployed contract.
func bindContractIPaymentCoordinator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractIPaymentCoordinatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractIPaymentCoordinator.Contract.ContractIPaymentCoordinatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractIPaymentCoordinator.Contract.ContractIPaymentCoordinatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractIPaymentCoordinator.Contract.ContractIPaymentCoordinatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractIPaymentCoordinator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractIPaymentCoordinator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractIPaymentCoordinator.Contract.contract.Transact(opts, method, params...)
}

// GENESISPAYMENTTIMESTAMP is a free data retrieval call binding the contract method 0x2cfd45eb.
//
// Solidity: function GENESIS_PAYMENT_TIMESTAMP() view returns(uint64)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorCaller) GENESISPAYMENTTIMESTAMP(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _ContractIPaymentCoordinator.contract.Call(opts, &out, "GENESIS_PAYMENT_TIMESTAMP")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GENESISPAYMENTTIMESTAMP is a free data retrieval call binding the contract method 0x2cfd45eb.
//
// Solidity: function GENESIS_PAYMENT_TIMESTAMP() view returns(uint64)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorSession) GENESISPAYMENTTIMESTAMP() (uint64, error) {
	return _ContractIPaymentCoordinator.Contract.GENESISPAYMENTTIMESTAMP(&_ContractIPaymentCoordinator.CallOpts)
}

// GENESISPAYMENTTIMESTAMP is a free data retrieval call binding the contract method 0x2cfd45eb.
//
// Solidity: function GENESIS_PAYMENT_TIMESTAMP() view returns(uint64)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorCallerSession) GENESISPAYMENTTIMESTAMP() (uint64, error) {
	return _ContractIPaymentCoordinator.Contract.GENESISPAYMENTTIMESTAMP(&_ContractIPaymentCoordinator.CallOpts)
}

// MAXFUTURELENGTH is a free data retrieval call binding the contract method 0x04a0c502.
//
// Solidity: function MAX_FUTURE_LENGTH() view returns(uint64)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorCaller) MAXFUTURELENGTH(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _ContractIPaymentCoordinator.contract.Call(opts, &out, "MAX_FUTURE_LENGTH")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// MAXFUTURELENGTH is a free data retrieval call binding the contract method 0x04a0c502.
//
// Solidity: function MAX_FUTURE_LENGTH() view returns(uint64)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorSession) MAXFUTURELENGTH() (uint64, error) {
	return _ContractIPaymentCoordinator.Contract.MAXFUTURELENGTH(&_ContractIPaymentCoordinator.CallOpts)
}

// MAXFUTURELENGTH is a free data retrieval call binding the contract method 0x04a0c502.
//
// Solidity: function MAX_FUTURE_LENGTH() view returns(uint64)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorCallerSession) MAXFUTURELENGTH() (uint64, error) {
	return _ContractIPaymentCoordinator.Contract.MAXFUTURELENGTH(&_ContractIPaymentCoordinator.CallOpts)
}

// MAXPAYMENTDURATION is a free data retrieval call binding the contract method 0xee619597.
//
// Solidity: function MAX_PAYMENT_DURATION() view returns(uint64)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorCaller) MAXPAYMENTDURATION(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _ContractIPaymentCoordinator.contract.Call(opts, &out, "MAX_PAYMENT_DURATION")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// MAXPAYMENTDURATION is a free data retrieval call binding the contract method 0xee619597.
//
// Solidity: function MAX_PAYMENT_DURATION() view returns(uint64)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorSession) MAXPAYMENTDURATION() (uint64, error) {
	return _ContractIPaymentCoordinator.Contract.MAXPAYMENTDURATION(&_ContractIPaymentCoordinator.CallOpts)
}

// MAXPAYMENTDURATION is a free data retrieval call binding the contract method 0xee619597.
//
// Solidity: function MAX_PAYMENT_DURATION() view returns(uint64)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorCallerSession) MAXPAYMENTDURATION() (uint64, error) {
	return _ContractIPaymentCoordinator.Contract.MAXPAYMENTDURATION(&_ContractIPaymentCoordinator.CallOpts)
}

// MAXRETROACTIVELENGTH is a free data retrieval call binding the contract method 0x37838ed0.
//
// Solidity: function MAX_RETROACTIVE_LENGTH() view returns(uint64)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorCaller) MAXRETROACTIVELENGTH(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _ContractIPaymentCoordinator.contract.Call(opts, &out, "MAX_RETROACTIVE_LENGTH")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// MAXRETROACTIVELENGTH is a free data retrieval call binding the contract method 0x37838ed0.
//
// Solidity: function MAX_RETROACTIVE_LENGTH() view returns(uint64)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorSession) MAXRETROACTIVELENGTH() (uint64, error) {
	return _ContractIPaymentCoordinator.Contract.MAXRETROACTIVELENGTH(&_ContractIPaymentCoordinator.CallOpts)
}

// MAXRETROACTIVELENGTH is a free data retrieval call binding the contract method 0x37838ed0.
//
// Solidity: function MAX_RETROACTIVE_LENGTH() view returns(uint64)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorCallerSession) MAXRETROACTIVELENGTH() (uint64, error) {
	return _ContractIPaymentCoordinator.Contract.MAXRETROACTIVELENGTH(&_ContractIPaymentCoordinator.CallOpts)
}

// ActivationDelay is a free data retrieval call binding the contract method 0x3a8c0786.
//
// Solidity: function activationDelay() view returns(uint64)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorCaller) ActivationDelay(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _ContractIPaymentCoordinator.contract.Call(opts, &out, "activationDelay")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ActivationDelay is a free data retrieval call binding the contract method 0x3a8c0786.
//
// Solidity: function activationDelay() view returns(uint64)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorSession) ActivationDelay() (uint64, error) {
	return _ContractIPaymentCoordinator.Contract.ActivationDelay(&_ContractIPaymentCoordinator.CallOpts)
}

// ActivationDelay is a free data retrieval call binding the contract method 0x3a8c0786.
//
// Solidity: function activationDelay() view returns(uint64)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorCallerSession) ActivationDelay() (uint64, error) {
	return _ContractIPaymentCoordinator.Contract.ActivationDelay(&_ContractIPaymentCoordinator.CallOpts)
}

// CalculateEarnerLeafHash is a free data retrieval call binding the contract method 0x149bc872.
//
// Solidity: function calculateEarnerLeafHash((address,bytes32) leaf) pure returns(bytes32)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorCaller) CalculateEarnerLeafHash(opts *bind.CallOpts, leaf IPaymentCoordinatorEarnerTreeMerkleLeaf) ([32]byte, error) {
	var out []interface{}
	err := _ContractIPaymentCoordinator.contract.Call(opts, &out, "calculateEarnerLeafHash", leaf)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateEarnerLeafHash is a free data retrieval call binding the contract method 0x149bc872.
//
// Solidity: function calculateEarnerLeafHash((address,bytes32) leaf) pure returns(bytes32)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorSession) CalculateEarnerLeafHash(leaf IPaymentCoordinatorEarnerTreeMerkleLeaf) ([32]byte, error) {
	return _ContractIPaymentCoordinator.Contract.CalculateEarnerLeafHash(&_ContractIPaymentCoordinator.CallOpts, leaf)
}

// CalculateEarnerLeafHash is a free data retrieval call binding the contract method 0x149bc872.
//
// Solidity: function calculateEarnerLeafHash((address,bytes32) leaf) pure returns(bytes32)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorCallerSession) CalculateEarnerLeafHash(leaf IPaymentCoordinatorEarnerTreeMerkleLeaf) ([32]byte, error) {
	return _ContractIPaymentCoordinator.Contract.CalculateEarnerLeafHash(&_ContractIPaymentCoordinator.CallOpts, leaf)
}

// CalculateTokenLeafHash is a free data retrieval call binding the contract method 0xf8cd8448.
//
// Solidity: function calculateTokenLeafHash((address,uint256) leaf) pure returns(bytes32)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorCaller) CalculateTokenLeafHash(opts *bind.CallOpts, leaf IPaymentCoordinatorTokenTreeMerkleLeaf) ([32]byte, error) {
	var out []interface{}
	err := _ContractIPaymentCoordinator.contract.Call(opts, &out, "calculateTokenLeafHash", leaf)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateTokenLeafHash is a free data retrieval call binding the contract method 0xf8cd8448.
//
// Solidity: function calculateTokenLeafHash((address,uint256) leaf) pure returns(bytes32)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorSession) CalculateTokenLeafHash(leaf IPaymentCoordinatorTokenTreeMerkleLeaf) ([32]byte, error) {
	return _ContractIPaymentCoordinator.Contract.CalculateTokenLeafHash(&_ContractIPaymentCoordinator.CallOpts, leaf)
}

// CalculateTokenLeafHash is a free data retrieval call binding the contract method 0xf8cd8448.
//
// Solidity: function calculateTokenLeafHash((address,uint256) leaf) pure returns(bytes32)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorCallerSession) CalculateTokenLeafHash(leaf IPaymentCoordinatorTokenTreeMerkleLeaf) ([32]byte, error) {
	return _ContractIPaymentCoordinator.Contract.CalculateTokenLeafHash(&_ContractIPaymentCoordinator.CallOpts, leaf)
}

// CalculationIntervalSeconds is a free data retrieval call binding the contract method 0x169bde2b.
//
// Solidity: function calculationIntervalSeconds() view returns(uint64)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorCaller) CalculationIntervalSeconds(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _ContractIPaymentCoordinator.contract.Call(opts, &out, "calculationIntervalSeconds")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// CalculationIntervalSeconds is a free data retrieval call binding the contract method 0x169bde2b.
//
// Solidity: function calculationIntervalSeconds() view returns(uint64)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorSession) CalculationIntervalSeconds() (uint64, error) {
	return _ContractIPaymentCoordinator.Contract.CalculationIntervalSeconds(&_ContractIPaymentCoordinator.CallOpts)
}

// CalculationIntervalSeconds is a free data retrieval call binding the contract method 0x169bde2b.
//
// Solidity: function calculationIntervalSeconds() view returns(uint64)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorCallerSession) CalculationIntervalSeconds() (uint64, error) {
	return _ContractIPaymentCoordinator.Contract.CalculationIntervalSeconds(&_ContractIPaymentCoordinator.CallOpts)
}

// CheckClaim is a free data retrieval call binding the contract method 0x5e9d8348.
//
// Solidity: function checkClaim((uint32,uint32,bytes,(address,bytes32),uint32[],bytes[],(address,uint256)[]) claim) view returns(bool)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorCaller) CheckClaim(opts *bind.CallOpts, claim IPaymentCoordinatorPaymentMerkleClaim) (bool, error) {
	var out []interface{}
	err := _ContractIPaymentCoordinator.contract.Call(opts, &out, "checkClaim", claim)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckClaim is a free data retrieval call binding the contract method 0x5e9d8348.
//
// Solidity: function checkClaim((uint32,uint32,bytes,(address,bytes32),uint32[],bytes[],(address,uint256)[]) claim) view returns(bool)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorSession) CheckClaim(claim IPaymentCoordinatorPaymentMerkleClaim) (bool, error) {
	return _ContractIPaymentCoordinator.Contract.CheckClaim(&_ContractIPaymentCoordinator.CallOpts, claim)
}

// CheckClaim is a free data retrieval call binding the contract method 0x5e9d8348.
//
// Solidity: function checkClaim((uint32,uint32,bytes,(address,bytes32),uint32[],bytes[],(address,uint256)[]) claim) view returns(bool)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorCallerSession) CheckClaim(claim IPaymentCoordinatorPaymentMerkleClaim) (bool, error) {
	return _ContractIPaymentCoordinator.Contract.CheckClaim(&_ContractIPaymentCoordinator.CallOpts, claim)
}

// ClaimerFor is a free data retrieval call binding the contract method 0x2b9f64a4.
//
// Solidity: function claimerFor(address earner) view returns(address)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorCaller) ClaimerFor(opts *bind.CallOpts, earner common.Address) (common.Address, error) {
	var out []interface{}
	err := _ContractIPaymentCoordinator.contract.Call(opts, &out, "claimerFor", earner)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ClaimerFor is a free data retrieval call binding the contract method 0x2b9f64a4.
//
// Solidity: function claimerFor(address earner) view returns(address)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorSession) ClaimerFor(earner common.Address) (common.Address, error) {
	return _ContractIPaymentCoordinator.Contract.ClaimerFor(&_ContractIPaymentCoordinator.CallOpts, earner)
}

// ClaimerFor is a free data retrieval call binding the contract method 0x2b9f64a4.
//
// Solidity: function claimerFor(address earner) view returns(address)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorCallerSession) ClaimerFor(earner common.Address) (common.Address, error) {
	return _ContractIPaymentCoordinator.Contract.ClaimerFor(&_ContractIPaymentCoordinator.CallOpts, earner)
}

// CumulativeClaimed is a free data retrieval call binding the contract method 0x865c6953.
//
// Solidity: function cumulativeClaimed(address claimer, address token) view returns(uint256)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorCaller) CumulativeClaimed(opts *bind.CallOpts, claimer common.Address, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ContractIPaymentCoordinator.contract.Call(opts, &out, "cumulativeClaimed", claimer, token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CumulativeClaimed is a free data retrieval call binding the contract method 0x865c6953.
//
// Solidity: function cumulativeClaimed(address claimer, address token) view returns(uint256)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorSession) CumulativeClaimed(claimer common.Address, token common.Address) (*big.Int, error) {
	return _ContractIPaymentCoordinator.Contract.CumulativeClaimed(&_ContractIPaymentCoordinator.CallOpts, claimer, token)
}

// CumulativeClaimed is a free data retrieval call binding the contract method 0x865c6953.
//
// Solidity: function cumulativeClaimed(address claimer, address token) view returns(uint256)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorCallerSession) CumulativeClaimed(claimer common.Address, token common.Address) (*big.Int, error) {
	return _ContractIPaymentCoordinator.Contract.CumulativeClaimed(&_ContractIPaymentCoordinator.CallOpts, claimer, token)
}

// CurrPaymentCalculationEndTimestamp is a free data retrieval call binding the contract method 0x67ef8585.
//
// Solidity: function currPaymentCalculationEndTimestamp() view returns(uint64)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorCaller) CurrPaymentCalculationEndTimestamp(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _ContractIPaymentCoordinator.contract.Call(opts, &out, "currPaymentCalculationEndTimestamp")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// CurrPaymentCalculationEndTimestamp is a free data retrieval call binding the contract method 0x67ef8585.
//
// Solidity: function currPaymentCalculationEndTimestamp() view returns(uint64)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorSession) CurrPaymentCalculationEndTimestamp() (uint64, error) {
	return _ContractIPaymentCoordinator.Contract.CurrPaymentCalculationEndTimestamp(&_ContractIPaymentCoordinator.CallOpts)
}

// CurrPaymentCalculationEndTimestamp is a free data retrieval call binding the contract method 0x67ef8585.
//
// Solidity: function currPaymentCalculationEndTimestamp() view returns(uint64)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorCallerSession) CurrPaymentCalculationEndTimestamp() (uint64, error) {
	return _ContractIPaymentCoordinator.Contract.CurrPaymentCalculationEndTimestamp(&_ContractIPaymentCoordinator.CallOpts)
}

// GetRootIndexFromHash is a free data retrieval call binding the contract method 0xe810ce21.
//
// Solidity: function getRootIndexFromHash(bytes32 rootHash) view returns(uint32)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorCaller) GetRootIndexFromHash(opts *bind.CallOpts, rootHash [32]byte) (uint32, error) {
	var out []interface{}
	err := _ContractIPaymentCoordinator.contract.Call(opts, &out, "getRootIndexFromHash", rootHash)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetRootIndexFromHash is a free data retrieval call binding the contract method 0xe810ce21.
//
// Solidity: function getRootIndexFromHash(bytes32 rootHash) view returns(uint32)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorSession) GetRootIndexFromHash(rootHash [32]byte) (uint32, error) {
	return _ContractIPaymentCoordinator.Contract.GetRootIndexFromHash(&_ContractIPaymentCoordinator.CallOpts, rootHash)
}

// GetRootIndexFromHash is a free data retrieval call binding the contract method 0xe810ce21.
//
// Solidity: function getRootIndexFromHash(bytes32 rootHash) view returns(uint32)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorCallerSession) GetRootIndexFromHash(rootHash [32]byte) (uint32, error) {
	return _ContractIPaymentCoordinator.Contract.GetRootIndexFromHash(&_ContractIPaymentCoordinator.CallOpts, rootHash)
}

// GlobalOperatorCommissionBips is a free data retrieval call binding the contract method 0x092db007.
//
// Solidity: function globalOperatorCommissionBips() view returns(uint16)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorCaller) GlobalOperatorCommissionBips(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _ContractIPaymentCoordinator.contract.Call(opts, &out, "globalOperatorCommissionBips")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GlobalOperatorCommissionBips is a free data retrieval call binding the contract method 0x092db007.
//
// Solidity: function globalOperatorCommissionBips() view returns(uint16)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorSession) GlobalOperatorCommissionBips() (uint16, error) {
	return _ContractIPaymentCoordinator.Contract.GlobalOperatorCommissionBips(&_ContractIPaymentCoordinator.CallOpts)
}

// GlobalOperatorCommissionBips is a free data retrieval call binding the contract method 0x092db007.
//
// Solidity: function globalOperatorCommissionBips() view returns(uint16)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorCallerSession) GlobalOperatorCommissionBips() (uint16, error) {
	return _ContractIPaymentCoordinator.Contract.GlobalOperatorCommissionBips(&_ContractIPaymentCoordinator.CallOpts)
}

// PaymentUpdater is a free data retrieval call binding the contract method 0x66d3b16b.
//
// Solidity: function paymentUpdater() view returns(address)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorCaller) PaymentUpdater(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIPaymentCoordinator.contract.Call(opts, &out, "paymentUpdater")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PaymentUpdater is a free data retrieval call binding the contract method 0x66d3b16b.
//
// Solidity: function paymentUpdater() view returns(address)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorSession) PaymentUpdater() (common.Address, error) {
	return _ContractIPaymentCoordinator.Contract.PaymentUpdater(&_ContractIPaymentCoordinator.CallOpts)
}

// PaymentUpdater is a free data retrieval call binding the contract method 0x66d3b16b.
//
// Solidity: function paymentUpdater() view returns(address)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorCallerSession) PaymentUpdater() (common.Address, error) {
	return _ContractIPaymentCoordinator.Contract.PaymentUpdater(&_ContractIPaymentCoordinator.CallOpts)
}

// PayAllForRange is a paid mutator transaction binding the contract method 0x1f525c0a.
//
// Solidity: function payAllForRange(((address,uint96)[],address,uint256,uint64,uint64)[] rangePayment) returns()
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorTransactor) PayAllForRange(opts *bind.TransactOpts, rangePayment []IPaymentCoordinatorRangePayment) (*types.Transaction, error) {
	return _ContractIPaymentCoordinator.contract.Transact(opts, "payAllForRange", rangePayment)
}

// PayAllForRange is a paid mutator transaction binding the contract method 0x1f525c0a.
//
// Solidity: function payAllForRange(((address,uint96)[],address,uint256,uint64,uint64)[] rangePayment) returns()
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorSession) PayAllForRange(rangePayment []IPaymentCoordinatorRangePayment) (*types.Transaction, error) {
	return _ContractIPaymentCoordinator.Contract.PayAllForRange(&_ContractIPaymentCoordinator.TransactOpts, rangePayment)
}

// PayAllForRange is a paid mutator transaction binding the contract method 0x1f525c0a.
//
// Solidity: function payAllForRange(((address,uint96)[],address,uint256,uint64,uint64)[] rangePayment) returns()
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorTransactorSession) PayAllForRange(rangePayment []IPaymentCoordinatorRangePayment) (*types.Transaction, error) {
	return _ContractIPaymentCoordinator.Contract.PayAllForRange(&_ContractIPaymentCoordinator.TransactOpts, rangePayment)
}

// PayForRange is a paid mutator transaction binding the contract method 0x42b5c010.
//
// Solidity: function payForRange(((address,uint96)[],address,uint256,uint64,uint64)[] rangePayments) returns()
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorTransactor) PayForRange(opts *bind.TransactOpts, rangePayments []IPaymentCoordinatorRangePayment) (*types.Transaction, error) {
	return _ContractIPaymentCoordinator.contract.Transact(opts, "payForRange", rangePayments)
}

// PayForRange is a paid mutator transaction binding the contract method 0x42b5c010.
//
// Solidity: function payForRange(((address,uint96)[],address,uint256,uint64,uint64)[] rangePayments) returns()
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorSession) PayForRange(rangePayments []IPaymentCoordinatorRangePayment) (*types.Transaction, error) {
	return _ContractIPaymentCoordinator.Contract.PayForRange(&_ContractIPaymentCoordinator.TransactOpts, rangePayments)
}

// PayForRange is a paid mutator transaction binding the contract method 0x42b5c010.
//
// Solidity: function payForRange(((address,uint96)[],address,uint256,uint64,uint64)[] rangePayments) returns()
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorTransactorSession) PayForRange(rangePayments []IPaymentCoordinatorRangePayment) (*types.Transaction, error) {
	return _ContractIPaymentCoordinator.Contract.PayForRange(&_ContractIPaymentCoordinator.TransactOpts, rangePayments)
}

// ProcessClaim is a paid mutator transaction binding the contract method 0x160bc83f.
//
// Solidity: function processClaim((uint32,uint32,bytes,(address,bytes32),uint32[],bytes[],(address,uint256)[]) claim) returns()
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorTransactor) ProcessClaim(opts *bind.TransactOpts, claim IPaymentCoordinatorPaymentMerkleClaim) (*types.Transaction, error) {
	return _ContractIPaymentCoordinator.contract.Transact(opts, "processClaim", claim)
}

// ProcessClaim is a paid mutator transaction binding the contract method 0x160bc83f.
//
// Solidity: function processClaim((uint32,uint32,bytes,(address,bytes32),uint32[],bytes[],(address,uint256)[]) claim) returns()
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorSession) ProcessClaim(claim IPaymentCoordinatorPaymentMerkleClaim) (*types.Transaction, error) {
	return _ContractIPaymentCoordinator.Contract.ProcessClaim(&_ContractIPaymentCoordinator.TransactOpts, claim)
}

// ProcessClaim is a paid mutator transaction binding the contract method 0x160bc83f.
//
// Solidity: function processClaim((uint32,uint32,bytes,(address,bytes32),uint32[],bytes[],(address,uint256)[]) claim) returns()
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorTransactorSession) ProcessClaim(claim IPaymentCoordinatorPaymentMerkleClaim) (*types.Transaction, error) {
	return _ContractIPaymentCoordinator.Contract.ProcessClaim(&_ContractIPaymentCoordinator.TransactOpts, claim)
}

// SetActivationDelay is a paid mutator transaction binding the contract method 0x96b896eb.
//
// Solidity: function setActivationDelay(uint64 _activationDelay) returns()
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorTransactor) SetActivationDelay(opts *bind.TransactOpts, _activationDelay uint64) (*types.Transaction, error) {
	return _ContractIPaymentCoordinator.contract.Transact(opts, "setActivationDelay", _activationDelay)
}

// SetActivationDelay is a paid mutator transaction binding the contract method 0x96b896eb.
//
// Solidity: function setActivationDelay(uint64 _activationDelay) returns()
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorSession) SetActivationDelay(_activationDelay uint64) (*types.Transaction, error) {
	return _ContractIPaymentCoordinator.Contract.SetActivationDelay(&_ContractIPaymentCoordinator.TransactOpts, _activationDelay)
}

// SetActivationDelay is a paid mutator transaction binding the contract method 0x96b896eb.
//
// Solidity: function setActivationDelay(uint64 _activationDelay) returns()
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorTransactorSession) SetActivationDelay(_activationDelay uint64) (*types.Transaction, error) {
	return _ContractIPaymentCoordinator.Contract.SetActivationDelay(&_ContractIPaymentCoordinator.TransactOpts, _activationDelay)
}

// SetClaimerFor is a paid mutator transaction binding the contract method 0xa0169ddd.
//
// Solidity: function setClaimerFor(address claimer) returns()
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorTransactor) SetClaimerFor(opts *bind.TransactOpts, claimer common.Address) (*types.Transaction, error) {
	return _ContractIPaymentCoordinator.contract.Transact(opts, "setClaimerFor", claimer)
}

// SetClaimerFor is a paid mutator transaction binding the contract method 0xa0169ddd.
//
// Solidity: function setClaimerFor(address claimer) returns()
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorSession) SetClaimerFor(claimer common.Address) (*types.Transaction, error) {
	return _ContractIPaymentCoordinator.Contract.SetClaimerFor(&_ContractIPaymentCoordinator.TransactOpts, claimer)
}

// SetClaimerFor is a paid mutator transaction binding the contract method 0xa0169ddd.
//
// Solidity: function setClaimerFor(address claimer) returns()
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorTransactorSession) SetClaimerFor(claimer common.Address) (*types.Transaction, error) {
	return _ContractIPaymentCoordinator.Contract.SetClaimerFor(&_ContractIPaymentCoordinator.TransactOpts, claimer)
}

// SetGlobalOperatorCommission is a paid mutator transaction binding the contract method 0xe221b245.
//
// Solidity: function setGlobalOperatorCommission(uint16 _globalCommissionBips) returns()
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorTransactor) SetGlobalOperatorCommission(opts *bind.TransactOpts, _globalCommissionBips uint16) (*types.Transaction, error) {
	return _ContractIPaymentCoordinator.contract.Transact(opts, "setGlobalOperatorCommission", _globalCommissionBips)
}

// SetGlobalOperatorCommission is a paid mutator transaction binding the contract method 0xe221b245.
//
// Solidity: function setGlobalOperatorCommission(uint16 _globalCommissionBips) returns()
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorSession) SetGlobalOperatorCommission(_globalCommissionBips uint16) (*types.Transaction, error) {
	return _ContractIPaymentCoordinator.Contract.SetGlobalOperatorCommission(&_ContractIPaymentCoordinator.TransactOpts, _globalCommissionBips)
}

// SetGlobalOperatorCommission is a paid mutator transaction binding the contract method 0xe221b245.
//
// Solidity: function setGlobalOperatorCommission(uint16 _globalCommissionBips) returns()
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorTransactorSession) SetGlobalOperatorCommission(_globalCommissionBips uint16) (*types.Transaction, error) {
	return _ContractIPaymentCoordinator.Contract.SetGlobalOperatorCommission(&_ContractIPaymentCoordinator.TransactOpts, _globalCommissionBips)
}

// SetPaymentUpdater is a paid mutator transaction binding the contract method 0x18190f53.
//
// Solidity: function setPaymentUpdater(address _paymentUpdater) returns()
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorTransactor) SetPaymentUpdater(opts *bind.TransactOpts, _paymentUpdater common.Address) (*types.Transaction, error) {
	return _ContractIPaymentCoordinator.contract.Transact(opts, "setPaymentUpdater", _paymentUpdater)
}

// SetPaymentUpdater is a paid mutator transaction binding the contract method 0x18190f53.
//
// Solidity: function setPaymentUpdater(address _paymentUpdater) returns()
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorSession) SetPaymentUpdater(_paymentUpdater common.Address) (*types.Transaction, error) {
	return _ContractIPaymentCoordinator.Contract.SetPaymentUpdater(&_ContractIPaymentCoordinator.TransactOpts, _paymentUpdater)
}

// SetPaymentUpdater is a paid mutator transaction binding the contract method 0x18190f53.
//
// Solidity: function setPaymentUpdater(address _paymentUpdater) returns()
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorTransactorSession) SetPaymentUpdater(_paymentUpdater common.Address) (*types.Transaction, error) {
	return _ContractIPaymentCoordinator.Contract.SetPaymentUpdater(&_ContractIPaymentCoordinator.TransactOpts, _paymentUpdater)
}

// SubmitRoot is a paid mutator transaction binding the contract method 0xa323aa31.
//
// Solidity: function submitRoot(bytes32 root, uint64 paymentCalculationEndTimestamp) returns()
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorTransactor) SubmitRoot(opts *bind.TransactOpts, root [32]byte, paymentCalculationEndTimestamp uint64) (*types.Transaction, error) {
	return _ContractIPaymentCoordinator.contract.Transact(opts, "submitRoot", root, paymentCalculationEndTimestamp)
}

// SubmitRoot is a paid mutator transaction binding the contract method 0xa323aa31.
//
// Solidity: function submitRoot(bytes32 root, uint64 paymentCalculationEndTimestamp) returns()
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorSession) SubmitRoot(root [32]byte, paymentCalculationEndTimestamp uint64) (*types.Transaction, error) {
	return _ContractIPaymentCoordinator.Contract.SubmitRoot(&_ContractIPaymentCoordinator.TransactOpts, root, paymentCalculationEndTimestamp)
}

// SubmitRoot is a paid mutator transaction binding the contract method 0xa323aa31.
//
// Solidity: function submitRoot(bytes32 root, uint64 paymentCalculationEndTimestamp) returns()
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorTransactorSession) SubmitRoot(root [32]byte, paymentCalculationEndTimestamp uint64) (*types.Transaction, error) {
	return _ContractIPaymentCoordinator.Contract.SubmitRoot(&_ContractIPaymentCoordinator.TransactOpts, root, paymentCalculationEndTimestamp)
}

// ContractIPaymentCoordinatorActivationDelaySetIterator is returned from FilterActivationDelaySet and is used to iterate over the raw logs and unpacked data for ActivationDelaySet events raised by the ContractIPaymentCoordinator contract.
type ContractIPaymentCoordinatorActivationDelaySetIterator struct {
	Event *ContractIPaymentCoordinatorActivationDelaySet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractIPaymentCoordinatorActivationDelaySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIPaymentCoordinatorActivationDelaySet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractIPaymentCoordinatorActivationDelaySet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractIPaymentCoordinatorActivationDelaySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIPaymentCoordinatorActivationDelaySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIPaymentCoordinatorActivationDelaySet represents a ActivationDelaySet event raised by the ContractIPaymentCoordinator contract.
type ContractIPaymentCoordinatorActivationDelaySet struct {
	OldActivationDelay uint64
	NewActivationDelay uint64
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterActivationDelaySet is a free log retrieval operation binding the contract event 0x29fabdd5e710c2b0282fcc042164a54b68bd783f561a1f82033b1270e5541dbb.
//
// Solidity: event ActivationDelaySet(uint64 oldActivationDelay, uint64 newActivationDelay)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorFilterer) FilterActivationDelaySet(opts *bind.FilterOpts) (*ContractIPaymentCoordinatorActivationDelaySetIterator, error) {

	logs, sub, err := _ContractIPaymentCoordinator.contract.FilterLogs(opts, "ActivationDelaySet")
	if err != nil {
		return nil, err
	}
	return &ContractIPaymentCoordinatorActivationDelaySetIterator{contract: _ContractIPaymentCoordinator.contract, event: "ActivationDelaySet", logs: logs, sub: sub}, nil
}

// WatchActivationDelaySet is a free log subscription operation binding the contract event 0x29fabdd5e710c2b0282fcc042164a54b68bd783f561a1f82033b1270e5541dbb.
//
// Solidity: event ActivationDelaySet(uint64 oldActivationDelay, uint64 newActivationDelay)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorFilterer) WatchActivationDelaySet(opts *bind.WatchOpts, sink chan<- *ContractIPaymentCoordinatorActivationDelaySet) (event.Subscription, error) {

	logs, sub, err := _ContractIPaymentCoordinator.contract.WatchLogs(opts, "ActivationDelaySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIPaymentCoordinatorActivationDelaySet)
				if err := _ContractIPaymentCoordinator.contract.UnpackLog(event, "ActivationDelaySet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseActivationDelaySet is a log parse operation binding the contract event 0x29fabdd5e710c2b0282fcc042164a54b68bd783f561a1f82033b1270e5541dbb.
//
// Solidity: event ActivationDelaySet(uint64 oldActivationDelay, uint64 newActivationDelay)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorFilterer) ParseActivationDelaySet(log types.Log) (*ContractIPaymentCoordinatorActivationDelaySet, error) {
	event := new(ContractIPaymentCoordinatorActivationDelaySet)
	if err := _ContractIPaymentCoordinator.contract.UnpackLog(event, "ActivationDelaySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIPaymentCoordinatorCalculationIntervalSecondsSetIterator is returned from FilterCalculationIntervalSecondsSet and is used to iterate over the raw logs and unpacked data for CalculationIntervalSecondsSet events raised by the ContractIPaymentCoordinator contract.
type ContractIPaymentCoordinatorCalculationIntervalSecondsSetIterator struct {
	Event *ContractIPaymentCoordinatorCalculationIntervalSecondsSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractIPaymentCoordinatorCalculationIntervalSecondsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIPaymentCoordinatorCalculationIntervalSecondsSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractIPaymentCoordinatorCalculationIntervalSecondsSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractIPaymentCoordinatorCalculationIntervalSecondsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIPaymentCoordinatorCalculationIntervalSecondsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIPaymentCoordinatorCalculationIntervalSecondsSet represents a CalculationIntervalSecondsSet event raised by the ContractIPaymentCoordinator contract.
type ContractIPaymentCoordinatorCalculationIntervalSecondsSet struct {
	OldCalculationIntervalSeconds uint64
	NewCalculationIntervalSeconds uint64
	Raw                           types.Log // Blockchain specific contextual infos
}

// FilterCalculationIntervalSecondsSet is a free log retrieval operation binding the contract event 0x247db8a3d2cd6fa855756642ee37072dc401a228c0c25bde1fa452a2e4739b3c.
//
// Solidity: event CalculationIntervalSecondsSet(uint64 oldCalculationIntervalSeconds, uint64 newCalculationIntervalSeconds)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorFilterer) FilterCalculationIntervalSecondsSet(opts *bind.FilterOpts) (*ContractIPaymentCoordinatorCalculationIntervalSecondsSetIterator, error) {

	logs, sub, err := _ContractIPaymentCoordinator.contract.FilterLogs(opts, "CalculationIntervalSecondsSet")
	if err != nil {
		return nil, err
	}
	return &ContractIPaymentCoordinatorCalculationIntervalSecondsSetIterator{contract: _ContractIPaymentCoordinator.contract, event: "CalculationIntervalSecondsSet", logs: logs, sub: sub}, nil
}

// WatchCalculationIntervalSecondsSet is a free log subscription operation binding the contract event 0x247db8a3d2cd6fa855756642ee37072dc401a228c0c25bde1fa452a2e4739b3c.
//
// Solidity: event CalculationIntervalSecondsSet(uint64 oldCalculationIntervalSeconds, uint64 newCalculationIntervalSeconds)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorFilterer) WatchCalculationIntervalSecondsSet(opts *bind.WatchOpts, sink chan<- *ContractIPaymentCoordinatorCalculationIntervalSecondsSet) (event.Subscription, error) {

	logs, sub, err := _ContractIPaymentCoordinator.contract.WatchLogs(opts, "CalculationIntervalSecondsSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIPaymentCoordinatorCalculationIntervalSecondsSet)
				if err := _ContractIPaymentCoordinator.contract.UnpackLog(event, "CalculationIntervalSecondsSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCalculationIntervalSecondsSet is a log parse operation binding the contract event 0x247db8a3d2cd6fa855756642ee37072dc401a228c0c25bde1fa452a2e4739b3c.
//
// Solidity: event CalculationIntervalSecondsSet(uint64 oldCalculationIntervalSeconds, uint64 newCalculationIntervalSeconds)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorFilterer) ParseCalculationIntervalSecondsSet(log types.Log) (*ContractIPaymentCoordinatorCalculationIntervalSecondsSet, error) {
	event := new(ContractIPaymentCoordinatorCalculationIntervalSecondsSet)
	if err := _ContractIPaymentCoordinator.contract.UnpackLog(event, "CalculationIntervalSecondsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIPaymentCoordinatorClaimerForSetIterator is returned from FilterClaimerForSet and is used to iterate over the raw logs and unpacked data for ClaimerForSet events raised by the ContractIPaymentCoordinator contract.
type ContractIPaymentCoordinatorClaimerForSetIterator struct {
	Event *ContractIPaymentCoordinatorClaimerForSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractIPaymentCoordinatorClaimerForSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIPaymentCoordinatorClaimerForSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractIPaymentCoordinatorClaimerForSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractIPaymentCoordinatorClaimerForSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIPaymentCoordinatorClaimerForSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIPaymentCoordinatorClaimerForSet represents a ClaimerForSet event raised by the ContractIPaymentCoordinator contract.
type ContractIPaymentCoordinatorClaimerForSet struct {
	Earner     common.Address
	OldClaimer common.Address
	Claimer    common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterClaimerForSet is a free log retrieval operation binding the contract event 0xbab947934d42e0ad206f25c9cab18b5bb6ae144acfb00f40b4e3aa59590ca312.
//
// Solidity: event ClaimerForSet(address indexed earner, address indexed oldClaimer, address indexed claimer)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorFilterer) FilterClaimerForSet(opts *bind.FilterOpts, earner []common.Address, oldClaimer []common.Address, claimer []common.Address) (*ContractIPaymentCoordinatorClaimerForSetIterator, error) {

	var earnerRule []interface{}
	for _, earnerItem := range earner {
		earnerRule = append(earnerRule, earnerItem)
	}
	var oldClaimerRule []interface{}
	for _, oldClaimerItem := range oldClaimer {
		oldClaimerRule = append(oldClaimerRule, oldClaimerItem)
	}
	var claimerRule []interface{}
	for _, claimerItem := range claimer {
		claimerRule = append(claimerRule, claimerItem)
	}

	logs, sub, err := _ContractIPaymentCoordinator.contract.FilterLogs(opts, "ClaimerForSet", earnerRule, oldClaimerRule, claimerRule)
	if err != nil {
		return nil, err
	}
	return &ContractIPaymentCoordinatorClaimerForSetIterator{contract: _ContractIPaymentCoordinator.contract, event: "ClaimerForSet", logs: logs, sub: sub}, nil
}

// WatchClaimerForSet is a free log subscription operation binding the contract event 0xbab947934d42e0ad206f25c9cab18b5bb6ae144acfb00f40b4e3aa59590ca312.
//
// Solidity: event ClaimerForSet(address indexed earner, address indexed oldClaimer, address indexed claimer)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorFilterer) WatchClaimerForSet(opts *bind.WatchOpts, sink chan<- *ContractIPaymentCoordinatorClaimerForSet, earner []common.Address, oldClaimer []common.Address, claimer []common.Address) (event.Subscription, error) {

	var earnerRule []interface{}
	for _, earnerItem := range earner {
		earnerRule = append(earnerRule, earnerItem)
	}
	var oldClaimerRule []interface{}
	for _, oldClaimerItem := range oldClaimer {
		oldClaimerRule = append(oldClaimerRule, oldClaimerItem)
	}
	var claimerRule []interface{}
	for _, claimerItem := range claimer {
		claimerRule = append(claimerRule, claimerItem)
	}

	logs, sub, err := _ContractIPaymentCoordinator.contract.WatchLogs(opts, "ClaimerForSet", earnerRule, oldClaimerRule, claimerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIPaymentCoordinatorClaimerForSet)
				if err := _ContractIPaymentCoordinator.contract.UnpackLog(event, "ClaimerForSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseClaimerForSet is a log parse operation binding the contract event 0xbab947934d42e0ad206f25c9cab18b5bb6ae144acfb00f40b4e3aa59590ca312.
//
// Solidity: event ClaimerForSet(address indexed earner, address indexed oldClaimer, address indexed claimer)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorFilterer) ParseClaimerForSet(log types.Log) (*ContractIPaymentCoordinatorClaimerForSet, error) {
	event := new(ContractIPaymentCoordinatorClaimerForSet)
	if err := _ContractIPaymentCoordinator.contract.UnpackLog(event, "ClaimerForSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIPaymentCoordinatorDistributionRootSubmittedIterator is returned from FilterDistributionRootSubmitted and is used to iterate over the raw logs and unpacked data for DistributionRootSubmitted events raised by the ContractIPaymentCoordinator contract.
type ContractIPaymentCoordinatorDistributionRootSubmittedIterator struct {
	Event *ContractIPaymentCoordinatorDistributionRootSubmitted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractIPaymentCoordinatorDistributionRootSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIPaymentCoordinatorDistributionRootSubmitted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractIPaymentCoordinatorDistributionRootSubmitted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractIPaymentCoordinatorDistributionRootSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIPaymentCoordinatorDistributionRootSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIPaymentCoordinatorDistributionRootSubmitted represents a DistributionRootSubmitted event raised by the ContractIPaymentCoordinator contract.
type ContractIPaymentCoordinatorDistributionRootSubmitted struct {
	RootIndex                      uint32
	Root                           [32]byte
	PaymentCalculationEndTimestamp uint64
	ActivatedAt                    uint64
	Raw                            types.Log // Blockchain specific contextual infos
}

// FilterDistributionRootSubmitted is a free log retrieval operation binding the contract event 0xe761fd6052ea06161bb50ec7768c4ae09bb59c30b15332240be2433c25ee12e4.
//
// Solidity: event DistributionRootSubmitted(uint32 indexed rootIndex, bytes32 indexed root, uint64 paymentCalculationEndTimestamp, uint64 activatedAt)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorFilterer) FilterDistributionRootSubmitted(opts *bind.FilterOpts, rootIndex []uint32, root [][32]byte) (*ContractIPaymentCoordinatorDistributionRootSubmittedIterator, error) {

	var rootIndexRule []interface{}
	for _, rootIndexItem := range rootIndex {
		rootIndexRule = append(rootIndexRule, rootIndexItem)
	}
	var rootRule []interface{}
	for _, rootItem := range root {
		rootRule = append(rootRule, rootItem)
	}

	logs, sub, err := _ContractIPaymentCoordinator.contract.FilterLogs(opts, "DistributionRootSubmitted", rootIndexRule, rootRule)
	if err != nil {
		return nil, err
	}
	return &ContractIPaymentCoordinatorDistributionRootSubmittedIterator{contract: _ContractIPaymentCoordinator.contract, event: "DistributionRootSubmitted", logs: logs, sub: sub}, nil
}

// WatchDistributionRootSubmitted is a free log subscription operation binding the contract event 0xe761fd6052ea06161bb50ec7768c4ae09bb59c30b15332240be2433c25ee12e4.
//
// Solidity: event DistributionRootSubmitted(uint32 indexed rootIndex, bytes32 indexed root, uint64 paymentCalculationEndTimestamp, uint64 activatedAt)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorFilterer) WatchDistributionRootSubmitted(opts *bind.WatchOpts, sink chan<- *ContractIPaymentCoordinatorDistributionRootSubmitted, rootIndex []uint32, root [][32]byte) (event.Subscription, error) {

	var rootIndexRule []interface{}
	for _, rootIndexItem := range rootIndex {
		rootIndexRule = append(rootIndexRule, rootIndexItem)
	}
	var rootRule []interface{}
	for _, rootItem := range root {
		rootRule = append(rootRule, rootItem)
	}

	logs, sub, err := _ContractIPaymentCoordinator.contract.WatchLogs(opts, "DistributionRootSubmitted", rootIndexRule, rootRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIPaymentCoordinatorDistributionRootSubmitted)
				if err := _ContractIPaymentCoordinator.contract.UnpackLog(event, "DistributionRootSubmitted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDistributionRootSubmitted is a log parse operation binding the contract event 0xe761fd6052ea06161bb50ec7768c4ae09bb59c30b15332240be2433c25ee12e4.
//
// Solidity: event DistributionRootSubmitted(uint32 indexed rootIndex, bytes32 indexed root, uint64 paymentCalculationEndTimestamp, uint64 activatedAt)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorFilterer) ParseDistributionRootSubmitted(log types.Log) (*ContractIPaymentCoordinatorDistributionRootSubmitted, error) {
	event := new(ContractIPaymentCoordinatorDistributionRootSubmitted)
	if err := _ContractIPaymentCoordinator.contract.UnpackLog(event, "DistributionRootSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIPaymentCoordinatorGlobalCommissionBipsSetIterator is returned from FilterGlobalCommissionBipsSet and is used to iterate over the raw logs and unpacked data for GlobalCommissionBipsSet events raised by the ContractIPaymentCoordinator contract.
type ContractIPaymentCoordinatorGlobalCommissionBipsSetIterator struct {
	Event *ContractIPaymentCoordinatorGlobalCommissionBipsSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractIPaymentCoordinatorGlobalCommissionBipsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIPaymentCoordinatorGlobalCommissionBipsSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractIPaymentCoordinatorGlobalCommissionBipsSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractIPaymentCoordinatorGlobalCommissionBipsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIPaymentCoordinatorGlobalCommissionBipsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIPaymentCoordinatorGlobalCommissionBipsSet represents a GlobalCommissionBipsSet event raised by the ContractIPaymentCoordinator contract.
type ContractIPaymentCoordinatorGlobalCommissionBipsSet struct {
	OldGlobalCommissionBips uint16
	NewGlobalCommissionBips uint16
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterGlobalCommissionBipsSet is a free log retrieval operation binding the contract event 0x8cdc428b0431b82d1619763f443a48197db344ba96905f3949643acd1c863a06.
//
// Solidity: event GlobalCommissionBipsSet(uint16 oldGlobalCommissionBips, uint16 newGlobalCommissionBips)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorFilterer) FilterGlobalCommissionBipsSet(opts *bind.FilterOpts) (*ContractIPaymentCoordinatorGlobalCommissionBipsSetIterator, error) {

	logs, sub, err := _ContractIPaymentCoordinator.contract.FilterLogs(opts, "GlobalCommissionBipsSet")
	if err != nil {
		return nil, err
	}
	return &ContractIPaymentCoordinatorGlobalCommissionBipsSetIterator{contract: _ContractIPaymentCoordinator.contract, event: "GlobalCommissionBipsSet", logs: logs, sub: sub}, nil
}

// WatchGlobalCommissionBipsSet is a free log subscription operation binding the contract event 0x8cdc428b0431b82d1619763f443a48197db344ba96905f3949643acd1c863a06.
//
// Solidity: event GlobalCommissionBipsSet(uint16 oldGlobalCommissionBips, uint16 newGlobalCommissionBips)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorFilterer) WatchGlobalCommissionBipsSet(opts *bind.WatchOpts, sink chan<- *ContractIPaymentCoordinatorGlobalCommissionBipsSet) (event.Subscription, error) {

	logs, sub, err := _ContractIPaymentCoordinator.contract.WatchLogs(opts, "GlobalCommissionBipsSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIPaymentCoordinatorGlobalCommissionBipsSet)
				if err := _ContractIPaymentCoordinator.contract.UnpackLog(event, "GlobalCommissionBipsSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseGlobalCommissionBipsSet is a log parse operation binding the contract event 0x8cdc428b0431b82d1619763f443a48197db344ba96905f3949643acd1c863a06.
//
// Solidity: event GlobalCommissionBipsSet(uint16 oldGlobalCommissionBips, uint16 newGlobalCommissionBips)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorFilterer) ParseGlobalCommissionBipsSet(log types.Log) (*ContractIPaymentCoordinatorGlobalCommissionBipsSet, error) {
	event := new(ContractIPaymentCoordinatorGlobalCommissionBipsSet)
	if err := _ContractIPaymentCoordinator.contract.UnpackLog(event, "GlobalCommissionBipsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIPaymentCoordinatorPayAllForRangeSubmitterSetIterator is returned from FilterPayAllForRangeSubmitterSet and is used to iterate over the raw logs and unpacked data for PayAllForRangeSubmitterSet events raised by the ContractIPaymentCoordinator contract.
type ContractIPaymentCoordinatorPayAllForRangeSubmitterSetIterator struct {
	Event *ContractIPaymentCoordinatorPayAllForRangeSubmitterSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractIPaymentCoordinatorPayAllForRangeSubmitterSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIPaymentCoordinatorPayAllForRangeSubmitterSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractIPaymentCoordinatorPayAllForRangeSubmitterSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractIPaymentCoordinatorPayAllForRangeSubmitterSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIPaymentCoordinatorPayAllForRangeSubmitterSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIPaymentCoordinatorPayAllForRangeSubmitterSet represents a PayAllForRangeSubmitterSet event raised by the ContractIPaymentCoordinator contract.
type ContractIPaymentCoordinatorPayAllForRangeSubmitterSet struct {
	PayAllForRangeSubmitter common.Address
	OldValue                bool
	NewValue                bool
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterPayAllForRangeSubmitterSet is a free log retrieval operation binding the contract event 0x17b06be02af2803593116eb96121b9c6e8bee1cc1b145e7c31c19c180e86189b.
//
// Solidity: event PayAllForRangeSubmitterSet(address indexed payAllForRangeSubmitter, bool indexed oldValue, bool indexed newValue)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorFilterer) FilterPayAllForRangeSubmitterSet(opts *bind.FilterOpts, payAllForRangeSubmitter []common.Address, oldValue []bool, newValue []bool) (*ContractIPaymentCoordinatorPayAllForRangeSubmitterSetIterator, error) {

	var payAllForRangeSubmitterRule []interface{}
	for _, payAllForRangeSubmitterItem := range payAllForRangeSubmitter {
		payAllForRangeSubmitterRule = append(payAllForRangeSubmitterRule, payAllForRangeSubmitterItem)
	}
	var oldValueRule []interface{}
	for _, oldValueItem := range oldValue {
		oldValueRule = append(oldValueRule, oldValueItem)
	}
	var newValueRule []interface{}
	for _, newValueItem := range newValue {
		newValueRule = append(newValueRule, newValueItem)
	}

	logs, sub, err := _ContractIPaymentCoordinator.contract.FilterLogs(opts, "PayAllForRangeSubmitterSet", payAllForRangeSubmitterRule, oldValueRule, newValueRule)
	if err != nil {
		return nil, err
	}
	return &ContractIPaymentCoordinatorPayAllForRangeSubmitterSetIterator{contract: _ContractIPaymentCoordinator.contract, event: "PayAllForRangeSubmitterSet", logs: logs, sub: sub}, nil
}

// WatchPayAllForRangeSubmitterSet is a free log subscription operation binding the contract event 0x17b06be02af2803593116eb96121b9c6e8bee1cc1b145e7c31c19c180e86189b.
//
// Solidity: event PayAllForRangeSubmitterSet(address indexed payAllForRangeSubmitter, bool indexed oldValue, bool indexed newValue)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorFilterer) WatchPayAllForRangeSubmitterSet(opts *bind.WatchOpts, sink chan<- *ContractIPaymentCoordinatorPayAllForRangeSubmitterSet, payAllForRangeSubmitter []common.Address, oldValue []bool, newValue []bool) (event.Subscription, error) {

	var payAllForRangeSubmitterRule []interface{}
	for _, payAllForRangeSubmitterItem := range payAllForRangeSubmitter {
		payAllForRangeSubmitterRule = append(payAllForRangeSubmitterRule, payAllForRangeSubmitterItem)
	}
	var oldValueRule []interface{}
	for _, oldValueItem := range oldValue {
		oldValueRule = append(oldValueRule, oldValueItem)
	}
	var newValueRule []interface{}
	for _, newValueItem := range newValue {
		newValueRule = append(newValueRule, newValueItem)
	}

	logs, sub, err := _ContractIPaymentCoordinator.contract.WatchLogs(opts, "PayAllForRangeSubmitterSet", payAllForRangeSubmitterRule, oldValueRule, newValueRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIPaymentCoordinatorPayAllForRangeSubmitterSet)
				if err := _ContractIPaymentCoordinator.contract.UnpackLog(event, "PayAllForRangeSubmitterSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePayAllForRangeSubmitterSet is a log parse operation binding the contract event 0x17b06be02af2803593116eb96121b9c6e8bee1cc1b145e7c31c19c180e86189b.
//
// Solidity: event PayAllForRangeSubmitterSet(address indexed payAllForRangeSubmitter, bool indexed oldValue, bool indexed newValue)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorFilterer) ParsePayAllForRangeSubmitterSet(log types.Log) (*ContractIPaymentCoordinatorPayAllForRangeSubmitterSet, error) {
	event := new(ContractIPaymentCoordinatorPayAllForRangeSubmitterSet)
	if err := _ContractIPaymentCoordinator.contract.UnpackLog(event, "PayAllForRangeSubmitterSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIPaymentCoordinatorPaymentClaimedIterator is returned from FilterPaymentClaimed and is used to iterate over the raw logs and unpacked data for PaymentClaimed events raised by the ContractIPaymentCoordinator contract.
type ContractIPaymentCoordinatorPaymentClaimedIterator struct {
	Event *ContractIPaymentCoordinatorPaymentClaimed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractIPaymentCoordinatorPaymentClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIPaymentCoordinatorPaymentClaimed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractIPaymentCoordinatorPaymentClaimed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractIPaymentCoordinatorPaymentClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIPaymentCoordinatorPaymentClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIPaymentCoordinatorPaymentClaimed represents a PaymentClaimed event raised by the ContractIPaymentCoordinator contract.
type ContractIPaymentCoordinatorPaymentClaimed struct {
	Root [32]byte
	Leaf IPaymentCoordinatorTokenTreeMerkleLeaf
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterPaymentClaimed is a free log retrieval operation binding the contract event 0x65c3ec778ea0d5ecd945eb3ccba33dbb4950974e8adb43029e234656486ccb35.
//
// Solidity: event PaymentClaimed(bytes32 indexed root, (address,uint256) leaf)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorFilterer) FilterPaymentClaimed(opts *bind.FilterOpts, root [][32]byte) (*ContractIPaymentCoordinatorPaymentClaimedIterator, error) {

	var rootRule []interface{}
	for _, rootItem := range root {
		rootRule = append(rootRule, rootItem)
	}

	logs, sub, err := _ContractIPaymentCoordinator.contract.FilterLogs(opts, "PaymentClaimed", rootRule)
	if err != nil {
		return nil, err
	}
	return &ContractIPaymentCoordinatorPaymentClaimedIterator{contract: _ContractIPaymentCoordinator.contract, event: "PaymentClaimed", logs: logs, sub: sub}, nil
}

// WatchPaymentClaimed is a free log subscription operation binding the contract event 0x65c3ec778ea0d5ecd945eb3ccba33dbb4950974e8adb43029e234656486ccb35.
//
// Solidity: event PaymentClaimed(bytes32 indexed root, (address,uint256) leaf)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorFilterer) WatchPaymentClaimed(opts *bind.WatchOpts, sink chan<- *ContractIPaymentCoordinatorPaymentClaimed, root [][32]byte) (event.Subscription, error) {

	var rootRule []interface{}
	for _, rootItem := range root {
		rootRule = append(rootRule, rootItem)
	}

	logs, sub, err := _ContractIPaymentCoordinator.contract.WatchLogs(opts, "PaymentClaimed", rootRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIPaymentCoordinatorPaymentClaimed)
				if err := _ContractIPaymentCoordinator.contract.UnpackLog(event, "PaymentClaimed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaymentClaimed is a log parse operation binding the contract event 0x65c3ec778ea0d5ecd945eb3ccba33dbb4950974e8adb43029e234656486ccb35.
//
// Solidity: event PaymentClaimed(bytes32 indexed root, (address,uint256) leaf)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorFilterer) ParsePaymentClaimed(log types.Log) (*ContractIPaymentCoordinatorPaymentClaimed, error) {
	event := new(ContractIPaymentCoordinatorPaymentClaimed)
	if err := _ContractIPaymentCoordinator.contract.UnpackLog(event, "PaymentClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIPaymentCoordinatorPaymentUpdaterSetIterator is returned from FilterPaymentUpdaterSet and is used to iterate over the raw logs and unpacked data for PaymentUpdaterSet events raised by the ContractIPaymentCoordinator contract.
type ContractIPaymentCoordinatorPaymentUpdaterSetIterator struct {
	Event *ContractIPaymentCoordinatorPaymentUpdaterSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractIPaymentCoordinatorPaymentUpdaterSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIPaymentCoordinatorPaymentUpdaterSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractIPaymentCoordinatorPaymentUpdaterSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractIPaymentCoordinatorPaymentUpdaterSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIPaymentCoordinatorPaymentUpdaterSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIPaymentCoordinatorPaymentUpdaterSet represents a PaymentUpdaterSet event raised by the ContractIPaymentCoordinator contract.
type ContractIPaymentCoordinatorPaymentUpdaterSet struct {
	OldPaymentUpdater common.Address
	NewPaymentUpdater common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterPaymentUpdaterSet is a free log retrieval operation binding the contract event 0x07d2890b3eb1206e7c3cb6bf8d46da31385ace3ce99abf85e5b690c83aa49678.
//
// Solidity: event PaymentUpdaterSet(address indexed oldPaymentUpdater, address indexed newPaymentUpdater)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorFilterer) FilterPaymentUpdaterSet(opts *bind.FilterOpts, oldPaymentUpdater []common.Address, newPaymentUpdater []common.Address) (*ContractIPaymentCoordinatorPaymentUpdaterSetIterator, error) {

	var oldPaymentUpdaterRule []interface{}
	for _, oldPaymentUpdaterItem := range oldPaymentUpdater {
		oldPaymentUpdaterRule = append(oldPaymentUpdaterRule, oldPaymentUpdaterItem)
	}
	var newPaymentUpdaterRule []interface{}
	for _, newPaymentUpdaterItem := range newPaymentUpdater {
		newPaymentUpdaterRule = append(newPaymentUpdaterRule, newPaymentUpdaterItem)
	}

	logs, sub, err := _ContractIPaymentCoordinator.contract.FilterLogs(opts, "PaymentUpdaterSet", oldPaymentUpdaterRule, newPaymentUpdaterRule)
	if err != nil {
		return nil, err
	}
	return &ContractIPaymentCoordinatorPaymentUpdaterSetIterator{contract: _ContractIPaymentCoordinator.contract, event: "PaymentUpdaterSet", logs: logs, sub: sub}, nil
}

// WatchPaymentUpdaterSet is a free log subscription operation binding the contract event 0x07d2890b3eb1206e7c3cb6bf8d46da31385ace3ce99abf85e5b690c83aa49678.
//
// Solidity: event PaymentUpdaterSet(address indexed oldPaymentUpdater, address indexed newPaymentUpdater)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorFilterer) WatchPaymentUpdaterSet(opts *bind.WatchOpts, sink chan<- *ContractIPaymentCoordinatorPaymentUpdaterSet, oldPaymentUpdater []common.Address, newPaymentUpdater []common.Address) (event.Subscription, error) {

	var oldPaymentUpdaterRule []interface{}
	for _, oldPaymentUpdaterItem := range oldPaymentUpdater {
		oldPaymentUpdaterRule = append(oldPaymentUpdaterRule, oldPaymentUpdaterItem)
	}
	var newPaymentUpdaterRule []interface{}
	for _, newPaymentUpdaterItem := range newPaymentUpdater {
		newPaymentUpdaterRule = append(newPaymentUpdaterRule, newPaymentUpdaterItem)
	}

	logs, sub, err := _ContractIPaymentCoordinator.contract.WatchLogs(opts, "PaymentUpdaterSet", oldPaymentUpdaterRule, newPaymentUpdaterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIPaymentCoordinatorPaymentUpdaterSet)
				if err := _ContractIPaymentCoordinator.contract.UnpackLog(event, "PaymentUpdaterSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaymentUpdaterSet is a log parse operation binding the contract event 0x07d2890b3eb1206e7c3cb6bf8d46da31385ace3ce99abf85e5b690c83aa49678.
//
// Solidity: event PaymentUpdaterSet(address indexed oldPaymentUpdater, address indexed newPaymentUpdater)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorFilterer) ParsePaymentUpdaterSet(log types.Log) (*ContractIPaymentCoordinatorPaymentUpdaterSet, error) {
	event := new(ContractIPaymentCoordinatorPaymentUpdaterSet)
	if err := _ContractIPaymentCoordinator.contract.UnpackLog(event, "PaymentUpdaterSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIPaymentCoordinatorRangePaymentCreatedIterator is returned from FilterRangePaymentCreated and is used to iterate over the raw logs and unpacked data for RangePaymentCreated events raised by the ContractIPaymentCoordinator contract.
type ContractIPaymentCoordinatorRangePaymentCreatedIterator struct {
	Event *ContractIPaymentCoordinatorRangePaymentCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractIPaymentCoordinatorRangePaymentCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIPaymentCoordinatorRangePaymentCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractIPaymentCoordinatorRangePaymentCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractIPaymentCoordinatorRangePaymentCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIPaymentCoordinatorRangePaymentCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIPaymentCoordinatorRangePaymentCreated represents a RangePaymentCreated event raised by the ContractIPaymentCoordinator contract.
type ContractIPaymentCoordinatorRangePaymentCreated struct {
	Avs              common.Address
	PaymentNonce     *big.Int
	RangePaymentHash [32]byte
	RangePayment     IPaymentCoordinatorRangePayment
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterRangePaymentCreated is a free log retrieval operation binding the contract event 0x6caa185f3cbe2749467bfbe7df9a7674bb65d8a78c5626b5464f880fa4bac48f.
//
// Solidity: event RangePaymentCreated(address indexed avs, uint256 indexed paymentNonce, bytes32 indexed rangePaymentHash, ((address,uint96)[],address,uint256,uint64,uint64) rangePayment)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorFilterer) FilterRangePaymentCreated(opts *bind.FilterOpts, avs []common.Address, paymentNonce []*big.Int, rangePaymentHash [][32]byte) (*ContractIPaymentCoordinatorRangePaymentCreatedIterator, error) {

	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}
	var paymentNonceRule []interface{}
	for _, paymentNonceItem := range paymentNonce {
		paymentNonceRule = append(paymentNonceRule, paymentNonceItem)
	}
	var rangePaymentHashRule []interface{}
	for _, rangePaymentHashItem := range rangePaymentHash {
		rangePaymentHashRule = append(rangePaymentHashRule, rangePaymentHashItem)
	}

	logs, sub, err := _ContractIPaymentCoordinator.contract.FilterLogs(opts, "RangePaymentCreated", avsRule, paymentNonceRule, rangePaymentHashRule)
	if err != nil {
		return nil, err
	}
	return &ContractIPaymentCoordinatorRangePaymentCreatedIterator{contract: _ContractIPaymentCoordinator.contract, event: "RangePaymentCreated", logs: logs, sub: sub}, nil
}

// WatchRangePaymentCreated is a free log subscription operation binding the contract event 0x6caa185f3cbe2749467bfbe7df9a7674bb65d8a78c5626b5464f880fa4bac48f.
//
// Solidity: event RangePaymentCreated(address indexed avs, uint256 indexed paymentNonce, bytes32 indexed rangePaymentHash, ((address,uint96)[],address,uint256,uint64,uint64) rangePayment)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorFilterer) WatchRangePaymentCreated(opts *bind.WatchOpts, sink chan<- *ContractIPaymentCoordinatorRangePaymentCreated, avs []common.Address, paymentNonce []*big.Int, rangePaymentHash [][32]byte) (event.Subscription, error) {

	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}
	var paymentNonceRule []interface{}
	for _, paymentNonceItem := range paymentNonce {
		paymentNonceRule = append(paymentNonceRule, paymentNonceItem)
	}
	var rangePaymentHashRule []interface{}
	for _, rangePaymentHashItem := range rangePaymentHash {
		rangePaymentHashRule = append(rangePaymentHashRule, rangePaymentHashItem)
	}

	logs, sub, err := _ContractIPaymentCoordinator.contract.WatchLogs(opts, "RangePaymentCreated", avsRule, paymentNonceRule, rangePaymentHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIPaymentCoordinatorRangePaymentCreated)
				if err := _ContractIPaymentCoordinator.contract.UnpackLog(event, "RangePaymentCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRangePaymentCreated is a log parse operation binding the contract event 0x6caa185f3cbe2749467bfbe7df9a7674bb65d8a78c5626b5464f880fa4bac48f.
//
// Solidity: event RangePaymentCreated(address indexed avs, uint256 indexed paymentNonce, bytes32 indexed rangePaymentHash, ((address,uint96)[],address,uint256,uint64,uint64) rangePayment)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorFilterer) ParseRangePaymentCreated(log types.Log) (*ContractIPaymentCoordinatorRangePaymentCreated, error) {
	event := new(ContractIPaymentCoordinatorRangePaymentCreated)
	if err := _ContractIPaymentCoordinator.contract.UnpackLog(event, "RangePaymentCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIPaymentCoordinatorRangePaymentForAllCreatedIterator is returned from FilterRangePaymentForAllCreated and is used to iterate over the raw logs and unpacked data for RangePaymentForAllCreated events raised by the ContractIPaymentCoordinator contract.
type ContractIPaymentCoordinatorRangePaymentForAllCreatedIterator struct {
	Event *ContractIPaymentCoordinatorRangePaymentForAllCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractIPaymentCoordinatorRangePaymentForAllCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIPaymentCoordinatorRangePaymentForAllCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractIPaymentCoordinatorRangePaymentForAllCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractIPaymentCoordinatorRangePaymentForAllCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIPaymentCoordinatorRangePaymentForAllCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIPaymentCoordinatorRangePaymentForAllCreated represents a RangePaymentForAllCreated event raised by the ContractIPaymentCoordinator contract.
type ContractIPaymentCoordinatorRangePaymentForAllCreated struct {
	Submitter        common.Address
	PaymentNonce     *big.Int
	RangePaymentHash [32]byte
	RangePayment     IPaymentCoordinatorRangePayment
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterRangePaymentForAllCreated is a free log retrieval operation binding the contract event 0xb4139d2f31d7b13944e26e3e632ccd1f8f4845ba446870f93f0ae0fda89e1eeb.
//
// Solidity: event RangePaymentForAllCreated(address indexed submitter, uint256 indexed paymentNonce, bytes32 indexed rangePaymentHash, ((address,uint96)[],address,uint256,uint64,uint64) rangePayment)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorFilterer) FilterRangePaymentForAllCreated(opts *bind.FilterOpts, submitter []common.Address, paymentNonce []*big.Int, rangePaymentHash [][32]byte) (*ContractIPaymentCoordinatorRangePaymentForAllCreatedIterator, error) {

	var submitterRule []interface{}
	for _, submitterItem := range submitter {
		submitterRule = append(submitterRule, submitterItem)
	}
	var paymentNonceRule []interface{}
	for _, paymentNonceItem := range paymentNonce {
		paymentNonceRule = append(paymentNonceRule, paymentNonceItem)
	}
	var rangePaymentHashRule []interface{}
	for _, rangePaymentHashItem := range rangePaymentHash {
		rangePaymentHashRule = append(rangePaymentHashRule, rangePaymentHashItem)
	}

	logs, sub, err := _ContractIPaymentCoordinator.contract.FilterLogs(opts, "RangePaymentForAllCreated", submitterRule, paymentNonceRule, rangePaymentHashRule)
	if err != nil {
		return nil, err
	}
	return &ContractIPaymentCoordinatorRangePaymentForAllCreatedIterator{contract: _ContractIPaymentCoordinator.contract, event: "RangePaymentForAllCreated", logs: logs, sub: sub}, nil
}

// WatchRangePaymentForAllCreated is a free log subscription operation binding the contract event 0xb4139d2f31d7b13944e26e3e632ccd1f8f4845ba446870f93f0ae0fda89e1eeb.
//
// Solidity: event RangePaymentForAllCreated(address indexed submitter, uint256 indexed paymentNonce, bytes32 indexed rangePaymentHash, ((address,uint96)[],address,uint256,uint64,uint64) rangePayment)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorFilterer) WatchRangePaymentForAllCreated(opts *bind.WatchOpts, sink chan<- *ContractIPaymentCoordinatorRangePaymentForAllCreated, submitter []common.Address, paymentNonce []*big.Int, rangePaymentHash [][32]byte) (event.Subscription, error) {

	var submitterRule []interface{}
	for _, submitterItem := range submitter {
		submitterRule = append(submitterRule, submitterItem)
	}
	var paymentNonceRule []interface{}
	for _, paymentNonceItem := range paymentNonce {
		paymentNonceRule = append(paymentNonceRule, paymentNonceItem)
	}
	var rangePaymentHashRule []interface{}
	for _, rangePaymentHashItem := range rangePaymentHash {
		rangePaymentHashRule = append(rangePaymentHashRule, rangePaymentHashItem)
	}

	logs, sub, err := _ContractIPaymentCoordinator.contract.WatchLogs(opts, "RangePaymentForAllCreated", submitterRule, paymentNonceRule, rangePaymentHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIPaymentCoordinatorRangePaymentForAllCreated)
				if err := _ContractIPaymentCoordinator.contract.UnpackLog(event, "RangePaymentForAllCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRangePaymentForAllCreated is a log parse operation binding the contract event 0xb4139d2f31d7b13944e26e3e632ccd1f8f4845ba446870f93f0ae0fda89e1eeb.
//
// Solidity: event RangePaymentForAllCreated(address indexed submitter, uint256 indexed paymentNonce, bytes32 indexed rangePaymentHash, ((address,uint96)[],address,uint256,uint64,uint64) rangePayment)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorFilterer) ParseRangePaymentForAllCreated(log types.Log) (*ContractIPaymentCoordinatorRangePaymentForAllCreated, error) {
	event := new(ContractIPaymentCoordinatorRangePaymentForAllCreated)
	if err := _ContractIPaymentCoordinator.contract.UnpackLog(event, "RangePaymentForAllCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
