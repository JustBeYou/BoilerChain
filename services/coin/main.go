package coin

import (
	"boilerchain/blocks"
	"fmt"
)

// Coin -
type Coin struct {
	Name          string
	InitialWallet Wallet
	Store         blocks.BlockStore
}

// NewCoin -
func NewCoin(name string, initialAmount uint64, passphrase string) Coin {
	wallet := NewWallet(
		fmt.Sprintf("Initial wallet of %s.", name),
		passphrase,
	)
	store := blocks.NewInMemoryStore(
		TransactionChain{
			[]Transaction{
				NewTransaction(
					wallet,
					passphrase,
					fmt.Sprintf("Created initial wallet of %s.", name),
					WalletAnnouncement,
					WalletAnnouncementTransaction{
						wallet,
					},
				),
				NewTransaction(
					wallet,
					passphrase,
					fmt.Sprintf("Intial coin offering of %s, consiting of %d coins.", name, initialAmount),
					Mine,
					MineTransaction{
						wallet.ID,
						initialAmount,
					},
				),
			},
		},
	)

	return Coin{
		name,
		wallet,
		store,
	}
}
