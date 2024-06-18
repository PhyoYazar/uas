package studentmarkgrp

import (
	"fmt"
	"time"

	"github.com/PhyoYazar/uas/business/core/studentmark"
	"github.com/PhyoYazar/uas/business/sys/validate"
	"github.com/google/uuid"
)

// AppMark represents information about an individual mark.
type AppStudentMark struct {
	ID          string `json:"id"`
	StudentID   string `json:"studentID"`
	SubjectID   string `json:"subjectID"`
	AttributeID string `json:"attributeID"`
	Mark        int    `json:"mark"`
	DateCreated string `json:"dateCreated"`
	DateUpdated string `json:"dateUpdated"`
}

func toAppStudentMark(mark studentmark.StudentMark) AppStudentMark {

	return AppStudentMark{
		ID:          mark.ID.String(),
		SubjectID:   mark.SubjectID.String(),
		StudentID:   mark.StudentID.String(),
		AttributeID: mark.AttributeID.String(),
		Mark:        mark.Mark,
		DateCreated: mark.DateCreated.Format(time.RFC3339),
		DateUpdated: mark.DateUpdated.Format(time.RFC3339),
	}
}

// =============================================================================

// AppNewMark contains information needed to create a new mark.
type AppNewStudentMark struct {
	StudentID   string `json:"studentID" validate:"required"`
	SubjectID   string `json:"subjectID" validate:"required"`
	AttributeID string `json:"attributeID" validate:"required"`
	Mark        int    `json:"mark" validate:"required"`
}

func toCoreNewStudentMark(app AppNewStudentMark) (studentmark.NewStudentMark, error) {

	var err error
	subID, err := uuid.Parse(app.SubjectID)
	if err != nil {
		return studentmark.NewStudentMark{}, fmt.Errorf("error parsing subjectid string to uuid: %w", err)

	}

	studID, err := uuid.Parse(app.StudentID)
	if err != nil {
		return studentmark.NewStudentMark{}, fmt.Errorf("error parsing studentid string string to uuid: %w", err)

	}

	attributeID, err := uuid.Parse(app.AttributeID)
	if err != nil {
		return studentmark.NewStudentMark{}, fmt.Errorf("error parsing attributeid string string to uuid: %w", err)

	}

	ss := studentmark.NewStudentMark{
		SubjectID:   subID,
		StudentID:   studID,
		AttributeID: attributeID,
		Mark:        app.Mark,
	}

	return ss, nil
}

// Validate checks the data in the model is considered clean.
func (app AppNewStudentMark) Validate() error {
	if err := validate.Check(app); err != nil {
		return err
	}
	return nil
}

// =============================================================================

// AppUpdateStudentMark contains information needed to update a studentMark.
type AppUpdateStudentMark struct {
	Mark        *int    `json:"mark"`
	SubjectID   *string `json:"subjectId"`
	StudentID   *string `json:"studentId"`
	AttributeID *string `json:"attributeId"`
}

func toCoreUpdateStudentMark(app AppUpdateStudentMark) (studentmark.UpdateStudentMark, error) {

	nSub := studentmark.UpdateStudentMark{
		Mark: app.Mark,
	}

	return nSub, nil
}

// Validate checks the data in the model is considered clean.
func (app AppUpdateStudentMark) Validate() error {
	if err := validate.Check(app); err != nil {
		return fmt.Errorf("validate: %w", err)
	}
	return nil
}
