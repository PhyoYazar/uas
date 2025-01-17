package subjectgrp

import (
	"fmt"
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
	Semester     string `json:"semester"`
	Instructor   string `json:"instructor"`
	Exam         int    `json:"exam"`
	Practical    int    `json:"practical"`
	Tutorial     int    `json:"tutorial"`
	Assignment   int    `json:"assignment"`
	Lab          int    `json:"lab"`
	DateCreated  string `json:"dateCreated"`
	DateUpdated  string `json:"dateUpdated"`
}

func toAppSubject(sub subject.Subject) AppSubject {

	return AppSubject{
		ID:           sub.ID.String(),
		Name:         sub.Name,
		Code:         sub.Code,
		Year:         sub.Year.Name(),
		AcademicYear: sub.AcademicYear,
		Semester:     sub.Semester.Name(),
		Instructor:   sub.Instructor,
		Exam:         sub.Exam,
		Practical:    sub.Practical,
		Lab:          sub.Lab,
		Tutorial:     sub.Tutorial,
		Assignment:   sub.Assignment,
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
	Semester     string `json:"semester" validate:"required"`
	Instructor   string `json:"instructor" validate:"required"`
	Exam         int    `json:"exam" validate:"required"`
	Practical    int    `json:"practical"`
	Tutorial     int    `json:"tutorial"`
	Assignment   int    `json:"assignment"`
	Lab          int    `json:"lab"`
}

func toCoreNewSubject(app AppNewSubject) (subject.NewSubject, error) {

	year, err := subject.ParseYear(app.Year)
	if err != nil {
		return subject.NewSubject{}, fmt.Errorf("error parsing year: %v", err)
	}

	semester, err := subject.ParseSemester(app.Semester)
	if err != nil {
		return subject.NewSubject{}, fmt.Errorf("error parsing semester: %v", err)
	}

	sub := subject.NewSubject{
		Name:         app.Name,
		Code:         app.Code,
		Year:         year,
		AcademicYear: app.AcademicYear,
		Semester:     semester,
		Instructor:   app.Instructor,
		Exam:         app.Exam,
		Lab:          app.Lab,
		Tutorial:     app.Tutorial,
		Assignment:   app.Assignment,
		Practical:    app.Practical,
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

// AppUpdateUser contains information needed to update a user.
type AppUpdateSubject struct {
	Name         *string          `json:"name"`
	Code         *string          `json:"code"`
	Year         subject.Year     `json:"year"`
	AcademicYear *string          `json:"academicYear"`
	Instructor   *string          `json:"instructor"`
	Semester     subject.Semester `json:"semester"`
	Exam         *int             `json:"exam"`
	Practical    *int             `json:"practical"`
	Tutorial     *int             `json:"tutorial"`
	Assignment   *int             `json:"assignment"`
	Lab          *int             `json:"lab"`
}

func toCoreUpdateSubject(app AppUpdateSubject) (subject.UpdateSubject, error) {

	// var year subject.Year
	// if app.Year != nil {
	// 	var err error
	// 	year, err = subject.ParseYear(*app.Year)
	// 	if err != nil {
	// 		return subject.UpdateSubject{}, fmt.Errorf("parse: %w", err)
	// 	}
	// }

	// var semester subject.Semester
	// if app.Semester != nil {
	// 	var err error
	// 	semester, err = subject.ParseSemester(*app.Semester)
	// 	if err != nil {
	// 		return subject.UpdateSubject{}, fmt.Errorf("parse: %w", err)
	// 	}
	// }

	// year, err := subject.ParseYear(app.Year)
	// if err != nil {
	// 	return subject.UpdateSubject{}, fmt.Errorf("error parsing year: %v", err)
	// }

	// semester, err := subject.ParseSemester(app.Semester)
	// if err != nil {
	// 	return subject.UpdateSubject{}, fmt.Errorf("error parsing semester: %v", err)
	// }

	nSub := subject.UpdateSubject{
		Name:         app.Name,
		Code:         app.Code,
		Year:         app.Year,
		AcademicYear: app.AcademicYear,
		Instructor:   app.Instructor,
		Semester:     app.Semester,
		Exam:         app.Exam,
		Practical:    app.Practical,
		Tutorial:     app.Tutorial,
		Assignment:   app.Assignment,
		Lab:          app.Lab,
	}

	return nSub, nil
}

// Validate checks the data in the model is considered clean.
func (app AppUpdateSubject) Validate() error {
	if err := validate.Check(app); err != nil {
		return fmt.Errorf("validate: %w", err)
	}
	return nil
}
