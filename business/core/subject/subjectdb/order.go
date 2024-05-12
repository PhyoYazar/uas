package subjectdb

import (
	"fmt"

	"github.com/PhyoYazar/uas/business/core/subject"
	"github.com/PhyoYazar/uas/business/data/order"
)

var orderByFields = map[string]string{
	subject.OrderByID:         "subject_id",
	subject.OrderByName:       "name",
	subject.OrderByInstructor: "instructor",
}

func orderByClause(orderBy order.By) (string, error) {
	by, exists := orderByFields[orderBy.Field]
	if !exists {
		return "", fmt.Errorf("field %q does not exist", orderBy.Field)
	}

	return " ORDER BY " + by + " " + orderBy.Direction, nil
}
