package markdb

import (
	"fmt"

	"github.com/PhyoYazar/uas/business/core/mark"
	"github.com/PhyoYazar/uas/business/data/order"
)

var orderByFields = map[string]string{
	mark.OrderByID:          "mark_id",
	mark.OrderByCoID:        "co_id",
	mark.OrderByGaID:        "ga_id",
	mark.OrderByAttributeID: "attribute_id",
	mark.OrderByMark:        "mark",
}

func orderByClause(orderBy order.By) (string, error) {
	by, exists := orderByFields[orderBy.Field]
	if !exists {
		return "", fmt.Errorf("field %q does not exist", orderBy.Field)
	}

	return " ORDER BY " + by + " " + orderBy.Direction, nil
}
