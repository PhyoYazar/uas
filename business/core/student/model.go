package student

import (
	"net/mail"
	"time"

	"github.com/google/uuid"
)

// Student represents information about an individual student.
type Student struct {
	ID           uuid.UUID
	Name         string
	Email        mail.Address
	RollNumber   int
	PhoneNumber  string
	Year         string
	AcademicYear string
	DateCreated  time.Time
	DateUpdated  time.Time
}

// NewStudent contains information needed to create a new student.
type NewStudent struct {
	Name         string
	Email        mail.Address
	RollNumber   int
	PhoneNumber  string
	Year         string
	AcademicYear string
}

// UpdateStudent contains information needed to update a student.
type UpdateStudent struct {
	Name         *string
	Email        *mail.Address
	RollNumber   *int
	PhoneNumber  *string
	Year         *string
	AcademicYear *string
}
