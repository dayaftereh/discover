package movement

import (
	"github.com/dayaftereh/discover/server/game/player"
	"github.com/dayaftereh/discover/server/mathf"
)

type Backend interface {
	Movement(player *player.Player, move *mathf.Vec3, rotation *mathf.Vec3)
}
