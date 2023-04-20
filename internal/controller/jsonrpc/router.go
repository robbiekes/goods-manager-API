package jsonrpc

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
	"github.com/robbiekes/goods-manager-api/internal/service"
	log "github.com/sirupsen/logrus"
)

type Router struct {
	service *service.GoodsManagerService
}

func NewRpcRouter(rpcServer *rpc.Server, s *service.GoodsManagerService) *mux.Router {
	rpcServer.RegisterCodec(json.NewCodec(), "application/json")
	rpcServer.RegisterCodec(json.NewCodec(), "application/json;charset=UTF-8")

	router := &Router{service: s}

	err := rpcServer.RegisterService(router, "")
	if err != nil {
		log.Info(err)
	}

	r := mux.NewRouter()
	r.Handle("/rpc", rpcServer)

	return r
}
