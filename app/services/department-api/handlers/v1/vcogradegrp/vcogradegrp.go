package vcogradegrp

import (
	"context"
	"fmt"
	"net/http"

	"github.com/PhyoYazar/uas/business/core/vcograde"
	"github.com/PhyoYazar/uas/business/web/v1/paging"
	"github.com/PhyoYazar/uas/foundation/web"
)

// Handlers manages the set of user endpoints.
type Handlers struct {
	vcograde *vcograde.Core
}

// New constructs a handlers for route access.
func New(vcograde *vcograde.Core) *Handlers {
	return &Handlers{
		vcograde: vcograde,
	}
}

// Query returns a list of subjects with paging.
func (h *Handlers) Query(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

	page, err := paging.ParseRequest(r)
	if err != nil {
		return err
	}

	filter, err := parseFilter(r)
	if err != nil {
		return err
	}

	orderBy, err := parseOrder(r)
	if err != nil {
		return err
	}

	stds, err := h.vcograde.Query(ctx, filter, orderBy, page.Number, page.RowsPerPage)
	if err != nil {
		return fmt.Errorf("query: %w", err)
	}

	return web.Respond(ctx, w, paging.NewResponse(stds, 0, page.Number, page.RowsPerPage), http.StatusOK)
}
