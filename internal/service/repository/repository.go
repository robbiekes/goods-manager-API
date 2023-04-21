package repository

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/robbiekes/goods-manager-api/pkg/postgres"
)

type Repository struct {
	pg *postgres.Postgres
}

func NewRepo(pg *postgres.Postgres) *Repository {
	return &Repository{pg: pg}
}

func (repo *Repository) ItemsAmount(ctx context.Context, storageID int) (int, error) {
	rowExists, err := repo.checkStorageExistence(ctx, storageID)
	if err != nil {
		return 0, errors.Wrap(err, "checking row existence")
	}

	if !rowExists {
		return 0, errors.New("no such storage")
	}

	sql := `
		select SUM(amount - reserved)
		from items_storages
		where storage_id = $1;
	`
	var items int

	rows := repo.pg.Pool.QueryRow(ctx, sql, storageID)

	err = rows.Scan(&items)
	if err != nil {
		return 0, fmt.Errorf("repo - ItemsAmount - rows.Scan: %w", err)
	}

	return items, nil
}

func (repo *Repository) ReserveItems(ctx context.Context, itemIDs []int64, storageID int) error {
	storageExists, err := repo.checkStorageExistence(ctx, storageID)
	if err != nil {
		return errors.Wrap(err, "checking row existance")
	}

	if !storageExists {
		return errors.New("no such storage")
	}

	sql := `
	update items_storages
		set reserved = 
		case reserved < amount and amount > 0
			when true then (reserved + 1)
			else reserved
		end
	where item_id = any($1) and storage_id = $2;
	`
	_, err = repo.pg.Pool.Exec(ctx, sql, itemIDs, storageID)
	if err != nil {
		return fmt.Errorf("repo - ReserveItems - a.Pool.Exec: %w", err)
	}

	return nil
}

func (repo *Repository) CancelReservation(ctx context.Context, itemIDs []int64, storageID int) error {
	storageExists, err := repo.checkStorageExistence(ctx, storageID)
	if err != nil {
		return errors.Wrap(err, "checking row existance")
	}

	if !storageExists {
		return errors.New("no such storage")
	}

	sql := `
		update items_storages
		set reserved = 
		case reserved > 0 
			when true then (reserved - 1)
			else reserved
		end
	where item_id = any($1) and storage_id = $2;
	`
	_, err = repo.pg.Pool.Exec(ctx, sql, itemIDs, storageID)
	if err != nil {
		return fmt.Errorf("repo - ReserveItems - a.Pool.Exec: %w", err)
	}

	return nil
}

func (repo *Repository) checkStorageExistence(ctx context.Context, id int) (bool, error) {
	var rowExists bool

	sqlCheck := `
		select exists (select true from storages where id = $1);
	`

	row := repo.pg.Pool.QueryRow(ctx, sqlCheck, id)
	err := row.Scan(&rowExists)
	if err != nil {
		return false, fmt.Errorf("repo - checkStorageExistence - rows.Scan: %w", err)
	}

	return rowExists, nil
}
