package fullmarkgrp

import (
	"fmt"
	"time"

	"github.com/PhyoYazar/uas/business/core/fullmark"
	"github.com/PhyoYazar/uas/business/sys/validate"
	"github.com/google/uuid"
)

// AppFullMark represents information about an individual fullmark.
type AppFullMark struct {
	ID          string `json:"id"`
	SubjectID   string `json:"subjectID"`
	AttributeID string `json:"attributeID"`
	Mark        int    `json:"mark"`
	DateCreated string `json:"dateCreated"`
	DateUpdated string `json:"dateUpdated"`
}

func toAppFullMark(fm fullmark.FullMark) AppFullMark {

	return AppFullMark{
		ID:          fm.ID.String(),
		SubjectID:   fm.SubjectID.String(),
		AttributeID: fm.AttributeID.String(),
		Mark:        fm.Mark,
		DateCreated: fm.DateCreated.Format(time.RFC3339),
		DateUpdated: fm.DateUpdated.Format(time.RFC3339),
	}
}

// =============================================================================

// AppNewFullMark contains information needed to create a new fullmark.
type AppNewFullMark struct {
	SubjectID   string `json:"subjectID" validate:"required"`
	AttributeID string `json:"attributeID" validate:"required"`
	Mark        int    `json:"mark"`
}

func toCoreNewFullMark(app AppNewFullMark) (fullmark.NewFullMark, error) {

	var err error
	subjectID, err := uuid.Parse(app.SubjectID)
	if err != nil {
		return fullmark.NewFullMark{}, fmt.Errorf("error parsing subjectid string to uuid: %w", err)

	}

	attID, err := uuid.Parse(app.AttributeID)
	if err != nil {
		return fullmark.NewFullMark{}, fmt.Errorf("error parsing attributeid string string to uuid: %w", err)

	}

	cg := fullmark.NewFullMark{
		SubjectID:   subjectID,
		AttributeID: attID,
		Mark:        app.Mark,
	}

	return cg, nil
}

// Validate checks the data in the model is considered clean.
func (app AppNewFullMark) Validate() error {
	if err := validate.Check(app); err != nil {
		return err
	}
	return nil
}
