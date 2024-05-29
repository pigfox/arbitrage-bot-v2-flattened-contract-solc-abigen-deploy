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

// PigfoxMetaData contains all meta data concerning the Pigfox contract.
var PigfoxMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_router0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_router1\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token1\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"slippageTolerancePctInt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadlineSecs\",\"type\":\"uint256\"}],\"name\":\"Swap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f80fd5b505f80546001600160a01b0319163317905561089e8061002d5f395ff3fe608060405234801561000f575f80fd5b5060043610610029575f3560e01c8063dba43ee91461002d575b5f80fd5b61004061003b3660046105d4565b610042565b005b5f546001600160a01b0316331461008c5760405162461bcd60e51b81526020600482015260096024820152682737ba1037bbb732b960b91b60448201526064015b60405180910390fd5b600180546001600160a01b03199081166001600160a01b038a8116918217909355600280549092168984161790915560405163095ea7b360e01b8152600481019190915260248101859052869186919083169063095ea7b3906044016020604051808303815f875af1158015610104573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610128919061063e565b6101745760405162461bcd60e51b815260206004820181905260248201527f746f6b656e30206f6e20726f757465723020617070726f7665206661696c65646044820152606401610083565b60025460405163095ea7b360e01b81526001600160a01b039182166004820152602481018790529082169063095ea7b3906044016020604051808303815f875af11580156101c4573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906101e8919061063e565b6102345760405162461bcd60e51b815260206004820181905260248201527f746f6b656e31206f6e20726f757465723120617070726f7665206661696c65646044820152606401610083565b6040805160028082526060820183525f9260208301908036833701905050905087815f8151811061026757610267610678565b60200260200101906001600160a01b031690816001600160a01b031681525050868160018151811061029b5761029b610678565b6001600160a01b03928316602091820292909201015260015460405163d06ca61f60e01b81525f92919091169063d06ca61f906102de908a9086906004016106cf565b5f60405180830381865afa1580156102f8573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261031f91908101906106ef565b90505f606461032e88826107cb565b8360018151811061034157610341610678565b602002602001015161035391906107e4565b61035d91906107fb565b6001549091506001600160a01b03166338ed17398983863061037f8c4261081a565b6040518663ffffffff1660e01b815260040161039f95949392919061082d565b5f604051808303815f875af11580156103ba573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526103e191908101906106ef565b5060408051600280825260608201835283925f9291906020830190803683370190505090508a815f8151811061041957610419610678565b60200260200101906001600160a01b031690816001600160a01b0316815250508b8160018151811061044d5761044d610678565b6001600160a01b03928316602091820292909201015260025460405163d06ca61f60e01b81525f92919091169063d06ca61f9061049090869086906004016106cf565b5f60405180830381865afa1580156104aa573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526104d191908101906106ef565b90505f60646104e08c826107cb565b836001815181106104f3576104f3610678565b602002602001015161050591906107e4565b61050f91906107fb565b905060025f9054906101000a90046001600160a01b03166001600160a01b03166338ed1739858386308f42610544919061081a565b6040518663ffffffff1660e01b815260040161056495949392919061082d565b5f604051808303815f875af115801561057f573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526105a691908101906106ef565b5050505050505050505050505050505050565b80356001600160a01b03811681146105cf575f80fd5b919050565b5f805f805f805f60e0888a0312156105ea575f80fd5b6105f3886105b9565b9650610601602089016105b9565b955061060f604089016105b9565b945061061d606089016105b9565b9699959850939660808101359560a0820135955060c0909101359350915050565b5f6020828403121561064e575f80fd5b8151801515811461065d575f80fd5b9392505050565b634e487b7160e01b5f52604160045260245ffd5b634e487b7160e01b5f52603260045260245ffd5b5f8151808452602084019350602083015f5b828110156106c55781516001600160a01b031686526020958601959091019060010161069e565b5093949350505050565b828152604060208201525f6106e7604083018461068c565b949350505050565b5f602082840312156106ff575f80fd5b815167ffffffffffffffff811115610715575f80fd5b8201601f81018413610725575f80fd5b805167ffffffffffffffff81111561073f5761073f610664565b8060051b604051601f19603f830116810181811067ffffffffffffffff8211171561076c5761076c610664565b604052918252602081840181019290810187841115610789575f80fd5b6020850194505b838510156107ac57845180825260209586019590935001610790565b509695505050505050565b634e487b7160e01b5f52601160045260245ffd5b818103818111156107de576107de6107b7565b92915050565b80820281158282048414176107de576107de6107b7565b5f8261081557634e487b7160e01b5f52601260045260245ffd5b500490565b808201808211156107de576107de6107b7565b85815284602082015260a060408201525f61084b60a083018661068c565b6001600160a01b039490941660608301525060800152939250505056fea26469706673582212200e7d509c916f932b14dad377b32ad5f188fd70123e58f7380397c76237465aea64736f6c634300081a0033",
}

// PigfoxABI is the input ABI used to generate the binding from.
// Deprecated: Use PigfoxMetaData.ABI instead.
var PigfoxABI = PigfoxMetaData.ABI

// PigfoxBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PigfoxMetaData.Bin instead.
var PigfoxBin = PigfoxMetaData.Bin

// DeployPigfox deploys a new Ethereum contract, binding an instance of Pigfox to it.
func DeployPigfox(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Pigfox, error) {
	parsed, err := PigfoxMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PigfoxBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Pigfox{PigfoxCaller: PigfoxCaller{contract: contract}, PigfoxTransactor: PigfoxTransactor{contract: contract}, PigfoxFilterer: PigfoxFilterer{contract: contract}}, nil
}

// Pigfox is an auto generated Go binding around an Ethereum contract.
type Pigfox struct {
	PigfoxCaller     // Read-only binding to the contract
	PigfoxTransactor // Write-only binding to the contract
	PigfoxFilterer   // Log filterer for contract events
}

// PigfoxCaller is an auto generated read-only Go binding around an Ethereum contract.
type PigfoxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PigfoxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PigfoxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PigfoxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PigfoxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PigfoxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PigfoxSession struct {
	Contract     *Pigfox           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PigfoxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PigfoxCallerSession struct {
	Contract *PigfoxCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PigfoxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PigfoxTransactorSession struct {
	Contract     *PigfoxTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PigfoxRaw is an auto generated low-level Go binding around an Ethereum contract.
type PigfoxRaw struct {
	Contract *Pigfox // Generic contract binding to access the raw methods on
}

// PigfoxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PigfoxCallerRaw struct {
	Contract *PigfoxCaller // Generic read-only contract binding to access the raw methods on
}

// PigfoxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PigfoxTransactorRaw struct {
	Contract *PigfoxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPigfox creates a new instance of Pigfox, bound to a specific deployed contract.
func NewPigfox(address common.Address, backend bind.ContractBackend) (*Pigfox, error) {
	contract, err := bindPigfox(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pigfox{PigfoxCaller: PigfoxCaller{contract: contract}, PigfoxTransactor: PigfoxTransactor{contract: contract}, PigfoxFilterer: PigfoxFilterer{contract: contract}}, nil
}

// NewPigfoxCaller creates a new read-only instance of Pigfox, bound to a specific deployed contract.
func NewPigfoxCaller(address common.Address, caller bind.ContractCaller) (*PigfoxCaller, error) {
	contract, err := bindPigfox(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PigfoxCaller{contract: contract}, nil
}

// NewPigfoxTransactor creates a new write-only instance of Pigfox, bound to a specific deployed contract.
func NewPigfoxTransactor(address common.Address, transactor bind.ContractTransactor) (*PigfoxTransactor, error) {
	contract, err := bindPigfox(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PigfoxTransactor{contract: contract}, nil
}

// NewPigfoxFilterer creates a new log filterer instance of Pigfox, bound to a specific deployed contract.
func NewPigfoxFilterer(address common.Address, filterer bind.ContractFilterer) (*PigfoxFilterer, error) {
	contract, err := bindPigfox(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PigfoxFilterer{contract: contract}, nil
}

// bindPigfox binds a generic wrapper to an already deployed contract.
func bindPigfox(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PigfoxMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pigfox *PigfoxRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pigfox.Contract.PigfoxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pigfox *PigfoxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pigfox.Contract.PigfoxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pigfox *PigfoxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pigfox.Contract.PigfoxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pigfox *PigfoxCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pigfox.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pigfox *PigfoxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pigfox.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pigfox *PigfoxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pigfox.Contract.contract.Transact(opts, method, params...)
}

// Swap is a paid mutator transaction binding the contract method 0xdba43ee9.
//
// Solidity: function Swap(address _router0, address _router1, address _token0, address _token1, uint256 amount, uint256 slippageTolerancePctInt, uint256 deadlineSecs) returns()
func (_Pigfox *PigfoxTransactor) Swap(opts *bind.TransactOpts, _router0 common.Address, _router1 common.Address, _token0 common.Address, _token1 common.Address, amount *big.Int, slippageTolerancePctInt *big.Int, deadlineSecs *big.Int) (*types.Transaction, error) {
	return _Pigfox.contract.Transact(opts, "Swap", _router0, _router1, _token0, _token1, amount, slippageTolerancePctInt, deadlineSecs)
}

// Swap is a paid mutator transaction binding the contract method 0xdba43ee9.
//
// Solidity: function Swap(address _router0, address _router1, address _token0, address _token1, uint256 amount, uint256 slippageTolerancePctInt, uint256 deadlineSecs) returns()
func (_Pigfox *PigfoxSession) Swap(_router0 common.Address, _router1 common.Address, _token0 common.Address, _token1 common.Address, amount *big.Int, slippageTolerancePctInt *big.Int, deadlineSecs *big.Int) (*types.Transaction, error) {
	return _Pigfox.Contract.Swap(&_Pigfox.TransactOpts, _router0, _router1, _token0, _token1, amount, slippageTolerancePctInt, deadlineSecs)
}

// Swap is a paid mutator transaction binding the contract method 0xdba43ee9.
//
// Solidity: function Swap(address _router0, address _router1, address _token0, address _token1, uint256 amount, uint256 slippageTolerancePctInt, uint256 deadlineSecs) returns()
func (_Pigfox *PigfoxTransactorSession) Swap(_router0 common.Address, _router1 common.Address, _token0 common.Address, _token1 common.Address, amount *big.Int, slippageTolerancePctInt *big.Int, deadlineSecs *big.Int) (*types.Transaction, error) {
	return _Pigfox.Contract.Swap(&_Pigfox.TransactOpts, _router0, _router1, _token0, _token1, amount, slippageTolerancePctInt, deadlineSecs)
}
