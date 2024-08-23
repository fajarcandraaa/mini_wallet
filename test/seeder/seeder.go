package seeder

import (
	"context"
	"fmt"
	"time"

	"github.com/fajarcandraaa/mini_wallet/test/faker"
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
)

func SeedWalletAccountFaker(db *gorm.DB, rds *redis.Client) (string, error) {
	var (
		rdsKey = fmt.Sprintf("customerfaker %s", faker.FakeCustomerXID)
	)
	fakeWalletAccount := faker.FakeWalletAccount()
	err := db.Create(&fakeWalletAccount).Error
	if err != nil {
		return "", err
	}
	err = rds.Set(context.Background(), rdsKey, faker.FakeTokenValue, 5*time.Minute).Err()
	if err != nil {
		return "", err
	}
	return rdsKey, nil
}

func SeedEnabledWalletAccountFaker(db *gorm.DB, rds *redis.Client) (string, error) {
	var (
		rdsKey = fmt.Sprintf("customerfaker %s", faker.FakeCustomerXID)
	)
	fakeWallet := faker.FakeWallet()
	err := db.Create(&fakeWallet).Error
	if err != nil {
		return "", err
	}

	fakeWalletAccount := faker.FakeWalletAccountEnable()
	err = db.Create(&fakeWalletAccount).Error
	if err != nil {
		return "", err
	}
	err = rds.Set(context.Background(), rdsKey, faker.FakeTokenValue, 5*time.Minute).Err()
	if err != nil {
		return "", err
	}
	return rdsKey, nil
}
