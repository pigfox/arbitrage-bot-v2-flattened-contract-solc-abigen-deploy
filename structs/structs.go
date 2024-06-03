package structs

import (
	"math/big"
)

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
}

type Transaction struct {
	GasPrice *big.Int
	GasLimit uint64
	Cost     *big.Int
	Hash     string
}
