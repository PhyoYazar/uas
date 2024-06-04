package gadb

import (
	"time"

	"github.com/PhyoYazar/uas/business/core/ga"
	"github.com/google/uuid"
)

// dbSubject represent the structure we need for moving data
// between the app and the database.
type dbGa struct {
	ID                 uuid.UUID `db:"ga_id"`
	Name               string    `db:"name"`
	Slug               string    `db:"slug"`
	IncrementingColumn int       `db:"incrementing_column"`
	DateCreated        time.Time `db:"date_created"`
	DateUpdated        time.Time `db:"date_updated"`
}

func toDBGa(g ga.Ga) dbGa {

	ga := dbGa{
		ID:          g.ID,
		Name:        g.Name,
		Slug:        g.Slug,
		DateCreated: g.DateCreated.UTC(),
		DateUpdated: g.DateUpdated.UTC(),
	}

	return ga
}

func toCoreGa(dbGa dbGa) ga.Ga {

	ga := ga.Ga{
		ID:          dbGa.ID,
		Name:        dbGa.Name,
		Slug:        dbGa.Slug,
		DateCreated: dbGa.DateCreated.In(time.Local),
		DateUpdated: dbGa.DateUpdated.In(time.Local),
	}

	return ga
}

func toCoreGaSlice(dbGas []dbGa) []ga.Ga {
	gas := make([]ga.Ga, len(dbGas))
	for i, dbGa := range dbGas {
		gas[i] = toCoreGa(dbGa)
	}
	return gas
}
