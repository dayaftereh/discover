package starsystem

import (
	"github.com/dayaftereh/discover/server/game/data"
	"github.com/dayaftereh/discover/server/game/engine/object"
	"github.com/dayaftereh/discover/server/mathf"
)

func (starSystem *StarSystem) loadData(starSystemData *data.StarSystem) {

	// load all planets
	for _, planetData := range starSystemData.Planets {
		// get the id of the object
		id := starSystem.nextID()

		// load the color for the planet
		color := int64(0xff82a1)
		if planetData.Color != nil {
			color = *planetData.Color
		}

		// create the planet
		planet := object.NewPlanet(id, color)

		// get the rigidbody of the planet
		rigidbody := planet.RigidBody()

		// set location if given
		if planetData.Position != nil {
			rigidbody.Position = planetData.Position
		}

		// set rotation if given
		if planetData.Rotation != nil {
			rigidbody.Rotation = mathf.QuaternionFromEuler(planetData.Rotation)
		}

		// add the planet to the world
		starSystem.world.AddGameObject(planet)
	}
}
