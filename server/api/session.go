package api

import (
	"context"

	"github.com/pkg/errors"
)

type Session struct{}

func ForkSessionContext(ctx context.Context, id string) context.Context {
	// add the session id to the context
	sessionContext := context.WithValue(ctx, Session{}, id)

	return sessionContext
}

func SessionIdFromContext(ctx context.Context) (string, error) {
	if ctx == nil {
		return "", errors.Errorf("unable to extract session-id from context, because context is nil")
	}

	val := ctx.Value(Session{})
	if val == nil {
		return "", errors.Errorf("unable to extract session-id from context, because value is nil")
	}

	return val.(string), nil
}
