package blocks

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"errors"
	"fmt"
	"strings"
)

// ValidateChain -
func ValidateChain(store BlockStore) (bool, error) {
	block := store.GetHead()
	for block.Index != store.GetNull().Index {
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
	if b.Hash[len(b.Hash)-int(b.Difficulty):] != strings.Repeat("0", int(b.Difficulty)) {
		return false, errors.New("Invalid proof of work")
	}
	return true, nil
}

// Work -
func Work(b Block) int64 {
	b.Nonce = 0
	b.Hash = Hash(b)
	isValid, _ := ValidatePoW(b)
	for isValid == false {
		b.Nonce++
		b.Hash = Hash(b)
		isValid, _ = ValidatePoW(b)
	}
	return b.Nonce
}

// Validate -
func Validate(child Block, parent Block) (bool, error) {
	isChildPoWValid, err := ValidatePoW(child)
	if isChildPoWValid == false {
		return false, err
	}

	if child.Index-1 != parent.Index {
		return false, errors.New("Invalid Index")
	}
	if child.Difficulty != parent.Difficulty {
		return false, errors.New("Invalid difficulty")
	}
	if child.Timestamp < parent.Timestamp {
		return false, errors.New("Invalid Timestamp")
	}
	if child.PrevHash != parent.Hash {
		return false, errors.New("Invalid PrevHash")
	}
	if child.Hash != Hash(child) {
		return false, errors.New("Invalid Hash")
	}
	if parent.Hash != Hash(parent) {
		return false, errors.New("Invalid parent Hash")
	}
	return true, nil
}

// Hash -
func Hash(b Block) string {
	blockCopy := b
	blockCopy.Hash = ""

	serialized := bytes.Buffer{}
	err := gob.NewEncoder(&serialized).Encode(blockCopy)
	if err != nil {
		panic(err)
	}

	hasher := sha256.New()
	hasher.Write(serialized.Bytes())
	return fmt.Sprintf("%x", hasher.Sum(nil))
}
