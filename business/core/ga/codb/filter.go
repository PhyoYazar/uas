package gadb

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PhyoYazar/uas/business/core/ga"
)

func (s *Store) applyFilter(filter ga.QueryFilter, data map[string]interface{}, buf *bytes.Buffer) {
	var wc []string

	if filter.ID != nil {
		data["ga_id"] = *filter.ID
		wc = append(wc, "ga_id = :ga_id")
	}

	if filter.Name != nil {
		data["name"] = fmt.Sprintf("%%%s%%", *filter.Name)
		wc = append(wc, "name LIKE :name")
	}

	if filter.Slug != nil {
		data["slug"] = fmt.Sprintf("%%%s%%", *filter.Name)
		wc = append(wc, "slug = :slug")
	}

	if len(wc) > 0 {
		buf.WriteString(" WHERE ")
		buf.WriteString(strings.Join(wc, " AND "))
	}
}
