package service

import (
	"context"
	"regexp"

	"github.com/fajarcandraaa/mini_wallet/helpers"
	"github.com/fajarcandraaa/mini_wallet/internal/dto"
	"github.com/fajarcandraaa/mini_wallet/internal/presentation"
	"github.com/fajarcandraaa/mini_wallet/internal/repositories"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-redis/redis/v8"
)

type walletTransactions struct {
	repo *repositories.Repository
	rds  *redis.Client
}

func NewWalletTransactionsService(repo *repositories.Repository, rds *redis.Client) *walletTransactions {
	return &walletTransactions{
		repo: repo,
		rds:  rds,
	}
}

var _ WalletTransactionContract = &walletTransactions{}

// ListTransactionsByWallerID implements WalletTransactionContract.
func (s *walletTransactions) ListTransactionsByWallerID(ctx context.Context, token string) (*presentation.ListTransactionResponse, error) {
	tokenKey, err := helpers.FindCustomerXidFromToken(ctx, s.rds, token)
	if err != nil {
		return nil, err
	}

	customerXid := helpers.GetCustomerXidFromToken(tokenKey)
	customerDetail, err := s.repo.Wallet.GetDataCustomerByToken(ctx, customerXid)
	if err != nil {
		return nil, err
	}

	listTrx, err := s.repo.WalletTransaction.GetListTransactions(ctx, customerDetail.WalletID)
	if err != nil {
		return nil, err
	}

	res := dto.ListWalletTransactionToResponse(listTrx)
	return res, nil
}

// TopUpVirtualMoney implements WalletTransactionContract.
func (s *walletTransactions) TopUpVirtualMoney(ctx context.Context, amount int, reffID, token string) (*presentation.DepositOrWithdrawlResponse, error) {
	tokenKey, err := helpers.FindCustomerXidFromToken(ctx, s.rds, token)
	if err != nil {
		return nil, err
	}

	customerXid := helpers.GetCustomerXidFromToken(tokenKey)
	customerDetail, err := s.repo.Wallet.GetDataCustomerByToken(ctx, customerXid)
	if err != nil {
		return nil, err
	}
	payloadBalance := dto.AddBalanceRequest(amount, reffID)
	err = validation.ValidateStruct(&payloadBalance,
		validation.Field(&payloadBalance.ReffID, validation.Required, validation.Match(regexp.MustCompile(`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`)), validation.Length(3, 100)),
	)
	if err != nil {
		return nil, err
	}

	payloadToDB := dto.AddBalanceRequestToDatabase(payloadBalance, *customerDetail)
	if err != nil {
		return nil, err
	}

	_, err = s.repo.WalletTransaction.AddBalance(ctx, *payloadToDB)
	if err != nil {
		return nil, err
	}

	response := dto.WalletTrxToResponse(*payloadToDB)
	return response, nil
}

// UseVirtualMoney implements WalletTransactionContract.
func (s *walletTransactions) UseVirtualMoney(ctx context.Context, amount int, reffID string, token string) (*presentation.DepositOrWithdrawlResponse, error) {
	tokenKey, err := helpers.FindCustomerXidFromToken(ctx, s.rds, token)
	if err != nil {
		return nil, err
	}

	customerXid := helpers.GetCustomerXidFromToken(tokenKey)
	customerDetail, err := s.repo.Wallet.GetDataCustomerByToken(ctx, customerXid)
	if err != nil {
		return nil, err
	}
	payloadBalance := dto.WithdrawlBalanceRequest(amount, reffID)
	err = validation.ValidateStruct(&payloadBalance,
		validation.Field(&payloadBalance.ReffID, validation.Required, validation.Match(regexp.MustCompile(`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`)), validation.Length(3, 100)),
	)
	if err != nil {
		return nil, err
	}

	payloadToDB := dto.WithdrawlBalanceRequestToDatabase(payloadBalance, *customerDetail)
	if err != nil {
		return nil, err
	}

	_, err = s.repo.WalletTransaction.SubtractBalance(ctx, *payloadToDB)
	if err != nil {
		return nil, err
	}

	response := dto.WalletTrxToResponse(*payloadToDB)
	return response, nil
}
