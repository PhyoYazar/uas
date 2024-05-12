package subject

import (
	"fmt"
	"time"

	"github.com/PhyoYazar/uas/business/sys/validate"
	"github.com/google/uuid"
)

// QueryFilter holds the available fields a query can be filtered on.
type QueryFilter struct {
	ID               *uuid.UUID `validate:"omitempty"`
	Name             *string    `validate:"omitempty,min=3"`
	Code             *string    `validate:"omitempty"`
	Year             *string    `validate:"omitempty"`
	AcademicYear     *string    `validate:"omitempty"`
	Instructor       *string    `validate:"omitempty"`
	StartCreatedDate *time.Time `validate:"omitempty"`
	EndCreatedDate   *time.Time `validate:"omitempty"`
}

// Validate checks the data in the model is considered clean.
func (qf *QueryFilter) Validate() error {
	if err := validate.Check(qf); err != nil {
		return fmt.Errorf("validate: %w", err)
	}
	return nil
}

// WithSubjectID sets the ID field of the QueryFilter value.
func (qf *QueryFilter) WithSubjectID(subjectID uuid.UUID) {
	qf.ID = &subjectID
}

// WithName sets the Name field of the QueryFilter value.
func (qf *QueryFilter) WithName(name string) {
	qf.Name = &name
}

// WithCode sets the Code field of the QueryFilter value.
func (qf *QueryFilter) WithCode(code string) {
	qf.Code = &code
}

// WithYear sets the Year field of the QueryFilter value.
func (qf *QueryFilter) WithYear(year string) {
	qf.Year = &year
}

// WithAcademicYear sets the AcademicYear field of the QueryFilter value.
func (qf *QueryFilter) WithAcademicYear(academicYear string) {
	qf.AcademicYear = &academicYear
}

// WithInstructor sets the Instructor field of the QueryFilter value.
func (qf *QueryFilter) WithInstructor(instructor string) {
	qf.Instructor = &instructor
}

// WithStartDateCreated sets the DateCreated field of the QueryFilter value.
func (qf *QueryFilter) WithStartDateCreated(startDate time.Time) {
	d := startDate.UTC()
	qf.StartCreatedDate = &d
}

// WithEndCreatedDate sets the DateCreated field of the QueryFilter value.
func (qf *QueryFilter) WithEndCreatedDate(endDate time.Time) {
	d := endDate.UTC()
	qf.EndCreatedDate = &d
}
