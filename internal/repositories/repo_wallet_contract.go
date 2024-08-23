package repositories

import (
	"context"

	"github.com/fajarcandraaa/mini_wallet/internal/presentation"
)

type WalletRepositoryContract interface {
	StoreNewWallet(ctx context.Context, payload presentation.NewWalletAccountRequest) (*string, error)
	GetDataCustomerByToken(ctx context.Context, customerXid string) (*presentation.CustomerDataByTokenResponse, error)
	DeleteWallet(ctx context.Context, customerXid, walletId string) error
}
