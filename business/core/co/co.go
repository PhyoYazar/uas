package co

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
	ErrUniqueCo              = errors.New("co already exists")
	ErrAuthenticationFailure = errors.New("authentication failed")
)

// Storer interface declares the behavior this package needs to perists and
// retrieve data.
type Storer interface {
	Create(ctx context.Context, co Co) error

	Query(ctx context.Context, filter QueryFilter, orderBy order.By, pageNumber int, rowsPerPage int) ([]Co, error)
	Count(ctx context.Context, filter QueryFilter) (int, error)
}

// // UserCore interface declares the behavior this package needs from the user
// // core domain.
// type SubjectCore interface {
// 	QueryByID(ctx context.Context, subjectID uuid.UUID) (subject.Subject, error)
// }

// Core manages the set of APIs for user access.
type Core struct {
	storer Storer
	// subjectCore SubjectCore
}

// NewCore constructs a core for user api access.
func NewCore(storer Storer) *Core {
	return &Core{
		storer: storer,
		// subjectCore: subCore,
	}
}

// Create inserts a new co into the database.
func (c *Core) Create(ctx context.Context, co NewCo) (Co, error) {

	now := time.Now()

	usr := Co{
		ID:          uuid.New(),
		Name:        co.Name,
		SubjectID:   co.SubjectID,
		DateCreated: now,
		DateUpdated: now,
	}

	if err := c.storer.Create(ctx, usr); err != nil {
		return Co{}, fmt.Errorf("create: %w", err)
	}

	return usr, nil
}

// Query retrieves a list of existing cos from the database.
func (c *Core) Query(ctx context.Context, filter QueryFilter, orderBy order.By, pageNumber int, rowsPerPage int) ([]Co, error) {
	cos, err := c.storer.Query(ctx, filter, orderBy, pageNumber, rowsPerPage)
	if err != nil {

		fmt.Printf("=============: %v", cos)
		fmt.Printf("=============: %v", err)

		return nil, fmt.Errorf("query: %w", err)
	}

	return cos, nil
}

// Count returns the total number of cos in the store.
func (c *Core) Count(ctx context.Context, filter QueryFilter) (int, error) {
	return c.storer.Count(ctx, filter)
}
