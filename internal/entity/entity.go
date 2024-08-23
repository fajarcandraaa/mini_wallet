package entity

type Error string

// Declare error messege
const (
	ErrPermissionNotAllowed = Error("permission.not_allowed")

	//User Error
	ErrWalletNotExist             = Error("domain.wallet.error.not_exist")
	ErrWalletAlreadyExist         = Error("domain.wallet.error.already_exist")
	ErrWalletsCredentialNotExist  = Error("domain.wallet.error.credential_not_exist")
	ErrWalletsUnprocessableEntity = Error("domain.wallet.error.unprocessable_entity")
	ErrWalletsIsDisable           = Error("domain.wallet.error.disable")
)

func (e Error) Error() string {
	return string(e)
}
