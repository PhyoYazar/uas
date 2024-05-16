package gagrp

import (
	"time"

	"github.com/PhyoYazar/uas/business/core/ga"
	"github.com/PhyoYazar/uas/business/sys/validate"
)

// AppGa represents information about an individual ga.
type AppGa struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	DateCreated string `json:"dateCreated"`
	DateUpdated string `json:"dateUpdated"`
}

func toAppGa(ga ga.Ga) AppGa {

	return AppGa{
		ID:          ga.ID.String(),
		Name:        ga.Name,
		Slug:        ga.Slug,
		DateCreated: ga.DateCreated.Format(time.RFC3339),
		DateUpdated: ga.DateUpdated.Format(time.RFC3339),
	}
}

// =============================================================================

// AppNewGa contains information needed to create a new ga.
type AppNewGa struct {
	Name string `json:"name" validate:"required"`
	Slug string `json:"slug" validate:"required"`
}

func toCoreNewGa(app AppNewGa) (ga.NewGa, error) {

	ga := ga.NewGa{
		Name: app.Name,
		Slug: app.Slug,
	}

	return ga, nil
}

// Validate checks the data in the model is considered clean.
func (app AppNewGa) Validate() error {
	if err := validate.Check(app); err != nil {
		return err
	}
	return nil
}

// =============================================================================
