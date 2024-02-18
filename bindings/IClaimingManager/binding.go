// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractIClaimingManager

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

// IClaimingManagerPaymentMerkleClaim is an auto generated low-level Go binding around an user-defined struct.
type IClaimingManagerPaymentMerkleClaim struct {
	Token     common.Address
	Amount    *big.Int
	RootIndex uint32
	LeafIndex uint32
	Proof     []byte
}

// ContractIClaimingManagerMetaData contains all meta data concerning the ContractIClaimingManager contract.
var ContractIClaimingManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"activationDelay\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claimers\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"cumulativeClaimed\",\"inputs\":[{\"name\":\"claimer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"globalCommissionBips\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_paymentUpdater\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_activationDelay\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paymentUpdater\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"processClaims\",\"inputs\":[{\"name\":\"claims\",\"type\":\"tuple[]\",\"internalType\":\"structIClaimingManager.PaymentMerkleClaim[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"rootIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"leafIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setActivationDelay\",\"inputs\":[{\"name\":\"_activationDelay\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setClaimer\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"claimer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setGlobalCommission\",\"inputs\":[{\"name\":\"_globalCommissionBips\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPaymentUpdater\",\"inputs\":[{\"name\":\"_paymentUpdater\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitRoot\",\"inputs\":[{\"name\":\"root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"paymentsCalculatedUntilTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ActivationDelaySet\",\"inputs\":[{\"name\":\"oldActivationDelay\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"newActivationDelay\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ClaimerSet\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"claimer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GlobalCommissionBipsSet\",\"inputs\":[{\"name\":\"oldGlobalCommissionBips\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"newGlobalCommissionBips\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PaymentClaimed\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIERC20\"},{\"name\":\"claimer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PaymentUpdaterSet\",\"inputs\":[{\"name\":\"oldPaymentUpdater\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newPaymentUpdater\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RootSubmitted\",\"inputs\":[{\"name\":\"root\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"paymentsCalculatedUntilTimestamp\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"activatedAfter\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false}]",
}

// ContractIClaimingManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractIClaimingManagerMetaData.ABI instead.
var ContractIClaimingManagerABI = ContractIClaimingManagerMetaData.ABI

// ContractIClaimingManager is an auto generated Go binding around an Ethereum contract.
type ContractIClaimingManager struct {
	ContractIClaimingManagerCaller     // Read-only binding to the contract
	ContractIClaimingManagerTransactor // Write-only binding to the contract
	ContractIClaimingManagerFilterer   // Log filterer for contract events
}

// ContractIClaimingManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractIClaimingManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractIClaimingManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractIClaimingManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractIClaimingManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractIClaimingManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractIClaimingManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractIClaimingManagerSession struct {
	Contract     *ContractIClaimingManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ContractIClaimingManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractIClaimingManagerCallerSession struct {
	Contract *ContractIClaimingManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// ContractIClaimingManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractIClaimingManagerTransactorSession struct {
	Contract     *ContractIClaimingManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// ContractIClaimingManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractIClaimingManagerRaw struct {
	Contract *ContractIClaimingManager // Generic contract binding to access the raw methods on
}

// ContractIClaimingManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractIClaimingManagerCallerRaw struct {
	Contract *ContractIClaimingManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ContractIClaimingManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractIClaimingManagerTransactorRaw struct {
	Contract *ContractIClaimingManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractIClaimingManager creates a new instance of ContractIClaimingManager, bound to a specific deployed contract.
func NewContractIClaimingManager(address common.Address, backend bind.ContractBackend) (*ContractIClaimingManager, error) {
	contract, err := bindContractIClaimingManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractIClaimingManager{ContractIClaimingManagerCaller: ContractIClaimingManagerCaller{contract: contract}, ContractIClaimingManagerTransactor: ContractIClaimingManagerTransactor{contract: contract}, ContractIClaimingManagerFilterer: ContractIClaimingManagerFilterer{contract: contract}}, nil
}

// NewContractIClaimingManagerCaller creates a new read-only instance of ContractIClaimingManager, bound to a specific deployed contract.
func NewContractIClaimingManagerCaller(address common.Address, caller bind.ContractCaller) (*ContractIClaimingManagerCaller, error) {
	contract, err := bindContractIClaimingManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractIClaimingManagerCaller{contract: contract}, nil
}

// NewContractIClaimingManagerTransactor creates a new write-only instance of ContractIClaimingManager, bound to a specific deployed contract.
func NewContractIClaimingManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractIClaimingManagerTransactor, error) {
	contract, err := bindContractIClaimingManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractIClaimingManagerTransactor{contract: contract}, nil
}

// NewContractIClaimingManagerFilterer creates a new log filterer instance of ContractIClaimingManager, bound to a specific deployed contract.
func NewContractIClaimingManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractIClaimingManagerFilterer, error) {
	contract, err := bindContractIClaimingManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractIClaimingManagerFilterer{contract: contract}, nil
}

// bindContractIClaimingManager binds a generic wrapper to an already deployed contract.
func bindContractIClaimingManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractIClaimingManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractIClaimingManager *ContractIClaimingManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractIClaimingManager.Contract.ContractIClaimingManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractIClaimingManager *ContractIClaimingManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractIClaimingManager.Contract.ContractIClaimingManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractIClaimingManager *ContractIClaimingManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractIClaimingManager.Contract.ContractIClaimingManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractIClaimingManager *ContractIClaimingManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractIClaimingManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractIClaimingManager *ContractIClaimingManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractIClaimingManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractIClaimingManager *ContractIClaimingManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractIClaimingManager.Contract.contract.Transact(opts, method, params...)
}

// ActivationDelay is a free data retrieval call binding the contract method 0x3a8c0786.
//
// Solidity: function activationDelay() view returns(uint32)
func (_ContractIClaimingManager *ContractIClaimingManagerCaller) ActivationDelay(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractIClaimingManager.contract.Call(opts, &out, "activationDelay")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ActivationDelay is a free data retrieval call binding the contract method 0x3a8c0786.
//
// Solidity: function activationDelay() view returns(uint32)
func (_ContractIClaimingManager *ContractIClaimingManagerSession) ActivationDelay() (uint32, error) {
	return _ContractIClaimingManager.Contract.ActivationDelay(&_ContractIClaimingManager.CallOpts)
}

// ActivationDelay is a free data retrieval call binding the contract method 0x3a8c0786.
//
// Solidity: function activationDelay() view returns(uint32)
func (_ContractIClaimingManager *ContractIClaimingManagerCallerSession) ActivationDelay() (uint32, error) {
	return _ContractIClaimingManager.Contract.ActivationDelay(&_ContractIClaimingManager.CallOpts)
}

// Claimers is a free data retrieval call binding the contract method 0xda62fba9.
//
// Solidity: function claimers(address account) view returns(address)
func (_ContractIClaimingManager *ContractIClaimingManagerCaller) Claimers(opts *bind.CallOpts, account common.Address) (common.Address, error) {
	var out []interface{}
	err := _ContractIClaimingManager.contract.Call(opts, &out, "claimers", account)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Claimers is a free data retrieval call binding the contract method 0xda62fba9.
//
// Solidity: function claimers(address account) view returns(address)
func (_ContractIClaimingManager *ContractIClaimingManagerSession) Claimers(account common.Address) (common.Address, error) {
	return _ContractIClaimingManager.Contract.Claimers(&_ContractIClaimingManager.CallOpts, account)
}

// Claimers is a free data retrieval call binding the contract method 0xda62fba9.
//
// Solidity: function claimers(address account) view returns(address)
func (_ContractIClaimingManager *ContractIClaimingManagerCallerSession) Claimers(account common.Address) (common.Address, error) {
	return _ContractIClaimingManager.Contract.Claimers(&_ContractIClaimingManager.CallOpts, account)
}

// CumulativeClaimed is a free data retrieval call binding the contract method 0x865c6953.
//
// Solidity: function cumulativeClaimed(address claimer, address token) view returns(uint256)
func (_ContractIClaimingManager *ContractIClaimingManagerCaller) CumulativeClaimed(opts *bind.CallOpts, claimer common.Address, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ContractIClaimingManager.contract.Call(opts, &out, "cumulativeClaimed", claimer, token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CumulativeClaimed is a free data retrieval call binding the contract method 0x865c6953.
//
// Solidity: function cumulativeClaimed(address claimer, address token) view returns(uint256)
func (_ContractIClaimingManager *ContractIClaimingManagerSession) CumulativeClaimed(claimer common.Address, token common.Address) (*big.Int, error) {
	return _ContractIClaimingManager.Contract.CumulativeClaimed(&_ContractIClaimingManager.CallOpts, claimer, token)
}

// CumulativeClaimed is a free data retrieval call binding the contract method 0x865c6953.
//
// Solidity: function cumulativeClaimed(address claimer, address token) view returns(uint256)
func (_ContractIClaimingManager *ContractIClaimingManagerCallerSession) CumulativeClaimed(claimer common.Address, token common.Address) (*big.Int, error) {
	return _ContractIClaimingManager.Contract.CumulativeClaimed(&_ContractIClaimingManager.CallOpts, claimer, token)
}

// GlobalCommissionBips is a free data retrieval call binding the contract method 0x2c088f0d.
//
// Solidity: function globalCommissionBips() view returns(uint16)
func (_ContractIClaimingManager *ContractIClaimingManagerCaller) GlobalCommissionBips(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _ContractIClaimingManager.contract.Call(opts, &out, "globalCommissionBips")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GlobalCommissionBips is a free data retrieval call binding the contract method 0x2c088f0d.
//
// Solidity: function globalCommissionBips() view returns(uint16)
func (_ContractIClaimingManager *ContractIClaimingManagerSession) GlobalCommissionBips() (uint16, error) {
	return _ContractIClaimingManager.Contract.GlobalCommissionBips(&_ContractIClaimingManager.CallOpts)
}

// GlobalCommissionBips is a free data retrieval call binding the contract method 0x2c088f0d.
//
// Solidity: function globalCommissionBips() view returns(uint16)
func (_ContractIClaimingManager *ContractIClaimingManagerCallerSession) GlobalCommissionBips() (uint16, error) {
	return _ContractIClaimingManager.Contract.GlobalCommissionBips(&_ContractIClaimingManager.CallOpts)
}

// PaymentUpdater is a free data retrieval call binding the contract method 0x66d3b16b.
//
// Solidity: function paymentUpdater() view returns(address)
func (_ContractIClaimingManager *ContractIClaimingManagerCaller) PaymentUpdater(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIClaimingManager.contract.Call(opts, &out, "paymentUpdater")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PaymentUpdater is a free data retrieval call binding the contract method 0x66d3b16b.
//
// Solidity: function paymentUpdater() view returns(address)
func (_ContractIClaimingManager *ContractIClaimingManagerSession) PaymentUpdater() (common.Address, error) {
	return _ContractIClaimingManager.Contract.PaymentUpdater(&_ContractIClaimingManager.CallOpts)
}

// PaymentUpdater is a free data retrieval call binding the contract method 0x66d3b16b.
//
// Solidity: function paymentUpdater() view returns(address)
func (_ContractIClaimingManager *ContractIClaimingManagerCallerSession) PaymentUpdater() (common.Address, error) {
	return _ContractIClaimingManager.Contract.PaymentUpdater(&_ContractIClaimingManager.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x6ecf2b22.
//
// Solidity: function initialize(address initialOwner, address _paymentUpdater, uint32 _activationDelay) returns()
func (_ContractIClaimingManager *ContractIClaimingManagerTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address, _paymentUpdater common.Address, _activationDelay uint32) (*types.Transaction, error) {
	return _ContractIClaimingManager.contract.Transact(opts, "initialize", initialOwner, _paymentUpdater, _activationDelay)
}

// Initialize is a paid mutator transaction binding the contract method 0x6ecf2b22.
//
// Solidity: function initialize(address initialOwner, address _paymentUpdater, uint32 _activationDelay) returns()
func (_ContractIClaimingManager *ContractIClaimingManagerSession) Initialize(initialOwner common.Address, _paymentUpdater common.Address, _activationDelay uint32) (*types.Transaction, error) {
	return _ContractIClaimingManager.Contract.Initialize(&_ContractIClaimingManager.TransactOpts, initialOwner, _paymentUpdater, _activationDelay)
}

// Initialize is a paid mutator transaction binding the contract method 0x6ecf2b22.
//
// Solidity: function initialize(address initialOwner, address _paymentUpdater, uint32 _activationDelay) returns()
func (_ContractIClaimingManager *ContractIClaimingManagerTransactorSession) Initialize(initialOwner common.Address, _paymentUpdater common.Address, _activationDelay uint32) (*types.Transaction, error) {
	return _ContractIClaimingManager.Contract.Initialize(&_ContractIClaimingManager.TransactOpts, initialOwner, _paymentUpdater, _activationDelay)
}

// ProcessClaims is a paid mutator transaction binding the contract method 0x1441f788.
//
// Solidity: function processClaims((address,uint256,uint32,uint32,bytes)[] claims) returns()
func (_ContractIClaimingManager *ContractIClaimingManagerTransactor) ProcessClaims(opts *bind.TransactOpts, claims []IClaimingManagerPaymentMerkleClaim) (*types.Transaction, error) {
	return _ContractIClaimingManager.contract.Transact(opts, "processClaims", claims)
}

// ProcessClaims is a paid mutator transaction binding the contract method 0x1441f788.
//
// Solidity: function processClaims((address,uint256,uint32,uint32,bytes)[] claims) returns()
func (_ContractIClaimingManager *ContractIClaimingManagerSession) ProcessClaims(claims []IClaimingManagerPaymentMerkleClaim) (*types.Transaction, error) {
	return _ContractIClaimingManager.Contract.ProcessClaims(&_ContractIClaimingManager.TransactOpts, claims)
}

// ProcessClaims is a paid mutator transaction binding the contract method 0x1441f788.
//
// Solidity: function processClaims((address,uint256,uint32,uint32,bytes)[] claims) returns()
func (_ContractIClaimingManager *ContractIClaimingManagerTransactorSession) ProcessClaims(claims []IClaimingManagerPaymentMerkleClaim) (*types.Transaction, error) {
	return _ContractIClaimingManager.Contract.ProcessClaims(&_ContractIClaimingManager.TransactOpts, claims)
}

// SetActivationDelay is a paid mutator transaction binding the contract method 0x58baaa3e.
//
// Solidity: function setActivationDelay(uint32 _activationDelay) returns()
func (_ContractIClaimingManager *ContractIClaimingManagerTransactor) SetActivationDelay(opts *bind.TransactOpts, _activationDelay uint32) (*types.Transaction, error) {
	return _ContractIClaimingManager.contract.Transact(opts, "setActivationDelay", _activationDelay)
}

// SetActivationDelay is a paid mutator transaction binding the contract method 0x58baaa3e.
//
// Solidity: function setActivationDelay(uint32 _activationDelay) returns()
func (_ContractIClaimingManager *ContractIClaimingManagerSession) SetActivationDelay(_activationDelay uint32) (*types.Transaction, error) {
	return _ContractIClaimingManager.Contract.SetActivationDelay(&_ContractIClaimingManager.TransactOpts, _activationDelay)
}

// SetActivationDelay is a paid mutator transaction binding the contract method 0x58baaa3e.
//
// Solidity: function setActivationDelay(uint32 _activationDelay) returns()
func (_ContractIClaimingManager *ContractIClaimingManagerTransactorSession) SetActivationDelay(_activationDelay uint32) (*types.Transaction, error) {
	return _ContractIClaimingManager.Contract.SetActivationDelay(&_ContractIClaimingManager.TransactOpts, _activationDelay)
}

// SetClaimer is a paid mutator transaction binding the contract method 0xf5cf673b.
//
// Solidity: function setClaimer(address account, address claimer) returns()
func (_ContractIClaimingManager *ContractIClaimingManagerTransactor) SetClaimer(opts *bind.TransactOpts, account common.Address, claimer common.Address) (*types.Transaction, error) {
	return _ContractIClaimingManager.contract.Transact(opts, "setClaimer", account, claimer)
}

// SetClaimer is a paid mutator transaction binding the contract method 0xf5cf673b.
//
// Solidity: function setClaimer(address account, address claimer) returns()
func (_ContractIClaimingManager *ContractIClaimingManagerSession) SetClaimer(account common.Address, claimer common.Address) (*types.Transaction, error) {
	return _ContractIClaimingManager.Contract.SetClaimer(&_ContractIClaimingManager.TransactOpts, account, claimer)
}

// SetClaimer is a paid mutator transaction binding the contract method 0xf5cf673b.
//
// Solidity: function setClaimer(address account, address claimer) returns()
func (_ContractIClaimingManager *ContractIClaimingManagerTransactorSession) SetClaimer(account common.Address, claimer common.Address) (*types.Transaction, error) {
	return _ContractIClaimingManager.Contract.SetClaimer(&_ContractIClaimingManager.TransactOpts, account, claimer)
}

// SetGlobalCommission is a paid mutator transaction binding the contract method 0x9d284a8a.
//
// Solidity: function setGlobalCommission(uint16 _globalCommissionBips) returns()
func (_ContractIClaimingManager *ContractIClaimingManagerTransactor) SetGlobalCommission(opts *bind.TransactOpts, _globalCommissionBips uint16) (*types.Transaction, error) {
	return _ContractIClaimingManager.contract.Transact(opts, "setGlobalCommission", _globalCommissionBips)
}

// SetGlobalCommission is a paid mutator transaction binding the contract method 0x9d284a8a.
//
// Solidity: function setGlobalCommission(uint16 _globalCommissionBips) returns()
func (_ContractIClaimingManager *ContractIClaimingManagerSession) SetGlobalCommission(_globalCommissionBips uint16) (*types.Transaction, error) {
	return _ContractIClaimingManager.Contract.SetGlobalCommission(&_ContractIClaimingManager.TransactOpts, _globalCommissionBips)
}

// SetGlobalCommission is a paid mutator transaction binding the contract method 0x9d284a8a.
//
// Solidity: function setGlobalCommission(uint16 _globalCommissionBips) returns()
func (_ContractIClaimingManager *ContractIClaimingManagerTransactorSession) SetGlobalCommission(_globalCommissionBips uint16) (*types.Transaction, error) {
	return _ContractIClaimingManager.Contract.SetGlobalCommission(&_ContractIClaimingManager.TransactOpts, _globalCommissionBips)
}

// SetPaymentUpdater is a paid mutator transaction binding the contract method 0x18190f53.
//
// Solidity: function setPaymentUpdater(address _paymentUpdater) returns()
func (_ContractIClaimingManager *ContractIClaimingManagerTransactor) SetPaymentUpdater(opts *bind.TransactOpts, _paymentUpdater common.Address) (*types.Transaction, error) {
	return _ContractIClaimingManager.contract.Transact(opts, "setPaymentUpdater", _paymentUpdater)
}

// SetPaymentUpdater is a paid mutator transaction binding the contract method 0x18190f53.
//
// Solidity: function setPaymentUpdater(address _paymentUpdater) returns()
func (_ContractIClaimingManager *ContractIClaimingManagerSession) SetPaymentUpdater(_paymentUpdater common.Address) (*types.Transaction, error) {
	return _ContractIClaimingManager.Contract.SetPaymentUpdater(&_ContractIClaimingManager.TransactOpts, _paymentUpdater)
}

// SetPaymentUpdater is a paid mutator transaction binding the contract method 0x18190f53.
//
// Solidity: function setPaymentUpdater(address _paymentUpdater) returns()
func (_ContractIClaimingManager *ContractIClaimingManagerTransactorSession) SetPaymentUpdater(_paymentUpdater common.Address) (*types.Transaction, error) {
	return _ContractIClaimingManager.Contract.SetPaymentUpdater(&_ContractIClaimingManager.TransactOpts, _paymentUpdater)
}

// SubmitRoot is a paid mutator transaction binding the contract method 0x3efe1db6.
//
// Solidity: function submitRoot(bytes32 root, uint32 paymentsCalculatedUntilTimestamp) returns()
func (_ContractIClaimingManager *ContractIClaimingManagerTransactor) SubmitRoot(opts *bind.TransactOpts, root [32]byte, paymentsCalculatedUntilTimestamp uint32) (*types.Transaction, error) {
	return _ContractIClaimingManager.contract.Transact(opts, "submitRoot", root, paymentsCalculatedUntilTimestamp)
}

// SubmitRoot is a paid mutator transaction binding the contract method 0x3efe1db6.
//
// Solidity: function submitRoot(bytes32 root, uint32 paymentsCalculatedUntilTimestamp) returns()
func (_ContractIClaimingManager *ContractIClaimingManagerSession) SubmitRoot(root [32]byte, paymentsCalculatedUntilTimestamp uint32) (*types.Transaction, error) {
	return _ContractIClaimingManager.Contract.SubmitRoot(&_ContractIClaimingManager.TransactOpts, root, paymentsCalculatedUntilTimestamp)
}

// SubmitRoot is a paid mutator transaction binding the contract method 0x3efe1db6.
//
// Solidity: function submitRoot(bytes32 root, uint32 paymentsCalculatedUntilTimestamp) returns()
func (_ContractIClaimingManager *ContractIClaimingManagerTransactorSession) SubmitRoot(root [32]byte, paymentsCalculatedUntilTimestamp uint32) (*types.Transaction, error) {
	return _ContractIClaimingManager.Contract.SubmitRoot(&_ContractIClaimingManager.TransactOpts, root, paymentsCalculatedUntilTimestamp)
}

// ContractIClaimingManagerActivationDelaySetIterator is returned from FilterActivationDelaySet and is used to iterate over the raw logs and unpacked data for ActivationDelaySet events raised by the ContractIClaimingManager contract.
type ContractIClaimingManagerActivationDelaySetIterator struct {
	Event *ContractIClaimingManagerActivationDelaySet // Event containing the contract specifics and raw log

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
func (it *ContractIClaimingManagerActivationDelaySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIClaimingManagerActivationDelaySet)
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
		it.Event = new(ContractIClaimingManagerActivationDelaySet)
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
func (it *ContractIClaimingManagerActivationDelaySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIClaimingManagerActivationDelaySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIClaimingManagerActivationDelaySet represents a ActivationDelaySet event raised by the ContractIClaimingManager contract.
type ContractIClaimingManagerActivationDelaySet struct {
	OldActivationDelay uint32
	NewActivationDelay uint32
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterActivationDelaySet is a free log retrieval operation binding the contract event 0xaf557c6c02c208794817a705609cfa935f827312a1adfdd26494b6b95dd2b4b3.
//
// Solidity: event ActivationDelaySet(uint32 oldActivationDelay, uint32 newActivationDelay)
func (_ContractIClaimingManager *ContractIClaimingManagerFilterer) FilterActivationDelaySet(opts *bind.FilterOpts) (*ContractIClaimingManagerActivationDelaySetIterator, error) {

	logs, sub, err := _ContractIClaimingManager.contract.FilterLogs(opts, "ActivationDelaySet")
	if err != nil {
		return nil, err
	}
	return &ContractIClaimingManagerActivationDelaySetIterator{contract: _ContractIClaimingManager.contract, event: "ActivationDelaySet", logs: logs, sub: sub}, nil
}

// WatchActivationDelaySet is a free log subscription operation binding the contract event 0xaf557c6c02c208794817a705609cfa935f827312a1adfdd26494b6b95dd2b4b3.
//
// Solidity: event ActivationDelaySet(uint32 oldActivationDelay, uint32 newActivationDelay)
func (_ContractIClaimingManager *ContractIClaimingManagerFilterer) WatchActivationDelaySet(opts *bind.WatchOpts, sink chan<- *ContractIClaimingManagerActivationDelaySet) (event.Subscription, error) {

	logs, sub, err := _ContractIClaimingManager.contract.WatchLogs(opts, "ActivationDelaySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIClaimingManagerActivationDelaySet)
				if err := _ContractIClaimingManager.contract.UnpackLog(event, "ActivationDelaySet", log); err != nil {
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

// ParseActivationDelaySet is a log parse operation binding the contract event 0xaf557c6c02c208794817a705609cfa935f827312a1adfdd26494b6b95dd2b4b3.
//
// Solidity: event ActivationDelaySet(uint32 oldActivationDelay, uint32 newActivationDelay)
func (_ContractIClaimingManager *ContractIClaimingManagerFilterer) ParseActivationDelaySet(log types.Log) (*ContractIClaimingManagerActivationDelaySet, error) {
	event := new(ContractIClaimingManagerActivationDelaySet)
	if err := _ContractIClaimingManager.contract.UnpackLog(event, "ActivationDelaySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIClaimingManagerClaimerSetIterator is returned from FilterClaimerSet and is used to iterate over the raw logs and unpacked data for ClaimerSet events raised by the ContractIClaimingManager contract.
type ContractIClaimingManagerClaimerSetIterator struct {
	Event *ContractIClaimingManagerClaimerSet // Event containing the contract specifics and raw log

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
func (it *ContractIClaimingManagerClaimerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIClaimingManagerClaimerSet)
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
		it.Event = new(ContractIClaimingManagerClaimerSet)
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
func (it *ContractIClaimingManagerClaimerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIClaimingManagerClaimerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIClaimingManagerClaimerSet represents a ClaimerSet event raised by the ContractIClaimingManager contract.
type ContractIClaimingManagerClaimerSet struct {
	Account common.Address
	Claimer common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterClaimerSet is a free log retrieval operation binding the contract event 0x4925eafc82d0c4d67889898eeed64b18488ab19811e61620f387026dec126a28.
//
// Solidity: event ClaimerSet(address account, address claimer)
func (_ContractIClaimingManager *ContractIClaimingManagerFilterer) FilterClaimerSet(opts *bind.FilterOpts) (*ContractIClaimingManagerClaimerSetIterator, error) {

	logs, sub, err := _ContractIClaimingManager.contract.FilterLogs(opts, "ClaimerSet")
	if err != nil {
		return nil, err
	}
	return &ContractIClaimingManagerClaimerSetIterator{contract: _ContractIClaimingManager.contract, event: "ClaimerSet", logs: logs, sub: sub}, nil
}

// WatchClaimerSet is a free log subscription operation binding the contract event 0x4925eafc82d0c4d67889898eeed64b18488ab19811e61620f387026dec126a28.
//
// Solidity: event ClaimerSet(address account, address claimer)
func (_ContractIClaimingManager *ContractIClaimingManagerFilterer) WatchClaimerSet(opts *bind.WatchOpts, sink chan<- *ContractIClaimingManagerClaimerSet) (event.Subscription, error) {

	logs, sub, err := _ContractIClaimingManager.contract.WatchLogs(opts, "ClaimerSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIClaimingManagerClaimerSet)
				if err := _ContractIClaimingManager.contract.UnpackLog(event, "ClaimerSet", log); err != nil {
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

// ParseClaimerSet is a log parse operation binding the contract event 0x4925eafc82d0c4d67889898eeed64b18488ab19811e61620f387026dec126a28.
//
// Solidity: event ClaimerSet(address account, address claimer)
func (_ContractIClaimingManager *ContractIClaimingManagerFilterer) ParseClaimerSet(log types.Log) (*ContractIClaimingManagerClaimerSet, error) {
	event := new(ContractIClaimingManagerClaimerSet)
	if err := _ContractIClaimingManager.contract.UnpackLog(event, "ClaimerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIClaimingManagerGlobalCommissionBipsSetIterator is returned from FilterGlobalCommissionBipsSet and is used to iterate over the raw logs and unpacked data for GlobalCommissionBipsSet events raised by the ContractIClaimingManager contract.
type ContractIClaimingManagerGlobalCommissionBipsSetIterator struct {
	Event *ContractIClaimingManagerGlobalCommissionBipsSet // Event containing the contract specifics and raw log

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
func (it *ContractIClaimingManagerGlobalCommissionBipsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIClaimingManagerGlobalCommissionBipsSet)
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
		it.Event = new(ContractIClaimingManagerGlobalCommissionBipsSet)
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
func (it *ContractIClaimingManagerGlobalCommissionBipsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIClaimingManagerGlobalCommissionBipsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIClaimingManagerGlobalCommissionBipsSet represents a GlobalCommissionBipsSet event raised by the ContractIClaimingManager contract.
type ContractIClaimingManagerGlobalCommissionBipsSet struct {
	OldGlobalCommissionBips uint16
	NewGlobalCommissionBips uint16
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterGlobalCommissionBipsSet is a free log retrieval operation binding the contract event 0x8cdc428b0431b82d1619763f443a48197db344ba96905f3949643acd1c863a06.
//
// Solidity: event GlobalCommissionBipsSet(uint16 oldGlobalCommissionBips, uint16 newGlobalCommissionBips)
func (_ContractIClaimingManager *ContractIClaimingManagerFilterer) FilterGlobalCommissionBipsSet(opts *bind.FilterOpts) (*ContractIClaimingManagerGlobalCommissionBipsSetIterator, error) {

	logs, sub, err := _ContractIClaimingManager.contract.FilterLogs(opts, "GlobalCommissionBipsSet")
	if err != nil {
		return nil, err
	}
	return &ContractIClaimingManagerGlobalCommissionBipsSetIterator{contract: _ContractIClaimingManager.contract, event: "GlobalCommissionBipsSet", logs: logs, sub: sub}, nil
}

// WatchGlobalCommissionBipsSet is a free log subscription operation binding the contract event 0x8cdc428b0431b82d1619763f443a48197db344ba96905f3949643acd1c863a06.
//
// Solidity: event GlobalCommissionBipsSet(uint16 oldGlobalCommissionBips, uint16 newGlobalCommissionBips)
func (_ContractIClaimingManager *ContractIClaimingManagerFilterer) WatchGlobalCommissionBipsSet(opts *bind.WatchOpts, sink chan<- *ContractIClaimingManagerGlobalCommissionBipsSet) (event.Subscription, error) {

	logs, sub, err := _ContractIClaimingManager.contract.WatchLogs(opts, "GlobalCommissionBipsSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIClaimingManagerGlobalCommissionBipsSet)
				if err := _ContractIClaimingManager.contract.UnpackLog(event, "GlobalCommissionBipsSet", log); err != nil {
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
func (_ContractIClaimingManager *ContractIClaimingManagerFilterer) ParseGlobalCommissionBipsSet(log types.Log) (*ContractIClaimingManagerGlobalCommissionBipsSet, error) {
	event := new(ContractIClaimingManagerGlobalCommissionBipsSet)
	if err := _ContractIClaimingManager.contract.UnpackLog(event, "GlobalCommissionBipsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIClaimingManagerPaymentClaimedIterator is returned from FilterPaymentClaimed and is used to iterate over the raw logs and unpacked data for PaymentClaimed events raised by the ContractIClaimingManager contract.
type ContractIClaimingManagerPaymentClaimedIterator struct {
	Event *ContractIClaimingManagerPaymentClaimed // Event containing the contract specifics and raw log

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
func (it *ContractIClaimingManagerPaymentClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIClaimingManagerPaymentClaimed)
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
		it.Event = new(ContractIClaimingManagerPaymentClaimed)
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
func (it *ContractIClaimingManagerPaymentClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIClaimingManagerPaymentClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIClaimingManagerPaymentClaimed represents a PaymentClaimed event raised by the ContractIClaimingManager contract.
type ContractIClaimingManagerPaymentClaimed struct {
	Token   common.Address
	Claimer common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaymentClaimed is a free log retrieval operation binding the contract event 0x6906788f9c6d5b8d1f449ea40ce9f59b59a825c15753633c28e35595b0a57659.
//
// Solidity: event PaymentClaimed(address token, address claimer, uint256 amount)
func (_ContractIClaimingManager *ContractIClaimingManagerFilterer) FilterPaymentClaimed(opts *bind.FilterOpts) (*ContractIClaimingManagerPaymentClaimedIterator, error) {

	logs, sub, err := _ContractIClaimingManager.contract.FilterLogs(opts, "PaymentClaimed")
	if err != nil {
		return nil, err
	}
	return &ContractIClaimingManagerPaymentClaimedIterator{contract: _ContractIClaimingManager.contract, event: "PaymentClaimed", logs: logs, sub: sub}, nil
}

// WatchPaymentClaimed is a free log subscription operation binding the contract event 0x6906788f9c6d5b8d1f449ea40ce9f59b59a825c15753633c28e35595b0a57659.
//
// Solidity: event PaymentClaimed(address token, address claimer, uint256 amount)
func (_ContractIClaimingManager *ContractIClaimingManagerFilterer) WatchPaymentClaimed(opts *bind.WatchOpts, sink chan<- *ContractIClaimingManagerPaymentClaimed) (event.Subscription, error) {

	logs, sub, err := _ContractIClaimingManager.contract.WatchLogs(opts, "PaymentClaimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIClaimingManagerPaymentClaimed)
				if err := _ContractIClaimingManager.contract.UnpackLog(event, "PaymentClaimed", log); err != nil {
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

// ParsePaymentClaimed is a log parse operation binding the contract event 0x6906788f9c6d5b8d1f449ea40ce9f59b59a825c15753633c28e35595b0a57659.
//
// Solidity: event PaymentClaimed(address token, address claimer, uint256 amount)
func (_ContractIClaimingManager *ContractIClaimingManagerFilterer) ParsePaymentClaimed(log types.Log) (*ContractIClaimingManagerPaymentClaimed, error) {
	event := new(ContractIClaimingManagerPaymentClaimed)
	if err := _ContractIClaimingManager.contract.UnpackLog(event, "PaymentClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIClaimingManagerPaymentUpdaterSetIterator is returned from FilterPaymentUpdaterSet and is used to iterate over the raw logs and unpacked data for PaymentUpdaterSet events raised by the ContractIClaimingManager contract.
type ContractIClaimingManagerPaymentUpdaterSetIterator struct {
	Event *ContractIClaimingManagerPaymentUpdaterSet // Event containing the contract specifics and raw log

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
func (it *ContractIClaimingManagerPaymentUpdaterSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIClaimingManagerPaymentUpdaterSet)
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
		it.Event = new(ContractIClaimingManagerPaymentUpdaterSet)
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
func (it *ContractIClaimingManagerPaymentUpdaterSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIClaimingManagerPaymentUpdaterSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIClaimingManagerPaymentUpdaterSet represents a PaymentUpdaterSet event raised by the ContractIClaimingManager contract.
type ContractIClaimingManagerPaymentUpdaterSet struct {
	OldPaymentUpdater common.Address
	NewPaymentUpdater common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterPaymentUpdaterSet is a free log retrieval operation binding the contract event 0x07d2890b3eb1206e7c3cb6bf8d46da31385ace3ce99abf85e5b690c83aa49678.
//
// Solidity: event PaymentUpdaterSet(address oldPaymentUpdater, address newPaymentUpdater)
func (_ContractIClaimingManager *ContractIClaimingManagerFilterer) FilterPaymentUpdaterSet(opts *bind.FilterOpts) (*ContractIClaimingManagerPaymentUpdaterSetIterator, error) {

	logs, sub, err := _ContractIClaimingManager.contract.FilterLogs(opts, "PaymentUpdaterSet")
	if err != nil {
		return nil, err
	}
	return &ContractIClaimingManagerPaymentUpdaterSetIterator{contract: _ContractIClaimingManager.contract, event: "PaymentUpdaterSet", logs: logs, sub: sub}, nil
}

// WatchPaymentUpdaterSet is a free log subscription operation binding the contract event 0x07d2890b3eb1206e7c3cb6bf8d46da31385ace3ce99abf85e5b690c83aa49678.
//
// Solidity: event PaymentUpdaterSet(address oldPaymentUpdater, address newPaymentUpdater)
func (_ContractIClaimingManager *ContractIClaimingManagerFilterer) WatchPaymentUpdaterSet(opts *bind.WatchOpts, sink chan<- *ContractIClaimingManagerPaymentUpdaterSet) (event.Subscription, error) {

	logs, sub, err := _ContractIClaimingManager.contract.WatchLogs(opts, "PaymentUpdaterSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIClaimingManagerPaymentUpdaterSet)
				if err := _ContractIClaimingManager.contract.UnpackLog(event, "PaymentUpdaterSet", log); err != nil {
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
// Solidity: event PaymentUpdaterSet(address oldPaymentUpdater, address newPaymentUpdater)
func (_ContractIClaimingManager *ContractIClaimingManagerFilterer) ParsePaymentUpdaterSet(log types.Log) (*ContractIClaimingManagerPaymentUpdaterSet, error) {
	event := new(ContractIClaimingManagerPaymentUpdaterSet)
	if err := _ContractIClaimingManager.contract.UnpackLog(event, "PaymentUpdaterSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIClaimingManagerRootSubmittedIterator is returned from FilterRootSubmitted and is used to iterate over the raw logs and unpacked data for RootSubmitted events raised by the ContractIClaimingManager contract.
type ContractIClaimingManagerRootSubmittedIterator struct {
	Event *ContractIClaimingManagerRootSubmitted // Event containing the contract specifics and raw log

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
func (it *ContractIClaimingManagerRootSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIClaimingManagerRootSubmitted)
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
		it.Event = new(ContractIClaimingManagerRootSubmitted)
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
func (it *ContractIClaimingManagerRootSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIClaimingManagerRootSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIClaimingManagerRootSubmitted represents a RootSubmitted event raised by the ContractIClaimingManager contract.
type ContractIClaimingManagerRootSubmitted struct {
	Root                             [32]byte
	PaymentsCalculatedUntilTimestamp uint32
	ActivatedAfter                   uint32
	Raw                              types.Log // Blockchain specific contextual infos
}

// FilterRootSubmitted is a free log retrieval operation binding the contract event 0x262191a0e015e84c4074af7ac4d2305db1490bf60340fbd04afa74cb37bcbdf1.
//
// Solidity: event RootSubmitted(bytes32 root, uint32 paymentsCalculatedUntilTimestamp, uint32 activatedAfter)
func (_ContractIClaimingManager *ContractIClaimingManagerFilterer) FilterRootSubmitted(opts *bind.FilterOpts) (*ContractIClaimingManagerRootSubmittedIterator, error) {

	logs, sub, err := _ContractIClaimingManager.contract.FilterLogs(opts, "RootSubmitted")
	if err != nil {
		return nil, err
	}
	return &ContractIClaimingManagerRootSubmittedIterator{contract: _ContractIClaimingManager.contract, event: "RootSubmitted", logs: logs, sub: sub}, nil
}

// WatchRootSubmitted is a free log subscription operation binding the contract event 0x262191a0e015e84c4074af7ac4d2305db1490bf60340fbd04afa74cb37bcbdf1.
//
// Solidity: event RootSubmitted(bytes32 root, uint32 paymentsCalculatedUntilTimestamp, uint32 activatedAfter)
func (_ContractIClaimingManager *ContractIClaimingManagerFilterer) WatchRootSubmitted(opts *bind.WatchOpts, sink chan<- *ContractIClaimingManagerRootSubmitted) (event.Subscription, error) {

	logs, sub, err := _ContractIClaimingManager.contract.WatchLogs(opts, "RootSubmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIClaimingManagerRootSubmitted)
				if err := _ContractIClaimingManager.contract.UnpackLog(event, "RootSubmitted", log); err != nil {
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

// ParseRootSubmitted is a log parse operation binding the contract event 0x262191a0e015e84c4074af7ac4d2305db1490bf60340fbd04afa74cb37bcbdf1.
//
// Solidity: event RootSubmitted(bytes32 root, uint32 paymentsCalculatedUntilTimestamp, uint32 activatedAfter)
func (_ContractIClaimingManager *ContractIClaimingManagerFilterer) ParseRootSubmitted(log types.Log) (*ContractIClaimingManagerRootSubmitted, error) {
	event := new(ContractIClaimingManagerRootSubmitted)
	if err := _ContractIClaimingManager.contract.UnpackLog(event, "RootSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
