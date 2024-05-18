package studentsubject

import (
	"time"

	"github.com/google/uuid"
)

// StudentSubject represents information about an individual StudentSubject.
type StudentSubject struct {
	ID          uuid.UUID
	StudentID   uuid.UUID
	SubjectID   uuid.UUID
	Mark        int
	DateCreated time.Time
	DateUpdated time.Time
}

// NewStudentSubject contains information needed to create a new StudentSubject.
type NewStudentSubject struct {
	StudentID uuid.UUID
	SubjectID uuid.UUID
	Mark      int
}

// UpdateStudentSubject contains information needed to update a StudentSubject.
type UpdateStudentSubject struct {
	StudentID *uuid.UUID
	SubjectID *uuid.UUID
	Mark      *int
}
