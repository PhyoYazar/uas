package gadb

import (
	"fmt"

	"github.com/PhyoYazar/uas/business/core/ga"
	"github.com/PhyoYazar/uas/business/data/order"
)

var orderByFields = map[string]string{
	ga.OrderByID:                 "ga_id",
	ga.OrderByName:               "name",
	ga.OrderBySlug:               "slug",
	ga.OrderByIncrementingColumn: "incrementing_column",
}

func orderByClause(orderBy order.By) (string, error) {
	by, exists := orderByFields[orderBy.Field]
	if !exists {
		return "", fmt.Errorf("field %q does not exist", orderBy.Field)
	}

	return " ORDER BY " + by + " " + orderBy.Direction, nil
}
