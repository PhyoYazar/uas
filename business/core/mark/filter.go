package mark

import (
	"fmt"

	"github.com/PhyoYazar/uas/business/sys/validate"
	"github.com/google/uuid"
)

// QueryFilter holds the available fields a query can be filtered on.
type QueryFilter struct {
	ID          *uuid.UUID `validate:"omitempty"`
	SubjectID   *uuid.UUID `validate:"omitempty"`
	GaID        *uuid.UUID `validate:"omitempty"`
	AttributeID *uuid.UUID `validate:"omitempty"`
	Mark        *int       `validate:"omitempty"`
}

// Validate checks the data in the model is considered clean.
func (qf *QueryFilter) Validate() error {
	if err := validate.Check(qf); err != nil {
		return fmt.Errorf("validate: %w", err)
	}
	return nil
}

// WithID sets the ID field of the QueryFilter value.
func (qf *QueryFilter) WithMarkID(markID uuid.UUID) {
	qf.ID = &markID
}

// WithSubjectID sets the Name field of the QueryFilter value.
func (qf *QueryFilter) WithSubjectID(subID uuid.UUID) {
	qf.SubjectID = &subID
}

// WithMarkID sets the Name field of the QueryFilter value.
func (qf *QueryFilter) WithGaID(gaID uuid.UUID) {
	qf.GaID = &gaID
}

// WithAttributeID sets the Name field of the QueryFilter value.
func (qf *QueryFilter) WithAttributeID(attID uuid.UUID) {
	qf.AttributeID = &attID
}

// WithMark sets the Name field of the QueryFilter value.
func (qf *QueryFilter) WithMark(mark int) {
	qf.Mark = &mark
}
