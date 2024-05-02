package wallet

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var Sepolia Wallet
var Ethereum Wallet

func SetUp() {
	fmt.Println("Wallet...")
	err := godotenv.Load() // Load() will read your .env file by default
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	Sepolia = newSepolia()
	Ethereum = newEthereum()
}

type Wallet struct {
	Address    string
	PrivateKey string
}

func newSepolia() Wallet {
	return Wallet{
		Address:    os.Getenv("SEPOLIA_ADDRESS"),
		PrivateKey: os.Getenv("SEPOLIA_PRIVATE_KEY"),
	}
}

func newEthereum() Wallet {
	return Wallet{
		Address:    os.Getenv("ETHEREUM_ADDRESS"),
		PrivateKey: os.Getenv("ETHEREUM_PRIVATE_KEY"),
	}
}
