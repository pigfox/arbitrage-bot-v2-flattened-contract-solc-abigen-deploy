package structs

import (
	"math/big"
)

const ContractName string = "Pigfox" //file name w/o _flat or .sol

var OnChainContract DeployedContract

type DeployedContract struct {
	Address  string
	TxHash   string
	TxCost   *big.Int
	Gas      uint64
	GasPrice *big.Int
}
