package service

import (
	"context"
	"github.com/robbiekes/goods-manager-api/internal/entity"
)

type GoodsManagerRepo interface {
	// получение кол-ва оставшихся товаров на складе // TODO: rename to GetItemsByStorage
	ItemsList(ctx context.Context, storageID int) ([]entity.Item, error)
	// резервирование товара на складе для доставки
	ReserveItems(ctx context.Context, itemIDs []int64, storageID int) error
	// освобождение резерва товаров
	CancelReservation(ctx context.Context, itemIDs []int64, storageID int) error
}
