package coin

import (
	"bytes"
	"encoding"
	"encoding/gob"
	"fmt"

	"github.com/rs/xid"
)

// TransactionType -
type TransactionType uint16

const (
	// Null -
	Null TransactionType = 0
	// Mine -
	Mine = 1
	// WalletAnnouncement -
	WalletAnnouncement = 2
	// Payment -
	Payment = 3
)

// Transaction -
type Transaction struct {
	ID          string
	PublisherID string
	Description string

	Type      TransactionType
	Data      interface{}
	Signature []byte
}

func (t Transaction) RemoveSignature() interface{} {
	copyT := t
	copyT.Signature = []byte{}
	return copyT
}

// NewTransaction -
func NewTransaction(
	publisher Wallet,
	passphrase string,
	description string,
	typeName TransactionType,
	data interface{},
) Transaction {
	t := Transaction{
		xid.New().String(),
		publisher.ID,
		description,
		typeName,
		data,
		[]byte{},
	}

	privateKey, _ := DecryptPrivateKey(
		publisher.Keys.PrivateKeyEncrypted,
		passphrase,
	)
	signature, _ := Sign(privateKey, t)
	t.Signature = signature

	return t
}

// TransactionChain -
type TransactionChain struct {
	Transactions []Transaction
}

// MineTransaction -
type MineTransaction struct {
	BeneficiaryID string
	Amount        uint64
}

// WalletAnnouncementTransaction -
type WalletAnnouncementTransaction struct {
	Created Wallet
}

// PaymentTransaction -
type PaymentTransaction struct {
	RecipientID string
	SenderID    string
	Amount      uint64
}

func init() {
	gob.Register(Transaction{})
	gob.Register(TransactionChain{})
	gob.Register(MineTransaction{})
	gob.Register(WalletAnnouncementTransaction{})
	gob.Register(PaymentTransaction{})
}

// MarshalText -
func (tc TransactionChain) MarshalText() ([]byte, error) {
	buf := bytes.Buffer{}
	buf.WriteString("\n")
	for _, t := range tc.Transactions {
		tb, _ := t.MarshalText()
		buf.Write(tb)
	}
	buf.WriteString("--- --- ---\n")
	return buf.Bytes(), nil
}

// MarshalText -
func (t Transaction) MarshalText() ([]byte, error) {
	buf := bytes.Buffer{}
	buf.WriteString(fmt.Sprintf("--- Transaction %s ---\n", t.ID))
	buf.WriteString(fmt.Sprintf("Description: %s\n", t.Description))
	buf.WriteString(fmt.Sprintf("Publisher: %s\n", t.PublisherID))

	text, _ := t.Data.(encoding.TextMarshaler).MarshalText()
	buf.Write(text)
	buf.WriteString("\n")
	return buf.Bytes(), nil
}

// MarshalText -
func (t MineTransaction) MarshalText() ([]byte, error) {
	text := fmt.Sprintf("Beneficiary ID: %s\nAmount: %d\n", t.BeneficiaryID, t.Amount)
	return []byte(text), nil
}

// MarshalText -
func (t WalletAnnouncementTransaction) MarshalText() ([]byte, error) {
	walletText, _ := t.Created.MarshalText()
	text := fmt.Sprintf("Wallet: %s\n", walletText)
	return []byte(text), nil
}

// MarshalText -
func (t PaymentTransaction) MarshalText() ([]byte, error) {
	text := fmt.Sprintf("From: %s\nTo: %s\nAmount: %d\n", t.SenderID, t.RecipientID, t.Amount)
	return []byte(text), nil
}
