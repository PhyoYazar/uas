package vstudentmark

import "github.com/google/uuid"

type VAttributes struct {
	StudentMarkID uuid.UUID `json:"studentMarkId"`
	AttributeID   uuid.UUID `json:"attributeId"`
	Mark          float64   `json:"mark"`
	Name          string    `json:"name"`
	// Instance      string
	// Type          string
}

type VStudentMark struct {
	ID          uuid.UUID     `json:"id"`
	RollNumber  int           `json:"rollNumber"`
	StudentName string        `json:"studentName"`
	Attributes  []VAttributes `json:"attributes"`
}
