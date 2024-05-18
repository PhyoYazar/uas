package cogadb

import (
	"fmt"
	"time"

	"github.com/PhyoYazar/uas/business/core/coga"
	"github.com/google/uuid"
)

// dbCoGa represent the structure we need for moving data
// between the app and the database.
type dbCoGa struct {
	ID          uuid.UUID `db:"co_ga_id"`
	CoID        uuid.UUID `db:"co_id"`
	GaID        uuid.UUID `db:"ga_id"`
	Mark        int       `db:"mark"`
	DateCreated time.Time `db:"date_created"`
	DateUpdated time.Time `db:"date_updated"`
}

func toDBCoGa(coga coga.CoGa) dbCoGa {

	cg := dbCoGa{
		ID:          coga.ID,
		CoID:        coga.CoID,
		GaID:        coga.GaID,
		Mark:        coga.Mark,
		DateCreated: coga.DateCreated.UTC(),
		DateUpdated: coga.DateUpdated.UTC(),
	}

	return cg
}

func toCoreCoGa(dbCoGa dbCoGa) (coga.CoGa, error) {

	cg := coga.CoGa{
		ID:          dbCoGa.ID,
		CoID:        dbCoGa.CoID,
		GaID:        dbCoGa.GaID,
		Mark:        dbCoGa.Mark,
		DateCreated: dbCoGa.DateCreated.In(time.Local),
		DateUpdated: dbCoGa.DateUpdated.In(time.Local),
	}

	return cg, nil
}

func toCoreCoGaSlice(cogas []dbCoGa) ([]coga.CoGa, error) {
	cg := make([]coga.CoGa, len(cogas))
	for i, dbCoGa := range cogas {
		var err error
		cg[i], err = toCoreCoGa(dbCoGa)
		if err != nil {
			return nil, fmt.Errorf("parse type: %w", err)
		}
	}
	return cg, nil
}
