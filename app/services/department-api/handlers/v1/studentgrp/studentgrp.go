package studentgrp

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/PhyoYazar/uas/business/core/student"
	"github.com/PhyoYazar/uas/business/sys/validate"
	v1 "github.com/PhyoYazar/uas/business/web/v1"
	"github.com/PhyoYazar/uas/business/web/v1/paging"
	"github.com/PhyoYazar/uas/foundation/web"
	"github.com/google/uuid"
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
		if errors.Is(err, student.ErrUniqueStudent) {
			return v1.NewRequestError(err, http.StatusConflict)
		}
		return fmt.Errorf("create: std[%+v]: %w", std, err)
	}

	return web.Respond(ctx, w, toAppStudent(std), http.StatusCreated)
}

// Update updates a student in the system.
func (h *Handlers) Update(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var app AppUpdateStudent
	if err := web.Decode(r, &app); err != nil {
		return err
	}

	studentID, err := uuid.Parse(web.Param(r, "student_id"))
	if err != nil {
		return validate.NewFieldsError("student_id", err)
	}

	std, err := h.student.QueryByID(ctx, studentID)
	if err != nil {
		switch {
		case errors.Is(err, student.ErrNotFound):
			return v1.NewRequestError(err, http.StatusNotFound)
		default:
			return fmt.Errorf("querybyid: studentID[%s]: %w", studentID, err)
		}
	}

	uSub, err := toCoreUpdateStudent(app)
	if err != nil {
		return v1.NewRequestError(err, http.StatusBadRequest)
	}

	std, err = h.student.Update(ctx, std, uSub)
	if err != nil {
		return fmt.Errorf("update: studentID[%s] uSub[%+v]: %w", studentID, uSub, err)
	}

	return web.Respond(ctx, w, toAppStudent(std), http.StatusOK)
}

// Delete removes a student from the system.
func (h *Handlers) Delete(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	studentID, err := uuid.Parse(web.Param(r, "student_id"))
	if err != nil {
		return validate.NewFieldsError("student_id", err)
	}

	std, err := h.student.QueryByID(ctx, studentID)
	if err != nil {
		switch {
		case errors.Is(err, student.ErrNotFound):
			return web.Respond(ctx, w, nil, http.StatusNoContent)
		default:
			return fmt.Errorf("querybyid: studentID[%s]: %w", studentID, err)
		}
	}

	if err := h.student.Delete(ctx, std); err != nil {
		return fmt.Errorf("delete: studentID[%s]: %w", studentID, err)
	}

	return web.Respond(ctx, w, nil, http.StatusNoContent)
}

// Query returns a list of students with paging.
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

// QueryByID returns a student by its ID.
func (h *Handlers) QueryByID(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	studentID, err := uuid.Parse(web.Param(r, "student_id"))
	if err != nil {
		return validate.NewFieldsError("student_id", err)
	}

	std, err := h.student.QueryByID(ctx, studentID)
	if err != nil {
		switch {
		case errors.Is(err, student.ErrNotFound):
			return v1.NewRequestError(err, http.StatusNotFound)
		default:
			return fmt.Errorf("querybyid: studentID[%s]: %w", studentID, err)
		}
	}

	return web.Respond(ctx, w, toAppStudent(std), http.StatusOK)
}
