package starsystem

import (
	"github.com/dayaftereh/discover/server/game/data"
	"github.com/dayaftereh/discover/server/game/engine/object"
)

func (starSystem *StarSystem) loadData(starSystemData *data.StarSystem) {

	if starSystemData.Sun != nil {
		starSystem.sunMass = starSystemData.Sun.Mass
		starSystem.sunColor = starSystemData.Sun.Color
	} else {
		starSystem.sunMass = 1e8
		starSystem.sunColor = 1234
	}

	// load all planets
	for _, planetData := range starSystemData.Planets {
		// get the id of the object
		id := starSystem.nextID()

		// create the planet
		planet := object.NewPlanet(id)

		// load the planet from data
		planet.Load(starSystem.sunMass, planetData)

		// add the planet to the world
		starSystem.world.AddGameObject(planet)
	}
}
