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
	DateCreated string `json:"dateCreated"`
	DateUpdated string `json:"dateUpdated"`
}

func toAppCoAttribute(mark coattribute.CoAttribute) AppCoAttribute {

	return AppCoAttribute{
		ID:          mark.ID.String(),
		CoID:        mark.CoID.String(),
		AttributeID: mark.AttributeID.String(),
		DateCreated: mark.DateCreated.Format(time.RFC3339),
		DateUpdated: mark.DateUpdated.Format(time.RFC3339),
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
