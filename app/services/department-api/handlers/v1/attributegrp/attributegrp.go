package attributegrp

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/PhyoYazar/uas/business/core/attribute"
	"github.com/PhyoYazar/uas/business/sys/validate"
	v1 "github.com/PhyoYazar/uas/business/web/v1"
	"github.com/PhyoYazar/uas/business/web/v1/paging"
	"github.com/PhyoYazar/uas/foundation/web"
	"github.com/google/uuid"
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

// Delete removes a subject from the system.
func (h *Handlers) Delete(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	attributeID, err := uuid.Parse(web.Param(r, "attribute_id"))
	if err != nil {
		return validate.NewFieldsError("attribute_id", err)
	}

	sub, err := h.attribute.QueryByID(ctx, attributeID)
	if err != nil {
		switch {
		case errors.Is(err, attribute.ErrNotFound):
			return web.Respond(ctx, w, nil, http.StatusNoContent)
		default:
			return fmt.Errorf("querybyid: attributeID[%s]: %w", attributeID, err)
		}
	}

	if err := h.attribute.Delete(ctx, sub); err != nil {
		return fmt.Errorf("delete: subjectID[%s]: %w", attributeID, err)
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

// QueryByID returns a subject by its ID.
func (h *Handlers) QueryByID(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	attributeID, err := uuid.Parse(web.Param(r, "attribute_id"))
	if err != nil {
		return validate.NewFieldsError("attribute_id", err)
	}

	sub, err := h.attribute.QueryByID(ctx, attributeID)
	if err != nil {
		switch {
		case errors.Is(err, attribute.ErrNotFound):
			return v1.NewRequestError(err, http.StatusNotFound)
		default:
			return fmt.Errorf("querybyid: attributeID[%s]: %w", attributeID, err)
		}
	}

	return web.Respond(ctx, w, toAppAttribute(sub), http.StatusOK)
}
