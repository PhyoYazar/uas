package markdb

import (
	"bytes"
	"context"
	"errors"
	"fmt"

	"github.com/PhyoYazar/uas/business/core/mark"
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
func (s *Store) Create(ctx context.Context, cm mark.Mark) error {

	const q = `
	INSERT INTO marks
		(mark_id, subject_id, ga_id, attribute_id, mark, date_created, date_updated)
	VALUES
		(:mark_id, :subject_id, :ga_id, :attribute_id, :mark, :date_created, :date_updated)`

	if err := database.NamedExecContext(ctx, s.log, s.db, q, toDBMark(cm)); err != nil {
		if errors.Is(err, database.ErrDBDuplicatedEntry) {
			return fmt.Errorf("namedexeccontext: %w", mark.ErrUniqueMark)
		}
		return fmt.Errorf("namedexeccontext: %w", err)
	}

	return nil
}

// Query retrieves a list of existing gas from the database.
func (s *Store) Query(ctx context.Context, filter mark.QueryFilter, orderBy order.By, pageNumber int, rowsPerPage int) ([]mark.Mark, error) {
	data := map[string]interface{}{
		"offset":        (pageNumber - 1) * rowsPerPage,
		"rows_per_page": rowsPerPage,
	}

	const q = `
	SELECT
	*
	FROM
	marks`

	buf := bytes.NewBufferString(q)
	s.applyFilter(filter, data, buf)

	orderByClause, err := orderByClause(orderBy)
	if err != nil {
		return nil, err
	}

	buf.WriteString(orderByClause)
	buf.WriteString(" OFFSET :offset ROWS FETCH NEXT :rows_per_page ROWS ONLY")

	var dbMark []dbMark
	if err := database.NamedQuerySlice(ctx, s.log, s.db, buf.String(), data, &dbMark); err != nil {
		return nil, fmt.Errorf("namedqueryslice: %w", err)
	}

	m, err := toCoreMarkSlice(dbMark)
	if err != nil {
		return nil, err
	}

	return m, nil
}

// Count returns the total number of cos in the DB.
func (s *Store) Count(ctx context.Context, filter mark.QueryFilter) (int, error) {
	data := map[string]interface{}{}

	const q = `
	SELECT
		count(1)
	FROM
	marks`

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
