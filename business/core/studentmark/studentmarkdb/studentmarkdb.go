package studentmarkdb

import (
	"bytes"
	"context"
	"errors"
	"fmt"

	"github.com/PhyoYazar/uas/business/core/studentmark"
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

// Create inserts a new ga into the database.
func (s *Store) Create(ctx context.Context, ss studentmark.StudentMark) error {

	const q = `
	INSERT INTO student_marks
		(student_mark_id, mark, student_id, subject_id, attribute_id,date_created, date_updated)
	VALUES
		(:student_mark_id, :mark, :student_id, :subject_id, :attribute_id, :date_created, :date_updated)`

	if err := database.NamedExecContext(ctx, s.log, s.db, q, toDBStudentMark(ss)); err != nil {
		if errors.Is(err, database.ErrDBDuplicatedEntry) {
			return fmt.Errorf("namedexeccontext: %w", studentmark.ErrUniqueStudentMark)
		}
		return fmt.Errorf("namedexeccontext: %w", err)
	}

	return nil
}

// Update replaces a user document in the database.
func (s *Store) Update(ctx context.Context, std studentmark.StudentMark) error {
	const q = `
	UPDATE
		student_marks
	SET
		"mark" = :mark,
		"date_updated" = :date_updated
	WHERE
		student_mark_id = :student_mark_id`

	if err := database.NamedExecContext(ctx, s.log, s.db, q, toDBStudentMark(std)); err != nil {
		if errors.Is(err, database.ErrDBDuplicatedEntry) {
			return studentmark.ErrUniqueStudentMark
		}
		return fmt.Errorf("namedexeccontext: %w", err)
	}

	return nil
}

// Delete removes a user from the database.
func (s *Store) Delete(ctx context.Context, std studentmark.StudentMark) error {
	data := struct {
		UserID string `db:"student_mark_id"`
	}{
		UserID: std.ID.String(),
	}

	const q = `
	DELETE FROM
		student_marks
	WHERE
		student_mark_id = :student_mark_id`

	if err := database.NamedExecContext(ctx, s.log, s.db, q, data); err != nil {
		return fmt.Errorf("namedexeccontext: %w", err)
	}

	return nil
}

// Query retrieves a list of existing gas from the database.
func (s *Store) Query(ctx context.Context, filter studentmark.QueryFilter, orderBy order.By, pageNumber int, rowsPerPage int) ([]studentmark.StudentMark, error) {
	data := map[string]interface{}{
		"offset":        (pageNumber - 1) * rowsPerPage,
		"rows_per_page": rowsPerPage,
	}

	const q = `
	SELECT
	*
	FROM
	student_marks`

	buf := bytes.NewBufferString(q)
	s.applyFilter(filter, data, buf)

	orderByClause, err := orderByClause(orderBy)
	if err != nil {
		return nil, err
	}

	buf.WriteString(orderByClause)
	buf.WriteString(" OFFSET :offset ROWS FETCH NEXT :rows_per_page ROWS ONLY")

	var dbSs []dbStudentMark
	if err := database.NamedQuerySlice(ctx, s.log, s.db, buf.String(), data, &dbSs); err != nil {
		return nil, fmt.Errorf("namedqueryslice: %w", err)
	}

	mark, err := toCoreStudentMarkSlice(dbSs)
	if err != nil {
		return nil, err
	}

	return mark, nil
}

// QueryByID gets the specified subject from the database.
func (s *Store) QueryByID(ctx context.Context, studentMarkID uuid.UUID) (studentmark.StudentMark, error) {
	data := struct {
		ID string `db:"student_mark_id"`
	}{
		ID: studentMarkID.String(),
	}

	const q = `
	SELECT
        student_mark_id, attribute_id, subject_id, student_id, date_created, date_updated
	FROM
		student_marks
	WHERE
		student_mark_id = :student_mark_id`

	var dbStd dbStudentMark
	if err := database.NamedQueryStruct(ctx, s.log, s.db, q, data, &dbStd); err != nil {
		if errors.Is(err, database.ErrDBNotFound) {
			return studentmark.StudentMark{}, fmt.Errorf("db: %w", studentmark.ErrNotFound)
		}
		return studentmark.StudentMark{}, fmt.Errorf("db: %w", err)
	}

	return toCoreStudentMark(dbStd)
}

// Count returns the total number of cos in the DB.
func (s *Store) Count(ctx context.Context, filter studentmark.QueryFilter) (int, error) {
	data := map[string]interface{}{}

	const q = `
	SELECT
		count(1)
	FROM
	student_marks`

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
