package starsystem

import (
	"github.com/dayaftereh/discover/server/game/data"
	"github.com/dayaftereh/discover/server/game/engine/object"
)

func (starSystem *StarSystem) WriteData() *data.StarSystem {
	// lock the star system
	starSystem.lock.Lock()
	defer starSystem.lock.Unlock()

	// write the planets
	planets := starSystem.writePlanets()

	// get the sun RigidBody
	sunRigidBody := starSystem.sun.RigidBody()

	// create the sun system data
	return &data.StarSystem{
		ID:      starSystem.ID,
		Name:    starSystem.Name,
		Planets: planets,
		Sun: &data.Sun{
			Mass:   sunRigidBody.Mass,
			Color:  starSystem.sun.Color(),
			Radius: starSystem.sun.Radius(),
		},
	}
}

func (starSystem *StarSystem) writePlanets() []*data.Planet {
	// list for the planet data
	planetsData := make([]*data.Planet, 0)

	// get all planets from the world
	planets := starSystem.world.GetGameObjectsByType(object.PlanetObject)

	for _, gameObject := range planets {
		// cast to a planet
		planet := gameObject.(*object.Planet)

		// add a new planet
		planetsData = append(planetsData, planet.Write())
	}

	return planetsData
}
