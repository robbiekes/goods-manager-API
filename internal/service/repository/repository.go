package repository

import (
	"context"
	"fmt"
	"github.com/robbiekes/goods-manager-api/internal/entity"
	"github.com/robbiekes/goods-manager-api/pkg/postgres"
)

type Repository struct {
	pg *postgres.Postgres
}

func NewRepo(pg *postgres.Postgres) *Repository {
	return &Repository{pg: pg}
}

func (repo *Repository) ItemsList(ctx context.Context, storageID int) ([]entity.Item, error) {
	sql := `
		select i.id, i.name, i.size
		from items as i
		join items_storages as i_s
		on i.id = i_s.item_id
		where i_s.storage_id = $1 and i_s.reserved = false
	`
	var items []entity.Item

	rows, err := repo.pg.Pool.Query(ctx, sql, storageID)
	if err != nil {
		return nil, fmt.Errorf("repo - ItemsList - a.Pool.QueryRow: %w", err)
	}

	defer rows.Close()

	err = rows.Scan(&items)
	if err != nil {
		return nil, fmt.Errorf("repo - ItemsList - rows.Scan: %w", err)
	}

	return items, nil
}

func (repo *Repository) ReserveItems(ctx context.Context, itemIDs []int64, storageID int) error {
	sql := `
		update items_storages
		set reserved = true
		where item_id in $1 and storage_id = $2
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
		set reserved = false
		where item_id in $1 and storage_id = $2
	`
	_, err := repo.pg.Pool.Exec(ctx, sql, itemIDs, storageID)
	if err != nil {
		return fmt.Errorf("repo - ReserveItems - a.Pool.Exec: %w", err)
	}

	return nil
}
