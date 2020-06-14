package main

import (
	"crypto/sha256"
	"encoding/binary"
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

	indexAsBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(indexAsBytes, uint64(index))

	nonceAsBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(nonceAsBytes, uint64(nonce))

	currentTimestampAsBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(currentTimestampAsBytes, uint64(currentTimestamp))

	hasher := sha256.New()
	hasher.Write(indexAsBytes)
	hasher.Write(nonceAsBytes)
	hasher.Write(currentTimestampAsBytes)
	hasher.Write([]byte(prevHash))
	hasher.Write([]byte(data))
	newHash := fmt.Sprintf("%x", hasher.Sum(nil))

	return Block{
		index,
		nonce,
		currentTimestamp,
		prevHash,
		data,
		newHash,
	}
}

// PrintBlock -
func PrintBlock(b Block) {
	fmt.Printf("--- Block %d ---\n", b.index)
	fmt.Printf("Nonce: %d\n", b.nonce)
	fmt.Printf("Timestamp: %d\n", b.timestamp)
	fmt.Printf("Prev. block hash: %s\n", b.prevHash)
	fmt.Printf("Data: %s\n", string(b.data))
	fmt.Printf("Hash: %s\n", b.hash)
	fmt.Printf("--- --- ---\n")
}
