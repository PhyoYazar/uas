package vgagrade

import "github.com/google/uuid"

type VGa struct {
	GaID       uuid.UUID `json:"gaId"`
	GaSlug     string    `json:"gaSlug"`
	TotalMarks float64   `json:"totalMarks"`
}

type VStudentMark struct {
	ID          uuid.UUID `json:"id"`
	RollNumber  int       `json:"rollNumber"`
	StudentName string    `json:"studentName"`
	Ga          []VGa     `json:"ga"`
}
