package student

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
	ErrUniqueEmail           = errors.New("email is not unique")
	ErrAuthenticationFailure = errors.New("authentication failed")
)

// Storer interface declares the behavior this package needs to perists and
// retrieve data.
type Storer interface {
	Create(ctx context.Context, std Student) error

	Query(ctx context.Context, filter QueryFilter, orderBy order.By, pageNumber int, rowsPerPage int) ([]Student, error)
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

// Create inserts a new student into the database.
func (c *Core) Create(ctx context.Context, ns NewStudent) (Student, error) {

	now := time.Now()

	std := Student{
		ID:           uuid.New(),
		Name:         ns.Name,
		Email:        ns.Email,
		RollNumber:   ns.RollNumber,
		PhoneNumber:  ns.PhoneNumber,
		Year:         ns.Year,
		AcademicYear: ns.AcademicYear,
		DateCreated:  now,
		DateUpdated:  now,
	}

	if err := c.storer.Create(ctx, std); err != nil {
		return Student{}, fmt.Errorf("create: %w", err)
	}

	return std, nil
}

// Query retrieves a list of existing students from the database.
func (c *Core) Query(ctx context.Context, filter QueryFilter, orderBy order.By, pageNumber int, rowsPerPage int) ([]Student, error) {
	stds, err := c.storer.Query(ctx, filter, orderBy, pageNumber, rowsPerPage)
	if err != nil {

		fmt.Printf("=============: %v", stds)
		fmt.Printf("=============: %v", err)

		return nil, fmt.Errorf("query: %w", err)
	}

	return stds, nil
}

// Count returns the total number of students in the store.
func (c *Core) Count(ctx context.Context, filter QueryFilter) (int, error) {
	return c.storer.Count(ctx, filter)
}
