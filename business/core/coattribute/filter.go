package coattribute

import (
	"fmt"

	"github.com/PhyoYazar/uas/business/sys/validate"
	"github.com/google/uuid"
)

// QueryFilter holds the available fields a query can be filtered on.
type QueryFilter struct {
	ID          *uuid.UUID `validate:"omitempty"`
	CoID        *uuid.UUID `validate:"omitempty"`
	AttributeID *uuid.UUID `validate:"omitempty"`
}

// Validate checks the data in the model is considered clean.
func (qf *QueryFilter) Validate() error {
	if err := validate.Check(qf); err != nil {
		return fmt.Errorf("validate: %w", err)
	}
	return nil
}

// WithID sets the ID field of the QueryFilter value.
func (qf *QueryFilter) WithCoAttributeID(coattID uuid.UUID) {
	qf.ID = &coattID
}

// WithStudentID sets the Name field of the QueryFilter value.
func (qf *QueryFilter) WithCoID(coID uuid.UUID) {
	qf.CoID = &coID
}

// WithSubjectID sets the Name field of the QueryFilter value.
func (qf *QueryFilter) WithAttributeID(attID uuid.UUID) {
	qf.AttributeID = &attID
}
