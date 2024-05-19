package attribute

import (
	"time"

	"github.com/google/uuid"
)

// Attribute represents information about an individual Attribute.
type Attribute struct {
	ID          uuid.UUID
	Name        string
	Type        Type
	Instance    int
	DateCreated time.Time
	DateUpdated time.Time
}

// NewAttribute contains information needed to create a new Attribute.
type NewAttribute struct {
	Name     string
	Type     Type
	Instance int
}

// UpdateAttribute contains information needed to update a Attribute.
type UpdateAttribute struct {
	Name     *string
	Type     *Type
	Instance *int
}
