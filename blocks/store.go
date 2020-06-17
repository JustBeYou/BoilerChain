package blocks

import (
	"errors"
)

// BlockStore -
// Genesis block is mandatory
type BlockStore interface {
	GetByIndex(int64) (Block, error)
	GetByHash(string) (Block, error)
	GetHead() Block
	GetTail() Block
	GetNull() Block
	GetNext(Block) (Block, error)
	GetPrev(Block) (Block, error)
	Add(interface{}) Block
	Append(Block) error
}

// InMemoryStore -
type InMemoryStore struct {
	blocks []Block
}

// NewInMemoryStore -
func NewInMemoryStore(genesisData interface{}) *InMemoryStore {
	newStore := new(InMemoryStore)
	newStore.Add(genesisData)
	return newStore
}

// GetByIndex -
func (store *InMemoryStore) GetByIndex(index int64) (Block, error) {
	for _, block := range store.blocks {
		if block.index == index {
			return block, nil
		}
	}
	return store.GetNull(), errors.New("Block not found")
}

// GetByHash -
func (store *InMemoryStore) GetByHash(hash string) (Block, error) {
	for _, block := range store.blocks {
		if block.hash == hash {
			return block, nil
		}
	}
	return store.GetNull(), errors.New("Block not found")
}

// GetHead -
func (store *InMemoryStore) GetHead() Block {
	if len(store.blocks) == 0 {
		return store.GetNull()
	}
	return store.blocks[0]
}

// GetTail -
func (store *InMemoryStore) GetTail() Block {
	if len(store.blocks) == 0 {
		return store.GetNull()
	}
	return store.blocks[len(store.blocks)-1]
}

// GetNull -
func (store *InMemoryStore) GetNull() Block {
	nullBlock := Block{}
	nullBlock.data = ByteContent{[]byte("Null block")}
	nullBlock.hash = Hash(nullBlock)
	return nullBlock
}

// GetNext -
func (store *InMemoryStore) GetNext(b Block) (Block, error) {
	block, err := store.GetByIndex(b.index + 1)
	if err != nil {
		return store.GetNull(), errors.New("Next block does not exist")
	}
	return block, nil
}

// GetPrev -
func (store *InMemoryStore) GetPrev(b Block) (Block, error) {
	block, err := store.GetByIndex(b.index - 1)
	if err != nil {
		return store.GetNull(), errors.New("Previous block does not exist")
	}
	return block, nil
}

// Add -
func (store *InMemoryStore) Add(data interface{}) Block {
	tail := store.GetTail()
	newBlock := NewBlock(
		tail.index+1,
		tail.hash,
		data,
	)
	store.blocks = append(store.blocks, newBlock)
	return newBlock
}

// Append -
func (store *InMemoryStore) Append(b Block) error {
	tail := store.GetTail()
	isValid, err := Validate(b, tail)
	if isValid == false {
		return err
	}
	store.blocks = append(store.blocks, b)
	return nil
}

// PrintChain -
func PrintChain(store BlockStore) {
	block := store.GetHead()
	for block.index != store.GetNull().index {
		PrintBlock(block)
		block, _ = store.GetNext(block)
	}
}
