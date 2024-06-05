package main

import (
	"arbitrage-bot-v2-flattened-contract-solc-abigen-deploy/compile"
	"arbitrage-bot-v2-flattened-contract-solc-abigen-deploy/config"
	"arbitrage-bot-v2-flattened-contract-solc-abigen-deploy/connection"
	"arbitrage-bot-v2-flattened-contract-solc-abigen-deploy/deploy"
	"arbitrage-bot-v2-flattened-contract-solc-abigen-deploy/structs"
	"arbitrage-bot-v2-flattened-contract-solc-abigen-deploy/verify"
	"arbitrage-bot-v2-flattened-contract-solc-abigen-deploy/wallet"
	"fmt"
)

func setUp() {
	fmt.Println("Setting up...")
	wallet.SetUp()
	config.SetUp(config.Sepolia)
	connection.SetUp()
}

func main() {
	contractName := "Trade" //file name w/o _flat or .sol
	setUp()
	compile.Run(contractName, "0.8.20", "london")
	deploy.Run(contractName)
	verify.Run(contractName)
	//cbase.Run()

	fmt.Println(fmt.Sprintf("+%v", structs.OnChainContract))
}
