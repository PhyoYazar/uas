package gagrp

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/PhyoYazar/uas/business/core/ga"
	v1 "github.com/PhyoYazar/uas/business/web/v1"
	"github.com/PhyoYazar/uas/business/web/v1/paging"
	"github.com/PhyoYazar/uas/foundation/web"
)

// Handlers manages the set of ga endpoints.
type Handlers struct {
	ga *ga.Core
}

// New constructs a handlers for route access.
func New(ga *ga.Core) *Handlers {
	return &Handlers{
		ga: ga,
	}
}

// Create adds a new ga to the system.
func (h *Handlers) Create(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var app AppNewGa
	if err := web.Decode(r, &app); err != nil {
		return err
	}

	ng, err := toCoreNewGa(app)
	if err != nil {
		return v1.NewRequestError(err, http.StatusBadRequest)
	}

	g, err := h.ga.Create(ctx, ng)
	if err != nil {
		if errors.Is(err, ga.ErrUniqueGa) {
			return v1.NewRequestError(err, http.StatusConflict)
		}
		return fmt.Errorf("create: ga[%+v]: %w", g, err)
	}

	return web.Respond(ctx, w, toAppGa(g), http.StatusCreated)
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

	gas, err := h.ga.Query(ctx, filter, orderBy, page.Number, page.RowsPerPage)
	if err != nil {
		return fmt.Errorf("query: %w", err)
	}

	items := make([]AppGa, len(gas))
	for i, ga := range gas {
		items[i] = toAppGa(ga)
	}

	total, err := h.ga.Count(ctx, filter)
	if err != nil {
		return fmt.Errorf("count: %w", err)
	}

	return web.Respond(ctx, w, paging.NewResponse(items, total, page.Number, page.RowsPerPage), http.StatusOK)
}
