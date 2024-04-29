// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindingSkateAVS

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

// ISignatureUtilsSignatureWithSaltAndExpiry is an auto generated low-level Go binding around an user-defined struct.
type ISignatureUtilsSignatureWithSaltAndExpiry struct {
	Signature []byte
	Salt      [32]byte
	Expiry    *big.Int
}

// BindingSkateAVSMetaData contains all meta data concerning the BindingSkateAVS contract.
var BindingSkateAVSMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_avsDirectory\",\"type\":\"address\",\"internalType\":\"contractIAVSDirectory\"},{\"name\":\"_registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"_stakeRegistry\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"},{\"name\":\"_taskManager\",\"type\":\"address\",\"internalType\":\"contractISkateTaskManager\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"avsDirectory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deregisterOperatorFromAVS\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"freezeOperator\",\"inputs\":[{\"name\":\"operatorAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getOperatorRestakedStrategies\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRestakeableStrategies\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerOperatorToAVS\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerOperatorWithAVS\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"taskManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractISkateTaskManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateAVSMetadataURI\",\"inputs\":[{\"name\":\"_metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false}]",
	Bin: "0x6101006040523480156200001257600080fd5b50604051620016653803806200166583398101604081905262000035916200014f565b6001600160a01b0380851660c052808416608052821660a0528383836200005b62000074565b5050506001600160a01b031660e05250620001b7915050565b600054610100900460ff1615620000e15760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff908116101562000134576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b03811681146200014c57600080fd5b50565b600080600080608085870312156200016657600080fd5b8451620001738162000136565b6020860151909450620001868162000136565b6040860151909350620001998162000136565b6060860151909250620001ac8162000136565b939692955090935050565b60805160a05160c05160e05161140c620002596000396000818161018801526106b301526000818160f90152818161075f0152818161087b01526108fa0152600081816103d20152818161052e015281816105c5015281816109f701528181610b7b0152610c1a0152600081816101fd0152818161028c0152818161030c015281816107d70152818161081f015281816109350152610ad6015261140c6000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c80639926ee7d116100715780639926ee7d1461015d578063a364f4da14610170578063a50a640e14610183578063a98fb355146101aa578063e481af9d146101bd578063f2fde38b146101c557600080fd5b806333cfb7b7146100b957806338c8ee64146100e25780636b3aa72e146100f7578063715018a6146101315780638303c4f4146101395780638da5cb5b1461014c575b600080fd5b6100cc6100c7366004610f24565b6101d8565b6040516100d99190610f48565b60405180910390f35b6100f56100f0366004610f24565b6106a8565b005b7f00000000000000000000000000000000000000000000000000000000000000005b6040516001600160a01b0390911681526020016100d9565b6100f5610734565b6100f561014736600461104a565b610748565b6033546001600160a01b0316610119565b6100f561016b36600461104a565b6107cc565b6100f561017e366004610f24565b610814565b6101197f000000000000000000000000000000000000000000000000000000000000000081565b6100f56101b83660046110f5565b6108db565b6100cc61092f565b6100f56101d3366004610f24565b610cf9565b6040516309aa152760e11b81526001600160a01b0382811660048301526060916000917f000000000000000000000000000000000000000000000000000000000000000016906313542a4e90602401602060405180830381865afa158015610244573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102689190611146565b60405163871ef04960e01b8152600481018290529091506000906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063871ef04990602401602060405180830381865afa1580156102d3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102f7919061115f565b90506001600160c01b038116158061039157507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610368573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061038c9190611188565b60ff16155b156103ad57505060408051600081526020810190915292915050565b60006103c1826001600160c01b0316610d6f565b90506000805b8251811015610497577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316633ca5a5f5848381518110610411576104116111ab565b01602001516040516001600160e01b031960e084901b16815260f89190911c6004820152602401602060405180830381865afa158015610455573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104799190611146565b61048390836111d7565b91508061048f816111ef565b9150506103c7565b5060008167ffffffffffffffff8111156104b3576104b3610f95565b6040519080825280602002602001820160405280156104dc578160200160208202803683370190505b5090506000805b845181101561069b576000858281518110610500576105006111ab565b0160200151604051633ca5a5f560e01b815260f89190911c6004820181905291506000906001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690633ca5a5f590602401602060405180830381865afa158015610575573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105999190611146565b905060005b81811015610685576040516356e4026d60e11b815260ff84166004820152602481018290527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063adc804da906044016040805180830381865afa158015610613573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610637919061120a565b6000015186868151811061064d5761064d6111ab565b6001600160a01b03909216602092830291909101909101528461066f816111ef565b955050808061067d906111ef565b91505061059e565b5050508080610693906111ef565b9150506104e3565b5090979650505050505050565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146107315760405162461bcd60e51b8152602060048201526024808201527f6f6e6c79536b6174655461736b4d616e616765723a20696e76616c6964206361604482015263363632b960e11b60648201526084015b60405180910390fd5b50565b61073c610e32565b6107466000610e8c565b565b604051639926ee7d60e01b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690639926ee7d9061079690859085906004016112c7565b600060405180830381600087803b1580156107b057600080fd5b505af11580156107c4573d6000803e3d6000fd5b505050505050565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146107485760405162461bcd60e51b815260040161072890611312565b336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461085c5760405162461bcd60e51b815260040161072890611312565b6040516351b27a6d60e11b81526001600160a01b0382811660048301527f0000000000000000000000000000000000000000000000000000000000000000169063a364f4da906024015b600060405180830381600087803b1580156108c057600080fd5b505af11580156108d4573d6000803e3d6000fd5b5050505050565b6108e3610e32565b60405163a98fb35560e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063a98fb355906108a690849060040161138a565b606060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610991573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109b59190611188565b60ff169050806109d357505060408051600081526020810190915290565b6000805b82811015610a8857604051633ca5a5f560e01b815260ff821660048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690633ca5a5f590602401602060405180830381865afa158015610a46573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a6a9190611146565b610a7490836111d7565b915080610a80816111ef565b9150506109d7565b5060008167ffffffffffffffff811115610aa457610aa4610f95565b604051908082528060200260200182016040528015610acd578160200160208202803683370190505b5090506000805b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610b32573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b569190611188565b60ff16811015610cef57604051633ca5a5f560e01b815260ff821660048201526000907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690633ca5a5f590602401602060405180830381865afa158015610bca573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610bee9190611146565b905060005b81811015610cda576040516356e4026d60e11b815260ff84166004820152602481018290527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063adc804da906044016040805180830381865afa158015610c68573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c8c919061120a565b60000151858581518110610ca257610ca26111ab565b6001600160a01b039092166020928302919091019091015283610cc4816111ef565b9450508080610cd2906111ef565b915050610bf3565b50508080610ce7906111ef565b915050610ad4565b5090949350505050565b610d01610e32565b6001600160a01b038116610d665760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610728565b61073181610e8c565b6060600080610d7d84610ede565b61ffff1667ffffffffffffffff811115610d9957610d99610f95565b6040519080825280601f01601f191660200182016040528015610dc3576020820181803683370190505b5090506000805b825182108015610ddb575061010081105b15610cef576001811b935085841615610e22578060f81b838381518110610e0457610e046111ab565b60200101906001600160f81b031916908160001a9053508160010191505b610e2b816111ef565b9050610dca565b6033546001600160a01b031633146107465760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610728565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6000805b8215610f0957610ef360018461139d565b9092169180610f01816113b4565b915050610ee2565b92915050565b6001600160a01b038116811461073157600080fd5b600060208284031215610f3657600080fd5b8135610f4181610f0f565b9392505050565b6020808252825182820181905260009190848201906040850190845b81811015610f895783516001600160a01b031683529284019291840191600101610f64565b50909695505050505050565b634e487b7160e01b600052604160045260246000fd5b6040516060810167ffffffffffffffff81118282101715610fce57610fce610f95565b60405290565b600067ffffffffffffffff80841115610fef57610fef610f95565b604051601f8501601f19908116603f0116810190828211818310171561101757611017610f95565b8160405280935085815286868601111561103057600080fd5b858560208301376000602087830101525050509392505050565b6000806040838503121561105d57600080fd5b823561106881610f0f565b9150602083013567ffffffffffffffff8082111561108557600080fd5b908401906060828703121561109957600080fd5b6110a1610fab565b8235828111156110b057600080fd5b83019150601f820187136110c357600080fd5b6110d287833560208501610fd4565b815260208301356020820152604083013560408201528093505050509250929050565b60006020828403121561110757600080fd5b813567ffffffffffffffff81111561111e57600080fd5b8201601f8101841361112f57600080fd5b61113e84823560208401610fd4565b949350505050565b60006020828403121561115857600080fd5b5051919050565b60006020828403121561117157600080fd5b81516001600160c01b0381168114610f4157600080fd5b60006020828403121561119a57600080fd5b815160ff81168114610f4157600080fd5b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b600082198211156111ea576111ea6111c1565b500190565b6000600019821415611203576112036111c1565b5060010190565b60006040828403121561121c57600080fd5b6040516040810181811067ffffffffffffffff8211171561123f5761123f610f95565b604052825161124d81610f0f565b815260208301516bffffffffffffffffffffffff8116811461126e57600080fd5b60208201529392505050565b6000815180845260005b818110156112a057602081850181015186830182015201611284565b818111156112b2576000602083870101525b50601f01601f19169290920160200192915050565b60018060a01b03831681526040602082015260008251606060408401526112f160a084018261127a565b90506020840151606084015260408401516080840152809150509392505050565b60208082526052908201527f536572766963654d616e61676572426173652e6f6e6c7952656769737472794360408201527f6f6f7264696e61746f723a2063616c6c6572206973206e6f742074686520726560608201527133b4b9ba393c9031b7b7b93234b730ba37b960711b608082015260a00190565b602081526000610f41602083018461127a565b6000828210156113af576113af6111c1565b500390565b600061ffff808316818114156113cc576113cc6111c1565b600101939250505056fea26469706673582212208b5b27940b3fc3bcdf7e0286e2b91fcf6e38843b8373aecb0de9363ead1454b164736f6c634300080c0033",
}

// BindingSkateAVSABI is the input ABI used to generate the binding from.
// Deprecated: Use BindingSkateAVSMetaData.ABI instead.
var BindingSkateAVSABI = BindingSkateAVSMetaData.ABI

// BindingSkateAVSBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BindingSkateAVSMetaData.Bin instead.
var BindingSkateAVSBin = BindingSkateAVSMetaData.Bin

// DeployBindingSkateAVS deploys a new Ethereum contract, binding an instance of BindingSkateAVS to it.
func DeployBindingSkateAVS(auth *bind.TransactOpts, backend bind.ContractBackend, _avsDirectory common.Address, _registryCoordinator common.Address, _stakeRegistry common.Address, _taskManager common.Address) (common.Address, *types.Transaction, *BindingSkateAVS, error) {
	parsed, err := BindingSkateAVSMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BindingSkateAVSBin), backend, _avsDirectory, _registryCoordinator, _stakeRegistry, _taskManager)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BindingSkateAVS{BindingSkateAVSCaller: BindingSkateAVSCaller{contract: contract}, BindingSkateAVSTransactor: BindingSkateAVSTransactor{contract: contract}, BindingSkateAVSFilterer: BindingSkateAVSFilterer{contract: contract}}, nil
}

// BindingSkateAVS is an auto generated Go binding around an Ethereum contract.
type BindingSkateAVS struct {
	BindingSkateAVSCaller     // Read-only binding to the contract
	BindingSkateAVSTransactor // Write-only binding to the contract
	BindingSkateAVSFilterer   // Log filterer for contract events
}

// BindingSkateAVSCaller is an auto generated read-only Go binding around an Ethereum contract.
type BindingSkateAVSCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingSkateAVSTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BindingSkateAVSTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingSkateAVSFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BindingSkateAVSFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingSkateAVSSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BindingSkateAVSSession struct {
	Contract     *BindingSkateAVS  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BindingSkateAVSCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BindingSkateAVSCallerSession struct {
	Contract *BindingSkateAVSCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// BindingSkateAVSTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BindingSkateAVSTransactorSession struct {
	Contract     *BindingSkateAVSTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// BindingSkateAVSRaw is an auto generated low-level Go binding around an Ethereum contract.
type BindingSkateAVSRaw struct {
	Contract *BindingSkateAVS // Generic contract binding to access the raw methods on
}

// BindingSkateAVSCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BindingSkateAVSCallerRaw struct {
	Contract *BindingSkateAVSCaller // Generic read-only contract binding to access the raw methods on
}

// BindingSkateAVSTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BindingSkateAVSTransactorRaw struct {
	Contract *BindingSkateAVSTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBindingSkateAVS creates a new instance of BindingSkateAVS, bound to a specific deployed contract.
func NewBindingSkateAVS(address common.Address, backend bind.ContractBackend) (*BindingSkateAVS, error) {
	contract, err := bindBindingSkateAVS(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BindingSkateAVS{BindingSkateAVSCaller: BindingSkateAVSCaller{contract: contract}, BindingSkateAVSTransactor: BindingSkateAVSTransactor{contract: contract}, BindingSkateAVSFilterer: BindingSkateAVSFilterer{contract: contract}}, nil
}

// NewBindingSkateAVSCaller creates a new read-only instance of BindingSkateAVS, bound to a specific deployed contract.
func NewBindingSkateAVSCaller(address common.Address, caller bind.ContractCaller) (*BindingSkateAVSCaller, error) {
	contract, err := bindBindingSkateAVS(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BindingSkateAVSCaller{contract: contract}, nil
}

// NewBindingSkateAVSTransactor creates a new write-only instance of BindingSkateAVS, bound to a specific deployed contract.
func NewBindingSkateAVSTransactor(address common.Address, transactor bind.ContractTransactor) (*BindingSkateAVSTransactor, error) {
	contract, err := bindBindingSkateAVS(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BindingSkateAVSTransactor{contract: contract}, nil
}

// NewBindingSkateAVSFilterer creates a new log filterer instance of BindingSkateAVS, bound to a specific deployed contract.
func NewBindingSkateAVSFilterer(address common.Address, filterer bind.ContractFilterer) (*BindingSkateAVSFilterer, error) {
	contract, err := bindBindingSkateAVS(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BindingSkateAVSFilterer{contract: contract}, nil
}

// bindBindingSkateAVS binds a generic wrapper to an already deployed contract.
func bindBindingSkateAVS(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BindingSkateAVSMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BindingSkateAVS *BindingSkateAVSRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BindingSkateAVS.Contract.BindingSkateAVSCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BindingSkateAVS *BindingSkateAVSRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BindingSkateAVS.Contract.BindingSkateAVSTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BindingSkateAVS *BindingSkateAVSRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BindingSkateAVS.Contract.BindingSkateAVSTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BindingSkateAVS *BindingSkateAVSCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BindingSkateAVS.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BindingSkateAVS *BindingSkateAVSTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BindingSkateAVS.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BindingSkateAVS *BindingSkateAVSTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BindingSkateAVS.Contract.contract.Transact(opts, method, params...)
}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_BindingSkateAVS *BindingSkateAVSCaller) AvsDirectory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BindingSkateAVS.contract.Call(opts, &out, "avsDirectory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_BindingSkateAVS *BindingSkateAVSSession) AvsDirectory() (common.Address, error) {
	return _BindingSkateAVS.Contract.AvsDirectory(&_BindingSkateAVS.CallOpts)
}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_BindingSkateAVS *BindingSkateAVSCallerSession) AvsDirectory() (common.Address, error) {
	return _BindingSkateAVS.Contract.AvsDirectory(&_BindingSkateAVS.CallOpts)
}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_BindingSkateAVS *BindingSkateAVSCaller) GetOperatorRestakedStrategies(opts *bind.CallOpts, operator common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _BindingSkateAVS.contract.Call(opts, &out, "getOperatorRestakedStrategies", operator)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_BindingSkateAVS *BindingSkateAVSSession) GetOperatorRestakedStrategies(operator common.Address) ([]common.Address, error) {
	return _BindingSkateAVS.Contract.GetOperatorRestakedStrategies(&_BindingSkateAVS.CallOpts, operator)
}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_BindingSkateAVS *BindingSkateAVSCallerSession) GetOperatorRestakedStrategies(operator common.Address) ([]common.Address, error) {
	return _BindingSkateAVS.Contract.GetOperatorRestakedStrategies(&_BindingSkateAVS.CallOpts, operator)
}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_BindingSkateAVS *BindingSkateAVSCaller) GetRestakeableStrategies(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _BindingSkateAVS.contract.Call(opts, &out, "getRestakeableStrategies")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_BindingSkateAVS *BindingSkateAVSSession) GetRestakeableStrategies() ([]common.Address, error) {
	return _BindingSkateAVS.Contract.GetRestakeableStrategies(&_BindingSkateAVS.CallOpts)
}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_BindingSkateAVS *BindingSkateAVSCallerSession) GetRestakeableStrategies() ([]common.Address, error) {
	return _BindingSkateAVS.Contract.GetRestakeableStrategies(&_BindingSkateAVS.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BindingSkateAVS *BindingSkateAVSCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BindingSkateAVS.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BindingSkateAVS *BindingSkateAVSSession) Owner() (common.Address, error) {
	return _BindingSkateAVS.Contract.Owner(&_BindingSkateAVS.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BindingSkateAVS *BindingSkateAVSCallerSession) Owner() (common.Address, error) {
	return _BindingSkateAVS.Contract.Owner(&_BindingSkateAVS.CallOpts)
}

// TaskManager is a free data retrieval call binding the contract method 0xa50a640e.
//
// Solidity: function taskManager() view returns(address)
func (_BindingSkateAVS *BindingSkateAVSCaller) TaskManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BindingSkateAVS.contract.Call(opts, &out, "taskManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TaskManager is a free data retrieval call binding the contract method 0xa50a640e.
//
// Solidity: function taskManager() view returns(address)
func (_BindingSkateAVS *BindingSkateAVSSession) TaskManager() (common.Address, error) {
	return _BindingSkateAVS.Contract.TaskManager(&_BindingSkateAVS.CallOpts)
}

// TaskManager is a free data retrieval call binding the contract method 0xa50a640e.
//
// Solidity: function taskManager() view returns(address)
func (_BindingSkateAVS *BindingSkateAVSCallerSession) TaskManager() (common.Address, error) {
	return _BindingSkateAVS.Contract.TaskManager(&_BindingSkateAVS.CallOpts)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_BindingSkateAVS *BindingSkateAVSTransactor) DeregisterOperatorFromAVS(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _BindingSkateAVS.contract.Transact(opts, "deregisterOperatorFromAVS", operator)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_BindingSkateAVS *BindingSkateAVSSession) DeregisterOperatorFromAVS(operator common.Address) (*types.Transaction, error) {
	return _BindingSkateAVS.Contract.DeregisterOperatorFromAVS(&_BindingSkateAVS.TransactOpts, operator)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_BindingSkateAVS *BindingSkateAVSTransactorSession) DeregisterOperatorFromAVS(operator common.Address) (*types.Transaction, error) {
	return _BindingSkateAVS.Contract.DeregisterOperatorFromAVS(&_BindingSkateAVS.TransactOpts, operator)
}

// FreezeOperator is a paid mutator transaction binding the contract method 0x38c8ee64.
//
// Solidity: function freezeOperator(address operatorAddr) returns()
func (_BindingSkateAVS *BindingSkateAVSTransactor) FreezeOperator(opts *bind.TransactOpts, operatorAddr common.Address) (*types.Transaction, error) {
	return _BindingSkateAVS.contract.Transact(opts, "freezeOperator", operatorAddr)
}

// FreezeOperator is a paid mutator transaction binding the contract method 0x38c8ee64.
//
// Solidity: function freezeOperator(address operatorAddr) returns()
func (_BindingSkateAVS *BindingSkateAVSSession) FreezeOperator(operatorAddr common.Address) (*types.Transaction, error) {
	return _BindingSkateAVS.Contract.FreezeOperator(&_BindingSkateAVS.TransactOpts, operatorAddr)
}

// FreezeOperator is a paid mutator transaction binding the contract method 0x38c8ee64.
//
// Solidity: function freezeOperator(address operatorAddr) returns()
func (_BindingSkateAVS *BindingSkateAVSTransactorSession) FreezeOperator(operatorAddr common.Address) (*types.Transaction, error) {
	return _BindingSkateAVS.Contract.FreezeOperator(&_BindingSkateAVS.TransactOpts, operatorAddr)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_BindingSkateAVS *BindingSkateAVSTransactor) RegisterOperatorToAVS(opts *bind.TransactOpts, operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _BindingSkateAVS.contract.Transact(opts, "registerOperatorToAVS", operator, operatorSignature)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_BindingSkateAVS *BindingSkateAVSSession) RegisterOperatorToAVS(operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _BindingSkateAVS.Contract.RegisterOperatorToAVS(&_BindingSkateAVS.TransactOpts, operator, operatorSignature)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_BindingSkateAVS *BindingSkateAVSTransactorSession) RegisterOperatorToAVS(operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _BindingSkateAVS.Contract.RegisterOperatorToAVS(&_BindingSkateAVS.TransactOpts, operator, operatorSignature)
}

// RegisterOperatorWithAVS is a paid mutator transaction binding the contract method 0x8303c4f4.
//
// Solidity: function registerOperatorWithAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_BindingSkateAVS *BindingSkateAVSTransactor) RegisterOperatorWithAVS(opts *bind.TransactOpts, operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _BindingSkateAVS.contract.Transact(opts, "registerOperatorWithAVS", operator, operatorSignature)
}

// RegisterOperatorWithAVS is a paid mutator transaction binding the contract method 0x8303c4f4.
//
// Solidity: function registerOperatorWithAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_BindingSkateAVS *BindingSkateAVSSession) RegisterOperatorWithAVS(operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _BindingSkateAVS.Contract.RegisterOperatorWithAVS(&_BindingSkateAVS.TransactOpts, operator, operatorSignature)
}

// RegisterOperatorWithAVS is a paid mutator transaction binding the contract method 0x8303c4f4.
//
// Solidity: function registerOperatorWithAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_BindingSkateAVS *BindingSkateAVSTransactorSession) RegisterOperatorWithAVS(operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _BindingSkateAVS.Contract.RegisterOperatorWithAVS(&_BindingSkateAVS.TransactOpts, operator, operatorSignature)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BindingSkateAVS *BindingSkateAVSTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BindingSkateAVS.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BindingSkateAVS *BindingSkateAVSSession) RenounceOwnership() (*types.Transaction, error) {
	return _BindingSkateAVS.Contract.RenounceOwnership(&_BindingSkateAVS.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BindingSkateAVS *BindingSkateAVSTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BindingSkateAVS.Contract.RenounceOwnership(&_BindingSkateAVS.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BindingSkateAVS *BindingSkateAVSTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BindingSkateAVS.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BindingSkateAVS *BindingSkateAVSSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BindingSkateAVS.Contract.TransferOwnership(&_BindingSkateAVS.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BindingSkateAVS *BindingSkateAVSTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BindingSkateAVS.Contract.TransferOwnership(&_BindingSkateAVS.TransactOpts, newOwner)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string _metadataURI) returns()
func (_BindingSkateAVS *BindingSkateAVSTransactor) UpdateAVSMetadataURI(opts *bind.TransactOpts, _metadataURI string) (*types.Transaction, error) {
	return _BindingSkateAVS.contract.Transact(opts, "updateAVSMetadataURI", _metadataURI)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string _metadataURI) returns()
func (_BindingSkateAVS *BindingSkateAVSSession) UpdateAVSMetadataURI(_metadataURI string) (*types.Transaction, error) {
	return _BindingSkateAVS.Contract.UpdateAVSMetadataURI(&_BindingSkateAVS.TransactOpts, _metadataURI)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string _metadataURI) returns()
func (_BindingSkateAVS *BindingSkateAVSTransactorSession) UpdateAVSMetadataURI(_metadataURI string) (*types.Transaction, error) {
	return _BindingSkateAVS.Contract.UpdateAVSMetadataURI(&_BindingSkateAVS.TransactOpts, _metadataURI)
}

// BindingSkateAVSInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the BindingSkateAVS contract.
type BindingSkateAVSInitializedIterator struct {
	Event *BindingSkateAVSInitialized // Event containing the contract specifics and raw log

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
func (it *BindingSkateAVSInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingSkateAVSInitialized)
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
		it.Event = new(BindingSkateAVSInitialized)
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
func (it *BindingSkateAVSInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingSkateAVSInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingSkateAVSInitialized represents a Initialized event raised by the BindingSkateAVS contract.
type BindingSkateAVSInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_BindingSkateAVS *BindingSkateAVSFilterer) FilterInitialized(opts *bind.FilterOpts) (*BindingSkateAVSInitializedIterator, error) {

	logs, sub, err := _BindingSkateAVS.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BindingSkateAVSInitializedIterator{contract: _BindingSkateAVS.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_BindingSkateAVS *BindingSkateAVSFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BindingSkateAVSInitialized) (event.Subscription, error) {

	logs, sub, err := _BindingSkateAVS.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingSkateAVSInitialized)
				if err := _BindingSkateAVS.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_BindingSkateAVS *BindingSkateAVSFilterer) ParseInitialized(log types.Log) (*BindingSkateAVSInitialized, error) {
	event := new(BindingSkateAVSInitialized)
	if err := _BindingSkateAVS.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingSkateAVSOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BindingSkateAVS contract.
type BindingSkateAVSOwnershipTransferredIterator struct {
	Event *BindingSkateAVSOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BindingSkateAVSOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingSkateAVSOwnershipTransferred)
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
		it.Event = new(BindingSkateAVSOwnershipTransferred)
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
func (it *BindingSkateAVSOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingSkateAVSOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingSkateAVSOwnershipTransferred represents a OwnershipTransferred event raised by the BindingSkateAVS contract.
type BindingSkateAVSOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BindingSkateAVS *BindingSkateAVSFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BindingSkateAVSOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BindingSkateAVS.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BindingSkateAVSOwnershipTransferredIterator{contract: _BindingSkateAVS.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BindingSkateAVS *BindingSkateAVSFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BindingSkateAVSOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BindingSkateAVS.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingSkateAVSOwnershipTransferred)
				if err := _BindingSkateAVS.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BindingSkateAVS *BindingSkateAVSFilterer) ParseOwnershipTransferred(log types.Log) (*BindingSkateAVSOwnershipTransferred, error) {
	event := new(BindingSkateAVSOwnershipTransferred)
	if err := _BindingSkateAVS.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
