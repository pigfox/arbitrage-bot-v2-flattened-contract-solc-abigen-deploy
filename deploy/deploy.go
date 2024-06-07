package deploy

import (
	"arbitrage-bot-v2-flattened-contract-solc-abigen-deploy/connection"
	"arbitrage-bot-v2-flattened-contract-solc-abigen-deploy/structs"
	"arbitrage-bot-v2-flattened-contract-solc-abigen-deploy/util"
	"arbitrage-bot-v2-flattened-contract-solc-abigen-deploy/wallet"
	"bytes"
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/jimlawless/whereami"
	"log"
	"math/big"
	"strings"
	"time"
)

func Run(params structs.DeploymentParam) {
	fmt.Println(params)
	value := big.NewInt(0)
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

	pendingBalance, err := connection.RPC.Client.PendingBalanceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err.Error() + " " + whereami.WhereAmI())
	}

	suggestedGasPrice, err := connection.RPC.Client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err.Error() + " " + whereami.WhereAmI())
	}
	actualGasPrice := new(big.Int).Mul(suggestedGasPrice, big.NewInt(20))
	//actualGasPrice = new(big.Int).Div(actualGasPrice, big.NewInt(10))

	fmt.Println("suggestedGasPrice: ", suggestedGasPrice)
	fmt.Println("actualGasPrice: ", actualGasPrice)

	//https://ethereum.stackexchange.com/questions/39401/how-do-you-calculate-gas-limit-for-transaction-with-data-in-ethereum
	contractString := util.GetContractBIN(params.Name)
	fmt.Println("contractString: ", contractString)
	contractBytecode := common.FromHex(contractString)

	contractABI, err := abi.JSON(strings.NewReader(util.GetContractABI(params.Name)))
	if err != nil {
		log.Fatalf("Failed to parse the contract ABI: %v", err)
	}
	fmt.Println("contractABI: ", contractABI)
	var data, constructorArguments []byte

	constructorArguments, err = contractABI.Pack("") //
	if err != nil {
		log.Fatalf("Failed to pack constructor arguments: %v", err)
	}

	data = append(contractBytecode, constructorArguments...)

	//fmt.Println("data: ", data)
	//fmt.Println("data: ", string(data))

	// Estimate gas limit
	msg := ethereum.CallMsg{
		From:     fromAddress,
		To:       nil, // To is nil for contract deployment
		Gas:      0,
		GasPrice: actualGasPrice,
		Value:    value,
		Data:     data,
	}

	gasLimit, err := connection.RPC.Client.EstimateGas(context.Background(), msg)
	if err != nil {
		log.Printf("Failed to estimate gas: %v", err)
		gasLimit = getBlockGasLimit()
	}

	//check sufficient funds
	cost := new(big.Int).Mul(actualGasPrice, big.NewInt(int64(gasLimit)))
	fmt.Println("pendingBalance: ", pendingBalance)
	fmt.Println("cost: ", cost)
	fmt.Println("ETH cost: ", util.WeiToETH(cost))
	if pendingBalance.Cmp(cost) < 0 { //pendingBalance < cost
		fmt.Println("Insufficient funds, have ", pendingBalance, " need ", cost, " ", whereami.WhereAmI())
	}

	fmt.Printf("Gas Price Gwei: %d\n", util.WeiToGwei(actualGasPrice))
	fmt.Printf("Gas Limit: %d\n", gasLimit)

	chainID, err := connection.RPC.Client.NetworkID(context.Background())
	if err != nil {
		log.Fatalf("Failed to get network ID: %v", err)
	}

	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		To:       nil,
		Value:    value,
		Gas:      gasLimit,
		GasPrice: actualGasPrice,
		Data:     data})

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatalf("Failed to sign transaction: %v", err)
	}

	// Send the transaction
	err = connection.RPC.Client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatalf("Failed to send transaction: %v", err)
	}

	fmt.Println(signedTx.Hash().Hex(), "waiting to be mined...")
	go network(signedTx.Hash().Hex())
	startTime := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Hour)
	defer cancel()
	receipt, err := bind.WaitMined(ctx, connection.RPC.Client, signedTx)
	if err != nil {
		log.Fatal(err.Error() + " " + whereami.WhereAmI())
	}
	fmt.Println("Mining took", time.Since(startTime))
	fmt.Println("Token tx:", tx.Hash(), "?????")
	fmt.Println("Token cost:", tx.Cost())
	fmt.Println("Token gas limit:", tx.Gas())
	fmt.Println("Token gas price:", tx.GasPrice())
	structs.OnChainContract.Tx.GasLimit = tx.Gas()
	structs.OnChainContract.Tx.GasPrice = tx.GasPrice()
	structs.OnChainContract.Tx.Hash = tx.Hash().Hex()
	structs.OnChainContract.Tx.Cost = tx.Cost()

	if receipt.Status == types.ReceiptStatusSuccessful {
		structs.OnChainContract.Receipt.Address = receipt.ContractAddress.Hex()
		structs.OnChainContract.Receipt.GasUsed = receipt.GasUsed
		structs.OnChainContract.Receipt.Hash = receipt.TxHash.Hex()
		structs.OnChainContract.Receipt.Block = receipt.BlockNumber
		fmt.Println("Token deployed successfully at", receipt.ContractAddress)
	} else {
		reason, err := getRevertReason(receipt)
		if err != nil {
			log.Fatal(err.Error() + whereami.WhereAmI())
		}

		if len(reason) == 0 {
			fmt.Println("No receipt reason found " + whereami.WhereAmI())
			network(receipt.TxHash.Hex())
		} else {
			log.Fatal(reason + " " + whereami.WhereAmI())
		}
	}
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
