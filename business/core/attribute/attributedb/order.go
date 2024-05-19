package attributedb

import (
	"fmt"

	"github.com/PhyoYazar/uas/business/core/attribute"
	"github.com/PhyoYazar/uas/business/data/order"
)

var orderByFields = map[string]string{
	attribute.OrderByID:       "attribute_id",
	attribute.OrderByName:     "name",
	attribute.OrderByInstance: "instance",
}

func orderByClause(orderBy order.By) (string, error) {
	by, exists := orderByFields[orderBy.Field]
	if !exists {
		return "", fmt.Errorf("field %q does not exist", orderBy.Field)
	}

	return " ORDER BY " + by + " " + orderBy.Direction, nil
}
