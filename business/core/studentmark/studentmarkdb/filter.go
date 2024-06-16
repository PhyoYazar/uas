package studentmarkdb

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PhyoYazar/uas/business/core/studentmark"
)

func (s *Store) applyFilter(filter studentmark.QueryFilter, data map[string]interface{}, buf *bytes.Buffer) {
	var wc []string

	if filter.ID != nil {
		data["student_mark_id"] = *filter.ID
		wc = append(wc, "student_mark_id = :student_mark_id")
	}

	if filter.Mark != nil {
		data["mark"] = fmt.Sprintf("%%%d%%", *filter.Mark)
		wc = append(wc, "mark = :mark")
	}

	if filter.StudentID != nil {
		data["student_id"] = *filter.StudentID
		wc = append(wc, "student_id = :student_id")
	}

	if filter.SubjectID != nil {
		data["subject_id"] = *filter.SubjectID
		wc = append(wc, "subject_id = :subject_id")
	}

	if filter.AttributeID != nil {
		data["attribute_id"] = *filter.AttributeID
		wc = append(wc, "attribute_id = :attribute_id")
	}

	if len(wc) > 0 {
		buf.WriteString(" WHERE ")
		buf.WriteString(strings.Join(wc, " AND "))
	}
}
