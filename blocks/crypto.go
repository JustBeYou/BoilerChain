package blocks

import (
	"crypto/sha256"
	"encoding/binary"
	"errors"
	"fmt"
)

// ValidateChain -
func ValidateChain(store BlockStore) (bool, error) {
	block := store.GetHead()
	for block.index != store.GetNull().index {
		prevBlock, _ := store.GetPrev(block)
		isValid, _ := Validate(block, prevBlock)
		if isValid == false {
			return false, errors.New("Invalid chain")
		}
		block, _ = store.GetNext(block)
	}
	return true, nil
}

// Validate -
func Validate(child Block, parent Block) (bool, error) {
	if child.index-1 != parent.index {
		return false, errors.New("Invalid index")
	}
	if child.timestamp < parent.timestamp {
		return false, errors.New("Invalid timestamp")
	}
	if child.prevHash != parent.hash {
		return false, errors.New("Invalid prevHash")
	}
	if child.hash != Hash(child) {
		return false, errors.New("Invalid hash")
	}
	if parent.hash != Hash(parent) {
		return false, errors.New("Invalid parent hash")
	}
	return true, nil
}

// Hash -
func Hash(b Block) string {
	indexAsBytes := toByteArray(b.index)
	nonceAsBytes := toByteArray(b.nonce)
	timestampAsBytes := toByteArray(b.timestamp)

	hasher := sha256.New()
	hasher.Write(indexAsBytes)
	hasher.Write(nonceAsBytes)
	hasher.Write(timestampAsBytes)
	hasher.Write([]byte(b.prevHash))
	hasher.Write([]byte(b.data))
	return fmt.Sprintf("%x", hasher.Sum(nil))
}

func toByteArray(i int64) []byte {
	iAsBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(iAsBytes, uint64(i))
	return iAsBytes
}
