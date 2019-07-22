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

	for _, gameObject := range planets {
		// cast to a planet
		planet := gameObject.(*object.Planet)

		// add a new planet
		planetsData = append(planetsData, planet.Write())
	}

	return &data.StarSystem{
		ID:      starSystem.ID,
		Name:    starSystem.Name,
		Planets: planetsData,
		Sun: &data.Sun{
			Color: starSystem.sunColor,
			Mass:  starSystem.sunMass,
		},
	}
}
