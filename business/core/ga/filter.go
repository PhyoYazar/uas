package ga

import (
	"fmt"

	"github.com/PhyoYazar/uas/business/sys/validate"
	"github.com/google/uuid"
)

// QueryFilter holds the available fields a query can be filtered on.
type QueryFilter struct {
	ID   *uuid.UUID `validate:"omitempty"`
	Name *string    `validate:"omitempty,min=3"`
	Slug *string    `validate:"omitempty"`
}

// Validate checks the data in the model is considered clean.
func (qf *QueryFilter) Validate() error {
	if err := validate.Check(qf); err != nil {
		return fmt.Errorf("validate: %w", err)
	}
	return nil
}

// WithGaID sets the ID field of the QueryFilter value.
func (qf *QueryFilter) WithGaID(gaID uuid.UUID) {
	qf.ID = &gaID
}

// WithName sets the Name field of the QueryFilter value.
func (qf *QueryFilter) WithName(name string) {
	qf.Name = &name
}

// WithSlug sets the Slug field of the QueryFilter value.
func (qf *QueryFilter) WithSlug(slug string) {
	qf.Slug = &slug
}
