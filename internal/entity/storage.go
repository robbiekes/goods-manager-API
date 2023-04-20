package entity

type Storage struct {
	ID      int    `db:"id"`
	Name    string `db:"name"`
	Allowed bool   `db:"allowed"`
}
