package studentmarkdb

import (
	"fmt"

	"github.com/PhyoYazar/uas/business/core/studentmark"
	"github.com/PhyoYazar/uas/business/data/order"
)

var orderByFields = map[string]string{
	studentmark.OrderByID:          "student_mark_id",
	studentmark.OrderByMark:        "mark",
	studentmark.OrderByStudentID:   "student_id",
	studentmark.OrderBySubjectID:   "subject_id",
	studentmark.OrderByAttributeID: "attribute_id",
}

func orderByClause(orderBy order.By) (string, error) {
	by, exists := orderByFields[orderBy.Field]
	if !exists {
		return "", fmt.Errorf("field %q does not exist", orderBy.Field)
	}

	return " ORDER BY " + by + " " + orderBy.Direction, nil
}
