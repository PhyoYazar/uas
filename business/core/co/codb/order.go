package codb

import (
	"fmt"

	"github.com/PhyoYazar/uas/business/core/co"
	"github.com/PhyoYazar/uas/business/data/order"
)

var orderByFields = map[string]string{
	co.OrderByID:       "co_id",
	co.OrderByName:     "name",
	co.OrderByInstance: "instance",
	co.OrderByMark:     "mark",
}

func orderByClause(orderBy order.By) (string, error) {
	by, exists := orderByFields[orderBy.Field]
	if !exists {
		return "", fmt.Errorf("field %q does not exist", orderBy.Field)
	}

	return " ORDER BY " + by + " " + orderBy.Direction, nil
}
