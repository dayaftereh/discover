package api

import (
	"context"
	"net/http"
)

// Function is an adapter to allow the use of ordinary functions as API endpoints.
type Function func(ctx context.Context, response http.ResponseWriter, request *http.Request, variables map[string]string) error
