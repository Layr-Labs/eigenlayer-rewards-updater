// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractIStrategyManager

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

// IStrategyManagerDeprecatedStructQueuedWithdrawal is an auto generated low-level Go binding around an user-defined struct.
type IStrategyManagerDeprecatedStructQueuedWithdrawal struct {
	Strategies           []common.Address
	Shares               []*big.Int
	Staker               common.Address
	WithdrawerAndNonce   IStrategyManagerDeprecatedStructWithdrawerAndNonce
	WithdrawalStartBlock uint32
	DelegatedAddress     common.Address
}

// IStrategyManagerDeprecatedStructWithdrawerAndNonce is an auto generated low-level Go binding around an user-defined struct.
type IStrategyManagerDeprecatedStructWithdrawerAndNonce struct {
	Withdrawer common.Address
	Nonce      *big.Int
}

// ContractIStrategyManagerMetaData contains all meta data concerning the ContractIStrategyManager contract.
var ContractIStrategyManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"addShares\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addStrategiesToDepositWhitelist\",\"inputs\":[{\"name\":\"strategiesToWhitelist\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"calculateWithdrawalRoot\",\"inputs\":[{\"name\":\"queuedWithdrawal\",\"type\":\"tuple\",\"internalType\":\"structIStrategyManager.DeprecatedStruct_QueuedWithdrawal\",\"components\":[{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"shares\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"withdrawerAndNonce\",\"type\":\"tuple\",\"internalType\":\"structIStrategyManager.DeprecatedStruct_WithdrawerAndNonce\",\"components\":[{\"name\":\"withdrawer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"withdrawalStartBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"delegatedAddress\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"delegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"depositIntoStrategy\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositIntoStrategyWithSignature\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"eigenPodManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIEigenPodManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDeposits\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"migrateQueuedWithdrawal\",\"inputs\":[{\"name\":\"queuedWithdrawal\",\"type\":\"tuple\",\"internalType\":\"structIStrategyManager.DeprecatedStruct_QueuedWithdrawal\",\"components\":[{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"shares\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"withdrawerAndNonce\",\"type\":\"tuple\",\"internalType\":\"structIStrategyManager.DeprecatedStruct_WithdrawerAndNonce\",\"components\":[{\"name\":\"withdrawer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"withdrawalStartBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"delegatedAddress\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeShares\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeStrategiesFromDepositWhitelist\",\"inputs\":[{\"name\":\"strategiesToRemoveFromWhitelist\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"slasher\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractISlasher\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stakerStrategyListLength\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stakerStrategyShares\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"strategyWhitelister\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawSharesAsTokens\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Deposit\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIERC20\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StrategyAddedToDepositWhitelist\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StrategyRemovedFromDepositWhitelist\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StrategyWhitelisterChanged\",\"inputs\":[{\"name\":\"previousAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false}]",
}

// ContractIStrategyManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractIStrategyManagerMetaData.ABI instead.
var ContractIStrategyManagerABI = ContractIStrategyManagerMetaData.ABI

// ContractIStrategyManager is an auto generated Go binding around an Ethereum contract.
type ContractIStrategyManager struct {
	ContractIStrategyManagerCaller     // Read-only binding to the contract
	ContractIStrategyManagerTransactor // Write-only binding to the contract
	ContractIStrategyManagerFilterer   // Log filterer for contract events
}

// ContractIStrategyManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractIStrategyManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractIStrategyManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractIStrategyManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractIStrategyManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractIStrategyManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractIStrategyManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractIStrategyManagerSession struct {
	Contract     *ContractIStrategyManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ContractIStrategyManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractIStrategyManagerCallerSession struct {
	Contract *ContractIStrategyManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// ContractIStrategyManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractIStrategyManagerTransactorSession struct {
	Contract     *ContractIStrategyManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// ContractIStrategyManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractIStrategyManagerRaw struct {
	Contract *ContractIStrategyManager // Generic contract binding to access the raw methods on
}

// ContractIStrategyManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractIStrategyManagerCallerRaw struct {
	Contract *ContractIStrategyManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ContractIStrategyManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractIStrategyManagerTransactorRaw struct {
	Contract *ContractIStrategyManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractIStrategyManager creates a new instance of ContractIStrategyManager, bound to a specific deployed contract.
func NewContractIStrategyManager(address common.Address, backend bind.ContractBackend) (*ContractIStrategyManager, error) {
	contract, err := bindContractIStrategyManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractIStrategyManager{ContractIStrategyManagerCaller: ContractIStrategyManagerCaller{contract: contract}, ContractIStrategyManagerTransactor: ContractIStrategyManagerTransactor{contract: contract}, ContractIStrategyManagerFilterer: ContractIStrategyManagerFilterer{contract: contract}}, nil
}

// NewContractIStrategyManagerCaller creates a new read-only instance of ContractIStrategyManager, bound to a specific deployed contract.
func NewContractIStrategyManagerCaller(address common.Address, caller bind.ContractCaller) (*ContractIStrategyManagerCaller, error) {
	contract, err := bindContractIStrategyManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractIStrategyManagerCaller{contract: contract}, nil
}

// NewContractIStrategyManagerTransactor creates a new write-only instance of ContractIStrategyManager, bound to a specific deployed contract.
func NewContractIStrategyManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractIStrategyManagerTransactor, error) {
	contract, err := bindContractIStrategyManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractIStrategyManagerTransactor{contract: contract}, nil
}

// NewContractIStrategyManagerFilterer creates a new log filterer instance of ContractIStrategyManager, bound to a specific deployed contract.
func NewContractIStrategyManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractIStrategyManagerFilterer, error) {
	contract, err := bindContractIStrategyManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractIStrategyManagerFilterer{contract: contract}, nil
}

// bindContractIStrategyManager binds a generic wrapper to an already deployed contract.
func bindContractIStrategyManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractIStrategyManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractIStrategyManager *ContractIStrategyManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractIStrategyManager.Contract.ContractIStrategyManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractIStrategyManager *ContractIStrategyManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractIStrategyManager.Contract.ContractIStrategyManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractIStrategyManager *ContractIStrategyManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractIStrategyManager.Contract.ContractIStrategyManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractIStrategyManager *ContractIStrategyManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractIStrategyManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractIStrategyManager *ContractIStrategyManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractIStrategyManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractIStrategyManager *ContractIStrategyManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractIStrategyManager.Contract.contract.Transact(opts, method, params...)
}

// CalculateWithdrawalRoot is a free data retrieval call binding the contract method 0xb43b514b.
//
// Solidity: function calculateWithdrawalRoot((address[],uint256[],address,(address,uint96),uint32,address) queuedWithdrawal) pure returns(bytes32)
func (_ContractIStrategyManager *ContractIStrategyManagerCaller) CalculateWithdrawalRoot(opts *bind.CallOpts, queuedWithdrawal IStrategyManagerDeprecatedStructQueuedWithdrawal) ([32]byte, error) {
	var out []interface{}
	err := _ContractIStrategyManager.contract.Call(opts, &out, "calculateWithdrawalRoot", queuedWithdrawal)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateWithdrawalRoot is a free data retrieval call binding the contract method 0xb43b514b.
//
// Solidity: function calculateWithdrawalRoot((address[],uint256[],address,(address,uint96),uint32,address) queuedWithdrawal) pure returns(bytes32)
func (_ContractIStrategyManager *ContractIStrategyManagerSession) CalculateWithdrawalRoot(queuedWithdrawal IStrategyManagerDeprecatedStructQueuedWithdrawal) ([32]byte, error) {
	return _ContractIStrategyManager.Contract.CalculateWithdrawalRoot(&_ContractIStrategyManager.CallOpts, queuedWithdrawal)
}

// CalculateWithdrawalRoot is a free data retrieval call binding the contract method 0xb43b514b.
//
// Solidity: function calculateWithdrawalRoot((address[],uint256[],address,(address,uint96),uint32,address) queuedWithdrawal) pure returns(bytes32)
func (_ContractIStrategyManager *ContractIStrategyManagerCallerSession) CalculateWithdrawalRoot(queuedWithdrawal IStrategyManagerDeprecatedStructQueuedWithdrawal) ([32]byte, error) {
	return _ContractIStrategyManager.Contract.CalculateWithdrawalRoot(&_ContractIStrategyManager.CallOpts, queuedWithdrawal)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractIStrategyManager *ContractIStrategyManagerCaller) Delegation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIStrategyManager.contract.Call(opts, &out, "delegation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractIStrategyManager *ContractIStrategyManagerSession) Delegation() (common.Address, error) {
	return _ContractIStrategyManager.Contract.Delegation(&_ContractIStrategyManager.CallOpts)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractIStrategyManager *ContractIStrategyManagerCallerSession) Delegation() (common.Address, error) {
	return _ContractIStrategyManager.Contract.Delegation(&_ContractIStrategyManager.CallOpts)
}

// EigenPodManager is a free data retrieval call binding the contract method 0x4665bcda.
//
// Solidity: function eigenPodManager() view returns(address)
func (_ContractIStrategyManager *ContractIStrategyManagerCaller) EigenPodManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIStrategyManager.contract.Call(opts, &out, "eigenPodManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EigenPodManager is a free data retrieval call binding the contract method 0x4665bcda.
//
// Solidity: function eigenPodManager() view returns(address)
func (_ContractIStrategyManager *ContractIStrategyManagerSession) EigenPodManager() (common.Address, error) {
	return _ContractIStrategyManager.Contract.EigenPodManager(&_ContractIStrategyManager.CallOpts)
}

// EigenPodManager is a free data retrieval call binding the contract method 0x4665bcda.
//
// Solidity: function eigenPodManager() view returns(address)
func (_ContractIStrategyManager *ContractIStrategyManagerCallerSession) EigenPodManager() (common.Address, error) {
	return _ContractIStrategyManager.Contract.EigenPodManager(&_ContractIStrategyManager.CallOpts)
}

// GetDeposits is a free data retrieval call binding the contract method 0x94f649dd.
//
// Solidity: function getDeposits(address staker) view returns(address[], uint256[])
func (_ContractIStrategyManager *ContractIStrategyManagerCaller) GetDeposits(opts *bind.CallOpts, staker common.Address) ([]common.Address, []*big.Int, error) {
	var out []interface{}
	err := _ContractIStrategyManager.contract.Call(opts, &out, "getDeposits", staker)

	if err != nil {
		return *new([]common.Address), *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	out1 := *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return out0, out1, err

}

// GetDeposits is a free data retrieval call binding the contract method 0x94f649dd.
//
// Solidity: function getDeposits(address staker) view returns(address[], uint256[])
func (_ContractIStrategyManager *ContractIStrategyManagerSession) GetDeposits(staker common.Address) ([]common.Address, []*big.Int, error) {
	return _ContractIStrategyManager.Contract.GetDeposits(&_ContractIStrategyManager.CallOpts, staker)
}

// GetDeposits is a free data retrieval call binding the contract method 0x94f649dd.
//
// Solidity: function getDeposits(address staker) view returns(address[], uint256[])
func (_ContractIStrategyManager *ContractIStrategyManagerCallerSession) GetDeposits(staker common.Address) ([]common.Address, []*big.Int, error) {
	return _ContractIStrategyManager.Contract.GetDeposits(&_ContractIStrategyManager.CallOpts, staker)
}

// Slasher is a free data retrieval call binding the contract method 0xb1344271.
//
// Solidity: function slasher() view returns(address)
func (_ContractIStrategyManager *ContractIStrategyManagerCaller) Slasher(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIStrategyManager.contract.Call(opts, &out, "slasher")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Slasher is a free data retrieval call binding the contract method 0xb1344271.
//
// Solidity: function slasher() view returns(address)
func (_ContractIStrategyManager *ContractIStrategyManagerSession) Slasher() (common.Address, error) {
	return _ContractIStrategyManager.Contract.Slasher(&_ContractIStrategyManager.CallOpts)
}

// Slasher is a free data retrieval call binding the contract method 0xb1344271.
//
// Solidity: function slasher() view returns(address)
func (_ContractIStrategyManager *ContractIStrategyManagerCallerSession) Slasher() (common.Address, error) {
	return _ContractIStrategyManager.Contract.Slasher(&_ContractIStrategyManager.CallOpts)
}

// StakerStrategyListLength is a free data retrieval call binding the contract method 0x8b8aac3c.
//
// Solidity: function stakerStrategyListLength(address staker) view returns(uint256)
func (_ContractIStrategyManager *ContractIStrategyManagerCaller) StakerStrategyListLength(opts *bind.CallOpts, staker common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ContractIStrategyManager.contract.Call(opts, &out, "stakerStrategyListLength", staker)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakerStrategyListLength is a free data retrieval call binding the contract method 0x8b8aac3c.
//
// Solidity: function stakerStrategyListLength(address staker) view returns(uint256)
func (_ContractIStrategyManager *ContractIStrategyManagerSession) StakerStrategyListLength(staker common.Address) (*big.Int, error) {
	return _ContractIStrategyManager.Contract.StakerStrategyListLength(&_ContractIStrategyManager.CallOpts, staker)
}

// StakerStrategyListLength is a free data retrieval call binding the contract method 0x8b8aac3c.
//
// Solidity: function stakerStrategyListLength(address staker) view returns(uint256)
func (_ContractIStrategyManager *ContractIStrategyManagerCallerSession) StakerStrategyListLength(staker common.Address) (*big.Int, error) {
	return _ContractIStrategyManager.Contract.StakerStrategyListLength(&_ContractIStrategyManager.CallOpts, staker)
}

// StakerStrategyShares is a free data retrieval call binding the contract method 0x7a7e0d92.
//
// Solidity: function stakerStrategyShares(address user, address strategy) view returns(uint256 shares)
func (_ContractIStrategyManager *ContractIStrategyManagerCaller) StakerStrategyShares(opts *bind.CallOpts, user common.Address, strategy common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ContractIStrategyManager.contract.Call(opts, &out, "stakerStrategyShares", user, strategy)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakerStrategyShares is a free data retrieval call binding the contract method 0x7a7e0d92.
//
// Solidity: function stakerStrategyShares(address user, address strategy) view returns(uint256 shares)
func (_ContractIStrategyManager *ContractIStrategyManagerSession) StakerStrategyShares(user common.Address, strategy common.Address) (*big.Int, error) {
	return _ContractIStrategyManager.Contract.StakerStrategyShares(&_ContractIStrategyManager.CallOpts, user, strategy)
}

// StakerStrategyShares is a free data retrieval call binding the contract method 0x7a7e0d92.
//
// Solidity: function stakerStrategyShares(address user, address strategy) view returns(uint256 shares)
func (_ContractIStrategyManager *ContractIStrategyManagerCallerSession) StakerStrategyShares(user common.Address, strategy common.Address) (*big.Int, error) {
	return _ContractIStrategyManager.Contract.StakerStrategyShares(&_ContractIStrategyManager.CallOpts, user, strategy)
}

// StrategyWhitelister is a free data retrieval call binding the contract method 0x967fc0d2.
//
// Solidity: function strategyWhitelister() view returns(address)
func (_ContractIStrategyManager *ContractIStrategyManagerCaller) StrategyWhitelister(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIStrategyManager.contract.Call(opts, &out, "strategyWhitelister")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StrategyWhitelister is a free data retrieval call binding the contract method 0x967fc0d2.
//
// Solidity: function strategyWhitelister() view returns(address)
func (_ContractIStrategyManager *ContractIStrategyManagerSession) StrategyWhitelister() (common.Address, error) {
	return _ContractIStrategyManager.Contract.StrategyWhitelister(&_ContractIStrategyManager.CallOpts)
}

// StrategyWhitelister is a free data retrieval call binding the contract method 0x967fc0d2.
//
// Solidity: function strategyWhitelister() view returns(address)
func (_ContractIStrategyManager *ContractIStrategyManagerCallerSession) StrategyWhitelister() (common.Address, error) {
	return _ContractIStrategyManager.Contract.StrategyWhitelister(&_ContractIStrategyManager.CallOpts)
}

// AddShares is a paid mutator transaction binding the contract method 0x50ff7225.
//
// Solidity: function addShares(address staker, address strategy, uint256 shares) returns()
func (_ContractIStrategyManager *ContractIStrategyManagerTransactor) AddShares(opts *bind.TransactOpts, staker common.Address, strategy common.Address, shares *big.Int) (*types.Transaction, error) {
	return _ContractIStrategyManager.contract.Transact(opts, "addShares", staker, strategy, shares)
}

// AddShares is a paid mutator transaction binding the contract method 0x50ff7225.
//
// Solidity: function addShares(address staker, address strategy, uint256 shares) returns()
func (_ContractIStrategyManager *ContractIStrategyManagerSession) AddShares(staker common.Address, strategy common.Address, shares *big.Int) (*types.Transaction, error) {
	return _ContractIStrategyManager.Contract.AddShares(&_ContractIStrategyManager.TransactOpts, staker, strategy, shares)
}

// AddShares is a paid mutator transaction binding the contract method 0x50ff7225.
//
// Solidity: function addShares(address staker, address strategy, uint256 shares) returns()
func (_ContractIStrategyManager *ContractIStrategyManagerTransactorSession) AddShares(staker common.Address, strategy common.Address, shares *big.Int) (*types.Transaction, error) {
	return _ContractIStrategyManager.Contract.AddShares(&_ContractIStrategyManager.TransactOpts, staker, strategy, shares)
}

// AddStrategiesToDepositWhitelist is a paid mutator transaction binding the contract method 0x5de08ff2.
//
// Solidity: function addStrategiesToDepositWhitelist(address[] strategiesToWhitelist) returns()
func (_ContractIStrategyManager *ContractIStrategyManagerTransactor) AddStrategiesToDepositWhitelist(opts *bind.TransactOpts, strategiesToWhitelist []common.Address) (*types.Transaction, error) {
	return _ContractIStrategyManager.contract.Transact(opts, "addStrategiesToDepositWhitelist", strategiesToWhitelist)
}

// AddStrategiesToDepositWhitelist is a paid mutator transaction binding the contract method 0x5de08ff2.
//
// Solidity: function addStrategiesToDepositWhitelist(address[] strategiesToWhitelist) returns()
func (_ContractIStrategyManager *ContractIStrategyManagerSession) AddStrategiesToDepositWhitelist(strategiesToWhitelist []common.Address) (*types.Transaction, error) {
	return _ContractIStrategyManager.Contract.AddStrategiesToDepositWhitelist(&_ContractIStrategyManager.TransactOpts, strategiesToWhitelist)
}

// AddStrategiesToDepositWhitelist is a paid mutator transaction binding the contract method 0x5de08ff2.
//
// Solidity: function addStrategiesToDepositWhitelist(address[] strategiesToWhitelist) returns()
func (_ContractIStrategyManager *ContractIStrategyManagerTransactorSession) AddStrategiesToDepositWhitelist(strategiesToWhitelist []common.Address) (*types.Transaction, error) {
	return _ContractIStrategyManager.Contract.AddStrategiesToDepositWhitelist(&_ContractIStrategyManager.TransactOpts, strategiesToWhitelist)
}

// DepositIntoStrategy is a paid mutator transaction binding the contract method 0xe7a050aa.
//
// Solidity: function depositIntoStrategy(address strategy, address token, uint256 amount) returns(uint256 shares)
func (_ContractIStrategyManager *ContractIStrategyManagerTransactor) DepositIntoStrategy(opts *bind.TransactOpts, strategy common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ContractIStrategyManager.contract.Transact(opts, "depositIntoStrategy", strategy, token, amount)
}

// DepositIntoStrategy is a paid mutator transaction binding the contract method 0xe7a050aa.
//
// Solidity: function depositIntoStrategy(address strategy, address token, uint256 amount) returns(uint256 shares)
func (_ContractIStrategyManager *ContractIStrategyManagerSession) DepositIntoStrategy(strategy common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ContractIStrategyManager.Contract.DepositIntoStrategy(&_ContractIStrategyManager.TransactOpts, strategy, token, amount)
}

// DepositIntoStrategy is a paid mutator transaction binding the contract method 0xe7a050aa.
//
// Solidity: function depositIntoStrategy(address strategy, address token, uint256 amount) returns(uint256 shares)
func (_ContractIStrategyManager *ContractIStrategyManagerTransactorSession) DepositIntoStrategy(strategy common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ContractIStrategyManager.Contract.DepositIntoStrategy(&_ContractIStrategyManager.TransactOpts, strategy, token, amount)
}

// DepositIntoStrategyWithSignature is a paid mutator transaction binding the contract method 0x32e89ace.
//
// Solidity: function depositIntoStrategyWithSignature(address strategy, address token, uint256 amount, address staker, uint256 expiry, bytes signature) returns(uint256 shares)
func (_ContractIStrategyManager *ContractIStrategyManagerTransactor) DepositIntoStrategyWithSignature(opts *bind.TransactOpts, strategy common.Address, token common.Address, amount *big.Int, staker common.Address, expiry *big.Int, signature []byte) (*types.Transaction, error) {
	return _ContractIStrategyManager.contract.Transact(opts, "depositIntoStrategyWithSignature", strategy, token, amount, staker, expiry, signature)
}

// DepositIntoStrategyWithSignature is a paid mutator transaction binding the contract method 0x32e89ace.
//
// Solidity: function depositIntoStrategyWithSignature(address strategy, address token, uint256 amount, address staker, uint256 expiry, bytes signature) returns(uint256 shares)
func (_ContractIStrategyManager *ContractIStrategyManagerSession) DepositIntoStrategyWithSignature(strategy common.Address, token common.Address, amount *big.Int, staker common.Address, expiry *big.Int, signature []byte) (*types.Transaction, error) {
	return _ContractIStrategyManager.Contract.DepositIntoStrategyWithSignature(&_ContractIStrategyManager.TransactOpts, strategy, token, amount, staker, expiry, signature)
}

// DepositIntoStrategyWithSignature is a paid mutator transaction binding the contract method 0x32e89ace.
//
// Solidity: function depositIntoStrategyWithSignature(address strategy, address token, uint256 amount, address staker, uint256 expiry, bytes signature) returns(uint256 shares)
func (_ContractIStrategyManager *ContractIStrategyManagerTransactorSession) DepositIntoStrategyWithSignature(strategy common.Address, token common.Address, amount *big.Int, staker common.Address, expiry *big.Int, signature []byte) (*types.Transaction, error) {
	return _ContractIStrategyManager.Contract.DepositIntoStrategyWithSignature(&_ContractIStrategyManager.TransactOpts, strategy, token, amount, staker, expiry, signature)
}

// MigrateQueuedWithdrawal is a paid mutator transaction binding the contract method 0xcd293f6f.
//
// Solidity: function migrateQueuedWithdrawal((address[],uint256[],address,(address,uint96),uint32,address) queuedWithdrawal) returns(bool, bytes32)
func (_ContractIStrategyManager *ContractIStrategyManagerTransactor) MigrateQueuedWithdrawal(opts *bind.TransactOpts, queuedWithdrawal IStrategyManagerDeprecatedStructQueuedWithdrawal) (*types.Transaction, error) {
	return _ContractIStrategyManager.contract.Transact(opts, "migrateQueuedWithdrawal", queuedWithdrawal)
}

// MigrateQueuedWithdrawal is a paid mutator transaction binding the contract method 0xcd293f6f.
//
// Solidity: function migrateQueuedWithdrawal((address[],uint256[],address,(address,uint96),uint32,address) queuedWithdrawal) returns(bool, bytes32)
func (_ContractIStrategyManager *ContractIStrategyManagerSession) MigrateQueuedWithdrawal(queuedWithdrawal IStrategyManagerDeprecatedStructQueuedWithdrawal) (*types.Transaction, error) {
	return _ContractIStrategyManager.Contract.MigrateQueuedWithdrawal(&_ContractIStrategyManager.TransactOpts, queuedWithdrawal)
}

// MigrateQueuedWithdrawal is a paid mutator transaction binding the contract method 0xcd293f6f.
//
// Solidity: function migrateQueuedWithdrawal((address[],uint256[],address,(address,uint96),uint32,address) queuedWithdrawal) returns(bool, bytes32)
func (_ContractIStrategyManager *ContractIStrategyManagerTransactorSession) MigrateQueuedWithdrawal(queuedWithdrawal IStrategyManagerDeprecatedStructQueuedWithdrawal) (*types.Transaction, error) {
	return _ContractIStrategyManager.Contract.MigrateQueuedWithdrawal(&_ContractIStrategyManager.TransactOpts, queuedWithdrawal)
}

// RemoveShares is a paid mutator transaction binding the contract method 0x8c80d4e5.
//
// Solidity: function removeShares(address staker, address strategy, uint256 shares) returns()
func (_ContractIStrategyManager *ContractIStrategyManagerTransactor) RemoveShares(opts *bind.TransactOpts, staker common.Address, strategy common.Address, shares *big.Int) (*types.Transaction, error) {
	return _ContractIStrategyManager.contract.Transact(opts, "removeShares", staker, strategy, shares)
}

// RemoveShares is a paid mutator transaction binding the contract method 0x8c80d4e5.
//
// Solidity: function removeShares(address staker, address strategy, uint256 shares) returns()
func (_ContractIStrategyManager *ContractIStrategyManagerSession) RemoveShares(staker common.Address, strategy common.Address, shares *big.Int) (*types.Transaction, error) {
	return _ContractIStrategyManager.Contract.RemoveShares(&_ContractIStrategyManager.TransactOpts, staker, strategy, shares)
}

// RemoveShares is a paid mutator transaction binding the contract method 0x8c80d4e5.
//
// Solidity: function removeShares(address staker, address strategy, uint256 shares) returns()
func (_ContractIStrategyManager *ContractIStrategyManagerTransactorSession) RemoveShares(staker common.Address, strategy common.Address, shares *big.Int) (*types.Transaction, error) {
	return _ContractIStrategyManager.Contract.RemoveShares(&_ContractIStrategyManager.TransactOpts, staker, strategy, shares)
}

// RemoveStrategiesFromDepositWhitelist is a paid mutator transaction binding the contract method 0xb5d8b5b8.
//
// Solidity: function removeStrategiesFromDepositWhitelist(address[] strategiesToRemoveFromWhitelist) returns()
func (_ContractIStrategyManager *ContractIStrategyManagerTransactor) RemoveStrategiesFromDepositWhitelist(opts *bind.TransactOpts, strategiesToRemoveFromWhitelist []common.Address) (*types.Transaction, error) {
	return _ContractIStrategyManager.contract.Transact(opts, "removeStrategiesFromDepositWhitelist", strategiesToRemoveFromWhitelist)
}

// RemoveStrategiesFromDepositWhitelist is a paid mutator transaction binding the contract method 0xb5d8b5b8.
//
// Solidity: function removeStrategiesFromDepositWhitelist(address[] strategiesToRemoveFromWhitelist) returns()
func (_ContractIStrategyManager *ContractIStrategyManagerSession) RemoveStrategiesFromDepositWhitelist(strategiesToRemoveFromWhitelist []common.Address) (*types.Transaction, error) {
	return _ContractIStrategyManager.Contract.RemoveStrategiesFromDepositWhitelist(&_ContractIStrategyManager.TransactOpts, strategiesToRemoveFromWhitelist)
}

// RemoveStrategiesFromDepositWhitelist is a paid mutator transaction binding the contract method 0xb5d8b5b8.
//
// Solidity: function removeStrategiesFromDepositWhitelist(address[] strategiesToRemoveFromWhitelist) returns()
func (_ContractIStrategyManager *ContractIStrategyManagerTransactorSession) RemoveStrategiesFromDepositWhitelist(strategiesToRemoveFromWhitelist []common.Address) (*types.Transaction, error) {
	return _ContractIStrategyManager.Contract.RemoveStrategiesFromDepositWhitelist(&_ContractIStrategyManager.TransactOpts, strategiesToRemoveFromWhitelist)
}

// WithdrawSharesAsTokens is a paid mutator transaction binding the contract method 0xc608c7f3.
//
// Solidity: function withdrawSharesAsTokens(address recipient, address strategy, uint256 shares, address token) returns()
func (_ContractIStrategyManager *ContractIStrategyManagerTransactor) WithdrawSharesAsTokens(opts *bind.TransactOpts, recipient common.Address, strategy common.Address, shares *big.Int, token common.Address) (*types.Transaction, error) {
	return _ContractIStrategyManager.contract.Transact(opts, "withdrawSharesAsTokens", recipient, strategy, shares, token)
}

// WithdrawSharesAsTokens is a paid mutator transaction binding the contract method 0xc608c7f3.
//
// Solidity: function withdrawSharesAsTokens(address recipient, address strategy, uint256 shares, address token) returns()
func (_ContractIStrategyManager *ContractIStrategyManagerSession) WithdrawSharesAsTokens(recipient common.Address, strategy common.Address, shares *big.Int, token common.Address) (*types.Transaction, error) {
	return _ContractIStrategyManager.Contract.WithdrawSharesAsTokens(&_ContractIStrategyManager.TransactOpts, recipient, strategy, shares, token)
}

// WithdrawSharesAsTokens is a paid mutator transaction binding the contract method 0xc608c7f3.
//
// Solidity: function withdrawSharesAsTokens(address recipient, address strategy, uint256 shares, address token) returns()
func (_ContractIStrategyManager *ContractIStrategyManagerTransactorSession) WithdrawSharesAsTokens(recipient common.Address, strategy common.Address, shares *big.Int, token common.Address) (*types.Transaction, error) {
	return _ContractIStrategyManager.Contract.WithdrawSharesAsTokens(&_ContractIStrategyManager.TransactOpts, recipient, strategy, shares, token)
}

// ContractIStrategyManagerDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the ContractIStrategyManager contract.
type ContractIStrategyManagerDepositIterator struct {
	Event *ContractIStrategyManagerDeposit // Event containing the contract specifics and raw log

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
func (it *ContractIStrategyManagerDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIStrategyManagerDeposit)
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
		it.Event = new(ContractIStrategyManagerDeposit)
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
func (it *ContractIStrategyManagerDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIStrategyManagerDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIStrategyManagerDeposit represents a Deposit event raised by the ContractIStrategyManager contract.
type ContractIStrategyManagerDeposit struct {
	Staker   common.Address
	Token    common.Address
	Strategy common.Address
	Shares   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x7cfff908a4b583f36430b25d75964c458d8ede8a99bd61be750e97ee1b2f3a96.
//
// Solidity: event Deposit(address staker, address token, address strategy, uint256 shares)
func (_ContractIStrategyManager *ContractIStrategyManagerFilterer) FilterDeposit(opts *bind.FilterOpts) (*ContractIStrategyManagerDepositIterator, error) {

	logs, sub, err := _ContractIStrategyManager.contract.FilterLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return &ContractIStrategyManagerDepositIterator{contract: _ContractIStrategyManager.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x7cfff908a4b583f36430b25d75964c458d8ede8a99bd61be750e97ee1b2f3a96.
//
// Solidity: event Deposit(address staker, address token, address strategy, uint256 shares)
func (_ContractIStrategyManager *ContractIStrategyManagerFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *ContractIStrategyManagerDeposit) (event.Subscription, error) {

	logs, sub, err := _ContractIStrategyManager.contract.WatchLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIStrategyManagerDeposit)
				if err := _ContractIStrategyManager.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x7cfff908a4b583f36430b25d75964c458d8ede8a99bd61be750e97ee1b2f3a96.
//
// Solidity: event Deposit(address staker, address token, address strategy, uint256 shares)
func (_ContractIStrategyManager *ContractIStrategyManagerFilterer) ParseDeposit(log types.Log) (*ContractIStrategyManagerDeposit, error) {
	event := new(ContractIStrategyManagerDeposit)
	if err := _ContractIStrategyManager.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIStrategyManagerStrategyAddedToDepositWhitelistIterator is returned from FilterStrategyAddedToDepositWhitelist and is used to iterate over the raw logs and unpacked data for StrategyAddedToDepositWhitelist events raised by the ContractIStrategyManager contract.
type ContractIStrategyManagerStrategyAddedToDepositWhitelistIterator struct {
	Event *ContractIStrategyManagerStrategyAddedToDepositWhitelist // Event containing the contract specifics and raw log

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
func (it *ContractIStrategyManagerStrategyAddedToDepositWhitelistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIStrategyManagerStrategyAddedToDepositWhitelist)
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
		it.Event = new(ContractIStrategyManagerStrategyAddedToDepositWhitelist)
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
func (it *ContractIStrategyManagerStrategyAddedToDepositWhitelistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIStrategyManagerStrategyAddedToDepositWhitelistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIStrategyManagerStrategyAddedToDepositWhitelist represents a StrategyAddedToDepositWhitelist event raised by the ContractIStrategyManager contract.
type ContractIStrategyManagerStrategyAddedToDepositWhitelist struct {
	Strategy common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStrategyAddedToDepositWhitelist is a free log retrieval operation binding the contract event 0x0c35b17d91c96eb2751cd456e1252f42a386e524ef9ff26ecc9950859fdc04fe.
//
// Solidity: event StrategyAddedToDepositWhitelist(address strategy)
func (_ContractIStrategyManager *ContractIStrategyManagerFilterer) FilterStrategyAddedToDepositWhitelist(opts *bind.FilterOpts) (*ContractIStrategyManagerStrategyAddedToDepositWhitelistIterator, error) {

	logs, sub, err := _ContractIStrategyManager.contract.FilterLogs(opts, "StrategyAddedToDepositWhitelist")
	if err != nil {
		return nil, err
	}
	return &ContractIStrategyManagerStrategyAddedToDepositWhitelistIterator{contract: _ContractIStrategyManager.contract, event: "StrategyAddedToDepositWhitelist", logs: logs, sub: sub}, nil
}

// WatchStrategyAddedToDepositWhitelist is a free log subscription operation binding the contract event 0x0c35b17d91c96eb2751cd456e1252f42a386e524ef9ff26ecc9950859fdc04fe.
//
// Solidity: event StrategyAddedToDepositWhitelist(address strategy)
func (_ContractIStrategyManager *ContractIStrategyManagerFilterer) WatchStrategyAddedToDepositWhitelist(opts *bind.WatchOpts, sink chan<- *ContractIStrategyManagerStrategyAddedToDepositWhitelist) (event.Subscription, error) {

	logs, sub, err := _ContractIStrategyManager.contract.WatchLogs(opts, "StrategyAddedToDepositWhitelist")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIStrategyManagerStrategyAddedToDepositWhitelist)
				if err := _ContractIStrategyManager.contract.UnpackLog(event, "StrategyAddedToDepositWhitelist", log); err != nil {
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

// ParseStrategyAddedToDepositWhitelist is a log parse operation binding the contract event 0x0c35b17d91c96eb2751cd456e1252f42a386e524ef9ff26ecc9950859fdc04fe.
//
// Solidity: event StrategyAddedToDepositWhitelist(address strategy)
func (_ContractIStrategyManager *ContractIStrategyManagerFilterer) ParseStrategyAddedToDepositWhitelist(log types.Log) (*ContractIStrategyManagerStrategyAddedToDepositWhitelist, error) {
	event := new(ContractIStrategyManagerStrategyAddedToDepositWhitelist)
	if err := _ContractIStrategyManager.contract.UnpackLog(event, "StrategyAddedToDepositWhitelist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIStrategyManagerStrategyRemovedFromDepositWhitelistIterator is returned from FilterStrategyRemovedFromDepositWhitelist and is used to iterate over the raw logs and unpacked data for StrategyRemovedFromDepositWhitelist events raised by the ContractIStrategyManager contract.
type ContractIStrategyManagerStrategyRemovedFromDepositWhitelistIterator struct {
	Event *ContractIStrategyManagerStrategyRemovedFromDepositWhitelist // Event containing the contract specifics and raw log

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
func (it *ContractIStrategyManagerStrategyRemovedFromDepositWhitelistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIStrategyManagerStrategyRemovedFromDepositWhitelist)
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
		it.Event = new(ContractIStrategyManagerStrategyRemovedFromDepositWhitelist)
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
func (it *ContractIStrategyManagerStrategyRemovedFromDepositWhitelistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIStrategyManagerStrategyRemovedFromDepositWhitelistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIStrategyManagerStrategyRemovedFromDepositWhitelist represents a StrategyRemovedFromDepositWhitelist event raised by the ContractIStrategyManager contract.
type ContractIStrategyManagerStrategyRemovedFromDepositWhitelist struct {
	Strategy common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStrategyRemovedFromDepositWhitelist is a free log retrieval operation binding the contract event 0x4074413b4b443e4e58019f2855a8765113358c7c72e39509c6af45fc0f5ba030.
//
// Solidity: event StrategyRemovedFromDepositWhitelist(address strategy)
func (_ContractIStrategyManager *ContractIStrategyManagerFilterer) FilterStrategyRemovedFromDepositWhitelist(opts *bind.FilterOpts) (*ContractIStrategyManagerStrategyRemovedFromDepositWhitelistIterator, error) {

	logs, sub, err := _ContractIStrategyManager.contract.FilterLogs(opts, "StrategyRemovedFromDepositWhitelist")
	if err != nil {
		return nil, err
	}
	return &ContractIStrategyManagerStrategyRemovedFromDepositWhitelistIterator{contract: _ContractIStrategyManager.contract, event: "StrategyRemovedFromDepositWhitelist", logs: logs, sub: sub}, nil
}

// WatchStrategyRemovedFromDepositWhitelist is a free log subscription operation binding the contract event 0x4074413b4b443e4e58019f2855a8765113358c7c72e39509c6af45fc0f5ba030.
//
// Solidity: event StrategyRemovedFromDepositWhitelist(address strategy)
func (_ContractIStrategyManager *ContractIStrategyManagerFilterer) WatchStrategyRemovedFromDepositWhitelist(opts *bind.WatchOpts, sink chan<- *ContractIStrategyManagerStrategyRemovedFromDepositWhitelist) (event.Subscription, error) {

	logs, sub, err := _ContractIStrategyManager.contract.WatchLogs(opts, "StrategyRemovedFromDepositWhitelist")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIStrategyManagerStrategyRemovedFromDepositWhitelist)
				if err := _ContractIStrategyManager.contract.UnpackLog(event, "StrategyRemovedFromDepositWhitelist", log); err != nil {
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

// ParseStrategyRemovedFromDepositWhitelist is a log parse operation binding the contract event 0x4074413b4b443e4e58019f2855a8765113358c7c72e39509c6af45fc0f5ba030.
//
// Solidity: event StrategyRemovedFromDepositWhitelist(address strategy)
func (_ContractIStrategyManager *ContractIStrategyManagerFilterer) ParseStrategyRemovedFromDepositWhitelist(log types.Log) (*ContractIStrategyManagerStrategyRemovedFromDepositWhitelist, error) {
	event := new(ContractIStrategyManagerStrategyRemovedFromDepositWhitelist)
	if err := _ContractIStrategyManager.contract.UnpackLog(event, "StrategyRemovedFromDepositWhitelist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIStrategyManagerStrategyWhitelisterChangedIterator is returned from FilterStrategyWhitelisterChanged and is used to iterate over the raw logs and unpacked data for StrategyWhitelisterChanged events raised by the ContractIStrategyManager contract.
type ContractIStrategyManagerStrategyWhitelisterChangedIterator struct {
	Event *ContractIStrategyManagerStrategyWhitelisterChanged // Event containing the contract specifics and raw log

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
func (it *ContractIStrategyManagerStrategyWhitelisterChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIStrategyManagerStrategyWhitelisterChanged)
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
		it.Event = new(ContractIStrategyManagerStrategyWhitelisterChanged)
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
func (it *ContractIStrategyManagerStrategyWhitelisterChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIStrategyManagerStrategyWhitelisterChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIStrategyManagerStrategyWhitelisterChanged represents a StrategyWhitelisterChanged event raised by the ContractIStrategyManager contract.
type ContractIStrategyManagerStrategyWhitelisterChanged struct {
	PreviousAddress common.Address
	NewAddress      common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterStrategyWhitelisterChanged is a free log retrieval operation binding the contract event 0x4264275e593955ff9d6146a51a4525f6ddace2e81db9391abcc9d1ca48047d29.
//
// Solidity: event StrategyWhitelisterChanged(address previousAddress, address newAddress)
func (_ContractIStrategyManager *ContractIStrategyManagerFilterer) FilterStrategyWhitelisterChanged(opts *bind.FilterOpts) (*ContractIStrategyManagerStrategyWhitelisterChangedIterator, error) {

	logs, sub, err := _ContractIStrategyManager.contract.FilterLogs(opts, "StrategyWhitelisterChanged")
	if err != nil {
		return nil, err
	}
	return &ContractIStrategyManagerStrategyWhitelisterChangedIterator{contract: _ContractIStrategyManager.contract, event: "StrategyWhitelisterChanged", logs: logs, sub: sub}, nil
}

// WatchStrategyWhitelisterChanged is a free log subscription operation binding the contract event 0x4264275e593955ff9d6146a51a4525f6ddace2e81db9391abcc9d1ca48047d29.
//
// Solidity: event StrategyWhitelisterChanged(address previousAddress, address newAddress)
func (_ContractIStrategyManager *ContractIStrategyManagerFilterer) WatchStrategyWhitelisterChanged(opts *bind.WatchOpts, sink chan<- *ContractIStrategyManagerStrategyWhitelisterChanged) (event.Subscription, error) {

	logs, sub, err := _ContractIStrategyManager.contract.WatchLogs(opts, "StrategyWhitelisterChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIStrategyManagerStrategyWhitelisterChanged)
				if err := _ContractIStrategyManager.contract.UnpackLog(event, "StrategyWhitelisterChanged", log); err != nil {
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

// ParseStrategyWhitelisterChanged is a log parse operation binding the contract event 0x4264275e593955ff9d6146a51a4525f6ddace2e81db9391abcc9d1ca48047d29.
//
// Solidity: event StrategyWhitelisterChanged(address previousAddress, address newAddress)
func (_ContractIStrategyManager *ContractIStrategyManagerFilterer) ParseStrategyWhitelisterChanged(log types.Log) (*ContractIStrategyManagerStrategyWhitelisterChanged, error) {
	event := new(ContractIStrategyManagerStrategyWhitelisterChanged)
	if err := _ContractIStrategyManager.contract.UnpackLog(event, "StrategyWhitelisterChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
