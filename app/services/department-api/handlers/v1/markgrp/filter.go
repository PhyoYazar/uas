package markgrp

import (
	"net/http"
	"strconv"

	"github.com/PhyoYazar/uas/business/core/mark"
	"github.com/PhyoYazar/uas/business/sys/validate"
	"github.com/google/uuid"
)

func parseFilter(r *http.Request) (mark.QueryFilter, error) {
	values := r.URL.Query()

	var filter mark.QueryFilter

	if cmId := values.Get("mark_id"); cmId != "" {
		id, err := uuid.Parse(cmId)
		if err != nil {
			return mark.QueryFilter{}, validate.NewFieldsError("mark_id", err)
		}
		filter.WithMarkID(id)
	}

	if cmId := values.Get("co_id"); cmId != "" {
		id, err := uuid.Parse(cmId)
		if err != nil {
			return mark.QueryFilter{}, validate.NewFieldsError("co_id", err)
		}
		filter.WithCoID(id)
	}

	if cmId := values.Get("ga_id"); cmId != "" {
		id, err := uuid.Parse(cmId)
		if err != nil {
			return mark.QueryFilter{}, validate.NewFieldsError("ga_id", err)
		}
		filter.WithMarkID(id)
	}

	if cmId := values.Get("attribute_id"); cmId != "" {
		id, err := uuid.Parse(cmId)
		if err != nil {
			return mark.QueryFilter{}, validate.NewFieldsError("attribute_id", err)
		}
		filter.WithAttributeID(id)
	}

	if mk := values.Get("mark"); mk != "" {
		mk, err := strconv.ParseInt(mk, 10, 64)
		if err != nil {
			return mark.QueryFilter{}, validate.NewFieldsError("mark", err)
		}
		filter.WithMark(int(mk))
	}

	if err := filter.Validate(); err != nil {
		return mark.QueryFilter{}, err
	}

	return filter, nil
}

// =============================================================================
