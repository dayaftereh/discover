package connection

import (
	"github.com/dayaftereh/discover/server/game/player"
)

type Backend interface {
	GetPlayer(id string) (*player.Player, error)
}
