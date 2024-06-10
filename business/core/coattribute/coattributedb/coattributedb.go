package coattributedb

import (
	"bytes"
	"context"
	"errors"
	"fmt"

	"github.com/PhyoYazar/uas/business/core/coattribute"
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

// Create inserts a new ga into the database.
func (s *Store) Create(ctx context.Context, cg coattribute.CoAttribute) error {

	const q = `
	INSERT INTO co_attributes
		(co_attribute_id, co_id, attribute_id, date_created, date_updated)
	VALUES
		(:co_attribute_id, :co_id, :attribute_id, :date_created, :date_updated)`

	if err := database.NamedExecContext(ctx, s.log, s.db, q, toDBCoAttribute(cg)); err != nil {
		if errors.Is(err, database.ErrDBDuplicatedEntry) {
			return fmt.Errorf("namedexeccontext: %w", coattribute.ErrUniqueCoAttribute)
		}
		return fmt.Errorf("namedexeccontext: %w", err)
	}

	return nil
}

// Delete removes a user from the database.
func (s *Store) Delete(ctx context.Context, caID string) error {
	data := struct {
		UserID string `db:"co_attribute_id"`
	}{
		UserID: caID,
	}

	const q = `
	DELETE FROM
		co_attributes
	WHERE
		co_attribute_id = :co_attribute_id`

	if err := database.NamedExecContext(ctx, s.log, s.db, q, data); err != nil {
		return fmt.Errorf("namedexeccontext: %w", err)
	}

	return nil
}

// Query retrieves a list of existing gas from the database.
func (s *Store) Query(ctx context.Context, filter coattribute.QueryFilter, orderBy order.By, pageNumber int, rowsPerPage int) ([]coattribute.CoAttribute, error) {
	data := map[string]interface{}{
		"offset":        (pageNumber - 1) * rowsPerPage,
		"rows_per_page": rowsPerPage,
	}

	const q = `
	SELECT
	*
	FROM
	co_attributes`

	buf := bytes.NewBufferString(q)
	s.applyFilter(filter, data, buf)

	orderByClause, err := orderByClause(orderBy)
	if err != nil {
		return nil, err
	}

	buf.WriteString(orderByClause)
	buf.WriteString(" OFFSET :offset ROWS FETCH NEXT :rows_per_page ROWS ONLY")

	var dbCoGa []dbCoAttribute
	if err := database.NamedQuerySlice(ctx, s.log, s.db, buf.String(), data, &dbCoGa); err != nil {
		return nil, fmt.Errorf("namedqueryslice: %w", err)
	}

	cg, err := toCoreCoGaSlice(dbCoGa)
	if err != nil {
		return nil, err
	}

	return cg, nil
}

// Count returns the total number of cos in the DB.
func (s *Store) Count(ctx context.Context, filter coattribute.QueryFilter) (int, error) {
	data := map[string]interface{}{}

	const q = `
	SELECT
		count(1)
	FROM
	co_attributes`

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
