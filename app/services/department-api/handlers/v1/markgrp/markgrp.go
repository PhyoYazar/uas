package markgrp

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/PhyoYazar/uas/business/core/coattribute"
	"github.com/PhyoYazar/uas/business/core/mark"
	v1 "github.com/PhyoYazar/uas/business/web/v1"
	"github.com/PhyoYazar/uas/business/web/v1/paging"
	"github.com/PhyoYazar/uas/foundation/web"
)

// Handlers manages the set of ga endpoints.
type Handlers struct {
	mark        *mark.Core
	coAttribute *coattribute.Core
}

// New constructs a handlers for route access.
func New(mark *mark.Core, coAttribute *coattribute.Core) *Handlers {
	return &Handlers{
		mark:        mark,
		coAttribute: coAttribute,
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
		})
		if err != nil {
			if errors.Is(err, mark.ErrUniqueMark) {
				return v1.NewRequestError(err, http.StatusConflict)
			}
			return fmt.Errorf("create: mark[%+v]: %w", m, err)
		}

	}

	return nil
}
