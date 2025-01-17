package cogrp

import (
	"fmt"
	"time"

	"github.com/PhyoYazar/uas/business/core/co"
	"github.com/PhyoYazar/uas/business/sys/validate"
	"github.com/google/uuid"
)

// AppUser represents information about an individual co.
type AppCo struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	SubjectID   string `json:"subjectID"`
	Instance    int    `json:"instance"`
	Mark        int    `json:"mark"`
	DateCreated string `json:"dateCreated"`
	DateUpdated string `json:"dateUpdated"`
}

func toAppCo(co co.Co) AppCo {

	return AppCo{
		ID:          co.ID.String(),
		Name:        co.Name,
		SubjectID:   co.SubjectID.String(),
		Instance:    co.Instance,
		Mark:        co.Mark,
		DateCreated: co.DateCreated.Format(time.RFC3339),
		DateUpdated: co.DateUpdated.Format(time.RFC3339),
	}
}

// =============================================================================

// AppNewCo contains information needed to create a new co.
type AppNewCo struct {
	Name      string `json:"name" validate:"required"`
	SubjectID string `json:"subjectID" validate:"required"`
	Instance  int    `json:"instance" validate:"required"`
	Mark      int    `json:"mark" validate:"required"`
}

func toCoreNewCo(app AppNewCo) (co.NewCo, error) {
	subjectId, err := uuid.Parse(app.SubjectID)
	if err != nil {
		return co.NewCo{}, fmt.Errorf("parsing subjectid: %w", err)
	}

	co := co.NewCo{
		Name:      app.Name,
		SubjectID: subjectId,
		Instance:  app.Instance,
		Mark:      app.Mark,
	}

	return co, nil
}

// Validate checks the data in the model is considered clean.
func (app AppNewCo) Validate() error {
	if err := validate.Check(app); err != nil {
		return err
	}
	return nil
}

// =============================================================================
