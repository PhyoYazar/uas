package vsubject

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type Storer interface {
	QueryByID(ctx context.Context, subjectID uuid.UUID) (VSubject, error)
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
	ErrNotFound              = errors.New("sub not found")
	ErrUniqueSubjectYear     = errors.New("name/code ,semester and academic year is not unique")
	ErrAuthenticationFailure = errors.New("authentication failed")
)

func (c *Core) QueryByID(ctx context.Context, subjectID uuid.UUID) (VSubject, error) {
	sub, err := c.storer.QueryByID(ctx, subjectID)
	if err != nil {
		return VSubject{}, fmt.Errorf("query: subjectID[%s]: %w", subjectID, err)
	}

	return sub, nil
}
