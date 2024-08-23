package routers

import (
	"fmt"
	"log"

	"github.com/fajarcandraaa/mini_wallet/internal/repositories"
	"github.com/fajarcandraaa/mini_wallet/internal/service"
	"github.com/fajarcandraaa/mini_wallet/pkg/storage/redis"
)

func (se *Serve) initializeRoutes() {
	rds, err := redis.NewRedisConfig()
	if err != nil {
		fmt.Printf("Cannot connect to redis")
		log.Fatal("This is the error:", err)
	}

	p := RouterConfigPrefix(se)            // set grouping prefix
	r := repositories.NewRepository(se.DB) //initiate repository
	s := service.NewService(r, rds)        //initiate service

	// //initiate endpoint
	walletRouter(p, s)

}
