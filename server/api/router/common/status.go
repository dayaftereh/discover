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
	found := common.backend.HasPlayer(sessionID)
	if !found {
		// response not found
		return api.NotFound(response)
	}

	// get the Player
	player, err := common.backend.GetPlayer(sessionID)
	if err != nil {
		return err
	}

	// create the status
	status := types.Status{
		Id:   player.Id,
		Name: player.Name,
	}

	// write the status as json response
	err = api.WriteJSON(response, http.StatusOK, &status)
	return err
}
