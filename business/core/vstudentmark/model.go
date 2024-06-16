package vstudentmark

import "github.com/google/uuid"

type VAttributes struct {
	StudentMarkID uuid.UUID
	AttributeID   uuid.UUID
	// Name          string
	// Instance      string
	// Type          string
	FullMark int
}

type VStudentMark struct {
	ID          uuid.UUID
	RollNumber  int
	StudentName string
	Attributes  []VAttributes
}

type VRemoveStudent struct {
	SubjectID   uuid.UUID
	StudentID   uuid.UUID
	AttributeID uuid.UUID
}
