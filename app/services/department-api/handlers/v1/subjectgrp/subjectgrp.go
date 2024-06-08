package subjectgrp

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/PhyoYazar/uas/business/core/subject"
	"github.com/PhyoYazar/uas/business/sys/validate"
	v1 "github.com/PhyoYazar/uas/business/web/v1"
	"github.com/PhyoYazar/uas/business/web/v1/paging"
	"github.com/PhyoYazar/uas/foundation/web"
	"github.com/google/uuid"
)

// Handlers manages the set of user endpoints.
type Handlers struct {
	subject *subject.Core
}

// New constructs a handlers for route access.
func New(subject *subject.Core) *Handlers {
	return &Handlers{
		subject: subject,
	}
}

// Create adds a new user to the system.
func (h *Handlers) Create(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var app AppNewSubject
	if err := web.Decode(r, &app); err != nil {
		return err
	}

	nc, err := toCoreNewSubject(app)
	if err != nil {
		return v1.NewRequestError(err, http.StatusBadRequest)
	}

	sub, err := h.subject.Create(ctx, nc)
	if err != nil {
		if errors.Is(err, subject.ErrUniqueSubjectYear) {
			return v1.NewRequestError(err, http.StatusConflict)
		}
		return fmt.Errorf("create: sub[%+v]: %w", sub, err)
	}

	return web.Respond(ctx, w, toAppSubject(sub), http.StatusCreated)
}

// Update updates a subject in the system.
func (h *Handlers) Update(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var app AppUpdateSubject
	if err := web.Decode(r, &app); err != nil {
		return err
	}

	subjectID, err := uuid.Parse(web.Param(r, "subject_id"))
	if err != nil {
		return validate.NewFieldsError("subject_id", err)
	}

	sub, err := h.subject.QueryByID(ctx, subjectID)
	if err != nil {
		switch {
		case errors.Is(err, subject.ErrNotFound):
			return v1.NewRequestError(err, http.StatusNotFound)
		default:
			return fmt.Errorf("querybyid: subjectID[%s]: %w", subjectID, err)
		}
	}

	uSub, err := toCoreUpdateSubject(app)
	if err != nil {
		return v1.NewRequestError(err, http.StatusBadRequest)
	}

	sub, err = h.subject.Update(ctx, sub, uSub)
	if err != nil {
		return fmt.Errorf("update: subjectID[%s] uSub[%+v]: %w", subjectID, uSub, err)
	}

	return web.Respond(ctx, w, toAppSubject(sub), http.StatusOK)
}

// Delete removes a subject from the system.
func (h *Handlers) Delete(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	subjectID, err := uuid.Parse(web.Param(r, "subject_id"))
	if err != nil {
		return validate.NewFieldsError("subject_id", err)
	}

	sub, err := h.subject.QueryByID(ctx, subjectID)
	if err != nil {
		switch {
		case errors.Is(err, subject.ErrNotFound):
			return web.Respond(ctx, w, nil, http.StatusNoContent)
		default:
			return fmt.Errorf("querybyid: subjectID[%s]: %w", subjectID, err)
		}
	}

	if err := h.subject.Delete(ctx, sub); err != nil {
		return fmt.Errorf("delete: subjectID[%s]: %w", subjectID, err)
	}

	return web.Respond(ctx, w, nil, http.StatusNoContent)
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

	subjects, err := h.subject.Query(ctx, filter, orderBy, page.Number, page.RowsPerPage)
	if err != nil {
		return fmt.Errorf("query: %w", err)
	}

	items := make([]AppSubject, len(subjects))
	for i, sub := range subjects {
		items[i] = toAppSubject(sub)
	}

	total, err := h.subject.Count(ctx, filter)
	if err != nil {
		return fmt.Errorf("count: %w", err)
	}

	return web.Respond(ctx, w, paging.NewResponse(items, total, page.Number, page.RowsPerPage), http.StatusOK)
}

// QueryByID returns a subject by its ID.
func (h *Handlers) QueryByID(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	subjectID, err := uuid.Parse(web.Param(r, "subject_id"))
	if err != nil {
		return validate.NewFieldsError("subject_id", err)
	}

	sub, err := h.subject.QueryByID(ctx, subjectID)
	if err != nil {
		switch {
		case errors.Is(err, subject.ErrNotFound):
			return v1.NewRequestError(err, http.StatusNotFound)
		default:
			return fmt.Errorf("querybyid: subjectID[%s]: %w", subjectID, err)
		}
	}

	return web.Respond(ctx, w, toAppSubject(sub), http.StatusOK)
}
