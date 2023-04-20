package goods

type GoodsGetter struct {
	StorageID int `json:"storageID" db:"storageID" binding:"required"`
}
