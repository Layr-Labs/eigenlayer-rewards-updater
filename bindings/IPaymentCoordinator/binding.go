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

// IPaymentCoordinatorRangePayment is an auto generated low-level Go binding around an user-defined struct.
type IPaymentCoordinatorRangePayment struct {
	Avs                 common.Address
	Strategy            common.Address
	Token               common.Address
	Amount              *big.Int
	StartRangeTimestamp *big.Int
	EndRangeTimestamp   *big.Int
}

// ContractIPaymentCoordinatorMetaData contains all meta data concerning the ContractIPaymentCoordinator contract.
var ContractIPaymentCoordinatorMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"LOWER_BOUND_START_RANGE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_FUTURE_RANGE_END\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claimingManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIClaimingManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"payForRange\",\"inputs\":[{\"name\":\"rangePayment\",\"type\":\"tuple\",\"internalType\":\"structIPaymentCoordinator.RangePayment\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startRangeTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"endRangeTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"RangePaymentCreated\",\"inputs\":[{\"name\":\"rangePayment\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIPaymentCoordinator.RangePayment\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startRangeTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"endRangeTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false}]",
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

// LOWERBOUNDSTARTRANGE is a free data retrieval call binding the contract method 0xfa3733f9.
//
// Solidity: function LOWER_BOUND_START_RANGE() view returns(uint32)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorCaller) LOWERBOUNDSTARTRANGE(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractIPaymentCoordinator.contract.Call(opts, &out, "LOWER_BOUND_START_RANGE")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LOWERBOUNDSTARTRANGE is a free data retrieval call binding the contract method 0xfa3733f9.
//
// Solidity: function LOWER_BOUND_START_RANGE() view returns(uint32)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorSession) LOWERBOUNDSTARTRANGE() (uint32, error) {
	return _ContractIPaymentCoordinator.Contract.LOWERBOUNDSTARTRANGE(&_ContractIPaymentCoordinator.CallOpts)
}

// LOWERBOUNDSTARTRANGE is a free data retrieval call binding the contract method 0xfa3733f9.
//
// Solidity: function LOWER_BOUND_START_RANGE() view returns(uint32)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorCallerSession) LOWERBOUNDSTARTRANGE() (uint32, error) {
	return _ContractIPaymentCoordinator.Contract.LOWERBOUNDSTARTRANGE(&_ContractIPaymentCoordinator.CallOpts)
}

// MAXFUTURERANGEEND is a free data retrieval call binding the contract method 0x3709210e.
//
// Solidity: function MAX_FUTURE_RANGE_END() view returns(uint32)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorCaller) MAXFUTURERANGEEND(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractIPaymentCoordinator.contract.Call(opts, &out, "MAX_FUTURE_RANGE_END")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// MAXFUTURERANGEEND is a free data retrieval call binding the contract method 0x3709210e.
//
// Solidity: function MAX_FUTURE_RANGE_END() view returns(uint32)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorSession) MAXFUTURERANGEEND() (uint32, error) {
	return _ContractIPaymentCoordinator.Contract.MAXFUTURERANGEEND(&_ContractIPaymentCoordinator.CallOpts)
}

// MAXFUTURERANGEEND is a free data retrieval call binding the contract method 0x3709210e.
//
// Solidity: function MAX_FUTURE_RANGE_END() view returns(uint32)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorCallerSession) MAXFUTURERANGEEND() (uint32, error) {
	return _ContractIPaymentCoordinator.Contract.MAXFUTURERANGEEND(&_ContractIPaymentCoordinator.CallOpts)
}

// ClaimingManager is a free data retrieval call binding the contract method 0x7409c962.
//
// Solidity: function claimingManager() view returns(address)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorCaller) ClaimingManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIPaymentCoordinator.contract.Call(opts, &out, "claimingManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ClaimingManager is a free data retrieval call binding the contract method 0x7409c962.
//
// Solidity: function claimingManager() view returns(address)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorSession) ClaimingManager() (common.Address, error) {
	return _ContractIPaymentCoordinator.Contract.ClaimingManager(&_ContractIPaymentCoordinator.CallOpts)
}

// ClaimingManager is a free data retrieval call binding the contract method 0x7409c962.
//
// Solidity: function claimingManager() view returns(address)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorCallerSession) ClaimingManager() (common.Address, error) {
	return _ContractIPaymentCoordinator.Contract.ClaimingManager(&_ContractIPaymentCoordinator.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address) (*types.Transaction, error) {
	return _ContractIPaymentCoordinator.contract.Transact(opts, "initialize", initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorSession) Initialize(initialOwner common.Address) (*types.Transaction, error) {
	return _ContractIPaymentCoordinator.Contract.Initialize(&_ContractIPaymentCoordinator.TransactOpts, initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorTransactorSession) Initialize(initialOwner common.Address) (*types.Transaction, error) {
	return _ContractIPaymentCoordinator.Contract.Initialize(&_ContractIPaymentCoordinator.TransactOpts, initialOwner)
}

// PayForRange is a paid mutator transaction binding the contract method 0x96b0f235.
//
// Solidity: function payForRange((address,address,address,uint256,uint256,uint256) rangePayment) returns()
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorTransactor) PayForRange(opts *bind.TransactOpts, rangePayment IPaymentCoordinatorRangePayment) (*types.Transaction, error) {
	return _ContractIPaymentCoordinator.contract.Transact(opts, "payForRange", rangePayment)
}

// PayForRange is a paid mutator transaction binding the contract method 0x96b0f235.
//
// Solidity: function payForRange((address,address,address,uint256,uint256,uint256) rangePayment) returns()
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorSession) PayForRange(rangePayment IPaymentCoordinatorRangePayment) (*types.Transaction, error) {
	return _ContractIPaymentCoordinator.Contract.PayForRange(&_ContractIPaymentCoordinator.TransactOpts, rangePayment)
}

// PayForRange is a paid mutator transaction binding the contract method 0x96b0f235.
//
// Solidity: function payForRange((address,address,address,uint256,uint256,uint256) rangePayment) returns()
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorTransactorSession) PayForRange(rangePayment IPaymentCoordinatorRangePayment) (*types.Transaction, error) {
	return _ContractIPaymentCoordinator.Contract.PayForRange(&_ContractIPaymentCoordinator.TransactOpts, rangePayment)
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
	RangePayment IPaymentCoordinatorRangePayment
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRangePaymentCreated is a free log retrieval operation binding the contract event 0x31a3379c66ceda2d9ff087d7adffc10edd7a4e7b7585c9370e55e15c43fcb58c.
//
// Solidity: event RangePaymentCreated((address,address,address,uint256,uint256,uint256) rangePayment)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorFilterer) FilterRangePaymentCreated(opts *bind.FilterOpts) (*ContractIPaymentCoordinatorRangePaymentCreatedIterator, error) {

	logs, sub, err := _ContractIPaymentCoordinator.contract.FilterLogs(opts, "RangePaymentCreated")
	if err != nil {
		return nil, err
	}
	return &ContractIPaymentCoordinatorRangePaymentCreatedIterator{contract: _ContractIPaymentCoordinator.contract, event: "RangePaymentCreated", logs: logs, sub: sub}, nil
}

// WatchRangePaymentCreated is a free log subscription operation binding the contract event 0x31a3379c66ceda2d9ff087d7adffc10edd7a4e7b7585c9370e55e15c43fcb58c.
//
// Solidity: event RangePaymentCreated((address,address,address,uint256,uint256,uint256) rangePayment)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorFilterer) WatchRangePaymentCreated(opts *bind.WatchOpts, sink chan<- *ContractIPaymentCoordinatorRangePaymentCreated) (event.Subscription, error) {

	logs, sub, err := _ContractIPaymentCoordinator.contract.WatchLogs(opts, "RangePaymentCreated")
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

// ParseRangePaymentCreated is a log parse operation binding the contract event 0x31a3379c66ceda2d9ff087d7adffc10edd7a4e7b7585c9370e55e15c43fcb58c.
//
// Solidity: event RangePaymentCreated((address,address,address,uint256,uint256,uint256) rangePayment)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorFilterer) ParseRangePaymentCreated(log types.Log) (*ContractIPaymentCoordinatorRangePaymentCreated, error) {
	event := new(ContractIPaymentCoordinatorRangePaymentCreated)
	if err := _ContractIPaymentCoordinator.contract.UnpackLog(event, "RangePaymentCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
