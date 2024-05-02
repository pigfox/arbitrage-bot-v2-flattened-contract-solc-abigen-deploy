// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package api

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

// BaseMetaData contains all meta data concerning the Base contract.
var BaseMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_y\",\"type\":\"uint256\"}],\"name\":\"add\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"s\",\"type\":\"string\"}],\"name\":\"concat\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getQ\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_q\",\"type\":\"uint256\"}],\"name\":\"setQ\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60c0604052600360809081526261626360e81b60a05260019061002290826100cc565b5034801561002e575f80fd5b5061018b565b634e487b7160e01b5f52604160045260245ffd5b600181811c9082168061005c57607f821691505b60208210810361007a57634e487b7160e01b5f52602260045260245ffd5b50919050565b601f8211156100c757805f5260205f20601f840160051c810160208510156100a55750805b601f840160051c820191505b818110156100c4575f81556001016100b1565b50505b505050565b81516001600160401b038111156100e5576100e5610034565b6100f9816100f38454610048565b84610080565b602080601f83116001811461012c575f84156101155750858301515b5f19600386901b1c1916600185901b178555610183565b5f85815260208120601f198616915b8281101561015a5788860151825594840194600190910190840161013b565b508582101561017757878501515f19600388901b60f8161c191681555b505060018460011b0185555b505050505050565b61032f806101985f395ff3fe608060405234801561000f575f80fd5b506004361061004a575f3560e01c8063771602f71461004e5780637881c417146100745780638f2d812814610094578063e8e4197a146100a8575b5f80fd5b61006161005c3660046100ef565b6100af565b6040519081526020015b60405180910390f35b610087610082366004610123565b6100c3565b60405161006b91906101ce565b6100a66100a2366004610203565b5f55565b005b5f54610061565b5f6100ba828461021a565b90505b92915050565b60606001826040516020016100d9929190610250565b6040516020818303038152906040529050919050565b5f8060408385031215610100575f80fd5b50508035926020909101359150565b634e487b7160e01b5f52604160045260245ffd5b5f60208284031215610133575f80fd5b813567ffffffffffffffff8082111561014a575f80fd5b818401915084601f83011261015d575f80fd5b81358181111561016f5761016f61010f565b604051601f8201601f19908116603f011681019083821181831017156101975761019761010f565b816040528281528760208487010111156101af575f80fd5b826020860160208301375f928101602001929092525095945050505050565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b5f60208284031215610213575f80fd5b5035919050565b808201808211156100bd57634e487b7160e01b5f52601160045260245ffd5b5f81518060208401855e5f93019283525090919050565b5f8084545f60018260011c9150600183168061026d57607f831692505b6020808410820361028c57634e487b7160e01b5f52602260045260245ffd5b8180156102a057600181146102b5576102e0565b60ff19861689528415158502890196506102e0565b5f8b8152602090205f5b868110156102d85781548b8201529085019083016102bf565b505084890196505b5050505050506102f08185610239565b9594505050505056fea264697066735822122038b2b42d1c12f77a7f89c634041f1601e34e56cd9d6e636e14c231a2390f6e6164736f6c63430008190033",
}

// BaseABI is the input ABI used to generate the binding from.
// Deprecated: Use BaseMetaData.ABI instead.
var BaseABI = BaseMetaData.ABI

// BaseBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BaseMetaData.Bin instead.
var BaseBin = BaseMetaData.Bin

// DeployBase deploys a new Ethereum contract, binding an instance of Base to it.
func DeployBase(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Base, error) {
	parsed, err := BaseMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BaseBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Base{BaseCaller: BaseCaller{contract: contract}, BaseTransactor: BaseTransactor{contract: contract}, BaseFilterer: BaseFilterer{contract: contract}}, nil
}

// Base is an auto generated Go binding around an Ethereum contract.
type Base struct {
	BaseCaller     // Read-only binding to the contract
	BaseTransactor // Write-only binding to the contract
	BaseFilterer   // Log filterer for contract events
}

// BaseCaller is an auto generated read-only Go binding around an Ethereum contract.
type BaseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BaseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BaseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BaseSession struct {
	Contract     *Base             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BaseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BaseCallerSession struct {
	Contract *BaseCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BaseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BaseTransactorSession struct {
	Contract     *BaseTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BaseRaw is an auto generated low-level Go binding around an Ethereum contract.
type BaseRaw struct {
	Contract *Base // Generic contract binding to access the raw methods on
}

// BaseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BaseCallerRaw struct {
	Contract *BaseCaller // Generic read-only contract binding to access the raw methods on
}

// BaseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BaseTransactorRaw struct {
	Contract *BaseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBase creates a new instance of Base, bound to a specific deployed contract.
func NewBase(address common.Address, backend bind.ContractBackend) (*Base, error) {
	contract, err := bindBase(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Base{BaseCaller: BaseCaller{contract: contract}, BaseTransactor: BaseTransactor{contract: contract}, BaseFilterer: BaseFilterer{contract: contract}}, nil
}

// NewBaseCaller creates a new read-only instance of Base, bound to a specific deployed contract.
func NewBaseCaller(address common.Address, caller bind.ContractCaller) (*BaseCaller, error) {
	contract, err := bindBase(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BaseCaller{contract: contract}, nil
}

// NewBaseTransactor creates a new write-only instance of Base, bound to a specific deployed contract.
func NewBaseTransactor(address common.Address, transactor bind.ContractTransactor) (*BaseTransactor, error) {
	contract, err := bindBase(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BaseTransactor{contract: contract}, nil
}

// NewBaseFilterer creates a new log filterer instance of Base, bound to a specific deployed contract.
func NewBaseFilterer(address common.Address, filterer bind.ContractFilterer) (*BaseFilterer, error) {
	contract, err := bindBase(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BaseFilterer{contract: contract}, nil
}

// bindBase binds a generic wrapper to an already deployed contract.
func bindBase(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BaseMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Base *BaseRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Base.Contract.BaseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Base *BaseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Base.Contract.BaseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Base *BaseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Base.Contract.BaseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Base *BaseCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Base.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Base *BaseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Base.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Base *BaseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Base.Contract.contract.Transact(opts, method, params...)
}

// Add is a free data retrieval call binding the contract method 0x771602f7.
//
// Solidity: function add(uint256 _x, uint256 _y) pure returns(uint256)
func (_Base *BaseCaller) Add(opts *bind.CallOpts, _x *big.Int, _y *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Base.contract.Call(opts, &out, "add", _x, _y)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Add is a free data retrieval call binding the contract method 0x771602f7.
//
// Solidity: function add(uint256 _x, uint256 _y) pure returns(uint256)
func (_Base *BaseSession) Add(_x *big.Int, _y *big.Int) (*big.Int, error) {
	return _Base.Contract.Add(&_Base.CallOpts, _x, _y)
}

// Add is a free data retrieval call binding the contract method 0x771602f7.
//
// Solidity: function add(uint256 _x, uint256 _y) pure returns(uint256)
func (_Base *BaseCallerSession) Add(_x *big.Int, _y *big.Int) (*big.Int, error) {
	return _Base.Contract.Add(&_Base.CallOpts, _x, _y)
}

// Concat is a free data retrieval call binding the contract method 0x7881c417.
//
// Solidity: function concat(string s) view returns(string)
func (_Base *BaseCaller) Concat(opts *bind.CallOpts, s string) (string, error) {
	var out []interface{}
	err := _Base.contract.Call(opts, &out, "concat", s)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Concat is a free data retrieval call binding the contract method 0x7881c417.
//
// Solidity: function concat(string s) view returns(string)
func (_Base *BaseSession) Concat(s string) (string, error) {
	return _Base.Contract.Concat(&_Base.CallOpts, s)
}

// Concat is a free data retrieval call binding the contract method 0x7881c417.
//
// Solidity: function concat(string s) view returns(string)
func (_Base *BaseCallerSession) Concat(s string) (string, error) {
	return _Base.Contract.Concat(&_Base.CallOpts, s)
}

// GetQ is a free data retrieval call binding the contract method 0xe8e4197a.
//
// Solidity: function getQ() view returns(uint256)
func (_Base *BaseCaller) GetQ(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Base.contract.Call(opts, &out, "getQ")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetQ is a free data retrieval call binding the contract method 0xe8e4197a.
//
// Solidity: function getQ() view returns(uint256)
func (_Base *BaseSession) GetQ() (*big.Int, error) {
	return _Base.Contract.GetQ(&_Base.CallOpts)
}

// GetQ is a free data retrieval call binding the contract method 0xe8e4197a.
//
// Solidity: function getQ() view returns(uint256)
func (_Base *BaseCallerSession) GetQ() (*big.Int, error) {
	return _Base.Contract.GetQ(&_Base.CallOpts)
}

// SetQ is a paid mutator transaction binding the contract method 0x8f2d8128.
//
// Solidity: function setQ(uint256 _q) returns()
func (_Base *BaseTransactor) SetQ(opts *bind.TransactOpts, _q *big.Int) (*types.Transaction, error) {
	return _Base.contract.Transact(opts, "setQ", _q)
}

// SetQ is a paid mutator transaction binding the contract method 0x8f2d8128.
//
// Solidity: function setQ(uint256 _q) returns()
func (_Base *BaseSession) SetQ(_q *big.Int) (*types.Transaction, error) {
	return _Base.Contract.SetQ(&_Base.TransactOpts, _q)
}

// SetQ is a paid mutator transaction binding the contract method 0x8f2d8128.
//
// Solidity: function setQ(uint256 _q) returns()
func (_Base *BaseTransactorSession) SetQ(_q *big.Int) (*types.Transaction, error) {
	return _Base.Contract.SetQ(&_Base.TransactOpts, _q)
}
