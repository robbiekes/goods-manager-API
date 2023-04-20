package jsonrpc

import (
	"context"
	"github.com/pkg/errors"
	"net/http"
)

func (r *Router) ReserveItems(_ *http.Request, req *ReserveItemsInput, resp *ReserveItemsOutput) error {
	ctx := context.Background()

	if req == nil {
		return errors.New("empty input")
	}

	err := r.service.ReserveItems(ctx, req.ItemIDs, req.StorageID)
	if err != nil {
		return errors.Wrap(err, "reserving items")
	}

	return nil
}

func (r *Router) CancelReservation(_ *http.Request, req *CancelReservationInput, resp *CancelReservationOutput) error {
	ctx := context.Background()

	if req == nil {
		return errors.New("empty input")
	}

	err := r.service.CancelReservation(ctx, req.ItemIDs, req.StorageID)
	if err != nil {
		return errors.Wrap(err, "cancelling items reservation")
	}

	return nil
}

func (r *Router) ItemsList(_ *http.Request, req *ItemsListInput, resp *ItemsListOutput) error {
	ctx := context.Background()

	if req == nil {
		return errors.New("empty input")
	}

	items, err := r.service.ItemsList(ctx, req.StorageID)
	if err != nil {
		return errors.Wrap(err, "getting items")
	}

	resp.Amount = len(items)

	return nil
}
