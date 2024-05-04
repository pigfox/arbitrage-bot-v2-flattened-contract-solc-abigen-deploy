package code

import (
	"arbitrage-bot-v2-flattened-contract-solc-abigen-deploy/connection"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
	"strings"
)

func subscribeEvent() {
	contractAddress := common.HexToAddress(currentAddress)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)
	sub, err := connection.RPC.Client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatalf("Failed to subscribe to filter logs: %v", err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(abiMap[currentAddress])) // Replace `ABI_STRING_HERE` with the actual ABI of the contract
	if err != nil {
		log.Fatalf("Failed to parse ABI: %v", err)
	}

	fmt.Println("Listening to logs...")
	for {
		select {
		case err := <-sub.Err():
			log.Fatalf("Subscription error: %v", err)
		case vLog := <-logs:
			fmt.Printf("New Log on Block #%d, Tx Hash: %s\n", vLog.BlockNumber, vLog.TxHash.Hex())
			processLog(&contractAbi, vLog)
		}
	}
}
