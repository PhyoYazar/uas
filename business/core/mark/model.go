package mark

import (
	"time"

	"github.com/google/uuid"
)

// Mark represents information about an individual Mark.
type Mark struct {
	ID          uuid.UUID
	Name        string
	Type        Type
	Instance    int
	DateCreated time.Time
	DateUpdated time.Time
}

// NewMark contains information needed to create a new Mark.
type NewMark struct {
	Name     string
	Type     Type
	Instance int
}

// UpdateMark contains information needed to update a Mark.
type UpdateMark struct {
	Name     *string
	Type     *Type
	Instance *int
}
