package service

import (
	"context"

	"github.com/fajarcandraaa/mini_wallet/internal/presentation"
)

type WalletTransactionContract interface {
	ListTransactionsByWallerID(ctx context.Context, token string) (*presentation.ListTransactionResponse, error)
	TopUpVirtualMoney(ctx context.Context, amount int, reffID, token string) (*presentation.DepositOrWithdrawlResponse, error)
	UseVirtualMoney(ctx context.Context, amount int, reffID, token string) (*presentation.DepositOrWithdrawlResponse, error)
}
