package vgagradegrp

import (
	"net/http"

	"github.com/PhyoYazar/uas/business/core/vgagrade"
	"github.com/PhyoYazar/uas/business/sys/validate"
	"github.com/google/uuid"
)

func parseFilter(r *http.Request) (vgagrade.QueryFilter, error) {
	values := r.URL.Query()

	var filter vgagrade.QueryFilter

	if studentID := values.Get("student_id"); studentID != "" {
		id, err := uuid.Parse(studentID)
		if err != nil {
			return vgagrade.QueryFilter{}, validate.NewFieldsError("student_id", err)
		}
		filter.WithStudentID(id)
	}

	if subjectID := values.Get("subject_id"); subjectID != "" {
		id, err := uuid.Parse(subjectID)
		if err != nil {
			return vgagrade.QueryFilter{}, validate.NewFieldsError("subject_id", err)
		}
		filter.WithSubjectID(id)
	}

	if year := values.Get("year"); year != "" {
		filter.WithYear(year)
	}

	if academicYear := values.Get("academic_year"); academicYear != "" {
		filter.WithAcademicYear(academicYear)
	}
	if err := filter.Validate(); err != nil {
		return vgagrade.QueryFilter{}, err
	}

	return filter, nil
}

// =============================================================================
