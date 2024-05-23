package co

import (
	"time"

	"github.com/google/uuid"
)

// Co represents information about an individual Co.
type Co struct {
	ID          uuid.UUID
	SubjectID   uuid.UUID
	Name        string
	Instance    int
	DateCreated time.Time
	DateUpdated time.Time
}

// NewCo contains information needed to create a new Co.
type NewCo struct {
	Name      string
	SubjectID uuid.UUID
	Instance  int
}

// UpdateCo contains information needed to update a Co.
type UpdateCo struct {
	Name      *string
	SubjectID uuid.UUID
	Instance  *int
}
