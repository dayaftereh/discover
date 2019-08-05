package admin

import (
	"github.com/dayaftereh/discover/server/game/player"
	"github.com/dayaftereh/discover/server/game/universe/starsystem"
)

type Backend interface {
	GetPlayerSession(id string) *player.Player
	GetStarSystems() []string
	GetStarSystem(name string) *starsystem.StarSystem
}
