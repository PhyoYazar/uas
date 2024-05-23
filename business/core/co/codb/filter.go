package codb

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PhyoYazar/uas/business/core/co"
)

func (s *Store) applyFilter(filter co.QueryFilter, data map[string]interface{}, buf *bytes.Buffer) {
	var wc []string

	if filter.ID != nil {
		data["co_id"] = *filter.ID
		wc = append(wc, "co_id = :co_id")
	}

	if filter.Name != nil {
		data["name"] = fmt.Sprintf("%%%s%%", *filter.Name)
		wc = append(wc, "name LIKE :name")
	}
	if filter.Instance != nil {
		data["instance"] = fmt.Sprintf("%%%d%%", *filter.Instance)
		wc = append(wc, "instance = :instance")
	}

	if len(wc) > 0 {
		buf.WriteString(" WHERE ")
		buf.WriteString(strings.Join(wc, " AND "))
	}

}
