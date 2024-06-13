package student

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
	RollNumber       *string    `validate:"omitempty,min=1"`
	Year             *string    `validate:"omitempty"`
	AcademicYear     *string    `validate:"omitempty"`
	StartCreatedDate *time.Time `validate:"omitempty"`
	EndCreatedDate   *time.Time `validate:"omitempty"`
	// Email            *mail.Address `validate:"omitempty"`
	// PhoneNumber      *string       `validate:"omitempty,min=3"`
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

// WithName sets the Name field of the QueryFilter value.
func (qf *QueryFilter) WithName(name string) {
	qf.Name = &name
}

// // WithEmail sets the Email field of the QueryFilter value.
// func (qf *QueryFilter) WithEmail(email mail.Address) {
// 	qf.Email = &email
// }

// // WithPhoneNumber sets the Instructor field of the QueryFilter value.
// func (qf *QueryFilter) WithPhoneNumber(phoneNumber string) {
// 	qf.PhoneNumber = &phoneNumber
// }

// WithYear sets the Year field of the QueryFilter value.
func (qf *QueryFilter) WithYear(year string) {
	qf.Year = &year
}

// WithAcademicYear sets the AcademicYear field of the QueryFilter value.
func (qf *QueryFilter) WithAcademicYear(academicYear string) {
	qf.AcademicYear = &academicYear
}

// WithRollNumber sets the Instructor field of the QueryFilter value.
func (qf *QueryFilter) WithRollNumber(rollNumber string) {
	qf.RollNumber = &rollNumber
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
