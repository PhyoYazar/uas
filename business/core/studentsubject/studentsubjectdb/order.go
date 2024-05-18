package studentsubjectdb

import (
	"fmt"

	"github.com/PhyoYazar/uas/business/core/studentsubject"
	"github.com/PhyoYazar/uas/business/data/order"
)

var orderByFields = map[string]string{
	studentsubject.OrderByMark:      "mark",
	studentsubject.OrderByStudentID: "student_id",
	studentsubject.OrderBySubjectID: "subject_id",
}

func orderByClause(orderBy order.By) (string, error) {
	by, exists := orderByFields[orderBy.Field]
	if !exists {
		return "", fmt.Errorf("field %q does not exist", orderBy.Field)
	}

	return " ORDER BY " + by + " " + orderBy.Direction, nil
}
