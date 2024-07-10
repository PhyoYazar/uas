package coattributegrp

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/PhyoYazar/uas/business/core/coattribute"
	"github.com/PhyoYazar/uas/business/sys/validate"
	v1 "github.com/PhyoYazar/uas/business/web/v1"
	"github.com/PhyoYazar/uas/business/web/v1/paging"
	"github.com/PhyoYazar/uas/foundation/web"
	"github.com/google/uuid"
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

// Update updates a student in the system.
func (h *Handlers) Update(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var app AppUpdateCoAttribute
	if err := web.Decode(r, &app); err != nil {
		return err
	}

	caID, err := uuid.Parse(web.Param(r, "co_attribute_id"))
	if err != nil {
		return validate.NewFieldsError("co_attribute_id", err)
	}

	ca, err := h.coattribute.QueryByID(ctx, caID)
	if err != nil {
		switch {
		case errors.Is(err, coattribute.ErrNotFound):
			return v1.NewRequestError(err, http.StatusNotFound)
		default:
			return fmt.Errorf("querybyid: caID[%s]: %w", caID, err)
		}
	}

	uCa, err := toCoreUpdateCoAttribute(app)
	if err != nil {
		return v1.NewRequestError(err, http.StatusBadRequest)
	}

	ca, err = h.coattribute.Update(ctx, ca, uCa)
	if err != nil {
		return fmt.Errorf("update: caID[%s] uCa[%+v]: %w", caID, uCa, err)
	}

	return web.Respond(ctx, w, toAppCoAttribute(ca), http.StatusOK)
}

// Delete removes a subject from the system.
func (h *Handlers) Delete(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	caID, err := uuid.Parse(web.Param(r, "co_attribute_id"))
	if err != nil {
		return validate.NewFieldsError("co_attribute_id", err)
	}

	if err := h.coattribute.Delete(ctx, caID.String()); err != nil {
		return fmt.Errorf("delete: coAttributeID[%s]: %w", caID, err)
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

// QueryByID returns a student by its ID.
func (h *Handlers) QueryByID(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	caID, err := uuid.Parse(web.Param(r, "co_attribute_id"))
	if err != nil {
		return validate.NewFieldsError("co_attribute_id", err)
	}

	ca, err := h.coattribute.QueryByID(ctx, caID)
	if err != nil {
		switch {
		case errors.Is(err, coattribute.ErrNotFound):
			return v1.NewRequestError(err, http.StatusNotFound)
		default:
			return fmt.Errorf("querybyid: caID[%s]: %w", caID, err)
		}
	}

	return web.Respond(ctx, w, toAppCoAttribute(ca), http.StatusOK)
}
