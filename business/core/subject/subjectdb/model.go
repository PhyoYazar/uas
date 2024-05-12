package subjectdb

import (
	"time"

	"github.com/PhyoYazar/uas/business/core/subject"
	"github.com/google/uuid"
)

// dbSubject represent the structure we need for moving data
// between the app and the database.
type dbSubject struct {
	ID           uuid.UUID `db:"subject_id"`
	Name         string    `db:"name"`
	Code         string    `db:"code"`
	Year         string    `db:"year"`
	AcademicYear string    `db:"academic_year"`
	Instructor   string    `db:"instructor"`
	Exam         int       `db:"exam"`
	Practical    int       `db:"practical"`
	DateCreated  time.Time `db:"date_created"`
	DateUpdated  time.Time `db:"date_updated"`
}

func toDBSubject(sub subject.Subject) dbSubject {

	subject := dbSubject{
		ID:           sub.ID,
		Name:         sub.Name,
		Code:         sub.Code,
		Year:         sub.Year,
		AcademicYear: sub.AcademicYear,
		Instructor:   sub.Instructor,
		Exam:         sub.Exam,
		Practical:    sub.Practical,
		DateCreated:  sub.DateCreated.UTC(),
		DateUpdated:  sub.DateUpdated.UTC(),
	}

	return subject
}

func toCoreSubject(dbSubject dbSubject) subject.Subject {

	sub := subject.Subject{
		ID:           dbSubject.ID,
		Name:         dbSubject.Name,
		Code:         dbSubject.Code,
		Year:         dbSubject.Year,
		AcademicYear: dbSubject.AcademicYear,
		Instructor:   dbSubject.Instructor,
		Exam:         dbSubject.Exam,
		Practical:    dbSubject.Practical,
		DateCreated:  dbSubject.DateCreated.In(time.Local),
		DateUpdated:  dbSubject.DateUpdated.In(time.Local),
	}

	return sub
}

func toCoreUserSlice(dbSubjects []dbSubject) []subject.Subject {
	subs := make([]subject.Subject, len(dbSubjects))
	for i, dbSubject := range dbSubjects {
		subs[i] = toCoreSubject(dbSubject)
	}
	return subs
}
