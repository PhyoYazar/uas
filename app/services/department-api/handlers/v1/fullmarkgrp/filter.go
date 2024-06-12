package fullmarkgrp

import (
	"net/http"
	"strconv"

	"github.com/PhyoYazar/uas/business/core/fullmark"
	"github.com/PhyoYazar/uas/business/sys/validate"
	"github.com/google/uuid"
)

func parseFilter(r *http.Request) (fullmark.QueryFilter, error) {
	values := r.URL.Query()

	var filter fullmark.QueryFilter

	if cmId := values.Get("full_mark_id"); cmId != "" {
		id, err := uuid.Parse(cmId)
		if err != nil {
			return fullmark.QueryFilter{}, validate.NewFieldsError("full_mark_id", err)
		}
		filter.WithFullMarkID(id)
	}

	if cmId := values.Get("subject_id"); cmId != "" {
		id, err := uuid.Parse(cmId)
		if err != nil {
			return fullmark.QueryFilter{}, validate.NewFieldsError("subject_id", err)
		}
		filter.WithSubjectID(id)
	}

	if cmId := values.Get("attribute_id"); cmId != "" {
		id, err := uuid.Parse(cmId)
		if err != nil {
			return fullmark.QueryFilter{}, validate.NewFieldsError("attribute_id", err)
		}
		filter.WithAttributeID(id)
	}

	if mk := values.Get("mark"); mk != "" {
		mk, err := strconv.ParseInt(mk, 10, 64)
		if err != nil {
			return fullmark.QueryFilter{}, validate.NewFieldsError("mark", err)
		}
		filter.WithMark(int(mk))
	}

	if err := filter.Validate(); err != nil {
		return fullmark.QueryFilter{}, err
	}

	return filter, nil
}

// =============================================================================
