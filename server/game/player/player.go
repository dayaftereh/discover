package player

import (
	"sync"

	"github.com/dayaftereh/discover/server/game/persistence/types"
)

type Player struct {
	ID    string
	Name  string
	Admin bool
	// game
	StarSystem *string
	// private
	look        sync.Mutex
	connections map[string]Connection
}

func NewPlayer(id string, data *types.Player) *Player {
	return &Player{
		ID:          id,
		Admin:       data.Admin,
		Name:        data.Name,
		StarSystem:  data.StarSystem,
		connections: make(map[string]Connection),
	}
}
