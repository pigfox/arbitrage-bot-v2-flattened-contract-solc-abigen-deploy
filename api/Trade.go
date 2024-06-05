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

// TradeMetaData contains all meta data concerning the Trade contract.
var TradeMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sourceDex\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"destinationDex\",\"type\":\"address\"}],\"name\":\"SwapExecuted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sourceDexAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_destinationDexAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_sunTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Swap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061047e806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063ffebebfb14610030575b600080fd5b61004361003e3660046103d4565b610045565b005b6000811161009a5760405162461bcd60e51b815260206004820152601d60248201527f416d6f756e74206d7573742062652067726561746572207468616e203000000060448201526064015b60405180910390fd5b6040516323b872dd60e01b8152336004820152306024820152604481018290528490849084906000906001600160a01b038316906323b872dd906064016020604051808303816000875af11580156100f6573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061011a919061041f565b9050806101615760405162461bcd60e51b8152602060048201526015602482015274151bdad95b881d1c985b9cd9995c8819985a5b1959605a1b6044820152606401610091565b60405163095ea7b360e01b81526001600160a01b0389811660048301526024820187905283169063095ea7b3906044016020604051808303816000875af11580156101b0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101d4919061041f565b5060405163b6b55f2560e01b8152600481018690526001600160a01b0385169063b6b55f2590602401600060405180830381600087803b15801561021757600080fd5b505af115801561022b573d6000803e3d6000fd5b5050604051632e1a7d4d60e01b8152600481018890526001600160a01b0387169250632e1a7d4d9150602401600060405180830381600087803b15801561027157600080fd5b505af1158015610285573d6000803e3d6000fd5b505060405163095ea7b360e01b81526001600160a01b038a81166004830152602482018990528516925063095ea7b391506044016020604051808303816000875af11580156102d8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102fc919061041f565b5060405163b6b55f2560e01b8152600481018690526001600160a01b0384169063b6b55f2590602401600060405180830381600087803b15801561033f57600080fd5b505af1158015610353573d6000803e3d6000fd5b50505050866001600160a01b0316886001600160a01b0316336001600160a01b03167f0e2cf19b6e7d9b18797e20fac66b57ebef5d92381316e9f6709d52408dee3d68886040516103a691815260200190565b60405180910390a45050505050505050565b80356001600160a01b03811681146103cf57600080fd5b919050565b600080600080608085870312156103ea57600080fd5b6103f3856103b8565b9350610401602086016103b8565b925061040f604086016103b8565b9396929550929360600135925050565b60006020828403121561043157600080fd5b8151801515811461044157600080fd5b939250505056fea2646970667358221220e28c095de53cf48dec069d099f214c9b07d0c764b09c7d233f1286c32d18450764736f6c63430008140033",
}

// TradeABI is the input ABI used to generate the binding from.
// Deprecated: Use TradeMetaData.ABI instead.
var TradeABI = TradeMetaData.ABI

// TradeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TradeMetaData.Bin instead.
var TradeBin = TradeMetaData.Bin

// DeployTrade deploys a new Ethereum contract, binding an instance of Trade to it.
func DeployTrade(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Trade, error) {
	parsed, err := TradeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TradeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Trade{TradeCaller: TradeCaller{contract: contract}, TradeTransactor: TradeTransactor{contract: contract}, TradeFilterer: TradeFilterer{contract: contract}}, nil
}

// Trade is an auto generated Go binding around an Ethereum contract.
type Trade struct {
	TradeCaller     // Read-only binding to the contract
	TradeTransactor // Write-only binding to the contract
	TradeFilterer   // Log filterer for contract events
}

// TradeCaller is an auto generated read-only Go binding around an Ethereum contract.
type TradeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TradeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TradeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TradeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TradeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TradeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TradeSession struct {
	Contract     *Trade            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TradeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TradeCallerSession struct {
	Contract *TradeCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TradeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TradeTransactorSession struct {
	Contract     *TradeTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TradeRaw is an auto generated low-level Go binding around an Ethereum contract.
type TradeRaw struct {
	Contract *Trade // Generic contract binding to access the raw methods on
}

// TradeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TradeCallerRaw struct {
	Contract *TradeCaller // Generic read-only contract binding to access the raw methods on
}

// TradeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TradeTransactorRaw struct {
	Contract *TradeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTrade creates a new instance of Trade, bound to a specific deployed contract.
func NewTrade(address common.Address, backend bind.ContractBackend) (*Trade, error) {
	contract, err := bindTrade(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Trade{TradeCaller: TradeCaller{contract: contract}, TradeTransactor: TradeTransactor{contract: contract}, TradeFilterer: TradeFilterer{contract: contract}}, nil
}

// NewTradeCaller creates a new read-only instance of Trade, bound to a specific deployed contract.
func NewTradeCaller(address common.Address, caller bind.ContractCaller) (*TradeCaller, error) {
	contract, err := bindTrade(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TradeCaller{contract: contract}, nil
}

// NewTradeTransactor creates a new write-only instance of Trade, bound to a specific deployed contract.
func NewTradeTransactor(address common.Address, transactor bind.ContractTransactor) (*TradeTransactor, error) {
	contract, err := bindTrade(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TradeTransactor{contract: contract}, nil
}

// NewTradeFilterer creates a new log filterer instance of Trade, bound to a specific deployed contract.
func NewTradeFilterer(address common.Address, filterer bind.ContractFilterer) (*TradeFilterer, error) {
	contract, err := bindTrade(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TradeFilterer{contract: contract}, nil
}

// bindTrade binds a generic wrapper to an already deployed contract.
func bindTrade(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TradeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Trade *TradeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Trade.Contract.TradeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Trade *TradeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Trade.Contract.TradeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Trade *TradeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Trade.Contract.TradeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Trade *TradeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Trade.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Trade *TradeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Trade.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Trade *TradeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Trade.Contract.contract.Transact(opts, method, params...)
}

// Swap is a paid mutator transaction binding the contract method 0xffebebfb.
//
// Solidity: function Swap(address _sourceDexAddress, address _destinationDexAddress, address _sunTokenAddress, uint256 amount) returns()
func (_Trade *TradeTransactor) Swap(opts *bind.TransactOpts, _sourceDexAddress common.Address, _destinationDexAddress common.Address, _sunTokenAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Trade.contract.Transact(opts, "Swap", _sourceDexAddress, _destinationDexAddress, _sunTokenAddress, amount)
}

// Swap is a paid mutator transaction binding the contract method 0xffebebfb.
//
// Solidity: function Swap(address _sourceDexAddress, address _destinationDexAddress, address _sunTokenAddress, uint256 amount) returns()
func (_Trade *TradeSession) Swap(_sourceDexAddress common.Address, _destinationDexAddress common.Address, _sunTokenAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Trade.Contract.Swap(&_Trade.TransactOpts, _sourceDexAddress, _destinationDexAddress, _sunTokenAddress, amount)
}

// Swap is a paid mutator transaction binding the contract method 0xffebebfb.
//
// Solidity: function Swap(address _sourceDexAddress, address _destinationDexAddress, address _sunTokenAddress, uint256 amount) returns()
func (_Trade *TradeTransactorSession) Swap(_sourceDexAddress common.Address, _destinationDexAddress common.Address, _sunTokenAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Trade.Contract.Swap(&_Trade.TransactOpts, _sourceDexAddress, _destinationDexAddress, _sunTokenAddress, amount)
}

// TradeSwapExecutedIterator is returned from FilterSwapExecuted and is used to iterate over the raw logs and unpacked data for SwapExecuted events raised by the Trade contract.
type TradeSwapExecutedIterator struct {
	Event *TradeSwapExecuted // Event containing the contract specifics and raw log

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
func (it *TradeSwapExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TradeSwapExecuted)
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
		it.Event = new(TradeSwapExecuted)
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
func (it *TradeSwapExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TradeSwapExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TradeSwapExecuted represents a SwapExecuted event raised by the Trade contract.
type TradeSwapExecuted struct {
	User           common.Address
	Amount         *big.Int
	SourceDex      common.Address
	DestinationDex common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSwapExecuted is a free log retrieval operation binding the contract event 0x0e2cf19b6e7d9b18797e20fac66b57ebef5d92381316e9f6709d52408dee3d68.
//
// Solidity: event SwapExecuted(address indexed user, uint256 amount, address indexed sourceDex, address indexed destinationDex)
func (_Trade *TradeFilterer) FilterSwapExecuted(opts *bind.FilterOpts, user []common.Address, sourceDex []common.Address, destinationDex []common.Address) (*TradeSwapExecutedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	var sourceDexRule []interface{}
	for _, sourceDexItem := range sourceDex {
		sourceDexRule = append(sourceDexRule, sourceDexItem)
	}
	var destinationDexRule []interface{}
	for _, destinationDexItem := range destinationDex {
		destinationDexRule = append(destinationDexRule, destinationDexItem)
	}

	logs, sub, err := _Trade.contract.FilterLogs(opts, "SwapExecuted", userRule, sourceDexRule, destinationDexRule)
	if err != nil {
		return nil, err
	}
	return &TradeSwapExecutedIterator{contract: _Trade.contract, event: "SwapExecuted", logs: logs, sub: sub}, nil
}

// WatchSwapExecuted is a free log subscription operation binding the contract event 0x0e2cf19b6e7d9b18797e20fac66b57ebef5d92381316e9f6709d52408dee3d68.
//
// Solidity: event SwapExecuted(address indexed user, uint256 amount, address indexed sourceDex, address indexed destinationDex)
func (_Trade *TradeFilterer) WatchSwapExecuted(opts *bind.WatchOpts, sink chan<- *TradeSwapExecuted, user []common.Address, sourceDex []common.Address, destinationDex []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	var sourceDexRule []interface{}
	for _, sourceDexItem := range sourceDex {
		sourceDexRule = append(sourceDexRule, sourceDexItem)
	}
	var destinationDexRule []interface{}
	for _, destinationDexItem := range destinationDex {
		destinationDexRule = append(destinationDexRule, destinationDexItem)
	}

	logs, sub, err := _Trade.contract.WatchLogs(opts, "SwapExecuted", userRule, sourceDexRule, destinationDexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TradeSwapExecuted)
				if err := _Trade.contract.UnpackLog(event, "SwapExecuted", log); err != nil {
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

// ParseSwapExecuted is a log parse operation binding the contract event 0x0e2cf19b6e7d9b18797e20fac66b57ebef5d92381316e9f6709d52408dee3d68.
//
// Solidity: event SwapExecuted(address indexed user, uint256 amount, address indexed sourceDex, address indexed destinationDex)
func (_Trade *TradeFilterer) ParseSwapExecuted(log types.Log) (*TradeSwapExecuted, error) {
	event := new(TradeSwapExecuted)
	if err := _Trade.contract.UnpackLog(event, "SwapExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
