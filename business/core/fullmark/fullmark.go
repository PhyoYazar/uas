package fullmark

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
	ErrNotFound              = errors.New("full mark not found")
	ErrUniqueFullMark        = errors.New("full mark is already exists")
	ErrAuthenticationFailure = errors.New("authentication failed")
)

// Storer interface declares the behavior this package needs to perists and
// retrieve data.
type Storer interface {
	Create(ctx context.Context, cm FullMark) error
	Delete(ctx context.Context, fullMarkID string) error

	Query(ctx context.Context, filter QueryFilter, orderBy order.By, pageNumber int, rowsPerPage int) ([]FullMark, error)
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

// Create inserts a new fm into the database.
func (c *Core) Create(ctx context.Context, fm NewFullMark) (FullMark, error) {

	now := time.Now()

	fmk := FullMark{
		ID:          uuid.New(),
		SubjectID:   fm.SubjectID,
		AttributeID: fm.AttributeID,
		Mark:        fm.Mark,
		DateCreated: now,
		DateUpdated: now,
	}

	if err := c.storer.Create(ctx, fmk); err != nil {
		return FullMark{}, fmt.Errorf("create: %w", err)
	}

	return fmk, nil
}

// Delete removes a user from the database.
func (c *Core) Delete(ctx context.Context, fullMarkID string) error {
	if err := c.storer.Delete(ctx, fullMarkID); err != nil {
		return fmt.Errorf("delete: %w", err)
	}

	return nil
}

// Query retrieves a list of existing ms from the database.
func (c *Core) Query(ctx context.Context, filter QueryFilter, orderBy order.By, pageNumber int, rowsPerPage int) ([]FullMark, error) {
	m, err := c.storer.Query(ctx, filter, orderBy, pageNumber, rowsPerPage)
	if err != nil {

		fmt.Printf("=============: %v", m)
		fmt.Printf("=============: %v", err)

		return nil, fmt.Errorf("query: %w", err)
	}

	return m, nil
}

// Count returns the total number of cos in the store.
func (c *Core) Count(ctx context.Context, filter QueryFilter) (int, error) {
	return c.storer.Count(ctx, filter)
}
