package blocks

import (
	"fmt"
	"math/rand"
	"time"
)

// Block - building unit of the blockchain
type Block struct {
	index     int64
	nonce     int64
	timestamp int64
	prevHash  string
	data      []byte
	hash      string
}

// NewBlock -
func NewBlock(index int64, prevHash string, data []byte) Block {
	nonce := rand.Int63()
	currentTimestamp := time.Now().UnixNano()

	newBlock := Block{
		index,
		nonce,
		currentTimestamp,
		prevHash,
		data,
		"",
	}
	newBlock.hash = Hash(newBlock)

	return newBlock
}

// PrintBlock -
func PrintBlock(b Block, unserialize func(Block) string) {
	fmt.Printf("--- Block %d ---\n", b.index)
	fmt.Printf("Nonce: %d\n", b.nonce)
	fmt.Printf("Timestamp: %d\n", b.timestamp)
	fmt.Printf("Prev. block hash: %s\n", b.prevHash)
	fmt.Printf("Data: %s\n", unserialize(b))
	fmt.Printf("Hash: %s\n", b.hash)
	fmt.Printf("--- --- ---\n")
}

// PrintBlockAsHex -
func PrintBlockAsHex(b Block) {
	PrintBlock(b, func(b Block) string {
		return fmt.Sprintf("%x", b.data)
	})
}

// PrintBlockAsString -
func PrintBlockAsString(b Block) {
	PrintBlock(b, func(b Block) string {
		return fmt.Sprintf("%s", b.data)
	})
}
