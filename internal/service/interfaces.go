package service

type GoodsManagerRepo interface {
	// получение кол-ва оставшихся товаров на складе
	ItemsList(storageID int) (int, error)
	// резервирование товара на складе для доставки
	ReserveItems(itemIDs []int64, storageID int) error
	// освобождение резерва товаров
	CancelReservation(itemIDs []int64, storageID int) error
}
