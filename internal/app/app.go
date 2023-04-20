package app

import (
	"fmt"
	"github.com/gorilla/rpc"
	"github.com/robbiekes/goods-manager-api/config"
	v1 "github.com/robbiekes/goods-manager-api/internal/controller/jsonrpc"
	"github.com/robbiekes/goods-manager-api/internal/service"
	"github.com/robbiekes/goods-manager-api/internal/service/repository"
	"github.com/robbiekes/goods-manager-api/pkg/httpserver"
	"github.com/robbiekes/goods-manager-api/pkg/postgres"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
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
	services := service.NewService(repository.NewRepo(nil))

	// RPC server
	log.Info("Initializing RPC server...")
	rpcServer := rpc.NewServer()
	v1.NewRpcRouter(rpcServer, services)
	httpServer := httpserver.New(rpcServer, httpserver.Port(cfg.HTTP.Port))

	// Waiting signal
	log.Info("Configuring graceful shutdown...")
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		log.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// // Graceful shutdown
	log.Info("Shutting down...")
	err = httpServer.Shutdown()
	if err != nil {
		log.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
