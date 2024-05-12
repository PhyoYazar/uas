package subject

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

// Set of error variables for CRUD operations.
var (
	ErrNotFound              = errors.New("user not found")
	ErrUniqueSubjectYear     = errors.New("subject and year is not unique")
	ErrAuthenticationFailure = errors.New("authentication failed")
)

// Storer interface declares the behavior this package needs to perists and
// retrieve data.
type Storer interface {
	Create(ctx context.Context, usr Subject) error
}

// Core manages the set of APIs for user access.
type Core struct {
	storer Storer
}

// NewCore constructs a core for user api access.
func NewCore(storer Storer) *Core {
	return &Core{
		storer: storer,
	}
}

// Create inserts a new user into the database.
func (c *Core) Create(ctx context.Context, ns NewSubject) (Subject, error) {

	now := time.Now()

	usr := Subject{
		ID:           uuid.New(),
		Name:         ns.Name,
		Code:         ns.Code,
		Year:         ns.Year,
		AcademicYear: ns.AcademicYear,
		Instructor:   ns.Instructor,
		Exam:         ns.Exam,
		Practical:    100 - ns.Exam,
		DateCreated:  now,
		DateUpdated:  now,
	}

	if err := c.storer.Create(ctx, usr); err != nil {
		return Subject{}, fmt.Errorf("create: %w", err)
	}

	return usr, nil
}
