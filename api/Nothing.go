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

// NothingMetaData contains all meta data concerning the Nothing contract.
var NothingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]",
	Bin: "0x6080604052348015600f57600080fd5b50603f80601d6000396000f3fe6080604052600080fdfea2646970667358221220d1b84972956a56ab82acb41e0c0c2332cf1fc742601dd4be78965861b4755d2b64736f6c63430008140033",
}

// NothingABI is the input ABI used to generate the binding from.
// Deprecated: Use NothingMetaData.ABI instead.
var NothingABI = NothingMetaData.ABI

// NothingBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use NothingMetaData.Bin instead.
var NothingBin = NothingMetaData.Bin

// DeployNothing deploys a new Ethereum contract, binding an instance of Nothing to it.
func DeployNothing(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Nothing, error) {
	parsed, err := NothingMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(NothingBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Nothing{NothingCaller: NothingCaller{contract: contract}, NothingTransactor: NothingTransactor{contract: contract}, NothingFilterer: NothingFilterer{contract: contract}}, nil
}

// Nothing is an auto generated Go binding around an Ethereum contract.
type Nothing struct {
	NothingCaller     // Read-only binding to the contract
	NothingTransactor // Write-only binding to the contract
	NothingFilterer   // Log filterer for contract events
}

// NothingCaller is an auto generated read-only Go binding around an Ethereum contract.
type NothingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NothingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NothingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NothingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NothingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NothingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NothingSession struct {
	Contract     *Nothing          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NothingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NothingCallerSession struct {
	Contract *NothingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// NothingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NothingTransactorSession struct {
	Contract     *NothingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// NothingRaw is an auto generated low-level Go binding around an Ethereum contract.
type NothingRaw struct {
	Contract *Nothing // Generic contract binding to access the raw methods on
}

// NothingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NothingCallerRaw struct {
	Contract *NothingCaller // Generic read-only contract binding to access the raw methods on
}

// NothingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NothingTransactorRaw struct {
	Contract *NothingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNothing creates a new instance of Nothing, bound to a specific deployed contract.
func NewNothing(address common.Address, backend bind.ContractBackend) (*Nothing, error) {
	contract, err := bindNothing(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Nothing{NothingCaller: NothingCaller{contract: contract}, NothingTransactor: NothingTransactor{contract: contract}, NothingFilterer: NothingFilterer{contract: contract}}, nil
}

// NewNothingCaller creates a new read-only instance of Nothing, bound to a specific deployed contract.
func NewNothingCaller(address common.Address, caller bind.ContractCaller) (*NothingCaller, error) {
	contract, err := bindNothing(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NothingCaller{contract: contract}, nil
}

// NewNothingTransactor creates a new write-only instance of Nothing, bound to a specific deployed contract.
func NewNothingTransactor(address common.Address, transactor bind.ContractTransactor) (*NothingTransactor, error) {
	contract, err := bindNothing(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NothingTransactor{contract: contract}, nil
}

// NewNothingFilterer creates a new log filterer instance of Nothing, bound to a specific deployed contract.
func NewNothingFilterer(address common.Address, filterer bind.ContractFilterer) (*NothingFilterer, error) {
	contract, err := bindNothing(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NothingFilterer{contract: contract}, nil
}

// bindNothing binds a generic wrapper to an already deployed contract.
func bindNothing(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NothingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Nothing *NothingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Nothing.Contract.NothingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Nothing *NothingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nothing.Contract.NothingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Nothing *NothingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Nothing.Contract.NothingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Nothing *NothingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Nothing.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Nothing *NothingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nothing.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Nothing *NothingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Nothing.Contract.contract.Transact(opts, method, params...)
}
