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

// Pigfox2MetaData contains all meta data concerning the Pigfox2 contract.
var Pigfox2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token1\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner1\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount1\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_token2\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner2\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount2\",\"type\":\"uint256\"}],\"name\":\"swap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061041b806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c80637ddf332014610030575b600080fd5b61004361003e36600461033d565b610045565b005b858585858585336001600160a01b038616148061006a5750336001600160a01b038316145b6100ac5760405162461bcd60e51b815260206004820152600e60248201526d139bdd08185d5d1a1bdc9a5e995960921b60448201526064015b60405180910390fd5b604051636eb1769f60e11b81526001600160a01b03868116600483015230602483015285919088169063dd62ed3e90604401602060405180830381865afa1580156100fb573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061011f91906103a3565b101561016d5760405162461bcd60e51b815260206004820152601960248201527f546f6b656e203120616c6c6f77616e636520746f6f206c6f770000000000000060448201526064016100a3565b604051636eb1769f60e11b81526001600160a01b03838116600483015230602483015282919085169063dd62ed3e90604401602060405180830381865afa1580156101bc573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101e091906103a3565b101561022e5760405162461bcd60e51b815260206004820152601960248201527f546f6b656e203220616c6c6f77616e636520746f6f206c6f770000000000000060448201526064016100a3565b61023a86868487610254565b61024683838784610254565b505050505050505050505050565b6040516323b872dd60e01b81526001600160a01b038481166004830152838116602483015260448201839052600091908616906323b872dd906064016020604051808303816000875af11580156102af573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102d391906103bc565b90508061031a5760405162461bcd60e51b8152602060048201526015602482015274151bdad95b881d1c985b9cd9995c8819985a5b1959605a1b60448201526064016100a3565b5050505050565b80356001600160a01b038116811461033857600080fd5b919050565b60008060008060008060c0878903121561035657600080fd5b61035f87610321565b955061036d60208801610321565b94506040870135935061038260608801610321565b925061039060808801610321565b915060a087013590509295509295509295565b6000602082840312156103b557600080fd5b5051919050565b6000602082840312156103ce57600080fd5b815180151581146103de57600080fd5b939250505056fea2646970667358221220c5077a422218497523656bacb1d20d265f55ebcbb0ff83d5ef472173493de58164736f6c63430008140033",
}

// Pigfox2ABI is the input ABI used to generate the binding from.
// Deprecated: Use Pigfox2MetaData.ABI instead.
var Pigfox2ABI = Pigfox2MetaData.ABI

// Pigfox2Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Pigfox2MetaData.Bin instead.
var Pigfox2Bin = Pigfox2MetaData.Bin

// DeployPigfox2 deploys a new Ethereum contract, binding an instance of Pigfox2 to it.
func DeployPigfox2(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Pigfox2, error) {
	parsed, err := Pigfox2MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Pigfox2Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Pigfox2{Pigfox2Caller: Pigfox2Caller{contract: contract}, Pigfox2Transactor: Pigfox2Transactor{contract: contract}, Pigfox2Filterer: Pigfox2Filterer{contract: contract}}, nil
}

// Pigfox2 is an auto generated Go binding around an Ethereum contract.
type Pigfox2 struct {
	Pigfox2Caller     // Read-only binding to the contract
	Pigfox2Transactor // Write-only binding to the contract
	Pigfox2Filterer   // Log filterer for contract events
}

// Pigfox2Caller is an auto generated read-only Go binding around an Ethereum contract.
type Pigfox2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Pigfox2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Pigfox2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Pigfox2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Pigfox2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Pigfox2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Pigfox2Session struct {
	Contract     *Pigfox2          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Pigfox2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Pigfox2CallerSession struct {
	Contract *Pigfox2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// Pigfox2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Pigfox2TransactorSession struct {
	Contract     *Pigfox2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// Pigfox2Raw is an auto generated low-level Go binding around an Ethereum contract.
type Pigfox2Raw struct {
	Contract *Pigfox2 // Generic contract binding to access the raw methods on
}

// Pigfox2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Pigfox2CallerRaw struct {
	Contract *Pigfox2Caller // Generic read-only contract binding to access the raw methods on
}

// Pigfox2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Pigfox2TransactorRaw struct {
	Contract *Pigfox2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewPigfox2 creates a new instance of Pigfox2, bound to a specific deployed contract.
func NewPigfox2(address common.Address, backend bind.ContractBackend) (*Pigfox2, error) {
	contract, err := bindPigfox2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pigfox2{Pigfox2Caller: Pigfox2Caller{contract: contract}, Pigfox2Transactor: Pigfox2Transactor{contract: contract}, Pigfox2Filterer: Pigfox2Filterer{contract: contract}}, nil
}

// NewPigfox2Caller creates a new read-only instance of Pigfox2, bound to a specific deployed contract.
func NewPigfox2Caller(address common.Address, caller bind.ContractCaller) (*Pigfox2Caller, error) {
	contract, err := bindPigfox2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Pigfox2Caller{contract: contract}, nil
}

// NewPigfox2Transactor creates a new write-only instance of Pigfox2, bound to a specific deployed contract.
func NewPigfox2Transactor(address common.Address, transactor bind.ContractTransactor) (*Pigfox2Transactor, error) {
	contract, err := bindPigfox2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Pigfox2Transactor{contract: contract}, nil
}

// NewPigfox2Filterer creates a new log filterer instance of Pigfox2, bound to a specific deployed contract.
func NewPigfox2Filterer(address common.Address, filterer bind.ContractFilterer) (*Pigfox2Filterer, error) {
	contract, err := bindPigfox2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Pigfox2Filterer{contract: contract}, nil
}

// bindPigfox2 binds a generic wrapper to an already deployed contract.
func bindPigfox2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Pigfox2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pigfox2 *Pigfox2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pigfox2.Contract.Pigfox2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pigfox2 *Pigfox2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pigfox2.Contract.Pigfox2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pigfox2 *Pigfox2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pigfox2.Contract.Pigfox2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pigfox2 *Pigfox2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pigfox2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pigfox2 *Pigfox2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pigfox2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pigfox2 *Pigfox2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pigfox2.Contract.contract.Transact(opts, method, params...)
}

// Swap is a paid mutator transaction binding the contract method 0x7ddf3320.
//
// Solidity: function swap(address _token1, address _owner1, uint256 _amount1, address _token2, address _owner2, uint256 _amount2) returns()
func (_Pigfox2 *Pigfox2Transactor) Swap(opts *bind.TransactOpts, _token1 common.Address, _owner1 common.Address, _amount1 *big.Int, _token2 common.Address, _owner2 common.Address, _amount2 *big.Int) (*types.Transaction, error) {
	return _Pigfox2.contract.Transact(opts, "swap", _token1, _owner1, _amount1, _token2, _owner2, _amount2)
}

// Swap is a paid mutator transaction binding the contract method 0x7ddf3320.
//
// Solidity: function swap(address _token1, address _owner1, uint256 _amount1, address _token2, address _owner2, uint256 _amount2) returns()
func (_Pigfox2 *Pigfox2Session) Swap(_token1 common.Address, _owner1 common.Address, _amount1 *big.Int, _token2 common.Address, _owner2 common.Address, _amount2 *big.Int) (*types.Transaction, error) {
	return _Pigfox2.Contract.Swap(&_Pigfox2.TransactOpts, _token1, _owner1, _amount1, _token2, _owner2, _amount2)
}

// Swap is a paid mutator transaction binding the contract method 0x7ddf3320.
//
// Solidity: function swap(address _token1, address _owner1, uint256 _amount1, address _token2, address _owner2, uint256 _amount2) returns()
func (_Pigfox2 *Pigfox2TransactorSession) Swap(_token1 common.Address, _owner1 common.Address, _amount1 *big.Int, _token2 common.Address, _owner2 common.Address, _amount2 *big.Int) (*types.Transaction, error) {
	return _Pigfox2.Contract.Swap(&_Pigfox2.TransactOpts, _token1, _owner1, _amount1, _token2, _owner2, _amount2)
}
