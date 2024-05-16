package gagrp

import (
	"net/http"

	"github.com/PhyoYazar/uas/business/core/ga"
	"github.com/PhyoYazar/uas/business/sys/validate"
	"github.com/google/uuid"
)

func parseFilter(r *http.Request) (ga.QueryFilter, error) {
	values := r.URL.Query()

	var filter ga.QueryFilter

	if gaId := values.Get("ga_id"); gaId != "" {
		id, err := uuid.Parse(gaId)
		if err != nil {
			return ga.QueryFilter{}, validate.NewFieldsError("ga_id", err)
		}
		filter.WithGaID(id)
	}

	if name := values.Get("name"); name != "" {
		filter.WithName(name)
	}

	if slug := values.Get("slug"); slug != "" {
		filter.WithSlug(slug)
	}

	if err := filter.Validate(); err != nil {
		return ga.QueryFilter{}, err
	}

	return filter, nil
}

// =============================================================================
