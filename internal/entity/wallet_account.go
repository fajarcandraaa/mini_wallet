package entity

import (
	"database/sql/driver"
	"time"
)

type WalletStatus string

const (
	Enable  WalletStatus = "enabled"
	Disable WalletStatus = "disabled"
)

// WalletAccount is representation from database wallet_account table
type WalletAccount struct {
	AccountID      string        `json:"account_id" gorm:"size:36;not null;unique index;primaryKey"`
	CustomerXid    string        `json:"customer_xid" gorm:"size:36;"`
	Status         *WalletStatus `json:"wallet_status" gorm:"column:wallet_status;"`
	WalletBallance int           `json:"wallet_ballance" gorm:"size:255;"`
	EnabledAt      *time.Time    `json:"enabled_at,omitempty"`
	DisabledAt     *time.Time    `json:"disabled_at,omitempty"`
	CreatedAt      time.Time     `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt      time.Time     `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}

func (ct *WalletStatus) WalletAccountScan(value interface{}) error {
	*ct = WalletStatus(value.([]byte))
	return nil
}

func (ct WalletStatus) WalletAccountValue() (driver.Value, error) {
	return string(ct), nil
}
