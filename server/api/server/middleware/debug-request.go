package middleware

import (
	"context"
	"net/http"

	"github.com/dayaftereh/discover/server/api"
)

type DebugRequestMiddleware struct{}

func NewDebugRequestMiddleware() DebugRequestMiddleware {
	return DebugRequestMiddleware{}
}

func (middleware DebugRequestMiddleware) WrapHandler(handler api.Function) api.Function {
	return func(ctx context.Context, response http.ResponseWriter, request *http.Request, variables map[string]string) error {
		return handler(ctx, response, request, variables)
	}
}
