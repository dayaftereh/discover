package stargen

import (
	"math"

	"github.com/dayaftereh/discover/server/mathf"

	"github.com/dayaftereh/discover/server/game/persistence/types"
)

type HabitableZoneMode string

const (
	RecentVenus               HabitableZoneMode = "recent-venus"
	RunawayGreenhouse         HabitableZoneMode = "runaway-greenhouse"
	MoistGreenhouse           HabitableZoneMode = "moist-greenhouse"
	EarthLike                 HabitableZoneMode = "earth-like"
	FirstCO2CondensationLimit HabitableZoneMode = "first-co2"
	MaximumGreenhouse         HabitableZoneMode = "max-greenhouse"
	EarlyMars                 HabitableZoneMode = "early-mars"
	TwoAUCloudLimit           HabitableZoneMode = "two-au"
)

func quadfix(x, y, w, z, p, q float64) (float64, float64, float64) {
	a := ((q * (w - x)) - (w * y) + (p * (y - z)) + (x * z)) / ((p - w) * (p - x) * (w - x))
	b := ((q * (math.Pow(x, 2.0) - math.Pow(w, 2.0))) + (math.Pow(w, 2.0) * y) - (math.Pow(x, 2.0) * z) + (math.Pow(p, 2.0) * (z - y))) / ((p - w) * (p - x) * (w - x))
	c := ((q * w * x * (w - x)) + (p * ((p * w * y) - (math.Pow(w, 2) * y) - (p * x * z) + (math.Pow(x, 2) * z)))) / ((p - w) * (p - x) * (w - x))
	return a, b, c
}

func quadTrend(a, b, c, x float64) float64 {
	return (a * math.Pow(x, 2.0)) + (b * x) + c
}

func planetRadiusHelper(planetMass, mass1, radius1, mass2, radius2, mass3, radius3 float64) float64 {
	a, b, c := quadfix(mass1, radius1, mass2, radius2, mass3, radius3)
	radius := quadTrend(a, b, c, planetMass)
	return radius
}

func logFix(x, y, w, z float64) (float64, float64) {
	a := ((y * math.Log(w)) - (z * math.Log(x))) / (math.Log(w) - math.Log(x))
	b := (z - y) / (math.Log(w) - math.Log(x))
	return a, b
}

func lnTrend(a, b, x float64) float64 {
	return a + (b * math.Log(x))
}

func planetRadiusHelper2(planetMass, mass1, radius1, mass2, radius2 float64) float64 {
	a, b := logFix(mass1, radius1, mass2, radius2)
	radius := lnTrend(a, b, planetMass)
	return radius
}

func calculateStellarFlux(a, b, c, d, seff, starTemp, starLuminosity float64) float64 {
	t := starTemp - 5780.0
	return seff + (a * t) + (b * math.Pow(t, 2.0)) + (c * math.Pow(t, 3.0)) + (d * math.Pow(t, 4.0))
}

func habitableZoneDistanceHelper(effectiveTemperature, luminosity float64, mode HabitableZoneMode, mass float64) float64 {
	mass = mathf.Clamp(mass, 0.1, 10.0)

	if mode == RecentVenus {
		return calculateStellarFlux(2.136E-4, 2.533E-8, -1.332E-11, -3.097E-15, 1.776, effectiveTemperature, luminosity)
	}

	if mode == RunawayGreenhouse {
		var a, b, c, d, seff float64
		if mass < 1.0 {
			a = planetRadiusHelper(mass, 0.1, 1.209E-4, 1.0, 1.332E-4, 5.0, 1.433E-4)
			b = planetRadiusHelper(mass, 0.1, 1.404E-8, 1.0, 1.58E-8, 5.0, 1.707E-8)
			c = planetRadiusHelper(mass, 0.1, -7.418E-12, 1.0, -8.308E-12, 5.0, -8.968E-12)
			d = planetRadiusHelper(mass, 0.1, -1.713E-15, 1.0, -1.931E-15, 5.0, -2.084E-15)
			seff = planetRadiusHelper(mass, 0.1, 0.99, 1.0, 1.107, 5.0, 1.188)
		} else {
			a = planetRadiusHelper2(mass, 1.0, 1.332E-4, 5.0, 1.433E-4)
			b = planetRadiusHelper2(mass, 1.0, 1.58E-8, 5.0, 1.707E-8)
			c = planetRadiusHelper2(mass, 1.0, -8.308E-12, 5.0, -8.968E-12)
			d = planetRadiusHelper2(mass, 1.0, -1.931E-15, 5.0, -2.084E-15)
			seff = planetRadiusHelper2(mass, 1.0, 1.107, 5.0, 1.188)
		}
		return calculateStellarFlux(a, b, c, d, seff, effectiveTemperature, luminosity)
	}

	if mode == MoistGreenhouse {
		stellarFluxGreen1 := calculateStellarFlux(1.332E-4, 1.58E-8, -8.308E-12, -1.931E-15, 1.107, effectiveTemperature, luminosity)
		stellarFluxMoist1 := calculateStellarFlux(8.1774E-5, 1.7063E-9, -4.3241E-12, -6.6462E-16, 1.0140, effectiveTemperature, luminosity)
		stellarFluxMax1 := calculateStellarFlux(5.8942E-5, 1.6558E-9, -3.0045E-12, -5.2983E-16, 0.3438, effectiveTemperature, luminosity)
		diff := stellarFluxGreen1 - stellarFluxMax1
		percent := (stellarFluxGreen1 - stellarFluxMoist1) / diff

		var a, b, c, d, seff float64
		if mass < 1.0 {
			a = planetRadiusHelper(mass, 0.1, 1.209E-4, 1.0, 1.332E-4, 5.0, 1.433E-4)
			b = planetRadiusHelper(mass, 0.1, 1.404E-8, 1.0, 1.58E-8, 5.0, 1.707E-8)
			c = planetRadiusHelper(mass, 0.1, -7.418E-12, 1.0, -8.308E-12, 5.0, -8.968E-12)
			d = planetRadiusHelper(mass, 0.1, -1.713E-15, 1.0, -1.931E-15, 5.0, -2.084E-15)
			seff = planetRadiusHelper(mass, 0.1, 0.99, 1.0, 1.107, 5.0, 1.188)
		} else {
			a = planetRadiusHelper2(mass, 1.0, 1.332E-4, 5.0, 1.433E-4)
			b = planetRadiusHelper2(mass, 1.0, 1.58E-8, 5.0, 1.707E-8)
			c = planetRadiusHelper2(mass, 1.0, -8.308E-12, 5.0, -8.968E-12)
			d = planetRadiusHelper2(mass, 1.0, -1.931E-15, 5.0, -2.084E-15)
			seff = planetRadiusHelper2(mass, 1.0, 1.107, 5.0, 1.188)
		}
		stellarFluxGreen2 := calculateStellarFlux(a, b, c, d, seff, effectiveTemperature, luminosity)
		stellarFluxMax2 := calculateStellarFlux(6.171E-5, 1.698E-9, -3.198E-12, -5.575E-16, 0.356, effectiveTemperature, luminosity)
		diff = stellarFluxGreen2 - stellarFluxMax2
		temp := diff * percent
		return stellarFluxGreen2 - (temp * diff)
	}

	if mode == EarthLike {
		stellarFluxGreen1 := calculateStellarFlux(1.332E-4, 1.58E-8, -8.308E-12, -1.931E-15, 1.107, effectiveTemperature, luminosity)
		stellarFluxEarth1 := calculateStellarFlux(8.3104E-5, 1.7677E-9, -4.39E-12, -6.79E-16, 1.0, effectiveTemperature, luminosity)
		stellarFluxMax1 := calculateStellarFlux(5.8942E-5, 1.6558E-9, -3.0045E-12, -5.2983E-16, 0.3438, effectiveTemperature, luminosity)
		diff := stellarFluxGreen1 - stellarFluxMax1
		percent := (stellarFluxGreen1 - stellarFluxEarth1) / diff

		var a, b, c, d, seff float64
		if mass < 1.0 {
			a = planetRadiusHelper(mass, 0.1, 1.209E-4, 1.0, 1.332E-4, 5.0, 1.433E-4)
			b = planetRadiusHelper(mass, 0.1, 1.404E-8, 1.0, 1.58E-8, 5.0, 1.707E-8)
			c = planetRadiusHelper(mass, 0.1, -7.418E-12, 1.0, -8.308E-12, 5.0, -8.968E-12)
			d = planetRadiusHelper(mass, 0.1, -1.713E-15, 1.0, -1.931E-15, 5.0, -2.084E-15)
			seff = planetRadiusHelper(mass, 0.1, 0.99, 1.0, 1.107, 5.0, 1.188)
		} else {
			a = planetRadiusHelper2(mass, 1.0, 1.332E-4, 5.0, 1.433E-4)
			b = planetRadiusHelper2(mass, 1.0, 1.58E-8, 5.0, 1.707E-8)
			c = planetRadiusHelper2(mass, 1.0, -8.308E-12, 5.0, -8.968E-12)
			d = planetRadiusHelper2(mass, 1.0, -1.931E-15, 5.0, -2.084E-15)
			seff = planetRadiusHelper2(mass, 1.0, 1.107, 5.0, 1.188)
		}

		stellarFluxGreen2 := calculateStellarFlux(a, b, c, d, seff, effectiveTemperature, luminosity)
		stellarFluxMax2 := calculateStellarFlux(6.171E-5, 1.698E-9, -3.198E-12, -5.575E-16, 0.356, effectiveTemperature, luminosity)
		diff = stellarFluxGreen2 - stellarFluxMax2
		temp := diff * percent
		return stellarFluxGreen2 - (temp * diff)
	}

	if mode == FirstCO2CondensationLimit {
		return calculateStellarFlux(4.4499e-5, 1.4065e-10, 2.2750e-12, -3.3509e-16, 0.5408, effectiveTemperature, luminosity)
	}

	if mode == MaximumGreenhouse {
		return calculateStellarFlux(6.171E-5, 1.698E-9, -3.198E-12, -5.575E-16, 0.356, effectiveTemperature, luminosity)
	}

	if mode == EarlyMars {
		return calculateStellarFlux(5.547E-5, 1.526E-9, -2.874E-12, -5.011E-16, 0.32, effectiveTemperature, luminosity)
	}

	if mode == TwoAUCloudLimit {
		return calculateStellarFlux(4.2588e-5, 1.1963e-9, -2.1709e-12, -3.8282e-16, 0.2484, effectiveTemperature, luminosity)
	}

	return 0.0
}

func habitableZoneDistance(sun *types.Sun, mode HabitableZoneMode, mass float64) float64 {
	if sun.EffectiveTemperature >= 2600 && sun.EffectiveTemperature <= 7200 {
		stellarFlux := habitableZoneDistanceHelper(sun.EffectiveTemperature, sun.Luminosity, mode, mass)
		return math.Sqrt(sun.Luminosity / stellarFlux)
	}

	if mode == RecentVenus || mode == RunawayGreenhouse || mode == MoistGreenhouse {
		return math.Sqrt(sun.Luminosity / 1.51)
	}

	if mode == EarthLike {
		return math.Sqrt(sun.Luminosity)
	}

	if mode == FirstCO2CondensationLimit || mode == MaximumGreenhouse || mode == EarlyMars || mode == TwoAUCloudLimit {
		return math.Sqrt(sun.Luminosity / 0.48)
	}

	return 0.0
}
