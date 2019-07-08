package common

import (
	"context"
	"net/http"

	"github.com/dayaftereh/discover/server/api"
	types "github.com/dayaftereh/discover/server/api/types/common"
)

func (common *commonRouter) status(ctx context.Context, response http.ResponseWriter, request *http.Request, variables map[string]string) error {
	// get the session id
	sessionID, err := api.SessionIdFromContext(ctx)
	if err != nil {
		return err
	}

	// create the status
	status := types.Status{
		Authenticated: false,
	}

	// check if the player exists
	player := common.backend.GetPlayerSession(sessionID)
	if player != nil {
		// set authenticated to true
		status.Authenticated = true
		status.ID = player.ID
		status.Name = &player.Name
		status.StarSystem = player.StarSystem
	}

	// write the status as json response
	err = api.WriteJSON(response, http.StatusOK, &status)
	return err
}
