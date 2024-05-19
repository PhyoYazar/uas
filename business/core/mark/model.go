package mark

import (
	"time"

	"github.com/google/uuid"
)

// Mark represents information about an individual Mark.
type Mark struct {
	ID          uuid.UUID
	CoID        uuid.UUID
	GaID        uuid.UUID
	AttributeID uuid.UUID
	Mark        int
	DateCreated time.Time
	DateUpdated time.Time
}

// NewMark contains information needed to create a new Mark.
type NewMark struct {
	CoID        uuid.UUID
	GaID        uuid.UUID
	AttributeID uuid.UUID
	Mark        int
}

// UpdateMark contains information needed to update a Mark.
type UpdateMark struct {
	CoID        *uuid.UUID
	GaID        *uuid.UUID
	AttributeID *uuid.UUID
	Mark        *int
}
