package goods

type GoodsManager struct {
	ItemIDs   []int64 `json:"itemIDs" db:"itemIDs" binding:"required"`
	StorageID int     `json:"storageID" db:"storageID" binding:"required"`
}
