package vattributedb

import (
	"github.com/PhyoYazar/uas/business/core/vattribute"
)

func existInSlice(structs []vattribute.VCo, newStruct vattribute.VCo) bool {
	for _, existingStruct := range structs {
		if existingStruct.ID == newStruct.ID {
			return true
		}
	}
	return false
}
