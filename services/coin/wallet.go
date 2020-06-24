package coin

import (
	"encoding/gob"
	"fmt"

	"github.com/rs/xid"
)

// Wallet -
type Wallet struct {
	ID          string
	Description string
	Amount      uint64
	Keys        KeyPair
	Signature   []byte
}

func (t Wallet) RemoveSignature() interface{} {
	copyT := t
	copyT.Signature = []byte{}
	return copyT
}

func init() {
	gob.Register(Wallet{})
}

// MarshalText -
func (w Wallet) MarshalText() (string, error) {
	text := fmt.Sprintf(`--- Wallet %s ---
Description: %s
Amount: %d
Signature: %x
--- --- ---`, w.ID, w.Description, w.Amount, w.Signature)
	return text, nil
}

// NewWallet -
func NewWallet(description string, passphrase string) Wallet {
	keys, _ := GenerateKeyPair(passphrase)
	newWallet := Wallet{
		xid.New().String(),
		description,
		0,
		keys,
		[]byte{},
	}

	privateKey, _ := DecryptPrivateKey(newWallet.Keys.PrivateKeyEncrypted, passphrase)
	signature, _ := Sign(privateKey, newWallet)
	newWallet.Signature = signature

	return newWallet
}

// PrintWallet -
func PrintWallet(w Wallet) {
	toPrint, _ := w.MarshalText()
	fmt.Print(toPrint)
}
