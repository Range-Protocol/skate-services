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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deregisterOperator\",\"inputs\":[{\"name\":\"op\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getMsg\",\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"messages\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"message\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"signer\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"postMsg\",\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"signer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerOperator\",\"inputs\":[{\"name\":\"op\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"}]",
	Bin: "0x608060405234801561001057600080fd5b506040516106fb3803806106fb83398101604081905261002f91610054565b600080546001600160a01b0319166001600160a01b0392909216919091179055610084565b60006020828403121561006657600080fd5b81516001600160a01b038116811461007d57600080fd5b9392505050565b610668806100936000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c80630d80fefd1461005c5780632b82735b146100865780633682a450146100a65780633fcfdc1d146100bb578063d8cf98ca146100ce575b600080fd5b61006f61006a36600461032e565b6100e1565b60405161007d92919061038d565b60405180910390f35b61009961009436600461032e565b61018e565b60405161007d91906103b7565b6100b96100b43660046103ed565b610230565b005b6100b96100c936600461041e565b61026b565b6100b96100dc3660046103ed565b6102f6565b6001602052600090815260409020805481906100fc906104e9565b80601f0160208091040260200160405190810160405280929190818152602001828054610128906104e9565b80156101755780601f1061014a57610100808354040283529160200191610175565b820191906000526020600020905b81548152906001019060200180831161015857829003601f168201915b505050600190930154919250506001600160a01b031682565b60008181526001602052604090208054606091906101ab906104e9565b80601f01602080910402602001604051908101604052809291908181526020018280546101d7906104e9565b80156102245780601f106101f957610100808354040283529160200191610224565b820191906000526020600020905b81548152906001019060200180831161020757829003601f168201915b50505050509050919050565b6000546001600160a01b0316331461024757600080fd5b6001600160a01b03166000908152600260205260409020805460ff19166001179055565b3360009081526002602052604090205460ff16151560011461028c57600080fd5b6040805180820182528381526001600160a01b0383166020808301919091526000868152600190915291909120815181906102c79082610572565b5060209190910151600190910180546001600160a01b0319166001600160a01b03909216919091179055505050565b6000546001600160a01b0316331461030d57600080fd5b6001600160a01b03166000908152600260205260409020805460ff19169055565b60006020828403121561034057600080fd5b5035919050565b6000815180845260005b8181101561036d57602081850181015186830182015201610351565b506000602082860101526020601f19601f83011685010191505092915050565b6040815260006103a06040830185610347565b905060018060a01b03831660208301529392505050565b6020815260006103ca6020830184610347565b9392505050565b80356001600160a01b03811681146103e857600080fd5b919050565b6000602082840312156103ff57600080fd5b6103ca826103d1565b634e487b7160e01b600052604160045260246000fd5b60008060006060848603121561043357600080fd5b83359250602084013567ffffffffffffffff8082111561045257600080fd5b818601915086601f83011261046657600080fd5b81358181111561047857610478610408565b604051601f8201601f19908116603f011681019083821181831017156104a0576104a0610408565b816040528281528960208487010111156104b957600080fd5b8260208601602083013760006020848301015280965050505050506104e0604085016103d1565b90509250925092565b600181811c908216806104fd57607f821691505b60208210810361051d57634e487b7160e01b600052602260045260246000fd5b50919050565b601f82111561056d57600081815260208120601f850160051c8101602086101561054a5750805b601f850160051c820191505b8181101561056957828155600101610556565b5050505b505050565b815167ffffffffffffffff81111561058c5761058c610408565b6105a08161059a84546104e9565b84610523565b602080601f8311600181146105d557600084156105bd5750858301515b600019600386901b1c1916600185901b178555610569565b600085815260208120601f198616915b82811015610604578886015182559484019460019091019084016105e5565b50858210156106225787850151600019600388901b60f8161c191681555b5050505050600190811b0190555056fea264697066735822122071cab6ca385ea364fc174bef9455b2b56276c3b6f27c9a82525e3f2cc1da91b164736f6c63430008140033",
}

// BindingSkateGatewayABI is the input ABI used to generate the binding from.
// Deprecated: Use BindingSkateGatewayMetaData.ABI instead.
var BindingSkateGatewayABI = BindingSkateGatewayMetaData.ABI

// BindingSkateGatewayBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BindingSkateGatewayMetaData.Bin instead.
var BindingSkateGatewayBin = BindingSkateGatewayMetaData.Bin

// DeployBindingSkateGateway deploys a new Ethereum contract, binding an instance of BindingSkateGateway to it.
func DeployBindingSkateGateway(auth *bind.TransactOpts, backend bind.ContractBackend, _owner common.Address) (common.Address, *types.Transaction, *BindingSkateGateway, error) {
	parsed, err := BindingSkateGatewayMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BindingSkateGatewayBin), backend, _owner)
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

// GetMsg is a free data retrieval call binding the contract method 0x2b82735b.
//
// Solidity: function getMsg(uint256 taskId) view returns(string)
func (_BindingSkateGateway *BindingSkateGatewayCaller) GetMsg(opts *bind.CallOpts, taskId *big.Int) (string, error) {
	var out []interface{}
	err := _BindingSkateGateway.contract.Call(opts, &out, "getMsg", taskId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetMsg is a free data retrieval call binding the contract method 0x2b82735b.
//
// Solidity: function getMsg(uint256 taskId) view returns(string)
func (_BindingSkateGateway *BindingSkateGatewaySession) GetMsg(taskId *big.Int) (string, error) {
	return _BindingSkateGateway.Contract.GetMsg(&_BindingSkateGateway.CallOpts, taskId)
}

// GetMsg is a free data retrieval call binding the contract method 0x2b82735b.
//
// Solidity: function getMsg(uint256 taskId) view returns(string)
func (_BindingSkateGateway *BindingSkateGatewayCallerSession) GetMsg(taskId *big.Int) (string, error) {
	return _BindingSkateGateway.Contract.GetMsg(&_BindingSkateGateway.CallOpts, taskId)
}

// Messages is a free data retrieval call binding the contract method 0x0d80fefd.
//
// Solidity: function messages(uint256 ) view returns(string message, address signer)
func (_BindingSkateGateway *BindingSkateGatewayCaller) Messages(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Message string
	Signer  common.Address
}, error) {
	var out []interface{}
	err := _BindingSkateGateway.contract.Call(opts, &out, "messages", arg0)

	outstruct := new(struct {
		Message string
		Signer  common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Message = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Signer = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Messages is a free data retrieval call binding the contract method 0x0d80fefd.
//
// Solidity: function messages(uint256 ) view returns(string message, address signer)
func (_BindingSkateGateway *BindingSkateGatewaySession) Messages(arg0 *big.Int) (struct {
	Message string
	Signer  common.Address
}, error) {
	return _BindingSkateGateway.Contract.Messages(&_BindingSkateGateway.CallOpts, arg0)
}

// Messages is a free data retrieval call binding the contract method 0x0d80fefd.
//
// Solidity: function messages(uint256 ) view returns(string message, address signer)
func (_BindingSkateGateway *BindingSkateGatewayCallerSession) Messages(arg0 *big.Int) (struct {
	Message string
	Signer  common.Address
}, error) {
	return _BindingSkateGateway.Contract.Messages(&_BindingSkateGateway.CallOpts, arg0)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0xd8cf98ca.
//
// Solidity: function deregisterOperator(address op) returns()
func (_BindingSkateGateway *BindingSkateGatewayTransactor) DeregisterOperator(opts *bind.TransactOpts, op common.Address) (*types.Transaction, error) {
	return _BindingSkateGateway.contract.Transact(opts, "deregisterOperator", op)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0xd8cf98ca.
//
// Solidity: function deregisterOperator(address op) returns()
func (_BindingSkateGateway *BindingSkateGatewaySession) DeregisterOperator(op common.Address) (*types.Transaction, error) {
	return _BindingSkateGateway.Contract.DeregisterOperator(&_BindingSkateGateway.TransactOpts, op)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0xd8cf98ca.
//
// Solidity: function deregisterOperator(address op) returns()
func (_BindingSkateGateway *BindingSkateGatewayTransactorSession) DeregisterOperator(op common.Address) (*types.Transaction, error) {
	return _BindingSkateGateway.Contract.DeregisterOperator(&_BindingSkateGateway.TransactOpts, op)
}

// PostMsg is a paid mutator transaction binding the contract method 0x3fcfdc1d.
//
// Solidity: function postMsg(uint256 taskId, string message, address signer) returns()
func (_BindingSkateGateway *BindingSkateGatewayTransactor) PostMsg(opts *bind.TransactOpts, taskId *big.Int, message string, signer common.Address) (*types.Transaction, error) {
	return _BindingSkateGateway.contract.Transact(opts, "postMsg", taskId, message, signer)
}

// PostMsg is a paid mutator transaction binding the contract method 0x3fcfdc1d.
//
// Solidity: function postMsg(uint256 taskId, string message, address signer) returns()
func (_BindingSkateGateway *BindingSkateGatewaySession) PostMsg(taskId *big.Int, message string, signer common.Address) (*types.Transaction, error) {
	return _BindingSkateGateway.Contract.PostMsg(&_BindingSkateGateway.TransactOpts, taskId, message, signer)
}

// PostMsg is a paid mutator transaction binding the contract method 0x3fcfdc1d.
//
// Solidity: function postMsg(uint256 taskId, string message, address signer) returns()
func (_BindingSkateGateway *BindingSkateGatewayTransactorSession) PostMsg(taskId *big.Int, message string, signer common.Address) (*types.Transaction, error) {
	return _BindingSkateGateway.Contract.PostMsg(&_BindingSkateGateway.TransactOpts, taskId, message, signer)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x3682a450.
//
// Solidity: function registerOperator(address op) returns()
func (_BindingSkateGateway *BindingSkateGatewayTransactor) RegisterOperator(opts *bind.TransactOpts, op common.Address) (*types.Transaction, error) {
	return _BindingSkateGateway.contract.Transact(opts, "registerOperator", op)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x3682a450.
//
// Solidity: function registerOperator(address op) returns()
func (_BindingSkateGateway *BindingSkateGatewaySession) RegisterOperator(op common.Address) (*types.Transaction, error) {
	return _BindingSkateGateway.Contract.RegisterOperator(&_BindingSkateGateway.TransactOpts, op)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x3682a450.
//
// Solidity: function registerOperator(address op) returns()
func (_BindingSkateGateway *BindingSkateGatewayTransactorSession) RegisterOperator(op common.Address) (*types.Transaction, error) {
	return _BindingSkateGateway.Contract.RegisterOperator(&_BindingSkateGateway.TransactOpts, op)
}
