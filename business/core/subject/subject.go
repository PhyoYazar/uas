package subject

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/PhyoYazar/uas/business/data/order"
	"github.com/google/uuid"
)

// Set of error variables for CRUD operations.
var (
	ErrNotFound              = errors.New("user not found")
	ErrUniqueSubjectYear     = errors.New("name/code ,semester and academic year is not unique")
	ErrAuthenticationFailure = errors.New("authentication failed")
)

// Storer interface declares the behavior this package needs to perists and
// retrieve data.
type Storer interface {
	Create(ctx context.Context, usr Subject) error

	Query(ctx context.Context, filter QueryFilter, orderBy order.By, pageNumber int, rowsPerPage int) ([]Subject, error)
	Count(ctx context.Context, filter QueryFilter) (int, error)

	QueryByID(ctx context.Context, subjectID uuid.UUID) (Subject, error)
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
		Semester:     ns.Semester,
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

// Query retrieves a list of existing subjects from the database.
func (c *Core) Query(ctx context.Context, filter QueryFilter, orderBy order.By, pageNumber int, rowsPerPage int) ([]Subject, error) {
	subjects, err := c.storer.Query(ctx, filter, orderBy, pageNumber, rowsPerPage)
	if err != nil {

		fmt.Printf("=============: %v", subjects)
		fmt.Printf("=============: %v", err)

		return nil, fmt.Errorf("query: %w", err)
	}

	return subjects, nil
}

// QueryByID finds the user by the specified ID.
func (c *Core) QueryByID(ctx context.Context, subjectID uuid.UUID) (Subject, error) {
	sub, err := c.storer.QueryByID(ctx, subjectID)
	if err != nil {
		return Subject{}, fmt.Errorf("query: subjectID[%s]: %w", subjectID, err)
	}

	return sub, nil
}

// Count returns the total number of subjects in the store.
func (c *Core) Count(ctx context.Context, filter QueryFilter) (int, error) {
	return c.storer.Count(ctx, filter)
}
