package vstudentmark

import (
	"context"
	"errors"
	"fmt"

	"github.com/PhyoYazar/uas/business/data/order"
	"github.com/google/uuid"
)

type Storer interface {
	Query(ctx context.Context, filter QueryFilter, orderBy order.By, pageNumber int, rowsPerPage int, subjectID uuid.UUID) ([]VStudentMark, error)

	// RemoveStudentMarks(ctx context.Context, rs VRemoveStudent) error

	Count(ctx context.Context, filter QueryFilter) (int, error)
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
func (c *Core) Query(ctx context.Context, filter QueryFilter, orderBy order.By, pageNumber int, rowsPerPage int, subjectID uuid.UUID) ([]VStudentMark, error) {
	atts, err := c.storer.Query(ctx, filter, orderBy, pageNumber, rowsPerPage, subjectID)
	if err != nil {

		fmt.Printf("=============: %v", atts)
		fmt.Printf("=============: %v", err)

		return nil, fmt.Errorf("query: %w", err)
	}

	return atts, nil
}

// func (c *Core) RemoveStudentMarks(ctx context.Context, ra VRemoveStudent) error {
// 	if err := c.storer.RemoveStudentMarks(ctx, ra); err != nil {
// 		return fmt.Errorf("remove marks by subjectID and attributeID: %w", err)
// 	}

// 	return nil
// }

// Count returns the total number of subjects in the store.
func (c *Core) Count(ctx context.Context, filter QueryFilter) (int, error) {
	return c.storer.Count(ctx, filter)
}
