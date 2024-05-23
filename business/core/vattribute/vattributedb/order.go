package vattributedb

import (
	"fmt"

	"github.com/PhyoYazar/uas/business/core/vattribute"
	"github.com/PhyoYazar/uas/business/data/order"
)

var orderByFields = map[string]string{
	vattribute.OrderByID:       "attribute_id",
	vattribute.OrderByName:     "name",
	vattribute.OrderByInstance: "instance",
	vattribute.OrderByType:     "type",
}

func orderByClause(orderBy order.By) (string, error) {
	by, exists := orderByFields[orderBy.Field]
	if !exists {
		return "", fmt.Errorf("field %q does not exist", orderBy.Field)
	}

	return " ORDER BY " + by + " " + orderBy.Direction, nil
}
