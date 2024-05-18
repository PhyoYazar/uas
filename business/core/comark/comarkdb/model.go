package comarkdb

import (
	"fmt"
	"time"

	"github.com/PhyoYazar/uas/business/core/comark"
	"github.com/google/uuid"
)

// dbCoMark represent the structure we need for moving data
// between the app and the database.
type dbCoMark struct {
	ID          uuid.UUID `db:"co_mark_id"`
	CoID        uuid.UUID `db:"co_id"`
	MarkID      uuid.UUID `db:"mark_id"`
	DateCreated time.Time `db:"date_created"`
	DateUpdated time.Time `db:"date_updated"`
}

func toDBCoMark(comark comark.CoMark) dbCoMark {

	cm := dbCoMark{
		ID:          comark.ID,
		CoID:        comark.CoID,
		MarkID:      comark.MarkID,
		DateCreated: comark.DateCreated.UTC(),
		DateUpdated: comark.DateUpdated.UTC(),
	}

	return cm
}

func toCoreCoMark(dbCoMark dbCoMark) (comark.CoMark, error) {

	cm := comark.CoMark{
		ID:          dbCoMark.ID,
		CoID:        dbCoMark.CoID,
		MarkID:      dbCoMark.MarkID,
		DateCreated: dbCoMark.DateCreated.In(time.Local),
		DateUpdated: dbCoMark.DateUpdated.In(time.Local),
	}

	return cm, nil
}

func toCoreCoMarkSlice(comarks []dbCoMark) ([]comark.CoMark, error) {
	cm := make([]comark.CoMark, len(comarks))
	for i, dbCoMark := range comarks {
		var err error
		cm[i], err = toCoreCoMark(dbCoMark)
		if err != nil {
			return nil, fmt.Errorf("parse type: %w", err)
		}
	}
	return cm, nil
}
