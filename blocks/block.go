package blocks

import (
	"encoding"
	"fmt"
	"time"
)

// Block - building unit of the blockchain
type Block struct {
	index     int64
	nonce     int64
	timestamp int64
	prevHash  string
	data      interface{}
	hash      string
}

// ByteContent -
type ByteContent struct {
	Data []byte
}

// MarshalBinary -
func (bc ByteContent) MarshalBinary() ([]byte, error) {
	return bc.Data, nil
}

// MarshalText -
func (bc ByteContent) MarshalText() (string, error) {
	return string(bc.Data), nil
}

// NewBlock -
func NewBlock(index int64, prevHash string, data interface{}) Block {
	currentTimestamp := time.Now().UnixNano()

	newBlock := Block{
		index,
		0,
		currentTimestamp,
		prevHash,
		data,
		"",
	}

	nonce := Work(newBlock)

	newBlock.nonce = nonce
	newBlock.hash = Hash(newBlock)

	return newBlock
}

// PrintBlock -
func PrintBlock(b Block) {
	fmt.Printf("--- Block %d ---\n", b.index)
	fmt.Printf("Nonce: %d\n", b.nonce)
	fmt.Printf("Timestamp: %d\n", b.timestamp)
	fmt.Printf("Prev. block hash: %s\n", b.prevHash)
	binaryData, _ := b.data.(encoding.TextMarshaler).MarshalText()
	fmt.Printf("Data: %s\n", binaryData)
	fmt.Printf("Hash: %s\n", b.hash)
	fmt.Printf("--- --- ---\n")
}
