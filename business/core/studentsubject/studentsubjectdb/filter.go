package studentsubjectdb

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PhyoYazar/uas/business/core/studentsubject"
)

func (s *Store) applyFilter(filter studentsubject.QueryFilter, data map[string]interface{}, buf *bytes.Buffer) {
	var wc []string

	if filter.ID != nil {
		data["student_subject_id"] = *filter.ID
		wc = append(wc, "student_subject_id = :student_subject_id")
	}

	if filter.Mark != nil {
		data["mark"] = fmt.Sprintf("%%%d%%", *filter.Mark)
		wc = append(wc, "mark = :,mark")
	}

	if filter.StudentID != nil {
		data["student_id"] = *filter.StudentID
		wc = append(wc, "student_id = :student_id")
	}

	if filter.SubjectID != nil {
		data["subject_id"] = *filter.SubjectID
		wc = append(wc, "subject_id = :subject_id")
	}

	if len(wc) > 0 {
		buf.WriteString(" WHERE ")
		buf.WriteString(strings.Join(wc, " AND "))
	}
}
