package studentgrp

import (
	"net/http"
	"net/mail"
	"time"

	"github.com/PhyoYazar/uas/business/core/student"
	"github.com/PhyoYazar/uas/business/sys/validate"
	"github.com/google/uuid"
)

func parseFilter(r *http.Request) (student.QueryFilter, error) {
	values := r.URL.Query()

	var filter student.QueryFilter

	if studentID := values.Get("student_id"); studentID != "" {
		id, err := uuid.Parse(studentID)
		if err != nil {
			return student.QueryFilter{}, validate.NewFieldsError("student_id", err)
		}
		filter.WithStudentID(id)
	}

	if email := values.Get("email"); email != "" {
		addr, err := mail.ParseAddress(email)
		if err != nil {
			return student.QueryFilter{}, validate.NewFieldsError("email", err)
		}
		filter.WithEmail(*addr)
	}

	if createdDate := values.Get("start_created_date"); createdDate != "" {
		t, err := time.Parse(time.RFC3339, createdDate)
		if err != nil {
			return student.QueryFilter{}, validate.NewFieldsError("start_created_date", err)
		}
		filter.WithStartDateCreated(t)
	}

	if createdDate := values.Get("end_created_date"); createdDate != "" {
		t, err := time.Parse(time.RFC3339, createdDate)
		if err != nil {
			return student.QueryFilter{}, validate.NewFieldsError("end_created_date", err)
		}
		filter.WithEndCreatedDate(t)
	}

	if name := values.Get("name"); name != "" {
		filter.WithName(name)
	}

	if year := values.Get("year"); year != "" {
		filter.WithYear(year)
	}

	if academicYear := values.Get("academicYear"); academicYear != "" {
		filter.WithAcademicYear(academicYear)
	}

	if rollNumber := values.Get("roll_number"); rollNumber != "" {
		filter.WithRollNumber(rollNumber)
	}

	if phoneNumber := values.Get("phone_number"); phoneNumber != "" {
		filter.WithPhoneNumber(phoneNumber)
	}

	if err := filter.Validate(); err != nil {
		return student.QueryFilter{}, err
	}

	return filter, nil
}

// =============================================================================
