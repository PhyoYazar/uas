package subject

import (
	"time"

	"github.com/google/uuid"
)

// Subject represents information about an individual subject.
type Subject struct {
	ID           uuid.UUID
	Name         string
	Code         string
	Year         string
	AcademicYear string
	Semester     string
	Instructor   string
	Exam         int
	Practical    int
	DateCreated  time.Time
	DateUpdated  time.Time
}

// NewSubject contains information needed to create a new subject.
type NewSubject struct {
	Name         string
	Code         string
	Year         string
	AcademicYear string
	Semester     string
	Instructor   string
	Exam         int
}

// UpdateSubject contains information needed to update a subject.
type UpdateSubject struct {
	Name         *string
	Code         *string
	Year         *string
	AcademicYear *string
	Instructor   *string
	Semester     *string
	Exam         *int
	Practical    *int
}
