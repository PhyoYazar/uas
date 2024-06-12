package student

import (
	"time"

	"github.com/google/uuid"
)

// Student represents information about an individual student.
type Student struct {
	ID           uuid.UUID
	Name         string
	RollNumber   int
	Year         Year
	AcademicYear string
	DateCreated  time.Time
	DateUpdated  time.Time
	// Email        mail.Address
	// PhoneNumber  string
}

// NewStudent contains information needed to create a new student.
type NewStudent struct {
	Name         string
	RollNumber   int
	Year         Year
	AcademicYear string
	// Email        mail.Address
	// PhoneNumber  string
}

// UpdateStudent contains information needed to update a student.
type UpdateStudent struct {
	Name         *string
	RollNumber   *int
	Year         Year
	AcademicYear *string
	// Email        *mail.Address
	// PhoneNumber  *string
}
