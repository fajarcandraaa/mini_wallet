package entity

import (
	"database/sql/driver"
	"time"
)

type WalletTrxType string

const (
	Deposit   WalletTrxType = "deposit"
	Withdrawl WalletTrxType = "withdrawl"
)

// WalletTrx is representation from database wallet_trx table
type WalletTransaction struct {
	WalletTrxID       string         `json:"wallet_trx_id" gorm:"size:36;not null;unique index;primaryKey"`
	WalletID          string         `json:"wallet_id" gorm:"size:36;"`
	WalletTrxType     *WalletTrxType `json:"wallet_trx_type" gorm:"column:wallet_trx_type;"`
	WalletBallanceTrx int            `json:"wallet_ballance_trx" gorm:"size:255;"`
	DepositedBy       string         `json:"deposited_by" gorm:"size:36;"`
	DepositedAt       *time.Time     `json:"deposited_at"`
	WithdrawnBy       string         `json:"withdrawn_by" gorm:"size:36;"`
	WithdrawnAt       *time.Time     `json:"withdrawn_at"`
	ReferenceID       string         `json:"reference_id" gorm:"size:36;"`
	CreatedAt         time.Time      `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt         time.Time      `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}

func (ct *WalletTrxType) Scan(value interface{}) error {
	*ct = WalletTrxType(value.([]byte))
	return nil
}

func (ct WalletTrxType) Value() (driver.Value, error) {
	return string(ct), nil
}

type WalletTransactions struct {
	WalletTrxID       string     `json:"wallet_trx_id"`
	WalletID          string     `json:"wallet_id"`
	WalletTrxType     string     `json:"wallet_trx_type"`
	WalletBallanceTrx int        `json:"wallet_ballance_trx"`
	DepositedBy       string     `json:"deposited_by"`
	DepositedAt       *time.Time `json:"deposited_at"`
	WithdrawnBy       string     `json:"withdrawn_by"`
	WithdrawnAt       *time.Time `json:"withdrawn_at"`
	ReferenceID       string     `json:"reference_id"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`
}
