// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindingBN254

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

// BindingBN254MetaData contains all meta data concerning the BindingBN254 contract.
var BindingBN254MetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212205d55efa37003e4b21913b72ba0c0770f1cd76c35c541591547eb776316586e6b64736f6c634300080c0033",
}

// BindingBN254ABI is the input ABI used to generate the binding from.
// Deprecated: Use BindingBN254MetaData.ABI instead.
var BindingBN254ABI = BindingBN254MetaData.ABI

// BindingBN254Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BindingBN254MetaData.Bin instead.
var BindingBN254Bin = BindingBN254MetaData.Bin

// DeployBindingBN254 deploys a new Ethereum contract, binding an instance of BindingBN254 to it.
func DeployBindingBN254(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BindingBN254, error) {
	parsed, err := BindingBN254MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BindingBN254Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BindingBN254{BindingBN254Caller: BindingBN254Caller{contract: contract}, BindingBN254Transactor: BindingBN254Transactor{contract: contract}, BindingBN254Filterer: BindingBN254Filterer{contract: contract}}, nil
}

// BindingBN254 is an auto generated Go binding around an Ethereum contract.
type BindingBN254 struct {
	BindingBN254Caller     // Read-only binding to the contract
	BindingBN254Transactor // Write-only binding to the contract
	BindingBN254Filterer   // Log filterer for contract events
}

// BindingBN254Caller is an auto generated read-only Go binding around an Ethereum contract.
type BindingBN254Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingBN254Transactor is an auto generated write-only Go binding around an Ethereum contract.
type BindingBN254Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingBN254Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BindingBN254Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingBN254Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BindingBN254Session struct {
	Contract     *BindingBN254     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BindingBN254CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BindingBN254CallerSession struct {
	Contract *BindingBN254Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// BindingBN254TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BindingBN254TransactorSession struct {
	Contract     *BindingBN254Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// BindingBN254Raw is an auto generated low-level Go binding around an Ethereum contract.
type BindingBN254Raw struct {
	Contract *BindingBN254 // Generic contract binding to access the raw methods on
}

// BindingBN254CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BindingBN254CallerRaw struct {
	Contract *BindingBN254Caller // Generic read-only contract binding to access the raw methods on
}

// BindingBN254TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BindingBN254TransactorRaw struct {
	Contract *BindingBN254Transactor // Generic write-only contract binding to access the raw methods on
}

// NewBindingBN254 creates a new instance of BindingBN254, bound to a specific deployed contract.
func NewBindingBN254(address common.Address, backend bind.ContractBackend) (*BindingBN254, error) {
	contract, err := bindBindingBN254(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BindingBN254{BindingBN254Caller: BindingBN254Caller{contract: contract}, BindingBN254Transactor: BindingBN254Transactor{contract: contract}, BindingBN254Filterer: BindingBN254Filterer{contract: contract}}, nil
}

// NewBindingBN254Caller creates a new read-only instance of BindingBN254, bound to a specific deployed contract.
func NewBindingBN254Caller(address common.Address, caller bind.ContractCaller) (*BindingBN254Caller, error) {
	contract, err := bindBindingBN254(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BindingBN254Caller{contract: contract}, nil
}

// NewBindingBN254Transactor creates a new write-only instance of BindingBN254, bound to a specific deployed contract.
func NewBindingBN254Transactor(address common.Address, transactor bind.ContractTransactor) (*BindingBN254Transactor, error) {
	contract, err := bindBindingBN254(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BindingBN254Transactor{contract: contract}, nil
}

// NewBindingBN254Filterer creates a new log filterer instance of BindingBN254, bound to a specific deployed contract.
func NewBindingBN254Filterer(address common.Address, filterer bind.ContractFilterer) (*BindingBN254Filterer, error) {
	contract, err := bindBindingBN254(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BindingBN254Filterer{contract: contract}, nil
}

// bindBindingBN254 binds a generic wrapper to an already deployed contract.
func bindBindingBN254(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BindingBN254MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BindingBN254 *BindingBN254Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BindingBN254.Contract.BindingBN254Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BindingBN254 *BindingBN254Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BindingBN254.Contract.BindingBN254Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BindingBN254 *BindingBN254Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BindingBN254.Contract.BindingBN254Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BindingBN254 *BindingBN254CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BindingBN254.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BindingBN254 *BindingBN254TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BindingBN254.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BindingBN254 *BindingBN254TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BindingBN254.Contract.contract.Transact(opts, method, params...)
}
