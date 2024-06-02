package deploy

import (
	"arbitrage-bot-v2-flattened-contract-solc-abigen-deploy/connection"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
	"time"
)

func network(txHashStr string) {
	// Check the number of pending transactions in the network
	pendingCount, err := connection.RPC.Client.PendingTransactionCount(context.Background())
	if err != nil {
		log.Fatalf("Failed to get pending transaction count: %v", err)
	}
	fmt.Printf("Pending transactions: %d\n", pendingCount)

	// Get the current gas price
	gasPrice, err := connection.RPC.Client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("Failed to suggest gas price: %v", err)
	}
	fmt.Printf("Suggested gas price: %s\n", gasPrice.String())

	// Your transaction hash
	txHash := common.HexToHash(txHashStr)

	// Check if the transaction is still in the pool
	// Check if the transaction is still in the pool
	_, isPending, err := connection.RPC.Client.TransactionByHash(context.Background(), txHash) //tx
	if err == ethereum.NotFound {
		fmt.Printf("Transaction %s not found\n", txHash.Hex())
	} else if err != nil {
		log.Fatalf("Failed to get transaction by hash: %v", err)
	} else {
		fmt.Printf("Transaction %s is pending: %v\n", txHash.Hex(), isPending)
		if !isPending {
			receipt, err := connection.RPC.Client.TransactionReceipt(context.Background(), txHash)
			if err == ethereum.NotFound {
				fmt.Printf("Transaction %s receipt not found\n", txHash.Hex())
			} else if err != nil {
				log.Fatalf("Failed to get transaction receipt: %v", err)
			} else {
				fmt.Printf("Transaction %s status: %v\n", txHash.Hex(), receipt.Status)
			}
		}
	}

	// Continuously monitor the transaction status
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()
	for range ticker.C {
		_, isPending, err := connection.RPC.Client.TransactionByHash(context.Background(), txHash) //tx
		if err == ethereum.NotFound {
			fmt.Printf("Transaction %s still not found\n", txHash.Hex())
		} else if err != nil {
			log.Fatalf("Failed to get transaction by hash: %v", err)
		} else {
			fmt.Printf("Transaction %s is pending: %v\n", txHash.Hex(), isPending)
			if !isPending {
				receipt, err := connection.RPC.Client.TransactionReceipt(context.Background(), txHash)
				if err == ethereum.NotFound {
					fmt.Printf("Transaction %s receipt not found\n", txHash.Hex())
				} else if err != nil {
					log.Fatalf("Failed to get transaction receipt: %v", err)
				} else {
					fmt.Printf("Transaction %s status: %v\n", txHash.Hex(), receipt.Status)
					if receipt.Status == types.ReceiptStatusSuccessful {
						fmt.Printf("Transaction %s mined successfully\n", txHash.Hex())
						break
					} else {
						fmt.Printf("Transaction %s failed\n", txHash.Hex())
						break
					}
				}
			}
		}
	}
}
