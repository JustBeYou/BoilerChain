package blocks

import (
	"encoding"
	"encoding/gob"
	"fmt"
	"time"
)

// Block - building unit of the blockchain
type Block struct {
	Index     int64
	Nonce     int64
	Timestamp int64
	PrevHash  string
	Data      interface{}
	Hash      string
}

// ByteContent -
type ByteContent struct {
	Data []byte
}

func init() {
	gob.Register(Block{})
	gob.Register(ByteContent{})
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

	newBlock.Nonce = nonce
	newBlock.Hash = Hash(newBlock)

	return newBlock
}

// PrintBlock -
func PrintBlock(b Block) {
	fmt.Printf("--- Block %d ---\n", b.Index)
	fmt.Printf("Nonce: %d\n", b.Nonce)
	fmt.Printf("Timestamp: %d\n", b.Timestamp)
	fmt.Printf("Prev. block Hash: %s\n", b.PrevHash)
	binaryData, _ := b.Data.(encoding.TextMarshaler).MarshalText()
	fmt.Printf("Data: %s\n", binaryData)
	fmt.Printf("Hash: %s\n", b.Hash)
	fmt.Printf("--- --- ---\n")
}
