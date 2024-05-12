package studentgrp

import (
	"fmt"
	"net/mail"
	"time"

	"github.com/PhyoYazar/uas/business/core/student"
	"github.com/PhyoYazar/uas/business/sys/validate"
)

// AppStudent represents information about an individual student.
type AppStudent struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Year         string `json:"year"`
	AcademicYear string `json:"academicYear"`
	PhoneNumber  string `json:"phoneNumber"`
	RollNumber   int    `json:"rollNumber"`
	DateCreated  string `json:"dateCreated"`
	DateUpdated  string `json:"dateUpdated"`
}

func toAppStudent(std student.Student) AppStudent {

	return AppStudent{
		ID:           std.ID.String(),
		Name:         std.Name,
		Email:        std.Email.Address,
		Year:         std.Year,
		AcademicYear: std.AcademicYear,
		PhoneNumber:  std.PhoneNumber,
		RollNumber:   std.RollNumber,
		DateCreated:  std.DateCreated.Format(time.RFC3339),
		DateUpdated:  std.DateUpdated.Format(time.RFC3339),
	}
}

// =============================================================================

// AppNewStudent contains information needed to create a new student.
type AppNewStudent struct {
	Name         string `json:"name" validate:"required"`
	Email        string `json:"email" validate:"required,email"`
	Year         string `json:"year" validate:"required"`
	AcademicYear string `json:"academicYear" validate:"required"`
	PhoneNumber  string `json:"phoneNumber" validate:"required"`
	RollNumber   int    `json:"rollNumber" validate:"required"`
}

func toCoreNewStudent(app AppNewStudent) (student.NewStudent, error) {

	addr, err := mail.ParseAddress(app.Email)
	if err != nil {
		return student.NewStudent{}, fmt.Errorf("parsing email: %w", err)
	}

	std := student.NewStudent{
		Name:         app.Name,
		Email:        *addr,
		Year:         app.Year,
		AcademicYear: app.AcademicYear,
		PhoneNumber:  app.PhoneNumber,
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
