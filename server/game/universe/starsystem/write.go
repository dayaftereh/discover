package starsystem

import (
	"github.com/dayaftereh/discover/server/game/data"
	"github.com/dayaftereh/discover/server/game/engine/object"
)

func (starSystem *StarSystem) WriteData() *data.StarSystem {
	// lock the star system
	starSystem.lock.Lock()
	defer starSystem.lock.Unlock()

	// list for the planet data
	planetsData := make([]*data.Planet, 0)

	// get all planets from the world
	planets := starSystem.world.GetGameObjectsByType(object.PlanetObject)

	for _, planet := range planets {
		// get the color
		color := planet.Color()
		// get the rigidbody
		rigidbody := planet.RigidBody()
		// get the ration as euler
		rotation := rigidbody.Rotation.ToEuler()

		// add a new planet
		planetsData = append(planetsData, &data.Planet{
			Color:    &color,
			Rotation: rotation,
			Position: rigidbody.Position,
		})
	}

	return &data.StarSystem{
		ID:      starSystem.ID,
		Name:    starSystem.Name,
		Planets: planetsData,
	}
}
