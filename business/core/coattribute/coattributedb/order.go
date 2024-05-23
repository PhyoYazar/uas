package coattributedb

import (
	"fmt"

	"github.com/PhyoYazar/uas/business/core/coattribute"
	"github.com/PhyoYazar/uas/business/data/order"
)

var orderByFields = map[string]string{
	coattribute.OrderByID:          "co_attribute_id",
	coattribute.OrderByCoID:        "co_id",
	coattribute.OrderByAttributeID: "attribute_id",
}

func orderByClause(orderBy order.By) (string, error) {
	by, exists := orderByFields[orderBy.Field]
	if !exists {
		return "", fmt.Errorf("field %q does not exist", orderBy.Field)
	}

	return " ORDER BY " + by + " " + orderBy.Direction, nil
}
