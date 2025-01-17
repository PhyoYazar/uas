package ga

import "github.com/PhyoYazar/uas/business/data/order"

// DefaultOrderBy represents the default way we sort.
var DefaultOrderBy = order.NewBy(OrderByIncrementingColumn, order.ASC)

// Set of fields that the results can be ordered by. These are the names
// that should be used by the application layer.
const (
	OrderByID                 = "gaid"
	OrderByName               = "name"
	OrderBySlug               = "slug"
	OrderByIncrementingColumn = "inc"
)
