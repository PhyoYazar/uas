package ga

import (
	"time"

	"github.com/google/uuid"
)

// Ga represents information about an individual Ga.
type Ga struct {
	ID          uuid.UUID
	Name        string
	Slug        string
	DateCreated time.Time
	DateUpdated time.Time
}

// NewGa contains information needed to create a new Ga.
type NewGa struct {
	Name string
	Slug string
}

// UpdateGa contains information needed to update a Ga.
type UpdateGa struct {
	Name *string
	Slug *string
}
