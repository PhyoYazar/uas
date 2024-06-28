package vattributedb

import (
	"bytes"
	"strings"

	"github.com/PhyoYazar/uas/business/core/vattribute"
)

func (s *Store) applyFilter(filter vattribute.QueryFilter, data map[string]interface{}, buf *bytes.Buffer) {
	var wc []string

	if filter.ID != nil {
		data["attribute_id"] = *filter.ID
		wc = append(wc, "ca.attribute_id = :attribute_id AND m.attribute_id = :attribute_id")
	}

	// if filter.SubID != nil {
	// 	data["subject_id"] = *filter.SubID
	// 	wc = append(wc, "m.subject_id = :subject_id AND co.subject_id = :subject_id") //! Careful comparison
	// }

	if filter.Name != nil {
		data["name"] = *filter.Name
		wc = append(wc, "a.name = :name")
	}

	if filter.Type != nil {
		data["type"] = filter.Type.Name()
		wc = append(wc, "type = :type")
	}

	if len(wc) > 0 {
		// buf.WriteString(" WHERE ")
		buf.WriteString(" AND ")
		buf.WriteString(strings.Join(wc, " AND "))
	}
}
