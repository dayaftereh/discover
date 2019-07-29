package generator_test

import (
	"testing"

	"github.com/dayaftereh/discover/server/game/universe/generator"
)

func TestPlanetEnvironmentGeneratePlanet(t *testing.T) {
	dust_density_coeff := 2.0e-3
	accretionProcessor := generator.NewAccretionProcessor()
	sun := generator.NewSun(0.4)
	outerDustLimit := accretionProcessor.StellarDustLimit(sun.Mass)

	accretionProcessor.DistPlanetaryMasses(
		sun.Mass,
		sun.Luminosity,
		0.0,
		outerDustLimit,
		0.0,
		dust_density_coeff,
		true,
	)

	// find the place to place the planet
	planet := accretionProcessor.Planet
	for planet != nil {
		planetEnvironment := generator.NewPlanetEnvironment()
		planetEnvironment.GeneratePlanet(sun, planet, true, true, true, false)

		planet = planet.NextPlanet
	}

	t.Log(accretionProcessor.Planet)
}
