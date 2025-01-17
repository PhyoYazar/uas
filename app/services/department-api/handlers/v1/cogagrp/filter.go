package cogagrp

import (
	"net/http"

	"github.com/PhyoYazar/uas/business/core/coga"
	"github.com/PhyoYazar/uas/business/sys/validate"
	"github.com/google/uuid"
)

func parseFilter(r *http.Request) (coga.QueryFilter, error) {
	values := r.URL.Query()

	var filter coga.QueryFilter

	if cgId := values.Get("co_ga_id"); cgId != "" {
		id, err := uuid.Parse(cgId)
		if err != nil {
			return coga.QueryFilter{}, validate.NewFieldsError("co_ga_id", err)
		}
		filter.WithCoGaID(id)
	}

	if cgId := values.Get("co_id"); cgId != "" {
		id, err := uuid.Parse(cgId)
		if err != nil {
			return coga.QueryFilter{}, validate.NewFieldsError("co_id", err)
		}
		filter.WithCoID(id)
	}

	if cgId := values.Get("ga_id"); cgId != "" {
		id, err := uuid.Parse(cgId)
		if err != nil {
			return coga.QueryFilter{}, validate.NewFieldsError("ga_id", err)
		}
		filter.WithGaID(id)
	}

	if err := filter.Validate(); err != nil {
		return coga.QueryFilter{}, err
	}

	return filter, nil
}

// =============================================================================
