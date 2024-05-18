package comarkgrp

import (
	"fmt"
	"time"

	"github.com/PhyoYazar/uas/business/core/comark"
	"github.com/PhyoYazar/uas/business/sys/validate"
	"github.com/google/uuid"
)

// AppCoGa represents information about an individual coga.
type AppCoMark struct {
	ID          string `json:"id"`
	CoID        string `json:"coID"`
	MarkID      string `json:"markID"`
	DateCreated string `json:"dateCreated"`
	DateUpdated string `json:"dateUpdated"`
}

func toAppCoMark(cm comark.CoMark) AppCoMark {

	return AppCoMark{
		ID:          cm.ID.String(),
		CoID:        cm.CoID.String(),
		MarkID:      cm.MarkID.String(),
		DateCreated: cm.DateCreated.Format(time.RFC3339),
		DateUpdated: cm.DateUpdated.Format(time.RFC3339),
	}
}

// =============================================================================

// AppNewCoGa contains information needed to create a new coga.
type AppNewCoMark struct {
	CoID   string `json:"coID" validate:"required"`
	MarkID string `json:"markID" validate:"required"`
}

func toCoreNewCoMark(app AppNewCoMark) (comark.NewCoMark, error) {

	var err error
	coID, err := uuid.Parse(app.CoID)
	if err != nil {
		return comark.NewCoMark{}, fmt.Errorf("error parsing coid string to uuid: %w", err)

	}

	markID, err := uuid.Parse(app.MarkID)
	if err != nil {
		return comark.NewCoMark{}, fmt.Errorf("error parsing gaid string string to uuid: %w", err)

	}

	cg := comark.NewCoMark{
		CoID:   coID,
		MarkID: markID,
	}

	return cg, nil
}

// Validate checks the data in the model is considered clean.
func (app AppNewCoMark) Validate() error {
	if err := validate.Check(app); err != nil {
		return err
	}
	return nil
}

// =============================================================================
