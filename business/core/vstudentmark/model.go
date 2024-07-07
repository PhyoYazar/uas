package vstudentmark

import "github.com/google/uuid"

type VAttributes struct {
	StudentMarkID uuid.UUID `json:"studentMarkId"`
	AttributeID   uuid.UUID `json:"attributeId"`
	Mark          int       `json:"mark"`
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

type VRemoveStudent struct {
	SubjectID   uuid.UUID `json:"subjectId"`
	StudentID   uuid.UUID `json:"studentId"`
	AttributeID uuid.UUID `json:"attributeId"`
}
