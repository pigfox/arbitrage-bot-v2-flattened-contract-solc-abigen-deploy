package connection

import (
	"flattened-contract-solc-abigen/config"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient" //nolint:goimports
	"log"
)

var RPC Connection

type Connection struct {
	Client *ethclient.Client
}

func SetUp() {
	fmt.Println("Connection...")
	client, err := ethclient.Dial("https://" + config.Map[config.Active].NetType + ".infura.io/v3/8cfea7aaa1384f1a87b6d5aa65119ea3")
	if err != nil {
		log.Fatal(err)
	}
	RPC.Client = client
}
