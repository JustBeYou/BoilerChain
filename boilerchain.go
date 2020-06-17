package main

import (
	"boilerchain/blocks"
	"boilerchain/services/coin"
	"fmt"
)

func main() {
	coin := coin.NewCoin("Buci", 1000000, "qwerty123456")
	isValid, err := blocks.ValidateChain(coin.Store)
	if isValid == false {
		fmt.Printf("Invalid chain: %s\n", err)
	}
	blocks.PrintChain(coin.Store)
}
