package app

import (
	"github.com/fajarcandraaa/mini_wallet/internal/entity"
)

// SetMigrationTable is used to register entity model which want to be migrate
func SetMigrationTable() []interface{} {
	var migrationData = []interface{}{
		&entity.Wallet{},
		&entity.WalletAccount{},
		&entity.WalletTransaction{},
	}

	return migrationData
}
