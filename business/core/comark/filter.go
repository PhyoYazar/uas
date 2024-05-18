package comark

import (
	"fmt"

	"github.com/PhyoYazar/uas/business/sys/validate"
	"github.com/google/uuid"
)

// QueryFilter holds the available fields a query can be filtered on.
type QueryFilter struct {
	ID     *uuid.UUID `validate:"omitempty"`
	CoID   *uuid.UUID `validate:"omitempty"`
	MarkID *uuid.UUID `validate:"omitempty"`
}

// Validate checks the data in the model is considered clean.
func (qf *QueryFilter) Validate() error {
	if err := validate.Check(qf); err != nil {
		return fmt.Errorf("validate: %w", err)
	}
	return nil
}

// WithID sets the ID field of the QueryFilter value.
func (qf *QueryFilter) WithCoMarkID(coMarkID uuid.UUID) {
	qf.ID = &coMarkID
}

// WithCoID sets the Name field of the QueryFilter value.
func (qf *QueryFilter) WithCoID(coID uuid.UUID) {
	qf.CoID = &coID
}

// WithMarkID sets the Name field of the QueryFilter value.
func (qf *QueryFilter) WithMarkID(markID uuid.UUID) {
	qf.MarkID = &markID
}
