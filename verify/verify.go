package verify

import (
	"arbitrage-bot-v2-flattened-contract-solc-abigen-deploy/connection"
	"arbitrage-bot-v2-flattened-contract-solc-abigen-deploy/structs"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"os"
)

func Run(contractName string) {
	fmt.Println("Contract verifying...")
	contractAddress := common.HexToAddress(structs.OnChainContract.Receipt.Address)
	addressBytecode, err := connection.RPC.Client.CodeAt(context.Background(), contractAddress, structs.OnChainContract.Receipt.Block) // nil is the latest block
	if err != nil {
		log.Fatalf("Failed to get contract addressBytecode: %v", err)
	}

	if len(addressBytecode) == 0 {
		log.Fatalf("No contract xcode found at the given address")
	}

	fmt.Printf("Compilded addressBytecode: %x\n", addressBytecode)
	fmt.Println("-----------------------------------------------------------------")

	rawCompiledBytecode := getBin(contractName)
	fmt.Printf("Compilded rawCompiledBytecode: %x\n", rawCompiledBytecode)
	fmt.Println("-----------------------------------------------------------------")
	compiledBytecodeBytes := common.FromHex(rawCompiledBytecode)
	fmt.Printf("Compilded compiledBytecodeBytes: %x\n", compiledBytecodeBytes)
	fmt.Println("-----------------------------------------------------------------")

	if string(addressBytecode) == string(compiledBytecodeBytes) {
		fmt.Println("The deployed contract addressBytecode matches the compiledBytecodeBytes.")
	} else {
		fmt.Println("The deployed contract addressBytecode does not match the compiledBytecodeBytes.")
		fmt.Println("Length deployed contract addressBytecode", len(string(addressBytecode)), "Length compiled addressBytecode", len(string(compiledBytecodeBytes)))
	}
}

func getBin(contractName string) string {
	path, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error:%s", err)
	}
	filePath := path + "/solc-output/" + fmt.Sprintf("%s.bin", contractName)

	// Read the file contents
	binData, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to read bin file: %v", err)
	}

	// Return the contents as a string
	return string(binData)
}
