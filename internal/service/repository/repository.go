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

func (repo *Repository) ItemsList(ctx context.Context, storageID int) (int, error) {
	sql, args, err := repo.pg.Builder.
		Select("itemID").
		From("items").
		Where("storageID = ? AND reserved = false", storageID).
		ToSql()

	if err != nil {
		return 0, fmt.Errorf("repo - ItemsList - a.Builder: %w", err)
	}

	var itemsLeft []int

	rows, err := repo.pg.Pool.Query(ctx, sql, args...)
	if err != nil {
		return 0, fmt.Errorf("repo - ItemsList - a.Pool.QueryRow: %w", err)
	}

	defer rows.Close()

	err = rows.Scan(&itemsLeft)
	if err != nil {
		return 0, fmt.Errorf("repo - ItemsList - rows.Scan: %w", err)
	}

	return len(itemsLeft), nil
}

func (repo *Repository) ReserveItems(ctx context.Context, itemIDs []int64, storageID int) error {
	// if item doesn't exist in storage, skip it or if it's single in the request, return error
	sql, args, err := repo.pg.Builder.
		Update("items").
		Set("reserved", "true").
		Where("itemID in $1 AND storageID = $2", itemIDs, storageID).
		ToSql()
	if err != nil {
		return fmt.Errorf("repo - ReserveItems - a.Builder: %w", err)
	}

	_, err = repo.pg.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("repo - ReserveItems - a.Pool.Exec: %w", err)
	}

	return nil
}

func (repo *Repository) CancelReservation(ctx context.Context, itemIDs []int64, storageID int) error {
	sql, args, err := repo.pg.Builder.
		Update("items").
		Set("reserved", "false").
		Where("itemID in ? AND storageID = ?", itemIDs, storageID).
		ToSql()
	if err != nil {
		return fmt.Errorf("repo - CancelReservation - a.Builder: %w", err)
	}

	_, err = repo.pg.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("repo - CancelReservation - a.Pool.Exec: %w", err)
	}

	return nil
}
