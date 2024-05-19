package comarkgrp

import (
	"errors"
	"net/http"

	"github.com/PhyoYazar/uas/business/core/comark"
	"github.com/PhyoYazar/uas/business/data/order"
	"github.com/PhyoYazar/uas/business/sys/validate"
)

var orderByFields = map[string]struct{}{
	comark.OrderByID:     {},
	comark.OrderByCoID:   {},
	comark.OrderByMarkID: {},
}

func parseOrder(r *http.Request) (order.By, error) {
	orderBy, err := order.Parse(r, comark.DefaultOrderBy)
	if err != nil {
		return order.By{}, err
	}

	if _, exists := orderByFields[orderBy.Field]; !exists {
		return order.By{}, validate.NewFieldsError(orderBy.Field, errors.New("order field does not exist"))
	}

	return orderBy, nil
}
