package starsystem

import (
	"github.com/dayaftereh/discover/server/game/persistence/types"
	"github.com/dayaftereh/discover/server/game/universe/starsystem/objects"
)

func (starSystem *StarSystem) Init() error {
	// set the starSystem running
	starSystem.running.Set(true)

	// lock the star system
	starSystem.lock.Lock()
	defer starSystem.lock.Unlock()

	// load the star system data
	starSystemData, err := starSystem.persistence.LoadStarSystem(starSystem.Name)
	if err != nil {
		return err
	}

	starSystem.Data = starSystemData

	// load the sun
	err = starSystem.loadAndAddSun()
	if err != nil {
		return err
	}

	// load all planets
	for _, planet := range starSystem.Data.Planets {
		// load and add the planet
		err = starSystem.loadAndAddPlanet(planet)
		if err != nil {
			return err
		}
	}

	// star the star system loop
	go starSystem.loop()

	return nil
}

func (starSystem *StarSystem) loadAndAddSun() error {
	// get the id for the sun
	id := starSystem.nextID()
	// create the sun
	starSystem.sun = objects.NewSun(id, starSystem.Data.Sun)
	// initialize the sun
	err := starSystem.sun.Init()
	if err != nil {
		return err
	}

	// add the sun to the world
	starSystem.world.AddObject(starSystem.sun)

	return nil
}

func (starSystem *StarSystem) loadAndAddPlanet(planetData *types.Planet) error {
	// get the id for the planet
	id := starSystem.nextID()
	// create the planet
	planet := objects.NewPlanet(id, planetData, starSystem.sun)
	// initialize the planet
	err := planet.Init()
	if err != nil {
		return err
	}

	// add the planet to the world
	starSystem.world.AddObject(planet)

	return nil
}
