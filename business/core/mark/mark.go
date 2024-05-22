package mark

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
	ErrUniqueMark            = errors.New("course outlines and mark type are already exists")
	ErrAuthenticationFailure = errors.New("authentication failed")
)

// Storer interface declares the behavior this package needs to perists and
// retrieve data.
type Storer interface {
	Create(ctx context.Context, cm Mark) error

	Query(ctx context.Context, filter QueryFilter, orderBy order.By, pageNumber int, rowsPerPage int) ([]Mark, error)
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

// Create inserts a new cm into the database.
func (c *Core) Create(ctx context.Context, m NewMark) (Mark, error) {

	now := time.Now()

	mk := Mark{
		ID:          uuid.New(),
		SubjectID:   m.SubjectID,
		GaID:        m.GaID,
		AttributeID: m.AttributeID,
		Mark:        m.Mark,
		DateCreated: now,
		DateUpdated: now,
	}

	if err := c.storer.Create(ctx, mk); err != nil {
		return Mark{}, fmt.Errorf("create: %w", err)
	}

	return mk, nil
}

// Query retrieves a list of existing ms from the database.
func (c *Core) Query(ctx context.Context, filter QueryFilter, orderBy order.By, pageNumber int, rowsPerPage int) ([]Mark, error) {
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
