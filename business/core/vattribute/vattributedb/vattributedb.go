package vattributedb

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/PhyoYazar/uas/business/core/vattribute"
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

// Delete removes a user from the database.
func (s *Store) RemoveMarks(ctx context.Context, ra vattribute.VRemoveAttribute) error {
	data := struct {
		SubjectID   string `db:"subject_id"`
		AttributeID string `db:"attribute_id"`
	}{
		SubjectID:   ra.SubjectID.String(),
		AttributeID: ra.AttributeID.String(),
	}

	const q = `
	DELETE FROM
		marks
	WHERE
		subject_id = :subject_id
	AND
		attribute_id = :attribute_id`

	if err := database.NamedExecContext(ctx, s.log, s.db, q, data); err != nil {
		return fmt.Errorf("namedexeccontext: %w", err)
	}

	return nil
}

// Delete removes a user from the database.
func (s *Store) RemoveFullMarks(ctx context.Context, ra vattribute.VRemoveAttribute) error {
	data := struct {
		SubjectID   string `db:"subject_id"`
		AttributeID string `db:"attribute_id"`
	}{
		SubjectID:   ra.SubjectID.String(),
		AttributeID: ra.AttributeID.String(),
	}

	const q = `
	DELETE FROM
		full_marks
	WHERE
		subject_id = :subject_id
	AND
		attribute_id = :attribute_id`

	if err := database.NamedExecContext(ctx, s.log, s.db, q, data); err != nil {
		return fmt.Errorf("namedexeccontext: %w", err)
	}

	return nil
}

// Delete removes a user from the database.
func (s *Store) RemoveCoAttributes(ctx context.Context, ra vattribute.VRemoveAttribute) error {
	data := struct {
		SubjectID   string `db:"subject_id"`
		AttributeID string `db:"attribute_id"`
	}{
		SubjectID:   ra.SubjectID.String(),
		AttributeID: ra.AttributeID.String(),
	}

	const q = `
	DELETE FROM
		co_attributes ca
	USING
		course_outlines co
	WHERE
		co.co_id = ca.co_id
	AND
		co.subject_id = :subject_id
	AND
		ca.attribute_id = :attribute_id`

	if err := database.NamedExecContext(ctx, s.log, s.db, q, data); err != nil {
		return fmt.Errorf("namedexeccontext: %w", err)
	}

	return nil
}

func (s *Store) QueryAttributeWithGaMark(ctx context.Context, filter vattribute.QueryFilter, orderBy order.By, pageNumber int, rowsPerPage int, subjectID uuid.UUID) ([]vattribute.VAttributeWithGaMark, error) {

	data := map[string]interface{}{
		"subject_id":    subjectID.String(),
		"offset":        (pageNumber - 1) * rowsPerPage,
		"rows_per_page": rowsPerPage,
	}

	const q = `
		SELECT
			a.attribute_id,
         a.name,
         a.instance,
         a.type,
			fm.mark full_mark,
         m.mark_id,
         m.ga_id,
         m.mark,
			ga.slug ga_slug,
			co.co_id,
			co.name co_name,
			co.instance co_instance
		FROM
			attributes a
		LEFT JOIN
	 		marks m ON m.attribute_id = a.attribute_id
		LEFT JOIN
			graduate_attributes ga ON ga.ga_id = m.ga_id
		LEFT JOIN
			co_ga cg ON cg.ga_id = ga.ga_id
		LEFT JOIN
			co_attributes ca ON ca.attribute_id = a.attribute_id
		LEFT JOIN
			course_outlines co ON co.co_id = ca.co_id AND co.co_id = cg.co_id
		LEFT JOIN
	 		full_marks fm ON fm.attribute_id = a.attribute_id AND fm.subject_id = :subject_id
		WHERE
			m.subject_id = :subject_id
		AND
			co.subject_id = :subject_id`

	buf := bytes.NewBufferString(q)
	s.applyFilter(filter, data, buf)

	orderByClause, err := orderByClause(orderBy)
	if err != nil {
		return nil, err
	}

	buf.WriteString(orderByClause)
	// buf.WriteString(" OFFSET :offset ROWS FETCH NEXT :rows_per_page ROWS ONLY")

	rows, err := database.NamedQueryRows(ctx, s.log, s.db, buf.String(), data)
	if err != nil {
		return nil, fmt.Errorf("namedqueryrows: %w", err)
	}

	var attributesMap = make(map[uuid.UUID]vattribute.VAttributeWithGaMark)
	var result []vattribute.VAttributeWithGaMark

	for rows.Next() {
		var (
			attributeID       uuid.UUID
			attributeName     string
			attributeInstance int
			attributeType     string
			fullMark          sql.NullInt64
			markID            sql.NullString
			gaID              sql.NullString
			mark              sql.NullInt64
			gaSlug            sql.NullString
			coID              sql.NullString
			coName            sql.NullString
			coInstance        int
		)

		err := rows.Scan(
			&attributeID, &attributeName, &attributeInstance, &attributeType,
			&fullMark,
			&markID, &gaID, &mark, &gaSlug,
			&coID, &coName, &coInstance,
		)

		if err != nil {
			return nil, err
		}

		attribute, ok := attributesMap[attributeID]
		if !ok {
			attribute = vattribute.VAttributeWithGaMark{
				ID:       attributeID,
				Name:     attributeName,
				Instance: attributeInstance,
				Type:     attributeType,
				FullMark: int(fullMark.Int64),
				Marks:    []vattribute.VMark{},
			}
			attributesMap[attributeID] = attribute
		}

		// Append Co if not NULL
		if coID.Valid && coName.Valid {
			co := vattribute.VCo{
				ID:       uuid.MustParse(coID.String),
				Name:     coName.String,
				Instance: coInstance,
			}

			if coIsExist := existInSlice(attribute.Co, co); !coIsExist {
				attribute.Co = append(attribute.Co, co)
			}
		}

		// Append Mark if not NULL
		if markID.Valid && gaID.Valid && mark.Valid && gaSlug.Valid {
			mark := vattribute.VMark{
				ID:     uuid.MustParse(markID.String),
				Mark:   int(mark.Int64),
				GaID:   uuid.MustParse(gaID.String),
				GaSlug: gaSlug.String,
			}

			if markIsExist := existInSlice(attribute.Marks, mark); !markIsExist {
				attribute.Marks = append(attribute.Marks, mark)
			}
		}

		attributesMap[attributeID] = attribute
	}

	for _, value := range attributesMap {
		result = append(result, value)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	// fmt.Printf("-------------------------------- ==========================>> ")
	// fmt.Printf("--------------------------------> %v", attributesMap)

	return result, nil
}

// Query retrieves a list of existing subjects from the database.
func (s *Store) Query(ctx context.Context, filter vattribute.QueryFilter, orderBy order.By, pageNumber int, rowsPerPage int, subjectID uuid.UUID) ([]vattribute.VAttribute, error) {
	data := struct {
		ID string `db:"subject_id"`
	}{
		ID: subjectID.String(),
	}

	const q = `
	SELECT
		a.attribute_id, a.name, a.instance, a.type, fm.mark, a.date_created, a.date_updated,
		co.co_id, co.name, co.instance, ca.co_mark, ca.co_attribute_id,
		ga.ga_id, ga.name, ga.slug, m.ga_mark, m.mark_id
	FROM
		attributes a
	JOIN
		co_attributes ca ON ca.attribute_id = a.attribute_id
	JOIN
		course_outlines co ON co.co_id = ca.co_id
	JOIN
		full_marks fm ON fm.attribute_id = a.attribute_id
	JOIN
		marks m ON m.attribute_id = a.attribute_id
	JOIN
		graduate_attributes ga ON ga.ga_id = m.ga_id
	JOIN
		co_ga cg ON cg.ga_id = ga.ga_id
	WHERE
		cg.co_id = ca.co_id
	AND
		co.subject_id = :subject_id
	AND
		m.subject_id = :subject_id`

	buf := bytes.NewBufferString(q)
	// s.applyFilter(filter, pdata, buf)

	// orderByClause, err := orderByClause(orderBy)
	// if err != nil {
	// 	return nil, err
	// }

	// buf.WriteString(orderByClause)
	buf.WriteString(" ORDER BY a.name, a.instance")
	// buf.WriteString(" OFFSET :offset ROWS FETCH NEXT :rows_per_page ROWS ONLY")

	rows, err := database.NamedQueryRows(ctx, s.log, s.db, buf.String(), data)
	if err != nil {
		return nil, fmt.Errorf("namedqueryrows: %w", err)
	}

	var attributesMap = make(map[uuid.UUID]vattribute.VAttribute)
	var result []vattribute.VAttribute

	for rows.Next() {
		var (
			attributeID   uuid.UUID
			attributeName string
			instance      int
			attributeType string
			fullMark      sql.NullInt64
			dateCreated   time.Time
			dateUpdated   time.Time
			coID          sql.NullString
			coAttributeID sql.NullString
			coName        sql.NullString
			coInstance    sql.NullInt64
			coMark        sql.NullInt64
			gaID          sql.NullString
			markID        sql.NullString
			gaName        sql.NullString
			gaSlug        sql.NullString
			gaMark        sql.NullInt64
		)

		err := rows.Scan(
			&attributeID, &attributeName, &instance, &attributeType, &fullMark, &dateCreated, &dateUpdated,
			&coID, &coName, &coInstance, &coMark, &coAttributeID,
			&gaID, &gaName, &gaSlug, &gaMark, &markID,
		)
		if err != nil {
			return nil, err
		}

		// Retrieve or create attribute
		attribute, ok := attributesMap[attributeID]
		if !ok {
			attribute = vattribute.VAttribute{
				ID:          attributeID,
				Name:        attributeName,
				Instance:    instance,
				Type:        attributeType,
				FullMark:    int(fullMark.Int64),
				DateCreated: dateCreated,
				DateUpdated: dateUpdated,
				Co:          []vattribute.VCo{},
				Ga:          []vattribute.VGa{},
			}
			attributesMap[attributeID] = attribute
		}

		// Append Co if not NULL
		if coID.Valid && coName.Valid && coInstance.Valid && coAttributeID.Valid {
			co := vattribute.VCo{
				ID:            uuid.MustParse(coID.String),
				Name:          coName.String,
				Instance:      int(coInstance.Int64),
				CoMark:        int(coMark.Int64),
				CoAttributeID: uuid.MustParse(coAttributeID.String),
			}

			if coIsExist := existInSlice(attribute.Co, co); !coIsExist {
				attribute.Co = append(attribute.Co, co)
			}
		}

		// Append Ga if not NULL
		if gaID.Valid && gaName.Valid && gaSlug.Valid && markID.Valid {
			ga := vattribute.VGa{
				ID:     uuid.MustParse(gaID.String),
				Name:   gaName.String,
				Slug:   gaSlug.String,
				GaMark: int(gaMark.Int64),
				MarkID: uuid.MustParse(markID.String),
			}

			attribute.Ga = append(attribute.Ga, ga)
		}

		attributesMap[attributeID] = attribute
	}

	for _, value := range attributesMap {
		result = append(result, value)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
