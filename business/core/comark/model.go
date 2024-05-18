package comark

import (
	"time"

	"github.com/google/uuid"
)

// CoMark represents information about an individual CoMark.
type CoMark struct {
	ID          uuid.UUID
	CoID        uuid.UUID
	MarkID      uuid.UUID
	DateCreated time.Time
	DateUpdated time.Time
}

// NewCoMark contains information needed to create a new CoMark.
type NewCoMark struct {
	CoID   uuid.UUID
	MarkID uuid.UUID
}

// UpdateCoMark contains information needed to update a CoMark.
type UpdateCoMark struct {
	CoID   *uuid.UUID
	MarkID *uuid.UUID
}
