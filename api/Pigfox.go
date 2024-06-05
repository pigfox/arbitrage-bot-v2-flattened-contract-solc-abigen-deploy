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
	ABI: "[{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_destination\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_router0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_router1\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token1\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"slippageTolerancePctInt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadlineSecs\",\"type\":\"uint256\"}],\"name\":\"Swap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_destination\",\"type\":\"address\"}],\"name\":\"setDestination\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b50604051610a6e380380610a6e83398101604081905261002e916100b5565b6001600160a01b0381166100885760405162461bcd60e51b815260206004820152601360248201527f496e76616c69642064657374696e6174696f6e00000000000000000000000000604482015260640160405180910390fd5b600180546001600160a01b039092166001600160a01b03199283161790555f8054909116331790556100e2565b5f602082840312156100c5575f80fd5b81516001600160a01b03811681146100db575f80fd5b9392505050565b61097f806100ef5f395ff3fe608060405234801561000f575f80fd5b5060043610610034575f3560e01c80630a0a05e614610038578063dba43ee91461004d575b5f80fd5b61004b6100463660046106a1565b610060565b005b61004b61005b3660046106c3565b610118565b5f546001600160a01b031633146100aa5760405162461bcd60e51b81526020600482015260096024820152682737ba1037bbb732b960b91b60448201526064015b60405180910390fd5b6001600160a01b0381166100f65760405162461bcd60e51b815260206004820152601360248201527224b73b30b634b2103232b9ba34b730ba34b7b760691b60448201526064016100a1565b600180546001600160a01b0319166001600160a01b0392909216919091179055565b5f546001600160a01b0316331461015d5760405162461bcd60e51b81526020600482015260096024820152682737ba1037bbb732b960b91b60448201526064016100a1565b600280546001600160a01b03199081166001600160a01b038a8116918217909355600380549092168984161790915560405163095ea7b360e01b8152600481019190915260248101859052869186919083169063095ea7b3906044016020604051808303815f875af11580156101d5573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906101f99190610735565b6102455760405162461bcd60e51b815260206004820181905260248201527f746f6b656e30206f6e20726f757465723020617070726f7665206661696c656460448201526064016100a1565b60035460405163095ea7b360e01b81526001600160a01b039182166004820152602481018790529082169063095ea7b3906044016020604051808303815f875af1158015610295573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906102b99190610735565b6103055760405162461bcd60e51b815260206004820181905260248201527f746f6b656e31206f6e20726f757465723120617070726f7665206661696c656460448201526064016100a1565b6040805160028082526060820183525f9260208301908036833701905050905087815f8151811061033857610338610768565b60200260200101906001600160a01b031690816001600160a01b031681525050868160018151811061036c5761036c610768565b6001600160a01b03928316602091820292909201015260025460405163d06ca61f60e01b81525f92919091169063d06ca61f906103af908a9086906004016107bf565b5f60405180830381865afa1580156103c9573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526103f091908101906107df565b90505f60646103ff88826108ac565b8360018151811061041257610412610768565b602002602001015161042491906108c5565b61042e91906108dc565b6002549091506001600160a01b03166338ed1739898386306104508c426108fb565b6040518663ffffffff1660e01b815260040161047095949392919061090e565b5f604051808303815f875af115801561048b573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526104b291908101906107df565b5060408051600280825260608201835283925f9291906020830190803683370190505090508a815f815181106104ea576104ea610768565b60200260200101906001600160a01b031690816001600160a01b0316815250508b8160018151811061051e5761051e610768565b6001600160a01b03928316602091820292909201015260035460405163d06ca61f60e01b81525f92919091169063d06ca61f9061056190869086906004016107bf565b5f60405180830381865afa15801561057b573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526105a291908101906107df565b90505f60646105b18c826108ac565b836001815181106105c4576105c4610768565b60200260200101516105d691906108c5565b6105e091906108dc565b905060035f9054906101000a90046001600160a01b03166001600160a01b03166338ed1739858386308f4261061591906108fb565b6040518663ffffffff1660e01b815260040161063595949392919061090e565b5f604051808303815f875af1158015610650573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261067791908101906107df565b5050505050505050505050505050505050565b6001600160a01b038116811461069e575f80fd5b50565b5f602082840312156106b1575f80fd5b81356106bc8161068a565b9392505050565b5f805f805f805f60e0888a0312156106d9575f80fd5b87356106e48161068a565b965060208801356106f48161068a565b955060408801356107048161068a565b945060608801356107148161068a565b9699959850939660808101359560a0820135955060c0909101359350915050565b5f60208284031215610745575f80fd5b815180151581146106bc575f80fd5b634e487b7160e01b5f52604160045260245ffd5b634e487b7160e01b5f52603260045260245ffd5b5f815180845260208085019450602084015f5b838110156107b45781516001600160a01b03168752958201959082019060010161078f565b509495945050505050565b828152604060208201525f6107d7604083018461077c565b949350505050565b5f60208083850312156107f0575f80fd5b825167ffffffffffffffff80821115610807575f80fd5b818501915085601f83011261081a575f80fd5b81518181111561082c5761082c610754565b8060051b604051601f19603f8301168101818110858211171561085157610851610754565b60405291825284820192508381018501918883111561086e575f80fd5b938501935b8285101561088c57845184529385019392850192610873565b98975050505050505050565b634e487b7160e01b5f52601160045260245ffd5b818103818111156108bf576108bf610898565b92915050565b80820281158282048414176108bf576108bf610898565b5f826108f657634e487b7160e01b5f52601260045260245ffd5b500490565b808201808211156108bf576108bf610898565b85815284602082015260a060408201525f61092c60a083018661077c565b6001600160a01b039490941660608301525060800152939250505056fea26469706673582212205110604f45c40ce7e9fd807de40ea390de147c300c80752af186b062e81768d764736f6c63430008180033",
}

// PigfoxABI is the input ABI used to generate the binding from.
// Deprecated: Use PigfoxMetaData.ABI instead.
var PigfoxABI = PigfoxMetaData.ABI

// PigfoxBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PigfoxMetaData.Bin instead.
var PigfoxBin = PigfoxMetaData.Bin

// DeployPigfox deploys a new Ethereum contract, binding an instance of Pigfox to it.
func DeployPigfox(auth *bind.TransactOpts, backend bind.ContractBackend, _destination common.Address) (common.Address, *types.Transaction, *Pigfox, error) {
	parsed, err := PigfoxMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PigfoxBin), backend, _destination)
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

// SetDestination is a paid mutator transaction binding the contract method 0x0a0a05e6.
//
// Solidity: function setDestination(address _destination) returns()
func (_Pigfox *PigfoxTransactor) SetDestination(opts *bind.TransactOpts, _destination common.Address) (*types.Transaction, error) {
	return _Pigfox.contract.Transact(opts, "setDestination", _destination)
}

// SetDestination is a paid mutator transaction binding the contract method 0x0a0a05e6.
//
// Solidity: function setDestination(address _destination) returns()
func (_Pigfox *PigfoxSession) SetDestination(_destination common.Address) (*types.Transaction, error) {
	return _Pigfox.Contract.SetDestination(&_Pigfox.TransactOpts, _destination)
}

// SetDestination is a paid mutator transaction binding the contract method 0x0a0a05e6.
//
// Solidity: function setDestination(address _destination) returns()
func (_Pigfox *PigfoxTransactorSession) SetDestination(_destination common.Address) (*types.Transaction, error) {
	return _Pigfox.Contract.SetDestination(&_Pigfox.TransactOpts, _destination)
}
