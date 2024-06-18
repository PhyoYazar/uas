package student

import (
	"time"

	"github.com/google/uuid"
)

// Student represents information about an individual student.
type Student struct {
	ID            uuid.UUID
	StudentNumber int
	RollNumber    int
	Year          Year
	AcademicYear  string
	DateCreated   time.Time
	DateUpdated   time.Time
}

// NewStudent contains information needed to create a new student.
type NewStudent struct {
	RollNumber    int
	StudentNumber int
	Year          Year
	AcademicYear  string
}

// UpdateStudent contains information needed to update a student.
type UpdateStudent struct {
	StudentNumber *int
	RollNumber    *int
	Year          Year
	AcademicYear  *string
}
