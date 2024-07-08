package vgagradedb

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"

	"github.com/PhyoYazar/uas/business/core/vgagrade"
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
func (s *Store) Query(ctx context.Context, filter vgagrade.QueryFilter, orderBy order.By, pageNumber int, rowsPerPage int) ([]vgagrade.VStudentMark, error) {
	data := map[string]interface{}{
		"offset":        (pageNumber - 1) * rowsPerPage,
		"rows_per_page": rowsPerPage,
	}

	const q = `
	SELECT
		s.student_id,
		s.student_name,
		s.roll_number,
		ga.ga_id,
		ga.slug,
		CAST(SUM(sm.mark) AS FLOAT) * 100.0 / CAST(SUM(fm.mark) AS FLOAT) AS percentage_marks
	FROM
		students s
	JOIN
		student_marks sm ON s.student_id = sm.student_id
	JOIN
		attributes a ON sm.attribute_id = a.attribute_id
	JOIN
		marks m ON m.attribute_id = a.attribute_id
	JOIN
		graduate_attributes ga ON m.ga_id = ga.ga_id
	JOIN
		full_marks fm ON fm.subject_id = sm.subject_id AND sm.attribute_id = fm.attribute_id`

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

	var studentsMap = make(map[uuid.UUID]vgagrade.VStudentMark)
	var result []vgagrade.VStudentMark

	for rows.Next() {
		var (
			studentID         uuid.UUID
			studentRollNumber int
			studentName       sql.NullString
			gaID              uuid.UUID
			gaSlug            string
			total_marks       sql.NullFloat64
		)

		err := rows.Scan(
			&studentID, &studentName, &studentRollNumber,
			&gaID, &gaSlug,
			&total_marks,
		)
		if err != nil {
			return nil, err
		}

		student, ok := studentsMap[studentID]
		if !ok {
			student = vgagrade.VStudentMark{
				ID:          studentID,
				RollNumber:  studentRollNumber,
				StudentName: studentName.String,
				Ga:          []vgagrade.VGa{},
			}
			studentsMap[studentID] = student
		}

		if total_marks.Valid {
			ga := vgagrade.VGa{
				GaID:       gaID,
				GaSlug:     gaSlug,
				TotalMarks: float64(total_marks.Float64),
			}
			student.Ga = append(student.Ga, ga)
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
