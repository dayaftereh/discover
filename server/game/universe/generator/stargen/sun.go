package stargen

import (
	"fmt"
	"math"

	"github.com/dayaftereh/discover/server/utils"
)

type Sun struct {
	Mass            float64 // star mass (units of solar masses)
	Luminosity      float64 // the stellar luminosity ratio is with respect to the sun
	EcosphereRadius float64 // the estimate range from the sun allowing existence of liquid water, Habitable ecosphere radius (AU)
	Life            float64 // The max life of the sun (years)
	Age             float64 // The age of the sun (years)
}

func luminosity(stellarMassRatio float64) float64 {
	n := 0.5*(2.0-stellarMassRatio) + 4.4
	if stellarMassRatio < 1.0 {
		n = 1.75*(stellarMassRatio-0.1) + 3.325
	}
	return (math.Pow(stellarMassRatio, n))
}

func NewSun(mass float64) *Sun {
	// calculate luminosity
	luminosity := luminosity(mass)

	// calculate the life of the sun
	life := 1e10 * (mass / luminosity)

	// check if life lager then max age
	maxAge := MaxSunAge
	if life > maxAge {
		maxAge = life
	}

	// generate a random age of the sun
	age := utils.RandFloat64(MinSunAge, maxAge)

	return &Sun{
		Mass:            mass,
		Luminosity:      luminosity,
		EcosphereRadius: math.Sqrt(luminosity),
		Life:            life,
		Age:             age,
	}
}

func (sun *Sun) String() string {
	s := fmt.Sprintf("Sun: [\n")
	s = fmt.Sprintf("%s Mass: %f\n", s, sun.Mass)
	s = fmt.Sprintf("%s Luminosity: %f\n", s, sun.Luminosity)
	s = fmt.Sprintf("%s EcosphereRadius: %f\n", s, sun.EcosphereRadius)
	s = fmt.Sprintf("%s Life: %f\n", s, sun.Life)
	s = fmt.Sprintf("%s Age: %f\n", s, sun.Age)
	s = fmt.Sprintf("%s]\n", s)
	return s
}
