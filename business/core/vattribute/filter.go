package vattribute

import (
	"fmt"

	"github.com/PhyoYazar/uas/business/core/attribute"
	"github.com/PhyoYazar/uas/business/sys/validate"
	"github.com/google/uuid"
)

// QueryFilter holds the available fields a query can be filtered on.
type QueryFilter struct {
	ID       *uuid.UUID      `validate:"omitempty"`
	SubID    *uuid.UUID      `validate:"omitempty"`
	Name     *string         `validate:"omitempty,min=3"`
	Type     *attribute.Type `validate:"omitempty"`
	Instance *int            `validate:"omitempty"`
}

// Validate checks the data in the model is considered clean.
func (qf *QueryFilter) Validate() error {
	if err := validate.Check(qf); err != nil {
		return fmt.Errorf("validate: %w", err)
	}
	return nil
}

func (qf *QueryFilter) WithAttributeID(attID uuid.UUID) {
	qf.ID = &attID
}

func (qf *QueryFilter) WithSubjectID(subID uuid.UUID) {
	qf.SubID = &subID
}

// WithName sets the Name field of the QueryFilter value.
func (qf *QueryFilter) WithName(name string) {
	qf.Name = &name
}

// WithType sets the Type field of the QueryFilter value.
func (qf *QueryFilter) WithAttributeType(attType attribute.Type) {
	qf.Type = &attType
}

// WithType sets the Type field of the QueryFilter value.
func (qf *QueryFilter) WithInstance(instance int) {
	qf.Instance = &instance
}
