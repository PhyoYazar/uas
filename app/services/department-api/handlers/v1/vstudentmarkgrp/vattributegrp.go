package vstudentmarkgrp

import (
	"context"
	"fmt"
	"net/http"

	"github.com/PhyoYazar/uas/business/core/vstudentmark"
	"github.com/PhyoYazar/uas/business/sys/validate"
	"github.com/PhyoYazar/uas/business/web/v1/paging"
	"github.com/PhyoYazar/uas/foundation/web"
	"github.com/google/uuid"
)

// Handlers manages the set of user endpoints.
type Handlers struct {
	vstudentmark *vstudentmark.Core
}

// New constructs a handlers for route access.
func New(vstudentmark *vstudentmark.Core) *Handlers {
	return &Handlers{
		vstudentmark: vstudentmark,
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

	atts, err := h.vstudentmark.Query(ctx, filter, orderBy, page.Number, page.RowsPerPage, subjectID)
	if err != nil {
		return fmt.Errorf("query: %w", err)
	}

	total, err := h.vstudentmark.Count(ctx, filter)
	if err != nil {
		return fmt.Errorf("count: %w", err)
	}

	return web.Respond(ctx, w, paging.NewResponse(atts, total, page.Number, page.RowsPerPage), http.StatusOK)
}

// // Delete removes a subject from the system.
// func (h *Handlers) RemoveAttribute(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
// 	attributeID, err := uuid.Parse(web.Param(r, "attribute_id"))
// 	if err != nil {
// 		return validate.NewFieldsError("attribute_id", err)
// 	}

// 	filter, err := parseFilter(r)
// 	if err != nil {
// 		return err
// 	}

// 	filter.ID = &attributeID

// 	ra := vstudentmark.VRemoveAttribute{
// 		SubjectID:   *filter.SubID,
// 		AttributeID: attributeID,
// 	}

// 	if err := h.vstudentmark.RemoveAttribute(ctx, ra); err != nil {
// 		return fmt.Errorf("remove: marks and ca[%s]: %w", attributeID, err)
// 	}

// 	return web.Respond(ctx, w, nil, http.StatusNoContent)
// }
