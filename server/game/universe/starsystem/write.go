package starsystem

import (
	"github.com/dayaftereh/discover/server/game/data"
	"github.com/dayaftereh/discover/server/game/universe/starsystem/objects"
)

func (starSystem *StarSystem) WriteData() *data.StarSystem {
	// lock the star system
	starSystem.lock.Lock()
	defer starSystem.lock.Unlock()

	// write the planets
	planets := starSystem.writePlanets()

	// convert the sun to data
	sun := starSystem.sun.Write()

	// create the sun system data
	return &data.StarSystem{
		ID:      starSystem.ID,
		Name:    starSystem.Name,
		Sun:     sun,
		Planets: planets,
	}
}

func (starSystem *StarSystem) writePlanets() []*data.Planet {
	// list for the planet data
	planetsData := make([]*data.Planet, 0)

	// get all planets from the world
	planets := starSystem.findGameObjectsByType(objects.GameObjectPlanet)

	for _, gameObject := range planets {
		// cast to a planet
		planet := gameObject.(*objects.Planet)

		// add a new planet
		planetsData = append(planetsData, planet.Write())
	}

	return planetsData
}
