package game

import (
	"github.com/dayaftereh/discover/server/game/data"
)

func (game *Game) Init(directory string) error {
	// load the game data
	gameData, err := data.Load(directory)
	if err != nil {
		return err
	}

	// load the players
	game.playerManager.LoadPlayersFromData(gameData)

	// load the universe
	game.universe.LoadUniverseFromData(gameData)

	return err
}
