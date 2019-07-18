package game

import "github.com/dayaftereh/discover/server/game/player"

type Backend interface {
	Ready(player *player.Player) error
	GetPlayerSession(id string) *player.Player
}
