package main

import (
	"boilerchain/blocks"
	"fmt"
)

func main() {
	store := blocks.NewInMemoryStore()

	for i := 2; i < 10; i++ {
		err := store.Append(blocks.NewBlock(
			int64(i),
			blocks.Hash(store.GetTail()),
			blocks.ByteContent{
				Data: []byte(fmt.Sprintf("Test block %d", i)),
			},
		))
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			panic("Invalid block")
		} else {
			fmt.Printf("Added block %d\n", i)
		}
	}

	isValid, err := blocks.ValidateChain(store)
	if isValid == false {
		fmt.Printf("Error: %v\n", err)
		panic("Invalid chain")
	} else {
		fmt.Println("Valid chain")
	}

	blocks.PrintChain(store)
}
