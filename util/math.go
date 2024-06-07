package util

import "math/big"

func AddPercentBigInt(value *big.Int, percent int) *big.Int {
	// Convert percent to *big.Int
	percentBigInt := big.NewInt(int64(percent))

	// Calculate the percentage of the value
	percentageValue := new(big.Int).Mul(value, percentBigInt)
	percentageValue.Div(percentageValue, big.NewInt(100))

	// Add the percentage to the original value

	return new(big.Int).Add(value, percentageValue)
}

func AddPercentUint64(value uint64, percent int) uint64 {
	// Convert value to *big.Int
	valueBigInt := big.NewInt(int64(value))

	// Calculate the percentage of the value
	percentageValue := AddPercentBigInt(valueBigInt, percent)

	return percentageValue.Uint64()
}

func WeiToETH(wei *big.Int) *big.Float {
	weiFloat := new(big.Float).SetInt(wei)
	ethValue := new(big.Float).Quo(weiFloat, big.NewFloat(1000000000000000000))

	return ethValue
}

func WeiToGwei(wei *big.Int) *big.Float {
	weiFloat := new(big.Float).SetInt(wei)
	gweiValue := new(big.Float).Quo(weiFloat, big.NewFloat(1000000000))

	return gweiValue
}
