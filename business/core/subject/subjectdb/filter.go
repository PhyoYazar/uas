package subjectdb

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PhyoYazar/uas/business/core/subject"
)

func (s *Store) applyFilter(filter subject.QueryFilter, data map[string]interface{}, buf *bytes.Buffer) {
	var wc []string

	if filter.ID != nil {
		data["subject_id"] = *filter.ID
		wc = append(wc, "subject_id = :subject_id")
	}

	if filter.Name != nil {
		data["name"] = fmt.Sprintf("%%%s%%", *filter.Name)
		wc = append(wc, "name LIKE :name")
	}

	if filter.Code != nil {
		data["code"] = fmt.Sprintf("%%%s%%", *filter.Code)
		wc = append(wc, "code LIKE :code")
	}

	if filter.Year != nil {
		data["year"] = fmt.Sprintf("%%%s%%", *filter.Year)
		wc = append(wc, "year = :year")
	}

	if filter.Semester != nil {
		data["semester"] = fmt.Sprintf("%%%s%%", *filter.Semester)
		wc = append(wc, "semester = :semester")
	}

	if filter.AcademicYear != nil {
		data["academic_year"] = fmt.Sprintf("%%%s%%", *filter.AcademicYear)
		wc = append(wc, "academic_year = :academic_year")
	}

	if filter.Instructor != nil {
		data["instructor"] = fmt.Sprintf("%%%s%%", *filter.Instructor)
		wc = append(wc, "instructor LIKE :instructor")
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
