package code

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
	"math/big"
)

func processLog(contractAbi *abi.ABI, vLog types.Log) {
	eventName := ""
	eventFound := false

	for _, event := range contractAbi.Events {
		if vLog.Topics[0] == event.ID { // Correct usage of event ID
			eventName = event.Name
			eventFound = true
			break
		}
	}

	if eventFound {
		fmt.Printf("Event [%s] detected\n", eventName)
		event := map[string]interface{}{}
		err := contractAbi.UnpackIntoMap(event, eventName, vLog.Data)
		if err != nil {
			log.Printf("Error unpacking event %s: %v\n", eventName, err)
		} else {
			for key, value := range event {
				switch val := value.(type) {
				case *big.Int:
					fmt.Printf("%s: %d\n", key, val)
				default:
					fmt.Printf("%s: %v\n", key, val)
				}
			}
		}
	} else {
		fmt.Println("Unknown event type")
	}
}
