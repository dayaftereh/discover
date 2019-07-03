package common

import (
	"context"
	"net/http"

	"github.com/dayaftereh/discover/server/api"
)

func (common *commonRouter) logout(ctx context.Context, response http.ResponseWriter, request *http.Request, variables map[string]string) error {
	// get the session id
	sessionID, err := api.SessionIdFromContext(ctx)
	if err != nil {
		return err
	}

	// drop the player from game
	common.backend.DropPlayerSession(sessionID)

	return api.SuccessEmpty(response)
}
