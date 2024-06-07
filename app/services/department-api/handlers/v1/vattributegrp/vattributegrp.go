package vattributegrp

import (
	"context"
	"fmt"
	"net/http"

	"github.com/PhyoYazar/uas/business/core/vattribute"
	"github.com/PhyoYazar/uas/business/sys/validate"
	"github.com/PhyoYazar/uas/business/web/v1/paging"
	"github.com/PhyoYazar/uas/foundation/web"
	"github.com/google/uuid"
)

// Handlers manages the set of user endpoints.
type Handlers struct {
	vattribute *vattribute.Core
}

// New constructs a handlers for route access.
func New(vattribute *vattribute.Core) *Handlers {
	return &Handlers{
		vattribute: vattribute,
	}
}

// Query returns a list of subjects with paging.
func (h *Handlers) Query(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	subjectID, err := uuid.Parse(web.Param(r, "subject_id"))
	if err != nil {
		return validate.NewFieldsError("subject_id", err)
	}

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

	atts, err := h.vattribute.Query(ctx, filter, orderBy, page.Number, page.RowsPerPage, subjectID)
	if err != nil {
		return fmt.Errorf("query: %w", err)
	}

	// total, err := h.vattribute.Count(ctx, filter)
	// if err != nil {
	// 	return fmt.Errorf("count: %w", err)
	// }

	return web.Respond(ctx, w, paging.NewResponse(atts, 0, page.Number, page.RowsPerPage), http.StatusOK)
}

// Query returns a list of subjects with paging.
func (h *Handlers) QueryAttributeWithGaMark(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	// subjectID, err := uuid.Parse(web.Param(r, "subject_id"))
	// if err != nil {
	// 	return validate.NewFieldsError("subject_id", err)
	// }

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

	atts, err := h.vattribute.QueryAttributeWithGaMark(ctx, filter, orderBy, page.Number, page.RowsPerPage)
	if err != nil {
		return fmt.Errorf("query: %w", err)
	}

	// total, err := h.vattribute.Count(ctx, filter)
	// if err != nil {
	// 	return fmt.Errorf("count: %w", err)
	// }

	return web.Respond(ctx, w, paging.NewResponse(atts, 0, page.Number, page.RowsPerPage), http.StatusOK)
}
