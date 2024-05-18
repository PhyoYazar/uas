package studentsubjectgrp

import (
	"net/http"
	"strconv"

	"github.com/PhyoYazar/uas/business/core/studentsubject"
	"github.com/PhyoYazar/uas/business/sys/validate"
	"github.com/google/uuid"
)

func parseFilter(r *http.Request) (studentsubject.QueryFilter, error) {
	values := r.URL.Query()

	var filter studentsubject.QueryFilter

	if ssId := values.Get("student_subject_id"); ssId != "" {
		id, err := uuid.Parse(ssId)
		if err != nil {
			return studentsubject.QueryFilter{}, validate.NewFieldsError("student_subject_id", err)
		}
		filter.WithStudentSubjectID(id)
	}

	if ssId := values.Get("student_id"); ssId != "" {
		id, err := uuid.Parse(ssId)
		if err != nil {
			return studentsubject.QueryFilter{}, validate.NewFieldsError("student_id", err)
		}
		filter.WithStudentID(id)
	}

	if ssId := values.Get("subject_id"); ssId != "" {
		id, err := uuid.Parse(ssId)
		if err != nil {
			return studentsubject.QueryFilter{}, validate.NewFieldsError("subject_id", err)
		}
		filter.WithSubjectID(id)
	}

	if mark := values.Get("mark"); mark != "" {
		mk, err := strconv.ParseInt(mark, 10, 64)
		if err != nil {
			return studentsubject.QueryFilter{}, validate.NewFieldsError("mark", err)
		}
		filter.WithMark(int(mk))
	}

	if err := filter.Validate(); err != nil {
		return studentsubject.QueryFilter{}, err
	}

	return filter, nil
}

// =============================================================================
