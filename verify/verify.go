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
	addressBytecode, err := connection.RPC.Client.CodeAt(context.Background(), contractAddress, nil) // nil is the latest block
	if err != nil {
		log.Fatalf("Failed to get contract addressBytecode: %v", err)
	}

	if len(addressBytecode) == 0 {
		log.Fatalf("No contract xcode found at the given address")
	}

	fmt.Printf("Compilded addressBytecode: %x\n", addressBytecode)
	fmt.Println("-----------------------------------------------------------------")

	//rawCompiledBytecode := "608060405234801561001057600080fd5b506040516020806101c78339810180604052810190808051906020019092919080519060200190929190505050806000819055505061017e806100616000396000f3fe6080604052600436106100565760003560e01c80631561ba021461005b5780633ccfd60b146100c7575b600080fd5b34801561006757600080fd5b50610070610108565b6040805160ff9092168252519081900360200190f35b6100cf61011c565b6040805192835260208301919091528051918290030190f35b60008054905090565b600080546001019055565b600080546100f09061011c565b80601f01602080910402602001604051908101604052809291908181526020018280546101289061011c565b80156101755780601f1061014a57610100808354040283529160200191610175565b820191906000526020600020905b81548152906001019060200180831161015857829003601f168201915b505050505090509056fea2646970667358221220f5abfd61846be815d1efb2e41e1d14b7be6b8cbb81852af5f88a280b1a71ea1e64736f6c63430006060033"
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
