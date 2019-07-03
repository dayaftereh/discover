package universe

import (
	"sync"

	"github.com/dayaftereh/discover/server/game/data"
	"github.com/dayaftereh/discover/server/game/universe/starsystem"
)

type Universe struct {
	lock        sync.RWMutex
	starSystems map[int64]*starsystem.StarSystem
}

func NewUniverse() *Universe {
	universe := &Universe{
		starSystems: make(map[int64]*starsystem.StarSystem),
	}

	return universe
}

func (universe *Universe) LoadStarSystems(starSystems []*data.StarSystem) {
	universe.lock.Lock()
	defer universe.lock.Unlock()

	for _, starSystem := range starSystems {
		universe.starSystems[starSystem.ID] = starsystem.NewStarSystem(starSystem.ID)
	}
}

func (universe *Universe) GetStarSystems() []*data.StarSystem {
	starSystems := make([]*data.StarSystem, 0)

	universe.lock.RLock()
	defer universe.lock.RUnlock()

	for _, starSystem := range universe.starSystems {
		starSystems = append(starSystems, &data.StarSystem{
			ID: starSystem.ID,
		})
	}

	return starSystems
}

func (universe *Universe) GetStarSystem(id int64) *starsystem.StarSystem {
	universe.lock.RLock()
	defer universe.lock.RUnlock()

	startSystem, ok := universe.starSystems[id]
	if !ok {
		return nil
	}
	return startSystem
}

func (universe *Universe) Shutdown() {
	universe.lock.RLock()
	defer universe.lock.RUnlock()

	for _, starSystem := range universe.starSystems {
		starSystem.Shutdown()
	}
}
