package studentdb

import (
	"net/mail"
	"time"

	"github.com/PhyoYazar/uas/business/core/student"
	"github.com/google/uuid"
)

// dbStudent represent the structure we need for moving data
// between the app and the database.
type dbStudent struct {
	ID           uuid.UUID `db:"student_id"`
	Name         string    `db:"name"`
	Email        string    `db:"email"`
	Year         string    `db:"year"`
	AcademicYear string    `db:"academic_year"`
	PhoneNumber  string    `db:"phone_number"`
	RollNumber   int       `db:"roll_number"`
	DateCreated  time.Time `db:"date_created"`
	DateUpdated  time.Time `db:"date_updated"`
}

func toDBStudent(std student.Student) dbStudent {

	student := dbStudent{
		ID:           std.ID,
		Name:         std.Name,
		Email:        std.Email.Address,
		Year:         std.Year,
		AcademicYear: std.AcademicYear,
		RollNumber:   std.RollNumber,
		PhoneNumber:  std.PhoneNumber,
		DateCreated:  std.DateCreated.UTC(),
		DateUpdated:  std.DateUpdated.UTC(),
	}

	return student
}

func toCoreStudent(dbStd dbStudent) student.Student {
	addr := mail.Address{
		Address: dbStd.Email,
	}

	std := student.Student{
		ID:           dbStd.ID,
		Name:         dbStd.Name,
		Email:        addr,
		Year:         dbStd.Year,
		AcademicYear: dbStd.AcademicYear,
		PhoneNumber:  dbStd.PhoneNumber,
		RollNumber:   dbStd.RollNumber,
		DateCreated:  dbStd.DateCreated.In(time.Local),
		DateUpdated:  dbStd.DateUpdated.In(time.Local),
	}

	return std
}

func toCoreStudentSlice(dbStudents []dbStudent) []student.Student {
	stds := make([]student.Student, len(dbStudents))
	for i, dbStudent := range dbStudents {
		stds[i] = toCoreStudent(dbStudent)
	}
	return stds
}
