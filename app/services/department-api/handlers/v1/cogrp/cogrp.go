package cogrp

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/PhyoYazar/uas/business/core/co"
	v1 "github.com/PhyoYazar/uas/business/web/v1"
	"github.com/PhyoYazar/uas/business/web/v1/paging"
	"github.com/PhyoYazar/uas/foundation/web"
)

// Handlers manages the set of co endpoints.
type Handlers struct {
	co *co.Core
}

// New constructs a handlers for route access.
func New(co *co.Core) *Handlers {
	return &Handlers{
		co: co,
	}
}

// Create adds a new co to the system.
func (h *Handlers) Create(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var app AppNewCo
	if err := web.Decode(r, &app); err != nil {
		return err
	}

	nc, err := toCoreNewCo(app)
	if err != nil {
		return v1.NewRequestError(err, http.StatusBadRequest)
	}

	c, err := h.co.Create(ctx, nc)
	if err != nil {
		if errors.Is(err, co.ErrUniqueCo) {
			return v1.NewRequestError(err, http.StatusConflict)
		}
		return fmt.Errorf("create: c[%+v]: %w", c, err)
	}

	return web.Respond(ctx, w, toAppCo(c), http.StatusCreated)
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

	cos, err := h.co.Query(ctx, filter, orderBy, page.Number, page.RowsPerPage)
	if err != nil {
		return fmt.Errorf("query: %w", err)
	}

	items := make([]AppCo, len(cos))
	for i, co := range cos {
		items[i] = toAppCo(co)
	}

	total, err := h.co.Count(ctx, filter)
	if err != nil {
		return fmt.Errorf("count: %w", err)
	}

	return web.Respond(ctx, w, paging.NewResponse(items, total, page.Number, page.RowsPerPage), http.StatusOK)
}
