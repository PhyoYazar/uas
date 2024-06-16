package studentmarkgrp

import (
	"net/http"
	"strconv"

	"github.com/PhyoYazar/uas/business/core/studentmark"
	"github.com/PhyoYazar/uas/business/sys/validate"
	"github.com/google/uuid"
)

func parseFilter(r *http.Request) (studentmark.QueryFilter, error) {
	values := r.URL.Query()

	var filter studentmark.QueryFilter

	if ssId := values.Get("student_subject_id"); ssId != "" {
		id, err := uuid.Parse(ssId)
		if err != nil {
			return studentmark.QueryFilter{}, validate.NewFieldsError("student_subject_id", err)
		}
		filter.WithStudentMarkID(id)
	}

	if ssId := values.Get("student_id"); ssId != "" {
		id, err := uuid.Parse(ssId)
		if err != nil {
			return studentmark.QueryFilter{}, validate.NewFieldsError("student_id", err)
		}
		filter.WithStudentID(id)
	}

	if ssId := values.Get("subject_id"); ssId != "" {
		id, err := uuid.Parse(ssId)
		if err != nil {
			return studentmark.QueryFilter{}, validate.NewFieldsError("subject_id", err)
		}
		filter.WithSubjectID(id)
	}

	if ssId := values.Get("attribute_id"); ssId != "" {
		id, err := uuid.Parse(ssId)
		if err != nil {
			return studentmark.QueryFilter{}, validate.NewFieldsError("attribute_id", err)
		}
		filter.WithAttributeID(id)
	}

	if mark := values.Get("mark"); mark != "" {
		mk, err := strconv.ParseInt(mark, 10, 64)
		if err != nil {
			return studentmark.QueryFilter{}, validate.NewFieldsError("mark", err)
		}
		filter.WithMark(int(mk))
	}

	if err := filter.Validate(); err != nil {
		return studentmark.QueryFilter{}, err
	}

	return filter, nil
}

// =============================================================================
