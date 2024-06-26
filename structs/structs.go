package structs

import (
	"math/big"
)

var DeploymentParams map[string]DeploymentParam
var OnChainContract DeployedContract

type DeployedContract struct {
	Receipt TxReceipt
	Tx      Transaction
}

type TxReceipt struct {
	Address string
	Hash    string
	Status  string
	GasUsed uint64
	Block   *big.Int
}

type Transaction struct {
	GasPrice *big.Int
	GasLimit uint64
	Cost     *big.Int
	Hash     string
}

type DeploymentParam struct {
	Name                 string
	EVM                  string
	CompilerVersion      string
	ConstructorArguments []interface{}
}
