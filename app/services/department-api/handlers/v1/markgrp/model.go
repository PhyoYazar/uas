package markgrp

import (
	"time"

	"github.com/PhyoYazar/uas/business/core/mark"
	"github.com/PhyoYazar/uas/business/sys/validate"
)

// AppMark represents information about an individual mark.
type AppMark struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Instance    int    `json:"instance"`
	DateCreated string `json:"dateCreated"`
	DateUpdated string `json:"dateUpdated"`
}

func toAppMark(mark mark.Mark) AppMark {

	return AppMark{
		ID:          mark.ID.String(),
		Name:        mark.Name,
		Type:        mark.Type.Name(),
		Instance:    mark.Instance,
		DateCreated: mark.DateCreated.Format(time.RFC3339),
		DateUpdated: mark.DateUpdated.Format(time.RFC3339),
	}
}

// =============================================================================

// AppNewMark contains information needed to create a new mark.
type AppNewMark struct {
	Name     string `json:"name" validate:"required"`
	Type     string `json:"type" validate:"required"`
	Instance int    `json:"instance" validate:"required"`
}

func toCoreNewMark(app AppNewMark) (mark.NewMark, error) {

	mark := mark.NewMark{
		Name:     app.Name,
		Type:     mark.MustParseMarkType(app.Type),
		Instance: app.Instance,
	}

	return mark, nil
}

// Validate checks the data in the model is considered clean.
func (app AppNewMark) Validate() error {
	if err := validate.Check(app); err != nil {
		return err
	}
	return nil
}

// =============================================================================
