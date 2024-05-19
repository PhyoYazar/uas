package coga

import (
	"fmt"

	"github.com/PhyoYazar/uas/business/sys/validate"
	"github.com/google/uuid"
)

// QueryFilter holds the available fields a query can be filtered on.
type QueryFilter struct {
	ID   *uuid.UUID `validate:"omitempty"`
	CoID *uuid.UUID `validate:"omitempty"`
	GaID *uuid.UUID `validate:"omitempty"`
	Mark *int       `validate:"omitempty"`
}

// Validate checks the data in the model is considered clean.
func (qf *QueryFilter) Validate() error {
	if err := validate.Check(qf); err != nil {
		return fmt.Errorf("validate: %w", err)
	}
	return nil
}

// WithID sets the ID field of the QueryFilter value.
func (qf *QueryFilter) WithCoGaID(cogaID uuid.UUID) {
	qf.ID = &cogaID
}

// WithStudentID sets the Name field of the QueryFilter value.
func (qf *QueryFilter) WithCoID(coID uuid.UUID) {
	qf.CoID = &coID
}

// WithSubjectID sets the Name field of the QueryFilter value.
func (qf *QueryFilter) WithGaID(gaID uuid.UUID) {
	qf.GaID = &gaID
}
