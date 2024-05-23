package coattributedb

import (
	"fmt"
	"time"

	"github.com/PhyoYazar/uas/business/core/coattribute"
	"github.com/google/uuid"
)

// dbCoGa represent the structure we need for moving data
// between the app and the database.
type dbCoAttribute struct {
	ID          uuid.UUID `db:"co_attribute_id"`
	CoID        uuid.UUID `db:"co_id"`
	AttributeID uuid.UUID `db:"attribute_id"`
	DateCreated time.Time `db:"date_created"`
	DateUpdated time.Time `db:"date_updated"`
}

func toDBCoAttribute(coga coattribute.CoAttribute) dbCoAttribute {

	cg := dbCoAttribute{
		ID:          coga.ID,
		CoID:        coga.CoID,
		AttributeID: coga.AttributeID,
		DateCreated: coga.DateCreated.UTC(),
		DateUpdated: coga.DateUpdated.UTC(),
	}

	return cg
}

func toCoreCoAttribute(dbCoAtt dbCoAttribute) (coattribute.CoAttribute, error) {

	cg := coattribute.CoAttribute{
		ID:          dbCoAtt.ID,
		CoID:        dbCoAtt.CoID,
		AttributeID: dbCoAtt.AttributeID,
		DateCreated: dbCoAtt.DateCreated.In(time.Local),
		DateUpdated: dbCoAtt.DateUpdated.In(time.Local),
	}

	return cg, nil
}

func toCoreCoGaSlice(cogas []dbCoAttribute) ([]coattribute.CoAttribute, error) {
	cg := make([]coattribute.CoAttribute, len(cogas))
	for i, dbCoGa := range cogas {
		var err error
		cg[i], err = toCoreCoAttribute(dbCoGa)
		if err != nil {
			return nil, fmt.Errorf("parse type: %w", err)
		}
	}
	return cg, nil
}
