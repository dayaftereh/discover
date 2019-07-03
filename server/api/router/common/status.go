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

	// check if the player exists
	player := common.backend.GetPlayerSession(sessionID)
	if player == nil {
		// response not found
		return api.NotFound(response)
	}

	// create the status
	status := types.Status{
		Id:         player.ID,
		Name:       &player.Name,
		StarSystem: player.StarSystem,
	}

	// write the status as json response
	err = api.WriteJSON(response, http.StatusOK, &status)
	return err
}
