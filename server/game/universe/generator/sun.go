package generator

import (
	"fmt"
	"math"

	"github.com/dayaftereh/discover/server/utils"
)

type Sun struct {
	Mass       float64
	Luminosity float64
	REcosphere float64
	Life       float64
	Age        float64
}

func luminosity(stellMassRatio float64) float64 {
	n := 0.5*(2.0-stellMassRatio) + 4.4
	if stellMassRatio < 1.0 {
		n = 1.75*(stellMassRatio-0.1) + 3.325
	}
	return (math.Pow(stellMassRatio, n))
}

func NewSun(mass float64) *Sun {
	// caluclate luminosity
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
		Mass:       mass,
		Luminosity: luminosity,
		REcosphere: math.Sqrt(luminosity),
		Life:       life,
		Age:        age,
	}
}

func (sun *Sun) String() string {
	s := fmt.Sprintf("Sun: [\n")
	s = fmt.Sprintf("%s Mass: %f\n", s, sun.Mass)
	s = fmt.Sprintf("%s Luminosity: %f\n", s, sun.Luminosity)
	s = fmt.Sprintf("%s REcosphere: %f\n", s, sun.REcosphere)
	s = fmt.Sprintf("%s Life: %f\n", s, sun.Life)
	s = fmt.Sprintf("%s Age: %f\n", s, sun.Age)
	s = fmt.Sprintf("%s]\n", s)
	return s
}
