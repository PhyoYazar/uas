package fullmarkgrp

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/PhyoYazar/uas/business/core/fullmark"
	"github.com/PhyoYazar/uas/business/sys/validate"
	v1 "github.com/PhyoYazar/uas/business/web/v1"
	"github.com/PhyoYazar/uas/business/web/v1/paging"
	"github.com/PhyoYazar/uas/foundation/web"
	"github.com/google/uuid"
)

// Handlers manages the set of ga endpoints.
type Handlers struct {
	fullmark *fullmark.Core
}

// New constructs a handlers for route access.
func New(fullmark *fullmark.Core) *Handlers {
	return &Handlers{
		fullmark: fullmark,
	}
}

// Create adds a new ga to the system.
func (h *Handlers) Create(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var app AppNewFullMark
	if err := web.Decode(r, &app); err != nil {
		return err
	}

	cms, err := toCoreNewFullMark(app)
	if err != nil {
		return v1.NewRequestError(err, http.StatusBadRequest)
	}

	cm, err := h.fullmark.Create(ctx, cms)
	if err != nil {
		if errors.Is(err, fullmark.ErrUniqueFullMark) {
			return v1.NewRequestError(err, http.StatusConflict)
		}
		return fmt.Errorf("create: fullmark[%+v]: %w", cm, err)
	}

	return web.Respond(ctx, w, toAppFullMark(cm), http.StatusCreated)
}

// Delete removes a subject from the system.
func (h *Handlers) Delete(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	fullMarkID, err := uuid.Parse(web.Param(r, "full_mark_id"))
	if err != nil {
		return validate.NewFieldsError("full_mark_id", err)
	}

	if err := h.fullmark.Delete(ctx, fullMarkID.String()); err != nil {
		return fmt.Errorf("delete: fullMarkID[%s]: %w", fullMarkID, err)
	}

	return web.Respond(ctx, w, nil, http.StatusNoContent)
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

	marks, err := h.fullmark.Query(ctx, filter, orderBy, page.Number, page.RowsPerPage)
	if err != nil {
		return fmt.Errorf("query: %w", err)
	}

	items := make([]AppFullMark, len(marks))
	for i, cm := range marks {
		items[i] = toAppFullMark(cm)
	}

	total, err := h.fullmark.Count(ctx, filter)
	if err != nil {
		return fmt.Errorf("count: %w", err)
	}

	return web.Respond(ctx, w, paging.NewResponse(items, total, page.Number, page.RowsPerPage), http.StatusOK)
}
