// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindingSkateGateway

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

// BindingSkateGatewayMetaData contains all meta data concerning the BindingSkateGateway contract.
var BindingSkateGatewayMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x6080604052348015600f57600080fd5b50603f80601d6000396000f3fe6080604052600080fdfea2646970667358221220764739d0cb7e4d83e3a390d098b01c5467754235867ecea0e3abb557eb52593b64736f6c63430008140033",
}

// BindingSkateGatewayABI is the input ABI used to generate the binding from.
// Deprecated: Use BindingSkateGatewayMetaData.ABI instead.
var BindingSkateGatewayABI = BindingSkateGatewayMetaData.ABI

// BindingSkateGatewayBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BindingSkateGatewayMetaData.Bin instead.
var BindingSkateGatewayBin = BindingSkateGatewayMetaData.Bin

// DeployBindingSkateGateway deploys a new Ethereum contract, binding an instance of BindingSkateGateway to it.
func DeployBindingSkateGateway(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BindingSkateGateway, error) {
	parsed, err := BindingSkateGatewayMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BindingSkateGatewayBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BindingSkateGateway{BindingSkateGatewayCaller: BindingSkateGatewayCaller{contract: contract}, BindingSkateGatewayTransactor: BindingSkateGatewayTransactor{contract: contract}, BindingSkateGatewayFilterer: BindingSkateGatewayFilterer{contract: contract}}, nil
}

// BindingSkateGateway is an auto generated Go binding around an Ethereum contract.
type BindingSkateGateway struct {
	BindingSkateGatewayCaller     // Read-only binding to the contract
	BindingSkateGatewayTransactor // Write-only binding to the contract
	BindingSkateGatewayFilterer   // Log filterer for contract events
}

// BindingSkateGatewayCaller is an auto generated read-only Go binding around an Ethereum contract.
type BindingSkateGatewayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingSkateGatewayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BindingSkateGatewayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingSkateGatewayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BindingSkateGatewayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingSkateGatewaySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BindingSkateGatewaySession struct {
	Contract     *BindingSkateGateway // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// BindingSkateGatewayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BindingSkateGatewayCallerSession struct {
	Contract *BindingSkateGatewayCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// BindingSkateGatewayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BindingSkateGatewayTransactorSession struct {
	Contract     *BindingSkateGatewayTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// BindingSkateGatewayRaw is an auto generated low-level Go binding around an Ethereum contract.
type BindingSkateGatewayRaw struct {
	Contract *BindingSkateGateway // Generic contract binding to access the raw methods on
}

// BindingSkateGatewayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BindingSkateGatewayCallerRaw struct {
	Contract *BindingSkateGatewayCaller // Generic read-only contract binding to access the raw methods on
}

// BindingSkateGatewayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BindingSkateGatewayTransactorRaw struct {
	Contract *BindingSkateGatewayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBindingSkateGateway creates a new instance of BindingSkateGateway, bound to a specific deployed contract.
func NewBindingSkateGateway(address common.Address, backend bind.ContractBackend) (*BindingSkateGateway, error) {
	contract, err := bindBindingSkateGateway(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BindingSkateGateway{BindingSkateGatewayCaller: BindingSkateGatewayCaller{contract: contract}, BindingSkateGatewayTransactor: BindingSkateGatewayTransactor{contract: contract}, BindingSkateGatewayFilterer: BindingSkateGatewayFilterer{contract: contract}}, nil
}

// NewBindingSkateGatewayCaller creates a new read-only instance of BindingSkateGateway, bound to a specific deployed contract.
func NewBindingSkateGatewayCaller(address common.Address, caller bind.ContractCaller) (*BindingSkateGatewayCaller, error) {
	contract, err := bindBindingSkateGateway(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BindingSkateGatewayCaller{contract: contract}, nil
}

// NewBindingSkateGatewayTransactor creates a new write-only instance of BindingSkateGateway, bound to a specific deployed contract.
func NewBindingSkateGatewayTransactor(address common.Address, transactor bind.ContractTransactor) (*BindingSkateGatewayTransactor, error) {
	contract, err := bindBindingSkateGateway(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BindingSkateGatewayTransactor{contract: contract}, nil
}

// NewBindingSkateGatewayFilterer creates a new log filterer instance of BindingSkateGateway, bound to a specific deployed contract.
func NewBindingSkateGatewayFilterer(address common.Address, filterer bind.ContractFilterer) (*BindingSkateGatewayFilterer, error) {
	contract, err := bindBindingSkateGateway(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BindingSkateGatewayFilterer{contract: contract}, nil
}

// bindBindingSkateGateway binds a generic wrapper to an already deployed contract.
func bindBindingSkateGateway(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BindingSkateGatewayMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BindingSkateGateway *BindingSkateGatewayRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BindingSkateGateway.Contract.BindingSkateGatewayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BindingSkateGateway *BindingSkateGatewayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BindingSkateGateway.Contract.BindingSkateGatewayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BindingSkateGateway *BindingSkateGatewayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BindingSkateGateway.Contract.BindingSkateGatewayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BindingSkateGateway *BindingSkateGatewayCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BindingSkateGateway.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BindingSkateGateway *BindingSkateGatewayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BindingSkateGateway.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BindingSkateGateway *BindingSkateGatewayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BindingSkateGateway.Contract.contract.Transact(opts, method, params...)
}
