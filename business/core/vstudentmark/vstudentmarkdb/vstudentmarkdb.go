package vstudentmarkdb

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"

	"github.com/PhyoYazar/uas/business/core/vstudentmark"
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

// Query retrieves a list of existing students from the database.
func (s *Store) Query(ctx context.Context, filter vstudentmark.QueryFilter, orderBy order.By, pageNumber int, rowsPerPage int) ([]vstudentmark.VStudentMark, error) {
	data := map[string]interface{}{
		"offset":        (pageNumber - 1) * rowsPerPage,
		"rows_per_page": rowsPerPage,
	}

	const q = `
	SELECT
		s.student_id,
		s.roll_number,
		s.student_name,
		sm.student_mark_id,
		sm.mark,
		a.name,
		a.attribute_id
	FROM
		students s
	LEFT JOIN
		student_marks sm ON sm.student_id = s.student_id
	LEFT JOIN
		attributes a ON a.attribute_id = sm.attribute_id`

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

	var studentsMap = make(map[uuid.UUID]vstudentmark.VStudentMark)
	var result []vstudentmark.VStudentMark

	for rows.Next() {
		var (
			studentID         uuid.UUID
			studentRollNumber int
			studentName       sql.NullString
			studentMarkID     sql.NullString
			studentMark       sql.NullInt64
			attributeName     sql.NullString
			attributeID       sql.NullString
		)

		err := rows.Scan(
			&studentID, &studentRollNumber, &studentName,
			&studentMarkID, &studentMark,
			&attributeName, &attributeID,
		)
		if err != nil {
			return nil, err
		}

		student, ok := studentsMap[studentID]
		if !ok {
			student = vstudentmark.VStudentMark{
				ID:          studentID,
				RollNumber:  studentRollNumber,
				StudentName: studentName.String,
				Attributes:  []vstudentmark.VAttributes{},
			}
			studentsMap[studentID] = student
		}

		if studentMark.Valid && studentMarkID.Valid && attributeID.Valid {
			attribute := vstudentmark.VAttributes{
				StudentMarkID: uuid.MustParse(studentMarkID.String),
				AttributeID:   uuid.MustParse(attributeID.String),
				Mark:          int(studentMark.Int64),
				Name:          attributeName.String,
			}

			student.Attributes = append(student.Attributes, attribute)
		}

		studentsMap[studentID] = student
	}

	for _, value := range studentsMap {
		result = append(result, value)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

// Count returns the total number of students in the DB.
func (s *Store) Count(ctx context.Context, filter vstudentmark.QueryFilter) (int, error) {
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
