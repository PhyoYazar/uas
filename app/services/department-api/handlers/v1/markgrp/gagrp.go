package markgrp

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/PhyoYazar/uas/business/core/mark"
	v1 "github.com/PhyoYazar/uas/business/web/v1"
	"github.com/PhyoYazar/uas/business/web/v1/paging"
	"github.com/PhyoYazar/uas/foundation/web"
)

// Handlers manages the set of ga endpoints.
type Handlers struct {
	mark *mark.Core
}

// New constructs a handlers for route access.
func New(mark *mark.Core) *Handlers {
	return &Handlers{
		mark: mark,
	}
}

// Create adds a new ga to the system.
func (h *Handlers) Create(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var app AppNewMark
	if err := web.Decode(r, &app); err != nil {
		return err
	}

	nm, err := toCoreNewMark(app)
	if err != nil {
		return v1.NewRequestError(err, http.StatusBadRequest)
	}

	mk, err := h.mark.Create(ctx, nm)
	if err != nil {
		if errors.Is(err, mark.ErrUniqueMark) {
			return v1.NewRequestError(err, http.StatusConflict)
		}
		return fmt.Errorf("create: ga[%+v]: %w", mk, err)
	}

	return web.Respond(ctx, w, toAppMark(mk), http.StatusCreated)
}

// Query returns a list of cos with paging.
func (h *Handlers) Query(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

	page, err := paging.ParseRequest(r)
	if err != nil {
		return err
	}

	filter, err := parseFilter(r)
	if err != nil {
		return err
	}

	orderBy, err := parseOrder(r)
	if err != nil {
		return err
	}

	mks, err := h.mark.Query(ctx, filter, orderBy, page.Number, page.RowsPerPage)
	if err != nil {
		return fmt.Errorf("query: %w", err)
	}

	items := make([]AppMark, len(mks))
	for i, mk := range mks {
		items[i] = toAppMark(mk)
	}

	total, err := h.mark.Count(ctx, filter)
	if err != nil {
		return fmt.Errorf("count: %w", err)
	}

	return web.Respond(ctx, w, paging.NewResponse(items, total, page.Number, page.RowsPerPage), http.StatusOK)
}
