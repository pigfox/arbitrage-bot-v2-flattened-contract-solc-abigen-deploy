package compile

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func Run(contract, compilerVersion, evm string) {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get current working directory: %v", err)
	}

	script := fmt.Sprintf("%s/compile/compile.sh", cwd)
	cmd := exec.Command(script, contract, compilerVersion, evm)

	// Connect cmd.Stdout and cmd.Stderr to os.Stdout and os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		log.Fatalf("Failed to execute ./compile.sh: %v", err)
	}
}
