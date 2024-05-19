package comarkdb

import (
	"fmt"

	"github.com/PhyoYazar/uas/business/core/comark"
	"github.com/PhyoYazar/uas/business/data/order"
)

var orderByFields = map[string]string{
	comark.OrderByID:     "co_mark_id",
	comark.OrderByCoID:   "co_id",
	comark.OrderByMarkID: "mark_id",
}

func orderByClause(orderBy order.By) (string, error) {
	by, exists := orderByFields[orderBy.Field]
	if !exists {
		return "", fmt.Errorf("field %q does not exist", orderBy.Field)
	}

	return " ORDER BY " + by + " " + orderBy.Direction, nil
}
