package vcogradedb

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"

	"github.com/PhyoYazar/uas/business/core/vcograde"
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
func (s *Store) Query(ctx context.Context, filter vcograde.QueryFilter, orderBy order.By, pageNumber int, rowsPerPage int) ([]vcograde.VStudentMark, error) {
	data := map[string]interface{}{
		"offset":        (pageNumber - 1) * rowsPerPage,
		"rows_per_page": rowsPerPage,
	}

	const q = `
	SELECT
     s.student_id,
     s.student_name,
     s.roll_number,
     co.co_id,
     co.name AS co_name,
     co.instance co_instance,
     SUM(fm.mark) AS total_full_marks,
     CAST(SUM(sm.mark) AS FLOAT) AS total_marks
 FROM
     students s
 JOIN
     student_marks sm ON s.student_id = sm.student_id
 JOIN
     attributes a ON sm.attribute_id = a.attribute_id
 JOIN
     co_attributes ca ON a.attribute_id = ca.attribute_id
 JOIN
     course_outlines co ON ca.co_id = co.co_id
 JOIN
     full_marks fm ON sm.subject_id = fm.subject_id AND sm.attribute_id = fm.attribute_id
 `

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

	var studentsMap = make(map[uuid.UUID]vcograde.VStudentMark)
	var result []vcograde.VStudentMark

	for rows.Next() {
		var (
			studentID         uuid.UUID
			studentRollNumber int
			studentName       sql.NullString
			coID              uuid.UUID
			coName            string
			coInstance        int
			total_full_marks  int
			total_marks       float64
		)

		err := rows.Scan(
			&studentID, &studentName, &studentRollNumber,
			&coID, &coName, &coInstance,
			&total_full_marks, &total_marks,
		)
		if err != nil {
			return nil, err
		}

		student, ok := studentsMap[studentID]
		if !ok {
			student = vcograde.VStudentMark{
				ID:          studentID,
				RollNumber:  studentRollNumber,
				StudentName: studentName.String,
				Co:          []vcograde.VCo{},
			}
			studentsMap[studentID] = student
		}

		co := vcograde.VCo{
			CoID:           coID,
			CoName:         coName,
			CoInstance:     coInstance,
			TotalFullMarks: total_full_marks,
			TotalMarks:     total_marks,
		}
		student.Co = append(student.Co, co)

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
