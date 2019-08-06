package universe

import (
	"log"
	"sync"

	"github.com/dayaftereh/discover/server/game/persistence"
	"github.com/dayaftereh/discover/server/game/persistence/types"

	"github.com/dayaftereh/discover/server/game/universe/generator"
	"github.com/dayaftereh/discover/server/game/universe/starsystem"
	"github.com/pkg/errors"
)

type Universe struct {
	lock              sync.RWMutex
	initialStarSystem string
	starSystems       map[string]*starsystem.StarSystem
	persistence       *persistence.PersistenceManager
	generator         *generator.Generator
}

func NewUniverse(persistenceManager *persistence.PersistenceManager) *Universe {
	universe := &Universe{
		persistence: persistenceManager,
		generator:   generator.NewGenerator(persistenceManager),
		starSystems: make(map[string]*starsystem.StarSystem),
	}

	return universe
}

func (universe *Universe) Init() error {
	universe.lock.Lock()
	defer universe.lock.Unlock()

	// load the universeData
	universeData, err := universe.persistence.LoadUniverse()
	if err != nil {
		return err
	}

	// load all star systems
	for name, _ := range universeData.StarSystems {
		// create the star system
		universe.starSystems[name] = starsystem.NewStarSystem(name, universe.persistence)
		// load the star system
		err = universe.starSystems[name].Init()
		if err != nil {
			return err
		}
	}

	if universeData.InitialStarSystem != nil {
		universe.initialStarSystem = *universeData.InitialStarSystem
	} else {
		// generate the initial star system
		err = universe.generateInitialStarSystem()
		if err != nil {
			return err
		}
	}
	// check if the initial star system exists
	_, initialStarSystemExists := universe.starSystems[universe.initialStarSystem]
	if !initialStarSystemExists {
		// generate a initialStarSystem
		universe.generateNewStarSystem(universe.initialStarSystem)
	}

	for len(universe.starSystems) < 15 {
		// find a new random name
		starSystemName, err := universe.findRandomStarSystemName()
		if err != nil {
			return err
		}
		// generate the new star system
		universe.generateNewStarSystem(starSystemName)
	}

	return nil
}

func (universe *Universe) generateInitialStarSystem() error {
	name, err := universe.findRandomStarSystemName()
	if err != nil {
		return err
	}

	log.Printf("generating initial star-system [ %s ]", name)

	// generate and create a new InitialStarSystem
	_, err = universe.generateNewStarSystem(name)
	if err != nil {
		return err
	}

	// set the initialStarSystem
	universe.initialStarSystem = name

	// storage the new universe
	err = universe.writeUniverse()
	return err
}

func (universe *Universe) generateNewStarSystem(name string) (*types.StarSystem, error) {
	log.Printf("generating star-system [ %s ]", name)

	// generate a new star system
	starSystem, err := universe.generator.Generate(name)
	// create the star system
	universe.starSystems[name] = starsystem.NewStarSystem(name, universe.persistence)
	// load the star system
	err = universe.starSystems[name].Init()
	if err != nil {
		return nil, err
	}

	// storage the new universe
	err = universe.writeUniverse()
	if err != nil {
		return nil, err
	}

	for _, p := range starSystem.Planets {
		if p.Atmosphere != nil && len(p.Atmosphere) > 0 {
			log.Println(p.Name)
		}
	}

	return starSystem, err
}

func (universe *Universe) findRandomStarSystemName() (string, error) {
	for {
		// generate a random name
		name := universe.generator.RandomStarSystemName()

		// check if the name not already used
		_, ok := universe.starSystems[name]
		if !ok {
			return name, nil
		}
	}
}

func (universe *Universe) GetInitialStarSystem() (*starsystem.StarSystem, error) {
	universe.lock.RLock()
	defer universe.lock.RUnlock()

	startSystem, ok := universe.starSystems[universe.initialStarSystem]
	if !ok {
		return nil, errors.Errorf("unable to find initial star system with id [ %d ]", universe.initialStarSystem)
	}
	return startSystem, nil
}

func (universe *Universe) GetStarSystem(name string) *starsystem.StarSystem {
	universe.lock.RLock()
	defer universe.lock.RUnlock()

	startSystem, ok := universe.starSystems[name]
	if !ok {
		return nil
	}
	return startSystem
}

func (universe *Universe) writeUniverse() error {
	starSystems := make(map[string][]string)
	// convert the star systems
	for name := range universe.starSystems {
		starSystems[name] = make([]string, 0)
	}

	// create data for the universe
	data := &types.Universe{
		StarSystems:       starSystems,
		InitialStarSystem: &universe.initialStarSystem,
	}

	// write the universe
	err := universe.persistence.WriteUniverse(data)
	return err
}

func (universe *Universe) StarSystemNames() []string {
	universe.lock.RLock()
	defer universe.lock.RUnlock()

	names := make([]string, 0)
	for name := range universe.starSystems {
		names = append(names, name)
	}
	return names
}

func (universe *Universe) Shutdown() {
	log.Println("shutdown universe...")

	universe.lock.RLock()
	defer universe.lock.RUnlock()

	for _, starSystem := range universe.starSystems {
		starSystem.Shutdown()
	}
}
