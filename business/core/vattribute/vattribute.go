package vattribute

import (
	"context"
	"errors"
	"fmt"

	"github.com/PhyoYazar/uas/business/data/order"
	"github.com/google/uuid"
)

type Storer interface {
	Query(ctx context.Context, filter QueryFilter, orderBy order.By, pageNumber int, rowsPerPage int, subjectID uuid.UUID) ([]VAttribute, error)

	// Count(ctx context.Context, filter QueryFilter) (int, error)
}

type Core struct {
	storer Storer
}

func NewCore(storer Storer) *Core {
	return &Core{
		storer: storer,
	}
}

// Set of error variables for CRUD operations.
var (
	ErrNotFound = errors.New("attributes not found")
)

// Query retrieves a list of existing subjects from the database.
func (c *Core) Query(ctx context.Context, filter QueryFilter, orderBy order.By, pageNumber int, rowsPerPage int, subjectID uuid.UUID) ([]VAttribute, error) {
	atts, err := c.storer.Query(ctx, filter, orderBy, pageNumber, rowsPerPage, subjectID)
	if err != nil {

		fmt.Printf("=============: %v", atts)
		fmt.Printf("=============: %v", err)

		return nil, fmt.Errorf("query: %w", err)
	}

	return atts, nil
}

// // Count returns the total number of subjects in the store.
// func (c *Core) Count(ctx context.Context, filter QueryFilter) (int, error) {
// 	return c.storer.Count(ctx, filter)
// }