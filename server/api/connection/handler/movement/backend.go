package movement

import (
	"github.com/dayaftereh/discover/server/game/player"
)

type Backend interface {
	Movement(player *player.Player, x float64, y float64, z float64)
}
