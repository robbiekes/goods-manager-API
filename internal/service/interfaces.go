package service

import (
	"context"
)

type GoodsManagerRepo interface {
	// получение кол-ва оставшихся товаров на складе
	ItemsAmount(ctx context.Context, storageID int) (int, error)
	// резервирование товара на складе для доставки
	ReserveItems(ctx context.Context, itemIDs []int64, storageID int) error
	// освобождение резерва товаров
	CancelReservation(ctx context.Context, itemIDs []int64, storageID int) error
}
