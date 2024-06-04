package cogagrp

import (
	"fmt"
	"time"

	"github.com/PhyoYazar/uas/business/core/coga"
	"github.com/PhyoYazar/uas/business/sys/validate"
	"github.com/google/uuid"
)

// AppCoGa represents information about an individual coga.
type AppCoGa struct {
	ID          string `json:"id"`
	CoID        string `json:"coID"`
	GaID        string `json:"gaID"`
	DateCreated string `json:"dateCreated"`
	DateUpdated string `json:"dateUpdated"`
}

func toAppCoGa(mark coga.CoGa) AppCoGa {

	return AppCoGa{
		ID:          mark.ID.String(),
		CoID:        mark.CoID.String(),
		GaID:        mark.GaID.String(),
		DateCreated: mark.DateCreated.Format(time.RFC3339),
		DateUpdated: mark.DateUpdated.Format(time.RFC3339),
	}
}

// =============================================================================

// AppNewCoGa contains information needed to create a new coga.
type AppNewCoGa struct {
	CoID string `json:"coID" validate:"required"`
	GaID string `json:"gaID" validate:"required"`
}

func toCoreNewCoGa(app AppNewCoGa) (coga.NewCoGa, error) {

	var err error
	coID, err := uuid.Parse(app.CoID)
	if err != nil {
		return coga.NewCoGa{}, fmt.Errorf("error parsing coid string to uuid: %w", err)

	}

	gaID, err := uuid.Parse(app.GaID)
	if err != nil {
		return coga.NewCoGa{}, fmt.Errorf("error parsing gaid string string to uuid: %w", err)

	}

	cg := coga.NewCoGa{
		CoID: coID,
		GaID: gaID,
	}

	return cg, nil
}

// Validate checks the data in the model is considered clean.
func (app AppNewCoGa) Validate() error {
	if err := validate.Check(app); err != nil {
		return err
	}
	return nil
}

// =============================================================================

// create Co and then connect created co with ga
type AppConnectCoGa struct {
	SubjectID  uuid.UUID   `json:"subjectID" validate:"required"`
	CoName     string      `json:"coName" validate:"required"`
	CoInstance int         `json:"coInstance" validate:"required"`
	GaID       []uuid.UUID `json:"gaIDs" validate:"required"`
}
