package coattributegrp

import (
	"net/http"

	"github.com/PhyoYazar/uas/business/core/coattribute"
	"github.com/PhyoYazar/uas/business/sys/validate"
	"github.com/google/uuid"
)

func parseFilter(r *http.Request) (coattribute.QueryFilter, error) {
	values := r.URL.Query()

	var filter coattribute.QueryFilter

	if cgId := values.Get("co_attribute_id"); cgId != "" {
		id, err := uuid.Parse(cgId)
		if err != nil {
			return coattribute.QueryFilter{}, validate.NewFieldsError("co_attribute_id", err)
		}
		filter.WithCoAttributeID(id)
	}

	if cgId := values.Get("co_id"); cgId != "" {
		id, err := uuid.Parse(cgId)
		if err != nil {
			return coattribute.QueryFilter{}, validate.NewFieldsError("co_id", err)
		}
		filter.WithCoID(id)
	}

	if cgId := values.Get("attribute_id"); cgId != "" {
		id, err := uuid.Parse(cgId)
		if err != nil {
			return coattribute.QueryFilter{}, validate.NewFieldsError("attribute_id", err)
		}
		filter.WithAttributeID(id)
	}

	if err := filter.Validate(); err != nil {
		return coattribute.QueryFilter{}, err
	}

	return filter, nil
}

// =============================================================================
