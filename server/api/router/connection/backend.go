package connection

import "github.com/dayaftereh/discover/server/game/player"

type Backend interface {
	GetPlayerSession(id string) *player.Player
}
