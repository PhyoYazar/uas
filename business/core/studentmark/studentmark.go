package studentmark

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
	ErrNotFound              = errors.New("name not found")
	ErrUniqueStudentMark     = errors.New("student mark is already exist")
	ErrAuthenticationFailure = errors.New("authentication failed")
)

// Storer interface declares the behavior this package needs to perists and
// retrieve data.
type Storer interface {
	Create(ctx context.Context, ss StudentMark) error

	Query(ctx context.Context, filter QueryFilter, orderBy order.By, pageNumber int, rowsPerPage int) ([]StudentMark, error)
	Count(ctx context.Context, filter QueryFilter) (int, error)
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

// Create inserts a new ga into the database.
func (c *Core) Create(ctx context.Context, sm NewStudentMark) (StudentMark, error) {

	now := time.Now()

	mk := StudentMark{
		ID:          uuid.New(),
		StudentID:   sm.StudentID,
		SubjectID:   sm.SubjectID,
		AttributeID: sm.AttributeID,
		Mark:        sm.Mark,
		DateCreated: now,
		DateUpdated: now,
	}

	if err := c.storer.Create(ctx, mk); err != nil {
		return StudentMark{}, fmt.Errorf("create: %w", err)
	}

	return mk, nil
}

// Query retrieves a list of existing gas from the database.
func (c *Core) Query(ctx context.Context, filter QueryFilter, orderBy order.By, pageNumber int, rowsPerPage int) ([]StudentMark, error) {
	sm, err := c.storer.Query(ctx, filter, orderBy, pageNumber, rowsPerPage)
	if err != nil {

		fmt.Printf("=============: %v", sm)
		fmt.Printf("=============: %v", err)

		return nil, fmt.Errorf("query: %w", err)
	}

	return sm, nil
}

// Count returns the total number of cos in the store.
func (c *Core) Count(ctx context.Context, filter QueryFilter) (int, error) {
	return c.storer.Count(ctx, filter)
}
