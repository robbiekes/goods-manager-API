package jsonrpc

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/robbiekes/goods-manager-api/internal/entity/jsonrpc"
)

var (
	ErrMalformedVersion = errors.New("malformed JSON-RPC version")
	ErrEmptyParams      = errors.New("empty params")
	ErrEmptyMethod      = errors.New("empty method")
	ErrEmptyID          = errors.New("empty id")
)

func ValidateRequest(req *jsonrpc.JsonRpcRequest) error {
	switch {
	case req.Version != jsonrpc.Version:
		return fmt.Errorf("%s: %s", ErrMalformedVersion, req.Version)
	case req.Params == nil:
		return ErrEmptyParams
	case req.Method == "":
		return ErrEmptyMethod
	case req.ID == "":
		return ErrEmptyID
	default:
		return nil
	}
}
