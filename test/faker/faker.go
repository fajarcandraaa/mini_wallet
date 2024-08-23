package faker

import (
	"time"

	"github.com/fajarcandraaa/mini_wallet/internal/entity"
)

const (
	FakeToken                      = "Token 1d54be82cbad3ec463dcd0ceab26fb409d5e4a52"
	FakeTokenValue                 = "1d54be82cbad3ec463dcd0ceab26fb409d5e4a52"
	FakeAccountID                  = "607bf44c-102a-4f4b-89ce-4b75d67ce415"
	FakeCustomerXID                = "ea0212d3-abd6-406f-8c67-868e814a2436"
	FakeWalletID                   = "fc70c26a-5922-4f83-843f-61859f9cf55c"
	FakeWalletAccountStatusDisable = entity.Disable
	FakeWalletAccountStatusEnable  = entity.Enable
	FakeReffID1                    = "49701238-82cd-40c6-9c26-2e546ecc3810"
	FakeReffID2                    = "01fb925f-74f4-4830-b703-36162cba53ba"
	FakeReffID3                    = "bc834b22-3fb5-412f-8325-c5710c1b3704"
)

func FakeWallet() *entity.Wallet {
	var (
		t = time.Now()
	)

	fakeWallet := &entity.Wallet{
		WalletID:       FakeWalletID,
		CustomerXid:    FakeCustomerXID,
		WalletBallance: 0,
		CreatedAt:      t,
		UpdatedAt:      t,
	}

	return fakeWallet
}

func FakeWalletAccount() *entity.WalletAccount {

	var (
		status entity.WalletStatus
		t      = time.Now()
	)
	status = FakeWalletAccountStatusDisable

	fakeWalletAccount := &entity.WalletAccount{
		AccountID:      FakeAccountID,
		CustomerXid:    FakeCustomerXID,
		Status:         &status,
		WalletBallance: 0,
		CreatedAt:      t,
		UpdatedAt:      t,
	}

	return fakeWalletAccount
}

func FakeWalletAccountEnable() *entity.WalletAccount {

	var (
		status entity.WalletStatus
		t      = time.Now()
	)
	status = FakeWalletAccountStatusEnable

	fakeWalletAccount := &entity.WalletAccount{
		AccountID:      FakeAccountID,
		CustomerXid:    FakeCustomerXID,
		Status:         &status,
		WalletBallance: 0,
		EnabledAt:      &t,
		CreatedAt:      t,
		UpdatedAt:      t,
	}

	return fakeWalletAccount
}
