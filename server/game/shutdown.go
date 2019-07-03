package game

import (
	"github.com/dayaftereh/discover/server/game/data"
)

func (game *Game) Shutdown(directory string) error {
	// shutdown the universe
	game.universe.Shutdown()

	err := game.writeGameData(directory)
	return err
}

func (game *Game) writeGameData(directory string) error {
	// create a new game data
	gameData := data.NewGame()

	// get the players for the game data
	err := game.writePlayers(gameData)
	if err != nil {
		return err
	}

	// get the star systems for the game data
	err = game.writeStarSystems(gameData)
	if err != nil {
		return err
	}

	// write the game data
	err = data.Write(directory, gameData)

	return err
}

func (game *Game) writePlayers(gameData *data.Game) error {
	// get the players
	players := game.playerManager.GetPlayers()

	for _, player := range players {
		gameData.Players[player.Name] = player
	}

	return nil
}

func (game *Game) writeStarSystems(gameData *data.Game) error {
	// get the starSystems
	starSystems := game.universe.GetStarSystems()

	for _, starSystem := range starSystems {
		gameData.StarSystems[starSystem.ID] = starSystem
	}

	return nil
}
