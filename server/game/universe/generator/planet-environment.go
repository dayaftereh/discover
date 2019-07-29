package generator

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

func (planetEnvironment *PlanetEnvironment) about(value float64, variation float64) float64 {
	return (value + (value * utils.RandFloat64(-variation, variation)))
}

func (planetEnvironment *PlanetEnvironment) inclination(orbitRadius float64) float64 {
	temp := int64(math.Round(math.Pow(orbitRadius, 0.2) * planetEnvironment.about(EarthAxialTilt, 0.4)))
	return (math.Mod(float64(temp), 360.0))
}

func (planetEnvironment *PlanetEnvironment) period(separation, smallMass, largeMass float64) float64 {
	eriodInYears := math.Sqrt((separation * separation * separation) / (smallMass + largeMass))
	return (eriodInYears * DaysInAYear)
}

func (planetEnvironment *PlanetEnvironment) rmsVelocity(molecularWeight, exosphericTemp float64) float64 {
	return math.Sqrt((3.0*MolarGasConst*exosphericTemp)/molecularWeight) * CMPerMeter
}

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

func (planetEnvironment *PlanetEnvironment) empiricalDensity(mass float64, orbitRadius float64, rEcosphere float64, gasGiant bool) float64 {
	temp := math.Pow(mass*SunMassInEarthMasses, (1.0 / 8.0))
	temp = temp * math.Sqrt(math.Sqrt(rEcosphere/orbitRadius))
	if gasGiant {
		return temp * 1.2
	}
	return temp * 5.5
}

func (planetEnvironment *PlanetEnvironment) volumeRadius(mass float64, density float64) float64 {
	mass = mass * SolarMassInGrams
	volume := mass / density
	return math.Pow((3.0*volume)/(4.0*math.Pi), (1.0/3.0)) / CMPerKM
}

func (planetEnvironment *PlanetEnvironment) volumeDensity(mass float64, radius float64) float64 {
	mass = mass * SolarMassInGrams
	radiusInCm := radius * CMPerKM
	volume := (4.0 * math.Pi * math.Pow(radiusInCm, 3.0)) / 3.0
	return mass / volume
}

func (planetEnvironment *PlanetEnvironment) dayLength(planetType PlanetType, mass, radius, orbitPeriod, density, a, e, sunMass, sunAge float64) (float64, bool) {
	planetaryMassInGrams := mass * SolarMassInGrams
	equatorialRadiusInCM := radius * CMPerKM
	yearInHours := orbitPeriod * 24.0
	giant := (planetType == PlanetGasGiant || planetType == PlanetSubGasGiant || planetType == PlanetSubSubGasGiant)

	stopped := false

	k2 := 0.33
	if giant {
		k2 = 0.24
	}

	baseAngularVelocity := math.Sqrt(2.0 * J * planetaryMassInGrams / (k2 * math.Pow(equatorialRadiusInCM, 2.0)))
	changeInAngularVelocity := ChangeInErthAngVel * (density / EarthDensity) * (equatorialRadiusInCM / EarthRadius) *
		(EarthMassInGrams / planetaryMassInGrams) * math.Pow(sunMass, 2.0) * (1.0 / math.Pow(a, 6.0))
	angVelocity := baseAngularVelocity + (changeInAngularVelocity * sunAge)

	var daysInAYear float64
	if angVelocity <= 0.0 {
		stopped = true
		daysInAYear = math.MaxFloat64
	} else {
		daysInAYear = (2.0 * math.Pi) / (SecondsPerHour * angVelocity)
	}

	if daysInAYear > yearInHours || stopped {
		if e > 0.1 {
			spinResonanceFactor := (1.0 - e) / (1.0 + e)
			return (spinResonanceFactor * yearInHours), true
		}
		return yearInHours, false
	}
	return daysInAYear, false
}

func (planetEnvironment *PlanetEnvironment) gravity(acceleration float64) float64 {
	return acceleration / EarthAcceleration
}

func (planetEnvironment *PlanetEnvironment) acceleration(mass float64, radius float64) float64 {
	return GravConstant * (mass * SolarMassInGrams) / math.Pow(radius*CMPerKM, 2.0)
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

func (planetEnvironment *PlanetEnvironment) gasLife(molecularWeight float64, exosphericTemp float64, surfGrav float64, radius float64) float64 {
	v := planetEnvironment.rmsVelocity(molecularWeight, exosphericTemp)
	g := surfGrav * EarthAcceleration
	r := radius * CMPerKM
	t := (math.Pow(v, 3.0) / (2.0 * math.Pow(g, 2.0) * r)) * math.Exp((3.0*g*r)/math.Pow(v, 2.0))
	years := t / (SecondsPerHour * 24.0 * DaysInAYear)

	if years > 2.0e10 {
		return math.MaxFloat64
	}
	return years
}

func (planetEnvironment *PlanetEnvironment) escapeVelocity(mass float64, radius float64) float64 {
	massInGrams := mass * SolarMassInGrams
	radiusInCm := radius * CMPerKM
	return math.Sqrt(2.0 * GravConstant * massInGrams / radiusInCm)
}

func (planetEnvironment *PlanetEnvironment) moleculeLimit(mass float64, equatRadius float64, exosphericTemp float64) float64 {
	escVelocity := planetEnvironment.escapeVelocity(mass, equatRadius)
	return (3.0 * MolarGasConst * exosphericTemp) / (math.Pow((escVelocity/GasRetentionThreshold)/CMPerMeter, 2.0))

}

func (planetEnvironment *PlanetEnvironment) effTemp(ecosphereRadius float64, orbitRadius float64, albedo float64) float64 {
	return math.Sqrt(ecosphereRadius/orbitRadius) * math.Sqrt(math.Sqrt((1.0-albedo)/(1.0-EarthAlbedo))) * EarthEffectiveTemp
}

func (planetEnvironment *PlanetEnvironment) estTemp(ecosphereRadius float64, orbitRadius float64, albedo float64) float64 {
	return math.Sqrt(ecosphereRadius/orbitRadius) * math.Sqrt(math.Sqrt((1.0-albedo)/(1.0-EarthAlbedo))) * EarthAverageKelvin
}

func (planetEnvironment *PlanetEnvironment) minMolecWeigth(mass float64, radius float64, exosphericTemp float64, sunAge float64, surfGrav float64) float64 {
	tagret := sunAge

	guess1 := planetEnvironment.moleculeLimit(mass, radius, exosphericTemp)
	guess2 := guess1

	life := planetEnvironment.gasLife(guess1, exosphericTemp, surfGrav, radius)

	if life > tagret {
		for loops := 0; loops < 25 && life > tagret; loops++ {
			guess1 = guess1 / 2.0
			life = planetEnvironment.gasLife(guess1, exosphericTemp, surfGrav, radius)
		}
	} else {
		for loops := 0; loops < 25 && life < tagret; loops++ {
			guess1 = guess1 * 2.0
			life = planetEnvironment.gasLife(guess1, exosphericTemp, surfGrav, radius)
		}
	}

	for loops := 0; loops < 25 && (guess2-guess1) > 0.1; loops++ {
		guess3 := (guess1 + guess2) / 2.0
		life = planetEnvironment.gasLife(guess3, exosphericTemp, surfGrav, radius)
		if life < tagret {
			guess1 = guess3
		} else {
			guess2 = guess3
		}
	}

	life = planetEnvironment.gasLife(guess2, exosphericTemp, surfGrav, radius)

	return guess2
}

func (planetEnvironment *PlanetEnvironment) greenHouse(rEcosphere, orbitRadius float64) bool {
	temp := planetEnvironment.effTemp(rEcosphere, orbitRadius, GreenhouseTriggerAlbedo)
	if temp > FreeezingPointOfWater {
		return true
	}
	return false
}

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

func (planetEnvironment *PlanetEnvironment) pressure(volatileGasInventory, equatRadius, gravity float64) float64 {
	equatRadius = EarthRadiusInKM / equatRadius
	return volatileGasInventory * gravity * (EarthSurfPersInMilliBars / MilliBarsPerBar) / math.Pow(equatRadius, 2.0)
}

func (planetEnvironment *PlanetEnvironment) boilingPoint(surfPressure float64) float64 {
	surfacePressureInBars := surfPressure / MilliBarsPerBar
	return 1.0 / ((math.Log(surfacePressureInBars) / -5050.5) + (1.0 / 373.0))
}

func (planetEnvironment *PlanetEnvironment) lim(x float64) float64 {
	return x / math.Sqrt(1.0+x*x*x*x)
}

func (planetEnvironment *PlanetEnvironment) soft(v, max, min float64) float64 {
	dv := v - min
	dm := max - min
	return (planetEnvironment.lim(2.0*dv/dm-1.0)+1.0)/2.0*dm + min
}

func (planetEnvironment *PlanetEnvironment) greenRise(opticalDepth, effectiveTemp, surfPressure float64) float64 {
	convectionFactor := EarthConvectionFactor * math.Pow(surfPressure/EarthSurfPersInMilliBars, 0.4)
	rise := (math.Sqrt(math.Sqrt(1.0+0.75*opticalDepth)) - 1.0) * effectiveTemp * convectionFactor
	if rise < 0.0 {
		return 0.0
	}
	return rise
}

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
	pressmod := 1.0 / math.Sqrt(1.0+20.0*planet.SurfPressure/1000.0)
	ppmod := 1.0 / math.Sqrt(10.0+5.0*planet.SurfPressure/1000.0)
	tiltmod := math.Abs(math.Cos(planet.AxialTilt*math.Pi/180.0) * math.Pow(1.0+planet.E, 2.0))
	daymod := 1.0 / (200.0/planet.Day + 1.0)
	mh := math.Pow(1.0+daymod, pressmod)
	ml := math.Pow(1.0-daymod, pressmod)
	hi := mh * planet.SurfTemp
	lo := ml * planet.SurfTemp
	sh := hi + math.Pow((100.0+hi)*tiltmod, math.Sqrt(ppmod))
	wl := lo - math.Pow((150.0+lo)*tiltmod, math.Sqrt(ppmod))
	max := planet.SurfTemp + math.Sqrt(planet.SurfTemp)*10.0
	min := planet.SurfTemp / math.Sqrt(planet.Day+24.0)

	if lo < min {
		lo = min
	}
	if wl < 0 {
		wl = 0
	}

	planet.HighTemp = planetEnvironment.soft(hi, max, min)
	planet.LowTemp = planetEnvironment.soft(lo, max, min)
	planet.MaxTemp = planetEnvironment.soft(sh, max, min)
	planet.MinTemp = planetEnvironment.soft(wl, max, min)
}

func (planetEnvironment *PlanetEnvironment) hydroFraction(volatileGasInventory, planetRadius float64) float64 {
	temp := (0.71 * volatileGasInventory / 1000.0) * math.Pow(EarthRadiusInKM/planetRadius, 2.0)
	return math.Min(temp, 1.0)
}

func (planetEnvironment *PlanetEnvironment) cloudFraction(surfTemp, smallestMWRetained, equatRadius, hydroFraction float64) float64 {
	if smallestMWRetained > WaterVapro {
		return 0.0
	}

	surfArea := 4.0 * math.Pi * math.Pow(equatRadius, 2.0)
	hydroMass := hydroFraction * surfArea * EarthWaterMassPerArea
	waterVaporInKG := (0.00000001 * hydroMass) * math.Exp(Q2_36*(surfTemp-EarthAverageKelvin))
	fraction := CloudCoverageFactor * waterVaporInKG / surfArea

	return math.Min(fraction, 1.0)
}
func (planetEnvironment *PlanetEnvironment) iceFraction(hydroFraction, surfTemp float64) float64 {
	surfTemp = math.Min(surfTemp, 328.0)
	temp := math.Pow(((328.0 - surfTemp) / 90.0), 5.0)
	temp = math.Min(temp, (1.5 * hydroFraction))
	return math.Min(temp, 1.0)
}

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

	cloudPart := cloudFraction * CloudAlbedo

	rockPart := rockFraction * RockyAlbedo
	icePart := iceFraction * IceAlbedo
	waterPart := waterFraction * WaterAlbedo
	if surfPressure == 0.0 {
		waterPart = 0.0
		rockPart = rockFraction * RockyAirlessAlbedo
		icePart = iceFraction * AirlessIceAlbedo
	}

	return cloudPart + rockPart + waterPart + icePart
}

func (planetEnvironment *PlanetEnvironment) calculateAndSetSurfaceTemp(planet *Planet, first bool, rEcosphere, lastWater, lastClouds, lastIce, lastTemp, lastAlbedo, sunMass float64) {
	if first {
		planet.Albedo = EarthAlbedo

		effectiveTemp := planetEnvironment.effTemp(rEcosphere, planet.A, planet.Albedo)
		opacity := planetEnvironment.opacity(planet.MolecWeight, planet.SurfPressure)
		greenhouseTemp := planetEnvironment.greenRise(opacity, effectiveTemp, planet.SurfPressure)
		planet.SurfTemp = effectiveTemp + greenhouseTemp

		planetEnvironment.setPlanetTempRange(planet)
	}

	if planet.GreenhouseEffect && planet.MaxTemp < planet.BoilPoint {
		planet.GreenhouseEffect = false
		accretedGas := (planet.GasMass / planet.Mass) < 0.000001
		planet.VolatileGasInventory = planetEnvironment.volInventory(
			planet.Mass,
			planet.ESCVelocity,
			planet.RMSVelocity,
			sunMass, planet.OrbitZone,
			planet.GreenhouseEffect,
			accretedGas)
		planet.SurfPressure = planetEnvironment.pressure(planet.VolatileGasInventory, planet.Radius, planet.SurfGrav)
		planet.BoilPoint = planetEnvironment.boilingPoint(planet.SurfPressure)
	}

	planet.Hydrosphere = planetEnvironment.hydroFraction(planet.VolatileGasInventory, planet.Radius)
	planet.CloudCover = planetEnvironment.cloudFraction(planet.SurfTemp, planet.MolecWeight, planet.Radius, planet.Hydrosphere)
	planet.IceCover = planetEnvironment.iceFraction(planet.Hydrosphere, planet.SurfTemp)

	if planet.GreenhouseEffect && planet.SurfPressure > 0.0 {
		planet.CloudCover = 1.0
	}

	boilOff := false
	if planet.HighTemp >= planet.BoilPoint && !first && !(mathf.CloseEquals(planet.Day, planet.OrbPeriod*24.0) || planet.ResonantPeriod) {
		boilOff = true
		planet.Hydrosphere = 0.0
		if planet.MolecWeight > WaterVapro {
			planet.CloudCover = 0.0
		} else {
			planet.CloudCover = 1.0
		}
	}

	if planet.SurfTemp < (FreeezingPointOfWater - 3.0) {
		planet.Hydrosphere = 0.0
		planet.Albedo = planetEnvironment.planetAlbedo(planet.Hydrosphere, planet.CloudCover, planet.IceCover, planet.SurfPressure)
		effectiveTemp := planetEnvironment.effTemp(rEcosphere, planet.A, planet.Albedo)
		opacity := planetEnvironment.opacity(planet.MolecWeight, planet.SurfPressure)
		greenhouseTemp := planetEnvironment.greenRise(opacity, effectiveTemp, planet.SurfPressure)
		planet.SurfTemp = effectiveTemp + greenhouseTemp
	}

	if !first {
		if !boilOff {
			planet.Hydrosphere = (planet.Hydrosphere + (lastWater * 2.0)) / 3.0
		}
		planet.CloudCover = (planet.CloudCover + (lastClouds * 2)) / 3
		planet.IceCover = (planet.IceCover + (lastIce * 2)) / 3
		planet.Albedo = (planet.Albedo + (lastAlbedo * 2)) / 3
		planet.SurfTemp = (planet.SurfTemp + (lastTemp * 2)) / 3
	}

	planetEnvironment.setPlanetTempRange(planet)
}

func (planetEnvironment *PlanetEnvironment) iterateAndSetSurfaceTemp(planet *Planet, sunMass, rEcosphere float64) {
	initialTemp := planetEnvironment.estTemp(rEcosphere, planet.A, planet.Albedo)

	planetEnvironment.calculateAndSetSurfaceTemp(planet, true, rEcosphere, 0.0, 0.0, 0.0, 0.0, 0.0, sunMass)

	for count := 0; count < 25; count++ {
		lastWater := planet.Hydrosphere
		lastClouds := planet.CloudCover
		lastIce := planet.IceCover
		lastTemp := planet.SurfTemp
		lastAlbedo := planet.Albedo

		planetEnvironment.calculateAndSetSurfaceTemp(planet, false, rEcosphere, lastWater, lastClouds, lastIce, lastTemp, lastAlbedo, sunMass)

		if math.Abs(planet.SurfTemp-lastTemp) < 0.25 {
			break
		}
	}

	planet.GreenhsRise = planet.SurfTemp - initialTemp
}

func (planetEnvironment *PlanetEnvironment) calculateGases(sunAge, molecWeight, escVelocity, gasMass, mass, surfPressure, surfTemp, exosphericTemp, lowTemp float64) []*Gas {
	atmosphere := make([]*Gas, 0)

	// no gases
	if !(surfPressure > 0.0) {
		return atmosphere
	}

	totalAmount := 0.0
	amount := make(map[int64]float64)
	pressure := surfPressure / MilliBarsPerBar
	for _, gas := range GasesTable {
		yp := gas.Boil / (373.0*((math.Log(pressure)+0.001)/-5050.5) + (10 / 373.0))
		// check if the gas stay on the planet
		if !((yp >= 0.0 && yp < lowTemp) && gas.Weigth >= molecWeight) {
			continue
		}

		vrms := planetEnvironment.rmsVelocity(gas.Weigth, exosphericTemp)
		pvrms := math.Pow(1.0/(1.0+vrms/escVelocity), sunAge/1e9)

		abund, react := gas.Abunds, 1.0
		if gas.Symbol == "Ar" {
			react = 0.15 * sunAge / 4e9
		} else if gas.Symbol == "He" {
			abund = abund * (0.001 + (gasMass / mass))
			pres2 := (0.75 + pressure)
			react = math.Pow(1.0/(1.0+gas.Reactivity), sunAge/2e9*pres2)
		} else if (gas.Symbol == "O" || gas.Symbol == "O2") && sunAge > 2e9 && surfTemp > 270.0 && surfTemp < 400.0 {
			// 	pres2 = (0.65 + pressure/2); //Breathable - M: .55-1.4
			pres2 := (0.89 + pressure/4.0) // Breathable - M: .6 -1.8
			react = math.Pow(1.0/(1.0+gas.Reactivity), math.Pow(sunAge/2e9, 0.25)*pres2)
		} else if gas.Symbol == "CO2" && sunAge > 2e9 && surfTemp > 270.0 && surfTemp < 400.0 {
			pres2 := (0.75 + pressure)
			react = math.Pow(1.0/(1.0+gas.Reactivity), math.Pow(sunAge/2e9, 0.5)*pres2)
			react *= 1.5
		} else {
			pres2 := (0.75 + pressure)
			react = math.Pow(1.0/(1.0+gas.Reactivity), sunAge/2e9*pres2)
		}

		fract := (1.0 - (molecWeight / gas.Weigth))
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
			SurfPressure: surfPressure * gasAmount / totalAmount,
		})
	}

	// sort after SurfPressure
	sort.SliceStable(atmosphere, func(i, j int) bool {
		return atmosphere[i].SurfPressure < atmosphere[j].SurfPressure
	})

	return atmosphere
}

func (planetEnvironment *PlanetEnvironment) inspiredPartialPressure(gasPressure, surfPressure float64) float64 {
	pH20 := H2OAssumedPressure
	fraction := gasPressure / surfPressure
	return (surfPressure - pH20) * fraction
}

func (planetEnvironment *PlanetEnvironment) breathability(atmosphere []*Gas, surfPressure float64) Oxygen {
	// no atmosphere
	if atmosphere == nil || len(atmosphere) < 1 {
		return None
	}

	oxygenOk := false

	for _, gas := range atmosphere {

		ipp := planetEnvironment.inspiredPartialPressure(gas.SurfPressure, surfPressure)

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

	planet.OrbitZone = planetEnvironment.orbitZone(sun.Luminosity, planet.A)
	planet.OrbPeriod = planetEnvironment.period(planet.A, planet.Mass, sun.Mass)
	if randomTilt {
		planet.AxialTilt = planetEnvironment.inclination(planet.A)
	}

	planet.ExosphericTemp = EarthExosphereTemp / math.Pow(planet.A/sun.REcosphere, 2.0)
	planet.RMSVelocity = planetEnvironment.rmsVelocity(MolNitrogen, planet.ExosphericTemp)
	planet.CoreRadius = planetEnvironment.kothariRadius(planet.DustMass, false, planet.OrbitZone)

	// Calculate the radius as a gas giant, to verify it will retain gas.
	// Then if mass > Earth, it's at least 5% gas and retains He, it's
	// some flavor of gas giant.

	planet.Density = planetEnvironment.empiricalDensity(planet.Mass, planet.A, sun.REcosphere, true)
	planet.Radius = planetEnvironment.volumeRadius(planet.Mass, planet.Density)

	planet.SurfAccel = planetEnvironment.acceleration(planet.Mass, planet.Radius)
	planet.SurfGrav = planetEnvironment.gravity(planet.SurfAccel)

	planet.MolecWeight = planetEnvironment.minMolecWeigth(planet.Mass, planet.Radius, planet.ExosphericTemp, sun.Age, planet.SurfGrav)

	// check if the playnet is a gas gaint
	if (planet.Mass*SunMassInEarthMasses) > 1.0 && (planet.GasMass/planet.Mass) > 0.05 && planet.MolecWeight <= 4.0 {
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

		planet.SurfAccel = planetEnvironment.acceleration(planet.Mass, planet.Radius)
		planet.SurfGrav = planetEnvironment.gravity(planet.SurfAccel)

		if (planet.GasMass / planet.Mass) > 0.00001 {
			h2Mass := planet.GasMass * 0.85
			heMass := (planet.GasMass - h2Mass) * 0.999
			h2Loss, heLoss := 0.0, 0.0
			h2Life := planetEnvironment.gasLife(MolHydrogen, planet.ExosphericTemp, planet.SurfGrav, planet.Radius)
			heLife := planetEnvironment.gasLife(Helium, planet.ExosphericTemp, planet.SurfGrav, planet.Radius)

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

			planet.SurfAccel = planetEnvironment.acceleration(planet.Mass, planet.Radius)
			planet.SurfGrav = planetEnvironment.gravity(planet.SurfAccel)
		}
	}

	planet.Day, planet.ResonantPeriod = planetEnvironment.dayLength(planet.Type, planet.Mass, planet.Radius, planet.OrbPeriod,
		planet.Density, planet.A, planet.E, sun.Mass, sun.Age)

	planet.ESCVelocity = planetEnvironment.escapeVelocity(planet.Mass, planet.Radius)

	// check if the planet is a gas gaint
	if planet.Type == PlanetGasGiant || planet.Type == PlanetSubGasGiant || planet.Type == PlanetSubSubGasGiant {
		planet.GreenhouseEffect = false
		planet.VolatileGasInventory = math.MaxFloat64
		planet.SurfPressure = math.MaxFloat64

		planet.BoilPoint = math.MaxFloat64

		planet.SurfTemp = math.MaxFloat64
		planet.GreenhsRise = 0.0
		planet.Albedo = planetEnvironment.about(GasGaintAlbedo, 0.1)
		planet.Hydrosphere = 1.0
		planet.CloudCover = 1.0
		planet.IceCover = 1.0
		planet.SurfGrav = planetEnvironment.gravity(planet.SurfAccel)
		planet.MolecWeight = planetEnvironment.minMolecWeigth(planet.Mass, planet.Radius, planet.ExosphericTemp, sun.Age, planet.SurfGrav)
		planet.SurfGrav = math.MaxFloat64
		planet.EstimatedTemp = planetEnvironment.estTemp(sun.REcosphere, planet.A, planet.Albedo)
		planet.EstimatedTerrTemp = planetEnvironment.estTemp(sun.REcosphere, planet.A, EarthAlbedo)

		temp := planet.EstimatedTerrTemp
		if temp >= FreeezingPointOfWater && temp <= (EarthAverageKelvin+10.0) && sun.Age > 2.0e9 {
			//habitable_jovians
		}
	} else {
		planet.EstimatedTemp = planetEnvironment.estTemp(sun.REcosphere, planet.A, EarthAlbedo)
		planet.EstimatedTerrTemp = planetEnvironment.estTemp(sun.REcosphere, planet.A, EarthAlbedo)

		planet.SurfGrav = planetEnvironment.gravity(planet.SurfAccel)
		planet.MolecWeight = planetEnvironment.minMolecWeigth(planet.Mass, planet.Radius, planet.ExosphericTemp, sun.Age, planet.SurfGrav)

		planet.GreenhouseEffect = planetEnvironment.greenHouse(sun.REcosphere, planet.A)

		accretedGas := (planet.GasMass / planet.Mass) > 0.000001
		planet.VolatileGasInventory = planetEnvironment.volInventory(planet.Mass, planet.ESCVelocity, planet.RMSVelocity, sun.Mass,
			planet.OrbitZone, planet.GreenhouseEffect, accretedGas)

		planet.SurfPressure = planetEnvironment.pressure(planet.VolatileGasInventory, planet.Radius, planet.SurfGrav)

		if mathf.CloseZero(planet.SurfPressure) {
			planet.BoilPoint = 0.0
		} else {
			planet.BoilPoint = planetEnvironment.boilingPoint(planet.SurfPressure)
		}

		/*	Sets:
		 *		planet->surf_temp
		 *		planet->greenhs_rise
		 *		planet->albedo
		 *		planet->hydrosphere
		 *		planet->cloud_cover
		 *		planet->ice_cover
		 */
		planetEnvironment.iterateAndSetSurfaceTemp(planet, sun.Mass, sun.REcosphere)

		if doGases && planet.MaxTemp >= FreeezingPointOfWater && planet.MinTemp <= planet.BoilPoint {
			planet.Atmosphere = planetEnvironment.calculateGases(sun.Age, planet.MolecWeight, planet.ESCVelocity, planet.GasMass,
				planet.Mass, planet.SurfPressure, planet.SurfTemp, planet.ExosphericTemp, planet.LowTemp)
		}

		/*
		 *	Next we assign a type to the planet.
		 */

		if planet.SurfPressure < 1.0 {
			if !isMoon && (planet.Mass*SunMassInEarthMasses) < AsteroidMassLimit {
				planet.Type = PlanetAsteroids
			} else {
				planet.Type = PlanetRock
			}
		} else if planet.SurfPressure > 6000.0 && planet.MolecWeight <= 2.0 { // Retains Hydrogen
			planet.Type = PlanetSubSubGasGiant
			planet.Atmosphere = make([]*Gas, 0)
		} else {
			// Atmospheres:
			if mathf.CloseEquals(planet.Day, planet.OrbPeriod*24.0) || planet.ResonantPeriod {
				planet.Type = Planet1Face
			} else if planet.Hydrosphere >= 0.95 {
				planet.Type = PlanetWater // >95% water
			} else if planet.IceCover >= 0.95 {
				planet.Type = PlanetIce // >95% ice
			} else if planet.Hydrosphere > 0.05 {
				planet.Type = PlanetTerrestrial // Terrestrial else <5% water
			} else if planet.MaxTemp > planet.BoilPoint {
				planet.Type = PlanetVenusian // Hot = Venusian
			} else if (planet.GasMass / planet.Mass) > 0.0001 {
				planet.Type = PlanetIce // Accreted gas, But no Greenhouse or liquid water => make it an Ice World
				planet.IceCover = 1.0
			} else if planet.SurfPressure <= 250.0 {
				planet.Type = PlanetMartian // Thin air = Martian
			} else if planet.SurfTemp < FreeezingPointOfWater {
				planet.Type = PlanetIce
			} else {
				planet.Type = PlanetUnknown
			}
		}

		if doMoons && !isMoon {
			if planet.Moons == nil || len(planet.Moons) < 1 {
				return
			}

			for _, moon := range planet.Moons {
				planetEnvironment.GeneratePlanet(sun, moon, randomTilt, doMoons, doGases, true) // Adjusts ptr->density

				rocheLimit := 2.44 * planet.Radius * math.Pow((planet.Density/moon.Density), (1.0/3.0))
				hillSphere := planet.A * KMPerAU * math.Pow((planet.Mass/(3.0*sun.Mass)), (1.0/3.0))
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
}
