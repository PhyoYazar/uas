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
	ErrNotFound              = errors.New("student not found")
	ErrUniqueStudent         = errors.New("student is not unique")
	ErrAuthenticationFailure = errors.New("authentication failed")
)

// Storer interface declares the behavior this package needs to perists and
// retrieve data.
type Storer interface {
	Create(ctx context.Context, std Student) error
	Update(ctx context.Context, sub Student) error
	Delete(ctx context.Context, sub Student) error

	Query(ctx context.Context, filter QueryFilter, orderBy order.By, pageNumber int, rowsPerPage int) ([]Student, error)
	Count(ctx context.Context, filter QueryFilter) (int, error)

	QueryByID(ctx context.Context, studentID uuid.UUID) (Student, error)
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
		ID:            uuid.New(),
		StudentNumber: ns.StudentNumber,
		RollNumber:    ns.RollNumber,
		Year:          ns.Year,
		AcademicYear:  ns.AcademicYear,
		DateCreated:   now,
		DateUpdated:   now,
		// Email:        ns.Email,
		// PhoneNumber:  ns.PhoneNumber,
	}

	if err := c.storer.Create(ctx, std); err != nil {
		return Student{}, fmt.Errorf("create: %w", err)
	}

	return std, nil
}

// Update replaces a user document in the database.
func (c *Core) Update(ctx context.Context, std Student, uStd UpdateStudent) (Student, error) {
	if uStd.StudentNumber != nil {
		std.StudentNumber = *uStd.StudentNumber
	}
	if uStd.AcademicYear != nil {
		std.AcademicYear = *uStd.AcademicYear
	}
	if uStd.Year != (Year{}) {
		std.Year = uStd.Year
	}
	if uStd.RollNumber != nil {
		std.RollNumber = *uStd.RollNumber
	}

	std.DateUpdated = time.Now()

	if err := c.storer.Update(ctx, std); err != nil {
		return Student{}, fmt.Errorf("update: %w", err)
	}

	return std, nil
}

// Delete removes a user from the database.
func (c *Core) Delete(ctx context.Context, std Student) error {
	if err := c.storer.Delete(ctx, std); err != nil {
		return fmt.Errorf("delete: %w", err)
	}

	return nil
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

// QueryByID finds the user by the specified ID.
func (c *Core) QueryByID(ctx context.Context, studentID uuid.UUID) (Student, error) {
	std, err := c.storer.QueryByID(ctx, studentID)
	if err != nil {
		return Student{}, fmt.Errorf("query: studentID[%s]: %w", studentID, err)
	}

	return std, nil
}

// Count returns the total number of students in the store.
func (c *Core) Count(ctx context.Context, filter QueryFilter) (int, error) {
	return c.storer.Count(ctx, filter)
}
