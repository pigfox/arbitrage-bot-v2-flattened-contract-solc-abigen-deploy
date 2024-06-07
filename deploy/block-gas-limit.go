package deploy

import (
	"arbitrage-bot-v2-flattened-contract-solc-abigen-deploy/connection"
	"context"
	"log"
)

func getBlockGasLimit() uint64 {
	// Get the latest block number
	header, err := connection.RPC.Client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	blockNumber := header.Number
	_ = blockNumber

	// Get the latest block
	block, err := connection.RPC.Client.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	return block.GasLimit()
}
