package vcograde

import (
	"context"
	"errors"
	"fmt"

	"github.com/PhyoYazar/uas/business/data/order"
)

type Storer interface {
	Query(ctx context.Context, filter QueryFilter, orderBy order.By, pageNumber int, rowsPerPage int) ([]VStudentMark, error)
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
	ErrNotFound = errors.New("student mark not found")
)

// Query retrieves a list of existing subjects from the database.
func (c *Core) Query(ctx context.Context, filter QueryFilter, orderBy order.By, pageNumber int, rowsPerPage int) ([]VStudentMark, error) {
	std, err := c.storer.Query(ctx, filter, orderBy, pageNumber, rowsPerPage)
	if err != nil {

		fmt.Printf("=============: %v", std)
		fmt.Printf("=============: %v", err)

		return nil, fmt.Errorf("query: %w", err)
	}

	return std, nil
}
