package studentmarkgrp

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/PhyoYazar/uas/business/core/studentmark"
	v1 "github.com/PhyoYazar/uas/business/web/v1"
	"github.com/PhyoYazar/uas/business/web/v1/paging"
	"github.com/PhyoYazar/uas/foundation/web"
)

// Handlers manages the set of ga endpoints.
type Handlers struct {
	studentMark *studentmark.Core
}

// New constructs a handlers for route access.
func New(studentMark *studentmark.Core) *Handlers {
	return &Handlers{
		studentMark: studentMark,
	}
}

// Create adds a new ga to the system.
func (h *Handlers) Create(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var app AppNewStudentMark
	if err := web.Decode(r, &app); err != nil {
		return err
	}

	nss, err := toCoreNewStudentMark(app)
	if err != nil {
		return v1.NewRequestError(err, http.StatusBadRequest)
	}

	stdsub, err := h.studentMark.Create(ctx, nss)
	if err != nil {
		if errors.Is(err, studentmark.ErrUniqueStudentMark) {
			return v1.NewRequestError(err, http.StatusConflict)
		}
		return fmt.Errorf("create: ga[%+v]: %w", stdsub, err)
	}

	return web.Respond(ctx, w, toAppStudentMark(stdsub), http.StatusCreated)
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

	ss, err := h.studentMark.Query(ctx, filter, orderBy, page.Number, page.RowsPerPage)
	if err != nil {
		return fmt.Errorf("query: %w", err)
	}

	items := make([]AppStudentMark, len(ss))
	for i, stdsub := range ss {
		items[i] = toAppStudentMark(stdsub)
	}

	total, err := h.studentMark.Count(ctx, filter)
	if err != nil {
		return fmt.Errorf("count: %w", err)
	}

	return web.Respond(ctx, w, paging.NewResponse(items, total, page.Number, page.RowsPerPage), http.StatusOK)
}
