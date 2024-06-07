package structs

import (
	"arbitrage-bot-v2-flattened-contract-solc-abigen-deploy/constants"
	"fmt"
)

func SetUp() {
	fmt.Println("Deployment params...")
	DeploymentParams = make(map[string]DeploymentParam)
	DeploymentParams["Trade"] = DeploymentParam{
		Name:                 "Trade",
		EVM:                  "london",
		CompilerVersion:      "0.8.20",
		ConstructorArguments: make([]interface{}, 0),
	}
	DeploymentParams["Arbitrage"] = DeploymentParam{
		Name:                 "Arbitrage",
		EVM:                  "london",
		CompilerVersion:      "0.8.20",
		ConstructorArguments: make([]interface{}, 0),
	}
	DeploymentParams["SUNToken"] = DeploymentParam{
		Name:                 "SUNToken",
		EVM:                  "cancun",
		CompilerVersion:      "0.8.24",
		ConstructorArguments: make([]interface{}, 0),
	}
	DeploymentParams["Pigfox"] = DeploymentParam{
		Name:                 "Pigfox",
		EVM:                  "london",
		CompilerVersion:      "0.8.20",
		ConstructorArguments: []interface{}{constants.KrakenVaultAddress},
	}
	DeploymentParams["Dex"] = DeploymentParam{
		Name:                 "Dex",
		EVM:                  "london",
		CompilerVersion:      "0.8.20",
		ConstructorArguments: make([]interface{}, 0),
	}
	DeploymentParams["Base"] = DeploymentParam{
		Name:                 "Base",
		EVM:                  "cancun",
		CompilerVersion:      "0.8.24",
		ConstructorArguments: make([]interface{}, 0),
	}
	DeploymentParams["Pigfox2"] = DeploymentParam{
		Name:                 "Pigfox2",
		EVM:                  "london",
		CompilerVersion:      "0.8.20",
		ConstructorArguments: make([]interface{}, 0),
	}
	DeploymentParams["Nothing"] = DeploymentParam{
		Name:                 "Nothing",
		EVM:                  "london",
		CompilerVersion:      "0.8.20",
		ConstructorArguments: make([]interface{}, 0),
	}
}
