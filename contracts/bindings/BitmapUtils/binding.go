// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindingBitmapUtils

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

// BindingBitmapUtilsMetaData contains all meta data concerning the BindingBitmapUtils contract.
var BindingBitmapUtilsMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122088645f035b9ba6d0823c1c8152cdc14b0df31ea619f9cc31c2e941f5b89cdb3864736f6c634300080c0033",
}

// BindingBitmapUtilsABI is the input ABI used to generate the binding from.
// Deprecated: Use BindingBitmapUtilsMetaData.ABI instead.
var BindingBitmapUtilsABI = BindingBitmapUtilsMetaData.ABI

// BindingBitmapUtilsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BindingBitmapUtilsMetaData.Bin instead.
var BindingBitmapUtilsBin = BindingBitmapUtilsMetaData.Bin

// DeployBindingBitmapUtils deploys a new Ethereum contract, binding an instance of BindingBitmapUtils to it.
func DeployBindingBitmapUtils(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BindingBitmapUtils, error) {
	parsed, err := BindingBitmapUtilsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BindingBitmapUtilsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BindingBitmapUtils{BindingBitmapUtilsCaller: BindingBitmapUtilsCaller{contract: contract}, BindingBitmapUtilsTransactor: BindingBitmapUtilsTransactor{contract: contract}, BindingBitmapUtilsFilterer: BindingBitmapUtilsFilterer{contract: contract}}, nil
}

// BindingBitmapUtils is an auto generated Go binding around an Ethereum contract.
type BindingBitmapUtils struct {
	BindingBitmapUtilsCaller     // Read-only binding to the contract
	BindingBitmapUtilsTransactor // Write-only binding to the contract
	BindingBitmapUtilsFilterer   // Log filterer for contract events
}

// BindingBitmapUtilsCaller is an auto generated read-only Go binding around an Ethereum contract.
type BindingBitmapUtilsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingBitmapUtilsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BindingBitmapUtilsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingBitmapUtilsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BindingBitmapUtilsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingBitmapUtilsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BindingBitmapUtilsSession struct {
	Contract     *BindingBitmapUtils // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BindingBitmapUtilsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BindingBitmapUtilsCallerSession struct {
	Contract *BindingBitmapUtilsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// BindingBitmapUtilsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BindingBitmapUtilsTransactorSession struct {
	Contract     *BindingBitmapUtilsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// BindingBitmapUtilsRaw is an auto generated low-level Go binding around an Ethereum contract.
type BindingBitmapUtilsRaw struct {
	Contract *BindingBitmapUtils // Generic contract binding to access the raw methods on
}

// BindingBitmapUtilsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BindingBitmapUtilsCallerRaw struct {
	Contract *BindingBitmapUtilsCaller // Generic read-only contract binding to access the raw methods on
}

// BindingBitmapUtilsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BindingBitmapUtilsTransactorRaw struct {
	Contract *BindingBitmapUtilsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBindingBitmapUtils creates a new instance of BindingBitmapUtils, bound to a specific deployed contract.
func NewBindingBitmapUtils(address common.Address, backend bind.ContractBackend) (*BindingBitmapUtils, error) {
	contract, err := bindBindingBitmapUtils(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BindingBitmapUtils{BindingBitmapUtilsCaller: BindingBitmapUtilsCaller{contract: contract}, BindingBitmapUtilsTransactor: BindingBitmapUtilsTransactor{contract: contract}, BindingBitmapUtilsFilterer: BindingBitmapUtilsFilterer{contract: contract}}, nil
}

// NewBindingBitmapUtilsCaller creates a new read-only instance of BindingBitmapUtils, bound to a specific deployed contract.
func NewBindingBitmapUtilsCaller(address common.Address, caller bind.ContractCaller) (*BindingBitmapUtilsCaller, error) {
	contract, err := bindBindingBitmapUtils(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BindingBitmapUtilsCaller{contract: contract}, nil
}

// NewBindingBitmapUtilsTransactor creates a new write-only instance of BindingBitmapUtils, bound to a specific deployed contract.
func NewBindingBitmapUtilsTransactor(address common.Address, transactor bind.ContractTransactor) (*BindingBitmapUtilsTransactor, error) {
	contract, err := bindBindingBitmapUtils(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BindingBitmapUtilsTransactor{contract: contract}, nil
}

// NewBindingBitmapUtilsFilterer creates a new log filterer instance of BindingBitmapUtils, bound to a specific deployed contract.
func NewBindingBitmapUtilsFilterer(address common.Address, filterer bind.ContractFilterer) (*BindingBitmapUtilsFilterer, error) {
	contract, err := bindBindingBitmapUtils(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BindingBitmapUtilsFilterer{contract: contract}, nil
}

// bindBindingBitmapUtils binds a generic wrapper to an already deployed contract.
func bindBindingBitmapUtils(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BindingBitmapUtilsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BindingBitmapUtils *BindingBitmapUtilsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BindingBitmapUtils.Contract.BindingBitmapUtilsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BindingBitmapUtils *BindingBitmapUtilsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BindingBitmapUtils.Contract.BindingBitmapUtilsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BindingBitmapUtils *BindingBitmapUtilsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BindingBitmapUtils.Contract.BindingBitmapUtilsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BindingBitmapUtils *BindingBitmapUtilsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BindingBitmapUtils.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BindingBitmapUtils *BindingBitmapUtilsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BindingBitmapUtils.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BindingBitmapUtils *BindingBitmapUtilsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BindingBitmapUtils.Contract.contract.Transact(opts, method, params...)
}
