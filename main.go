package main

import (
	"arbitrage-bot-v2-flattened-contract-solc-abigen-deploy/compile"
	"arbitrage-bot-v2-flattened-contract-solc-abigen-deploy/config"
	"arbitrage-bot-v2-flattened-contract-solc-abigen-deploy/connection"
	"arbitrage-bot-v2-flattened-contract-solc-abigen-deploy/wallet"
	"fmt"
)

var contractName string = "Base" //file name w/o _flat or .sol
func setUp() {
	fmt.Println("Setting up...")
	wallet.SetUp()
	config.SetUp(config.Sepolia)
	connection.SetUp()
}

func main() {
	setUp()
	compile.Run("Base", "0.8.24")
	//deploy.Run(contractName)
	//verify.Run(contractName)
	/*
		if contractName == "Base" {
			cbase.Run()
		}

	*/
}
