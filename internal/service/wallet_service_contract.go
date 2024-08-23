package service

import (
	"context"

	"github.com/fajarcandraaa/mini_wallet/internal/presentation"
)

type WalletServiceContract interface {
	CreateAccount(ctx context.Context, payload presentation.InitiateWalletAccountRequest) (*presentation.InitiateWalletAccountResponse, error)
	DeleteAccount(ctx context.Context, token string) error
}
