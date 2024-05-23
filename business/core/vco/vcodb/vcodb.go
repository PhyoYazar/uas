package vcodb

import (
	"context"
	"errors"
	"fmt"

	"github.com/PhyoYazar/uas/business/core/vco"
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

// QueryByID gets the specified subject from the database.
func (s *Store) QueryByID(ctx context.Context, coID uuid.UUID) (vco.VCo, error) {
	data := struct {
		ID string `db:"co_id"`
	}{
		ID: coID.String(),
	}

	const q = `
	SELECT
		s.subject_id, s.name, s.code, s.academic_year, s.instructor, s.semester, co.co_id, co.name,co.instance, ga.ga_id, ga.name, ga.slug, co.date_created, co.date_updated
	FROM
		course_outlines co
	LEFT JOIN
		subjects s ON co.subject_id = s.subject_id
	LEFT JOIN
		co_ga cg ON cg.co_id = co.co_id
	LEFT JOIN
		graduate_attributes ga ON ga.ga_id = cg.ga_id
	WHERE
		co.co_id = :co_id`

	rows, err := database.NamedQueryRows(ctx, s.log, s.db, q, data)
	if err != nil {
		if errors.Is(err, database.ErrDBNotFound) {
			return vco.VCo{}, fmt.Errorf("db: %w", vco.ErrNotFound)
		}
		return vco.VCo{}, fmt.Errorf("db: %w", err)
	}

	defer rows.Close()

	// var subject vsubject.VSubject
	// coMap := make(map[uuid.UUID]*vsubject.VCo)

	// for rows.Next() {
	// 	var subjectID, coID, gaID uuid.UUID
	// 	var subjectName, subjectCode, academicYear, instructor, semester, coName, gaName, gaSlug sql.NullString

	// 	err := rows.Scan(&subjectID, &subjectName, &subjectCode, &academicYear, &instructor, &semester, &coID, &coName, &gaID, &gaName, &gaSlug)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	// Initialize subject if it hasn't been initialized yet
	// 	if subject.ID == uuid.Nil {
	// 		subject = vsubject.VSubject{
	// 			ID:           subjectID,
	// 			Name:         subjectName.String,
	// 			Code:         subjectCode.String,
	// 			AcademicYear: academicYear.String,
	// 			Instructor:   instructor.String,
	// 			Semester:     semester.String,
	// 			Co:           []vsubject.VCo{},
	// 		}
	// 	}

	// 	// Initialize course outline if it doesn't exist in the map
	// 	if _, exists := coMap[coID]; !exists {
	// 		coMap[coID] = &vsubject.VCo{
	// 			ID:   coID,
	// 			Name: coName.String,
	// 			Ga:   []vsubject.VGa{},
	// 		}
	// 		subject.Co = append(subject.Co, *coMap[coID])
	// 	}

	// 	// Add graduate attribute to the course outline
	// 	if gaID != uuid.Nil {
	// 		ga := vsubject.VGa{
	// 			ID:   gaID,
	// 			Name: gaName.String,
	// 			Slug: gaSlug.String,
	// 		}
	// 		// Find the course outline in the subject's list and update its Ga field
	// 		for i := range subject.Co {
	// 			if subject.Co[i].ID == coID {
	// 				subject.Co[i].Ga = append(subject.Co[i].Ga, ga)
	// 				break
	// 			}
	// 		}
	// 	}
	// }

	// // Check for errors during row iteration
	// if err := rows.Err(); err != nil {
	// 	log.Fatal(err)
	// }

	// if subject.ID == uuid.Nil {
	// 	return vsubject.VSubject{}, errors.New("subject not found")
	// }

	var co vco.VCo
	var gaMap = make(map[uuid.UUID]vco.VGa)

	for rows.Next() {
		var row vco.CourseOutlineRow
		err := rows.Scan(
			&row.SubjectID, &row.SubjectName, &row.SubjectCode, &row.AcademicYear, &row.Instructor, &row.Semester,
			&row.CoID, &row.CoName, &row.CoInstance,
			&row.GaID, &row.GaName, &row.GaSlug,
			&row.DateCreated, &row.DateUpdated,
		)
		if err != nil {
			return vco.VCo{}, err
		}

		if co.ID == uuid.Nil {
			co = vco.VCo{
				ID:       row.CoID,
				Name:     row.CoName,
				Instance: row.CoInstance,
				Subject: vco.VSubject{
					ID:           row.SubjectID,
					Name:         row.SubjectName,
					Code:         row.SubjectCode,
					AcademicYear: row.AcademicYear,
					Instructor:   row.Instructor,
					Semester:     row.Semester,
				},
				DateCreated: row.DateCreated,
				DateUpdated: row.DateUpdated,
				Ga:          []vco.VGa{},
			}
		}

		if row.GaID != uuid.Nil {
			if _, exists := gaMap[row.GaID]; !exists {
				ga := vco.VGa{
					ID:   row.GaID,
					Name: row.GaName,
					Slug: row.GaSlug,
				}
				co.Ga = append(co.Ga, ga)
				gaMap[row.GaID] = ga
			}
		}
	}

	if err := rows.Err(); err != nil {
		return vco.VCo{}, err
	}

	return co, nil
}
