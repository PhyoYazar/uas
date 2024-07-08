package vgagradedb

import (
	"fmt"

	"github.com/PhyoYazar/uas/business/core/vstudentmark"
	"github.com/PhyoYazar/uas/business/data/order"
)

var orderByFields = map[string]string{
	vstudentmark.OrderByID:         "student_id",
	vstudentmark.OrderByRollNumber: "roll_number",
}

func orderByClause(orderBy order.By) (string, error) {
	by, exists := orderByFields[orderBy.Field]
	if !exists {
		return "", fmt.Errorf("field %q does not exist", orderBy.Field)
	}

	return " ORDER BY " + by + " " + orderBy.Direction, nil
}
