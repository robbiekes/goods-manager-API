package jsonrpc

type ReserveItemsInput struct {
	ItemIDs []int64 `json:"item_ids,omitempty"`
	// StorageID нужен для работы с товарами, одновременно хранящимися на нескольких складах
	StorageID int `json:"storage_id"`
}

type ReserveItemsOutput struct {
}

type CancelReservationInput struct {
	ItemIDs []int64 `json:"item_ids,omitempty"`
	// StorageID нужен для работы с товарами, одновременно хранящимися на нескольких складах
	StorageID int `json:"storage_id"`
}

type CancelReservationOutput struct {
}

type ItemsListInput struct {
	StorageID int `json:"storage_id"`
}

type ItemsListOutput struct {
	Amount int `json:"amount"`
}
