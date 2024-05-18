package coga

import (
	"time"

	"github.com/google/uuid"
)

// CoGa represents information about an individual CoGa.
type CoGa struct {
	ID          uuid.UUID
	CoID        uuid.UUID
	GaID        uuid.UUID
	Mark        int
	DateCreated time.Time
	DateUpdated time.Time
}

// NewCoGa contains information needed to create a new CoGa.
type NewCoGa struct {
	CoID uuid.UUID
	GaID uuid.UUID
	Mark int
}

// UpdateStudentSubject contains information needed to update a CoGa.
type UpdateCoGa struct {
	CoID *uuid.UUID
	GaID *uuid.UUID
	Mark *int
}
