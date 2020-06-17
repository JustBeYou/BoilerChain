package coin

import (
	"bytes"
	"encoding"
	"encoding/binary"
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
	transactions []Transaction
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

// MarshalBinary -
func (tc TransactionChain) MarshalBinary() ([]byte, error) {
	bytes := []byte{}
	for _, t := range tc.transactions {
		tb, _ := t.MarshalBinary()
		bytes = append(bytes, tb...)
	}
	return bytes, nil
}

// MarshalText -
func (tc TransactionChain) MarshalText() ([]byte, error) {
	buf := bytes.Buffer{}
	buf.WriteString("\n")
	for _, t := range tc.transactions {
		tb, _ := t.MarshalText()
		buf.Write(tb)
	}
	buf.WriteString("--- --- ---\n")
	return buf.Bytes(), nil
}

// MarshalBinary -
func (t Transaction) MarshalBinary() ([]byte, error) {
	bytes := []byte{}
	bytes = append(bytes, []byte(t.ID)...)
	bytes = append(bytes, []byte(t.PublisherID)...)
	bytes = append(bytes, []byte(t.Description)...)
	bytes = append(bytes, toByteArray(int64(t.Type))...)
	binaryData, _ := t.Data.(encoding.BinaryMarshaler).MarshalBinary()
	bytes = append(bytes, binaryData...)
	return bytes, nil
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

// MarshalBinary -
func (t MineTransaction) MarshalBinary() ([]byte, error) {
	bytes := []byte{}
	bytes = append(bytes, []byte(t.BeneficiaryID)...)
	bytes = append(bytes, toByteArray(int64(t.Amount))...)
	return bytes, nil
}

// MarshalText -
func (t MineTransaction) MarshalText() ([]byte, error) {
	text := fmt.Sprintf("Beneficiary ID: %s\nAmount: %d\n", t.BeneficiaryID, t.Amount)
	return []byte(text), nil
}

// MarshalBinary -
func (t WalletAnnouncementTransaction) MarshalBinary() ([]byte, error) {
	bytes := []byte{}
	binaryWallet, _ := t.Created.MarshalBinary()
	bytes = append(bytes, binaryWallet...)
	return bytes, nil
}

// MarshalText -
func (t WalletAnnouncementTransaction) MarshalText() ([]byte, error) {
	walletText, _ := t.Created.MarshalText()
	text := fmt.Sprintf("Wallet: %s\n", walletText)
	return []byte(text), nil
}

// MarshalBinary -
func (t PaymentTransaction) MarshalBinary() ([]byte, error) {
	bytes := []byte{}
	bytes = append(bytes, []byte(t.RecipientID)...)
	bytes = append(bytes, []byte(t.SenderID)...)
	bytes = append(bytes, toByteArray(int64(t.Amount))...)
	return bytes, nil
}

// MarshalText -
func (t PaymentTransaction) MarshalText() ([]byte, error) {
	text := fmt.Sprintf("From: %s\nTo: %s\nAmount: %d\n", t.SenderID, t.RecipientID, t.Amount)
	return []byte(text), nil
}

func toByteArray(i int64) []byte {
	iAsBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(iAsBytes, uint64(i))
	return iAsBytes
}
