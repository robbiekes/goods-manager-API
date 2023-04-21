package service

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/robbiekes/goods-manager-api/internal/service/mock"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
	"testing"
)

var testError = errors.New("some error")

func TestGoodsManagerService_ItemsAmount(t *testing.T) {
	ctx := context.Background()

	type TestCases struct {
		Name         string
		StorageID    int
		Expectations func(r *mock.MockGoodsManagerRepo)
		Expected     int
		Error        assert.ErrorAssertionFunc
	}

	testcases := []TestCases{
		{
			Name: "OK",
			Expectations: func(r *mock.MockGoodsManagerRepo) {
				r.EXPECT().ItemsAmount(ctx, 1).Return(5, nil)
			},
			StorageID: 1,
			Expected:  5,
			Error:     assert.NoError,
		},
		{
			Name: "Error getting items amount",
			Expectations: func(r *mock.MockGoodsManagerRepo) {
				r.EXPECT().ItemsAmount(ctx, 1).Return(0, testError)
			},
			StorageID: 1,
			Expected:  0,
			Error:     assert.Error,
		},
	}
	for _, tc := range testcases {
		t.Run(tc.Name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			goodsManagerMock := mock.NewMockGoodsManagerRepo(ctrl)

			tc.Expectations(goodsManagerMock)

			amount, err := goodsManagerMock.ItemsAmount(ctx, tc.StorageID)
			if err != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tc.Expected, amount)

		})
	}
}

func TestGoodsManagerService_ReserveItems(t *testing.T) {
	ctx := context.Background()

	type TestCases struct {
		Name         string
		StorageID    int
		ItemIDs      []int64
		Expectations func(r *mock.MockGoodsManagerRepo)
		Error        assert.ErrorAssertionFunc
	}

	items := []int64{1, 2}

	testcases := []TestCases{
		{
			Name: "OK",
			Expectations: func(r *mock.MockGoodsManagerRepo) {
				r.EXPECT().ReserveItems(ctx, items, 1).Return(nil)
			},
			StorageID: 1,
			ItemIDs:   items,
			Error:     assert.NoError,
		},
		{
			Name: "Error empty items list",
			Expectations: func(r *mock.MockGoodsManagerRepo) {
				r.EXPECT().ReserveItems(ctx, []int64{}, 1).Return(ErrEmptyItemsList)
			},
			StorageID: 1,
			ItemIDs:   []int64{},
			Error:     assert.Error,
		},
		{
			Name: "Error reserving items",
			Expectations: func(r *mock.MockGoodsManagerRepo) {
				r.EXPECT().ReserveItems(ctx, items, 1).Return(testError)
			},
			StorageID: 1,
			ItemIDs:   items,
			Error:     assert.Error,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.Name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			goodsManagerMock := mock.NewMockGoodsManagerRepo(ctrl)

			tc.Expectations(goodsManagerMock)

			err := goodsManagerMock.ReserveItems(ctx, tc.ItemIDs, tc.StorageID)
			if err != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestGoodsManagerService_CancelReservation(t *testing.T) {
	ctx := context.Background()

	type TestCases struct {
		Name         string
		StorageID    int
		ItemIDs      []int64
		Expectations func(r *mock.MockGoodsManagerRepo)
		Error        assert.ErrorAssertionFunc
	}

	items := []int64{1, 2}

	testcases := []TestCases{
		{
			Name: "OK",
			Expectations: func(r *mock.MockGoodsManagerRepo) {
				r.EXPECT().CancelReservation(ctx, items, 1).Return(nil)
			},
			StorageID: 1,
			ItemIDs:   items,
			Error:     assert.NoError,
		},
		{
			Name: "Error empty items list",
			Expectations: func(r *mock.MockGoodsManagerRepo) {
				r.EXPECT().CancelReservation(ctx, []int64{}, 1).Return(ErrEmptyItemsList)
			},
			StorageID: 1,
			ItemIDs:   []int64{},
			Error:     assert.Error,
		},
		{
			Name: "Error cancelling items reservation",
			Expectations: func(r *mock.MockGoodsManagerRepo) {
				r.EXPECT().CancelReservation(ctx, items, 1).Return(testError)
			},
			StorageID: 1,
			ItemIDs:   items,
			Error:     assert.Error,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.Name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			goodsManagerMock := mock.NewMockGoodsManagerRepo(ctrl)

			tc.Expectations(goodsManagerMock)

			err := goodsManagerMock.CancelReservation(ctx, tc.ItemIDs, tc.StorageID)
			if err != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
