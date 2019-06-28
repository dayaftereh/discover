package middleware

import (
	"context"
	"net/http"

	"github.com/pkg/errors"

	"github.com/dayaftereh/discover/server/api"
	"github.com/dayaftereh/discover/server/api/session"
)

type SessionMiddleware struct {
	manager *session.Manager
}

func NewSessionMiddleware(manager *session.Manager) SessionMiddleware {
	return SessionMiddleware{
		manager,
	}
}

func (middleware SessionMiddleware) WrapHandler(handler api.Function) api.Function {
	return func(ctx context.Context, response http.ResponseWriter, request *http.Request, variables map[string]string) error {
		// try to get the session for the request
		session, err := middleware.manager.Get(response, request)

		// check for error
		if err != nil {
			return errors.Wrapf(err, "fail to extract session from incoming request")
		}

		// create the session context
		ctx = api.ForkSessionContext(ctx, session.Id)

		return handler(ctx, response, request, variables)
	}
}
