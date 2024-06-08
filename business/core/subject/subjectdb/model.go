package subjectdb

import (
	"fmt"
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
	Semester     string    `db:"semester"`
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
		Year:         sub.Year.Name(),
		AcademicYear: sub.AcademicYear,
		Semester:     sub.Semester.Name(),
		Instructor:   sub.Instructor,
		Exam:         sub.Exam,
		Practical:    sub.Practical,
		DateCreated:  sub.DateCreated.UTC(),
		DateUpdated:  sub.DateUpdated.UTC(),
	}

	return subject
}

func toCoreSubject(dbSubject dbSubject) (subject.Subject, error) {

	semester, err := subject.ParseSemester(dbSubject.Semester)
	if err != nil {
		return subject.Subject{}, fmt.Errorf("parse type: %w", err)
	}

	year, err := subject.ParseYear(dbSubject.Year)
	if err != nil {
		return subject.Subject{}, fmt.Errorf("parse type: %w", err)
	}

	sub := subject.Subject{
		ID:           dbSubject.ID,
		Name:         dbSubject.Name,
		Code:         dbSubject.Code,
		Year:         year,
		AcademicYear: dbSubject.AcademicYear,
		Semester:     semester,
		Instructor:   dbSubject.Instructor,
		Exam:         dbSubject.Exam,
		Practical:    dbSubject.Practical,
		DateCreated:  dbSubject.DateCreated.In(time.Local),
		DateUpdated:  dbSubject.DateUpdated.In(time.Local),
	}

	return sub, nil
}

func toCoreSubjectSlice(dbSubjects []dbSubject) ([]subject.Subject, error) {
	subs := make([]subject.Subject, len(dbSubjects))

	for i, dbSubject := range dbSubjects {
		var err error

		subs[i], err = toCoreSubject(dbSubject)
		if err != nil {
			return nil, fmt.Errorf("parse type: %w", err)
		}
	}
	return subs, nil
}
