package structs

import "fmt"

func SetUp() {
	fmt.Println("Deployment params...")
	DeploymentParams = make(map[string]DeploymentParam)
	DeploymentParams["Trade"] = DeploymentParam{
		Name:           "Trade",
		EVM:            "london",
		CompileVersion: "0.8.20",
	}
	DeploymentParams["Arbitrage"] = DeploymentParam{
		Name:           "Arbitrage",
		EVM:            "london",
		CompileVersion: "0.8.20",
	}
	DeploymentParams["SUNToken"] = DeploymentParam{
		Name:           "SUNToken",
		EVM:            "cancun",
		CompileVersion: "0.8.24",
	}
	DeploymentParams["Pigfox"] = DeploymentParam{
		Name:           "Pigfox",
		EVM:            "london",
		CompileVersion: "0.8.20",
	}
	DeploymentParams["Dex"] = DeploymentParam{
		Name:           "Dex",
		EVM:            "london",
		CompileVersion: "0.8.20",
	}
	DeploymentParams["Base"] = DeploymentParam{
		Name:           "Base",
		EVM:            "cancun",
		CompileVersion: "0.8.24",
	}
}
