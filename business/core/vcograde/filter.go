package vcograde

import (
	"fmt"

	"github.com/PhyoYazar/uas/business/sys/validate"
	"github.com/google/uuid"
)

// QueryFilter holds the available fields a query can be filtered on.
type QueryFilter struct {
	ID           *uuid.UUID `validate:"omitempty"`
	Year         *string    `validate:"omitempty,required"`
	AcademicYear *string    `validate:"omitempty,required"`
}

// Validate checks the data in the model is considered clean.
func (qf *QueryFilter) Validate() error {
	if err := validate.Check(qf); err != nil {
		return fmt.Errorf("validate: %w", err)
	}
	return nil
}

// WithStudentID sets the ID field of the QueryFilter value.
func (qf *QueryFilter) WithStudentID(studentID uuid.UUID) {
	qf.ID = &studentID
}

// WithYear sets the Year field of the QueryFilter value.
func (qf *QueryFilter) WithYear(year string) {
	qf.Year = &year
}

// WithAcademicYear sets the AcademicYear field of the QueryFilter value.
func (qf *QueryFilter) WithAcademicYear(academicYear string) {
	qf.AcademicYear = &academicYear
}
