package mark

import (
	"time"

	"github.com/google/uuid"
)

// Mark represents information about an individual Mark.
type Mark struct {
	ID          uuid.UUID
	SubjectID   uuid.UUID
	GaID        uuid.UUID
	AttributeID uuid.UUID
	Mark        int
	GaMark      int
	DateCreated time.Time
	DateUpdated time.Time
}

// NewMark contains information needed to create a new Mark.
type NewMark struct {
	SubjectID   uuid.UUID
	GaID        uuid.UUID
	AttributeID uuid.UUID
	Mark        int
	GaMark      int
}

// UpdateMark contains information needed to update a Mark.
type UpdateMark struct {
	GaMark *int
}
