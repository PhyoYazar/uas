package cogrp

import (
	"net/http"
	"strconv"

	"github.com/PhyoYazar/uas/business/core/co"
	"github.com/PhyoYazar/uas/business/sys/validate"
	"github.com/google/uuid"
)

func parseFilter(r *http.Request) (co.QueryFilter, error) {
	values := r.URL.Query()

	var filter co.QueryFilter

	if coId := values.Get("co_id"); coId != "" {
		id, err := uuid.Parse(coId)
		if err != nil {
			return co.QueryFilter{}, validate.NewFieldsError("co_id", err)
		}
		filter.WithCoID(id)
	}

	if name := values.Get("name"); name != "" {
		filter.WithName(name)
	}

	if instance := values.Get("instance"); instance != "" {
		inst, err := strconv.ParseInt(instance, 10, 64)
		if err != nil {
			return co.QueryFilter{}, validate.NewFieldsError("instance", err)
		}
		filter.WithInstance(int(inst))
	}

	if mark := values.Get("mark"); mark != "" {
		mark, err := strconv.ParseInt(mark, 10, 64)
		if err != nil {
			return co.QueryFilter{}, validate.NewFieldsError("mark", err)
		}
		filter.WithInstance(int(mark))
	}

	if err := filter.Validate(); err != nil {
		return co.QueryFilter{}, err
	}

	return filter, nil
}

// =============================================================================
