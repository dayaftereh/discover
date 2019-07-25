package generator

import (
	"github.com/dayaftereh/discover/server/game/data"
	"github.com/dayaftereh/discover/server/utils"
)

const (
	// SolarMass is the factor to convert M to kg
	SolarMass float64 = 1.98847e30
	// SolarRadius is the factor to convert R to meter
	SolarRadius float64 = 6.957e8
)

func randSun() *data.Sun {
	// get a random stellar class
	class := RandStellarClass()

	// generate a random mass based on the stellar classification in kg
	mass := utils.RandFromRange(class.Mass) * SolarMass

	// generate a random radius based on the stellar classification in meter
	radius := utils.RandFromRange(class.Radius) * SolarRadius

	// generate a random temperature based on the stellar classification in kelvins
	temperature := utils.RandFromRange(class.Temperature) * 1000.0

	// generate a random luminosity based on the stellar classification in L
	luminosity := utils.RandFromRange(class.Luminosity)

	// create the sun data
	return &data.Sun{
		Class:       class.Class,
		Color:       class.Color,
		Mass:        mass,
		Radius:      radius,
		Temperature: temperature,
		Luminosity:  luminosity,
	}
}
