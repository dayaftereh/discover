package starsystem

import (
	"github.com/dayaftereh/discover/server/game/data"
	"github.com/dayaftereh/discover/server/game/engine/object"
	"github.com/dayaftereh/discover/server/mathf"
	"github.com/dayaftereh/discover/server/mathf/orbit"
)

func (starSystem *StarSystem) loadData(starSystemData *data.StarSystem) {
	// get the id for the sun
	id := starSystem.nextID()
	// create the sun
	starSystem.sun = object.NewSun(id)
	// load the sun
	starSystem.sun.Load(starSystemData.Sun)
	// add the sun to the world
	starSystem.world.AddGameObject(starSystem.sun)

	// load all planets
	for _, planetData := range starSystemData.Planets {
		// load the planet
		starSystem.loadAndAddPlanet(planetData)
	}
}

func (starSystem *StarSystem) loadAndAddPlanet(planetData *data.Planet) {
	// get next id for planet
	id := starSystem.nextID()
	// load the orbit for the planet
	orbit := starSystem.loadOrbitForPlanet(planetData.Orbit)

	// create a new planet
	planet := object.NewPlanet(id)

	// load the planet from data
	planet.Load(orbit, planetData)

	// add the planet to the world
	starSystem.world.AddGameObject(planet)
}

func (starSystem *StarSystem) loadOrbitForPlanet(orbitData *data.Orbit) *orbit.Orbit {
	// get the RigidBody of the sun
	sunRigidBody := starSystem.sun.RigidBody()
	// get th radius of the sun
	centralBodyRadius := starSystem.sun.Radius()
	// calucate mu
	mu := mathf.GravitationalConstant * sunRigidBody.Mass

	// create the orbit
	orbit := orbit.OrbitFromParams(&orbit.OrbitParameter{
		MU:                  &mu,
		CentralBodyRadius:   &centralBodyRadius,
		Apogee:              &orbitData.Apogee,
		Perigee:             &orbitData.Perigee,
		Eccentricity:        &orbitData.Eccentricity,
		Inclination:         &orbitData.Inclination,
		RightAscension:      &orbitData.RightAscension,
		ArgumentOfPeriapsis: &orbitData.ArgumentOfPeriapsis,
	})

	return orbit
}
