package service

import (
	"context"

	"github.com/fajarcandraaa/mini_wallet/helpers"
	"github.com/fajarcandraaa/mini_wallet/internal/dto"
	"github.com/fajarcandraaa/mini_wallet/internal/entity"
	"github.com/fajarcandraaa/mini_wallet/internal/presentation"
	"github.com/fajarcandraaa/mini_wallet/internal/repositories"
	"github.com/go-redis/redis/v8"
)

type walletAccount struct {
	repo *repositories.Repository
	rds  *redis.Client
}

func NewWalletAccountService(repo *repositories.Repository, rds *redis.Client) *walletAccount {
	return &walletAccount{
		repo: repo,
		rds:  rds,
	}
}

var _ WalletAccountContract = &walletAccount{}

// DisableWallet implements WalletAccountContract.
func (s *walletAccount) DisableWallet(ctx context.Context, token string) (*presentation.WalletDataResponse, error) {
	tokenKey, err := helpers.FindCustomerXidFromToken(ctx, s.rds, token)
	if err != nil {
		return nil, err
	}

	customerXid := helpers.GetCustomerXidFromToken(tokenKey)
	disabledWallet, err := s.repo.WalletAccount.UpdateStatus(ctx, entity.Disable, customerXid)
	if err != nil {
		return nil, err
	}

	rsp := dto.WalletAccountToResponse(*disabledWallet)

	return &rsp, nil
}

// EnableWallet implements WalletAccountContract.
func (s *walletAccount) EnableWallet(ctx context.Context, token string) (*presentation.WalletDataResponse, error) {
	tokenKey, err := helpers.FindCustomerXidFromToken(ctx, s.rds, token)
	if err != nil {
		return nil, err
	}

	customerXid := helpers.GetCustomerXidFromToken(tokenKey)
	enabledWallet, err := s.repo.WalletAccount.UpdateStatus(ctx, entity.Enable, customerXid)
	if err != nil {
		return nil, err
	}

	rsp := dto.WalletAccountToResponse(*enabledWallet)

	return &rsp, nil
}

// ViewBallance implements WalletAccountContract.
func (s *walletAccount) ViewBallanceOnWallet(ctx context.Context, token string) (*presentation.WalletDataResponse, error) {
	tokenKey, err := helpers.FindCustomerXidFromToken(ctx, s.rds, token)
	if err != nil {
		return nil, err
	}

	customerXid := helpers.GetCustomerXidFromToken(tokenKey)
	viewBallance, err := s.repo.WalletAccount.GetBalanceByCustomerXID(ctx, customerXid)
	if err != nil {
		return nil, err
	}

	rsp := dto.WalletAccountToResponse(*viewBallance)

	return &rsp, nil
}
