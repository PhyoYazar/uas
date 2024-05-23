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
