package routers

import "github.com/gorilla/mux"

type PathPrefix struct {
	V1     *mux.Router
	Wallet *mux.Router
}

func RouterConfigPrefix(se *Serve) *PathPrefix {
	var (
		api    = se.Router.PathPrefix("/api").Subrouter()
		v1     = api.PathPrefix("/v1").Subrouter()
		wallet = v1.PathPrefix("/wallet").Subrouter()
	)

	result := &PathPrefix{
		V1:     v1,
		Wallet: wallet,
	}

	return result
}
