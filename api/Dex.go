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

// DexMetaData contains all meta data concerning the Dex contract.
var DexMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Trade\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"trade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506106b3806100206000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c806347e7ef241461005c57806368f8c2a414610071578063c23f001f14610084578063d4fac45d146100be578063f3fef3a3146100d1575b600080fd5b61006f61006a366004610584565b6100e4565b005b61006f61007f366004610584565b61024d565b6100ac6100923660046105ae565b600060208181529281526040808220909352908152205481565b60405190815260200160405180910390f35b6100ac6100cc3660046105ae565b610385565b61006f6100df366004610584565b6103b0565b6000811161010d5760405162461bcd60e51b8152600401610104906105e1565b60405180910390fd5b6040516323b872dd60e01b81523360048201523060248201526044810182905282906000906001600160a01b038316906323b872dd906064016020604051808303816000875af1158015610165573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101899190610618565b9050806101d05760405162461bcd60e51b8152602060048201526015602482015274151bdad95b881d1c985b9cd9995c8819985a5b1959605a1b6044820152606401610104565b6001600160a01b03841660009081526020818152604080832033845290915281208054859290610201908490610657565b909155505060405183815233906001600160a01b038616907f5548c837ab068cf56a2c2479df0882a4922fd203edb7517321831d95078c5f62906020015b60405180910390a350505050565b6000811161026d5760405162461bcd60e51b8152600401610104906105e1565b6001600160a01b0382166000908152602081815260408083203384529091529020548111156102d55760405162461bcd60e51b8152602060048201526014602482015273496e73756666696369656e742062616c616e636560601b6044820152606401610104565b6001600160a01b0382166000908152602081815260408083203384529091528120805483929061030690849061066a565b90915550506001600160a01b0382166000908152602081815260408083203384529091528120805483929061033c908490610657565b909155505060405181815233906001600160a01b038416907f0d49f50d6dd629f4e4070731edb0abd680823ab1cbd43d89590522da4b5b32e79060200160405180910390a35050565b6001600160a01b03808316600090815260208181526040808320938516835292905220545b92915050565b600081116103d05760405162461bcd60e51b8152600401610104906105e1565b6001600160a01b0382166000908152602081815260408083203384529091529020548111156104385760405162461bcd60e51b8152602060048201526014602482015273496e73756666696369656e742062616c616e636560601b6044820152606401610104565b6001600160a01b0382166000908152602081815260408083203384529091528120805483929061046990849061066a565b909155505060405163a9059cbb60e01b81523360048201526024810182905282906000906001600160a01b0383169063a9059cbb906044016020604051808303816000875af11580156104c0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104e49190610618565b90508061052b5760405162461bcd60e51b8152602060048201526015602482015274151bdad95b881d1c985b9cd9995c8819985a5b1959605a1b6044820152606401610104565b60405183815233906001600160a01b038616907f9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb9060200161023f565b80356001600160a01b038116811461057f57600080fd5b919050565b6000806040838503121561059757600080fd5b6105a083610568565b946020939093013593505050565b600080604083850312156105c157600080fd5b6105ca83610568565b91506105d860208401610568565b90509250929050565b6020808252601d908201527f416d6f756e74206d7573742062652067726561746572207468616e2030000000604082015260600190565b60006020828403121561062a57600080fd5b8151801515811461063a57600080fd5b9392505050565b634e487b7160e01b600052601160045260246000fd5b808201808211156103aa576103aa610641565b818103818111156103aa576103aa61064156fea26469706673582212200f3bab58a98ad94849f6bb32fbd3af79d6cf46aead6d95e468ca30f99072be2064736f6c63430008140033",
}

// DexABI is the input ABI used to generate the binding from.
// Deprecated: Use DexMetaData.ABI instead.
var DexABI = DexMetaData.ABI

// DexBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DexMetaData.Bin instead.
var DexBin = DexMetaData.Bin

// DeployDex deploys a new Ethereum contract, binding an instance of Dex to it.
func DeployDex(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Dex, error) {
	parsed, err := DexMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DexBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Dex{DexCaller: DexCaller{contract: contract}, DexTransactor: DexTransactor{contract: contract}, DexFilterer: DexFilterer{contract: contract}}, nil
}

// Dex is an auto generated Go binding around an Ethereum contract.
type Dex struct {
	DexCaller     // Read-only binding to the contract
	DexTransactor // Write-only binding to the contract
	DexFilterer   // Log filterer for contract events
}

// DexCaller is an auto generated read-only Go binding around an Ethereum contract.
type DexCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DexTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DexTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DexFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DexFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DexSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DexSession struct {
	Contract     *Dex              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DexCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DexCallerSession struct {
	Contract *DexCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// DexTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DexTransactorSession struct {
	Contract     *DexTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DexRaw is an auto generated low-level Go binding around an Ethereum contract.
type DexRaw struct {
	Contract *Dex // Generic contract binding to access the raw methods on
}

// DexCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DexCallerRaw struct {
	Contract *DexCaller // Generic read-only contract binding to access the raw methods on
}

// DexTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DexTransactorRaw struct {
	Contract *DexTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDex creates a new instance of Dex, bound to a specific deployed contract.
func NewDex(address common.Address, backend bind.ContractBackend) (*Dex, error) {
	contract, err := bindDex(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Dex{DexCaller: DexCaller{contract: contract}, DexTransactor: DexTransactor{contract: contract}, DexFilterer: DexFilterer{contract: contract}}, nil
}

// NewDexCaller creates a new read-only instance of Dex, bound to a specific deployed contract.
func NewDexCaller(address common.Address, caller bind.ContractCaller) (*DexCaller, error) {
	contract, err := bindDex(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DexCaller{contract: contract}, nil
}

// NewDexTransactor creates a new write-only instance of Dex, bound to a specific deployed contract.
func NewDexTransactor(address common.Address, transactor bind.ContractTransactor) (*DexTransactor, error) {
	contract, err := bindDex(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DexTransactor{contract: contract}, nil
}

// NewDexFilterer creates a new log filterer instance of Dex, bound to a specific deployed contract.
func NewDexFilterer(address common.Address, filterer bind.ContractFilterer) (*DexFilterer, error) {
	contract, err := bindDex(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DexFilterer{contract: contract}, nil
}

// bindDex binds a generic wrapper to an already deployed contract.
func bindDex(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DexMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dex *DexRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dex.Contract.DexCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dex *DexRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dex.Contract.DexTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dex *DexRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dex.Contract.DexTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dex *DexCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dex.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dex *DexTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dex.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dex *DexTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dex.Contract.contract.Transact(opts, method, params...)
}

// Balances is a free data retrieval call binding the contract method 0xc23f001f.
//
// Solidity: function balances(address , address ) view returns(uint256)
func (_Dex *DexCaller) Balances(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Dex.contract.Call(opts, &out, "balances", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0xc23f001f.
//
// Solidity: function balances(address , address ) view returns(uint256)
func (_Dex *DexSession) Balances(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Dex.Contract.Balances(&_Dex.CallOpts, arg0, arg1)
}

// Balances is a free data retrieval call binding the contract method 0xc23f001f.
//
// Solidity: function balances(address , address ) view returns(uint256)
func (_Dex *DexCallerSession) Balances(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Dex.Contract.Balances(&_Dex.CallOpts, arg0, arg1)
}

// GetBalance is a free data retrieval call binding the contract method 0xd4fac45d.
//
// Solidity: function getBalance(address token, address user) view returns(uint256)
func (_Dex *DexCaller) GetBalance(opts *bind.CallOpts, token common.Address, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Dex.contract.Call(opts, &out, "getBalance", token, user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0xd4fac45d.
//
// Solidity: function getBalance(address token, address user) view returns(uint256)
func (_Dex *DexSession) GetBalance(token common.Address, user common.Address) (*big.Int, error) {
	return _Dex.Contract.GetBalance(&_Dex.CallOpts, token, user)
}

// GetBalance is a free data retrieval call binding the contract method 0xd4fac45d.
//
// Solidity: function getBalance(address token, address user) view returns(uint256)
func (_Dex *DexCallerSession) GetBalance(token common.Address, user common.Address) (*big.Int, error) {
	return _Dex.Contract.GetBalance(&_Dex.CallOpts, token, user)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) returns()
func (_Dex *DexTransactor) Deposit(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Dex.contract.Transact(opts, "deposit", token, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) returns()
func (_Dex *DexSession) Deposit(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Dex.Contract.Deposit(&_Dex.TransactOpts, token, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) returns()
func (_Dex *DexTransactorSession) Deposit(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Dex.Contract.Deposit(&_Dex.TransactOpts, token, amount)
}

// Trade is a paid mutator transaction binding the contract method 0x68f8c2a4.
//
// Solidity: function trade(address token, uint256 amount) returns()
func (_Dex *DexTransactor) Trade(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Dex.contract.Transact(opts, "trade", token, amount)
}

// Trade is a paid mutator transaction binding the contract method 0x68f8c2a4.
//
// Solidity: function trade(address token, uint256 amount) returns()
func (_Dex *DexSession) Trade(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Dex.Contract.Trade(&_Dex.TransactOpts, token, amount)
}

// Trade is a paid mutator transaction binding the contract method 0x68f8c2a4.
//
// Solidity: function trade(address token, uint256 amount) returns()
func (_Dex *DexTransactorSession) Trade(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Dex.Contract.Trade(&_Dex.TransactOpts, token, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address token, uint256 amount) returns()
func (_Dex *DexTransactor) Withdraw(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Dex.contract.Transact(opts, "withdraw", token, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address token, uint256 amount) returns()
func (_Dex *DexSession) Withdraw(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Dex.Contract.Withdraw(&_Dex.TransactOpts, token, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address token, uint256 amount) returns()
func (_Dex *DexTransactorSession) Withdraw(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Dex.Contract.Withdraw(&_Dex.TransactOpts, token, amount)
}

// DexDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Dex contract.
type DexDepositIterator struct {
	Event *DexDeposit // Event containing the contract specifics and raw log

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
func (it *DexDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DexDeposit)
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
		it.Event = new(DexDeposit)
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
func (it *DexDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DexDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DexDeposit represents a Deposit event raised by the Dex contract.
type DexDeposit struct {
	Token  common.Address
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x5548c837ab068cf56a2c2479df0882a4922fd203edb7517321831d95078c5f62.
//
// Solidity: event Deposit(address indexed token, address indexed user, uint256 amount)
func (_Dex *DexFilterer) FilterDeposit(opts *bind.FilterOpts, token []common.Address, user []common.Address) (*DexDepositIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Dex.contract.FilterLogs(opts, "Deposit", tokenRule, userRule)
	if err != nil {
		return nil, err
	}
	return &DexDepositIterator{contract: _Dex.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x5548c837ab068cf56a2c2479df0882a4922fd203edb7517321831d95078c5f62.
//
// Solidity: event Deposit(address indexed token, address indexed user, uint256 amount)
func (_Dex *DexFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *DexDeposit, token []common.Address, user []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Dex.contract.WatchLogs(opts, "Deposit", tokenRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DexDeposit)
				if err := _Dex.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x5548c837ab068cf56a2c2479df0882a4922fd203edb7517321831d95078c5f62.
//
// Solidity: event Deposit(address indexed token, address indexed user, uint256 amount)
func (_Dex *DexFilterer) ParseDeposit(log types.Log) (*DexDeposit, error) {
	event := new(DexDeposit)
	if err := _Dex.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DexTradeIterator is returned from FilterTrade and is used to iterate over the raw logs and unpacked data for Trade events raised by the Dex contract.
type DexTradeIterator struct {
	Event *DexTrade // Event containing the contract specifics and raw log

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
func (it *DexTradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DexTrade)
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
		it.Event = new(DexTrade)
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
func (it *DexTradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DexTradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DexTrade represents a Trade event raised by the Dex contract.
type DexTrade struct {
	Token  common.Address
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTrade is a free log retrieval operation binding the contract event 0x0d49f50d6dd629f4e4070731edb0abd680823ab1cbd43d89590522da4b5b32e7.
//
// Solidity: event Trade(address indexed token, address indexed user, uint256 amount)
func (_Dex *DexFilterer) FilterTrade(opts *bind.FilterOpts, token []common.Address, user []common.Address) (*DexTradeIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Dex.contract.FilterLogs(opts, "Trade", tokenRule, userRule)
	if err != nil {
		return nil, err
	}
	return &DexTradeIterator{contract: _Dex.contract, event: "Trade", logs: logs, sub: sub}, nil
}

// WatchTrade is a free log subscription operation binding the contract event 0x0d49f50d6dd629f4e4070731edb0abd680823ab1cbd43d89590522da4b5b32e7.
//
// Solidity: event Trade(address indexed token, address indexed user, uint256 amount)
func (_Dex *DexFilterer) WatchTrade(opts *bind.WatchOpts, sink chan<- *DexTrade, token []common.Address, user []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Dex.contract.WatchLogs(opts, "Trade", tokenRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DexTrade)
				if err := _Dex.contract.UnpackLog(event, "Trade", log); err != nil {
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

// ParseTrade is a log parse operation binding the contract event 0x0d49f50d6dd629f4e4070731edb0abd680823ab1cbd43d89590522da4b5b32e7.
//
// Solidity: event Trade(address indexed token, address indexed user, uint256 amount)
func (_Dex *DexFilterer) ParseTrade(log types.Log) (*DexTrade, error) {
	event := new(DexTrade)
	if err := _Dex.contract.UnpackLog(event, "Trade", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DexWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Dex contract.
type DexWithdrawIterator struct {
	Event *DexWithdraw // Event containing the contract specifics and raw log

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
func (it *DexWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DexWithdraw)
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
		it.Event = new(DexWithdraw)
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
func (it *DexWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DexWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DexWithdraw represents a Withdraw event raised by the Dex contract.
type DexWithdraw struct {
	Token  common.Address
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address indexed token, address indexed user, uint256 amount)
func (_Dex *DexFilterer) FilterWithdraw(opts *bind.FilterOpts, token []common.Address, user []common.Address) (*DexWithdrawIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Dex.contract.FilterLogs(opts, "Withdraw", tokenRule, userRule)
	if err != nil {
		return nil, err
	}
	return &DexWithdrawIterator{contract: _Dex.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address indexed token, address indexed user, uint256 amount)
func (_Dex *DexFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *DexWithdraw, token []common.Address, user []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Dex.contract.WatchLogs(opts, "Withdraw", tokenRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DexWithdraw)
				if err := _Dex.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address indexed token, address indexed user, uint256 amount)
func (_Dex *DexFilterer) ParseWithdraw(log types.Log) (*DexWithdraw, error) {
	event := new(DexWithdraw)
	if err := _Dex.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
