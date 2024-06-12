package fullmarkdb

import (
	"bytes"
	"context"
	"errors"
	"fmt"

	"github.com/PhyoYazar/uas/business/core/fullmark"
	"github.com/PhyoYazar/uas/business/data/order"
	database "github.com/PhyoYazar/uas/business/sys/database/pgx"
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

// Create inserts a new cm into the database.
func (s *Store) Create(ctx context.Context, fm fullmark.FullMark) error {

	const q = `
	INSERT INTO full_marks
		(full_mark_id, subject_id, attribute_id, mark, date_created, date_updated)
	VALUES
		(:full_mark_id, :subject_id, :attribute_id, :mark, :date_created, :date_updated)`

	if err := database.NamedExecContext(ctx, s.log, s.db, q, toDBFullMark(fm)); err != nil {
		if errors.Is(err, database.ErrDBDuplicatedEntry) {
			return fmt.Errorf("namedexeccontext: %w", fullmark.ErrUniqueFullMark)
		}
		return fmt.Errorf("namedexeccontext: %w", err)
	}

	return nil
}

// Delete removes a user from the database.
func (s *Store) Delete(ctx context.Context, fullMarkID string) error {
	data := struct {
		UserID string `db:"full_mark_id"`
	}{
		UserID: fullMarkID,
	}

	const q = `
	DELETE FROM
		full_marks
	WHERE
		full_mark_id = :full_mark_id`

	if err := database.NamedExecContext(ctx, s.log, s.db, q, data); err != nil {
		return fmt.Errorf("namedexeccontext: %w", err)
	}

	return nil
}

// Query retrieves a list of existing gas from the database.
func (s *Store) Query(ctx context.Context, filter fullmark.QueryFilter, orderBy order.By, pageNumber int, rowsPerPage int) ([]fullmark.FullMark, error) {
	data := map[string]interface{}{
		"offset":        (pageNumber - 1) * rowsPerPage,
		"rows_per_page": rowsPerPage,
	}

	const q = `
	SELECT
	*
	FROM
	full_marks`

	buf := bytes.NewBufferString(q)
	s.applyFilter(filter, data, buf)

	orderByClause, err := orderByClause(orderBy)
	if err != nil {
		return nil, err
	}

	buf.WriteString(orderByClause)
	buf.WriteString(" OFFSET :offset ROWS FETCH NEXT :rows_per_page ROWS ONLY")

	var dbFullMark []dbFullMark
	if err := database.NamedQuerySlice(ctx, s.log, s.db, buf.String(), data, &dbFullMark); err != nil {
		return nil, fmt.Errorf("namedqueryslice: %w", err)
	}

	m, err := toCoreFullMarkSlice(dbFullMark)
	if err != nil {
		return nil, err
	}

	return m, nil
}

// Count returns the total number of cos in the DB.
func (s *Store) Count(ctx context.Context, filter fullmark.QueryFilter) (int, error) {
	data := map[string]interface{}{}

	const q = `
	SELECT
		count(1)
	FROM
	full_marks`

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
