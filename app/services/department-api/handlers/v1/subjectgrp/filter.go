package subjectgrp

import (
	"net/http"
	"time"

	"github.com/PhyoYazar/uas/business/core/subject"
	"github.com/PhyoYazar/uas/business/sys/validate"
	"github.com/google/uuid"
)

func parseFilter(r *http.Request) (subject.QueryFilter, error) {
	values := r.URL.Query()

	var filter subject.QueryFilter

	if subjectID := values.Get("subject_id"); subjectID != "" {
		id, err := uuid.Parse(subjectID)
		if err != nil {
			return subject.QueryFilter{}, validate.NewFieldsError("subject_id", err)
		}
		filter.WithSubjectID(id)
	}

	if createdDate := values.Get("start_created_date"); createdDate != "" {
		t, err := time.Parse(time.RFC3339, createdDate)
		if err != nil {
			return subject.QueryFilter{}, validate.NewFieldsError("start_created_date", err)
		}
		filter.WithStartDateCreated(t)
	}

	if createdDate := values.Get("end_created_date"); createdDate != "" {
		t, err := time.Parse(time.RFC3339, createdDate)
		if err != nil {
			return subject.QueryFilter{}, validate.NewFieldsError("end_created_date", err)
		}
		filter.WithEndCreatedDate(t)
	}

	if name := values.Get("name"); name != "" {
		filter.WithName(name)
	}

	if code := values.Get("code"); code != "" {
		filter.WithCode(code)
	}

	if year := values.Get("year"); year != "" {
		filter.WithYear(year)
	}

	if academicYear := values.Get("academicYear"); academicYear != "" {
		filter.WithAcademicYear(academicYear)
	}

	if instructor := values.Get("instructor"); instructor != "" {
		filter.WithInstructor(instructor)
	}

	if err := filter.Validate(); err != nil {
		return subject.QueryFilter{}, err
	}

	return filter, nil
}

// =============================================================================
