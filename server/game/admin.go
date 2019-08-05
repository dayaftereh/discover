package game

import "github.com/dayaftereh/discover/server/game/universe/starsystem"

func (game *Game) GetStarSystems() []string {
	return game.universe.StarSystemNames()
}

func (game *Game) GetStarSystem(name string) *starsystem.StarSystem {
	return game.universe.GetStarSystem(name)
}
