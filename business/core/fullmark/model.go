package fullmark

import (
	"time"

	"github.com/google/uuid"
)

// FullMark represents information about an individual FullMark.
type FullMark struct {
	ID          uuid.UUID
	SubjectID   uuid.UUID
	AttributeID uuid.UUID
	Mark        int
	DateCreated time.Time
	DateUpdated time.Time
}

// NewFullMark contains information needed to create a new FullMark.
type NewFullMark struct {
	SubjectID   uuid.UUID
	AttributeID uuid.UUID
	Mark        int
}
