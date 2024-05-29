package structs

import (
	"math/big"
)

var OnChainContract DeployedContract

type DeployedContract struct {
	Address  string
	TxHash   string
	TxCost   *big.Int
	Gas      uint64
	GasPrice *big.Int
}
