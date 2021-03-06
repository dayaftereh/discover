package common

import (
	"context"
	"log"
	"net/http"

	"github.com/dayaftereh/discover/server/api"
	types "github.com/dayaftereh/discover/server/api/types/common"
	"github.com/pkg/errors"
)

func (common *commonRouter) login(ctx context.Context, response http.ResponseWriter, request *http.Request, variables map[string]string) error {
	var login types.Login
	// read the received json
	err := api.ReadJSON(request, &login)
	if err != nil {
		return err
	}

	// get the sessionId
	sessionID, err := api.SessionIdFromContext(ctx)
	if err != nil {
		return err
	}

	// check if name exists
	if login.Name == nil {
		return errors.Errorf("fail to login user, because name is missing")
	}

	// get the player
	player, err := common.backend.SessionByName(sessionID, *login.Name)
	if err != nil {
		return err
	}

	// create the status
	status := types.Status{
		Authenticated: true,
		ID:            player.ID,
		Name:          &player.Name,
		StarSystem:    player.StarSystem,
		Admin:         player.Admin,
	}

	// log the login
	log.Printf("user [ %s ] with session [ %s ] successful login\n", player.Name, player.ID)

	// write the status as json response
	err = api.WriteJSON(response, http.StatusOK, &status)

	return err
}
