package service

import (
	"context"
	"github.com/pkg/errors"
	"github.com/robbiekes/goods-manager-api/internal/entity"
)

var ErrEmptyItemsList = errors.New("empty items list")

type GoodsManagerService struct {
	repo GoodsManagerRepo
}

func NewService(repo GoodsManagerRepo) *GoodsManagerService {
	return &GoodsManagerService{repo: repo}
}

func (s *GoodsManagerService) ItemsList(ctx context.Context, storageID int) ([]entity.Item, error) {
	items, err := s.repo.ItemsList(ctx, storageID)
	if err != nil {
		return nil, errors.Wrap(err, "getting items")
	}

	if items == nil {
		return nil, errors.Wrap(err, "empty items list")
	}

	return items, nil
}

func (s *GoodsManagerService) ReserveItems(ctx context.Context, itemIDs []int64, storageID int) error {
	if itemIDs == nil && len(itemIDs) == 0 {
		return ErrEmptyItemsList
	}

	err := s.repo.ReserveItems(ctx, itemIDs, storageID)
	if err != nil {
		return errors.Wrap(err, "reserving items")
	}

	return nil
}

func (s *GoodsManagerService) CancelReservation(ctx context.Context, itemIDs []int64, storageID int) error {
	if itemIDs == nil && len(itemIDs) == 0 {
		return ErrEmptyItemsList
	}

	err := s.repo.CancelReservation(ctx, itemIDs, storageID)
	if err != nil {
		return errors.Wrap(err, "cancelling items reservation")
	}

	return nil
}
