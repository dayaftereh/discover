package starsystem

import (
	"github.com/dayaftereh/discover/server/game/data"
	"github.com/dayaftereh/discover/server/game/universe/starsystem/objects"
	object "github.com/dayaftereh/discover/server/game/universe/starsystem/objects"
)

func (starSystem *StarSystem) loadData(starSystemData *data.StarSystem) {
	// get the id for the sun
	id := starSystem.nextID()
	// create the sun
	starSystem.sun = objects.NewSun(id)
	// load the sun
	starSystem.sun.Load(starSystemData.Sun)
	// add the sun to the world
	starSystem.world.AddObject(starSystem.sun)

	// load all planets
	//for _, planetData := range starSystemData.Planets {
	// load the planet
	//starSystem.loadAndAddPlanet(planetData)
	//}
}

func (starSystem *StarSystem) loadAndAddPlanet(planetData *data.Planet) {
	// get next id for planet
	id := starSystem.nextID()
	// create a new planet
	planet := object.NewPlanet(id)

	// load the planet from data and sun
	planet.Load(starSystem.sun, planetData)

	// add the planet to the world
	starSystem.world.AddObject(planet)
}
