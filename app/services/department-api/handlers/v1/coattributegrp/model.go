package coattributegrp

import (
	"fmt"
	"time"

	"github.com/PhyoYazar/uas/business/core/coattribute"
	"github.com/PhyoYazar/uas/business/sys/validate"
	"github.com/google/uuid"
)

// AppCoGa represents information about an individual coga.
type AppCoAttribute struct {
	ID          string `json:"id"`
	CoID        string `json:"coID"`
	AttributeID string `json:"attributeID"`
	CoMark      int    `json:"coMark"`
	DateCreated string `json:"dateCreated"`
	DateUpdated string `json:"dateUpdated"`
}

func toAppCoAttribute(ca coattribute.CoAttribute) AppCoAttribute {

	return AppCoAttribute{
		ID:          ca.ID.String(),
		CoID:        ca.CoID.String(),
		AttributeID: ca.AttributeID.String(),
		CoMark:      ca.CoMark,
		DateCreated: ca.DateCreated.Format(time.RFC3339),
		DateUpdated: ca.DateUpdated.Format(time.RFC3339),
	}
}

// =============================================================================

// AppNewCoCoAttribute contains information needed to create a new coga.
type AppNewCoAttribute struct {
	CoID        string `json:"coID" validate:"required"`
	AttributeID string `json:"attributeID" validate:"required"`
}

func toCoreNewCoAttribute(app AppNewCoAttribute) (coattribute.NewCoAttribute, error) {

	var err error
	coID, err := uuid.Parse(app.CoID)
	if err != nil {
		return coattribute.NewCoAttribute{}, fmt.Errorf("error parsing coid string to uuid: %w", err)

	}

	attID, err := uuid.Parse(app.AttributeID)
	if err != nil {
		return coattribute.NewCoAttribute{}, fmt.Errorf("error parsing attributeID string string to uuid: %w", err)

	}

	cg := coattribute.NewCoAttribute{
		CoID:        coID,
		AttributeID: attID,
	}

	return cg, nil
}

// Validate checks the data in the model is considered clean.
func (app AppNewCoAttribute) Validate() error {
	if err := validate.Check(app); err != nil {
		return err
	}
	return nil
}

// =============================================================================

// AppUpdateStudent contains information needed to update a student.
type AppUpdateCoAttribute struct {
	CoMark *int `json:"coMark"`
}

func toCoreUpdateCoAttribute(app AppUpdateCoAttribute) (coattribute.UpdateCoAttribute, error) {

	ca := coattribute.UpdateCoAttribute{
		CoMark: app.CoMark,
	}

	return ca, nil
}

// Validate checks the data in the model is considered clean.
func (app AppUpdateCoAttribute) Validate() error {
	if err := validate.Check(app); err != nil {
		return fmt.Errorf("validate: %w", err)
	}
	return nil
}
