package blocks

import (
	"crypto/sha256"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"
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

// ValidatePoW - validate proof of work
func ValidatePoW(b Block) (bool, error) {
	const numberOfZeros = 3
	if b.hash[len(b.hash)-numberOfZeros:] != strings.Repeat("0", numberOfZeros) {
		return false, errors.New("Invalid proof of work")
	}
	return true, nil
}

// Work -
func Work(b Block) int64 {
	b.nonce = 0
	b.hash = Hash(b)
	isValid, _ := ValidatePoW(b)
	for isValid == false {
		b.nonce++
		b.hash = Hash(b)
		isValid, _ = ValidatePoW(b)
	}
	return b.nonce
}

// Validate -
func Validate(child Block, parent Block) (bool, error) {
	isChildPoWValid, err := ValidatePoW(child)
	if isChildPoWValid == false {
		return false, err
	}

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
	hasher.Write(b.data.(Serializer).Serialize())
	return fmt.Sprintf("%x", hasher.Sum(nil))
}

func toByteArray(i int64) []byte {
	iAsBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(iAsBytes, uint64(i))
	return iAsBytes
}
