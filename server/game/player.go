package game

import (
	"github.com/dayaftereh/discover/server/game/player"
	"github.com/pkg/errors"
)

func (game *Game) HasPlayer(id string) bool {
	// get or create the player
	player := game.playerManager.Get(id)
	return player != nil
}

// GetPlayer or create for given id
func (game *Game) GetPlayer(id string) (*player.Player, error) {
	// get or create the player
	player := game.playerManager.GetOrCreate(id)

	// check if player found
	if player == nil {
		return nil, errors.Errorf("unable to find player for id [ %s ]", id)
	}

	return player, nil
}

// DropPlayer for given id from Game
func (game *Game) DropPlayer(id string) error {
	// remove the player
	player := game.playerManager.Remove(id)
	// check for player found
	if player == nil {
		return errors.Errorf("unable to drop player, because current player with id [ %s ] not found", id)
	}

	return nil
}
