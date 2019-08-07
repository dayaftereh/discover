package stargen

import (
	"github.com/dayaftereh/discover/server/game/persistence/types"
	"github.com/dayaftereh/discover/server/utils/container"
)

func GenerateStellarSystem(randomTilt bool, doMoons bool, doGases bool) (*types.Sun, []*types.Planet) {
	// create a new sun for the stellar system
	sun := NewRandomSun()

	// create the accretion processor for the planetesimal
	accretionProcessor := NewAccretionProcessor()
	// calculate the limit of stellar dust based on the mass of the sun
	outerDustLimit := accretionProcessor.StellarDustLimit(sun.Mass)

	// create the planetesimal
	accretionProcessor.DistPlanetaryMasses(
		sun.Mass,
		sun.Luminosity,
		0.0,
		outerDustLimit,
		0.0,
		DustDensityCoefficient,
		true,
	)
	// get the planetesimal
	planets := make([]*types.Planet, 0)
	// get each planet
	container.ForEach(accretionProcessor.Planets, func(value interface{}, index int64) {
		planet := value.(*types.Planet)
		// create planet environment builder
		planetEnvironment := NewPlanetEnvironment()
		// generate the planetesimal to a planet
		planetEnvironment.GeneratePlanet(sun, planet, randomTilt, doMoons, doGases, false)
		// append the found planet
		planets = append(planets, planet)
	})

	return sun, planets
}
