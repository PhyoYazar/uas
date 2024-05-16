package markdb

import (
	"time"

	"github.com/PhyoYazar/uas/business/core/mark"
	"github.com/google/uuid"
)

// dbSubject represent the structure we need for moving data
// between the app and the database.
type dbMark struct {
	ID          uuid.UUID `db:"mark_id"`
	Name        string    `db:"name"`
	Type        string    `db:"type"`
	Instance    int       `db:"instance"`
	DateCreated time.Time `db:"date_created"`
	DateUpdated time.Time `db:"date_updated"`
}

func toDBMark(mk mark.Mark) dbMark {

	mark := dbMark{
		ID:          mk.ID,
		Name:        mk.Name,
		Type:        mk.Type.Name(),
		Instance:    mk.Instance,
		DateCreated: mk.DateCreated.UTC(),
		DateUpdated: mk.DateUpdated.UTC(),
	}

	return mark
}

func toCoreMark(dbMark dbMark) mark.Mark {

	mark := mark.Mark{
		ID:          dbMark.ID,
		Name:        dbMark.Name,
		Type:        mark.MustParseMarkType(dbMark.Type),
		Instance:    dbMark.Instance,
		DateCreated: dbMark.DateCreated.In(time.Local),
		DateUpdated: dbMark.DateUpdated.In(time.Local),
	}

	return mark
}

func toCoreMarkSlice(dbMarks []dbMark) []mark.Mark {
	mks := make([]mark.Mark, len(dbMarks))
	for i, dbGa := range dbMarks {
		mks[i] = toCoreMark(dbGa)
	}
	return mks
}
