package repository

import (
	"context"
	"fmt"
	"github.com/robbiekes/goods-manager-api/pkg/postgres"
)

type Repository struct {
	pg *postgres.Postgres
}

func NewRepo(pg *postgres.Postgres) *Repository {
	return &Repository{pg: pg}
}

func (repo *Repository) ItemsAmount(ctx context.Context, storageID int) (int, error) {
	sql := `
		select SUM(amount - reserved)
		from items_storages
		where storage_id = $1;
	`
	var items int

	rows := repo.pg.Pool.QueryRow(ctx, sql, storageID)

	err := rows.Scan(&items)
	if err != nil {
		return 0, fmt.Errorf("repo - ItemsAmount - rows.Scan: %w", err)
	}

	return items, nil
}

func (repo *Repository) ReserveItems(ctx context.Context, itemIDs []int64, storageID int) error {

	sql := `
	update items_storages
		set reserved = 
		case reserved < amount and amount > 0
			when true then (reserved + 1)
			else reserved
		end
	where item_id = any($1) and storage_id = $2;
	`
	_, err := repo.pg.Pool.Exec(ctx, sql, itemIDs, storageID)
	if err != nil {
		return fmt.Errorf("repo - ReserveItems - a.Pool.Exec: %w", err)
	}

	return nil
}

func (repo *Repository) CancelReservation(ctx context.Context, itemIDs []int64, storageID int) error {
	sql := `
		update items_storages
		set reserved = 
		case reserved > 0 
			when true then (reserved - 1)
			else reserved
		end
	where item_id = any($1) and storage_id = $2;
	`
	_, err := repo.pg.Pool.Exec(ctx, sql, itemIDs, storageID)
	if err != nil {
		return fmt.Errorf("repo - ReserveItems - a.Pool.Exec: %w", err)
	}

	return nil
}
