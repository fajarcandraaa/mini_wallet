package entity

import "time"

// Wallet is representation from database wallet table
type Wallet struct {
	WalletID       string    `json:"wallet_id" gorm:"size:36;not null;unique index;primaryKey"`
	CustomerXid    string    `json:"customer_xid" gorm:"size:36;"`
	WalletBallance int       `json:"wallet_ballance" gorm:"size:255;"`
	CreatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
