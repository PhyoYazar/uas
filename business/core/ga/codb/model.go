package codb

import (
	"time"

	"github.com/PhyoYazar/uas/business/core/co"
	"github.com/google/uuid"
)

// dbSubject represent the structure we need for moving data
// between the app and the database.
type dbCo struct {
	ID          uuid.UUID `db:"co_id"`
	Name        string    `db:"name"`
	SubjectID   uuid.UUID `db:"subject_id"` // ID of the user who created the product.
	DateCreated time.Time `db:"date_created"`
	DateUpdated time.Time `db:"date_updated"`
}

func toDBCo(c co.Co) dbCo {

	co := dbCo{
		ID:          c.ID,
		Name:        c.Name,
		SubjectID:   c.SubjectID,
		DateCreated: c.DateCreated.UTC(),
		DateUpdated: c.DateUpdated.UTC(),
	}

	return co
}

func toCoreCo(dbCo dbCo) co.Co {

	co := co.Co{
		ID:          dbCo.ID,
		Name:        dbCo.Name,
		SubjectID:   dbCo.SubjectID,
		DateCreated: dbCo.DateCreated.In(time.Local),
		DateUpdated: dbCo.DateUpdated.In(time.Local),
	}

	return co
}

func toCoreCoSlice(dbCos []dbCo) []co.Co {
	cos := make([]co.Co, len(dbCos))
	for i, dbCo := range dbCos {
		cos[i] = toCoreCo(dbCo)
	}
	return cos
}