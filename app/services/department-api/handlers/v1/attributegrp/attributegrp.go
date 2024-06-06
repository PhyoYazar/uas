package attributegrp

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/PhyoYazar/uas/business/core/attribute"
	v1 "github.com/PhyoYazar/uas/business/web/v1"
	"github.com/PhyoYazar/uas/business/web/v1/paging"
	"github.com/PhyoYazar/uas/foundation/web"
)

// Handlers manages the set of ga endpoints.
type Handlers struct {
	attribute *attribute.Core
}

// New constructs a handlers for route access.
func New(attribute *attribute.Core) *Handlers {
	return &Handlers{
		attribute: attribute,
	}
}

// Create adds a new ga to the system.
func (h *Handlers) Create(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var app AppNewAttribute
	if err := web.Decode(r, &app); err != nil {
		return err
	}

	nm, err := toCoreNewAttribute(app)
	if err != nil {
		return v1.NewRequestError(err, http.StatusBadRequest)
	}

	mk, err := h.attribute.Create(ctx, nm)
	if err != nil {
		if errors.Is(err, attribute.ErrUniqueAttribute) {
			return v1.NewRequestError(err, http.StatusConflict)
		}
		return fmt.Errorf("create: attribute[%+v]: %w", mk, err)
	}

	return web.Respond(ctx, w, toAppAttribute(mk), http.StatusCreated)
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

	attri, err := h.attribute.Query(ctx, filter, orderBy, page.Number, page.RowsPerPage)
	if err != nil {
		return fmt.Errorf("query: %w", err)
	}

	items := make([]AppAttribute, len(attri))
	for i, mk := range attri {
		items[i] = toAppAttribute(mk)
	}

	total, err := h.attribute.Count(ctx, filter)
	if err != nil {
		return fmt.Errorf("count: %w", err)
	}

	return web.Respond(ctx, w, paging.NewResponse(items, total, page.Number, page.RowsPerPage), http.StatusOK)
}
