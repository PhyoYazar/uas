package attributegrp

import (
	"fmt"
	"time"

	"github.com/PhyoYazar/uas/business/core/attribute"
	"github.com/PhyoYazar/uas/business/sys/validate"
)

// AppAttribute represents information about an individual attribute.
type AppAttribute struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Instance    int    `json:"instance"`
	DateCreated string `json:"dateCreated"`
	DateUpdated string `json:"dateUpdated"`
}

func toAppAttribute(attribute attribute.Attribute) AppAttribute {

	return AppAttribute{
		ID:          attribute.ID.String(),
		Name:        attribute.Name,
		Type:        attribute.Type.Name(),
		Instance:    attribute.Instance,
		DateCreated: attribute.DateCreated.Format(time.RFC3339),
		DateUpdated: attribute.DateUpdated.Format(time.RFC3339),
	}
}

// =============================================================================

// AppNewAttribute contains information needed to create a new attribute.
type AppNewAttribute struct {
	Name     string `json:"name" validate:"required"`
	Type     string `json:"type" validate:"required"`
	Instance int    `json:"instance" validate:"required"`
}

func toCoreNewAttribute(app AppNewAttribute) (attribute.NewAttribute, error) {

	att, err := attribute.ParseType(app.Type)
	if err != nil {
		return attribute.NewAttribute{}, fmt.Errorf("parse type: %w", err)
	}

	attribute := attribute.NewAttribute{
		Name:     app.Name,
		Type:     att,
		Instance: app.Instance,
	}

	return attribute, nil
}

// Validate checks the data in the model is considered clean.
func (app AppNewAttribute) Validate() error {
	if err := validate.Check(app); err != nil {
		return err
	}
	return nil
}

// =============================================================================
