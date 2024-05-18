package studentsubjectgrp

import (
	"fmt"
	"time"

	"github.com/PhyoYazar/uas/business/core/studentsubject"
	"github.com/PhyoYazar/uas/business/sys/validate"
	"github.com/google/uuid"
)

// AppMark represents information about an individual mark.
type AppStudentSubject struct {
	ID          string `json:"id"`
	StudentID   string `json:"studentID"`
	SubjectID   string `json:"subjectID"`
	Mark        int    `json:"mark"`
	DateCreated string `json:"dateCreated"`
	DateUpdated string `json:"dateUpdated"`
}

func toAppStudentSubject(mark studentsubject.StudentSubject) AppStudentSubject {

	return AppStudentSubject{
		ID:          mark.ID.String(),
		SubjectID:   mark.SubjectID.String(),
		StudentID:   mark.StudentID.String(),
		Mark:        mark.Mark,
		DateCreated: mark.DateCreated.Format(time.RFC3339),
		DateUpdated: mark.DateUpdated.Format(time.RFC3339),
	}
}

// =============================================================================

// AppNewMark contains information needed to create a new mark.
type AppNewStudentSubject struct {
	StudentID string `json:"studentID" validate:"required"`
	SubjectID string `json:"subjectID" validate:"required"`
	Mark      int    `json:"mark" validate:"required"`
}

func toCoreNewStudentSubject(app AppNewStudentSubject) (studentsubject.NewStudentSubject, error) {

	var err error
	subID, err := uuid.Parse(app.SubjectID)
	if err != nil {
		return studentsubject.NewStudentSubject{}, fmt.Errorf("error parsing subjectid string to uuid: %w", err)

	}

	studID, err := uuid.Parse(app.StudentID)
	if err != nil {
		return studentsubject.NewStudentSubject{}, fmt.Errorf("error parsing studentid string string to uuid: %w", err)

	}

	ss := studentsubject.NewStudentSubject{
		SubjectID: subID,
		StudentID: studID,
		Mark:      app.Mark,
	}

	return ss, nil
}

// Validate checks the data in the model is considered clean.
func (app AppNewStudentSubject) Validate() error {
	if err := validate.Check(app); err != nil {
		return err
	}
	return nil
}

// =============================================================================
