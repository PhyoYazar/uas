package attributegrp

import (
	"net/http"

	"github.com/PhyoYazar/uas/business/core/attribute"
	"github.com/PhyoYazar/uas/business/sys/validate"
	"github.com/google/uuid"
)

func parseFilter(r *http.Request) (attribute.QueryFilter, error) {
	values := r.URL.Query()

	var filter attribute.QueryFilter

	if attributeId := values.Get("attribute_id"); attributeId != "" {
		id, err := uuid.Parse(attributeId)
		if err != nil {
			return attribute.QueryFilter{}, validate.NewFieldsError("attribute_id", err)
		}
		filter.WithAttributeID(id)
	}

	if name := values.Get("name"); name != "" {
		filter.WithName(name)
	}

	if attributeTyp := values.Get("type"); attributeTyp != "" {
		typ, err := attribute.ParseType(attributeTyp)
		if err != nil {
			return attribute.QueryFilter{}, validate.NewFieldsError("type", err)
		}
		filter.WithAttributeType(typ)
	}

	if err := filter.Validate(); err != nil {
		return attribute.QueryFilter{}, err
	}

	return filter, nil
}

// =============================================================================
