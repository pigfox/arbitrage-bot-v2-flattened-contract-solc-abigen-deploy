package compile

import (
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
)

func baseFile(outputDir, apiDir, generatedFileName, generatedPackageName string) {
	// Generate Go bindings using abigen
	outFile := filepath.Join(apiDir, generatedFileName+".go")
	compiledFilename := outputDir + "/" + generatedFileName
	cmd := exec.Command("abigen", "--bin="+compiledFilename+".bin", "--abi="+compiledFilename+".abi", "--pkg="+generatedPackageName, "--type="+generatedFileName, "--out="+outFile) //nolint:lll,gosec
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Failed to generate Go bindings: %v\nOutput: %s", err, string(output))
	}
	fmt.Println("Go bindings generated successfully")
}
