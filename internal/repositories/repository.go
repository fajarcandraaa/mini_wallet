package repositories

import "github.com/jinzhu/gorm"

type Repository struct {
	Wallet            WalletRepositoryContract
	WalletAccount     WalletAccountRepositoryContract
	WalletTransaction WalletTransactionsRepositoryContract
}

func NewWallet(db *gorm.DB) WalletRepositoryContract {
	return NewWalletRepository(db)
}

func NewWalletAccount(db *gorm.DB) WalletAccountRepositoryContract {
	return NewWalletAccountRepository(db)
}

func NewWalletTransactions(db *gorm.DB) WalletTransactionsRepositoryContract  {
	return NewWalletTransactionsRepository(db)
}

// NewRepository to setting repositories
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Wallet:            NewWallet(db),
		WalletAccount:     NewWalletAccount(db),
		WalletTransaction: NewWalletTransactions(db),
	}
}
