package cogrp

import (
	"net/http"

	"github.com/PhyoYazar/uas/business/core/co"
	"github.com/PhyoYazar/uas/business/sys/validate"
	"github.com/google/uuid"
)

func parseFilter(r *http.Request) (co.QueryFilter, error) {
	values := r.URL.Query()

	var filter co.QueryFilter

	if coId := values.Get("co_id"); coId != "" {
		id, err := uuid.Parse(coId)
		if err != nil {
			return co.QueryFilter{}, validate.NewFieldsError("co_id", err)
		}
		filter.WithCoID(id)
	}

	if name := values.Get("name"); name != "" {
		filter.WithName(name)
	}

	if err := filter.Validate(); err != nil {
		return co.QueryFilter{}, err
	}

	return filter, nil
}

// =============================================================================
