package studentgrp

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/PhyoYazar/uas/business/core/student"
	v1 "github.com/PhyoYazar/uas/business/web/v1"
	"github.com/PhyoYazar/uas/business/web/v1/paging"
	"github.com/PhyoYazar/uas/foundation/web"
)

// Handlers manages the set of user endpoints.
type Handlers struct {
	student *student.Core
}

// New constructs a handlers for route access.
func New(student *student.Core) *Handlers {
	return &Handlers{
		student: student,
	}
}

// Create adds a new student to the system.
func (h *Handlers) Create(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var app AppNewStudent
	if err := web.Decode(r, &app); err != nil {
		return err
	}

	ns, err := toCoreNewStudent(app)
	if err != nil {
		return v1.NewRequestError(err, http.StatusBadRequest)
	}

	std, err := h.student.Create(ctx, ns)
	if err != nil {
		if errors.Is(err, student.ErrUniqueEmail) {
			return v1.NewRequestError(err, http.StatusConflict)
		}
		return fmt.Errorf("create: std[%+v]: %w", std, err)
	}

	return web.Respond(ctx, w, toAppStudent(std), http.StatusCreated)
}

// Query returns a list of subjects with paging.
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

	students, err := h.student.Query(ctx, filter, orderBy, page.Number, page.RowsPerPage)
	if err != nil {
		return fmt.Errorf("query: %w", err)
	}

	items := make([]AppStudent, len(students))
	for i, std := range students {
		items[i] = toAppStudent(std)
	}

	total, err := h.student.Count(ctx, filter)
	if err != nil {
		return fmt.Errorf("count: %w", err)
	}

	return web.Respond(ctx, w, paging.NewResponse(items, total, page.Number, page.RowsPerPage), http.StatusOK)
}
