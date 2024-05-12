package studentdb

import (
	"fmt"

	"github.com/PhyoYazar/uas/business/core/student"
	"github.com/PhyoYazar/uas/business/data/order"
)

var orderByFields = map[string]string{
	student.OrderByID:         "student_id",
	student.OrderByName:       "name",
	student.OrderByEmail:      "email",
	student.OrderByRollNumber: "roll_number",
}

func orderByClause(orderBy order.By) (string, error) {
	by, exists := orderByFields[orderBy.Field]
	if !exists {
		return "", fmt.Errorf("field %q does not exist", orderBy.Field)
	}

	return " ORDER BY " + by + " " + orderBy.Direction, nil
}
