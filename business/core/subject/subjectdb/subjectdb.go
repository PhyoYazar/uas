package subjectdb

import (
	"context"
	"errors"
	"fmt"

	"github.com/PhyoYazar/uas/business/core/subject"
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

// Create inserts a new subject into the database.
func (s *Store) Create(ctx context.Context, sub subject.Subject) error {
	const q = `
	INSERT INTO subjects
		(subject_id, name, code, year, academic_year, instructor, exam, practical, date_created, date_updated)
	VALUES
		(:subject_id, :name, :code, :year, :academic_year, :instructor, :exam, :practical, :date_created, :date_updated)`

	if err := database.NamedExecContext(ctx, s.log, s.db, q, toDBSubject(sub)); err != nil {
		if errors.Is(err, database.ErrDBDuplicatedEntry) {
			return fmt.Errorf("namedexeccontext: %w", subject.ErrUniqueSubjectYear)
		}
		return fmt.Errorf("namedexeccontext: %w", err)
	}

	return nil
}
