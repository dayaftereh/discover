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

func massToLuminosity(mass float64) float64 {
	if mass <= 0.6224 {
		return 0.3815 * math.Pow(mass, 2.5185)
	}

	if mass <= 1.0 {
		return math.Pow(mass, 4.551)
	}

	if mass <= 3.1623 {
		return math.Pow(mass, 4.351)
	}

	if mass <= 16.0 {
		return 2.7563 * math.Pow(mass, 3.4704)
	}

	return 42.321 * math.Pow(mass, 2.4853)
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
	luminosity := massToLuminosity(mass)

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
