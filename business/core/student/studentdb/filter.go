package studentdb

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PhyoYazar/uas/business/core/student"
)

func (s *Store) applyFilter(filter student.QueryFilter, data map[string]interface{}, buf *bytes.Buffer) {
	var wc []string

	if filter.ID != nil {
		data["student_id"] = *filter.ID
		wc = append(wc, "student_id = :student_id")
	}

	if filter.Name != nil {
		data["name"] = fmt.Sprintf("%%%s%%", *filter.Name)
		wc = append(wc, "name LIKE :name")
	}

	if filter.Email != nil {
		data["email"] = (*filter.Email).String()
		wc = append(wc, "email = :email")
	}

	if filter.Year != nil {
		data["year"] = fmt.Sprintf("%%%s%%", *filter.Year)
		wc = append(wc, "year = :year")
	}

	if filter.AcademicYear != nil {
		data["academic_year"] = fmt.Sprintf("%%%s%%", *filter.AcademicYear)
		wc = append(wc, "academic_year = :academic_year")
	}

	if filter.PhoneNumber != nil {
		data["phone_number"] = fmt.Sprintf("%%%s%%", *filter.PhoneNumber)
		wc = append(wc, "phone_number LIKE :phone_number")
	}

	if filter.RollNumber != nil {
		data["roll_number"] = fmt.Sprintf("%%%s%%", *filter.RollNumber)
		wc = append(wc, "roll_number LIKE :roll_number")
	}

	if filter.StartCreatedDate != nil {
		data["start_date_created"] = *filter.StartCreatedDate
		wc = append(wc, "date_created >= :start_date_created")
	}

	if filter.EndCreatedDate != nil {
		data["end_date_created"] = *filter.EndCreatedDate
		wc = append(wc, "date_created <= :end_date_created")
	}

	if len(wc) > 0 {
		buf.WriteString(" WHERE ")
		buf.WriteString(strings.Join(wc, " AND "))
	}
}
