package cogadb

import (
	"bytes"
	"strings"

	"github.com/PhyoYazar/uas/business/core/coga"
)

func (s *Store) applyFilter(filter coga.QueryFilter, data map[string]interface{}, buf *bytes.Buffer) {
	var wc []string

	if filter.ID != nil {
		data["co_ga_id"] = *filter.ID
		wc = append(wc, "co_ga_id = :co_ga_id")
	}

	if filter.CoID != nil {
		data["co_id"] = *filter.CoID
		wc = append(wc, "co_id = :co_id")
	}

	if filter.GaID != nil {
		data["ga_id"] = *filter.GaID
		wc = append(wc, "ga_id = :ga_id")
	}

	if len(wc) > 0 {
		buf.WriteString(" WHERE ")
		buf.WriteString(strings.Join(wc, " AND "))
	}
}
