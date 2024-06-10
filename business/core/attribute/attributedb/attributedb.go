package attributedb

import (
	"bytes"
	"context"
	"errors"
	"fmt"

	"github.com/PhyoYazar/uas/business/core/attribute"
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
func (s *Store) Create(ctx context.Context, att attribute.Attribute) error {
	const q = `
	INSERT INTO attributes
		(attribute_id, name, type, instance, date_created, date_updated)
	VALUES
		(:attribute_id, :name, :type, :instance, :date_created, :date_updated)`

	if err := database.NamedExecContext(ctx, s.log, s.db, q, toDBAttribute(att)); err != nil {
		if errors.Is(err, database.ErrDBDuplicatedEntry) {
			return fmt.Errorf("namedexeccontext: %w", attribute.ErrUniqueAttribute)
		}
		return fmt.Errorf("namedexeccontext: %w", err)
	}

	return nil
}

// Delete removes a user from the database.
func (s *Store) Delete(ctx context.Context, attribute attribute.Attribute) error {
	data := struct {
		UserID string `db:"attribute_id"`
	}{
		UserID: attribute.ID.String(),
	}

	const q = `
	DELETE FROM
		attributes
	WHERE
		attribute_id = :attribute_id`

	if err := database.NamedExecContext(ctx, s.log, s.db, q, data); err != nil {
		return fmt.Errorf("namedexeccontext: %w", err)
	}

	return nil
}

// Query retrieves a list of existing gas from the database.
func (s *Store) Query(ctx context.Context, filter attribute.QueryFilter, orderBy order.By, pageNumber int, rowsPerPage int) ([]attribute.Attribute, error) {
	data := map[string]interface{}{
		"offset":        (pageNumber - 1) * rowsPerPage,
		"rows_per_page": rowsPerPage,
	}

	const q = `
	SELECT
	*
	FROM
	attributes`

	buf := bytes.NewBufferString(q)
	s.applyFilter(filter, data, buf)

	orderByClause, err := orderByClause(orderBy)
	if err != nil {
		return nil, err
	}

	buf.WriteString(orderByClause)
	buf.WriteString(" OFFSET :offset ROWS FETCH NEXT :rows_per_page ROWS ONLY")

	var dbAttribute []dbAttribute
	if err := database.NamedQuerySlice(ctx, s.log, s.db, buf.String(), data, &dbAttribute); err != nil {
		return nil, fmt.Errorf("namedqueryslice: %w", err)
	}

	att, err := toCoreAttributeSlice(dbAttribute)
	if err != nil {
		return nil, err
	}

	return att, nil
}

// QueryByID gets the specified subject from the database.
func (s *Store) QueryByID(ctx context.Context, attributeID uuid.UUID) (attribute.Attribute, error) {
	data := struct {
		ID string `db:"attribute_id"`
	}{
		ID: attributeID.String(),
	}

	const q = `
	SELECT
        attribute_id, name, type, instance, date_created, date_updated
	FROM
		attributes
	WHERE
		attribute_id = :attribute_id`

	var dbAtt dbAttribute
	if err := database.NamedQueryStruct(ctx, s.log, s.db, q, data, &dbAtt); err != nil {
		if errors.Is(err, database.ErrDBNotFound) {
			return attribute.Attribute{}, fmt.Errorf("db: %w", attribute.ErrNotFound)
		}
		return attribute.Attribute{}, fmt.Errorf("db: %w", err)
	}

	return toCoreAttribute(dbAtt)
}

// Count returns the total number of cos in the DB.
func (s *Store) Count(ctx context.Context, filter attribute.QueryFilter) (int, error) {
	data := map[string]interface{}{}

	const q = `
	SELECT
		count(1)
	FROM
	attributes`

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
