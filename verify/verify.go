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
	deployedBytecode, err := connection.RPC.Client.CodeAt(context.Background(), contractAddress, structs.OnChainContract.Receipt.Block) // nil is the latest block
	if err != nil {
		log.Fatalf("Failed to get contract deployedBytecode: %v", err)
	}

	if len(deployedBytecode) == 0 {
		log.Fatalf("No contract xcode found at the given address")
	}

	fmt.Printf("Compiled deployedBytecode: %x\n", deployedBytecode)
	fmt.Println("-----------------------------------------------------------------")

	rawCompiledBytecode := getBin(contractName)
	fmt.Printf("Compiled rawCompiledBytecode: %x\n", rawCompiledBytecode)
	fmt.Println("-----------------------------------------------------------------")
	compiledBytecodeBytes := common.FromHex(rawCompiledBytecode)
	fmt.Printf("Compiled compiledBytecodeBytes: %x\n", compiledBytecodeBytes)
	fmt.Println("-----------------------------------------------------------------")

	if string(deployedBytecode) == string(compiledBytecodeBytes) {
		fmt.Println("The deployedBytecode matches the compiledBytecodeBytes.")
	} else if string(deployedBytecode) == rawCompiledBytecode {
		fmt.Println("The deployedBytecode matches the rawCompiledBytecode.")
	} else {
		fmt.Println("The deployedBytecode does not match the compiledBytecodeBytes or the rawCompiledBytecode.")
		fmt.Println("Length compiled deployedBytecode", len(string(deployedBytecode)))
		fmt.Println("Length compiled compiledBytecodeBytes", len(string(compiledBytecodeBytes)))
		fmt.Println("Length compiled  rawCompiledBytecode", len(rawCompiledBytecode))
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
