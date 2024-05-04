package code

import "fmt"

var abiMap = map[string]string{}

func setABIs() {
	fmt.Println("ABI...")
	abiMap = make(map[string]string)
	abiMap["0x938c5cC616EDA96642D2E1AEa7680aDEAFA93B26"] = `[{"inputs":[],"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"uint256","name":"x","type":"uint256"},{"indexed":false,"internalType":"uint256","name":"y","type":"uint256"}],"name":"Adding","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"address","name":"owner","type":"address"},{"indexed":false,"internalType":"address","name":"delployedAt","type":"address"}],"name":"Created","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"uint256","name":"q","type":"uint256"}],"name":"SettingQ","type":"event"},{"inputs":[{"internalType":"uint256","name":"_x","type":"uint256"},{"internalType":"uint256","name":"_y","type":"uint256"}],"name":"add","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"string","name":"s","type":"string"}],"name":"concat","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"getQ","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint256","name":"_q","type":"uint256"}],"name":"setQ","outputs":[],"stateMutability":"nonpayable","type":"function"}]`
}
