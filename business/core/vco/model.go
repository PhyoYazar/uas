package vco

import (
	"time"

	"github.com/google/uuid"
)

type VGa struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Slug string    `json:"slug"`
}

type VSubject struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Code         string    `json:"code"`
	AcademicYear string    `json:"academic_year"`
	Instructor   string    `json:"instructor"`
	Semester     string    `json:"semester"`
}

type VCo struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Instance    int       `json:"instance"`
	Subject     VSubject  `json:"subject"`
	Ga          []VGa     `json:"ga"`
	DateCreated time.Time `json:"date_created"`
	DateUpdated time.Time `json:"date_updated"`
}

type CourseOutlineRow struct {
	SubjectID    uuid.UUID
	SubjectName  string
	SubjectCode  string
	AcademicYear string
	Instructor   string
	Semester     string
	CoID         uuid.UUID
	CoName       string
	CoInstance   int
	GaID         uuid.UUID
	GaName       string
	GaSlug       string
	DateCreated  time.Time
	DateUpdated  time.Time
}
