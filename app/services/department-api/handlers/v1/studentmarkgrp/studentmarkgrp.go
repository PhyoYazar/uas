package studentmarkgrp

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/PhyoYazar/uas/business/core/studentmark"
	"github.com/PhyoYazar/uas/business/sys/validate"
	v1 "github.com/PhyoYazar/uas/business/web/v1"
	"github.com/PhyoYazar/uas/business/web/v1/paging"
	"github.com/PhyoYazar/uas/foundation/web"
	"github.com/google/uuid"
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

// Update updates a student in the system.
func (h *Handlers) Update(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var app AppUpdateStudentMark
	if err := web.Decode(r, &app); err != nil {
		return err
	}

	studentMarkID, err := uuid.Parse(web.Param(r, "student_mark_id"))
	if err != nil {
		return validate.NewFieldsError("student_mark_id", err)
	}

	std, err := h.studentMark.QueryByID(ctx, studentMarkID)
	if err != nil {
		switch {
		case errors.Is(err, studentmark.ErrNotFound):
			return v1.NewRequestError(err, http.StatusNotFound)
		default:
			return fmt.Errorf("querybyid: studentMarkID[%s]: %w", studentMarkID, err)
		}
	}

	uSub, err := toCoreUpdateStudentMark(app)
	if err != nil {
		return v1.NewRequestError(err, http.StatusBadRequest)
	}

	std, err = h.studentMark.Update(ctx, std, uSub)
	if err != nil {
		return fmt.Errorf("update: studentMarkID[%s] uSub[%+v]: %w", studentMarkID, uSub, err)
	}

	return web.Respond(ctx, w, toAppStudentMark(std), http.StatusOK)
}

// Delete removes a student from the system.
func (h *Handlers) Delete(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	studentMarkID, err := uuid.Parse(web.Param(r, "student_mark_id"))
	if err != nil {
		return validate.NewFieldsError("student_mark_id", err)
	}

	std, err := h.studentMark.QueryByID(ctx, studentMarkID)
	if err != nil {
		switch {
		case errors.Is(err, studentmark.ErrNotFound):
			return web.Respond(ctx, w, nil, http.StatusNoContent)
		default:
			return fmt.Errorf("querybyid: studentMarkID[%s]: %w", studentMarkID, err)
		}
	}

	if err := h.studentMark.Delete(ctx, std); err != nil {
		return fmt.Errorf("delete: studentMarkID[%s]: %w", studentMarkID, err)
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

// QueryByID returns a student by its ID.
func (h *Handlers) QueryByID(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	studentMarkID, err := uuid.Parse(web.Param(r, "student_mark_id"))
	if err != nil {
		return validate.NewFieldsError("student_mark_id", err)
	}

	std, err := h.studentMark.QueryByID(ctx, studentMarkID)
	if err != nil {
		switch {
		case errors.Is(err, studentmark.ErrNotFound):
			return v1.NewRequestError(err, http.StatusNotFound)
		default:
			return fmt.Errorf("querybyid: studentMarkID[%s]: %w", studentMarkID, err)
		}
	}

	return web.Respond(ctx, w, toAppStudentMark(std), http.StatusOK)
}
