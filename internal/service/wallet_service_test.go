package service_test

import (
	"context"
	"os"
	"testing"

	"github.com/fajarcandraaa/mini_wallet/internal/entity"
	"github.com/fajarcandraaa/mini_wallet/internal/presentation"
	"github.com/fajarcandraaa/mini_wallet/internal/repositories"
	"github.com/fajarcandraaa/mini_wallet/internal/service"
	rds "github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInsetNewWalletAccount(t *testing.T) {
	var (
		ctx = context.Background()
		dsn = "host=localhost user=postgres dbname=julo_backend_test sslmode=disable password=postgres port=5433"
	)

	db, err := gorm.Open("postgres", dsn) // initiate database for testing
	require.NoError(t, err)
	db.AutoMigrate(&entity.Wallet{}, &entity.WalletAccount{})
	defer db.DropTable(&entity.Wallet{}, &entity.WalletAccount{})

	err = godotenv.Load("../../.env") // Update the path accordingly
	require.NoError(t, err)

	// Redis connection options
	options := &rds.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0, // Default DB
	}

	// // initiate redis client for testing
	client := rds.NewClient(options)
	require.NoError(t, err)

	r := repositories.NewRepository(db)
	walletService := service.NewWalletService(r, client)

	t.Run("if store data given valid data, expected no error", func(t *testing.T) {
		var (
			payload = presentation.InitiateWalletAccountRequest{
				CustomerXid: "ea0212d3-abd6-406f-8c67-868e814a2436",
			}
		)
		res, err := walletService.CreateAccount(ctx, payload)
		require.NoError(t, err)
		require.Equal(t, err, nil)
		assert.NotNil(t, res)
	})

	t.Run("if store data given not valid data, expected  error", func(t *testing.T) {
		var (
			payload = presentation.InitiateWalletAccountRequest{
				CustomerXid: "123456",
			}
		)
		res, err := walletService.CreateAccount(ctx, payload)
		require.Error(t, err)
		assert.Nil(t, res)
	})

}
