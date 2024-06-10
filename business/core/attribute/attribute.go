package attribute

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
	ErrUniqueAttribute       = errors.New("attribute already exists")
	ErrAuthenticationFailure = errors.New("authentication failed")
)

// Storer interface declares the behavior this package needs to perists and
// retrieve data.
type Storer interface {
	Create(ctx context.Context, attribute Attribute) error
	Delete(ctx context.Context, attribute Attribute) error

	Query(ctx context.Context, filter QueryFilter, orderBy order.By, pageNumber int, rowsPerPage int) ([]Attribute, error)
	Count(ctx context.Context, filter QueryFilter) (int, error)

	QueryByID(ctx context.Context, attributeID uuid.UUID) (Attribute, error)
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
func (c *Core) Create(ctx context.Context, attribute NewAttribute) (Attribute, error) {

	now := time.Now()

	att := Attribute{
		ID:          uuid.New(),
		Name:        attribute.Name,
		Type:        attribute.Type,
		Instance:    attribute.Instance,
		DateCreated: now,
		DateUpdated: now,
	}

	if err := c.storer.Create(ctx, att); err != nil {
		return Attribute{}, fmt.Errorf("create: %w", err)
	}

	return att, nil
}

// Delete removes a user from the database.
func (c *Core) Delete(ctx context.Context, attribute Attribute) error {
	if err := c.storer.Delete(ctx, attribute); err != nil {
		return fmt.Errorf("delete: %w", err)
	}

	return nil
}

// Query retrieves a list of existing gas from the database.
func (c *Core) Query(ctx context.Context, filter QueryFilter, orderBy order.By, pageNumber int, rowsPerPage int) ([]Attribute, error) {
	att, err := c.storer.Query(ctx, filter, orderBy, pageNumber, rowsPerPage)
	if err != nil {

		fmt.Printf("=============: %v", att)
		fmt.Printf("=============: %v", err)

		return nil, fmt.Errorf("query: %w", err)
	}

	return att, nil
}

// QueryByID finds the user by the specified ID.
func (c *Core) QueryByID(ctx context.Context, attributeID uuid.UUID) (Attribute, error) {
	att, err := c.storer.QueryByID(ctx, attributeID)
	if err != nil {
		return Attribute{}, fmt.Errorf("query: attributeID[%s]: %w", attributeID, err)
	}

	return att, nil
}

// Count returns the total number of cos in the store.
func (c *Core) Count(ctx context.Context, filter QueryFilter) (int, error) {
	return c.storer.Count(ctx, filter)
}
