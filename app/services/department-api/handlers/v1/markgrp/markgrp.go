package markgrp

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/PhyoYazar/uas/business/core/coattribute"
	"github.com/PhyoYazar/uas/business/core/fullmark"
	"github.com/PhyoYazar/uas/business/core/mark"
	"github.com/PhyoYazar/uas/business/sys/validate"
	v1 "github.com/PhyoYazar/uas/business/web/v1"
	"github.com/PhyoYazar/uas/business/web/v1/paging"
	"github.com/PhyoYazar/uas/foundation/web"
	"github.com/google/uuid"
)

// Handlers manages the set of ga endpoints.
type Handlers struct {
	mark        *mark.Core
	coAttribute *coattribute.Core
	fullmark    *fullmark.Core
}

// New constructs a handlers for route access.
func New(mark *mark.Core, coAttribute *coattribute.Core, fullMark *fullmark.Core) *Handlers {
	return &Handlers{
		mark:        mark,
		coAttribute: coAttribute,
		fullmark:    fullMark,
	}
}

// Create adds a new ga to the system.
func (h *Handlers) Create(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var app AppNewMark
	if err := web.Decode(r, &app); err != nil {
		return err
	}

	cms, err := toCoreNewMark(app)
	if err != nil {
		return v1.NewRequestError(err, http.StatusBadRequest)
	}

	cm, err := h.mark.Create(ctx, cms)
	if err != nil {
		if errors.Is(err, mark.ErrUniqueMark) {
			return v1.NewRequestError(err, http.StatusConflict)
		}
		return fmt.Errorf("create: mark[%+v]: %w", cm, err)
	}

	return web.Respond(ctx, w, toAppMark(cm), http.StatusCreated)
}

// Update updates a student in the system.
func (h *Handlers) Update(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var app AppUpdateMark
	if err := web.Decode(r, &app); err != nil {
		return err
	}

	markID, err := uuid.Parse(web.Param(r, "mark_id"))
	if err != nil {
		return validate.NewFieldsError("mark_id", err)
	}

	mk, err := h.mark.QueryByID(ctx, markID)
	if err != nil {
		switch {
		case errors.Is(err, mark.ErrNotFound):
			return v1.NewRequestError(err, http.StatusNotFound)
		default:
			return fmt.Errorf("querybyid: markID[%s]: %w", markID, err)
		}
	}

	uMark, err := toCoreUpdateMark(app)
	if err != nil {
		return v1.NewRequestError(err, http.StatusBadRequest)
	}

	mk, err = h.mark.Update(ctx, mk, uMark)
	if err != nil {
		return fmt.Errorf("update: markID[%s] uMark[%+v]: %w", markID, uMark, err)
	}

	return web.Respond(ctx, w, toAppMark(mk), http.StatusOK)
}

// Delete removes a subject from the system.
func (h *Handlers) Delete(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	markID, err := uuid.Parse(web.Param(r, "mark_id"))
	if err != nil {
		return validate.NewFieldsError("mark_id", err)
	}

	if err := h.mark.Delete(ctx, markID.String()); err != nil {
		return fmt.Errorf("delete: markID[%s]: %w", markID, err)
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

	marks, err := h.mark.Query(ctx, filter, orderBy, page.Number, page.RowsPerPage)
	if err != nil {
		return fmt.Errorf("query: %w", err)
	}

	items := make([]AppMark, len(marks))
	for i, cm := range marks {
		items[i] = toAppMark(cm)
	}

	total, err := h.mark.Count(ctx, filter)
	if err != nil {
		return fmt.Errorf("count: %w", err)
	}

	return web.Respond(ctx, w, paging.NewResponse(items, total, page.Number, page.RowsPerPage), http.StatusOK)
}

func (h *Handlers) CreateMarkByConnectingCOGA(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var App MarkByConnectingCOGA

	if err := web.Decode(r, &App); err != nil {
		return err
	}

	// 1. create Co-Attribute
	for _, coId := range App.CoIDs {
		cg, err := h.coAttribute.Create(ctx, coattribute.NewCoAttribute{CoID: coId, AttributeID: App.AttributeID})
		if err != nil {
			if errors.Is(err, coattribute.ErrUniqueCoAttribute) {
				return v1.NewRequestError(err, http.StatusConflict)
			}
			return fmt.Errorf("create: coattribute[%+v]: %w", cg, err)
		}
	}

	// 2. create Mark
	for _, ga := range App.Gas {
		m, err := h.mark.Create(ctx, mark.NewMark{
			SubjectID:   App.SubjectID,
			AttributeID: App.AttributeID,
			GaID:        ga.GaId,
			Mark:        ga.Mark,
			GaMark:      0,
		})
		if err != nil {
			if errors.Is(err, mark.ErrUniqueMark) {
				return v1.NewRequestError(err, http.StatusConflict)
			}
			return fmt.Errorf("create: mark[%+v]: %w", m, err)
		}
	}

	// 3. full mark of the attributes
	fm, err := h.fullmark.Create(ctx,
		fullmark.NewFullMark{
			SubjectID:   App.SubjectID,
			AttributeID: App.AttributeID,
			Mark:        App.FullMark,
		})
	if err != nil {
		if errors.Is(err, fullmark.ErrUniqueFullMark) {
			return v1.NewRequestError(err, http.StatusConflict)
		}
		return fmt.Errorf("create: fullmark with cogaattribute[%+v]: %w", fm, err)
	}

	return nil
}

// QueryByID returns a student by its ID.
func (h *Handlers) QueryByID(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	markID, err := uuid.Parse(web.Param(r, "mark_id"))
	if err != nil {
		return validate.NewFieldsError("mark_id", err)
	}

	mk, err := h.mark.QueryByID(ctx, markID)
	if err != nil {
		switch {
		case errors.Is(err, mark.ErrNotFound):
			return v1.NewRequestError(err, http.StatusNotFound)
		default:
			return fmt.Errorf("querybyid: markID[%s]: %w", markID, err)
		}
	}

	return web.Respond(ctx, w, toAppMark(mk), http.StatusOK)
}
