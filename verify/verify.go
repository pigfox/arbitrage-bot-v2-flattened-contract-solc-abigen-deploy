package verify

import (
	"arbitrage-bot-v2-flattened-contract-solc-abigen-deploy/connection"
	"arbitrage-bot-v2-flattened-contract-solc-abigen-deploy/structs"
	"arbitrage-bot-v2-flattened-contract-solc-abigen-deploy/util"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"log"
)

func Run(params structs.DeploymentParam) {
	fmt.Println("Contract verifying...")
	contractAddress := common.HexToAddress(structs.OnChainContract.Receipt.Address)
	deployedBytecode, err := connection.RPC.Client.CodeAt(context.Background(), contractAddress, structs.OnChainContract.Receipt.Block) // nil is the latest block
	if err != nil {
		log.Fatalf("Failed to get contract deployedBytecode: %v", err)
	}

	if len(deployedBytecode) == 0 {
		log.Fatalf("No contract code found at the given address")
	}

	fmt.Printf("Compiled deployedBytecode: %x\n", deployedBytecode)
	fmt.Println("-----------------------------------------------------------------")

	rawCompiledBytecode := util.GetContractBIN(params.Name)
	compiledBytecodeBytes := common.FromHex(rawCompiledBytecode)
	fmt.Printf("Compiled compiledBytecodeBytes: %x\n", compiledBytecodeBytes)
	fmt.Println("-----------------------------------------------------------------")

	if string(deployedBytecode) == string(compiledBytecodeBytes) {
		fmt.Println("The deployedBytecode matches the compiledBytecodeBytes.")
	} else if string(deployedBytecode) == rawCompiledBytecode {
		fmt.Println("The deployedBytecode matches the rawCompiledBytecode.")
	} else {
		fmt.Println("The deployedBytecode does not match the compiledBytecodeBytes.")
		fmt.Println("Length compiled deployedBytecode", len(string(deployedBytecode)))
		fmt.Println("Length compiled compiledBytecodeBytes", len(string(compiledBytecodeBytes)))
	}
}
