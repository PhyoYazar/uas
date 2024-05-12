package subjectgrp

import (
	"time"

	"github.com/PhyoYazar/uas/business/core/subject"
	"github.com/PhyoYazar/uas/business/sys/validate"
)

// AppUser represents information about an individual user.
type AppSubject struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Code         string `json:"code"`
	Year         string `json:"year"`
	AcademicYear string `json:"academicYear"`
	Instructor   string `json:"instructor"`
	Exam         int    `json:"exam"`
	Practical    int    `json:"practical"`
	DateCreated  string `json:"dateCreated"`
	DateUpdated  string `json:"dateUpdated"`
}

func toAppSubject(sub subject.Subject) AppSubject {

	return AppSubject{
		ID:           sub.ID.String(),
		Name:         sub.Name,
		Code:         sub.Code,
		Year:         sub.Year,
		AcademicYear: sub.AcademicYear,
		Instructor:   sub.Instructor,
		Exam:         sub.Exam,
		Practical:    sub.Practical,
		DateCreated:  sub.DateCreated.Format(time.RFC3339),
		DateUpdated:  sub.DateUpdated.Format(time.RFC3339),
	}
}

// =============================================================================

// AppNewSubject contains information needed to create a new subject.
type AppNewSubject struct {
	Name         string `json:"name" validate:"required"`
	Code         string `json:"code" validate:"required"`
	Year         string `json:"year" validate:"required"`
	AcademicYear string `json:"academicYear" validate:"required"`
	Instructor   string `json:"instructor" validate:"required"`
	Exam         int    `json:"exam" validate:"required"`
}

func toCoreNewSubject(app AppNewSubject) (subject.NewSubject, error) {

	sub := subject.NewSubject{
		Name:         app.Name,
		Code:         app.Code,
		Year:         app.Year,
		AcademicYear: app.AcademicYear,
		Instructor:   app.Instructor,
		Exam:         app.Exam,
	}

	return sub, nil
}

// Validate checks the data in the model is considered clean.
func (app AppNewSubject) Validate() error {
	if err := validate.Check(app); err != nil {
		return err
	}
	return nil
}

// =============================================================================
