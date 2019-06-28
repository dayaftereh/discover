package common

import (
	"context"
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

	// get the player
	player, err := common.backend.GetPlayer(sessionID)
	if err != nil {
		return err
	}
	// check if name exists
	if login.Name == nil {
		return errors.Errorf("fail to login user, because name is missing")
	}

	// set the name for the player
	player.SetName(login.Name)

	return api.SuccessEmpty(response)
}
