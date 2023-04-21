package service

import (
	"context"
	"github.com/pkg/errors"
)

var ErrEmptyItemsList = errors.New("empty items list")

type GoodsManagerService struct {
	repo GoodsManagerRepo
}

func NewService(repo GoodsManagerRepo) *GoodsManagerService {
	return &GoodsManagerService{repo: repo}
}

func (s *GoodsManagerService) ItemsAmount(ctx context.Context, storageID int) (int, error) {
	items, err := s.repo.ItemsAmount(ctx, storageID)
	if err != nil {
		return 0, errors.Wrap(err, "getting items")
	}

	// if items == 0 {
	// 	return 0, errors.Wrap(err, "empty items list")
	// }

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
