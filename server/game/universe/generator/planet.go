package generator

import (
	"math"

	"github.com/dayaftereh/discover/server/utils"

	"github.com/dayaftereh/discover/server/game/data"
)

const EccentricityCoeff float64 = 0.077

func randPlanets(sun *data.Sun) {
	startOrbitRadius := nearestPlanet(sun)
	massLeft := sun.Mass * 0.98 // use the mass of the sun

	planets := make([]*data.Planet, 0)

	for massLeft < 1000 {
		mass := planetMass(massLeft)
		massLeft -= mass

	}
}

func planetMass(massLeft float64) float64 {
	return utils.RandFloat64(0, massLeft)
}

func planetRadius(mass /*in kg*/ float64, density float64) float64 {
	massInGrams := mass * 1000.0
	volume := massInGrams / density
	cmPerMeter := 1000.0
	return math.Pow((3.0*volume)/(4.0*math.Pi), (1.0/3.0)) / cmPerMeter
}

func planetDensity(mass float64, orbitRadius /*in m*/ float64, sun *data.Sun) float64 {
	rEcosphere := math.Sqrt(sun.Luminosity)

	solarMass := sun.Mass / SolarMass
	sunMassInEarthMasses := 332775.64
	orbitRadiusAU := orbitRadius / AU

	tmp := math.Pow(mass*sunMassInEarthMasses, (1.0 / 8.0))
	tmp = tmp * math.Sqrt(math.Sqrt(rEcosphere/orbitRadiusAU))
	return tmp * 5.5
}

func nextPlanetOrbitRadius(sun *data.Sun, last float64) float64 {
	max := farthestPlanet(sun)
	return utils.RandFloat64(last, max)
}

func nearestPlanet(sun *data.Sun) float64 {
	stellarMassRatio := sun.Mass / SolarMass
	return (0.3 * math.Pow(stellarMassRatio, (1.0/3.0)))
}

func farthestPlanet(sun *data.Sun) float64 {
	stellarMassRatio := sun.Mass / SolarMass
	return (50.0 * math.Pow(stellarMassRatio, (1.0/3.0)))
}

func randInclination() float64 {
	return 0
}

func randEccentricity() float64 {
	E := 1.0 - math.Pow(utils.RandFloat64(0.0, 1.0), EccentricityCoeff)
	if E > 0.9999 {
		E := 0.999
	}
	return E
}
