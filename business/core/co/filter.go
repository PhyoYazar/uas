package co

import (
	"fmt"

	"github.com/PhyoYazar/uas/business/sys/validate"
	"github.com/google/uuid"
)

// QueryFilter holds the available fields a query can be filtered on.
type QueryFilter struct {
	ID       *uuid.UUID `validate:"omitempty"`
	Name     *string    `validate:"omitempty,min=3"`
	Instance *int       `validate:"omitempty"`
	Mark     *int       `validate:"omitempty"`
}

// Validate checks the data in the model is considered clean.
func (qf *QueryFilter) Validate() error {
	if err := validate.Check(qf); err != nil {
		return fmt.Errorf("validate: %w", err)
	}
	return nil
}

// WithCoID sets the ID field of the QueryFilter value.
func (qf *QueryFilter) WithCoID(coID uuid.UUID) {
	qf.ID = &coID
}

// WithName sets the Name field of the QueryFilter value.
func (qf *QueryFilter) WithName(name string) {
	qf.Name = &name
}

// WithInstance sets the Instance field of the QueryFilter value.
func (qf *QueryFilter) WithInstance(instance int) {
	qf.Instance = &instance
}

// WithMark sets the Instance field of the QueryFilter value.
func (qf *QueryFilter) WithMark(mark int) {
	qf.Mark = &mark
}
