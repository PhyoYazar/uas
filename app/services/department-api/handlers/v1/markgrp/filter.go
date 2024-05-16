package markgrp

import (
	"net/http"

	"github.com/PhyoYazar/uas/business/core/mark"
	"github.com/PhyoYazar/uas/business/sys/validate"
	"github.com/google/uuid"
)

func parseFilter(r *http.Request) (mark.QueryFilter, error) {
	values := r.URL.Query()

	var filter mark.QueryFilter

	if markId := values.Get("mark_id"); markId != "" {
		id, err := uuid.Parse(markId)
		if err != nil {
			return mark.QueryFilter{}, validate.NewFieldsError("mark_id", err)
		}
		filter.WithMarkID(id)
	}

	if name := values.Get("name"); name != "" {
		filter.WithName(name)
	}

	if markTyp := values.Get("type"); markTyp != "" {
		typ, err := mark.ParseType(markTyp)
		if err != nil {
			return mark.QueryFilter{}, validate.NewFieldsError("type", err)
		}
		filter.WithMarkType(typ)
	}

	if err := filter.Validate(); err != nil {
		return mark.QueryFilter{}, err
	}

	return filter, nil
}

// =============================================================================
