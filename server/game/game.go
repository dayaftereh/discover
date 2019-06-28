package game

import "github.com/dayaftereh/discover/server/game/player"

type Game struct {
	// player
	playerManager *player.Manager
}

func NewGame() *Game {
	return &Game{
		playerManager: player.NewPlayerManager(),
	}
}
