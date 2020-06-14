package main

import (
	"errors"
)

// BlockStore -
// Genesis block is mandatory
type BlockStore interface {
	GetByIndex(int64) (Block, error)
	GetByHash(string) (Block, error)
	GetTail() Block
	NewBlock([]byte) Block
}

// InMemoryBlockStore -
type InMemoryBlockStore struct {
	blocks []Block
}

// NewInMemoryBlockStore -
func NewInMemoryBlockStore() InMemoryBlockStore {
	newStore := InMemoryBlockStore{}
	newStore.NewBlock([]byte("Genesis block"))
	return newStore
}

// GetByIndex -
func (store InMemoryBlockStore) GetByIndex(index int64) (Block, error) {
	for _, block := range store.blocks {
		if block.index == index {
			return block, nil
		}
	}
	return Block{}, errors.New("Block not found")
}

// GetByHash -
func (store InMemoryBlockStore) GetByHash(hash string) (Block, error) {
	for _, block := range store.blocks {
		if block.hash == hash {
			return block, nil
		}
	}
	return Block{}, errors.New("Block not found")
}

// GetTail -
func (store InMemoryBlockStore) GetTail() Block {
	if len(store.blocks) == 0 {
		return Block{}
	}
	return store.blocks[len(store.blocks)-1]
}

// NewBlock -
func (store *InMemoryBlockStore) NewBlock(data []byte) Block {
	tail := store.GetTail()
	newBlock := NewBlock(
		tail.index+1,
		tail.hash,
		data,
	)
	store.blocks = append(store.blocks, newBlock)
	return newBlock
}
