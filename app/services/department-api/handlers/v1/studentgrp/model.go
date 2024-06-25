package studentgrp

import (
	"fmt"
	"time"

	"github.com/PhyoYazar/uas/business/core/student"
	"github.com/PhyoYazar/uas/business/sys/validate"
)

// AppStudent represents information about an individual student.
type AppStudent struct {
	ID           string `json:"id"`
	StudentName  string `json:"studentName"`
	Year         string `json:"year"`
	AcademicYear string `json:"academicYear"`
	RollNumber   int    `json:"rollNumber"`
	DateCreated  string `json:"dateCreated"`
	DateUpdated  string `json:"dateUpdated"`
}

func toAppStudent(std student.Student) AppStudent {

	return AppStudent{
		ID:           std.ID.String(),
		StudentName:  std.StudentName,
		Year:         std.Year.Name(),
		AcademicYear: std.AcademicYear,
		RollNumber:   std.RollNumber,
		DateCreated:  std.DateCreated.Format(time.RFC3339),
		DateUpdated:  std.DateUpdated.Format(time.RFC3339),
	}
}

// =============================================================================

// AppNewStudent contains information needed to create a new student.
type AppNewStudent struct {
	StudentName  string `json:"studentName" validate:"required"`
	Year         string `json:"year" validate:"required"`
	AcademicYear string `json:"academicYear" validate:"required"`
	RollNumber   int    `json:"rollNumber" validate:"required"`
}

func toCoreNewStudent(app AppNewStudent) (student.NewStudent, error) {

	year, err := student.ParseYear(app.Year)
	if err != nil {
		return student.NewStudent{}, fmt.Errorf("error parsing year: %v", err)
	}

	std := student.NewStudent{
		StudentName:  app.StudentName,
		Year:         year,
		AcademicYear: app.AcademicYear,
		RollNumber:   app.RollNumber,
	}

	return std, nil
}

// Validate checks the data in the model is considered clean.
func (app AppNewStudent) Validate() error {
	if err := validate.Check(app); err != nil {
		return err
	}
	return nil
}

// =============================================================================

// AppUpdateStudent contains information needed to update a student.
type AppUpdateStudent struct {
	StudentName  *string      `json:"studentName"`
	RollNumber   *int         `json:"rollNumber"`
	Year         student.Year `json:"year"`
	AcademicYear *string      `json:"academicYear"`
}

func toCoreUpdateStudent(app AppUpdateStudent) (student.UpdateStudent, error) {

	nSub := student.UpdateStudent{
		StudentName:  app.StudentName,
		RollNumber:   app.RollNumber,
		Year:         app.Year,
		AcademicYear: app.AcademicYear,
	}

	return nSub, nil
}

// Validate checks the data in the model is considered clean.
func (app AppUpdateStudent) Validate() error {
	if err := validate.Check(app); err != nil {
		return fmt.Errorf("validate: %w", err)
	}
	return nil
}
