package studentmarkdb

import (
	"fmt"
	"time"

	"github.com/PhyoYazar/uas/business/core/studentmark"
	"github.com/google/uuid"
)

// dbStudentMark represent the structure we need for moving data
// between the app and the database.
type dbStudentMark struct {
	ID          uuid.UUID `db:"student_mark_id"`
	StudentID   uuid.UUID `db:"student_id"`
	SubjectID   uuid.UUID `db:"subject_id"`
	AttributeID uuid.UUID `db:"attribute_id"`
	Mark        int       `db:"mark"`
	DateCreated time.Time `db:"date_created"`
	DateUpdated time.Time `db:"date_updated"`
}

func toDBStudentMark(ss studentmark.StudentMark) dbStudentMark {

	mark := dbStudentMark{
		ID:          ss.ID,
		StudentID:   ss.StudentID,
		SubjectID:   ss.SubjectID,
		AttributeID: ss.AttributeID,
		Mark:        ss.Mark,
		DateCreated: ss.DateCreated.UTC(),
		DateUpdated: ss.DateUpdated.UTC(),
	}

	return mark
}

func toCoreStudentMark(dbStudentMark dbStudentMark) (studentmark.StudentMark, error) {

	ss := studentmark.StudentMark{
		ID:          dbStudentMark.ID,
		StudentID:   dbStudentMark.StudentID,
		SubjectID:   dbStudentMark.SubjectID,
		AttributeID: dbStudentMark.AttributeID,
		Mark:        dbStudentMark.Mark,
		DateCreated: dbStudentMark.DateCreated.In(time.Local),
		DateUpdated: dbStudentMark.DateUpdated.In(time.Local),
	}

	return ss, nil
}

func toCoreStudentMarkSlice(dbStudentMarks []dbStudentMark) ([]studentmark.StudentMark, error) {
	ss := make([]studentmark.StudentMark, len(dbStudentMarks))
	for i, dbSm := range dbStudentMarks {
		var err error
		ss[i], err = toCoreStudentMark(dbSm)
		if err != nil {
			return nil, fmt.Errorf("parse type: %w", err)
		}
	}
	return ss, nil
}
