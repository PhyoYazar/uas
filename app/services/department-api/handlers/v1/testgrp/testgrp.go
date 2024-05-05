package testgrp

import (
	"context"
	"errors"
	"math/rand"
	"net/http"

	v1 "github.com/PhyoYazar/uas/business/web/v1"
	"github.com/PhyoYazar/uas/foundation/web"
)

func Test(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	if n := rand.Intn(100); n%2 == 0 {
		return v1.NewRequestError(errors.New("TRUSTED ERROR"), http.StatusBadRequest)
		// panic("OHH NOO PANIC")
	}

	// 1. Validate the data
	// 2. Call into the business layer

	status := struct {
		Status string `json:"status"`
	}{
		Status: "Ok",
	}

	return web.Respond(ctx, w, status, http.StatusOK)
}
