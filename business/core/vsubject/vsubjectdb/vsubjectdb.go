package vsubjectdb

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/PhyoYazar/uas/business/core/subject"
	"github.com/PhyoYazar/uas/business/core/vsubject"
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
func (s *Store) QueryByID(ctx context.Context, subjectID uuid.UUID) (vsubject.VSubject, error) {
	data := struct {
		ID string `db:"subject_id"`
	}{
		ID: subjectID.String(),
	}

	const q = `
	SELECT
		s.subject_id, s.name, s.code, s.academic_year, s.instructor, s.semester, co.co_id, co.name, ga.ga_id, ga.name, ga.slug
	FROM
		subjects s
	LEFT JOIN
		course_outlines co ON co.subject_id = s.subject_id
	LEFT JOIN
		co_ga cg ON co.co_id = cg.co_id
	LEFT JOIN
		graduate_attributes ga ON cg.ga_id = ga.ga_id
	WHERE
		s.subject_id = :subject_id`

	rows, err := database.NamedQueryRows(ctx, s.log, s.db, q, data)
	if err != nil {
		if errors.Is(err, database.ErrDBNotFound) {
			return vsubject.VSubject{}, fmt.Errorf("db: %w", subject.ErrNotFound)
		}
		return vsubject.VSubject{}, fmt.Errorf("db: %w", err)
	}

	defer rows.Close()

	// subjectMap := make(map[uuid.UUID]*vsubject.VSubject)
	// coMap := make(map[uuid.UUID]*vsubject.VCo)

	// for rows.Next() {
	// 	var subjectID, coID, gaID uuid.UUID
	// 	var subjectName, subjectCode, academicYear, instructor, semester, coName, gaName, gaSlug sql.NullString

	// 	err := rows.Scan(&subjectID, &subjectName, &subjectCode, &academicYear, &instructor, &semester, &coID, &coName, &gaID, &gaName, &gaSlug)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	// Initialize subject if it doesn't exist in the map
	// 	if _, exists := subjectMap[subjectID]; !exists {
	// 		subjectMap[subjectID] = &vsubject.VSubject{
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
	// 		subjectMap[subjectID].Co = append(subjectMap[subjectID].Co, *coMap[coID])
	// 	}

	// 	// Add graduate attribute to the course outline
	// 	if gaID != uuid.Nil {
	// 		ga := vsubject.VGa{
	// 			ID:   gaID,
	// 			Name: gaName.String,
	// 			Slug: gaSlug.String,
	// 		}
	// 		coMap[coID].Ga = append(coMap[coID].Ga, ga)
	// 	}
	// }

	// // Convert map to slice
	// var subjects []vsubject.VSubject
	// for _, subject := range subjectMap {
	// 	subjects = append(subjects, *subject)
	// }

	var subject vsubject.VSubject
	coMap := make(map[uuid.UUID]*vsubject.VCo)

	for rows.Next() {
		var subjectID, coID, gaID uuid.UUID
		var subjectName, subjectCode, academicYear, instructor, semester, coName, gaName, gaSlug sql.NullString

		err := rows.Scan(&subjectID, &subjectName, &subjectCode, &academicYear, &instructor, &semester, &coID, &coName, &gaID, &gaName, &gaSlug)
		if err != nil {
			log.Fatal(err)
		}

		// Initialize subject if it hasn't been initialized yet
		if subject.ID == uuid.Nil {
			subject = vsubject.VSubject{
				ID:           subjectID,
				Name:         subjectName.String,
				Code:         subjectCode.String,
				AcademicYear: academicYear.String,
				Instructor:   instructor.String,
				Semester:     semester.String,
				Co:           []vsubject.VCo{},
			}
		}

		// Initialize course outline if it doesn't exist in the map
		if _, exists := coMap[coID]; !exists {
			coMap[coID] = &vsubject.VCo{
				ID:   coID,
				Name: coName.String,
				Ga:   []vsubject.VGa{},
			}
			subject.Co = append(subject.Co, *coMap[coID])
		}

		// Add graduate attribute to the course outline
		if gaID != uuid.Nil {
			ga := vsubject.VGa{
				ID:   gaID,
				Name: gaName.String,
				Slug: gaSlug.String,
			}
			// Find the course outline in the subject's list and update its Ga field
			for i := range subject.Co {
				if subject.Co[i].ID == coID {
					subject.Co[i].Ga = append(subject.Co[i].Ga, ga)
					break
				}
			}
		}
	}

	// Check for errors during row iteration
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	if subject.ID == uuid.Nil {
		return vsubject.VSubject{}, errors.New("subject not found")
	}

	return subject, nil
}
