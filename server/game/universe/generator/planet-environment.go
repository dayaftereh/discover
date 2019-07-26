package generator

import (
	"math"

	"github.com/dayaftereh/discover/server/utils"
)

type PlanetEnvironment struct{}

func (planetEnvironment *PlanetEnvironment) about(value, variation, float64) float64 {
	return (value + (value * utils.RandFloat64(-variation, variation)))
}

func (planetEnvironment *PlanetEnvironment) inclination(orbitRadius float64) float64 {
	temp := (int)(math.Pow(orbitRadius, 0.2) * planetEnvironment.about(EarthAxialTilt, 0.4))
	return (math.Mod(temp, 360.0))
}

func (planetEnvironment *PlanetEnvironment) period(separation, smallMass, largeMass float64) float64 {
	eriodInYears := math.Sqrt((separation * separation * separation) / (smallMass + largeMass))
	return (eriodInYears * DaysInAYear)
}

func (planetEnvironment *PlanetEnvironment) orbitZone(luminosity, orbitRadius float64) OrbitZone {
	if orbitRadius < (4.0 * math.Sqrt(luminosity)) {
		return Orbit1
	} else if orbitRadius < (15.0 * math.Sqrt(luminosity)) {
		return Orbit2
	} else {
		return Orbit3
	}
}

func (planetEnvironment *PlanetEnvironment) GeneratePlanet(sun *Sun, planet *Planet, randomTilt bool) {

}
