package vsubjectgrp

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/PhyoYazar/uas/business/core/vsubject"
	"github.com/PhyoYazar/uas/business/sys/validate"
	v1 "github.com/PhyoYazar/uas/business/web/v1"
	"github.com/PhyoYazar/uas/foundation/web"
	"github.com/google/uuid"
)

// Handlers manages the set of user endpoints.
type Handlers struct {
	vsubject *vsubject.Core
}

// New constructs a handlers for route access.
func New(vsubject *vsubject.Core) *Handlers {
	return &Handlers{
		vsubject: vsubject,
	}
}

// QueryByID returns a subject by its ID.
func (h *Handlers) QueryByID(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	subjectID, err := uuid.Parse(web.Param(r, "subject_id"))
	if err != nil {
		return validate.NewFieldsError("subject_id", err)
	}

	sub, err := h.vsubject.QueryByID(ctx, subjectID)
	if err != nil {
		switch {
		case errors.Is(err, vsubject.ErrNotFound):
			return v1.NewRequestError(err, http.StatusNotFound)
		default:
			return fmt.Errorf("querybyid: subjectID[%s]: %w", subjectID, err)
		}
	}

	return web.Respond(ctx, w, sub, http.StatusOK)
}
