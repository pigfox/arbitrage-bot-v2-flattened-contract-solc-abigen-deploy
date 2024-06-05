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
	Bin: "0x608060405234801562000010575f80fd5b506040518060400160405280600981526020016829aaa7102a37b5b2b760b91b8152506040518060400160405280600381526020016229aaa760e91b8152508160039081620000609190620002af565b5060046200006f8282620002af565b506200009c9150339050620000876012600a6200048a565b6200009690620f42406200049e565b620000a2565b620004ce565b6001600160a01b038216620000d15760405163ec442f0560e01b81525f60048201526024015b60405180910390fd5b620000de5f8383620000e2565b5050565b6001600160a01b03831662000110578060025f828254620001049190620004b8565b90915550620001829050565b6001600160a01b0383165f9081526020819052604090205481811015620001645760405163391434e360e21b81526001600160a01b03851660048201526024810182905260448101839052606401620000c8565b6001600160a01b0384165f9081526020819052604090209082900390555b6001600160a01b038216620001a057600280548290039055620001be565b6001600160a01b0382165f9081526020819052604090208054820190555b816001600160a01b0316836001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516200020491815260200190565b60405180910390a3505050565b634e487b7160e01b5f52604160045260245ffd5b600181811c908216806200023a57607f821691505b6020821081036200025957634e487b7160e01b5f52602260045260245ffd5b50919050565b601f821115620002aa57805f5260205f20601f840160051c81016020851015620002865750805b601f840160051c820191505b81811015620002a7575f815560010162000292565b50505b505050565b81516001600160401b03811115620002cb57620002cb62000211565b620002e381620002dc845462000225565b846200025f565b602080601f83116001811462000319575f8415620003015750858301515b5f19600386901b1c1916600185901b17855562000373565b5f85815260208120601f198616915b82811015620003495788860151825594840194600190910190840162000328565b50858210156200036757878501515f19600388901b60f8161c191681555b505060018460011b0185555b505050505050565b634e487b7160e01b5f52601160045260245ffd5b600181815b80851115620003cf57815f1904821115620003b357620003b36200037b565b80851615620003c157918102915b93841c939080029062000394565b509250929050565b5f82620003e75750600162000484565b81620003f557505f62000484565b81600181146200040e5760028114620004195762000439565b600191505062000484565b60ff8411156200042d576200042d6200037b565b50506001821b62000484565b5060208310610133831016604e8410600b84101617156200045e575081810a62000484565b6200046a83836200038f565b805f19048211156200048057620004806200037b565b0290505b92915050565b5f620004978383620003d7565b9392505050565b80820281158282048414176200048457620004846200037b565b808201808211156200048457620004846200037b565b6108ba80620004dc5f395ff3fe608060405234801561000f575f80fd5b50600436106100cb575f3560e01c80632ff2e9dc1161008857806370a082311161006357806370a08231146101a457806395d89b41146101cc578063a9059cbb146101d4578063dd62ed3e146101e7575f80fd5b80632ff2e9dc1461017f578063313ce567146101875780635b7f415c1461019c575f80fd5b806306fdde03146100cf578063095ea7b3146100ed57806318160ddd14610110578063188214001461012257806323b872dd1461014a5780632a9053181461015d575b5f80fd5b6100d761021f565b6040516100e4919061060a565b60405180910390f35b6101006100fb366004610671565b6102af565b60405190151581526020016100e4565b6002545b6040519081526020016100e4565b6100d76040518060400160405280600981526020016829aaa7102a37b5b2b760b91b81525081565b610100610158366004610699565b6102c8565b6100d76040518060400160405280600381526020016229aaa760e91b81525081565b6101146102eb565b60125b60405160ff90911681526020016100e4565b61018a601281565b6101146101b23660046106d2565b6001600160a01b03165f9081526020819052604090205490565b6100d7610307565b6101006101e2366004610671565b610316565b6101146101f53660046106f2565b6001600160a01b039182165f90815260016020908152604080832093909416825291909152205490565b60606003805461022e90610723565b80601f016020809104026020016040519081016040528092919081815260200182805461025a90610723565b80156102a55780601f1061027c576101008083540402835291602001916102a5565b820191905f5260205f20905b81548152906001019060200180831161028857829003601f168201915b5050505050905090565b5f336102bc818585610323565b60019150505b92915050565b5f336102d5858285610335565b6102e08585856103b5565b506001949350505050565b6102f76012600a61084f565b61030490620f424061085a565b81565b60606004805461022e90610723565b5f336102bc8185856103b5565b6103308383836001610412565b505050565b6001600160a01b038381165f908152600160209081526040808320938616835292905220545f1981146103af57818110156103a157604051637dc7a0d960e11b81526001600160a01b038416600482015260248101829052604481018390526064015b60405180910390fd5b6103af84848484035f610412565b50505050565b6001600160a01b0383166103de57604051634b637e8f60e11b81525f6004820152602401610398565b6001600160a01b0382166104075760405163ec442f0560e01b81525f6004820152602401610398565b6103308383836104e4565b6001600160a01b03841661043b5760405163e602df0560e01b81525f6004820152602401610398565b6001600160a01b03831661046457604051634a1406b160e11b81525f6004820152602401610398565b6001600160a01b038085165f90815260016020908152604080832093871683529290522082905580156103af57826001600160a01b0316846001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040516104d691815260200190565b60405180910390a350505050565b6001600160a01b03831661050e578060025f8282546105039190610871565b9091555061057e9050565b6001600160a01b0383165f90815260208190526040902054818110156105605760405163391434e360e21b81526001600160a01b03851660048201526024810182905260448101839052606401610398565b6001600160a01b0384165f9081526020819052604090209082900390555b6001600160a01b03821661059a576002805482900390556105b8565b6001600160a01b0382165f9081526020819052604090208054820190555b816001600160a01b0316836001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516105fd91815260200190565b60405180910390a3505050565b5f602080835283518060208501525f5b818110156106365785810183015185820160400152820161061a565b505f604082860101526040601f19601f8301168501019250505092915050565b80356001600160a01b038116811461066c575f80fd5b919050565b5f8060408385031215610682575f80fd5b61068b83610656565b946020939093013593505050565b5f805f606084860312156106ab575f80fd5b6106b484610656565b92506106c260208501610656565b9150604084013590509250925092565b5f602082840312156106e2575f80fd5b6106eb82610656565b9392505050565b5f8060408385031215610703575f80fd5b61070c83610656565b915061071a60208401610656565b90509250929050565b600181811c9082168061073757607f821691505b60208210810361075557634e487b7160e01b5f52602260045260245ffd5b50919050565b634e487b7160e01b5f52601160045260245ffd5b600181815b808511156107a957815f190482111561078f5761078f61075b565b8085161561079c57918102915b93841c9390800290610774565b509250929050565b5f826107bf575060016102c2565b816107cb57505f6102c2565b81600181146107e157600281146107eb57610807565b60019150506102c2565b60ff8411156107fc576107fc61075b565b50506001821b6102c2565b5060208310610133831016604e8410600b841016171561082a575081810a6102c2565b610834838361076f565b805f19048211156108475761084761075b565b029392505050565b5f6106eb83836107b1565b80820281158282048414176102c2576102c261075b565b808201808211156102c2576102c261075b56fea2646970667358221220b0e1c866d87cd7c53bb12b5ec6f15f41d53e151b2c96eaadf1b6e3785c5f461b64736f6c63430008180033",
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
