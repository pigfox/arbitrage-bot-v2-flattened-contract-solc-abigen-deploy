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

// SUNTokenMetaData contains all meta data concerning the SUNToken contract.
var SUNTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"INITIAL_SUPPLY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOKEN_DECIMALS\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOKEN_NAME\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOKEN_SYMBOL\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040518060400160405280600981526020016829aaa7102a37b5b2b760b91b8152506040518060400160405280600381526020016229aaa760e91b8152508160039081620000619190620002bd565b506004620000708282620002bd565b506200009d9150339050620000886012600a6200049e565b6200009790620f4240620004b3565b620000a3565b620004e3565b6001600160a01b038216620000d35760405163ec442f0560e01b8152600060048201526024015b60405180910390fd5b620000e160008383620000e5565b5050565b6001600160a01b03831662000114578060026000828254620001089190620004cd565b90915550620001889050565b6001600160a01b03831660009081526020819052604090205481811015620001695760405163391434e360e21b81526001600160a01b03851660048201526024810182905260448101839052606401620000ca565b6001600160a01b03841660009081526020819052604090209082900390555b6001600160a01b038216620001a657600280548290039055620001c5565b6001600160a01b03821660009081526020819052604090208054820190555b816001600160a01b0316836001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516200020b91815260200190565b60405180910390a3505050565b634e487b7160e01b600052604160045260246000fd5b600181811c908216806200024357607f821691505b6020821081036200026457634e487b7160e01b600052602260045260246000fd5b50919050565b601f821115620002b857600081815260208120601f850160051c81016020861015620002935750805b601f850160051c820191505b81811015620002b4578281556001016200029f565b5050505b505050565b81516001600160401b03811115620002d957620002d962000218565b620002f181620002ea84546200022e565b846200026a565b602080601f831160018114620003295760008415620003105750858301515b600019600386901b1c1916600185901b178555620002b4565b600085815260208120601f198616915b828110156200035a5788860151825594840194600190910190840162000339565b5085821015620003795787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b634e487b7160e01b600052601160045260246000fd5b600181815b80851115620003e0578160001904821115620003c457620003c462000389565b80851615620003d257918102915b93841c9390800290620003a4565b509250929050565b600082620003f95750600162000498565b81620004085750600062000498565b81600181146200042157600281146200042c576200044c565b600191505062000498565b60ff84111562000440576200044062000389565b50506001821b62000498565b5060208310610133831016604e8410600b841016171562000471575081810a62000498565b6200047d83836200039f565b806000190482111562000494576200049462000389565b0290505b92915050565b6000620004ac8383620003e8565b9392505050565b808202811582820484141762000498576200049862000389565b8082018082111562000498576200049862000389565b6108e780620004f36000396000f3fe608060405234801561001057600080fd5b50600436106100cf5760003560e01c80632ff2e9dc1161008c57806370a082311161006657806370a08231146101a957806395d89b41146101d2578063a9059cbb146101da578063dd62ed3e146101ed57600080fd5b80632ff2e9dc14610184578063313ce5671461018c5780635b7f415c146101a157600080fd5b806306fdde03146100d4578063095ea7b3146100f257806318160ddd14610115578063188214001461012757806323b872dd1461014f5780632a90531814610162575b600080fd5b6100dc610226565b6040516100e99190610622565b60405180910390f35b61010561010036600461068c565b6102b8565b60405190151581526020016100e9565b6002545b6040519081526020016100e9565b6100dc6040518060400160405280600981526020016829aaa7102a37b5b2b760b91b81525081565b61010561015d3660046106b6565b6102d2565b6100dc6040518060400160405280600381526020016229aaa760e91b81525081565b6101196102f6565b60125b60405160ff90911681526020016100e9565b61018f601281565b6101196101b73660046106f2565b6001600160a01b031660009081526020819052604090205490565b6100dc610312565b6101056101e836600461068c565b610321565b6101196101fb366004610714565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b60606003805461023590610747565b80601f016020809104026020016040519081016040528092919081815260200182805461026190610747565b80156102ae5780601f10610283576101008083540402835291602001916102ae565b820191906000526020600020905b81548152906001019060200180831161029157829003601f168201915b5050505050905090565b6000336102c681858561032f565b60019150505b92915050565b6000336102e0858285610341565b6102eb8585856103c4565b506001949350505050565b6103026012600a61087b565b61030f90620f4240610887565b81565b60606004805461023590610747565b6000336102c68185856103c4565b61033c8383836001610423565b505050565b6001600160a01b0383811660009081526001602090815260408083209386168352929052205460001981146103be57818110156103af57604051637dc7a0d960e11b81526001600160a01b038416600482015260248101829052604481018390526064015b60405180910390fd5b6103be84848484036000610423565b50505050565b6001600160a01b0383166103ee57604051634b637e8f60e11b8152600060048201526024016103a6565b6001600160a01b0382166104185760405163ec442f0560e01b8152600060048201526024016103a6565b61033c8383836104f8565b6001600160a01b03841661044d5760405163e602df0560e01b8152600060048201526024016103a6565b6001600160a01b03831661047757604051634a1406b160e11b8152600060048201526024016103a6565b6001600160a01b03808516600090815260016020908152604080832093871683529290522082905580156103be57826001600160a01b0316846001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040516104ea91815260200190565b60405180910390a350505050565b6001600160a01b038316610523578060026000828254610518919061089e565b909155506105959050565b6001600160a01b038316600090815260208190526040902054818110156105765760405163391434e360e21b81526001600160a01b038516600482015260248101829052604481018390526064016103a6565b6001600160a01b03841660009081526020819052604090209082900390555b6001600160a01b0382166105b1576002805482900390556105d0565b6001600160a01b03821660009081526020819052604090208054820190555b816001600160a01b0316836001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8360405161061591815260200190565b60405180910390a3505050565b600060208083528351808285015260005b8181101561064f57858101830151858201604001528201610633565b506000604082860101526040601f19601f8301168501019250505092915050565b80356001600160a01b038116811461068757600080fd5b919050565b6000806040838503121561069f57600080fd5b6106a883610670565b946020939093013593505050565b6000806000606084860312156106cb57600080fd5b6106d484610670565b92506106e260208501610670565b9150604084013590509250925092565b60006020828403121561070457600080fd5b61070d82610670565b9392505050565b6000806040838503121561072757600080fd5b61073083610670565b915061073e60208401610670565b90509250929050565b600181811c9082168061075b57607f821691505b60208210810361077b57634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052601160045260246000fd5b600181815b808511156107d25781600019048211156107b8576107b8610781565b808516156107c557918102915b93841c939080029061079c565b509250929050565b6000826107e9575060016102cc565b816107f6575060006102cc565b816001811461080c576002811461081657610832565b60019150506102cc565b60ff84111561082757610827610781565b50506001821b6102cc565b5060208310610133831016604e8410600b8410161715610855575081810a6102cc565b61085f8383610797565b806000190482111561087357610873610781565b029392505050565b600061070d83836107da565b80820281158282048414176102cc576102cc610781565b808201808211156102cc576102cc61078156fea2646970667358221220d4acd9868a760db5cdad8df9117df313ced1d11cf500575f2059b2bbb926cedf64736f6c63430008140033",
}

// SUNTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use SUNTokenMetaData.ABI instead.
var SUNTokenABI = SUNTokenMetaData.ABI

// SUNTokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SUNTokenMetaData.Bin instead.
var SUNTokenBin = SUNTokenMetaData.Bin

// DeploySUNToken deploys a new Ethereum contract, binding an instance of SUNToken to it.
func DeploySUNToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SUNToken, error) {
	parsed, err := SUNTokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SUNTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SUNToken{SUNTokenCaller: SUNTokenCaller{contract: contract}, SUNTokenTransactor: SUNTokenTransactor{contract: contract}, SUNTokenFilterer: SUNTokenFilterer{contract: contract}}, nil
}

// SUNToken is an auto generated Go binding around an Ethereum contract.
type SUNToken struct {
	SUNTokenCaller     // Read-only binding to the contract
	SUNTokenTransactor // Write-only binding to the contract
	SUNTokenFilterer   // Log filterer for contract events
}

// SUNTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type SUNTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SUNTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SUNTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SUNTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SUNTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SUNTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SUNTokenSession struct {
	Contract     *SUNToken         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SUNTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SUNTokenCallerSession struct {
	Contract *SUNTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SUNTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SUNTokenTransactorSession struct {
	Contract     *SUNTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SUNTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type SUNTokenRaw struct {
	Contract *SUNToken // Generic contract binding to access the raw methods on
}

// SUNTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SUNTokenCallerRaw struct {
	Contract *SUNTokenCaller // Generic read-only contract binding to access the raw methods on
}

// SUNTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SUNTokenTransactorRaw struct {
	Contract *SUNTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSUNToken creates a new instance of SUNToken, bound to a specific deployed contract.
func NewSUNToken(address common.Address, backend bind.ContractBackend) (*SUNToken, error) {
	contract, err := bindSUNToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SUNToken{SUNTokenCaller: SUNTokenCaller{contract: contract}, SUNTokenTransactor: SUNTokenTransactor{contract: contract}, SUNTokenFilterer: SUNTokenFilterer{contract: contract}}, nil
}

// NewSUNTokenCaller creates a new read-only instance of SUNToken, bound to a specific deployed contract.
func NewSUNTokenCaller(address common.Address, caller bind.ContractCaller) (*SUNTokenCaller, error) {
	contract, err := bindSUNToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SUNTokenCaller{contract: contract}, nil
}

// NewSUNTokenTransactor creates a new write-only instance of SUNToken, bound to a specific deployed contract.
func NewSUNTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*SUNTokenTransactor, error) {
	contract, err := bindSUNToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SUNTokenTransactor{contract: contract}, nil
}

// NewSUNTokenFilterer creates a new log filterer instance of SUNToken, bound to a specific deployed contract.
func NewSUNTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*SUNTokenFilterer, error) {
	contract, err := bindSUNToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SUNTokenFilterer{contract: contract}, nil
}

// bindSUNToken binds a generic wrapper to an already deployed contract.
func bindSUNToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SUNTokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SUNToken *SUNTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SUNToken.Contract.SUNTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SUNToken *SUNTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SUNToken.Contract.SUNTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SUNToken *SUNTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SUNToken.Contract.SUNTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SUNToken *SUNTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SUNToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SUNToken *SUNTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SUNToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SUNToken *SUNTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SUNToken.Contract.contract.Transact(opts, method, params...)
}

// INITIALSUPPLY is a free data retrieval call binding the contract method 0x2ff2e9dc.
//
// Solidity: function INITIAL_SUPPLY() view returns(uint256)
func (_SUNToken *SUNTokenCaller) INITIALSUPPLY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SUNToken.contract.Call(opts, &out, "INITIAL_SUPPLY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// INITIALSUPPLY is a free data retrieval call binding the contract method 0x2ff2e9dc.
//
// Solidity: function INITIAL_SUPPLY() view returns(uint256)
func (_SUNToken *SUNTokenSession) INITIALSUPPLY() (*big.Int, error) {
	return _SUNToken.Contract.INITIALSUPPLY(&_SUNToken.CallOpts)
}

// INITIALSUPPLY is a free data retrieval call binding the contract method 0x2ff2e9dc.
//
// Solidity: function INITIAL_SUPPLY() view returns(uint256)
func (_SUNToken *SUNTokenCallerSession) INITIALSUPPLY() (*big.Int, error) {
	return _SUNToken.Contract.INITIALSUPPLY(&_SUNToken.CallOpts)
}

// TOKENDECIMALS is a free data retrieval call binding the contract method 0x5b7f415c.
//
// Solidity: function TOKEN_DECIMALS() view returns(uint8)
func (_SUNToken *SUNTokenCaller) TOKENDECIMALS(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _SUNToken.contract.Call(opts, &out, "TOKEN_DECIMALS")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// TOKENDECIMALS is a free data retrieval call binding the contract method 0x5b7f415c.
//
// Solidity: function TOKEN_DECIMALS() view returns(uint8)
func (_SUNToken *SUNTokenSession) TOKENDECIMALS() (uint8, error) {
	return _SUNToken.Contract.TOKENDECIMALS(&_SUNToken.CallOpts)
}

// TOKENDECIMALS is a free data retrieval call binding the contract method 0x5b7f415c.
//
// Solidity: function TOKEN_DECIMALS() view returns(uint8)
func (_SUNToken *SUNTokenCallerSession) TOKENDECIMALS() (uint8, error) {
	return _SUNToken.Contract.TOKENDECIMALS(&_SUNToken.CallOpts)
}

// TOKENNAME is a free data retrieval call binding the contract method 0x18821400.
//
// Solidity: function TOKEN_NAME() view returns(string)
func (_SUNToken *SUNTokenCaller) TOKENNAME(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SUNToken.contract.Call(opts, &out, "TOKEN_NAME")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TOKENNAME is a free data retrieval call binding the contract method 0x18821400.
//
// Solidity: function TOKEN_NAME() view returns(string)
func (_SUNToken *SUNTokenSession) TOKENNAME() (string, error) {
	return _SUNToken.Contract.TOKENNAME(&_SUNToken.CallOpts)
}

// TOKENNAME is a free data retrieval call binding the contract method 0x18821400.
//
// Solidity: function TOKEN_NAME() view returns(string)
func (_SUNToken *SUNTokenCallerSession) TOKENNAME() (string, error) {
	return _SUNToken.Contract.TOKENNAME(&_SUNToken.CallOpts)
}

// TOKENSYMBOL is a free data retrieval call binding the contract method 0x2a905318.
//
// Solidity: function TOKEN_SYMBOL() view returns(string)
func (_SUNToken *SUNTokenCaller) TOKENSYMBOL(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SUNToken.contract.Call(opts, &out, "TOKEN_SYMBOL")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TOKENSYMBOL is a free data retrieval call binding the contract method 0x2a905318.
//
// Solidity: function TOKEN_SYMBOL() view returns(string)
func (_SUNToken *SUNTokenSession) TOKENSYMBOL() (string, error) {
	return _SUNToken.Contract.TOKENSYMBOL(&_SUNToken.CallOpts)
}

// TOKENSYMBOL is a free data retrieval call binding the contract method 0x2a905318.
//
// Solidity: function TOKEN_SYMBOL() view returns(string)
func (_SUNToken *SUNTokenCallerSession) TOKENSYMBOL() (string, error) {
	return _SUNToken.Contract.TOKENSYMBOL(&_SUNToken.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_SUNToken *SUNTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SUNToken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_SUNToken *SUNTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _SUNToken.Contract.Allowance(&_SUNToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_SUNToken *SUNTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _SUNToken.Contract.Allowance(&_SUNToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_SUNToken *SUNTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SUNToken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_SUNToken *SUNTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _SUNToken.Contract.BalanceOf(&_SUNToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_SUNToken *SUNTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _SUNToken.Contract.BalanceOf(&_SUNToken.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_SUNToken *SUNTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _SUNToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_SUNToken *SUNTokenSession) Decimals() (uint8, error) {
	return _SUNToken.Contract.Decimals(&_SUNToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_SUNToken *SUNTokenCallerSession) Decimals() (uint8, error) {
	return _SUNToken.Contract.Decimals(&_SUNToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SUNToken *SUNTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SUNToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SUNToken *SUNTokenSession) Name() (string, error) {
	return _SUNToken.Contract.Name(&_SUNToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SUNToken *SUNTokenCallerSession) Name() (string, error) {
	return _SUNToken.Contract.Name(&_SUNToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SUNToken *SUNTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SUNToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SUNToken *SUNTokenSession) Symbol() (string, error) {
	return _SUNToken.Contract.Symbol(&_SUNToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SUNToken *SUNTokenCallerSession) Symbol() (string, error) {
	return _SUNToken.Contract.Symbol(&_SUNToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_SUNToken *SUNTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SUNToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_SUNToken *SUNTokenSession) TotalSupply() (*big.Int, error) {
	return _SUNToken.Contract.TotalSupply(&_SUNToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_SUNToken *SUNTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _SUNToken.Contract.TotalSupply(&_SUNToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_SUNToken *SUNTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _SUNToken.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_SUNToken *SUNTokenSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _SUNToken.Contract.Approve(&_SUNToken.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_SUNToken *SUNTokenTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _SUNToken.Contract.Approve(&_SUNToken.TransactOpts, spender, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_SUNToken *SUNTokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _SUNToken.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_SUNToken *SUNTokenSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _SUNToken.Contract.Transfer(&_SUNToken.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_SUNToken *SUNTokenTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _SUNToken.Contract.Transfer(&_SUNToken.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_SUNToken *SUNTokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _SUNToken.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_SUNToken *SUNTokenSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _SUNToken.Contract.TransferFrom(&_SUNToken.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_SUNToken *SUNTokenTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _SUNToken.Contract.TransferFrom(&_SUNToken.TransactOpts, from, to, value)
}

// SUNTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the SUNToken contract.
type SUNTokenApprovalIterator struct {
	Event *SUNTokenApproval // Event containing the contract specifics and raw log

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
func (it *SUNTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SUNTokenApproval)
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
		it.Event = new(SUNTokenApproval)
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
func (it *SUNTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SUNTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SUNTokenApproval represents a Approval event raised by the SUNToken contract.
type SUNTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_SUNToken *SUNTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*SUNTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _SUNToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &SUNTokenApprovalIterator{contract: _SUNToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_SUNToken *SUNTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *SUNTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _SUNToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SUNTokenApproval)
				if err := _SUNToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_SUNToken *SUNTokenFilterer) ParseApproval(log types.Log) (*SUNTokenApproval, error) {
	event := new(SUNTokenApproval)
	if err := _SUNToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SUNTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the SUNToken contract.
type SUNTokenTransferIterator struct {
	Event *SUNTokenTransfer // Event containing the contract specifics and raw log

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
func (it *SUNTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SUNTokenTransfer)
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
		it.Event = new(SUNTokenTransfer)
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
func (it *SUNTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SUNTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SUNTokenTransfer represents a Transfer event raised by the SUNToken contract.
type SUNTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_SUNToken *SUNTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*SUNTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SUNToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &SUNTokenTransferIterator{contract: _SUNToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_SUNToken *SUNTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *SUNTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SUNToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SUNTokenTransfer)
				if err := _SUNToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_SUNToken *SUNTokenFilterer) ParseTransfer(log types.Log) (*SUNTokenTransfer, error) {
	event := new(SUNTokenTransfer)
	if err := _SUNToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
