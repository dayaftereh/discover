package generator_test

import (
	"testing"

	"github.com/dayaftereh/discover/server/utils"

	"github.com/dayaftereh/discover/server/game/universe/generator"
)

func TestAccretionProcessorDistPlanetaryMasses(t *testing.T) {

	timeSum := 0.0
	planetsCountSum := 0
	moonSum := 0

	runs := 1000
	for i := 0; i < runs; i++ {

		accretionProcessor := generator.NewAccretionProcessor()

		sun := generator.NewSun(0.4)
		outerDustLimit := accretionProcessor.StellarDustLimit(sun.Mass)

		dust_density_coeff := 2.0e-3

		now := utils.SystemMillis()

		accretionProcessor.DistPlanetaryMasses(
			sun.Mass,
			sun.Luminosity,
			0.0,
			outerDustLimit,
			0.0,
			dust_density_coeff,
			true,
		)
		timeSum += utils.SystemMillis() - now

		planet := accretionProcessor.Planet
		for planet != nil {
			planetsCountSum++
			for _, _ = range planet.Moons {
				moonSum++
			}

			planet = planet.NextPlanet
		}

	}

	t.Logf("Avg. Moon: %f", (float64(moonSum) / float64(runs)))

	t.Logf("Avg. Planets: %f", (float64(planetsCountSum) / float64(runs)))
	t.Logf("Avg. Time: %f ms", (float64(timeSum) / float64(runs)))

}
