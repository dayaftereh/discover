package game

import (
	"github.com/dayaftereh/discover/server/game/data"
)

func (game *Game) Shutdown(directory string) error {
	// shutdown the universe
	game.universe.Shutdown()

	// write the game data
	err := game.writeGameData(directory)
	return err
}

func (game *Game) writeGameData(directory string) error {
	// create a new game data
	gameData := &data.Game{}

	// write the player to game data
	game.playerManager.WritePlayersToData(gameData)

	// write the universe to game data
	game.universe.WriteUniverseToData(gameData)

	// write the game data to file-system
	err := data.Write(directory, gameData)

	return err
}
