package connection

import (
	"context"
	"net/http"

	"github.com/dayaftereh/discover/server/utils"

	"github.com/dayaftereh/discover/server/api"
	wsconn "github.com/dayaftereh/discover/server/api/connection"
	"github.com/pkg/errors"
)

func (connection *connectionRouter) websocket(ctx context.Context, response http.ResponseWriter, request *http.Request, variables map[string]string) error {
	// get the sessionId
	sessionID, err := api.SessionIdFromContext(ctx)
	if err != nil {
		return err
	}

	// get player for session
	player := connection.backend.GetPlayerSession(sessionID)

	// check if player login
	if player == nil {
		// do not allow befor login
		return api.Forbidden(response)
	}

	// upgrade the connection to a websocket
	conn, err := connection.upgrader.Upgrade(response, request, nil)
	if err != nil {
		return errors.Wrapf(err, "fail to upgrade incoming reguest to a stable websocket")
	}

	// genrade a random connection id
	connectionID, err := utils.RandString(128)
	if err != nil {
		return errors.Wrapf(err, "unable to generated a random connection id")
	}

	// create a new connection
	websocketConn := wsconn.NewConnection(connectionID, player, conn)
	// notify dispatcher about new websocket connection
	connection.dispatcher.EmitOpen(websocketConn)

	// add the connection to the player
	player.AddConnection(websocketConn)

	return nil
}
