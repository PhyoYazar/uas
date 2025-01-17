package coattribute

import (
	"time"

	"github.com/google/uuid"
)

// CoGa represents information about an individual CoAttribute.
type CoAttribute struct {
	ID          uuid.UUID
	CoID        uuid.UUID
	AttributeID uuid.UUID
	CoMark      int
	DateCreated time.Time
	DateUpdated time.Time
}

// NewCoAttribute contains information needed to create a new CoAttribute.
type NewCoAttribute struct {
	CoID        uuid.UUID
	AttributeID uuid.UUID
	CoMark      int
}

// UpdateCoAttribute contains information needed to update a CoAttribute.
type UpdateCoAttribute struct {
	CoMark *int
}
