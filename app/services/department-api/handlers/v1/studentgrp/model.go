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
	Name         string `json:"name"`
	Year         string `json:"year"`
	AcademicYear string `json:"academicYear"`
	RollNumber   int    `json:"rollNumber"`
	DateCreated  string `json:"dateCreated"`
	DateUpdated  string `json:"dateUpdated"`
	// Email        string `json:"email"`
	// PhoneNumber  string `json:"phoneNumber"`
}

func toAppStudent(std student.Student) AppStudent {

	return AppStudent{
		ID:           std.ID.String(),
		Name:         std.Name,
		Year:         std.Year.Name(),
		AcademicYear: std.AcademicYear,
		RollNumber:   std.RollNumber,
		DateCreated:  std.DateCreated.Format(time.RFC3339),
		DateUpdated:  std.DateUpdated.Format(time.RFC3339),
		// Email:        std.Email.Address,
		// PhoneNumber:  std.PhoneNumber,
	}
}

// =============================================================================

// AppNewStudent contains information needed to create a new student.
type AppNewStudent struct {
	Name         string `json:"name" validate:"required"`
	Year         string `json:"year" validate:"required"`
	AcademicYear string `json:"academicYear" validate:"required"`
	RollNumber   int    `json:"rollNumber" validate:"required"`
	// PhoneNumber  string `json:"phoneNumber" validate:"required"`
	// Email        string `json:"email" validate:"required,email"`
}

func toCoreNewStudent(app AppNewStudent) (student.NewStudent, error) {

	// addr, err := mail.ParseAddress(app.Email)
	// if err != nil {
	// 	return student.NewStudent{}, fmt.Errorf("parsing email: %w", err)
	// }

	year, err := student.ParseYear(app.Year)
	if err != nil {
		return student.NewStudent{}, fmt.Errorf("error parsing year: %v", err)
	}

	std := student.NewStudent{
		Name:         app.Name,
		Year:         year,
		AcademicYear: app.AcademicYear,
		RollNumber:   app.RollNumber,
		// Email:        *addr,
		// PhoneNumber:  app.PhoneNumber,
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
	Name         *string      `json:"name"`
	RollNumber   *int         `json:"rollNumber"`
	Year         student.Year `json:"year"`
	AcademicYear *string      `json:"academicYear"`
}

func toCoreUpdateStudent(app AppUpdateStudent) (student.UpdateStudent, error) {

	nSub := student.UpdateStudent{
		Name:         app.Name,
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
