package studentdb

import (
	"fmt"
	"time"

	"github.com/PhyoYazar/uas/business/core/student"
	"github.com/google/uuid"
)

// dbStudent represent the structure we need for moving data
// between the app and the database.
type dbStudent struct {
	ID           uuid.UUID `db:"student_id"`
	StudentName  string    `db:"student_name"`
	Year         string    `db:"year"`
	AcademicYear string    `db:"academic_year"`
	RollNumber   int       `db:"roll_number"`
	DateCreated  time.Time `db:"date_created"`
	DateUpdated  time.Time `db:"date_updated"`
}

func toDBStudent(std student.Student) dbStudent {

	student := dbStudent{
		ID:           std.ID,
		StudentName:  std.StudentName,
		Year:         std.Year.Name(),
		AcademicYear: std.AcademicYear,
		RollNumber:   std.RollNumber,
		DateCreated:  std.DateCreated.UTC(),
		DateUpdated:  std.DateUpdated.UTC(),
	}

	return student
}

func toCoreStudent(dbStd dbStudent) (student.Student, error) {

	year, err := student.ParseYear(dbStd.Year)
	if err != nil {
		return student.Student{}, fmt.Errorf("parse type: %w", err)
	}

	std := student.Student{
		ID:           dbStd.ID,
		StudentName:  dbStd.StudentName,
		Year:         year,
		AcademicYear: dbStd.AcademicYear,
		RollNumber:   dbStd.RollNumber,
		DateCreated:  dbStd.DateCreated.In(time.Local),
		DateUpdated:  dbStd.DateUpdated.In(time.Local),
	}

	return std, nil
}

func toCoreStudentSlice(dbStudents []dbStudent) ([]student.Student, error) {
	stds := make([]student.Student, len(dbStudents))
	for i, dbStudent := range dbStudents {
		var err error
		stds[i], err = toCoreStudent(dbStudent)

		if err != nil {
			return nil, fmt.Errorf("parse type: %w", err)
		}
	}
	return stds, nil
}
