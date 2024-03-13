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

// IPaymentCoordinatorMerkleLeaf is an auto generated low-level Go binding around an user-defined struct.
type IPaymentCoordinatorMerkleLeaf struct {
	Token     common.Address
	Amount    *big.Int
	Recipient common.Address
}

// IPaymentCoordinatorPaymentMerkleClaim is an auto generated low-level Go binding around an user-defined struct.
type IPaymentCoordinatorPaymentMerkleClaim struct {
	Leaf      IPaymentCoordinatorMerkleLeaf
	RootIndex uint32
	LeafIndex uint32
	Proof     []byte
}

// IPaymentCoordinatorRangePayment is an auto generated low-level Go binding around an user-defined struct.
type IPaymentCoordinatorRangePayment struct {
	Strategies     []common.Address
	Weights        []*big.Int
	Token          common.Address
	Amount         *big.Int
	StartTimestamp *big.Int
	Duration       *big.Int
}

// ContractIPaymentCoordinatorMetaData contains all meta data concerning the ContractIPaymentCoordinator contract.
var ContractIPaymentCoordinatorMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"LOWER_BOUND_START_RANGE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_PAYMENT_DURATION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"activationDelay\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculationIntervalSeconds\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkClaim\",\"inputs\":[{\"name\":\"claim\",\"type\":\"tuple\",\"internalType\":\"structIPaymentCoordinator.PaymentMerkleClaim\",\"components\":[{\"name\":\"leaf\",\"type\":\"tuple\",\"internalType\":\"structIPaymentCoordinator.MerkleLeaf\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"rootIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"leafIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"cumulativeClaimed\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"findLeafHash\",\"inputs\":[{\"name\":\"leaf\",\"type\":\"tuple\",\"internalType\":\"structIPaymentCoordinator.MerkleLeaf\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"globalOperatorCommissionBips\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_activationDelay\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_globalCommissionBips\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"payAllForRange\",\"inputs\":[{\"name\":\"rangePayment\",\"type\":\"tuple\",\"internalType\":\"structIPaymentCoordinator.RangePayment\",\"components\":[{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"weights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"duration\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"payForRange\",\"inputs\":[{\"name\":\"rangePayment\",\"type\":\"tuple\",\"internalType\":\"structIPaymentCoordinator.RangePayment\",\"components\":[{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"weights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"duration\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paymentUpdater\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"processClaims\",\"inputs\":[{\"name\":\"claims\",\"type\":\"tuple[]\",\"internalType\":\"structIPaymentCoordinator.PaymentMerkleClaim[]\",\"components\":[{\"name\":\"leaf\",\"type\":\"tuple\",\"internalType\":\"structIPaymentCoordinator.MerkleLeaf\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"rootIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"leafIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"recipientOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setActivationDelay\",\"inputs\":[{\"name\":\"_activationDelay\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setGlobalOperatorCommission\",\"inputs\":[{\"name\":\"_globalCommissionBips\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRecipient\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitRoot\",\"inputs\":[{\"name\":\"root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"paymentsCalculatedUntilTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ActivationDelaySet\",\"inputs\":[{\"name\":\"oldActivationDelay\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"newActivationDelay\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GlobalCommissionBipsSet\",\"inputs\":[{\"name\":\"oldGlobalCommissionBips\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"newGlobalCommissionBips\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PaymentClaimed\",\"inputs\":[{\"name\":\"leaf\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIPaymentCoordinator.MerkleLeaf\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PaymentUpdaterSet\",\"inputs\":[{\"name\":\"oldPaymentUpdater\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPaymentUpdater\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RangePaymentCreated\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"rangePayment\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIPaymentCoordinator.RangePayment\",\"components\":[{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"weights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"duration\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RecipientSet\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RootSubmitted\",\"inputs\":[{\"name\":\"root\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"paymentsCalculatedUntilTimestamp\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"activatedAfter\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false}]",
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

// MAXPAYMENTDURATION is a free data retrieval call binding the contract method 0xee619597.
//
// Solidity: function MAX_PAYMENT_DURATION() view returns(uint32)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorCaller) MAXPAYMENTDURATION(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractIPaymentCoordinator.contract.Call(opts, &out, "MAX_PAYMENT_DURATION")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// MAXPAYMENTDURATION is a free data retrieval call binding the contract method 0xee619597.
//
// Solidity: function MAX_PAYMENT_DURATION() view returns(uint32)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorSession) MAXPAYMENTDURATION() (uint32, error) {
	return _ContractIPaymentCoordinator.Contract.MAXPAYMENTDURATION(&_ContractIPaymentCoordinator.CallOpts)
}

// MAXPAYMENTDURATION is a free data retrieval call binding the contract method 0xee619597.
//
// Solidity: function MAX_PAYMENT_DURATION() view returns(uint32)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorCallerSession) MAXPAYMENTDURATION() (uint32, error) {
	return _ContractIPaymentCoordinator.Contract.MAXPAYMENTDURATION(&_ContractIPaymentCoordinator.CallOpts)
}

// ActivationDelay is a free data retrieval call binding the contract method 0x3a8c0786.
//
// Solidity: function activationDelay() view returns(uint32)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorCaller) ActivationDelay(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractIPaymentCoordinator.contract.Call(opts, &out, "activationDelay")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ActivationDelay is a free data retrieval call binding the contract method 0x3a8c0786.
//
// Solidity: function activationDelay() view returns(uint32)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorSession) ActivationDelay() (uint32, error) {
	return _ContractIPaymentCoordinator.Contract.ActivationDelay(&_ContractIPaymentCoordinator.CallOpts)
}

// ActivationDelay is a free data retrieval call binding the contract method 0x3a8c0786.
//
// Solidity: function activationDelay() view returns(uint32)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorCallerSession) ActivationDelay() (uint32, error) {
	return _ContractIPaymentCoordinator.Contract.ActivationDelay(&_ContractIPaymentCoordinator.CallOpts)
}

// CalculationIntervalSeconds is a free data retrieval call binding the contract method 0x169bde2b.
//
// Solidity: function calculationIntervalSeconds() view returns(uint32)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorCaller) CalculationIntervalSeconds(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractIPaymentCoordinator.contract.Call(opts, &out, "calculationIntervalSeconds")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// CalculationIntervalSeconds is a free data retrieval call binding the contract method 0x169bde2b.
//
// Solidity: function calculationIntervalSeconds() view returns(uint32)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorSession) CalculationIntervalSeconds() (uint32, error) {
	return _ContractIPaymentCoordinator.Contract.CalculationIntervalSeconds(&_ContractIPaymentCoordinator.CallOpts)
}

// CalculationIntervalSeconds is a free data retrieval call binding the contract method 0x169bde2b.
//
// Solidity: function calculationIntervalSeconds() view returns(uint32)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorCallerSession) CalculationIntervalSeconds() (uint32, error) {
	return _ContractIPaymentCoordinator.Contract.CalculationIntervalSeconds(&_ContractIPaymentCoordinator.CallOpts)
}

// CheckClaim is a free data retrieval call binding the contract method 0xf82b28f3.
//
// Solidity: function checkClaim(((address,uint256,address),uint32,uint32,bytes) claim) view returns(bool)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorCaller) CheckClaim(opts *bind.CallOpts, claim IPaymentCoordinatorPaymentMerkleClaim) (bool, error) {
	var out []interface{}
	err := _ContractIPaymentCoordinator.contract.Call(opts, &out, "checkClaim", claim)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckClaim is a free data retrieval call binding the contract method 0xf82b28f3.
//
// Solidity: function checkClaim(((address,uint256,address),uint32,uint32,bytes) claim) view returns(bool)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorSession) CheckClaim(claim IPaymentCoordinatorPaymentMerkleClaim) (bool, error) {
	return _ContractIPaymentCoordinator.Contract.CheckClaim(&_ContractIPaymentCoordinator.CallOpts, claim)
}

// CheckClaim is a free data retrieval call binding the contract method 0xf82b28f3.
//
// Solidity: function checkClaim(((address,uint256,address),uint32,uint32,bytes) claim) view returns(bool)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorCallerSession) CheckClaim(claim IPaymentCoordinatorPaymentMerkleClaim) (bool, error) {
	return _ContractIPaymentCoordinator.Contract.CheckClaim(&_ContractIPaymentCoordinator.CallOpts, claim)
}

// CumulativeClaimed is a free data retrieval call binding the contract method 0x865c6953.
//
// Solidity: function cumulativeClaimed(address recipient, address token) view returns(uint256)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorCaller) CumulativeClaimed(opts *bind.CallOpts, recipient common.Address, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ContractIPaymentCoordinator.contract.Call(opts, &out, "cumulativeClaimed", recipient, token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CumulativeClaimed is a free data retrieval call binding the contract method 0x865c6953.
//
// Solidity: function cumulativeClaimed(address recipient, address token) view returns(uint256)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorSession) CumulativeClaimed(recipient common.Address, token common.Address) (*big.Int, error) {
	return _ContractIPaymentCoordinator.Contract.CumulativeClaimed(&_ContractIPaymentCoordinator.CallOpts, recipient, token)
}

// CumulativeClaimed is a free data retrieval call binding the contract method 0x865c6953.
//
// Solidity: function cumulativeClaimed(address recipient, address token) view returns(uint256)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorCallerSession) CumulativeClaimed(recipient common.Address, token common.Address) (*big.Int, error) {
	return _ContractIPaymentCoordinator.Contract.CumulativeClaimed(&_ContractIPaymentCoordinator.CallOpts, recipient, token)
}

// FindLeafHash is a free data retrieval call binding the contract method 0x21fea2c6.
//
// Solidity: function findLeafHash((address,uint256,address) leaf) view returns(bytes32)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorCaller) FindLeafHash(opts *bind.CallOpts, leaf IPaymentCoordinatorMerkleLeaf) ([32]byte, error) {
	var out []interface{}
	err := _ContractIPaymentCoordinator.contract.Call(opts, &out, "findLeafHash", leaf)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// FindLeafHash is a free data retrieval call binding the contract method 0x21fea2c6.
//
// Solidity: function findLeafHash((address,uint256,address) leaf) view returns(bytes32)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorSession) FindLeafHash(leaf IPaymentCoordinatorMerkleLeaf) ([32]byte, error) {
	return _ContractIPaymentCoordinator.Contract.FindLeafHash(&_ContractIPaymentCoordinator.CallOpts, leaf)
}

// FindLeafHash is a free data retrieval call binding the contract method 0x21fea2c6.
//
// Solidity: function findLeafHash((address,uint256,address) leaf) view returns(bytes32)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorCallerSession) FindLeafHash(leaf IPaymentCoordinatorMerkleLeaf) ([32]byte, error) {
	return _ContractIPaymentCoordinator.Contract.FindLeafHash(&_ContractIPaymentCoordinator.CallOpts, leaf)
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

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIPaymentCoordinator.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorSession) Owner() (common.Address, error) {
	return _ContractIPaymentCoordinator.Contract.Owner(&_ContractIPaymentCoordinator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorCallerSession) Owner() (common.Address, error) {
	return _ContractIPaymentCoordinator.Contract.Owner(&_ContractIPaymentCoordinator.CallOpts)
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

// RecipientOf is a free data retrieval call binding the contract method 0x695b7f23.
//
// Solidity: function recipientOf(address account) view returns(address)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorCaller) RecipientOf(opts *bind.CallOpts, account common.Address) (common.Address, error) {
	var out []interface{}
	err := _ContractIPaymentCoordinator.contract.Call(opts, &out, "recipientOf", account)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RecipientOf is a free data retrieval call binding the contract method 0x695b7f23.
//
// Solidity: function recipientOf(address account) view returns(address)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorSession) RecipientOf(account common.Address) (common.Address, error) {
	return _ContractIPaymentCoordinator.Contract.RecipientOf(&_ContractIPaymentCoordinator.CallOpts, account)
}

// RecipientOf is a free data retrieval call binding the contract method 0x695b7f23.
//
// Solidity: function recipientOf(address account) view returns(address)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorCallerSession) RecipientOf(account common.Address) (common.Address, error) {
	return _ContractIPaymentCoordinator.Contract.RecipientOf(&_ContractIPaymentCoordinator.CallOpts, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd2ba69c.
//
// Solidity: function initialize(address initialOwner, uint32 _activationDelay, uint16 _globalCommissionBips) returns()
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address, _activationDelay uint32, _globalCommissionBips uint16) (*types.Transaction, error) {
	return _ContractIPaymentCoordinator.contract.Transact(opts, "initialize", initialOwner, _activationDelay, _globalCommissionBips)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd2ba69c.
//
// Solidity: function initialize(address initialOwner, uint32 _activationDelay, uint16 _globalCommissionBips) returns()
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorSession) Initialize(initialOwner common.Address, _activationDelay uint32, _globalCommissionBips uint16) (*types.Transaction, error) {
	return _ContractIPaymentCoordinator.Contract.Initialize(&_ContractIPaymentCoordinator.TransactOpts, initialOwner, _activationDelay, _globalCommissionBips)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd2ba69c.
//
// Solidity: function initialize(address initialOwner, uint32 _activationDelay, uint16 _globalCommissionBips) returns()
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorTransactorSession) Initialize(initialOwner common.Address, _activationDelay uint32, _globalCommissionBips uint16) (*types.Transaction, error) {
	return _ContractIPaymentCoordinator.Contract.Initialize(&_ContractIPaymentCoordinator.TransactOpts, initialOwner, _activationDelay, _globalCommissionBips)
}

// PayAllForRange is a paid mutator transaction binding the contract method 0xc61c8ff0.
//
// Solidity: function payAllForRange((address[],uint256[],address,uint256,uint256,uint256) rangePayment) returns()
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorTransactor) PayAllForRange(opts *bind.TransactOpts, rangePayment IPaymentCoordinatorRangePayment) (*types.Transaction, error) {
	return _ContractIPaymentCoordinator.contract.Transact(opts, "payAllForRange", rangePayment)
}

// PayAllForRange is a paid mutator transaction binding the contract method 0xc61c8ff0.
//
// Solidity: function payAllForRange((address[],uint256[],address,uint256,uint256,uint256) rangePayment) returns()
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorSession) PayAllForRange(rangePayment IPaymentCoordinatorRangePayment) (*types.Transaction, error) {
	return _ContractIPaymentCoordinator.Contract.PayAllForRange(&_ContractIPaymentCoordinator.TransactOpts, rangePayment)
}

// PayAllForRange is a paid mutator transaction binding the contract method 0xc61c8ff0.
//
// Solidity: function payAllForRange((address[],uint256[],address,uint256,uint256,uint256) rangePayment) returns()
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorTransactorSession) PayAllForRange(rangePayment IPaymentCoordinatorRangePayment) (*types.Transaction, error) {
	return _ContractIPaymentCoordinator.Contract.PayAllForRange(&_ContractIPaymentCoordinator.TransactOpts, rangePayment)
}

// PayForRange is a paid mutator transaction binding the contract method 0x7652d39f.
//
// Solidity: function payForRange((address[],uint256[],address,uint256,uint256,uint256) rangePayment) returns()
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorTransactor) PayForRange(opts *bind.TransactOpts, rangePayment IPaymentCoordinatorRangePayment) (*types.Transaction, error) {
	return _ContractIPaymentCoordinator.contract.Transact(opts, "payForRange", rangePayment)
}

// PayForRange is a paid mutator transaction binding the contract method 0x7652d39f.
//
// Solidity: function payForRange((address[],uint256[],address,uint256,uint256,uint256) rangePayment) returns()
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorSession) PayForRange(rangePayment IPaymentCoordinatorRangePayment) (*types.Transaction, error) {
	return _ContractIPaymentCoordinator.Contract.PayForRange(&_ContractIPaymentCoordinator.TransactOpts, rangePayment)
}

// PayForRange is a paid mutator transaction binding the contract method 0x7652d39f.
//
// Solidity: function payForRange((address[],uint256[],address,uint256,uint256,uint256) rangePayment) returns()
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorTransactorSession) PayForRange(rangePayment IPaymentCoordinatorRangePayment) (*types.Transaction, error) {
	return _ContractIPaymentCoordinator.Contract.PayForRange(&_ContractIPaymentCoordinator.TransactOpts, rangePayment)
}

// ProcessClaims is a paid mutator transaction binding the contract method 0x1d3138e5.
//
// Solidity: function processClaims(((address,uint256,address),uint32,uint32,bytes)[] claims) returns()
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorTransactor) ProcessClaims(opts *bind.TransactOpts, claims []IPaymentCoordinatorPaymentMerkleClaim) (*types.Transaction, error) {
	return _ContractIPaymentCoordinator.contract.Transact(opts, "processClaims", claims)
}

// ProcessClaims is a paid mutator transaction binding the contract method 0x1d3138e5.
//
// Solidity: function processClaims(((address,uint256,address),uint32,uint32,bytes)[] claims) returns()
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorSession) ProcessClaims(claims []IPaymentCoordinatorPaymentMerkleClaim) (*types.Transaction, error) {
	return _ContractIPaymentCoordinator.Contract.ProcessClaims(&_ContractIPaymentCoordinator.TransactOpts, claims)
}

// ProcessClaims is a paid mutator transaction binding the contract method 0x1d3138e5.
//
// Solidity: function processClaims(((address,uint256,address),uint32,uint32,bytes)[] claims) returns()
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorTransactorSession) ProcessClaims(claims []IPaymentCoordinatorPaymentMerkleClaim) (*types.Transaction, error) {
	return _ContractIPaymentCoordinator.Contract.ProcessClaims(&_ContractIPaymentCoordinator.TransactOpts, claims)
}

// SetActivationDelay is a paid mutator transaction binding the contract method 0x58baaa3e.
//
// Solidity: function setActivationDelay(uint32 _activationDelay) returns()
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorTransactor) SetActivationDelay(opts *bind.TransactOpts, _activationDelay uint32) (*types.Transaction, error) {
	return _ContractIPaymentCoordinator.contract.Transact(opts, "setActivationDelay", _activationDelay)
}

// SetActivationDelay is a paid mutator transaction binding the contract method 0x58baaa3e.
//
// Solidity: function setActivationDelay(uint32 _activationDelay) returns()
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorSession) SetActivationDelay(_activationDelay uint32) (*types.Transaction, error) {
	return _ContractIPaymentCoordinator.Contract.SetActivationDelay(&_ContractIPaymentCoordinator.TransactOpts, _activationDelay)
}

// SetActivationDelay is a paid mutator transaction binding the contract method 0x58baaa3e.
//
// Solidity: function setActivationDelay(uint32 _activationDelay) returns()
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorTransactorSession) SetActivationDelay(_activationDelay uint32) (*types.Transaction, error) {
	return _ContractIPaymentCoordinator.Contract.SetActivationDelay(&_ContractIPaymentCoordinator.TransactOpts, _activationDelay)
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

// SetRecipient is a paid mutator transaction binding the contract method 0x8bc8407a.
//
// Solidity: function setRecipient(address account, address recipient) returns()
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorTransactor) SetRecipient(opts *bind.TransactOpts, account common.Address, recipient common.Address) (*types.Transaction, error) {
	return _ContractIPaymentCoordinator.contract.Transact(opts, "setRecipient", account, recipient)
}

// SetRecipient is a paid mutator transaction binding the contract method 0x8bc8407a.
//
// Solidity: function setRecipient(address account, address recipient) returns()
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorSession) SetRecipient(account common.Address, recipient common.Address) (*types.Transaction, error) {
	return _ContractIPaymentCoordinator.Contract.SetRecipient(&_ContractIPaymentCoordinator.TransactOpts, account, recipient)
}

// SetRecipient is a paid mutator transaction binding the contract method 0x8bc8407a.
//
// Solidity: function setRecipient(address account, address recipient) returns()
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorTransactorSession) SetRecipient(account common.Address, recipient common.Address) (*types.Transaction, error) {
	return _ContractIPaymentCoordinator.Contract.SetRecipient(&_ContractIPaymentCoordinator.TransactOpts, account, recipient)
}

// SubmitRoot is a paid mutator transaction binding the contract method 0x3efe1db6.
//
// Solidity: function submitRoot(bytes32 root, uint32 paymentsCalculatedUntilTimestamp) returns()
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorTransactor) SubmitRoot(opts *bind.TransactOpts, root [32]byte, paymentsCalculatedUntilTimestamp uint32) (*types.Transaction, error) {
	return _ContractIPaymentCoordinator.contract.Transact(opts, "submitRoot", root, paymentsCalculatedUntilTimestamp)
}

// SubmitRoot is a paid mutator transaction binding the contract method 0x3efe1db6.
//
// Solidity: function submitRoot(bytes32 root, uint32 paymentsCalculatedUntilTimestamp) returns()
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorSession) SubmitRoot(root [32]byte, paymentsCalculatedUntilTimestamp uint32) (*types.Transaction, error) {
	return _ContractIPaymentCoordinator.Contract.SubmitRoot(&_ContractIPaymentCoordinator.TransactOpts, root, paymentsCalculatedUntilTimestamp)
}

// SubmitRoot is a paid mutator transaction binding the contract method 0x3efe1db6.
//
// Solidity: function submitRoot(bytes32 root, uint32 paymentsCalculatedUntilTimestamp) returns()
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorTransactorSession) SubmitRoot(root [32]byte, paymentsCalculatedUntilTimestamp uint32) (*types.Transaction, error) {
	return _ContractIPaymentCoordinator.Contract.SubmitRoot(&_ContractIPaymentCoordinator.TransactOpts, root, paymentsCalculatedUntilTimestamp)
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
	OldActivationDelay uint32
	NewActivationDelay uint32
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterActivationDelaySet is a free log retrieval operation binding the contract event 0xaf557c6c02c208794817a705609cfa935f827312a1adfdd26494b6b95dd2b4b3.
//
// Solidity: event ActivationDelaySet(uint32 oldActivationDelay, uint32 newActivationDelay)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorFilterer) FilterActivationDelaySet(opts *bind.FilterOpts) (*ContractIPaymentCoordinatorActivationDelaySetIterator, error) {

	logs, sub, err := _ContractIPaymentCoordinator.contract.FilterLogs(opts, "ActivationDelaySet")
	if err != nil {
		return nil, err
	}
	return &ContractIPaymentCoordinatorActivationDelaySetIterator{contract: _ContractIPaymentCoordinator.contract, event: "ActivationDelaySet", logs: logs, sub: sub}, nil
}

// WatchActivationDelaySet is a free log subscription operation binding the contract event 0xaf557c6c02c208794817a705609cfa935f827312a1adfdd26494b6b95dd2b4b3.
//
// Solidity: event ActivationDelaySet(uint32 oldActivationDelay, uint32 newActivationDelay)
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

// ParseActivationDelaySet is a log parse operation binding the contract event 0xaf557c6c02c208794817a705609cfa935f827312a1adfdd26494b6b95dd2b4b3.
//
// Solidity: event ActivationDelaySet(uint32 oldActivationDelay, uint32 newActivationDelay)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorFilterer) ParseActivationDelaySet(log types.Log) (*ContractIPaymentCoordinatorActivationDelaySet, error) {
	event := new(ContractIPaymentCoordinatorActivationDelaySet)
	if err := _ContractIPaymentCoordinator.contract.UnpackLog(event, "ActivationDelaySet", log); err != nil {
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
	Leaf IPaymentCoordinatorMerkleLeaf
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterPaymentClaimed is a free log retrieval operation binding the contract event 0x401edde99091dc11fe6dccbcce7adf5bd8f37eef1c3dc2620ed71bac579386b9.
//
// Solidity: event PaymentClaimed((address,uint256,address) leaf)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorFilterer) FilterPaymentClaimed(opts *bind.FilterOpts) (*ContractIPaymentCoordinatorPaymentClaimedIterator, error) {

	logs, sub, err := _ContractIPaymentCoordinator.contract.FilterLogs(opts, "PaymentClaimed")
	if err != nil {
		return nil, err
	}
	return &ContractIPaymentCoordinatorPaymentClaimedIterator{contract: _ContractIPaymentCoordinator.contract, event: "PaymentClaimed", logs: logs, sub: sub}, nil
}

// WatchPaymentClaimed is a free log subscription operation binding the contract event 0x401edde99091dc11fe6dccbcce7adf5bd8f37eef1c3dc2620ed71bac579386b9.
//
// Solidity: event PaymentClaimed((address,uint256,address) leaf)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorFilterer) WatchPaymentClaimed(opts *bind.WatchOpts, sink chan<- *ContractIPaymentCoordinatorPaymentClaimed) (event.Subscription, error) {

	logs, sub, err := _ContractIPaymentCoordinator.contract.WatchLogs(opts, "PaymentClaimed")
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

// ParsePaymentClaimed is a log parse operation binding the contract event 0x401edde99091dc11fe6dccbcce7adf5bd8f37eef1c3dc2620ed71bac579386b9.
//
// Solidity: event PaymentClaimed((address,uint256,address) leaf)
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
	Avs          common.Address
	RangePayment IPaymentCoordinatorRangePayment
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRangePaymentCreated is a free log retrieval operation binding the contract event 0x48334dd4edfcc5d854afde6a0f65f799fe8ab0975a1f76ddb8e89f2984512d18.
//
// Solidity: event RangePaymentCreated(address avs, (address[],uint256[],address,uint256,uint256,uint256) rangePayment)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorFilterer) FilterRangePaymentCreated(opts *bind.FilterOpts) (*ContractIPaymentCoordinatorRangePaymentCreatedIterator, error) {

	logs, sub, err := _ContractIPaymentCoordinator.contract.FilterLogs(opts, "RangePaymentCreated")
	if err != nil {
		return nil, err
	}
	return &ContractIPaymentCoordinatorRangePaymentCreatedIterator{contract: _ContractIPaymentCoordinator.contract, event: "RangePaymentCreated", logs: logs, sub: sub}, nil
}

// WatchRangePaymentCreated is a free log subscription operation binding the contract event 0x48334dd4edfcc5d854afde6a0f65f799fe8ab0975a1f76ddb8e89f2984512d18.
//
// Solidity: event RangePaymentCreated(address avs, (address[],uint256[],address,uint256,uint256,uint256) rangePayment)
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

// ParseRangePaymentCreated is a log parse operation binding the contract event 0x48334dd4edfcc5d854afde6a0f65f799fe8ab0975a1f76ddb8e89f2984512d18.
//
// Solidity: event RangePaymentCreated(address avs, (address[],uint256[],address,uint256,uint256,uint256) rangePayment)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorFilterer) ParseRangePaymentCreated(log types.Log) (*ContractIPaymentCoordinatorRangePaymentCreated, error) {
	event := new(ContractIPaymentCoordinatorRangePaymentCreated)
	if err := _ContractIPaymentCoordinator.contract.UnpackLog(event, "RangePaymentCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIPaymentCoordinatorRecipientSetIterator is returned from FilterRecipientSet and is used to iterate over the raw logs and unpacked data for RecipientSet events raised by the ContractIPaymentCoordinator contract.
type ContractIPaymentCoordinatorRecipientSetIterator struct {
	Event *ContractIPaymentCoordinatorRecipientSet // Event containing the contract specifics and raw log

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
func (it *ContractIPaymentCoordinatorRecipientSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIPaymentCoordinatorRecipientSet)
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
		it.Event = new(ContractIPaymentCoordinatorRecipientSet)
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
func (it *ContractIPaymentCoordinatorRecipientSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIPaymentCoordinatorRecipientSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIPaymentCoordinatorRecipientSet represents a RecipientSet event raised by the ContractIPaymentCoordinator contract.
type ContractIPaymentCoordinatorRecipientSet struct {
	Account   common.Address
	Recipient common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRecipientSet is a free log retrieval operation binding the contract event 0xc1416b5cdab50a9fbc872236e1aa54566c6deb40024e63a4b1737ecacf09d6f9.
//
// Solidity: event RecipientSet(address indexed account, address indexed recipient)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorFilterer) FilterRecipientSet(opts *bind.FilterOpts, account []common.Address, recipient []common.Address) (*ContractIPaymentCoordinatorRecipientSetIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ContractIPaymentCoordinator.contract.FilterLogs(opts, "RecipientSet", accountRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &ContractIPaymentCoordinatorRecipientSetIterator{contract: _ContractIPaymentCoordinator.contract, event: "RecipientSet", logs: logs, sub: sub}, nil
}

// WatchRecipientSet is a free log subscription operation binding the contract event 0xc1416b5cdab50a9fbc872236e1aa54566c6deb40024e63a4b1737ecacf09d6f9.
//
// Solidity: event RecipientSet(address indexed account, address indexed recipient)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorFilterer) WatchRecipientSet(opts *bind.WatchOpts, sink chan<- *ContractIPaymentCoordinatorRecipientSet, account []common.Address, recipient []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ContractIPaymentCoordinator.contract.WatchLogs(opts, "RecipientSet", accountRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIPaymentCoordinatorRecipientSet)
				if err := _ContractIPaymentCoordinator.contract.UnpackLog(event, "RecipientSet", log); err != nil {
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

// ParseRecipientSet is a log parse operation binding the contract event 0xc1416b5cdab50a9fbc872236e1aa54566c6deb40024e63a4b1737ecacf09d6f9.
//
// Solidity: event RecipientSet(address indexed account, address indexed recipient)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorFilterer) ParseRecipientSet(log types.Log) (*ContractIPaymentCoordinatorRecipientSet, error) {
	event := new(ContractIPaymentCoordinatorRecipientSet)
	if err := _ContractIPaymentCoordinator.contract.UnpackLog(event, "RecipientSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIPaymentCoordinatorRootSubmittedIterator is returned from FilterRootSubmitted and is used to iterate over the raw logs and unpacked data for RootSubmitted events raised by the ContractIPaymentCoordinator contract.
type ContractIPaymentCoordinatorRootSubmittedIterator struct {
	Event *ContractIPaymentCoordinatorRootSubmitted // Event containing the contract specifics and raw log

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
func (it *ContractIPaymentCoordinatorRootSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIPaymentCoordinatorRootSubmitted)
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
		it.Event = new(ContractIPaymentCoordinatorRootSubmitted)
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
func (it *ContractIPaymentCoordinatorRootSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIPaymentCoordinatorRootSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIPaymentCoordinatorRootSubmitted represents a RootSubmitted event raised by the ContractIPaymentCoordinator contract.
type ContractIPaymentCoordinatorRootSubmitted struct {
	Root                             [32]byte
	PaymentsCalculatedUntilTimestamp uint32
	ActivatedAfter                   uint32
	Raw                              types.Log // Blockchain specific contextual infos
}

// FilterRootSubmitted is a free log retrieval operation binding the contract event 0x262191a0e015e84c4074af7ac4d2305db1490bf60340fbd04afa74cb37bcbdf1.
//
// Solidity: event RootSubmitted(bytes32 root, uint32 paymentsCalculatedUntilTimestamp, uint32 activatedAfter)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorFilterer) FilterRootSubmitted(opts *bind.FilterOpts) (*ContractIPaymentCoordinatorRootSubmittedIterator, error) {

	logs, sub, err := _ContractIPaymentCoordinator.contract.FilterLogs(opts, "RootSubmitted")
	if err != nil {
		return nil, err
	}
	return &ContractIPaymentCoordinatorRootSubmittedIterator{contract: _ContractIPaymentCoordinator.contract, event: "RootSubmitted", logs: logs, sub: sub}, nil
}

// WatchRootSubmitted is a free log subscription operation binding the contract event 0x262191a0e015e84c4074af7ac4d2305db1490bf60340fbd04afa74cb37bcbdf1.
//
// Solidity: event RootSubmitted(bytes32 root, uint32 paymentsCalculatedUntilTimestamp, uint32 activatedAfter)
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorFilterer) WatchRootSubmitted(opts *bind.WatchOpts, sink chan<- *ContractIPaymentCoordinatorRootSubmitted) (event.Subscription, error) {

	logs, sub, err := _ContractIPaymentCoordinator.contract.WatchLogs(opts, "RootSubmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIPaymentCoordinatorRootSubmitted)
				if err := _ContractIPaymentCoordinator.contract.UnpackLog(event, "RootSubmitted", log); err != nil {
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
func (_ContractIPaymentCoordinator *ContractIPaymentCoordinatorFilterer) ParseRootSubmitted(log types.Log) (*ContractIPaymentCoordinatorRootSubmitted, error) {
	event := new(ContractIPaymentCoordinatorRootSubmitted)
	if err := _ContractIPaymentCoordinator.contract.UnpackLog(event, "RootSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
