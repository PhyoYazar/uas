package studentmark

import "github.com/PhyoYazar/uas/business/data/order"

// DefaultOrderBy represents the default way we sort.
var DefaultOrderBy = order.NewBy(OrderByID, order.ASC)

// Set of fields that the results can be ordered by. These are the names
// that should be used by the application layer.
const (
	OrderByID          = "studentmarkid"
	OrderByMark        = "mark"
	OrderByStudentID   = "student_id"
	OrderBySubjectID   = "subject_id"
	OrderByAttributeID = "attribute_id"
)