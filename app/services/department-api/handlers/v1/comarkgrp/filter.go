package comarkgrp

import (
	"net/http"

	"github.com/PhyoYazar/uas/business/core/comark"
	"github.com/PhyoYazar/uas/business/sys/validate"
	"github.com/google/uuid"
)

func parseFilter(r *http.Request) (comark.QueryFilter, error) {
	values := r.URL.Query()

	var filter comark.QueryFilter

	if cmId := values.Get("co_mark_id"); cmId != "" {
		id, err := uuid.Parse(cmId)
		if err != nil {
			return comark.QueryFilter{}, validate.NewFieldsError("co_mark_id", err)
		}
		filter.WithCoMarkID(id)
	}

	if cmId := values.Get("co_id"); cmId != "" {
		id, err := uuid.Parse(cmId)
		if err != nil {
			return comark.QueryFilter{}, validate.NewFieldsError("co_id", err)
		}
		filter.WithCoID(id)
	}

	if cmId := values.Get("mark_id"); cmId != "" {
		id, err := uuid.Parse(cmId)
		if err != nil {
			return comark.QueryFilter{}, validate.NewFieldsError("mark_id", err)
		}
		filter.WithMarkID(id)
	}

	if err := filter.Validate(); err != nil {
		return comark.QueryFilter{}, err
	}

	return filter, nil
}

// =============================================================================
