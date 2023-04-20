package service

import "github.com/pkg/errors"

var ErrEmptyItemsList = errors.New("empty items list")

type GoodsManagerService struct {
	repo GoodsManagerRepo
}

func NewService(repo GoodsManagerRepo) *GoodsManagerService {
	return &GoodsManagerService{repo: repo}
}

func (s *GoodsManagerService) ItemsList(storageID int) (int, error) {
	return s.repo.ItemsList(storageID)
}

func (s *GoodsManagerService) ReserveItems(itemIDs []int64, storageID int) error {
	if itemIDs == nil && len(itemIDs) == 0 {
		return ErrEmptyItemsList
	}

	return s.repo.ReserveItems(itemIDs, storageID)
}

func (s *GoodsManagerService) CancelReservation(itemIDs []int64, storageID int) error {
	if itemIDs == nil && len(itemIDs) == 0 {
		return ErrEmptyItemsList
	}

	return s.repo.CancelReservation(itemIDs, storageID)
}
