package stargen

import (
	"math"

	"github.com/dayaftereh/discover/server/game/persistence/types"
	"github.com/dayaftereh/discover/server/utils"
)

func luminosity(stellarMassRatio float64) float64 {
	n := 0.5*(2.0-stellarMassRatio) + 4.4
	if stellarMassRatio < 1.0 {
		n = 1.75*(stellarMassRatio-0.1) + 3.325
	}
	return (math.Pow(stellarMassRatio, n))
}

func NewRandomSun() *types.Sun {
	// get a random stellar class
	stellarClass := RandStellarClass()
	// generate a random star mass based on the stellar class
	mass := utils.RandFromRange(stellarClass.Mass)
	// create a new sun for the stellar system
	sun := NewSun(mass)

	sun.Class = stellarClass.Class
	sun.Color = stellarClass.Color

	sun.EffectiveTemperature = utils.RandFromRange(stellarClass.Temperature) * 1000.0

	return sun
}

func NewSun(mass float64) *types.Sun {
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

	return &types.Sun{
		Mass:            mass,
		Luminosity:      luminosity,
		EcosphereRadius: math.Sqrt(luminosity),
		Life:            life,
		Age:             age,
	}
}
