package stargen

import (
	"math"
	"sort"

	"github.com/dayaftereh/discover/server/mathf"

	"github.com/dayaftereh/discover/server/utils"
)

type PlanetEnvironment struct{}

func NewPlanetEnvironment() *PlanetEnvironment {
	return &PlanetEnvironment{}
}

// about returns a value within a certain variation of the exact value given it in 'value'.
func (planetEnvironment *PlanetEnvironment) about(value float64, variation float64) float64 {
	return (value + (value * utils.RandFloat64(-variation, variation)))
}

// inclination calculates the orbital radius is expected in units of Astronomical Units (AU)
// Inclination is returned in units of degrees.
func (planetEnvironment *PlanetEnvironment) inclination(orbitRadius float64) float64 {
	temp := int64(math.Round(math.Pow(orbitRadius, 0.2) * planetEnvironment.about(EarthAxialTilt, 0.4)))
	return (math.Mod(float64(temp), 360.0))
}

// period calculates the period in terms of Earth days. The separation is in units of AU, and both masses are in units of solar masses.
func (planetEnvironment *PlanetEnvironment) period(separation, smallMass, largeMass float64) float64 {
	periodInYears := math.Sqrt(math.Pow(separation, 3.0) / (smallMass + largeMass))
	return (periodInYears * DaysInAYear)
}

// rootMeanSquareVelocity calculates the Root Mean Square (RMS) velocity of the	molecule or atom. This is Fogg's eq.16.
// The velocity returned is in cm/sec. Orbital radius is in A.U.(ie: in units of the earth's orbital radius).
func (planetEnvironment *PlanetEnvironment) rootMeanSquareVelocity(molecularWeight, exosphericTemp float64) float64 {
	return math.Sqrt((3.0*MolarGasConst*exosphericTemp)/molecularWeight) * CMPerMeter
}

// kothariRadius returns the radius of the planet in kilometers.  The mass passed in is in units of solar masses.
// This formula is listed as eq.9 in Fogg's article, although some typos crop up in that eq.
// See "The Internal Constitution of Planets", by Dr. D. S. Kothari, Mon. Not. of the Royal Astronomical Society, vol 96
// pp.833-843, 1936 for the derivation.  Specifically, this is Kothari's eq.23, which appears on page 840.
func (planetEnvironment *PlanetEnvironment) kothariRadius(mass float64, gaint bool, zone OrbitZone) float64 {
	var atomicWeight, atomicNum float64
	if zone == Orbit1 {
		if gaint {
			atomicWeight = 9.5
			atomicNum = 4.5
		} else {
			atomicWeight = 15.0
			atomicNum = 8.0
		}
	} else {
		if zone == Orbit2 {
			if gaint {
				atomicWeight = 2.47
				atomicNum = 2.0
			} else {
				atomicWeight = 10.0
				atomicNum = 5.0
			}
		} else {
			if gaint {
				atomicWeight = 7.0
				atomicNum = 4.0
			} else {
				atomicWeight = 10.0
				atomicNum = 5.0
			}
		}
	}

	temp1 := atomicWeight * atomicNum
	temp := (2.0 * Beta20 * math.Pow(SolarMassInGrams, (1.0/3.0))) / (A1_20 * math.Pow(temp1, (1.0/3.0)))

	temp2 := A2_20 * math.Pow(atomicWeight, (4.0/3.0)) * math.Pow(SolarMassInGrams, (2.0/3.0))
	temp2 = temp2 * math.Pow(mass, (2.0/3.0))
	temp2 = temp2 / (A1_20 * math.Pow(atomicNum, 2.0))
	temp2 = 1.0 + temp2
	temp = temp / temp2
	temp = (temp * math.Pow(mass, (1.0/3.0))) / CMPerKM

	temp /= JimsFudge /* Make Earth = actual earth */

	return temp
}

// empiricalDensity calculates the density. The mass passed in is in units of solar masses, and the orbital radius is in units of AU.
// The density is returned in units of grams/cc.
func (planetEnvironment *PlanetEnvironment) empiricalDensity(mass float64, orbitRadius float64, rEcosphere float64, gasGiant bool) float64 {
	temp := math.Pow(mass*SunMassInEarthMasses, (1.0 / 8.0))
	temp = temp * math.Sqrt(math.Sqrt(rEcosphere/orbitRadius))
	if gasGiant {
		return temp * 1.2
	}
	return temp * 5.5
}

// volumeRadius calculates the radius from the volume. The mass is in units of solar masses, and the density is in units of grams/cc.
// The radius returned is in units of km.
func (planetEnvironment *PlanetEnvironment) volumeRadius(mass float64, density float64) float64 {
	mass = mass * SolarMassInGrams
	volume := mass / density
	return math.Pow((3.0*volume)/(4.0*math.Pi), (1.0/3.0)) / CMPerKM
}

// volumeDensity calculates the volume density. The mass passed in is in units of solar masses, and the equatorial radius is in km.
// The density is returned in units of grams/cc.
func (planetEnvironment *PlanetEnvironment) volumeDensity(mass float64, radius float64) float64 {
	mass = mass * SolarMassInGrams
	radiusInCm := radius * CMPerKM
	volume := (4.0 * math.Pi * math.Pow(radiusInCm, 3.0)) / 3.0
	return mass / volume
}

// dayLength calculates the length of the day in units of hours.
// Input parameters are mass (in solar masses), radius (in Km), orbital period (in days), orbital radius (in AU), density (in g/cc), ccentricity, and whether it is a gas giant or not.
func (planetEnvironment *PlanetEnvironment) dayLength(planet *Planet, sun *Sun) (float64, bool) {
	/*--------------------------------------------------------------------------*/
	/*	 Fogg's information for this routine came from Dole "Habitable Planets	*/
	/* for Man", Blaisdell Publishing Company, NY, 1964.  From this, he came	*/
	/* up with his eq.12, which is the equation for the 'base_angular_velocity' */
	/* below.  He then used an equation for the change in angular velocity per	*/
	/* time (dw/dt) from P. Goldreich and S. Soter's paper "Q in the Solar		*/
	/* System" in Icarus, vol 5, pp.375-389 (1966).	 Using as a comparison the	*/
	/* change in angular velocity for the Earth, Fogg has come up with an		*/
	/* approximation for our new planet (his eq.13) and take that into account. */
	/* This is used to find 'change_in_angular_velocity' below.					*/
	/*--------------------------------------------------------------------------*/

	planetaryMassInGrams := planet.Mass * SolarMassInGrams
	equatorialRadiusInCM := planet.Radius * CMPerKM
	yearInHours := planet.OrbitPeriod * 24.0
	giant := (planet.Type == PlanetGasGiant || planet.Type == PlanetSubGasGiant || planet.Type == PlanetSubSubGasGiant)

	stopped := false

	k2 := 0.33
	if giant {
		k2 = 0.24
	}

	baseAngularVelocity := math.Sqrt(2.0 * J * planetaryMassInGrams / (k2 * math.Pow(equatorialRadiusInCM, 2.0)))

	/*	This next calculation determines how much the planet's rotation is	 */
	/*	slowed by the presence of the star.								 */

	changeInAngularVelocity := ChangeInErthAngVel * (planet.Density / EarthDensity) * (equatorialRadiusInCM / EarthRadius) *
		(EarthMassInGrams / planetaryMassInGrams) * math.Pow(sun.Mass, 2.0) * (1.0 / math.Pow(planet.SemiMajorAxis, 6.0))
	angVelocity := baseAngularVelocity + (changeInAngularVelocity * sun.Age)

	var daysInAYear float64
	if angVelocity <= 0.0 {
		stopped = true
		daysInAYear = math.MaxFloat64
	} else {
		daysInAYear = (2.0 * math.Pi) / (SecondsPerHour * angVelocity)
	}

	if daysInAYear > yearInHours || stopped {
		if planet.Eccentricity > 0.1 {
			spinResonanceFactor := (1.0 - planet.Eccentricity) / (1.0 + planet.Eccentricity)
			return (spinResonanceFactor * yearInHours), true
		}
		return yearInHours, false
	}
	return daysInAYear, false
}

// gravity calculates the surface gravity of a planet.
// The acceleration is in units of cm/sec2, and the gravity is returned in units of Earth gravities.
func (planetEnvironment *PlanetEnvironment) gravity(acceleration float64) float64 {
	return acceleration / EarthAcceleration
}

// acceleration calculates the surface acceleration of a planet.
// The mass is in units of solar masses, the radius in terms of km, and the acceleration is returned in units of cm/sec2.
func (planetEnvironment *PlanetEnvironment) acceleration(mass float64, radius float64) float64 {
	return GravityConstant * (mass * SolarMassInGrams) / math.Pow(radius*CMPerKM, 2.0)
}

// orbitZone returns the orbital 'zone' of the particle from the given the orbital radius of a planet in AU
func (planetEnvironment *PlanetEnvironment) orbitZone(luminosity, orbitRadius float64) OrbitZone {
	if orbitRadius < (4.0 * math.Sqrt(luminosity)) {
		return Orbit1
	} else if orbitRadius < (15.0 * math.Sqrt(luminosity)) {
		return Orbit2
	} else {
		return Orbit3
	}
}

// gasLife calculates the number of years it takes for 1/e of a gas to escape from a planet's atmosphere.
func (planetEnvironment *PlanetEnvironment) gasLife(molecularWeight float64, exosphericTemp float64, surfGrav float64, radius float64) float64 {
	// Taken from Dole p. 34. He cites Jeans (1916) & Jones (1923)

	v := planetEnvironment.rootMeanSquareVelocity(molecularWeight, exosphericTemp)
	g := surfGrav * EarthAcceleration
	r := radius * CMPerKM
	t := (math.Pow(v, 3.0) / (2.0 * math.Pow(g, 2.0) * r)) * math.Exp((3.0*g*r)/math.Pow(v, 2.0))
	years := t / (SecondsPerHour * 24.0 * DaysInAYear)

	if years > 2.0e10 {
		return math.MaxFloat64
	}
	return years
}

// escapeVelocity calculates the escape velocity.
// The mass is in units of solar mass, the radius in kilometers, and the velocity returned is in cm/sec.
func (planetEnvironment *PlanetEnvironment) escapeVelocity(mass float64, radius float64) float64 {
	// Note that it appears that Fogg's eq.15 is incorrect
	massInGrams := mass * SolarMassInGrams
	radiusInCm := radius * CMPerKM
	return math.Sqrt(2.0 * GravityConstant * massInGrams / radiusInCm)
}

// moleculeLimit returns the smallest molecular weight retained by the body, which is useful for determining the atmosphere composition.
// Mass is in units of solar masses, and equatorial radius is in units of kilometers.
func (planetEnvironment *PlanetEnvironment) moleculeLimit(mass float64, equatRadius float64, exosphericTemp float64) float64 {
	escVelocity := planetEnvironment.escapeVelocity(mass, equatRadius)
	return (3.0 * MolarGasConst * exosphericTemp) / (math.Pow((escVelocity/GasRetentionThreshold)/CMPerMeter, 2.0))

}

// effectiveTemp calculates the effective temperature.
// The ecosphere radius is given in AU, the orbital radius in AU, and the temperature returned is in Kelvin.
func (planetEnvironment *PlanetEnvironment) effectiveTemp(ecosphereRadius float64, orbitRadius float64, albedo float64) float64 {
	// This is Fogg's eq.19.
	return math.Sqrt(ecosphereRadius/orbitRadius) * math.Sqrt(math.Sqrt((1.0-albedo)/(1.0-EarthAlbedo))) * EarthEffectiveTemp
}

func (planetEnvironment *PlanetEnvironment) estimatedTemp(ecosphereRadius float64, orbitRadius float64, albedo float64) float64 {
	return math.Sqrt(ecosphereRadius/orbitRadius) * math.Sqrt(math.Sqrt((1.0-albedo)/(1.0-EarthAlbedo))) * EarthAverageKelvin
}

func (planetEnvironment *PlanetEnvironment) minMolecularWeight(planet *Planet, sun *Sun) float64 {

	guess1 := planetEnvironment.moleculeLimit(planet.Mass, planet.Radius, planet.ExosphericTemperature)
	guess2 := guess1

	life := planetEnvironment.gasLife(guess1, planet.ExosphericTemperature, planet.SurfaceGravity, planet.Radius)

	if life > sun.Age {
		for loops := 0; loops < 25 && life > sun.Age; loops++ {
			guess1 = guess1 / 2.0
			life = planetEnvironment.gasLife(guess1, planet.ExosphericTemperature, planet.SurfaceGravity, planet.Radius)
		}
	} else {
		for loops := 0; loops < 25 && life < sun.Age; loops++ {
			guess2 = guess2 * 2.0
			life = planetEnvironment.gasLife(guess2, planet.ExosphericTemperature, planet.SurfaceGravity, planet.Radius)
		}
	}

	for loops := 0; loops < 25 && (guess2-guess1) > 0.1; loops++ {
		guess3 := (guess1 + guess2) / 2.0
		life = planetEnvironment.gasLife(guess3, planet.ExosphericTemperature, planet.SurfaceGravity, planet.Radius)
		if life < sun.Age {
			guess1 = guess3
		} else {
			guess2 = guess3
		}
	}

	life = planetEnvironment.gasLife(guess2, planet.ExosphericTemperature, planet.SurfaceGravity, planet.Radius)

	return guess2
}

func (planetEnvironment *PlanetEnvironment) greenHouse(ecosphereRadius, orbitRadius float64) bool {
	/*--------------------------------------------------------------------------*/
	/* Old grnhouse:                                                            */
	/*	Note that if the orbital radius of the planet is greater than or equal	*/
	/*	to R_inner, 99% of it's volatiles are assumed to have been deposited in */
	/*	surface reservoirs (otherwise, it suffers from the greenhouse effect).	*/
	/*--------------------------------------------------------------------------*/
	/*	if ((orb_radius < r_greenhouse) && (zone == 1)) */
	/*--------------------------------------------------------------------------*/
	/*	The new definition is based on the initial surface temperature and what	*/
	/*	state water is in. If it's too hot, the water will never condense out	*/
	/*	of the atmosphere, rain down and form an ocean. The albedo used here	*/
	/*	was chosen so that the boundary is about the same as the old method		*/
	/*	Neither zone, nor r_greenhouse are used in this version				JLB	*/
	/*--------------------------------------------------------------------------*/
	temp := planetEnvironment.effectiveTemp(ecosphereRadius, orbitRadius, GreenhouseTriggerAlbedo)
	if temp > FreezingPointOfWater {
		return true
	}
	return false
}

// This implements Fogg's eq.17.  The 'inventory' returned is unitless.
func (planetEnvironment *PlanetEnvironment) volInventory(mass, escapeVel, rmsVel, stellarMass float64, zone OrbitZone, greenhouseEffect, accretedGas bool) float64 {
	velocityRatio := escapeVel / rmsVel
	if velocityRatio < GasRetentionThreshold {
		return 0.0
	}

	var proportionConst float64
	switch zone {
	case Orbit1:
		proportionConst = 140000.0 /* 100 -> 140 JLB */
	case Orbit2:
		proportionConst = 75000.0
	case Orbit3:
		proportionConst = 250.0
	default:
		proportionConst = 0.0
	}
	erathUnit := mass * SunMassInEarthMasses
	temp1 := (proportionConst * erathUnit) / stellarMass
	temp2 := planetEnvironment.about(temp1, 0.2)
	temp2 = temp1
	if greenhouseEffect || accretedGas {
		return temp2
	}
	return temp2 / 140.0 /* 100 -> 140 JLB */
}

// pressure calculates the surface pressure.
// The pressure returned is in units of	millibars (mb).	The gravity is in units of Earth gravities, the radius in units of kilometers.
func (planetEnvironment *PlanetEnvironment) pressure(volatileGasInventory, equatRadius, gravity float64) float64 {
	// This implements Fogg's eq.18.
	// JLB: Apparently this assumed that earth pressure = 1000mb. I've added a fudge factor (EARTH_SURF_PRES_IN_MILLIBARS / 1000.) to correct for that
	equatRadius = EarthRadiusInKM / equatRadius
	return volatileGasInventory * gravity * (EarthSurfPersInMilliBars / MilliBarsPerBar) / math.Pow(equatRadius, 2.0)
}

// boilingPoint returns the boiling point of water in an atmosphere of surface pressure in millibars.
// The boiling point is returned in units of Kelvin.
//
func (planetEnvironment *PlanetEnvironment) boilingPoint(surfPressure float64) float64 {
	// This is Fogg's eq.21.
	surfacePressureInBars := surfPressure / MilliBarsPerBar
	return 1.0 / ((math.Log(surfacePressureInBars) / -5050.5) + (1.0 / 373.0))
}

func (planetEnvironment *PlanetEnvironment) lim(x float64) float64 {
	return x / math.Sqrt(math.Sqrt(1.0+x*x*x*x))
}

func (planetEnvironment *PlanetEnvironment) soft(v, max, min float64) float64 {
	dv := v - min
	dm := max - min
	return (planetEnvironment.lim(2.0*dv/dm-1.0)+1.0)/2.0*dm + min
}

// greenRise calculates the rise in temperature produced by the greenhouse effect
// The effective temperature given is in units of Kelvin, as is the rise in temperature produced by the, which is returned.
// I tuned this by changing a pow(x,.25) to pow(x,.4) to match Venus - JLB
func (planetEnvironment *PlanetEnvironment) greenRise(opticalDepth, effectiveTemp, surfPressure float64) float64 {
	// This is Fogg's eq.20, and is also Hart's eq.20 in his "Evolution of Earth's Atmosphere" article.
	convectionFactor := EarthConvectionFactor * math.Pow(surfPressure/EarthSurfPersInMilliBars, 0.4)
	rise := (math.Sqrt(math.Sqrt(1.0+0.75*opticalDepth)) - 1.0) * effectiveTemp * convectionFactor
	return math.Max(rise, 0.0)
}

// opacity returns the dimensionless quantity of optical depth, which is useful in determining the amount of greenhouse effect on a planet.
func (planetEnvironment *PlanetEnvironment) opacity(molecularWeight, surfPressure float64) float64 {
	opticalDepth := 0.0
	if molecularWeight >= 0.0 && molecularWeight < 10.0 {
		opticalDepth = opticalDepth + 3.0
	}
	if molecularWeight >= 10.0 && molecularWeight < 20.0 {
		opticalDepth = opticalDepth + 2.34
	}
	if molecularWeight >= 20.0 && molecularWeight < 30.0 {
		opticalDepth = opticalDepth + 1.0
	}
	if molecularWeight >= 30.0 && molecularWeight < 45.0 {
		opticalDepth = opticalDepth + 0.15
	}
	if molecularWeight >= 45 && molecularWeight < 100.0 {
		opticalDepth = opticalDepth + 0.05
	}
	if surfPressure >= (70.0 * EarthSurfPersInMilliBars) {
		return opticalDepth * 8.333
	}

	if surfPressure >= (50.0 * EarthSurfPersInMilliBars) {
		return opticalDepth * 6.666
	}

	if surfPressure >= (30.0 * EarthSurfPersInMilliBars) {
		return opticalDepth * 3.33
	}

	if surfPressure >= (10.0 * EarthSurfPersInMilliBars) {
		return opticalDepth * 2.0
	}

	if surfPressure >= (5.0 * EarthSurfPersInMilliBars) {
		return opticalDepth * 1.5
	}

	return opticalDepth
}

func (planetEnvironment *PlanetEnvironment) setPlanetTempRange(planet *Planet) {
	pressmod := 1.0 / math.Sqrt(1.0+20.0*planet.SurfacePressure/1000.0)
	ppmod := 1.0 / math.Sqrt(10.0+5.0*planet.SurfacePressure/1000.0)
	tiltmod := math.Abs(math.Cos(planet.AxialTilt*math.Pi/180.0) * math.Pow(1.0+planet.Eccentricity, 2.0))
	daymod := 1.0 / (200.0/planet.Day + 1.0)
	mh := math.Pow(1.0+daymod, pressmod)
	ml := math.Pow(1.0-daymod, pressmod)
	hi := mh * planet.SurfaceTemperature
	lo := ml * planet.SurfaceTemperature
	sh := hi + math.Pow((100.0+hi)*tiltmod, math.Sqrt(ppmod))
	wl := lo - math.Pow((150.0+lo)*tiltmod, math.Sqrt(ppmod))
	max := planet.SurfaceTemperature + math.Sqrt(planet.SurfaceTemperature)*10.0
	min := planet.SurfaceTemperature / math.Sqrt(planet.Day+24.0)

	lo = math.Max(lo, min)
	wl = math.Max(wl, 0.0)

	planet.HighTemperature = planetEnvironment.soft(hi, max, min)
	planet.LowTemperature = planetEnvironment.soft(lo, max, min)
	planet.MaxTemperature = planetEnvironment.soft(sh, max, min)
	planet.MinTemperature = planetEnvironment.soft(wl, max, min)
}

// hydroFraction returns the fraction of the planet covered with water.
// Given the volatile gas inventory and planetary radius of a planet (in Km).
func (planetEnvironment *PlanetEnvironment) hydroFraction(volatileGasInventory, planetRadius float64) float64 {
	// This function is Fogg's eq.22.
	// I have changed the function very slightly: the fraction of Earth's urface covered by water is 71%, not 75% as Fogg used.
	temp := (0.71 * volatileGasInventory / 1000.0) * math.Pow(EarthRadiusInKM/planetRadius, 2.0)
	return math.Min(temp, 1.0)
}

// cloudFraction returns the fraction of cloud cover available
// Given the surface temperature of a planet (in Kelvin)
func (planetEnvironment *PlanetEnvironment) cloudFraction(surfTemp, smallestMWRetained, equatRadius, hydroFraction float64) float64 {
	// This is Fogg's eq.23.
	// See Hart in "Icarus" (vol 33, pp23 - 39, 1978) for an explanation. This equation is Hart's eq.3.
	// I have modified it slightly using constants and relationships from Glass's book "Introduction to Planetary Geology", p.46.
	//  The 'CloudCoverageFactor' is the amount of surface area on Earth covered by one Kg. of cloud.
	if smallestMWRetained > WaterVapro {
		return 0.0
	}

	surfArea := 4.0 * math.Pi * math.Pow(equatRadius, 2.0)
	hydroMass := hydroFraction * surfArea * EarthWaterMassPerArea
	waterVaporInKG := (0.00000001 * hydroMass) * math.Exp(Q2_36*(surfTemp-EarthAverageKelvin))
	fraction := CloudCoverageFactor * waterVaporInKG / surfArea

	return math.Min(fraction, 1.0)
}

// iceFraction returns the fraction of the planet's surface covered by ice.
//  Given the surface temperature of a planet (in Kelvin)
func (planetEnvironment *PlanetEnvironment) iceFraction(hydroFraction, surfTemp float64) float64 {
	// This is Fogg's eq.24. See Hart[24] in Icarus vol.33, p.28 for an explanation.
	// I have changed a constant from 70 to 90 in order to bring it more in line with
	// the fraction of the Earth's surface covered with ice, which is approximately .016 (=1.6%).
	surfTemp = math.Min(surfTemp, 328.0)
	temp := math.Pow(((328.0 - surfTemp) / 90.0), 5.0)
	temp = math.Min(temp, (1.5 * hydroFraction))
	return math.Min(temp, 1.0)
}

// planetAlbedo
// The surface temperature passed in is in units of Kelvin.
// The cloud adjustment is the fraction of cloud cover obscuring each of the three major components of albedo that lie below the clouds.
func (planetEnvironment *PlanetEnvironment) planetAlbedo(waterFraction, cloudFraction, iceFraction, surfPressure float64) float64 {
	rockFraction := 1.0 - waterFraction - iceFraction
	components := 0.0
	if waterFraction > 0.0 {
		components += 1.0
	}
	if iceFraction > 0.0 {
		components += 1.0
	}
	if rockFraction > 0.0 {
		components += 1.0
	}
	cloudAdjustment := cloudFraction / components

	if rockFraction >= cloudAdjustment {
		rockFraction -= cloudAdjustment
	} else {
		rockFraction = 0.0
	}

	if waterFraction > cloudAdjustment {
		waterFraction -= cloudAdjustment
	} else {
		waterFraction = 0.0
	}

	if iceFraction > cloudAdjustment {
		iceFraction -= cloudAdjustment
	} else {
		iceFraction = 0.0
	}

	cloudPart := cloudFraction * CloudAlbedo /* about(...,0.2); */

	rockPart := rockFraction * RockyAlbedo   /* about(...,0.1); */
	icePart := iceFraction * IceAlbedo       /* about(...,0.1); */
	waterPart := waterFraction * WaterAlbedo /* about(...,0.2); */
	if surfPressure == 0.0 {
		waterPart = 0.0
		rockPart = rockFraction * RockyAirlessAlbedo /* about(...,0.3); */
		icePart = iceFraction * AirlessIceAlbedo     /* about(...,0.4); */
	}

	return cloudPart + rockPart + waterPart + icePart
}

func (planetEnvironment *PlanetEnvironment) calculateAndSetSurfaceTemp(planet *Planet, sun *Sun, first bool, lastWater, lastClouds, lastIce, lastTemp, lastAlbedo float64) {
	if first {
		planet.Albedo = EarthAlbedo

		effectiveTemp := planetEnvironment.effectiveTemp(sun.EcosphereRadius, planet.SemiMajorAxis, planet.Albedo)
		opacity := planetEnvironment.opacity(planet.MolecularWeight, planet.SurfacePressure)
		greenhouseTemp := planetEnvironment.greenRise(opacity, effectiveTemp, planet.SurfacePressure)
		planet.SurfaceTemperature = effectiveTemp + greenhouseTemp

		planetEnvironment.setPlanetTempRange(planet)
	}

	if planet.GreenhouseEffect && planet.MaxTemperature < planet.BoilPoint {
		planet.GreenhouseEffect = false
		accretedGas := (planet.GasMass / planet.Mass) < 0.000001
		planet.VolatileGasInventory = planetEnvironment.volInventory(
			planet.Mass,
			planet.EscapeVelocity,
			planet.RootMeanSquareVelocity,
			sun.Mass,
			planet.OrbitZone,
			planet.GreenhouseEffect,
			accretedGas)
		planet.SurfacePressure = planetEnvironment.pressure(planet.VolatileGasInventory, planet.Radius, planet.SurfaceGravity)
		planet.BoilPoint = planetEnvironment.boilingPoint(planet.SurfacePressure)
	}

	planet.Hydrosphere = planetEnvironment.hydroFraction(planet.VolatileGasInventory, planet.Radius)
	planet.CloudCover = planetEnvironment.cloudFraction(planet.SurfaceTemperature, planet.MolecularWeight, planet.Radius, planet.Hydrosphere)
	planet.IceCover = planetEnvironment.iceFraction(planet.Hydrosphere, planet.SurfaceTemperature)

	if planet.GreenhouseEffect && planet.SurfacePressure > 0.0 {
		planet.CloudCover = 1.0
	}

	boilOff := false
	if planet.HighTemperature >= planet.BoilPoint && !first && !(mathf.CloseEquals(planet.Day, planet.OrbitPeriod*24.0) || planet.ResonantPeriod) {
		boilOff = true
		planet.Hydrosphere = 0.0
		if planet.MolecularWeight > WaterVapro {
			planet.CloudCover = 0.0
		} else {
			planet.CloudCover = 1.0
		}
	}

	if planet.SurfaceTemperature < (FreezingPointOfWater - 3.0) {
		planet.Hydrosphere = 0.0
		planet.Albedo = planetEnvironment.planetAlbedo(planet.Hydrosphere, planet.CloudCover, planet.IceCover, planet.SurfacePressure)
		effectiveTemp := planetEnvironment.effectiveTemp(sun.EcosphereRadius, planet.SemiMajorAxis, planet.Albedo)
		opacity := planetEnvironment.opacity(planet.MolecularWeight, planet.SurfacePressure)
		greenhouseTemp := planetEnvironment.greenRise(opacity, effectiveTemp, planet.SurfacePressure)
		planet.SurfaceTemperature = effectiveTemp + greenhouseTemp
	}

	if !first {
		if !boilOff {
			planet.Hydrosphere = (planet.Hydrosphere + (lastWater * 2.0)) / 3.0
		}
		planet.CloudCover = (planet.CloudCover + (lastClouds * 2.0)) / 3.0
		planet.IceCover = (planet.IceCover + (lastIce * 2.0)) / 3.0
		planet.Albedo = (planet.Albedo + (lastAlbedo * 2.0)) / 3.0
		planet.SurfaceTemperature = (planet.SurfaceTemperature + (lastTemp * 2.0)) / 3.0
	}

	planetEnvironment.setPlanetTempRange(planet)
}

func (planetEnvironment *PlanetEnvironment) iterateAndSetSurfaceTemp(planet *Planet, sun *Sun) {
	initialTemp := planetEnvironment.estimatedTemp(sun.EcosphereRadius, planet.SemiMajorAxis, planet.Albedo)

	planetEnvironment.calculateAndSetSurfaceTemp(planet, sun, true, 0.0, 0.0, 0.0, 0.0, 0.0)

	for count := 0; count < 25; count++ {
		lastWater := planet.Hydrosphere
		lastClouds := planet.CloudCover
		lastIce := planet.IceCover
		lastTemp := planet.SurfaceTemperature
		lastAlbedo := planet.Albedo

		planetEnvironment.calculateAndSetSurfaceTemp(planet, sun, false, lastWater, lastClouds, lastIce, lastTemp, lastAlbedo)

		if math.Abs(planet.SurfaceTemperature-lastTemp) < 0.25 {
			break
		}
	}

	planet.GreenhouseRise = planet.SurfaceTemperature - initialTemp
}

func (planetEnvironment *PlanetEnvironment) calculateGases(planet *Planet, sun *Sun) []*Gas {
	atmosphere := make([]*Gas, 0)

	// no gases
	if !(planet.SurfacePressure > 0.0) {
		return atmosphere
	}

	totalAmount := 0.0
	amount := make(map[int64]float64)
	pressure := planet.SurfacePressure / MilliBarsPerBar
	for _, gas := range GasesTable {

		yp := gas.Boil / (373.0*(math.Log(pressure+0.001)/-5050.5) + (1.0 / 373.0))
		// check if the gas stay on the planet
		if !((yp >= 0.0 && yp < planet.LowTemperature) && gas.Weight >= planet.MolecularWeight) {
			continue
		}

		vrms := planetEnvironment.rootMeanSquareVelocity(gas.Weight, planet.ExosphericTemperature)
		pvrms := math.Pow(1.0/(1.0+vrms/planet.EscapeVelocity), sun.Age/1e9)

		abund, react := gas.Abunds, 1.0
		if gas.Symbol == "Ar" {
			react = 0.15 * sun.Age / 4e9
		} else if gas.Symbol == "He" {
			abund = abund * (0.001 + (planet.GasMass / planet.Mass))
			pres2 := (0.75 + pressure)
			react = math.Pow(1.0/(1.0+gas.Reactivity), sun.Age/2e9*pres2)
		} else if (gas.Symbol == "O" || gas.Symbol == "O2") && sun.Age > 2e9 && planet.SurfaceTemperature > 270.0 && planet.SurfaceTemperature < 400.0 {
			// 	pres2 = (0.65 + pressure/2); //Breathable - M: .55-1.4
			pres2 := (0.89 + pressure/4.0) // Breathable - M: .6 -1.8
			react = math.Pow(1.0/(1.0+gas.Reactivity), math.Pow(sun.Age/2e9, 0.25)*pres2)
		} else if gas.Symbol == "CO2" && sun.Age > 2e9 && planet.SurfaceTemperature > 270.0 && planet.SurfaceTemperature < 400.0 {
			pres2 := (0.75 + pressure)
			react = math.Pow(1.0/(1.0+gas.Reactivity), math.Pow(sun.Age/2e9, 0.5)*pres2)
			react *= 1.5
		} else {
			pres2 := (0.75 + pressure)
			react = math.Pow(1.0/(1.0+gas.Reactivity), sun.Age/2e9*pres2)
		}

		fract := (1.0 - (planet.MolecularWeight / gas.Weight))
		gasAmount := abund * pvrms * react * fract

		if gasAmount > 0.0 {
			totalAmount += gasAmount
			amount[gas.Num] = gasAmount
		}
	}

	for num, gasAmount := range amount {
		gas, ok := GasesTable[num]

		if !ok {
			continue
		}

		atmosphere = append(atmosphere, &Gas{
			Num:          gas.Num,
			SurfPressure: planet.SurfacePressure * gasAmount / totalAmount,
		})
	}

	// sort after SurfPressure
	sort.SliceStable(atmosphere, func(i, j int) bool {
		return atmosphere[i].SurfPressure < atmosphere[j].SurfPressure
	})

	return atmosphere
}

// inspiredPartialPressure taking into account humidification of the air in the nasal passage and throat
func (planetEnvironment *PlanetEnvironment) inspiredPartialPressure(surfPressure, gasPressure float64) float64 {
	// This formula is on Dole's p. 14
	pH20 := H2OAssumedPressure
	fraction := gasPressure / surfPressure
	return (surfPressure - pH20) * fraction
}

// breathability verifies if the planet has breathability gases
func (planetEnvironment *PlanetEnvironment) breathability(atmosphere []*Gas, surfPressure float64) Oxygen {
	/*--------------------------------------------------------------------------*/
	/*	 This function uses figures on the maximum inspired partial pressures   */
	/*   of Oxygen, other atmospheric and traces gases as laid out on pages 15, */
	/*   16 and 18 of Dole's Habitable Planets for Man to derive breathability  */
	/*   of the planet's atmosphere.                                       JLB  */
	/*--------------------------------------------------------------------------*/

	// no atmosphere
	if atmosphere == nil || len(atmosphere) < 1 {
		return None
	}

	oxygenOk := false

	// get each gas on the planet
	for _, gas := range atmosphere {

		ipp := planetEnvironment.inspiredPartialPressure(surfPressure, gas.SurfPressure)

		gasAtom, ok := GasesTable[gas.Num]
		if !ok {
			continue
		}

		if ipp > gasAtom.MaxIpp {
			return Toxic // POISONOUS
		}

		if gasAtom.Symbol == "O" {
			oxygenOk = (ipp >= MinO2IPP) && (ipp <= MaxO2IPP)
		}
	}

	if oxygenOk {
		return Breathable
	}
	return Unbreathable
}

func (planetEnvironment *PlanetEnvironment) GeneratePlanet(sun *Sun, planet *Planet, randomTilt bool, doMoons, doGases, isMoon bool) {
	// initialize the planet
	planet.Atmosphere = make([]*Gas, 0)
	planet.Breathability = None

	// start caluclating
	planet.OrbitZone = planetEnvironment.orbitZone(sun.Luminosity, planet.SemiMajorAxis)
	planet.OrbitPeriod = planetEnvironment.period(planet.SemiMajorAxis, planet.Mass, sun.Mass)
	if randomTilt {
		planet.AxialTilt = planetEnvironment.inclination(planet.SemiMajorAxis)
	}

	planet.ExosphericTemperature = EarthExosphereTemp / math.Pow(planet.SemiMajorAxis/sun.EcosphereRadius, 2.0)
	planet.RootMeanSquareVelocity = planetEnvironment.rootMeanSquareVelocity(MolNitrogen, planet.ExosphericTemperature)
	planet.CoreRadius = planetEnvironment.kothariRadius(planet.DustMass, false, planet.OrbitZone)

	// Calculate the radius as a gas giant, to verify it will retain gas.
	// Then if mass > Earth, it's at least 5% gas and retains He, it's
	// some flavor of gas giant.

	planet.Density = planetEnvironment.empiricalDensity(planet.Mass, planet.SemiMajorAxis, sun.EcosphereRadius, true)
	planet.Radius = planetEnvironment.volumeRadius(planet.Mass, planet.Density)

	planet.SurfaceAcceleration = planetEnvironment.acceleration(planet.Mass, planet.Radius)
	planet.SurfaceGravity = planetEnvironment.gravity(planet.SurfaceAcceleration)

	planet.MolecularWeight = planetEnvironment.minMolecularWeight(planet, sun)

	// check if the planet is a gas gaint
	if (planet.Mass*SunMassInEarthMasses) > 1.0 && (planet.GasMass/planet.Mass) > 0.05 && planet.MolecularWeight <= 4.0 {
		if (planet.GasMass / planet.Mass) < 0.20 {
			planet.Type = PlanetSubSubGasGiant
		} else if (planet.Mass * SunMassInEarthMasses) < 20.0 {
			planet.Type = PlanetSubGasGiant
		} else {
			planet.Type = PlanetGasGiant
		}
	} else { // If not, it's rocky.
		planet.Radius = planetEnvironment.kothariRadius(planet.Mass, false, planet.OrbitZone)
		planet.Density = planetEnvironment.volumeDensity(planet.Mass, planet.Radius)

		planet.SurfaceAcceleration = planetEnvironment.acceleration(planet.Mass, planet.Radius)
		planet.SurfaceGravity = planetEnvironment.gravity(planet.SurfaceAcceleration)

		if (planet.GasMass / planet.Mass) > 0.00001 {
			h2Mass := planet.GasMass * 0.85
			heMass := (planet.GasMass - h2Mass) * 0.999
			h2Loss, heLoss := 0.0, 0.0
			h2Life := planetEnvironment.gasLife(MolHydrogen, planet.ExosphericTemperature, planet.SurfaceGravity, planet.Radius)
			heLife := planetEnvironment.gasLife(Helium, planet.ExosphericTemperature, planet.SurfaceGravity, planet.Radius)

			if h2Life < sun.Age {
				h2Loss = ((1.0 - (1.0 / math.Exp(sun.Age/h2Life))) * h2Mass)
				planet.GasMass -= h2Loss
				planet.Mass -= h2Loss
			}

			if heLife < sun.Age {
				heLoss = ((1.0 - (1.0 / math.Exp(sun.Age/heLife))) * heMass)
				planet.GasMass -= heLoss
				planet.Mass -= heLoss

			}

			planet.SurfaceAcceleration = planetEnvironment.acceleration(planet.Mass, planet.Radius)
			planet.SurfaceGravity = planetEnvironment.gravity(planet.SurfaceAcceleration)
		}
	}

	planet.Day, planet.ResonantPeriod = planetEnvironment.dayLength(planet, sun)

	planet.EscapeVelocity = planetEnvironment.escapeVelocity(planet.Mass, planet.Radius)

	// check if the planet is a gas gaint
	if planet.Type == PlanetGasGiant || planet.Type == PlanetSubGasGiant || planet.Type == PlanetSubSubGasGiant {
		planet.GreenhouseEffect = false
		planet.VolatileGasInventory = math.MaxFloat64
		planet.SurfacePressure = math.MaxFloat64

		planet.BoilPoint = math.MaxFloat64

		planet.SurfaceTemperature = math.MaxFloat64
		planet.GreenhouseRise = 0.0
		planet.Albedo = planetEnvironment.about(GasGaintAlbedo, 0.1)
		planet.Hydrosphere = 1.0
		planet.CloudCover = 1.0
		planet.IceCover = 0.0
		planet.SurfaceGravity = planetEnvironment.gravity(planet.SurfaceAcceleration)
		planet.MolecularWeight = planetEnvironment.minMolecularWeight(planet, sun)
		planet.SurfaceGravity = math.MaxFloat64
		planet.EstimatedTemperature = planetEnvironment.estimatedTemp(sun.EcosphereRadius, planet.SemiMajorAxis, planet.Albedo)
		planet.EstimatedTerrestrialTemperature = planetEnvironment.estimatedTemp(sun.EcosphereRadius, planet.SemiMajorAxis, EarthAlbedo)

		temp := planet.EstimatedTerrestrialTemperature
		if temp >= FreezingPointOfWater && temp <= (EarthAverageKelvin+10.0) && sun.Age > 2.0e9 {
			//habitable_jovians
		}
	} else {
		planet.EstimatedTemperature = planetEnvironment.estimatedTemp(sun.EcosphereRadius, planet.SemiMajorAxis, EarthAlbedo)
		planet.EstimatedTerrestrialTemperature = planetEnvironment.estimatedTemp(sun.EcosphereRadius, planet.SemiMajorAxis, EarthAlbedo)

		planet.SurfaceGravity = planetEnvironment.gravity(planet.SurfaceAcceleration)
		planet.MolecularWeight = planetEnvironment.minMolecularWeight(planet, sun)

		planet.GreenhouseEffect = planetEnvironment.greenHouse(sun.EcosphereRadius, planet.SemiMajorAxis)

		accretedGas := (planet.GasMass / planet.Mass) > 0.000001
		planet.VolatileGasInventory = planetEnvironment.volInventory(planet.Mass, planet.EscapeVelocity, planet.RootMeanSquareVelocity,
			sun.Mass, planet.OrbitZone, planet.GreenhouseEffect, accretedGas)

		planet.SurfacePressure = planetEnvironment.pressure(planet.VolatileGasInventory, planet.Radius, planet.SurfaceGravity)

		if mathf.CloseZero(planet.SurfacePressure) {
			planet.BoilPoint = 0.0
		} else {
			planet.BoilPoint = planetEnvironment.boilingPoint(planet.SurfacePressure)
		}

		/*	Sets:
		 *		planet->surf_temp
		 *		planet->greenhs_rise
		 *		planet->albedo
		 *		planet->hydrosphere
		 *		planet->cloud_cover
		 *		planet->ice_cover
		 */
		planetEnvironment.iterateAndSetSurfaceTemp(planet, sun)

		if doGases && planet.MaxTemperature >= FreezingPointOfWater && planet.MinTemperature <= planet.BoilPoint {
			planet.Atmosphere = planetEnvironment.calculateGases(planet, sun)
		}

		/*
		 *	Next we assign a type to the planet.
		 */

		if planet.SurfacePressure < 1.0 {
			if !isMoon && (planet.Mass*SunMassInEarthMasses) < AsteroidMassLimit {
				planet.Type = PlanetAsteroids
			} else {
				planet.Type = PlanetRock
			}
		} else if planet.SurfacePressure > 6000.0 && planet.MolecularWeight <= 2.0 { // Retains Hydrogen
			planet.Type = PlanetSubSubGasGiant
			planet.Atmosphere = make([]*Gas, 0)
		} else {
			// Atmospheres:
			if mathf.CloseEquals(planet.Day, planet.OrbitPeriod*24.0) || planet.ResonantPeriod {
				planet.Type = Planet1Face
			} else if planet.Hydrosphere >= 0.95 {
				planet.Type = PlanetWater // >95% water
			} else if planet.IceCover >= 0.95 {
				planet.Type = PlanetIce // >95% ice
			} else if planet.Hydrosphere > 0.05 {
				planet.Type = PlanetTerrestrial // Terrestrial else <5% water
			} else if planet.MaxTemperature > planet.BoilPoint {
				planet.Type = PlanetVenusian // Hot = Venusian
			} else if (planet.GasMass / planet.Mass) > 0.0001 {
				planet.Type = PlanetIce // Accreted gas, But no Greenhouse or liquid water => make it an Ice World
				planet.IceCover = 1.0
			} else if planet.SurfacePressure <= 250.0 {
				planet.Type = PlanetMartian // Thin air = Martian
			} else if planet.SurfaceTemperature < FreezingPointOfWater {
				planet.Type = PlanetIce
			} else {
				planet.Type = PlanetUnknown
			}
		}
	}

	// verify if the planet Breathability
	planet.Breathability = planetEnvironment.breathability(planet.Atmosphere, planet.SurfaceAcceleration)

	// check if do moons and a moon is available
	if doMoons && !isMoon && planet.Moons != nil && len(planet.Moons) > 0 {
		// generate all moons for the planet
		for _, moon := range planet.Moons {
			// move the moon to the planet
			moon.SemiMajorAxis = planet.SemiMajorAxis
			moon.Eccentricity = planet.Eccentricity

			// generate the moon
			planetEnvironment.GeneratePlanet(sun, moon, randomTilt, doMoons, doGases, true) // Adjusts ptr->density

			rocheLimit := 2.44 * planet.Radius * math.Pow((planet.Density/moon.Density), (1.0/3.0))
			hillSphere := planet.SemiMajorAxis * KMPerAU * math.Pow((planet.Mass/(3.0*sun.Mass)), (1.0/3.0))
			if (rocheLimit * 3.0) < hillSphere {
				moon.MoonA = utils.RandFloat64(rocheLimit*1.5, hillSphere/2.0) / KMPerAU
				moon.MoonE = randEccentricity()
			} else {
				moon.MoonA = 0.0
				moon.MoonE = 0.0
			}
		}
	}

}