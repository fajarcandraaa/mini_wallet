package repositories

import (
	"context"

	"github.com/fajarcandraaa/mini_wallet/internal/entity"
	"github.com/jinzhu/gorm"
)

type WalletTransactionsRepository struct {
	db *gorm.DB
}

func NewWalletTransactionsRepository(db *gorm.DB) *WalletTransactionsRepository {
	return &WalletTransactionsRepository{
		db: db,
	}
}

var _ WalletTransactionsRepositoryContract = &WalletTransactionsRepository{}

// GetTransaction implements WalletTransactionsRepositoryContract.
func (r *WalletTransactionsRepository) GetListTransactions(ctx context.Context, walletId string) ([]entity.WalletTransactions, error) {
	var (
		table_name         = "wallet_transactions"
		walletTransactions []entity.WalletTransactions
	)

	err := r.db.Table(table_name).Find(&walletTransactions, "wallet_id = ?", walletId).Error
	if err != nil {
		return nil, err
	}

	return walletTransactions, nil
}

// AddBalance implements WalletTransactionsRepositoryContract.
func (r *WalletTransactionsRepository) AddBalance(ctx context.Context, p entity.WalletTransaction) (*entity.WalletTransaction, error) {
	var (
		account     entity.WalletAccount
		reffID      = p.ReferenceID
		querySearch = `SELECT wallet_trx_id FROM wallet_transactions WHERE reference_id = $1 AND wallet_trx_type = 'deposit'`
		queryTrx    = `
			INSERT INTO wallet_transactions
			(wallet_trx_id, wallet_id, wallet_trx_type, wallet_ballance_trx, deposited_by, deposited_at, reference_id)
			VALUES ($1,$2,$3,$4,$5,$6,$7);
		`
		queryUpdate = `
			UPDATE wallets SET wallet_ballance = wallet_ballance + $1 WHERE customer_xid = $2;
		`
	)

	// Cek wallet status
	err := r.db.First(&account, "customer_xid = ?", p.DepositedBy).Error
	if err != nil {
		return nil, err
	}
	if *account.Status == entity.Disable {
		return nil, entity.ErrWalletsIsDisable
	}

	// preparation for database transaction
	arg := []interface{}{
		&p.WalletTrxID,
		&p.WalletID,
		&p.WalletTrxType,
		&p.WalletBallanceTrx,
		&p.DepositedBy,
		&p.DepositedAt,
		&p.ReferenceID,
	}
	argUpdate := []interface{}{
		&p.WalletBallanceTrx,
		&p.DepositedBy,
	}
	result, err := databaseTrxImplementor(r.db, reffID, querySearch, queryTrx, queryUpdate, arg, argUpdate)
	if err != nil {
		return nil, err
	}

	return result, nil

}

// SubtractBalance implements WalletTransactionsRepositoryContract.
func (r *WalletTransactionsRepository) SubtractBalance(ctx context.Context, p entity.WalletTransaction) (*entity.WalletTransaction, error) {
	var (
		account     entity.WalletAccount
		reffID      = p.ReferenceID
		querySearch = `SELECT wallet_trx_id FROM wallet_transactions WHERE reference_id = $1 AND wallet_trx_type = 'withdrawl'`
		queryTrx    = `
			INSERT INTO wallet_transactions
			(wallet_trx_id, wallet_id, wallet_trx_type, wallet_ballance_trx, withdrawn_by, withdrawn_at, reference_id)
			VALUES ($1,$2,$3,$4,$5,$6,$7);
		`
		queryUpdate = `
			UPDATE wallets SET wallet_ballance = wallet_ballance - $1 WHERE customer_xid = $2;
		`
	)

	// Cek wallet status
	err := r.db.First(&account, "customer_xid = ?", p.WithdrawnBy).Error
	if err != nil {
		return nil, err
	}
	if *account.Status == entity.Disable {
		return nil, entity.ErrWalletsIsDisable
	}

	// preparation for database transaction
	arg := []interface{}{
		&p.WalletTrxID,
		&p.WalletID,
		&p.WalletTrxType,
		&p.WalletBallanceTrx,
		&p.WithdrawnBy,
		&p.WithdrawnAt,
		&p.ReferenceID,
	}
	argUpdate := []interface{}{
		&p.WalletBallanceTrx,
		&p.WithdrawnBy,
	}
	result, err := databaseTrxImplementor(r.db, reffID, querySearch, queryTrx, queryUpdate, arg, argUpdate)
	if err != nil {
		return nil, err
	}

	return result, nil

}

// databaseTrx is handle transaction proccess on database
func databaseTrxImplementor(db *gorm.DB, reffID, querySearch, queryTrx, queryUpdate string, arg, argUpdate []interface{}) (*entity.WalletTransaction, error) {
	var (
		result entity.WalletTransaction
		exist  = false
	)

	tx := db.Begin()
	err := tx.Exec("SET TRANSACTION ISOLATION LEVEL SERIALIZABLE").Error
	if err != nil {
		return nil, err
	}

	ex := tx.Exec(querySearch, reffID)
	if ex.RowsAffected > 0 {
		exist = true
	}

	if exist {
		tx.Rollback()
		return nil, entity.ErrWalletAlreadyExist
	}

	err = tx.Exec(queryTrx, arg...).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Exec(queryUpdate, argUpdate...).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit().Error
	if err != nil {
		return nil, err
	}

	return &result, nil
}
