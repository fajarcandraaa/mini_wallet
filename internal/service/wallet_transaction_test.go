package service_test

import (
	"context"
	"testing"

	"github.com/fajarcandraaa/mini_wallet/helpers"
	"github.com/fajarcandraaa/mini_wallet/internal/entity"
	"github.com/fajarcandraaa/mini_wallet/internal/repositories"
	"github.com/fajarcandraaa/mini_wallet/internal/service"
	"github.com/fajarcandraaa/mini_wallet/test/faker"
	"github.com/fajarcandraaa/mini_wallet/test/seeder"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTopUpVirtualMoney(t *testing.T) {
	db, rds, err := testConfig(t)
	require.NoError(t, err)
	defer db.DropTable(&entity.Wallet{}, &entity.WalletAccount{}, &entity.WalletTransaction{})

	r := repositories.NewRepository(db)
	walletTransactionService := service.NewWalletTransactionsService(r, rds)

	key, err := seeder.SeedEnabledWalletAccountFaker(db, rds)
	require.NoError(t, err)

	t.Run("feature view topup virtual money : if data is valid, expected no error", func(t *testing.T) {
		var (
			ctx    = context.Background()
			token  = faker.FakeToken
			amount = 15000
			reffID = faker.FakeReffID1
		)
		tokenString, err := helpers.ParseTokenHex(token)
		require.NoError(t, err)

		res, err := walletTransactionService.TopUpVirtualMoney(ctx, amount, reffID, tokenString)
		require.NoError(t, err)
		require.Equal(t, err, nil)
		assert.NotNil(t, res)
	})

	t.Run("feature view topup virtual money : if data is duplicate, expected error", func(t *testing.T) {
		var (
			ctx    = context.Background()
			token  = faker.FakeToken
			amount = 15000
			reffID = faker.FakeReffID1
		)
		tokenString, err := helpers.ParseTokenHex(token)
		require.NoError(t, err)

		rsp, err := walletTransactionService.TopUpVirtualMoney(ctx, amount, reffID, tokenString)
		require.Error(t, err)
		assert.Nil(t, rsp)
	})

	_, err = rds.Del(context.Background(), key).Result()
	require.NoError(t, err)
}

func TestWithdrawlVirtualMoney(t *testing.T) {
	db, rds, err := testConfig(t)
	require.NoError(t, err)
	defer db.DropTable(&entity.Wallet{}, &entity.WalletAccount{}, &entity.WalletTransaction{})

	r := repositories.NewRepository(db)
	walletTransactionService := service.NewWalletTransactionsService(r, rds)

	key, err := seeder.SeedEnabledWalletAccountFaker(db, rds)
	require.NoError(t, err)

	t.Run("feature view use virtual money : if data is valid, expected no error", func(t *testing.T) {
		var (
			ctx    = context.Background()
			token  = faker.FakeToken
			amount = 15000
			reffID = faker.FakeReffID1
		)
		tokenString, err := helpers.ParseTokenHex(token)
		require.NoError(t, err)

		res, err := walletTransactionService.UseVirtualMoney(ctx, amount, reffID, tokenString)
		require.NoError(t, err)
		require.Equal(t, err, nil)
		assert.NotNil(t, res)
	})

	t.Run("feature view use virtual money : if data is duplicate, expected error", func(t *testing.T) {
		var (
			ctx    = context.Background()
			token  = faker.FakeToken
			amount = 15000
			reffID = faker.FakeReffID1
		)
		tokenString, err := helpers.ParseTokenHex(token)
		require.NoError(t, err)

		rsp, err := walletTransactionService.UseVirtualMoney(ctx, amount, reffID, tokenString)
		require.Error(t, err)
		assert.Nil(t, rsp)
	})

	_, err = rds.Del(context.Background(), key).Result()
	require.NoError(t, err)
}
