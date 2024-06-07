package vattributedb

// Define an interface that requires a GetID method
type HasID interface {
	GetID() interface{}
}

func existInSlice[T HasID](existingItems []T, compareItem T) bool {
	compareItemID := compareItem.GetID()

	for _, item := range existingItems {
		if item.GetID() == compareItemID {
			return true
		}
	}
	return false
}
