package markdb

import (
	"fmt"
	"time"

	"github.com/PhyoYazar/uas/business/core/mark"
	"github.com/google/uuid"
)

// dbMark represent the structure we need for moving data
// between the app and the database.
type dbMark struct {
	ID          uuid.UUID `db:"mark_id"`
	CoID        uuid.UUID `db:"co_id"`
	GaID        uuid.UUID `db:"ga_id"`
	AttributeID uuid.UUID `db:"attribute_id"`
	Mark        int       `db:"mark"`
	DateCreated time.Time `db:"date_created"`
	DateUpdated time.Time `db:"date_updated"`
}

func toDBMark(mark mark.Mark) dbMark {

	m := dbMark{
		ID:          mark.ID,
		CoID:        mark.CoID,
		GaID:        mark.GaID,
		AttributeID: mark.AttributeID,
		Mark:        mark.Mark,
		DateCreated: mark.DateCreated.UTC(),
		DateUpdated: mark.DateUpdated.UTC(),
	}

	return m
}

func toCoreMark(dbMark dbMark) (mark.Mark, error) {

	m := mark.Mark{
		ID:          dbMark.ID,
		CoID:        dbMark.CoID,
		GaID:        dbMark.GaID,
		AttributeID: dbMark.AttributeID,
		Mark:        dbMark.Mark,
		DateCreated: dbMark.DateCreated.In(time.Local),
		DateUpdated: dbMark.DateUpdated.In(time.Local),
	}

	return m, nil
}

func toCoreMarkSlice(marks []dbMark) ([]mark.Mark, error) {
	m := make([]mark.Mark, len(marks))
	for i, dbMark := range marks {
		var err error
		m[i], err = toCoreMark(dbMark)
		if err != nil {
			return nil, fmt.Errorf("parse type: %w", err)
		}
	}
	return m, nil
}
