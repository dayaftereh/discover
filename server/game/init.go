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
	err = game.loadPlayers(gameData)
	if err != nil {
		return err
	}

	// load the star systems
	err = game.loadStarSystems(gameData)

	return err
}

func (game *Game) loadPlayers(gameData *data.Game) error {
	// convert players to array
	players := make([]*data.Player, 0)
	for _, player := range gameData.Players {
		players = append(players, player)
	}

	// load the players
	game.playerManager.LoadPlayers(players)

	return nil
}

func (game *Game) loadStarSystems(gameData *data.Game) error {
	// convert star-systems to array
	starSystems := make([]*data.StarSystem, 0)
	for _, starSystem := range gameData.StarSystems {
		starSystems = append(starSystems, starSystem)
	}

	// load the star systems
	game.universe.LoadStarSystems(starSystems)

	return nil
}
