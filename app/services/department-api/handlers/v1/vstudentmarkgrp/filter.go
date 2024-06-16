package vstudentmarkgrp

import (
	"net/http"

	"github.com/PhyoYazar/uas/business/core/vstudentmark"
	"github.com/PhyoYazar/uas/business/sys/validate"
	"github.com/google/uuid"
)

func parseFilter(r *http.Request) (vstudentmark.QueryFilter, error) {
	values := r.URL.Query()

	var filter vstudentmark.QueryFilter

	if studentID := values.Get("student_id"); studentID != "" {
		id, err := uuid.Parse(studentID)
		if err != nil {
			return vstudentmark.QueryFilter{}, validate.NewFieldsError("student_id", err)
		}
		filter.WithStudentID(id)
	}

	if year := values.Get("year"); year != "" {
		filter.WithYear(year)
	}

	if academicYear := values.Get("academicYear"); academicYear != "" {
		filter.WithAcademicYear(academicYear)
	}
	if err := filter.Validate(); err != nil {
		return vstudentmark.QueryFilter{}, err
	}

	return filter, nil
}

// =============================================================================
