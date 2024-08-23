package presentation

import (
	"time"

	"github.com/fajarcandraaa/mini_wallet/internal/entity"
)

// Request
type (
	InitiateWalletAccountRequest struct {
		CustomerXid string `json:"customer_xid"`
	}

	NewWalletAccountRequest struct {
		AccountID      string              `json:"account_id"`
		WalletID       string              `json:"wallet_id"`
		CustomerXid    string              `json:"customer_xid"`
		WalletStatus   entity.WalletStatus `json:"wallet_status"`
		WalletBallance int                 `json:"wallet_ballance"`
	}

	AddBalanceRequest struct {
		Amount  int                  `json:"amount"`
		ReffID  string               `json:"reference_id"`
		TrxType entity.WalletTrxType `json:"trx_type"`
	}

	DetailBalanceRequest struct {
		ReffID            string               `json:"reference_id"`
		WalletID          string               `json:"wallet_id"`
		WalletTrxType     entity.WalletTrxType `json:"wallet_trx_type"`
		WalletBallanceTrx int                  `json:"wallet_ballance_trx"`
		DepositedBy       string               `json:"deposited_by"`
		DepositedAt       *time.Time           `json:"deposited_at"`
		WithdrawnBy       string               `json:"withdrawn_by"`
		WithdrawnAt       *time.Time           `json:"withdrawn_at"`
	}
)

// Response
type (
	InitiateWalletAccountResponse struct {
		Token string `json:"token"`
	}

	CustomerDataByTokenResponse struct {
		WalletID   string `json:"wallet_id"`
		AccountID  string `json:"account_id"`
		CustomerID string `json:"customer_xid"`
	}

	WalletDataResponse struct {
		Wallet WalletDetailDataResponse `json:"wallet"`
	}

	WalletDetailDataResponse struct {
		ID         string    `json:"id"`
		OwnedBy    string    `json:"owned_by"`
		Status     string    `json:"status"`
		EnabledAt  time.Time `json:"enabled_at,omitempty"`
		DisabledAt time.Time `json:"disabled_at,omitempty"`
		Balance    int       `json:"balance"`
	}

	TrxDetailResponse struct {
		ID          string     `json:"id,omitempty"`
		DepositedBy string     `json:"deposited_by,omitempty"`
		WithdrawnBy string     `json:"withdrawn_by,omitempty"`
		Status      string     `json:"status,omitempty"`
		DepositedAt *time.Time `json:"deposited_at,omitempty"`
		WithdrawnAt *time.Time `json:"withdrawn_at,omitempty"`
		Amount      int        `json:"amount,omitempty"`
		ReffID      string     `json:"reference_id,omitempty"`
	}

	DepositOrWithdrawlResponse struct {
		Deposit   *TrxDetailResponse `json:"deposit,omitempty"`
		Withdrawl *TrxDetailResponse `json:"withdrawl,omitempty"`
	}

	ListTransactionResponse struct {
		Transactions []TrxDetailResponse `json:"transactions"`
	}

	WalletTransactionDetailResponse struct {
		ID                string                `json:"id"`
		WalletID          string                `json:"wallet_id"`
		WalletTrxType     *entity.WalletTrxType `json:"wallet_trx_type"`
		WalletBallanceTrx int                   `json:"wallet_ballance_trx"`
		DepositedBy       string                `json:"deposited_by,omitempty"`
		DepositedAt       *time.Time            `json:"deposited_at,omitempty"`
		WithdrawnBy       string                `json:"withdrawn_by,omitempty"`
		WithdrawnAt       *time.Time            `json:"withdrawn_at,omitempty"`
		CreatedAt         time.Time             `json:"created_at,omitempty"`
		UpdatedAt         time.Time             `json:"updated_at,omitempty"`
	}

	WalletTransactionListResponse struct {
		Transactions []WalletTransactionDetailResponse `json:"transactions"`
	}
)
