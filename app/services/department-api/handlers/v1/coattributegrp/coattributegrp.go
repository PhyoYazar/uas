package coattributegrp

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/PhyoYazar/uas/business/core/coattribute"
	v1 "github.com/PhyoYazar/uas/business/web/v1"
	"github.com/PhyoYazar/uas/business/web/v1/paging"
	"github.com/PhyoYazar/uas/foundation/web"
)

// Handlers manages the set of ga endpoints.
type Handlers struct {
	coattribute *coattribute.Core
}

// New constructs a handlers for route access.
func New(coattribute *coattribute.Core) *Handlers {
	return &Handlers{
		coattribute: coattribute,
	}
}

// Create adds a new ga to the system.
func (h *Handlers) Create(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var app AppNewCoAttribute
	if err := web.Decode(r, &app); err != nil {
		return err
	}

	cgs, err := toCoreNewCoAttribute(app)
	if err != nil {
		return v1.NewRequestError(err, http.StatusBadRequest)
	}

	cg, err := h.coattribute.Create(ctx, cgs)
	if err != nil {
		if errors.Is(err, coattribute.ErrUniqueCoAttribute) {
			return v1.NewRequestError(err, http.StatusConflict)
		}
		return fmt.Errorf("create: ga[%+v]: %w", cg, err)
	}

	return web.Respond(ctx, w, toAppCoAttribute(cg), http.StatusCreated)
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

	cogas, err := h.coattribute.Query(ctx, filter, orderBy, page.Number, page.RowsPerPage)
	if err != nil {
		return fmt.Errorf("query: %w", err)
	}

	items := make([]AppCoAttribute, len(cogas))
	for i, cg := range cogas {
		items[i] = toAppCoAttribute(cg)
	}

	total, err := h.coattribute.Count(ctx, filter)
	if err != nil {
		return fmt.Errorf("count: %w", err)
	}

	return web.Respond(ctx, w, paging.NewResponse(items, total, page.Number, page.RowsPerPage), http.StatusOK)
}
