package repositories

import (
	"context"

	"github.com/fajarcandraaa/mini_wallet/internal/entity"
)

type WalletAccountRepositoryContract interface {
	UpdateStatus(ctx context.Context, status entity.WalletStatus, custromerXid string) (*entity.WalletAccount, error)
	GetBalanceByCustomerXID(ctx context.Context, customerXID string) (*entity.WalletAccount, error)
}
