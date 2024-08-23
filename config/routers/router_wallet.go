package routers

import (
	"github.com/fajarcandraaa/mini_wallet/internal/service"
	"github.com/fajarcandraaa/mini_wallet/usecase"
)

func walletRouter(p *PathPrefix, s *service.Service) {
	var (
		walletUseCase = usecase.NewWalletUseCase(s)
	)

	p.V1.HandleFunc("/init", walletUseCase.InitializeAccountWallet).Methods("POST")

	p.Wallet.HandleFunc("", walletUseCase.EnabledWallet).Methods("POST")
	p.Wallet.HandleFunc("", walletUseCase.ViewBallance).Methods("GET")
	p.Wallet.HandleFunc("", walletUseCase.DisabledWallet).Methods("PATCH")
	p.Wallet.HandleFunc("/deposits", walletUseCase.TopUpBalance).Methods("POST")
	p.Wallet.HandleFunc("/withdrawals", walletUseCase.WithdrawlBalance).Methods("POST")
	p.Wallet.HandleFunc("/transactions", walletUseCase.MyWalletTransactions).Methods("GET")
	p.Wallet.HandleFunc("", walletUseCase.PermanentlyDeleteWallet).Methods("DELETE")
}
