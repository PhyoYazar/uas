package coattribute

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
	ErrUniqueCoAttribute     = errors.New("course outlines and attributes connection are already exists")
	ErrAuthenticationFailure = errors.New("authentication failed")
)

// Storer interface declares the behavior this package needs to perists and
// retrieve data.
type Storer interface {
	Create(ctx context.Context, sa CoAttribute) error
	Delete(ctx context.Context, caID string) error

	Query(ctx context.Context, filter QueryFilter, orderBy order.By, pageNumber int, rowsPerPage int) ([]CoAttribute, error)
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
func (c *Core) Create(ctx context.Context, cg NewCoAttribute) (CoAttribute, error) {

	now := time.Now()

	coga := CoAttribute{
		ID:          uuid.New(),
		CoID:        cg.CoID,
		AttributeID: cg.AttributeID,
		DateCreated: now,
		DateUpdated: now,
	}

	if err := c.storer.Create(ctx, coga); err != nil {
		return CoAttribute{}, fmt.Errorf("create: %w", err)
	}

	return coga, nil
}

// Query retrieves a list of existing gas from the database.
func (c *Core) Query(ctx context.Context, filter QueryFilter, orderBy order.By, pageNumber int, rowsPerPage int) ([]CoAttribute, error) {
	coga, err := c.storer.Query(ctx, filter, orderBy, pageNumber, rowsPerPage)
	if err != nil {

		fmt.Printf("=============: %v", coga)
		fmt.Printf("=============: %v", err)

		return nil, fmt.Errorf("query: %w", err)
	}

	return coga, nil
}

// Delete removes a user from the database.
func (c *Core) Delete(ctx context.Context, caID string) error {
	if err := c.storer.Delete(ctx, caID); err != nil {
		return fmt.Errorf("delete: %w", err)
	}

	return nil
}

// Count returns the total number of cos in the store.
func (c *Core) Count(ctx context.Context, filter QueryFilter) (int, error) {
	return c.storer.Count(ctx, filter)
}
