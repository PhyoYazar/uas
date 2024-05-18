package studentsubjectdb

import (
	"fmt"
	"time"

	"github.com/PhyoYazar/uas/business/core/studentsubject"
	"github.com/google/uuid"
)

// dbStudentSubject represent the structure we need for moving data
// between the app and the database.
type dbStudentSubject struct {
	ID          uuid.UUID `db:"student_subject_id"`
	StudentID   uuid.UUID `db:"student_id"`
	SubjectID   uuid.UUID `db:"subject_id"`
	Mark        int       `db:"mark"`
	DateCreated time.Time `db:"date_created"`
	DateUpdated time.Time `db:"date_updated"`
}

func toDBStudentSubject(ss studentsubject.StudentSubject) dbStudentSubject {

	mark := dbStudentSubject{
		ID:          ss.ID,
		StudentID:   ss.StudentID,
		SubjectID:   ss.SubjectID,
		Mark:        ss.Mark,
		DateCreated: ss.DateCreated.UTC(),
		DateUpdated: ss.DateUpdated.UTC(),
	}

	return mark
}

func toCoreStudentSubject(dbStudentSubject dbStudentSubject) (studentsubject.StudentSubject, error) {

	ss := studentsubject.StudentSubject{
		ID:          dbStudentSubject.ID,
		StudentID:   dbStudentSubject.StudentID,
		SubjectID:   dbStudentSubject.SubjectID,
		Mark:        dbStudentSubject.Mark,
		DateCreated: dbStudentSubject.DateCreated.In(time.Local),
		DateUpdated: dbStudentSubject.DateUpdated.In(time.Local),
	}

	return ss, nil
}

func toCoreStudentSubjectSlice(dbStudentSubjects []dbStudentSubject) ([]studentsubject.StudentSubject, error) {
	ss := make([]studentsubject.StudentSubject, len(dbStudentSubjects))
	for i, dbSs := range dbStudentSubjects {
		var err error
		ss[i], err = toCoreStudentSubject(dbSs)
		if err != nil {
			return nil, fmt.Errorf("parse type: %w", err)
		}
	}
	return ss, nil
}
