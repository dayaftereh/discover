package universe

import (
	"log"
	"sync"

	"github.com/dayaftereh/discover/server/game/data"
	"github.com/dayaftereh/discover/server/game/universe/starsystem"
	"github.com/pkg/errors"
)

type Universe struct {
	lock              sync.RWMutex
	initialStarSystem *int64
	starSystems       map[int64]*starsystem.StarSystem
}

func NewUniverse() *Universe {
	universe := &Universe{
		starSystems: make(map[int64]*starsystem.StarSystem),
	}

	return universe
}

func (universe *Universe) LoadUniverseFromData(gameData *data.Game) {
	universe.lock.Lock()
	defer universe.lock.Unlock()

	// set the initialStarSystem
	universe.initialStarSystem = gameData.Universe.InitialStarSystem

	// create the starSystems
	for _, starSystemData := range gameData.Universe.StarSystems {
		// create star system from data
		universe.starSystems[starSystemData.ID] = starsystem.NewStarSystem(starSystemData)
	}
}

func (universe *Universe) WriteUniverseToData(gameData *data.Game) {
	universe.lock.RLock()
	defer universe.lock.RUnlock()

	// create the universe data
	gameData.Universe = &data.Universe{
		InitialStarSystem: universe.initialStarSystem,
		StarSystems:       make(map[int64]*data.StarSystem),
	}

	// write the star systems to data
	for id, gameStarSystem := range universe.starSystems {
		// write the data for the star system
		gameData.Universe.StarSystems[id] = gameStarSystem.WriteData()
	}
}

func (universe *Universe) GetInitialStarSystem() (*starsystem.StarSystem, error) {
	universe.lock.RLock()
	defer universe.lock.RUnlock()

	if universe.initialStarSystem == nil {
		return nil, errors.Errorf("missing initial star system id")
	}

	startSystem, ok := universe.starSystems[*universe.initialStarSystem]
	if !ok {
		return nil, errors.Errorf("unable to find initial star system with id [ %d ]", universe.initialStarSystem)
	}
	return startSystem, nil
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
	log.Println("shutdown universe...")

	universe.lock.RLock()
	defer universe.lock.RUnlock()

	for _, starSystem := range universe.starSystems {
		starSystem.Shutdown()
	}
}
