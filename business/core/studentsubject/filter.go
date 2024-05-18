package studentsubject

import (
	"fmt"

	"github.com/PhyoYazar/uas/business/sys/validate"
	"github.com/google/uuid"
)

// QueryFilter holds the available fields a query can be filtered on.
type QueryFilter struct {
	ID        *uuid.UUID `validate:"omitempty"`
	StudentID *uuid.UUID `validate:"omitempty"`
	SubjectID *uuid.UUID `validate:"omitempty"`
	Mark      *int       `validate:"omitempty"`
}

// Validate checks the data in the model is considered clean.
func (qf *QueryFilter) Validate() error {
	if err := validate.Check(qf); err != nil {
		return fmt.Errorf("validate: %w", err)
	}
	return nil
}

// WithID sets the ID field of the QueryFilter value.
func (qf *QueryFilter) WithStudentSubjectID(stdSubID uuid.UUID) {
	qf.ID = &stdSubID
}

// WithStudentID sets the Name field of the QueryFilter value.
func (qf *QueryFilter) WithStudentID(studentID uuid.UUID) {
	qf.StudentID = &studentID
}

// WithSubjectID sets the Name field of the QueryFilter value.
func (qf *QueryFilter) WithSubjectID(subjectID uuid.UUID) {
	qf.SubjectID = &subjectID
}

// WithMark sets the Name field of the QueryFilter value.
func (qf *QueryFilter) WithMark(mark int) {
	qf.Mark = &mark
}
