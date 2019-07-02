package game

import (
	"github.com/dayaftereh/discover/server/game/player"
	"github.com/dayaftereh/discover/server/game/universe"
)

type Game struct {
	// player
	playerManager *player.Manager
	// Universe
	universe *universe.Universe
}

func NewGame() *Game {
	return &Game{
		universe:      universe.NewUniverse(),
		playerManager: player.NewPlayerManager(),
	}
}
