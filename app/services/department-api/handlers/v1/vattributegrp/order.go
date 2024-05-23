package vattributegrp

import (
	"errors"
	"net/http"

	"github.com/PhyoYazar/uas/business/core/vattribute"
	"github.com/PhyoYazar/uas/business/data/order"
	"github.com/PhyoYazar/uas/business/sys/validate"
)

var orderByFields = map[string]struct{}{
	vattribute.OrderByID:       {},
	vattribute.OrderByName:     {},
	vattribute.OrderByInstance: {},
	vattribute.OrderByType:     {},
}

func parseOrder(r *http.Request) (order.By, error) {
	orderBy, err := order.Parse(r, vattribute.DefaultOrderBy)
	if err != nil {
		return order.By{}, err
	}

	if _, exists := orderByFields[orderBy.Field]; !exists {
		return order.By{}, validate.NewFieldsError(orderBy.Field, errors.New("order field does not exist"))
	}

	return orderBy, nil
}
