package markgrp

import (
	"fmt"
	"time"

	"github.com/PhyoYazar/uas/business/core/mark"
	"github.com/PhyoYazar/uas/business/sys/validate"
	"github.com/google/uuid"
)

// AppCoGa represents information about an individual coga.
type AppMark struct {
	ID          string `json:"id"`
	SubjectID   string `json:"subjectID"`
	GaID        string `json:"gaID"`
	AttributeID string `json:"attributeID"`
	Mark        int    `json:"mark"`
	DateCreated string `json:"dateCreated"`
	DateUpdated string `json:"dateUpdated"`
}

func toAppMark(m mark.Mark) AppMark {

	return AppMark{
		ID:          m.ID.String(),
		SubjectID:   m.SubjectID.String(),
		GaID:        m.GaID.String(),
		AttributeID: m.AttributeID.String(),
		Mark:        m.Mark,
		DateCreated: m.DateCreated.Format(time.RFC3339),
		DateUpdated: m.DateUpdated.Format(time.RFC3339),
	}
}

// =============================================================================

// AppNewCoGa contains information needed to create a new coga.
type AppNewMark struct {
	SubjectID   string `json:"subjectID" validate:"required"`
	GaId        string `json:"gaID" validate:"required"`
	AttributeID string `json:"attributeID" validate:"required"`
	Mark        int    `json:"mark"`
}

func toCoreNewMark(app AppNewMark) (mark.NewMark, error) {

	var err error
	subjectID, err := uuid.Parse(app.SubjectID)
	if err != nil {
		return mark.NewMark{}, fmt.Errorf("error parsing subjectid string to uuid: %w", err)

	}

	gaID, err := uuid.Parse(app.GaId)
	if err != nil {
		return mark.NewMark{}, fmt.Errorf("error parsing gaid string string to uuid: %w", err)

	}

	attID, err := uuid.Parse(app.AttributeID)
	if err != nil {
		return mark.NewMark{}, fmt.Errorf("error parsing attributeid string string to uuid: %w", err)

	}

	cg := mark.NewMark{
		SubjectID:   subjectID,
		GaID:        gaID,
		AttributeID: attID,
		Mark:        app.Mark,
	}

	return cg, nil
}

// Validate checks the data in the model is considered clean.
func (app AppNewMark) Validate() error {
	if err := validate.Check(app); err != nil {
		return err
	}
	return nil
}

// =============================================================================

type GaMark struct {
	GaId uuid.UUID `json:"gaID"`
	Mark int       `json:"mark"`
}

type MarkByConnectingCOGA struct {
	CoIDs       []uuid.UUID `json:"coIDs"`
	Gas         []GaMark    `json:"gas"`
	SubjectID   uuid.UUID   `json:"subjectID"`
	AttributeID uuid.UUID   `json:"attributeID"`
	FullMark    int         `json:"fullMark"`
}
