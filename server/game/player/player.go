package player

import (
	"sync"

	"github.com/dayaftereh/discover/server/game/data"
)

type Player struct {
	ID   string
	Name string
	// game
	StarSystem *int64
	// private
	look        sync.Mutex
	connections map[string]Connection
}

func NewPlayer(id string, data *data.Player) *Player {
	return &Player{
		ID:          id,
		Name:        data.Name,
		StarSystem:  data.StarSystem,
		connections: make(map[string]Connection),
	}
}
