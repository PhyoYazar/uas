package studentgrp

import (
	"errors"
	"net/http"

	"github.com/PhyoYazar/uas/business/core/student"
	"github.com/PhyoYazar/uas/business/data/order"
	"github.com/PhyoYazar/uas/business/sys/validate"
)

var orderByFields = map[string]struct{}{
	student.OrderByID:            {},
	student.OrderByStudentNumber: {},
	student.OrderByRollNumber:    {},
	// student.OrderByEmail:      {},
}

func parseOrder(r *http.Request) (order.By, error) {
	orderBy, err := order.Parse(r, student.DefaultOrderBy)
	if err != nil {
		return order.By{}, err
	}

	if _, exists := orderByFields[orderBy.Field]; !exists {
		return order.By{}, validate.NewFieldsError(orderBy.Field, errors.New("order field does not exist"))
	}

	return orderBy, nil
}
