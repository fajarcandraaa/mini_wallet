package service_test

import (
	"os"
	"testing"

	"github.com/fajarcandraaa/mini_wallet/internal/entity"
	rds "github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
)

func testConfig(t *testing.T) (*gorm.DB, *rds.Client, error) {
	var (
		dsn = "host=localhost user=postgres dbname=julo_backend_test sslmode=disable password=postgres port=5433"
	)

	db, err := gorm.Open("postgres", dsn) // initiate database for testing
	require.NoError(t, err)
	db.AutoMigrate(&entity.Wallet{}, &entity.WalletAccount{}, &entity.WalletTransaction{})

	err = godotenv.Load("../../.env") // Update the path accordingly
	if err != nil {
		return nil, nil, err
	}

	// Redis connection options
	options := &rds.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0, // Default DB
	}

	// // initiate redis client for testing
	client := rds.NewClient(options)

	return db, client, nil

}
