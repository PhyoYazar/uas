package vco

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type Storer interface {
	QueryByID(ctx context.Context, subjectID uuid.UUID) (VCo, error)
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
	ErrNotFound = errors.New("co not found")
)

func (c *Core) QueryByID(ctx context.Context, coID uuid.UUID) (VCo, error) {
	sub, err := c.storer.QueryByID(ctx, coID)
	if err != nil {
		return VCo{}, fmt.Errorf("query: subjectID[%s]: %w", coID, err)
	}

	return sub, nil
}
