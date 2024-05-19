package markdb

import (
	"bytes"
	"strings"

	"github.com/PhyoYazar/uas/business/core/mark"
)

func (s *Store) applyFilter(filter mark.QueryFilter, data map[string]interface{}, buf *bytes.Buffer) {
	var wc []string

	if filter.ID != nil {
		data["mark_id"] = *filter.ID
		wc = append(wc, "mark_id = :mark_id")
	}

	if filter.CoID != nil {
		data["co_id"] = *filter.CoID
		wc = append(wc, "co_id = :co_id")
	}

	if filter.GaID != nil {
		data["ga_id"] = *filter.GaID
		wc = append(wc, "ga_id = :ga_id")
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
