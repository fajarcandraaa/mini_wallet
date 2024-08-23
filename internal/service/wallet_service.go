package service

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"time"

	"github.com/fajarcandraaa/mini_wallet/helpers"
	"github.com/fajarcandraaa/mini_wallet/internal/dto"
	"github.com/fajarcandraaa/mini_wallet/internal/presentation"
	"github.com/fajarcandraaa/mini_wallet/internal/repositories"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-redis/redis/v8"
)

type walletService struct {
	repo *repositories.Repository
	rds  *redis.Client
}

func NewWalletService(repo *repositories.Repository, rds *redis.Client) *walletService {
	return &walletService{
		repo: repo,
		rds:  rds,
	}
}

var _ WalletServiceContract = &walletService{}

// DeleteAccount implements WalletServiceContract.
func (s *walletService) DeleteAccount(ctx context.Context, token string) error {
	tokenKey, err := helpers.FindCustomerXidFromToken(ctx, s.rds, token)
	if err != nil {
		return err
	}

	customerXid := helpers.GetCustomerXidFromToken(tokenKey)
	customerDetail, err := s.repo.Wallet.GetDataCustomerByToken(ctx, customerXid)
	if err != nil {
		return err
	}
	// walletDetail, err := s.repo.WalletAccount.GetBalanceByCustomerXID(ctx, customerXid)
	// if err != nil {
	// 	return err
	// }

	err = s.repo.Wallet.DeleteWallet(ctx, customerDetail.CustomerID, customerDetail.WalletID)
	if err != nil {
		return err
	}

	return nil
}

// CreateAccount implements WalletServiceContract.
func (s *walletService) CreateAccount(ctx context.Context, payload presentation.InitiateWalletAccountRequest) (*presentation.InitiateWalletAccountResponse, error) {
	var (
		rdsKey     = fmt.Sprintf("customer %s", payload.CustomerXid)
		expiration = 7 * 24 * time.Hour
	)

	err := validation.ValidateStruct(&payload,
		validation.Field(&payload.CustomerXid, validation.Required, validation.Match(regexp.MustCompile(`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`)), validation.Length(3, 100)),
	)
	if err != nil {
		return nil, err
	}

	payloadData := dto.CustomerXidToDatabase(payload)
	result, err := s.repo.Wallet.StoreNewWallet(ctx, *payloadData)
	if err != nil {
		return nil, err
	}

	go func() {
		err = s.rds.Set(context.Background(), rdsKey, *result, expiration).Err()
		if err != nil {
			log.Printf("ERROR Redis Set : %s", err.Error())
		}

	}()

	resp := dto.TokenToResponse(*result)
	return &resp, nil
}
