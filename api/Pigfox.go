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
	Bin: "0x608060405234801561001057600080fd5b50604051610ab7380380610ab783398101604081905261002f916100b7565b6001600160a01b0381166100895760405162461bcd60e51b815260206004820152601360248201527f496e76616c69642064657374696e6174696f6e00000000000000000000000000604482015260640160405180910390fd5b600180546001600160a01b039092166001600160a01b031992831617905560008054909116331790556100e7565b6000602082840312156100c957600080fd5b81516001600160a01b03811681146100e057600080fd5b9392505050565b6109c1806100f66000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80630a0a05e61461003b578063dba43ee914610050575b600080fd5b61004e6100493660046106c8565b610063565b005b61004e61005e3660046106ec565b61011c565b6000546001600160a01b031633146100ae5760405162461bcd60e51b81526020600482015260096024820152682737ba1037bbb732b960b91b60448201526064015b60405180910390fd5b6001600160a01b0381166100fa5760405162461bcd60e51b815260206004820152601360248201527224b73b30b634b2103232b9ba34b730ba34b7b760691b60448201526064016100a5565b600180546001600160a01b0319166001600160a01b0392909216919091179055565b6000546001600160a01b031633146101625760405162461bcd60e51b81526020600482015260096024820152682737ba1037bbb732b960b91b60448201526064016100a5565b600280546001600160a01b03199081166001600160a01b038a8116918217909355600380549092168984161790915560405163095ea7b360e01b8152600481019190915260248101859052869186919083169063095ea7b3906044016020604051808303816000875af11580156101dd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102019190610763565b61024d5760405162461bcd60e51b815260206004820181905260248201527f746f6b656e30206f6e20726f757465723020617070726f7665206661696c656460448201526064016100a5565b60035460405163095ea7b360e01b81526001600160a01b039182166004820152602481018790529082169063095ea7b3906044016020604051808303816000875af11580156102a0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102c49190610763565b6103105760405162461bcd60e51b815260206004820181905260248201527f746f6b656e31206f6e20726f757465723120617070726f7665206661696c656460448201526064016100a5565b60408051600280825260608201835260009260208301908036833701905050905087816000815181106103455761034561079b565b60200260200101906001600160a01b031690816001600160a01b03168152505086816001815181106103795761037961079b565b6001600160a01b03928316602091820292909201015260025460405163d06ca61f60e01b8152600092919091169063d06ca61f906103bd908a9086906004016107f5565b600060405180830381865afa1580156103da573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526104029190810190610816565b90506000606461041288826108ea565b836001815181106104255761042561079b565b60200260200101516104379190610903565b610441919061091a565b6002549091506001600160a01b03166338ed1739898386306104638c4261093c565b6040518663ffffffff1660e01b815260040161048395949392919061094f565b6000604051808303816000875af11580156104a2573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526104ca9190810190610816565b50604080516002808252606082018352839260009291906020830190803683370190505090508a816000815181106105045761050461079b565b60200260200101906001600160a01b031690816001600160a01b0316815250508b816001815181106105385761053861079b565b6001600160a01b03928316602091820292909201015260035460405163d06ca61f60e01b8152600092919091169063d06ca61f9061057c90869086906004016107f5565b600060405180830381865afa158015610599573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526105c19190810190610816565b9050600060646105d18c826108ea565b836001815181106105e4576105e461079b565b60200260200101516105f69190610903565b610600919061091a565b9050600360009054906101000a90046001600160a01b03166001600160a01b03166338ed1739858386308f42610636919061093c565b6040518663ffffffff1660e01b815260040161065695949392919061094f565b6000604051808303816000875af1158015610675573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261069d9190810190610816565b5050505050505050505050505050505050565b6001600160a01b03811681146106c557600080fd5b50565b6000602082840312156106da57600080fd5b81356106e5816106b0565b9392505050565b600080600080600080600060e0888a03121561070757600080fd5b8735610712816106b0565b96506020880135610722816106b0565b95506040880135610732816106b0565b94506060880135610742816106b0565b9699959850939660808101359560a0820135955060c0909101359350915050565b60006020828403121561077557600080fd5b815180151581146106e557600080fd5b634e487b7160e01b600052604160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b600081518084526020808501945080840160005b838110156107ea5781516001600160a01b0316875295820195908201906001016107c5565b509495945050505050565b82815260406020820152600061080e60408301846107b1565b949350505050565b6000602080838503121561082957600080fd5b825167ffffffffffffffff8082111561084157600080fd5b818501915085601f83011261085557600080fd5b81518181111561086757610867610785565b8060051b604051601f19603f8301168101818110858211171561088c5761088c610785565b6040529182528482019250838101850191888311156108aa57600080fd5b938501935b828510156108c8578451845293850193928501926108af565b98975050505050505050565b634e487b7160e01b600052601160045260246000fd5b818103818111156108fd576108fd6108d4565b92915050565b80820281158282048414176108fd576108fd6108d4565b60008261093757634e487b7160e01b600052601260045260246000fd5b500490565b808201808211156108fd576108fd6108d4565b85815284602082015260a06040820152600061096e60a08301866107b1565b6001600160a01b039490941660608301525060800152939250505056fea2646970667358221220aeec4a1dc30b028e49a7f972611022d99ca57227e43ac6030c17fe6e92c81a6f64736f6c63430008140033",
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
