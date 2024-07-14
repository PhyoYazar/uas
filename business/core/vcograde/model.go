package vcograde

import "github.com/google/uuid"

type VCo struct {
	CoID           uuid.UUID `json:"coId"`
	CoName         string    `json:"coName"`
	CoInstance     int       `json:"coInstance"`
	TotalFullMarks int       `json:"totalFullMarks"`
	TotalMarks     float64   `json:"totalMarks"`
}

type VStudentMark struct {
	ID          uuid.UUID `json:"id"`
	RollNumber  int       `json:"rollNumber"`
	StudentName string    `json:"studentName"`
	Co          []VCo     `json:"co"`
}
