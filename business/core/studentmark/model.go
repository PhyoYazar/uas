package studentmark

import (
	"time"

	"github.com/google/uuid"
)

// StudentMark represents information about an individual StudentMark.
type StudentMark struct {
	ID          uuid.UUID
	StudentID   uuid.UUID
	SubjectID   uuid.UUID
	AttributeID uuid.UUID
	Mark        int
	DateCreated time.Time
	DateUpdated time.Time
}

// NewStudentMark contains information needed to create a new StudentMark.
type NewStudentMark struct {
	StudentID   uuid.UUID
	SubjectID   uuid.UUID
	AttributeID uuid.UUID
	Mark        int
}

// UpdateStudentMark contains information needed to update a StudentMark.
type UpdateStudentMark struct {
	// StudentID   *uuid.UUID
	// SubjectID   *uuid.UUID
	// AttributeID *uuid.UUID
	Mark *int
}
