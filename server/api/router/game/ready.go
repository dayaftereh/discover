package game

import (
	"context"
	"log"
	"net/http"

	"github.com/dayaftereh/discover/server/api"
)

func (game *commonGame) ready(ctx context.Context, response http.ResponseWriter, request *http.Request, variables map[string]string) error {
	// get the sessionId
	sessionID, err := api.SessionIdFromContext(ctx)
	if err != nil {
		return err
	}

	// get the player session
	player := game.backend.GetPlayerSession(sessionID)

	// response player not found
	if player == nil {
		return api.NotFound(response)
	}

	// log about player ready
	log.Printf("player [ %s ] notified about ready\n", player.Name)

	// mark the player as ready
	game.backend.Ready(player)

	// response ok
	return api.SuccessEmpty(response)
}
