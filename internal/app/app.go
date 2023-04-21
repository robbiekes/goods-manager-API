package app

import (
	"fmt"
	"github.com/gorilla/rpc"
	"github.com/robbiekes/goods-manager-api/config"
	v1 "github.com/robbiekes/goods-manager-api/internal/controller/jsonrpc"
	"github.com/robbiekes/goods-manager-api/internal/service"
	"github.com/robbiekes/goods-manager-api/internal/service/repository"
	"github.com/robbiekes/goods-manager-api/pkg/postgres"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func Run(cfg *config.Config) {
	// Repository
	log.Info("Initializing postgres storage")
	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		log.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer pg.Close()

	// Service
	log.Info("Initializing service...")
	services := service.NewService(repository.NewRepo(pg))

	// RPC server
	log.Info("Initializing RPC server...")
	rpcServer := rpc.NewServer()
	r := v1.NewRpcRouter(rpcServer, services)

	http.ListenAndServe(":8080", r)
}
