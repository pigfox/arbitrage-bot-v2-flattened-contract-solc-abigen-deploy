package deploy

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"flattened-contract-solc-abigen/api"
	"flattened-contract-solc-abigen/config"
	"flattened-contract-solc-abigen/connection"
	"flattened-contract-solc-abigen/wallet"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/jimlawless/whereami"
	"log"
	"math/big"
)

func Run() {
	privateKey, err := crypto.HexToECDSA(wallet.Sepolia.PrivateKey)
	if err != nil {
		log.Fatal(err.Error() + " " + whereami.WhereAmI())
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal(err.Error() + " " + whereami.WhereAmI())
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

	// Set the gas price and gas limit
	suggestedGasPrice, err := connection.RPC.Client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err.Error() + " " + whereami.WhereAmI())
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, config.Map[config.Active].ChainID)
	if err != nil {
		log.Fatal(err.Error() + " " + whereami.WhereAmI())
	}

	auth.Nonce = big.NewInt(int64(nonce))

	auth.Value = big.NewInt(0)
	auth.GasPrice = suggestedGasPrice

	contractByteLength := 15429
	//gasLimit calculation below
	//https://ethereum.stackexchange.com/questions/39401/how-do-you-calculate-gas-limit-for-transaction-with-data-in-ethereum
	gasLimit := 21000 + (68 * contractByteLength)
	auth.GasLimit = uint64(gasLimit)
	fmt.Println("auth.GasLimit: ", auth.GasLimit)

	blockGasLimit := getBlockGasLimit()
	if blockGasLimit < auth.GasLimit {
		auth.GasLimit = blockGasLimit
		fmt.Println("using block gas limit: ", blockGasLimit)
	}

	gasLimitBigInt := new(big.Int).SetUint64(auth.GasLimit)
	calculatedCost := new(big.Int).Mul(gasLimitBigInt, auth.GasPrice)
	authValue := new(big.Int).SetUint64(auth.Value.Uint64())
	calculatedCost = new(big.Int).Add(calculatedCost, authValue)

	fmt.Println("calculatedCost: ", calculatedCost)
	fmt.Println("pendingBalance: ", pendingBalance)
	fmt.Println("suggestedGasPrice: ", suggestedGasPrice)
	fmt.Println("gasLimit: ", gasLimit)
	fmt.Println("auth.GasLimit: ", auth.GasLimit)
	fmt.Println("auth.GasPrice: ", auth.GasPrice)
	fmt.Println("nonce: ", auth.Nonce)

	address, tx, instance, err := api.DeployBase(auth, connection.RPC.Client)
	if err != nil {
		log.Fatal(err.Error() + " " + whereami.WhereAmI())
	}

	fmt.Println("----After deploy----")
	fmt.Println("Token deployed at:", address)
	fmt.Println("Token tx:", tx.Hash())
	fmt.Println("Token cost:", tx.Cost())
	fmt.Println("Token gas limit:", tx.Gas())
	fmt.Println("Token gas price:", tx.GasPrice())
	fmt.Println("Waiting to be mined...")

	receipt, err := bind.WaitMined(context.Background(), connection.RPC.Client, tx)
	if err != nil {
		log.Fatal(err.Error() + " " + whereami.WhereAmI())
	}
	fmt.Println("Mined...")

	if receipt.Status == types.ReceiptStatusSuccessful {
		fmt.Println("Token deployed successfully")
	} else {
		reason, err := getRevertReason(receipt) //nolint:govet
		log.Fatal(err.Error() + " " + reason + " " + whereami.WhereAmI())
	}

	_ = instance
}

func getRevertReason(receipt *types.Receipt) (string, error) {
	if len(receipt.Logs) == 0 {
		return "", nil
	}

	// Get the revert event signature
	revertEventSig := []byte("Error(string)")

	// Iterate through the receipt logs to find the revert event
	for _, log := range receipt.Logs {
		if len(log.Topics) > 0 && bytes.Equal(log.Topics[0].Bytes(), revertEventSig) {
			if len(log.Data) < 32 {
				return "", fmt.Errorf("Invalid event data")
			}

			// Extract the revert reason from the event data
			revertReason := string(log.Data[32:])
			return revertReason, nil
		}
	}

	return "", nil
}
