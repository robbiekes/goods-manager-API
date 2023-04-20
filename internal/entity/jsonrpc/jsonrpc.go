package jsonrpc

import "encoding/json"

const Version = "2.0"

// TODO: create request structure for get request

type RequestParams struct {
	ItemIDs []int64 `json:"itemIDs,omitempty"`

	// Для работы с товарами, которые одновременно могут находиться на нескольких складах
	StorageID int `json:"storageID"`
}

type JsonRpcRequest struct {
	Version string          `json:"jsonrpc"`
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params"`
	ID      string          `json:"id"`
}

type JsonRpcResponse struct {
	Version string          `json:"jsonrpc"`
	Error   *JsonRpcError   `json:"error,omitempty"`
	ID      string          `json:"id"`
	Result  json.RawMessage `json:"result,omitempty"`
}

type JsonRpcError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
