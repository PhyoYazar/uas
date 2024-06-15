package vattribute

import (
	"time"

	"github.com/google/uuid"
)

type VGa struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Slug string    `json:"slug"`
}

type VCo struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Instance int       `json:"instance"`
}

// Implement the GetID method for VCo
func (co VCo) GetID() interface{} {
	return co.ID
}

type VAttribute struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Instance    int       `json:"instance"`
	Type        string    `json:"type"`
	DateCreated time.Time `json:"date_created"`
	DateUpdated time.Time `json:"date_updated"`
	Co          []VCo     `json:"co"`
	Ga          []VGa     `json:"ga"`
}

//================================================================

type VMark struct {
	ID     uuid.UUID `json:"id"`
	Mark   int       `json:"mark"`
	GaID   uuid.UUID `json:"gaID"`
	GaSlug string    `json:"gaSlug"`
}

// Implement the GetID method for VMark
func (mark VMark) GetID() interface{} {
	return mark.ID
}

type VAttributeWithGaMark struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Instance int       `json:"instance"`
	Type     string    `json:"type"`
	FullMark int       `json:"fullMark"`
	Marks    []VMark   `json:"marks"`
	Co       []VCo     `json:"co"`
}

//================================================================

type VRemoveAttribute struct {
	SubjectID   uuid.UUID `json:"subject_id"`
	AttributeID uuid.UUID `json:"attribute_id"`
}
