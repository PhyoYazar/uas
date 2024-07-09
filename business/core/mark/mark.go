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
	Update(ctx context.Context, cm Mark) error
	Delete(ctx context.Context, mark string) error

	Query(ctx context.Context, filter QueryFilter, orderBy order.By, pageNumber int, rowsPerPage int) ([]Mark, error)
	Count(ctx context.Context, filter QueryFilter) (int, error)

	QueryByID(ctx context.Context, markID uuid.UUID) (Mark, error)
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

// Update replaces a user document in the database.
func (c *Core) Update(ctx context.Context, mark Mark, uMark UpdateMark) (Mark, error) {

	if uMark.GaMark != nil {
		mark.GaMark = *uMark.GaMark
	}

	mark.DateUpdated = time.Now()

	if err := c.storer.Update(ctx, mark); err != nil {
		return Mark{}, fmt.Errorf("update: %w", err)
	}

	return mark, nil
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

// Delete removes a user from the database.
func (c *Core) Delete(ctx context.Context, mark string) error {
	if err := c.storer.Delete(ctx, mark); err != nil {
		return fmt.Errorf("delete: %w", err)
	}

	return nil
}

// QueryByID finds the user by the specified ID.
func (c *Core) QueryByID(ctx context.Context, markID uuid.UUID) (Mark, error) {
	std, err := c.storer.QueryByID(ctx, markID)

	if err != nil {
		return Mark{}, fmt.Errorf("query: markID[%s]: %w", markID, err)
	}

	return std, nil
}

// Count returns the total number of cos in the store.
func (c *Core) Count(ctx context.Context, filter QueryFilter) (int, error) {
	return c.storer.Count(ctx, filter)
}
