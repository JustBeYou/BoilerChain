package main

import (
	"boilerchain/blocks"
	"fmt"
)

func main() {
	store := blocks.NewInMemoryStore()

	err := store.Append(blocks.NewBlock(
		2,
		blocks.Hash(store.GetTail()),
		[]byte("Test block"),
	))
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		panic("Invalid block")
	} else {
		fmt.Println("Added block")
	}

	isValid, err := blocks.ValidateChain(store)
	if isValid == false {
		fmt.Printf("Error: %v\n", err)
		panic("Invalid chain")
	} else {
		fmt.Println("Valid chain")
	}

	blocks.PrintBlockAsString(store.GetTail())
}
