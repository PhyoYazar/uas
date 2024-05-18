package comarkdb

import (
	"bytes"
	"strings"

	"github.com/PhyoYazar/uas/business/core/comark"
)

func (s *Store) applyFilter(filter comark.QueryFilter, data map[string]interface{}, buf *bytes.Buffer) {
	var wc []string

	if filter.ID != nil {
		data["co_mark_id"] = *filter.ID
		wc = append(wc, "co_mark_id = :co_mark_id")
	}

	if filter.CoID != nil {
		data["co_id"] = *filter.CoID
		wc = append(wc, "co_id = :co_id")
	}

	if filter.MarkID != nil {
		data["mark_id"] = *filter.MarkID
		wc = append(wc, "mark_id = :mark_id")
	}

	if len(wc) > 0 {
		buf.WriteString(" WHERE ")
		buf.WriteString(strings.Join(wc, " AND "))
	}
}
