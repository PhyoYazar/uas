package fullmarkdb

import (
	"bytes"
	"strings"

	"github.com/PhyoYazar/uas/business/core/fullmark"
)

func (s *Store) applyFilter(filter fullmark.QueryFilter, data map[string]interface{}, buf *bytes.Buffer) {
	var wc []string

	if filter.ID != nil {
		data["full_mark_id"] = *filter.ID
		wc = append(wc, "full_mark_id = :full_mark_id")
	}

	if filter.SubjectID != nil {
		data["subject_id"] = *filter.SubjectID
		wc = append(wc, "subject_id = :subject_id")
	}

	if filter.AttributeID != nil {
		data["attribute_id"] = *filter.AttributeID
		wc = append(wc, "attribute_id = :attribute_id")
	}

	if filter.Mark != nil {
		data["mark"] = *filter.Mark
		wc = append(wc, "mark = :mark")
	}

	if len(wc) > 0 {
		buf.WriteString(" WHERE ")
		buf.WriteString(strings.Join(wc, " AND "))
	}
}
