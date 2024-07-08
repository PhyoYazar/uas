package vgagradedb

import (
	"bytes"
	"strings"

	"github.com/PhyoYazar/uas/business/core/vgagrade"
)

func (s *Store) applyFilter(filter vgagrade.QueryFilter, data map[string]interface{}, buf *bytes.Buffer) {
	var wc []string

	if filter.ID != nil {
		data["student_id"] = *filter.ID
		wc = append(wc, "student_id = :student_id")
	}

	if filter.Year != nil {
		data["year"] = *filter.Year
		wc = append(wc, "year = :year")
	}

	if filter.AcademicYear != nil {
		data["academic_year"] = *filter.AcademicYear
		wc = append(wc, "academic_year = :academic_year")
	}

	if len(wc) > 0 {
		buf.WriteString(" WHERE ")
		buf.WriteString(strings.Join(wc, " AND "))
		buf.WriteString(" GROUP BY s.student_id, s.student_name, s.roll_number, ga.ga_id, ga.slug ")
	}
}
