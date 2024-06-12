package fullmarkdb

import (
	"fmt"
	"time"

	"github.com/PhyoYazar/uas/business/core/fullmark"
	"github.com/google/uuid"
)

// dbMark represent the structure we need for moving data
// between the app and the database.
type dbFullMark struct {
	ID          uuid.UUID `db:"full_mark_id"`
	SubjectID   uuid.UUID `db:"subject_id"`
	AttributeID uuid.UUID `db:"attribute_id"`
	Mark        int       `db:"mark"`
	DateCreated time.Time `db:"date_created"`
	DateUpdated time.Time `db:"date_updated"`
}

func toDBFullMark(mark fullmark.FullMark) dbFullMark {

	m := dbFullMark{
		ID:          mark.ID,
		SubjectID:   mark.SubjectID,
		AttributeID: mark.AttributeID,
		Mark:        mark.Mark,
		DateCreated: mark.DateCreated.UTC(),
		DateUpdated: mark.DateUpdated.UTC(),
	}

	return m
}

func toCoreFullMark(dbFullMark dbFullMark) (fullmark.FullMark, error) {

	m := fullmark.FullMark{
		ID:          dbFullMark.ID,
		SubjectID:   dbFullMark.SubjectID,
		AttributeID: dbFullMark.AttributeID,
		Mark:        dbFullMark.Mark,
		DateCreated: dbFullMark.DateCreated.In(time.Local),
		DateUpdated: dbFullMark.DateUpdated.In(time.Local),
	}

	return m, nil
}

func toCoreFullMarkSlice(fullMarks []dbFullMark) ([]fullmark.FullMark, error) {
	fm := make([]fullmark.FullMark, len(fullMarks))
	for i, dbMark := range fullMarks {
		var err error
		fm[i], err = toCoreFullMark(dbMark)
		if err != nil {
			return nil, fmt.Errorf("parse type: %w", err)
		}
	}
	return fm, nil
}
