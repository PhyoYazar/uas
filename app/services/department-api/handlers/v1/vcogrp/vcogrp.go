package vcogrp

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/PhyoYazar/uas/business/core/vco"
	"github.com/PhyoYazar/uas/business/sys/validate"
	v1 "github.com/PhyoYazar/uas/business/web/v1"
	"github.com/PhyoYazar/uas/foundation/web"
	"github.com/google/uuid"
)

// Handlers manages the set of user endpoints.
type Handlers struct {
	vco *vco.Core
}

// New constructs a handlers for route access.
func New(vco *vco.Core) *Handlers {
	return &Handlers{
		vco: vco,
	}
}

// QueryByID returns a subject by its ID.
func (h *Handlers) QueryByID(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	vcoID, err := uuid.Parse(web.Param(r, "co_id"))
	if err != nil {
		return validate.NewFieldsError("co_id", err)
	}

	sub, err := h.vco.QueryByID(ctx, vcoID)
	if err != nil {
		switch {
		case errors.Is(err, vco.ErrNotFound):
			return v1.NewRequestError(err, http.StatusNotFound)
		default:
			return fmt.Errorf("querybyid: vcoID[%s]: %w", vcoID, err)
		}
	}

	return web.Respond(ctx, w, sub, http.StatusOK)
}
