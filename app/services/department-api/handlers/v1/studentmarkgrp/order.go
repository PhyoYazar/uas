package studentmarkgrp

import (
	"errors"
	"net/http"

	"github.com/PhyoYazar/uas/business/core/studentmark"
	"github.com/PhyoYazar/uas/business/data/order"
	"github.com/PhyoYazar/uas/business/sys/validate"
)

var orderByFields = map[string]struct{}{
	studentmark.OrderByID:        {},
	studentmark.OrderByMark:      {},
	studentmark.OrderBySubjectID: {},
	studentmark.OrderByStudentID: {},
}

func parseOrder(r *http.Request) (order.By, error) {
	orderBy, err := order.Parse(r, studentmark.DefaultOrderBy)
	if err != nil {
		return order.By{}, err
	}

	if _, exists := orderByFields[orderBy.Field]; !exists {
		return order.By{}, validate.NewFieldsError(orderBy.Field, errors.New("order field does not exist"))
	}

	return orderBy, nil
}
