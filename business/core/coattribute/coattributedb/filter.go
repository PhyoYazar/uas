package coattributedb

import (
	"bytes"
	"strings"

	"github.com/PhyoYazar/uas/business/core/coattribute"
)

func (s *Store) applyFilter(filter coattribute.QueryFilter, data map[string]interface{}, buf *bytes.Buffer) {
	var wc []string

	if filter.ID != nil {
		data["co_attribute_id"] = *filter.ID
		wc = append(wc, "co_attribute_id = :co_attribute_id")
	}

	if filter.CoID != nil {
		data["co_id"] = *filter.CoID
		wc = append(wc, "co_id = :co_id")
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
