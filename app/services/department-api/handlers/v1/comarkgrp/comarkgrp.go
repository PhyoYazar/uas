package comarkgrp

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/PhyoYazar/uas/business/core/comark"
	v1 "github.com/PhyoYazar/uas/business/web/v1"
	"github.com/PhyoYazar/uas/business/web/v1/paging"
	"github.com/PhyoYazar/uas/foundation/web"
)

// Handlers manages the set of ga endpoints.
type Handlers struct {
	comark *comark.Core
}

// New constructs a handlers for route access.
func New(comark *comark.Core) *Handlers {
	return &Handlers{
		comark: comark,
	}
}

// Create adds a new ga to the system.
func (h *Handlers) Create(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var app AppNewCoMark
	if err := web.Decode(r, &app); err != nil {
		return err
	}

	cms, err := toCoreNewCoMark(app)
	if err != nil {
		return v1.NewRequestError(err, http.StatusBadRequest)
	}

	cm, err := h.comark.Create(ctx, cms)
	if err != nil {
		if errors.Is(err, comark.ErrUniqueCoMark) {
			return v1.NewRequestError(err, http.StatusConflict)
		}
		return fmt.Errorf("create: ga[%+v]: %w", cm, err)
	}

	return web.Respond(ctx, w, toAppCoMark(cm), http.StatusCreated)
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

	comarks, err := h.comark.Query(ctx, filter, orderBy, page.Number, page.RowsPerPage)
	if err != nil {
		return fmt.Errorf("query: %w", err)
	}

	items := make([]AppCoMark, len(comarks))
	for i, cm := range comarks {
		items[i] = toAppCoMark(cm)
	}

	total, err := h.comark.Count(ctx, filter)
	if err != nil {
		return fmt.Errorf("count: %w", err)
	}

	return web.Respond(ctx, w, paging.NewResponse(items, total, page.Number, page.RowsPerPage), http.StatusOK)
}
