package subjectgrp

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/PhyoYazar/uas/business/core/subject"
	v1 "github.com/PhyoYazar/uas/business/web/v1"
	"github.com/PhyoYazar/uas/foundation/web"
)

// Handlers manages the set of user endpoints.
type Handlers struct {
	subject *subject.Core
}

// New constructs a handlers for route access.
func New(subject *subject.Core) *Handlers {
	return &Handlers{
		subject: subject,
	}
}

// Create adds a new user to the system.
func (h *Handlers) Create(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var app AppNewSubject
	if err := web.Decode(r, &app); err != nil {
		return err
	}

	nc, err := toCoreNewSubject(app)
	if err != nil {
		return v1.NewRequestError(err, http.StatusBadRequest)
	}

	sub, err := h.subject.Create(ctx, nc)
	if err != nil {
		if errors.Is(err, subject.ErrUniqueSubjectYear) {
			return v1.NewRequestError(err, http.StatusConflict)
		}
		return fmt.Errorf("create: sub[%+v]: %w", sub, err)
	}

	return web.Respond(ctx, w, toAppSubject(sub), http.StatusCreated)
}
