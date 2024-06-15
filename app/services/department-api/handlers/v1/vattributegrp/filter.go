package vattributegrp

import (
	"net/http"
	"strconv"

	"github.com/PhyoYazar/uas/business/core/attribute"
	"github.com/PhyoYazar/uas/business/core/vattribute"
	"github.com/PhyoYazar/uas/business/sys/validate"
	"github.com/google/uuid"
)

func parseFilter(r *http.Request) (vattribute.QueryFilter, error) {
	values := r.URL.Query()

	var filter vattribute.QueryFilter

	if attID := values.Get("attribute_id"); attID != "" {
		id, err := uuid.Parse(attID)
		if err != nil {
			return vattribute.QueryFilter{}, validate.NewFieldsError("attribute_id", err)
		}
		filter.WithAttributeID(id)
	}

	if subID := values.Get("subject_id"); subID != "" {
		id, err := uuid.Parse(subID)
		if err != nil {
			return vattribute.QueryFilter{}, validate.NewFieldsError("subject_id", err)
		}
		filter.WithSubjectID(id)
	}

	if name := values.Get("name"); name != "" {
		filter.WithName(name)
	}

	if instance := values.Get("instance"); instance != "" {
		inst, err := strconv.ParseInt(instance, 10, 64)
		if err != nil {
			return vattribute.QueryFilter{}, validate.NewFieldsError("mark", err)
		}
		filter.WithInstance(int(inst))
	}

	if attributeTyp := values.Get("type"); attributeTyp != "" {
		typ, err := attribute.ParseType(attributeTyp)
		if err != nil {
			return vattribute.QueryFilter{}, validate.NewFieldsError("type", err)
		}
		filter.WithAttributeType(typ)
	}

	if err := filter.Validate(); err != nil {
		return vattribute.QueryFilter{}, err
	}

	return filter, nil
}

// =============================================================================
