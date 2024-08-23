package dto

import (
	"time"

	"github.com/fajarcandraaa/mini_wallet/internal/entity"
	"github.com/fajarcandraaa/mini_wallet/internal/presentation"
	"github.com/google/uuid"
)

func ListWalletTransactionToResponse(payload []entity.WalletTransactions) *presentation.ListTransactionResponse {
	var (
		transactions = []presentation.TrxDetailResponse{}
		transaction  presentation.TrxDetailResponse
	)

	for _, p := range payload {
		switch p.WalletTrxType {
		case "deposit":
			detail := &presentation.TrxDetailResponse{
				ID:          p.WalletTrxID,
				DepositedBy: p.DepositedBy,
				Status:      "success",
				DepositedAt: p.DepositedAt,
				Amount:      p.WalletBallanceTrx,
				ReffID:      p.ReferenceID,
			}
			transaction = *detail
		default:
			detail := &presentation.TrxDetailResponse{
				ID:          p.WalletTrxID,
				WithdrawnBy: p.WithdrawnBy,
				Status:      "success",
				WithdrawnAt: p.WithdrawnAt,
				Amount:      p.WalletBallanceTrx,
				ReffID:      p.ReferenceID,
			}
			transaction = *detail
		}
		transactions = append(transactions, transaction)
	}

	res := &presentation.ListTransactionResponse{
		Transactions: transactions,
	}

	return res
}

func WalletTrxToResponse(p entity.WalletTransaction) *presentation.DepositOrWithdrawlResponse {
	var (
		resp    presentation.DepositOrWithdrawlResponse
		trxType = p.WalletTrxType
	)
	switch *trxType {
	case entity.Deposit:
		detail := &presentation.TrxDetailResponse{
			ID:          p.WalletTrxID,
			DepositedBy: p.DepositedBy,
			Status:      "success",
			DepositedAt: p.DepositedAt,
			Amount:      p.WalletBallanceTrx,
			ReffID:      p.ReferenceID,
		}

		resp = presentation.DepositOrWithdrawlResponse{
			Deposit: detail,
		}
	default:
		detail := &presentation.TrxDetailResponse{
			ID:          p.WalletTrxID,
			WithdrawnBy: p.WithdrawnBy,
			Status:      "success",
			WithdrawnAt: p.WithdrawnAt,
			Amount:      p.WalletBallanceTrx,
			ReffID:      p.ReferenceID,
		}
		resp = presentation.DepositOrWithdrawlResponse{
			Withdrawl: detail,
		}
	}

	return &resp
}

func AddBalanceRequest(amount int, reffID string) presentation.AddBalanceRequest {
	resp := presentation.AddBalanceRequest{
		Amount:  amount,
		ReffID:  reffID,
		TrxType: entity.Deposit,
	}

	return resp
}

func WithdrawlBalanceRequest(amount int, reffID string) presentation.AddBalanceRequest {
	resp := presentation.AddBalanceRequest{
		Amount:  amount,
		ReffID:  reffID,
		TrxType: entity.Withdrawl,
	}

	return resp
}

func AddBalanceRequestToDatabase(p presentation.AddBalanceRequest, d presentation.CustomerDataByTokenResponse) *entity.WalletTransaction {
	t := time.Now()
	res := entity.WalletTransaction{
		WalletTrxID:       uuid.NewString(),
		WalletID:          d.WalletID,
		WalletTrxType:     &p.TrxType,
		WalletBallanceTrx: p.Amount,
		DepositedBy:       d.CustomerID,
		DepositedAt:       &t,
		ReferenceID:       p.ReffID,
	}

	return &res
}

func WithdrawlBalanceRequestToDatabase(p presentation.AddBalanceRequest, d presentation.CustomerDataByTokenResponse) *entity.WalletTransaction {
	t := time.Now()
	res := entity.WalletTransaction{
		WalletTrxID:       uuid.NewString(),
		WalletID:          d.WalletID,
		WalletTrxType:     &p.TrxType,
		WalletBallanceTrx: p.Amount,
		WithdrawnBy:       d.CustomerID,
		WithdrawnAt:       &t,
		ReferenceID:       p.ReffID,
	}

	return &res
}

func CustomerXidToDatabase(p presentation.InitiateWalletAccountRequest) *presentation.NewWalletAccountRequest {
	res := &presentation.NewWalletAccountRequest{
		AccountID:      uuid.NewString(),
		WalletID:       uuid.NewString(),
		CustomerXid:    p.CustomerXid,
		WalletStatus:   entity.Disable,
		WalletBallance: 0,
	}

	return res
}

func WalletAccountToResponse(p entity.WalletAccount) presentation.WalletDataResponse {
	var (
		res presentation.WalletDetailDataResponse
	)
	switch *p.Status {
	case "enabled":
		res = presentation.WalletDetailDataResponse{
			ID:        p.AccountID,
			OwnedBy:   p.CustomerXid,
			Status:    string(*p.Status),
			EnabledAt: *p.EnabledAt,
			Balance:   p.WalletBallance,
		}
	default:
		res = presentation.WalletDetailDataResponse{
			ID:         p.AccountID,
			OwnedBy:    p.CustomerXid,
			Status:     string(*p.Status),
			DisabledAt: *p.DisabledAt,
			Balance:    p.WalletBallance,
		}
	}

	result := presentation.WalletDataResponse{
		Wallet: res,
	}

	return result
}

func TokenToResponse(t string) presentation.InitiateWalletAccountResponse {
	res := presentation.InitiateWalletAccountResponse{
		Token: t,
	}

	return res
}
