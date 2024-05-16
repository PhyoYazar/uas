package markdb

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PhyoYazar/uas/business/core/mark"
)

func (s *Store) applyFilter(filter mark.QueryFilter, data map[string]interface{}, buf *bytes.Buffer) {
	var wc []string

	if filter.ID != nil {
		data["mark_id"] = *filter.ID
		wc = append(wc, "mark_id = :mark_id")
	}

	if filter.Name != nil {
		data["name"] = fmt.Sprintf("%%%s%%", *filter.Name)
		wc = append(wc, "name LIKE :name")
	}

	if filter.Type != nil {
		data["type"] = filter.Type.Name()
		wc = append(wc, "type = :type")
	}

	if len(wc) > 0 {
		buf.WriteString(" WHERE ")
		buf.WriteString(strings.Join(wc, " AND "))
	}
}
