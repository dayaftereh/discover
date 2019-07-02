package universe

import (
	"github.com/dayaftereh/discover/server/game/player"
	"github.com/dayaftereh/discover/server/game/universe/starsystem"
	"github.com/dayaftereh/discover/server/utils"
)

type Universe struct {
	starSystem  int64
	counter     *utils.IDCounter
	players     map[string]int64
	starSystems map[int64]*starsystem.StarSystem
}

func NewUniverse() *Universe {
	universe := &Universe{
		starSystem:  0,
		counter:     utils.NewIDCounter(),
		players:     make(map[string]int64),
		starSystems: make(map[int64]*starsystem.StarSystem),
	}

	universe.starSystem = universe.GenerateStarSystem()

	return universe
}

func (universe *Universe) DefaultStarSystem() *starsystem.StarSystem {
	return universe.GetStarSystem(universe.starSystem)
}

func (universe *Universe) GenerateStarSystem() int64 {
	id := universe.counter.Next()
	universe.starSystems[id] = starsystem.NewStarSystem(id)
	return id
}

func (universe *Universe) GetStarSystem(id int64) *starsystem.StarSystem {
	startSystem, ok := universe.starSystems[id]
	if !ok {
		return nil
	}
	return startSystem
}

func (universe *Universe) GetPlayerStarSystem(player *player.Player) *starsystem.StarSystem {
	starSystemID, ok := universe.players[player.ID]
	if !ok {
		return nil
	}

	return universe.GetStarSystem(starSystemID)
}

func (universe *Universe) SetPlayerStarSystem(player *player.Player, id int64) {
	universe.players[player.ID] = id
}

func (universe *Universe) DropPlayerStarSystem(player *player.Player) {
	_, ok := universe.players[player.ID]
	if ok {
		delete(universe.players, player.ID)
	}
}
