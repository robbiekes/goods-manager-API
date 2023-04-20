package jsonrpc

import (
	json2 "encoding/json"
	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
	"github.com/robbiekes/goods-manager-api/internal/entity/jsonrpc"
	"github.com/robbiekes/goods-manager-api/internal/service"
	log "github.com/sirupsen/logrus"
	"net/http"
)

const Version = "2.0"

type Router struct {
	service *service.GoodsManagerService
}

func NewRpcRouter(rpcServer *rpc.Server, s *service.GoodsManagerService) {
	rpcServer.RegisterCodec(json.NewCodec(), "application/json")
	rpcServer.RegisterCodec(json.NewCodec(), "application/json;charset=UTF-8")

	router := &Router{service: s}

	err := rpcServer.RegisterService(router, "router")
	if err != nil {
		log.Info(err)
	}

	r := mux.NewRouter()
	r.Handle("/rpc", rpcServer)
}

func (r *Router) ReserveItems(_ *http.Request, req *jsonrpc.JsonRpcRequest, resp *jsonrpc.JsonRpcResponse) error {
	var params jsonrpc.RequestParams

	err := ValidateRequest(req)
	if err != nil {
		resp = &jsonrpc.JsonRpcResponse{
			Version: req.Version,
			Error: &jsonrpc.JsonRpcError{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			},
			ID:     req.ID,
			Result: nil,
		}
		return err
	}

	log.Info(req)

	err = json2.Unmarshal(req.Params, &params)
	if err != nil {
		resp = &jsonrpc.JsonRpcResponse{
			Version: req.Version,
			Error: &jsonrpc.JsonRpcError{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			},
			ID:     req.ID,
			Result: nil,
		}

		return err
	}

	// call method
	if req.Method == "router.ReserveItems" {
		err = r.service.ReserveItems(params.ItemIDs, params.StorageID)
		if err != nil {
			resp = &jsonrpc.JsonRpcResponse{
				Version: req.Version,
				Error: &jsonrpc.JsonRpcError{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				},
				ID:     req.ID,
				Result: nil,
			}
			return err
		}
	} else {
		resp = &jsonrpc.JsonRpcResponse{
			Version: req.Version,
			Error: &jsonrpc.JsonRpcError{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			},
			ID:     req.ID,
			Result: nil,
		}
		return err
	}

	resp = &jsonrpc.JsonRpcResponse{
		Version: req.Version,
		Error:   nil,
		ID:      req.ID,
		Result:  nil,
	}

	return nil
}

func (r *Router) CancelReservation(_ *http.Request, req *jsonrpc.JsonRpcRequest, resp *jsonrpc.JsonRpcResponse) error {
	var params jsonrpc.RequestParams

	err := json2.Unmarshal(req.Params, &params)
	if err != nil {
		resp = &jsonrpc.JsonRpcResponse{
			Version: req.Version,
			Error: &jsonrpc.JsonRpcError{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			},
			ID:     req.ID,
			Result: nil,
		}

		return err
	}

	// call method
	if req.Method == "router.CancelReservation" {
		err = r.service.CancelReservation(params.ItemIDs, params.StorageID)
		if err != nil {
			resp = &jsonrpc.JsonRpcResponse{
				Version: req.Version,
				Error: &jsonrpc.JsonRpcError{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				},
				ID:     req.ID,
				Result: nil,
			}

			return err
		}
	} else {
		resp = &jsonrpc.JsonRpcResponse{
			Version: req.Version,
			Error: &jsonrpc.JsonRpcError{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			},
			ID:     req.ID,
			Result: nil,
		}

		return err
	}

	resp = &jsonrpc.JsonRpcResponse{
		Version: req.Version,
		Error:   nil,
		ID:      req.ID,
		Result:  nil,
	}

	return nil
}

func (r *Router) ItemsList(httpReq *http.Request, req *jsonrpc.JsonRpcRequest, resp *jsonrpc.JsonRpcResponse) error {
	var (
		params jsonrpc.RequestParams
		result int
	)

	log.Info(httpReq)

	err := json2.Unmarshal(req.Params, &params)
	if err != nil {
		resp = &jsonrpc.JsonRpcResponse{
			Version: req.Version,
			Error: &jsonrpc.JsonRpcError{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			},
			ID:     req.ID,
			Result: nil,
		}

		return err
	}

	// call method
	if req.Method == "router.ItemsList" {
		result, err = r.service.ItemsList(params.StorageID)
		if err != nil {
			resp = &jsonrpc.JsonRpcResponse{
				Version: req.Version,
				Error: &jsonrpc.JsonRpcError{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				},
				ID:     req.ID,
				Result: nil,
			}

			return err
		}
	} else {
		resp = &jsonrpc.JsonRpcResponse{
			Version: req.Version,
			Error: &jsonrpc.JsonRpcError{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			},
			ID:     req.ID,
			Result: nil,
		}

		return err
	}

	// marshalling request
	resultBytes, err := json2.Marshal(result)
	if err != nil {
		resp = &jsonrpc.JsonRpcResponse{
			Version: req.Version,
			Error: &jsonrpc.JsonRpcError{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			},
			ID:     req.ID,
			Result: nil,
		}

		return err
	}

	resp = &jsonrpc.JsonRpcResponse{
		Version: req.Version,
		Error:   nil,
		ID:      req.ID,
		Result:  resultBytes,
	}

	return nil
}
