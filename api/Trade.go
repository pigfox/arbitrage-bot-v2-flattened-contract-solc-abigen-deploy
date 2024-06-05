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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_profitDestination\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sourceDex\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"destinationDex\",\"type\":\"address\"}],\"name\":\"SwapExecuted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProfitDestination\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_profitDestination\",\"type\":\"address\"}],\"name\":\"setProfitDestination\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sourceDexAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_destinationDexAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"swap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5060405161090338038061090383398101604081905261002f91610062565b60008054336001600160a01b031991821617909155600180549091166001600160a01b0392909216919091179055610092565b60006020828403121561007457600080fd5b81516001600160a01b038116811461008b57600080fd5b9392505050565b610862806100a16000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c806313af40351461005c5780632690ac4b14610071578063893d20e81461009a578063a6c70871146100ab578063a9678a18146100be575b600080fd5b61006f61006a36600461071c565b6100d1565b005b6001546001600160a01b03165b6040516001600160a01b03909116815260200160405180910390f35b6000546001600160a01b031661007e565b61006f6100b936600461071c565b610126565b61006f6100cc36600461073e565b610172565b6000546001600160a01b031633146101045760405162461bcd60e51b81526004016100fb90610789565b60405180910390fd5b600080546001600160a01b0319166001600160a01b0392909216919091179055565b6000546001600160a01b031633146101505760405162461bcd60e51b81526004016100fb90610789565b600180546001600160a01b0319166001600160a01b0392909216919091179055565b6000546001600160a01b0316331461019c5760405162461bcd60e51b81526004016100fb90610789565b600081116101ec5760405162461bcd60e51b815260206004820152601d60248201527f416d6f756e74206d7573742062652067726561746572207468616e203000000060448201526064016100fb565b6040516323b872dd60e01b8152336004820152306024820152604481018290528490849084906000906001600160a01b038316906323b872dd906064016020604051808303816000875af1158015610248573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061026c91906107ca565b9050806102b35760405162461bcd60e51b8152602060048201526015602482015274151bdad95b881d1c985b9cd9995c8819985a5b1959605a1b60448201526064016100fb565b60405163095ea7b360e01b81526001600160a01b0389811660048301526024820187905283169063095ea7b3906044016020604051808303816000875af1158015610302573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061032691906107ca565b506040516311f9fbc960e21b81526001600160a01b038781166004830152602482018790528516906347e7ef2490604401600060405180830381600087803b15801561037157600080fd5b505af1158015610385573d6000803e3d6000fd5b505060405163f3fef3a360e01b81526001600160a01b038981166004830152602482018990528716925063f3fef3a39150604401600060405180830381600087803b1580156103d357600080fd5b505af11580156103e7573d6000803e3d6000fd5b505060405163d4fac45d60e01b81526001600160a01b038981166004830152306024830152600093508616915063d4fac45d90604401602060405180830381865afa15801561043a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061045e91906107ec565b60405163095ea7b360e01b81526001600160a01b038a81166004830152602482018990529192509084169063095ea7b3906044016020604051808303816000875af11580156104b1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104d591906107ca565b506040516311f9fbc960e21b81526001600160a01b038881166004830152602482018890528516906347e7ef2490604401600060405180830381600087803b15801561052057600080fd5b505af1158015610534573d6000803e3d6000fd5b505060405163f3fef3a360e01b81526001600160a01b038a81166004830152602482018a90528716925063f3fef3a39150604401600060405180830381600087803b15801561058257600080fd5b505af1158015610596573d6000803e3d6000fd5b505060405163d4fac45d60e01b81526001600160a01b038a81166004830152306024830152600093508716915063d4fac45d90604401602060405180830381865afa1580156105e9573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061060d91906107ec565b9050600061061b8383610805565b9050801561069c5760015460405163a9059cbb60e01b81526001600160a01b039182166004820152602481018390529086169063a9059cbb906044016020604051808303816000875af1158015610676573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061069a91906107ca565b505b896001600160a01b03168b6001600160a01b0316336001600160a01b03167f0e2cf19b6e7d9b18797e20fac66b57ebef5d92381316e9f6709d52408dee3d688b6040516106eb91815260200190565b60405180910390a45050505050505050505050565b80356001600160a01b038116811461071757600080fd5b919050565b60006020828403121561072e57600080fd5b61073782610700565b9392505050565b6000806000806080858703121561075457600080fd5b61075d85610700565b935061076b60208601610700565b925061077960408601610700565b9396929550929360600135925050565b60208082526021908201527f4f6e6c79206f776e65722063616e2063616c6c20746869732066756e6374696f6040820152603760f91b606082015260800190565b6000602082840312156107dc57600080fd5b8151801515811461073757600080fd5b6000602082840312156107fe57600080fd5b5051919050565b8181038181111561082657634e487b7160e01b600052601160045260246000fd5b9291505056fea26469706673582212205504a2280c1026fda8697c0270be3fddc392979fa881564cd0462d0f8ce403af64736f6c63430008140033",
}

// TradeABI is the input ABI used to generate the binding from.
// Deprecated: Use TradeMetaData.ABI instead.
var TradeABI = TradeMetaData.ABI

// TradeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TradeMetaData.Bin instead.
var TradeBin = TradeMetaData.Bin

// DeployTrade deploys a new Ethereum contract, binding an instance of Trade to it.
func DeployTrade(auth *bind.TransactOpts, backend bind.ContractBackend, _profitDestination common.Address) (common.Address, *types.Transaction, *Trade, error) {
	parsed, err := TradeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TradeBin), backend, _profitDestination)
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

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Trade *TradeCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Trade.contract.Call(opts, &out, "getOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Trade *TradeSession) GetOwner() (common.Address, error) {
	return _Trade.Contract.GetOwner(&_Trade.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Trade *TradeCallerSession) GetOwner() (common.Address, error) {
	return _Trade.Contract.GetOwner(&_Trade.CallOpts)
}

// GetProfitDestination is a free data retrieval call binding the contract method 0x2690ac4b.
//
// Solidity: function getProfitDestination() view returns(address)
func (_Trade *TradeCaller) GetProfitDestination(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Trade.contract.Call(opts, &out, "getProfitDestination")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetProfitDestination is a free data retrieval call binding the contract method 0x2690ac4b.
//
// Solidity: function getProfitDestination() view returns(address)
func (_Trade *TradeSession) GetProfitDestination() (common.Address, error) {
	return _Trade.Contract.GetProfitDestination(&_Trade.CallOpts)
}

// GetProfitDestination is a free data retrieval call binding the contract method 0x2690ac4b.
//
// Solidity: function getProfitDestination() view returns(address)
func (_Trade *TradeCallerSession) GetProfitDestination() (common.Address, error) {
	return _Trade.Contract.GetProfitDestination(&_Trade.CallOpts)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_Trade *TradeTransactor) SetOwner(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _Trade.contract.Transact(opts, "setOwner", _owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_Trade *TradeSession) SetOwner(_owner common.Address) (*types.Transaction, error) {
	return _Trade.Contract.SetOwner(&_Trade.TransactOpts, _owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_Trade *TradeTransactorSession) SetOwner(_owner common.Address) (*types.Transaction, error) {
	return _Trade.Contract.SetOwner(&_Trade.TransactOpts, _owner)
}

// SetProfitDestination is a paid mutator transaction binding the contract method 0xa6c70871.
//
// Solidity: function setProfitDestination(address _profitDestination) returns()
func (_Trade *TradeTransactor) SetProfitDestination(opts *bind.TransactOpts, _profitDestination common.Address) (*types.Transaction, error) {
	return _Trade.contract.Transact(opts, "setProfitDestination", _profitDestination)
}

// SetProfitDestination is a paid mutator transaction binding the contract method 0xa6c70871.
//
// Solidity: function setProfitDestination(address _profitDestination) returns()
func (_Trade *TradeSession) SetProfitDestination(_profitDestination common.Address) (*types.Transaction, error) {
	return _Trade.Contract.SetProfitDestination(&_Trade.TransactOpts, _profitDestination)
}

// SetProfitDestination is a paid mutator transaction binding the contract method 0xa6c70871.
//
// Solidity: function setProfitDestination(address _profitDestination) returns()
func (_Trade *TradeTransactorSession) SetProfitDestination(_profitDestination common.Address) (*types.Transaction, error) {
	return _Trade.Contract.SetProfitDestination(&_Trade.TransactOpts, _profitDestination)
}

// Swap is a paid mutator transaction binding the contract method 0xa9678a18.
//
// Solidity: function swap(address _sourceDexAddress, address _destinationDexAddress, address _tokenAddress, uint256 _amount) returns()
func (_Trade *TradeTransactor) Swap(opts *bind.TransactOpts, _sourceDexAddress common.Address, _destinationDexAddress common.Address, _tokenAddress common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Trade.contract.Transact(opts, "swap", _sourceDexAddress, _destinationDexAddress, _tokenAddress, _amount)
}

// Swap is a paid mutator transaction binding the contract method 0xa9678a18.
//
// Solidity: function swap(address _sourceDexAddress, address _destinationDexAddress, address _tokenAddress, uint256 _amount) returns()
func (_Trade *TradeSession) Swap(_sourceDexAddress common.Address, _destinationDexAddress common.Address, _tokenAddress common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Trade.Contract.Swap(&_Trade.TransactOpts, _sourceDexAddress, _destinationDexAddress, _tokenAddress, _amount)
}

// Swap is a paid mutator transaction binding the contract method 0xa9678a18.
//
// Solidity: function swap(address _sourceDexAddress, address _destinationDexAddress, address _tokenAddress, uint256 _amount) returns()
func (_Trade *TradeTransactorSession) Swap(_sourceDexAddress common.Address, _destinationDexAddress common.Address, _tokenAddress common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Trade.Contract.Swap(&_Trade.TransactOpts, _sourceDexAddress, _destinationDexAddress, _tokenAddress, _amount)
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
