package vsubject

import (
	"github.com/google/uuid"
)

type VGa struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Slug string    `json:"slug"`
}

type VCo struct {
	ID       uuid.UUID `json:"id"`
	Instance int       `json:"instance"`
	Name     string    `json:"name"`
	Ga       []VGa     `json:"ga"`
}

type VSubject struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Code         string    `json:"code"`
	AcademicYear string    `json:"academic_year"`
	Instructor   string    `json:"instructor"`
	Semester     string    `json:"semester"`
	Co           []VCo     `json:"co"`
}
