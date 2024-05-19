package cogadb

import (
	"fmt"

	"github.com/PhyoYazar/uas/business/core/coga"
	"github.com/PhyoYazar/uas/business/data/order"
)

var orderByFields = map[string]string{
	coga.OrderByID:   "co_ga_id",
	coga.OrderByCoID: "co_id",
	coga.OrderByGaID: "ga_id",
}

func orderByClause(orderBy order.By) (string, error) {
	by, exists := orderByFields[orderBy.Field]
	if !exists {
		return "", fmt.Errorf("field %q does not exist", orderBy.Field)
	}

	return " ORDER BY " + by + " " + orderBy.Direction, nil
}
