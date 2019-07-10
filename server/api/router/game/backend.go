package game

import "github.com/dayaftereh/discover/server/game/player"

type Backend interface {
	Ready(player *player.Player)
	GetPlayerSession(id string) *player.Player
}
