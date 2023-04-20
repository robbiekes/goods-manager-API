package entity

type Item struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
	Size string `db:"size"`
}
