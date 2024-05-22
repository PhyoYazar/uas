package studentsubjectgrp

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/PhyoYazar/uas/business/core/studentsubject"
	v1 "github.com/PhyoYazar/uas/business/web/v1"
	"github.com/PhyoYazar/uas/business/web/v1/paging"
	"github.com/PhyoYazar/uas/foundation/web"
)

// Handlers manages the set of ga endpoints.
type Handlers struct {
	studentSubject *studentsubject.Core
}

// New constructs a handlers for route access.
func New(studentSubject *studentsubject.Core) *Handlers {
	return &Handlers{
		studentSubject: studentSubject,
	}
}

// Create adds a new ga to the system.
func (h *Handlers) Create(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var app AppNewStudentSubject
	if err := web.Decode(r, &app); err != nil {
		return err
	}

	nss, err := toCoreNewStudentSubject(app)
	if err != nil {
		return v1.NewRequestError(err, http.StatusBadRequest)
	}

	stdsub, err := h.studentSubject.Create(ctx, nss)
	if err != nil {
		if errors.Is(err, studentsubject.ErrUniqueStudentSubject) {
			return v1.NewRequestError(err, http.StatusConflict)
		}
		return fmt.Errorf("create: ga[%+v]: %w", stdsub, err)
	}

	return web.Respond(ctx, w, toAppStudentSubject(stdsub), http.StatusCreated)
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

	ss, err := h.studentSubject.Query(ctx, filter, orderBy, page.Number, page.RowsPerPage)
	if err != nil {
		return fmt.Errorf("query: %w", err)
	}

	items := make([]AppStudentSubject, len(ss))
	for i, stdsub := range ss {
		items[i] = toAppStudentSubject(stdsub)
	}

	total, err := h.studentSubject.Count(ctx, filter)
	if err != nil {
		return fmt.Errorf("count: %w", err)
	}

	return web.Respond(ctx, w, paging.NewResponse(items, total, page.Number, page.RowsPerPage), http.StatusOK)
}