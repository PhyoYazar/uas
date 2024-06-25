package studentdb

import (
	"bytes"
	"context"
	"errors"
	"fmt"

	"github.com/PhyoYazar/uas/business/core/student"
	"github.com/PhyoYazar/uas/business/data/order"
	database "github.com/PhyoYazar/uas/business/sys/database/pgx"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// Store manages the set of APIs for user database access.
type Store struct {
	log *zap.SugaredLogger
	db  *sqlx.DB
}

// NewStore constructs the api for data access.
func NewStore(log *zap.SugaredLogger, db *sqlx.DB) *Store {
	return &Store{
		log: log,
		db:  db,
	}
}

// Create inserts a new student into the database.
func (s *Store) Create(ctx context.Context, std student.Student) error {
	const q = `
	INSERT INTO students
		(student_id, student_name, year, academic_year, roll_number, date_created, date_updated)
	VALUES
		(:student_id, :student_name, :year, :academic_year, :roll_number, :date_created, :date_updated)`

	if err := database.NamedExecContext(ctx, s.log, s.db, q, toDBStudent(std)); err != nil {
		if errors.Is(err, database.ErrDBDuplicatedEntry) {
			return fmt.Errorf("namedexeccontext: %w", student.ErrUniqueStudent)
		}
		return fmt.Errorf("namedexeccontext: %w", err)
	}

	return nil
}

// Update replaces a user document in the database.
func (s *Store) Update(ctx context.Context, std student.Student) error {
	const q = `
	UPDATE
		students
	SET
		"student_name" = :student_name,
		"academic_year" = :academic_year,
		"year" = :year,
		"roll_number" = :roll_number,
		"date_updated" = :date_updated
	WHERE
		student_id = :student_id`

	if err := database.NamedExecContext(ctx, s.log, s.db, q, toDBStudent(std)); err != nil {
		if errors.Is(err, database.ErrDBDuplicatedEntry) {
			return student.ErrUniqueStudent
		}
		return fmt.Errorf("namedexeccontext: %w", err)
	}

	return nil
}

// Delete removes a user from the database.
func (s *Store) Delete(ctx context.Context, std student.Student) error {
	data := struct {
		UserID string `db:"student_id"`
	}{
		UserID: std.ID.String(),
	}

	const q = `
	DELETE FROM
		students
	WHERE
		student_id = :student_id`

	if err := database.NamedExecContext(ctx, s.log, s.db, q, data); err != nil {
		return fmt.Errorf("namedexeccontext: %w", err)
	}

	return nil
}

// Query retrieves a list of existing students from the database.
func (s *Store) Query(ctx context.Context, filter student.QueryFilter, orderBy order.By, pageNumber int, rowsPerPage int) ([]student.Student, error) {
	data := map[string]interface{}{
		"offset":        (pageNumber - 1) * rowsPerPage,
		"rows_per_page": rowsPerPage,
	}

	const q = `
	SELECT
		*
	FROM
		students`

	buf := bytes.NewBufferString(q)
	s.applyFilter(filter, data, buf)

	orderByClause, err := orderByClause(orderBy)
	if err != nil {
		return nil, err
	}

	buf.WriteString(orderByClause)
	buf.WriteString(" OFFSET :offset ROWS FETCH NEXT :rows_per_page ROWS ONLY")

	var dbStudents []dbStudent
	if err := database.NamedQuerySlice(ctx, s.log, s.db, buf.String(), data, &dbStudents); err != nil {
		return nil, fmt.Errorf("namedqueryslice: %w", err)
	}

	return toCoreStudentSlice(dbStudents)
}

// QueryByID gets the specified subject from the database.
func (s *Store) QueryByID(ctx context.Context, studentID uuid.UUID) (student.Student, error) {
	data := struct {
		ID string `db:"student_id"`
	}{
		ID: studentID.String(),
	}

	const q = `
	SELECT
        student_id, student_name, roll_number, year, academic_year, date_created, date_updated
	FROM
		students
	WHERE
		student_id = :student_id`

	var dbStd dbStudent
	if err := database.NamedQueryStruct(ctx, s.log, s.db, q, data, &dbStd); err != nil {
		if errors.Is(err, database.ErrDBNotFound) {
			return student.Student{}, fmt.Errorf("db: %w", student.ErrNotFound)
		}
		return student.Student{}, fmt.Errorf("db: %w", err)
	}

	return toCoreStudent(dbStd)
}

// Count returns the total number of students in the DB.
func (s *Store) Count(ctx context.Context, filter student.QueryFilter) (int, error) {
	data := map[string]interface{}{}

	const q = `
	SELECT
		count(1)
	FROM
		students`

	buf := bytes.NewBufferString(q)
	s.applyFilter(filter, data, buf)

	var count struct {
		Count int `db:"count"`
	}
	if err := database.NamedQueryStruct(ctx, s.log, s.db, buf.String(), data, &count); err != nil {
		return 0, fmt.Errorf("namedquerystruct: %w", err)
	}

	return count.Count, nil
}
