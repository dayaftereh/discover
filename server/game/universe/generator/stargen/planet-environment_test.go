package stargen_test

import (
	"testing"

	"github.com/dayaftereh/discover/server/utils/container"

	"github.com/dayaftereh/discover/server/game/universe/generator/stargen"
)

func TestPlanetEnvironmentGeneratePlanet(t *testing.T) {
	sun := stargen.NewSun(0.4)
	accretionProcessor := stargen.NewAccretionProcessor()
	outerDustLimit := accretionProcessor.StellarDustLimit(sun.Mass)

	accretionProcessor.DistPlanetaryMasses(
		sun.Mass,
		sun.Luminosity,
		0.0,
		outerDustLimit,
		0.0,
		stargen.DustDensityCoefficient,
		true,
	)

	container.ForEach(accretionProcessor.Planets, func(value interface{}, index int64) {
		planet := value.(*stargen.Planet)

		planetEnvironment := stargen.NewPlanetEnvironment()
		planetEnvironment.GeneratePlanet(sun, planet, true, true, true, false)

		t.Log(planet)
	})
}
