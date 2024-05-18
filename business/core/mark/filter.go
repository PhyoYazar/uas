package mark

import (
	"fmt"

	"github.com/PhyoYazar/uas/business/sys/validate"
	"github.com/google/uuid"
)

// QueryFilter holds the available fields a query can be filtered on.
type QueryFilter struct {
	ID       *uuid.UUID `validate:"omitempty"`
	Name     *string    `validate:"omitempty,min=3"`
	Type     *Type      `validate:"omitempty"`
	Instance *int       `validate:"omitempty"`
}

// Validate checks the data in the model is considered clean.
func (qf *QueryFilter) Validate() error {
	if err := validate.Check(qf); err != nil {
		return fmt.Errorf("validate: %w", err)
	}
	return nil
}

// WithMarkID sets the ID field of the QueryFilter value.
func (qf *QueryFilter) WithMarkID(markID uuid.UUID) {
	qf.ID = &markID
}

// WithName sets the Name field of the QueryFilter value.
func (qf *QueryFilter) WithName(name string) {
	qf.Name = &name
}

// WithType sets the Type field of the QueryFilter value.
func (qf *QueryFilter) WithMarkType(markType Type) {
	qf.Type = &markType
}