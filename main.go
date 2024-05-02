package main

import (
	"flattened-contract-solc-abigen/code"
	"flattened-contract-solc-abigen/config"
	"flattened-contract-solc-abigen/connection"
	"flattened-contract-solc-abigen/wallet"
	"fmt"
)

func setUp() {
	fmt.Println("Setting up...")
	wallet.SetUp()
	config.SetUp(config.Sepolia)
	connection.SetUp()
}

func main() {
	setUp()
	//compile.SolC()
	//deploy.Run()
	code.Run()
}
