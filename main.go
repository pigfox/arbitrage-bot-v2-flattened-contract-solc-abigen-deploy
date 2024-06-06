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
	structs.SetUp()
	wallet.SetUp()
	config.SetUp(config.Sepolia)
	connection.SetUp()
}

func main() {
	setUp()
	contractName := "Base" //file name w/o _flat or .sol
	params := structs.DeploymentParams[contractName]
	compile.Run(params.Name, params.CompileVersion, params.EVM)
	deploy.Run(params.Name)
	verify.Run(params.Name)
	//cbase.Run()

	fmt.Println(fmt.Sprintf("+%v", structs.OnChainContract))
}
