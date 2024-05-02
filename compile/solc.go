package compile

import (
	"flattened-contract-solc-abigen/util"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func SolC() {
	currentPath := util.GetCurrentPath()
	contractFileName := "Base"
	contractFile := currentPath + "/sol/" + contractFileName + ".sol" // Path to the Solidity contract
	outputDir := currentPath + "/solc-output"                         // Directory to store ABI and BIN files
	apiDir := currentPath + "/api"                                    // Directory to store generated Go bindings
	generatedFileName := "Base"
	generatedPackageName := "api"

	err := util.EmptyFolder(outputDir)
	if err != nil {
		log.Fatalf(" %v\n: %s can't empty folder", err, outputDir)
	}
	err = util.EmptyFolder(apiDir)
	if err != nil {
		log.Fatalf(" %v\n: %s can't empty folder", err, apiDir)
	}

	// Ensure output and API directories exist
	err = os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		log.Fatalf(" %v\n: %s not writable", err, outputDir)
	}
	err = os.MkdirAll(apiDir, os.ModePerm)
	if err != nil {
		log.Fatalf(" %v\n: %s not writable", err, outputDir)
	}

	// Compile the contract using solc
	cmd := exec.Command("solc", "--optimize", "--optimize-runs 200", "--abi", "--bin", "--overwrite", "--output-dir", outputDir, contractFile)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Failed to compile contract: %v\nOutput: %s", err, string(output))
	}
	fmt.Println("Contract compiled successfully")
	//mergedFiles(outputDir, apiDir, generatedFileName, generatedPackageName)
	baseFile(outputDir, apiDir, generatedFileName, generatedPackageName)
}
