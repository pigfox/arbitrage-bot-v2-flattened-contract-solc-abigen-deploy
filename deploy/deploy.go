package deploy

import (
	"arbitrage-bot-v2-flattened-contract-solc-abigen-deploy/connection"
	"arbitrage-bot-v2-flattened-contract-solc-abigen-deploy/util"
	"arbitrage-bot-v2-flattened-contract-solc-abigen-deploy/wallet"
	"bytes"
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/jimlawless/whereami"
	"log"
	"math/big"
	"time"
)

func Run(contractName string) {
	fmt.Println("Contract deploying...")
	privateKey, err := crypto.HexToECDSA(wallet.Sepolia.PrivateKey)
	if err != nil {
		log.Fatal(err.Error() + " " + whereami.WhereAmI())
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Not OK " + whereami.WhereAmI())
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := connection.RPC.Client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err.Error() + " " + whereami.WhereAmI())
	}
	/*
		pendingBalance, err := connection.RPC.Client.PendingBalanceAt(context.Background(), fromAddress)
		if err != nil {
			log.Fatal(err.Error() + " " + whereami.WhereAmI())
		}
		fmt.Println("pendingBalance: ", pendingBalance)
	*/
	// Set the gas price and gas limit
	suggestedGasPrice, err := connection.RPC.Client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err.Error() + " " + whereami.WhereAmI())
	}
	fmt.Println("suggestedGasPrice: ", suggestedGasPrice)
	/*
		auth, err := bind.NewKeyedTransactorWithChainID(privateKey, config.Map[config.Active].ChainID)
		if err != nil {
			log.Fatal(err.Error() + " " + whereami.WhereAmI())
		}
	*/
	/*
		1349 May 31 11:32 Base.abi
		2856 May 31 11:32 Base.bin
	*/
	//https://ethereum.stackexchange.com/questions/39401/how-do-you-calculate-gas-limit-for-transaction-with-data-in-ethereum
	contractString := util.ContractBytes(contractName)
	fmt.Println("contractString: ", contractString)
	contractBytecode := common.FromHex(contractString)

	// Estimate gas limit
	msg := ethereum.CallMsg{
		From:     fromAddress,
		To:       nil, // To is nil for contract deployment
		Gas:      0,
		GasPrice: suggestedGasPrice,
		Value:    big.NewInt(0),
		Data:     contractBytecode,
	}

	gasLimit, err := connection.RPC.Client.EstimateGas(context.Background(), msg)
	if err != nil {
		log.Fatalf("Failed to estimate gas: %v", err)
	}

	fmt.Printf("Gas Price: %s\n", suggestedGasPrice.String())
	fmt.Printf("Gas Limit: %d\n", gasLimit)

	// Example of creating a transaction for contract deployment
	chainID, err := connection.RPC.Client.NetworkID(context.Background())
	if err != nil {
		log.Fatalf("Failed to get network ID: %v", err)
	}

	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		To:       nil,
		Value:    big.NewInt(0),
		Gas:      gasLimit,
		GasPrice: suggestedGasPrice,
		Data:     contractBytecode})

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatalf("Failed to sign transaction: %v", err)
	}

	// Send the transaction
	err = connection.RPC.Client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatalf("Failed to send transaction: %v", err)
	}

	/*
			blockGasLimit := getBlockGasLimit()
			fmt.Println("blockGasLimit: ", blockGasLimit)
			if auth.GasLimit < blockGasLimit {
				auth.GasLimit = blockGasLimit
				fmt.Println("using block gas limit: ", blockGasLimit)
			}

			auth.Nonce = big.NewInt(int64(nonce))
			auth.Value = big.NewInt(0)
			auth.GasPrice = util.AddPercentBigInt(suggestedGasPrice, 30)

			fmt.Println("auth.GasLimit: ", auth.GasLimit)
			fmt.Println("auth.GasPrice: ", auth.GasPrice)
			fmt.Println("nonce: ", auth.Nonce)

			//address, tx, instance, err := api.DeployPigfox(auth, connection.RPC.Client)
			address, tx, instance, err := api.DeployBase(auth, connection.RPC.Client)
			if err != nil {
				log.Fatal(err.Error() + " " + whereami.WhereAmI())
			}

		structs.OnChainContract.Address = address.Hex()
		structs.OnChainContract.TxHash = tx.Hash().Hex()
		structs.OnChainContract.TxCost = tx.Cost()
		structs.OnChainContract.Gas = tx.Gas()
		structs.OnChainContract.GasPrice = tx.GasPrice()
	*/
	fmt.Println(signedTx.Hash().Hex(), "waiting to be mined...")
	startTime := time.Now()
	receipt, err := bind.WaitMined(context.Background(), connection.RPC.Client, signedTx)
	if err != nil {
		log.Fatal(err.Error() + " " + whereami.WhereAmI())
	}
	fmt.Println("Mining took", time.Since(startTime))
	//fmt.Println("Token deployed at:", tx.address)
	fmt.Println("Token tx:", tx.Hash())
	fmt.Println("Token cost:", tx.Cost())
	fmt.Println("Token gas limit:", tx.Gas())
	fmt.Println("Token gas price:", tx.GasPrice())

	if receipt.Status == types.ReceiptStatusSuccessful {
		fmt.Println("Token deployed successfully")
	} else {
		reason, err := getRevertReason(receipt)
		if err != nil {
			log.Fatal(err.Error() + whereami.WhereAmI())
		}
		log.Fatal(reason + " " + whereami.WhereAmI())
	}

	//_ = instance
}

func getRevertReason(receipt *types.Receipt) (string, error) {
	if len(receipt.Logs) == 0 {
		return "", nil
	}

	// Get the revert event signature
	revertEventSig := []byte("Error(string)")

	// Iterate through the receipt logs to find the revert event
	for _, receiptLog := range receipt.Logs {
		if len(receiptLog.Topics) > 0 && bytes.Equal(receiptLog.Topics[0].Bytes(), revertEventSig) {
			if len(receiptLog.Data) < 32 {
				return "", fmt.Errorf("invalid event data")
			}

			// Extract the revert reason from the event data
			revertReason := string(receiptLog.Data[32:])
			return revertReason, nil
		}
	}

	return "", nil
}
