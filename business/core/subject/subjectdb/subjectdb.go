package subjectdb

import (
	"bytes"
	"context"
	"errors"
	"fmt"

	"github.com/PhyoYazar/uas/business/core/subject"
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

// Create inserts a new subject into the database.
func (s *Store) Create(ctx context.Context, sub subject.Subject) error {
	const q = `
	INSERT INTO subjects
		(subject_id, name, code, year, semester, academic_year, instructor, exam, practical, tutorial, lab, assignment, date_created, date_updated)
	VALUES
		(:subject_id, :name, :code, :year, :semester, :academic_year, :instructor, :exam, :practical, :tutorial, :lab, :assignment, :date_created, :date_updated)`

	if err := database.NamedExecContext(ctx, s.log, s.db, q, toDBSubject(sub)); err != nil {
		if errors.Is(err, database.ErrDBDuplicatedEntry) {
			return fmt.Errorf("namedexeccontext: %w", subject.ErrUniqueSubjectYear)
		}
		return fmt.Errorf("namedexeccontext: %w", err)
	}

	return nil
}

// Update replaces a user document in the database.
func (s *Store) Update(ctx context.Context, sub subject.Subject) error {
	const q = `
	UPDATE
		subjects
	SET
		"name" = :name,
		"code" = :code,
		"academic_year" = :academic_year,
		"year" = :year,
		"instructor" = :instructor,
		"semester" = :semester,
		"exam" = :exam,
		"tutorial" = :tutorial,
      "lab" = :lab,
      "assignment" = :assignment,
		"practical" = :practical,
		"date_updated" = :date_updated
	WHERE
		subject_id = :subject_id`

	if err := database.NamedExecContext(ctx, s.log, s.db, q, toDBSubject(sub)); err != nil {
		if errors.Is(err, database.ErrDBDuplicatedEntry) {
			return subject.ErrUniqueSubjectYear
		}
		return fmt.Errorf("namedexeccontext: %w", err)
	}

	return nil
}

// Delete removes a user from the database.
func (s *Store) Delete(ctx context.Context, sub subject.Subject) error {
	data := struct {
		UserID string `db:"subject_id"`
	}{
		UserID: sub.ID.String(),
	}

	const q = `
	DELETE FROM
		subjects
	WHERE
		subject_id = :subject_id`

	if err := database.NamedExecContext(ctx, s.log, s.db, q, data); err != nil {
		return fmt.Errorf("namedexeccontext: %w", err)
	}

	return nil
}

// Query retrieves a list of existing subjects from the database.
func (s *Store) Query(ctx context.Context, filter subject.QueryFilter, orderBy order.By, pageNumber int, rowsPerPage int) ([]subject.Subject, error) {
	data := map[string]interface{}{
		"offset":        (pageNumber - 1) * rowsPerPage,
		"rows_per_page": rowsPerPage,
	}

	const q = `
	SELECT
		*
	FROM
		subjects`

	buf := bytes.NewBufferString(q)
	s.applyFilter(filter, data, buf)

	orderByClause, err := orderByClause(orderBy)
	if err != nil {
		return nil, err
	}

	buf.WriteString(orderByClause)
	buf.WriteString(" OFFSET :offset ROWS FETCH NEXT :rows_per_page ROWS ONLY")

	var dbSubjects []dbSubject
	if err := database.NamedQuerySlice(ctx, s.log, s.db, buf.String(), data, &dbSubjects); err != nil {
		return nil, fmt.Errorf("namedqueryslice: %w", err)
	}

	sub, err := toCoreSubjectSlice(dbSubjects)
	if err != nil {
		return nil, err
	}

	return sub, nil
}

// QueryByID gets the specified subject from the database.
func (s *Store) QueryByID(ctx context.Context, subjectID uuid.UUID) (subject.Subject, error) {
	data := struct {
		ID string `db:"subject_id"`
	}{
		ID: subjectID.String(),
	}

	const q = `
	SELECT
        subject_id, name, instructor, year, academic_year, code, semester, exam, practical, tutorial, lab, assignment, date_created, date_updated
	FROM
		subjects
	WHERE
		subject_id = :subject_id`

	var dbSub dbSubject
	if err := database.NamedQueryStruct(ctx, s.log, s.db, q, data, &dbSub); err != nil {
		if errors.Is(err, database.ErrDBNotFound) {
			return subject.Subject{}, fmt.Errorf("db: %w", subject.ErrNotFound)
		}
		return subject.Subject{}, fmt.Errorf("db: %w", err)
	}

	return toCoreSubject(dbSub)
}

// Count returns the total number of subjects in the DB.
func (s *Store) Count(ctx context.Context, filter subject.QueryFilter) (int, error) {
	data := map[string]interface{}{}

	const q = `
	SELECT
		count(1)
	FROM
		subjects`

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
