package generator

import (
	"math"

	"github.com/dayaftereh/discover/server/utils"
)

type AccretionProcessor struct {
	CloudEccentricity float64
	DustBand          *DustBand
	Planet            *Planet
}

func NewAccretionProcessor() *AccretionProcessor {
	return &AccretionProcessor{
		CloudEccentricity: 0.2,
		Planet:            nil,
		DustBand:          nil,
	}
}

func (accretionProcessor *AccretionProcessor) StellarDustLimit(stellMassRatio float64) float64 {
	return (200.0 * math.Pow(stellMassRatio, (1.0/3.0)))
}

func (accretionProcessor *AccretionProcessor) InitialDustBand(innerDust float64, outerDust float64) {
	accretionProcessor.DustBand = &DustBand{
		InnerEdge:   innerDust,
		OuterEdge:   outerDust,
		DustPresent: true,
		GasPresent:  true,
	}
}

func (accretionProcessor *AccretionProcessor) nearestPlanet(stellMassRatio float64) float64 {
	return (0.3 * math.Pow(stellMassRatio, (1.0/3.0)))
}

func (accretionProcessor *AccretionProcessor) farthestPlanet(stellMassRatio float64) float64 {
	return (50.0 * math.Pow(stellMassRatio, (1.0/3.0)))
}

func (accretionProcessor *AccretionProcessor) randEccentricity() float64 {
	E := 1.0 - math.Pow(utils.RandFloat64(0.0, 1.0), EccentricityCoeff)
	if E > 0.9999 {
		E = 0.999
	}
	return E
}

func (accretionProcessor *AccretionProcessor) innerEffectLimit(a, e, mass float64) float64 {
	return (a * (1.0 - e) * (1.0 - mass) / (1.0 + accretionProcessor.CloudEccentricity))

}

func (accretionProcessor *AccretionProcessor) outerEffectLimit(a, e, mass float64) float64 {
	return (a * (1.0 + e) * (1.0 + mass) / (1.0 - accretionProcessor.CloudEccentricity))

}

func (accretionProcessor *AccretionProcessor) isDustAvailable(insideRange, outsideRange float64) bool {

	dustBand := accretionProcessor.DustBand
	dustHere := false

	for dustBand != nil && dustBand.OuterEdge < insideRange {
		dustBand = dustBand.NextBand
	}

	// no dust band found
	if dustBand == nil {
		return false
	}

	dustHere = dustBand.DustPresent

	for dustBand != nil && dustBand.InnerEdge < outsideRange {
		dustHere = dustHere || dustBand.DustPresent
		dustBand = dustBand.NextBand
	}
	return dustHere
}

func (accretionProcessor *AccretionProcessor) criticalLimit(orbRadius, eccentricity, stellLuminosityRatio float64) float64 {
	perihelionDist := (orbRadius - orbRadius*eccentricity)
	temp := perihelionDist * math.Sqrt(stellLuminosityRatio)
	return (B * math.Pow(temp, -0.75))
}

func (accretionProcessor *AccretionProcessor) collectDust(lastMass, a, e, critMass, dustDensity float64, dustBand *DustBand) (float64, float64, float64, float64, float64, float64) {
	temp := lastMass / (1.0 + lastMass)
	reducedMass := math.Pow(temp, (1.0 / 4.0))

	rInner := accretionProcessor.innerEffectLimit(a, e, reducedMass)
	rOuter := accretionProcessor.outerEffectLimit(a, e, reducedMass)

	if rInner < 0.0 {
		rInner = 0.0
	}

	if dustBand == nil {
		return 0.0, 0.0, 0.0, rInner, rOuter, reducedMass
	}

	tempDensity := dustDensity
	if !dustBand.DustPresent {
		tempDensity = 0.0
	}

	gasDensity := 0.0
	massDensity := tempDensity
	if !(lastMass < critMass || !dustBand.GasPresent) {
		massDensity = K * tempDensity / (1.0 + math.Sqrt(critMass/lastMass)*(K-1.0))
		gasDensity = massDensity - tempDensity
	}

	if dustBand.OuterEdge <= rInner || dustBand.InnerEdge >= rOuter {
		return accretionProcessor.collectDust(lastMass, a, e, critMass, dustDensity, dustBand.NextBand)
	}

	bandwidth := (rOuter - rInner)
	temp1 := rOuter - dustBand.OuterEdge
	if temp1 < 0.0 {
		temp1 = 0.0
	}

	width := bandwidth - temp1

	temp2 := dustBand.InnerEdge - rInner
	if temp2 < 0.0 {
		temp2 = 0.0
	}

	width = width - temp2

	temp = 4.0 * math.Pi * math.Pow(a, 2.0) * reducedMass * (1.0 - e*(temp1-temp2)/bandwidth)
	volume := temp * width

	newMass := (volume * massDensity)

	newGas := (volume * gasDensity)
	newDust := (newMass - newGas)

	nextMass, nextDust, nextGas, rInner, rOuter, reducedMass := accretionProcessor.collectDust(lastMass, a, e, critMass, dustDensity, dustBand.NextBand)

	newMass += nextMass
	newGas += nextGas
	newDust += nextDust

	return newMass, newDust, newGas, rInner, rOuter, reducedMass
}

func (accretionProcessor *AccretionProcessor) updateDustLanes(min, max, mass, critMass, bodyInnerBound, bodyOuterBound float64) bool {
	gas := !(mass > critMass)

	node1 := accretionProcessor.DustBand

	for node1 != nil {
		// between the dust band
		if node1.InnerEdge < min && node1.OuterEdge > max {
			node2 := &DustBand{
				InnerEdge:   min,
				OuterEdge:   max,
				DustPresent: false,
			}

			if node1.GasPresent {
				node2.GasPresent = gas
			}

			node3 := &DustBand{
				InnerEdge:   max,
				OuterEdge:   node1.OuterEdge,
				GasPresent:  node1.GasPresent,
				DustPresent: node1.DustPresent,
				NextBand:    node1.NextBand,
			}

			node1.NextBand = node2
			node2.NextBand = node3
			node1.OuterEdge = min
			node1 = node3.NextBand
			continue
		}

		if node1.InnerEdge < max && node1.OuterEdge > max {
			node2 := &DustBand{
				NextBand:    node1.NextBand,
				InnerEdge:   max,
				OuterEdge:   node1.OuterEdge,
				DustPresent: node1.DustPresent,
				GasPresent:  node1.GasPresent,
			}

			node1.NextBand = node2
			node1.OuterEdge = max
			if node1.GasPresent {
				node1.GasPresent = gas
			} else {
				node1.GasPresent = false
			}
			node1.DustPresent = false
			node1 = node2.NextBand
			continue
		}

		if node1.InnerEdge < min && node1.OuterEdge > min {
			node2 := &DustBand{
				NextBand:    node1.NextBand,
				InnerEdge:   min,
				OuterEdge:   node1.OuterEdge,
				DustPresent: false,
				GasPresent:  false,
			}
			if node1.GasPresent {
				node2.GasPresent = gas
			}
			node1.NextBand = node2
			node1.OuterEdge = min
			node1 = node2.NextBand
			continue
		}

		if node1.InnerEdge >= min && node1.OuterEdge <= max {
			if node1.GasPresent {
				node1.GasPresent = gas
			}
			node1.DustPresent = false
			node1 = node1.NextBand
			continue
		}

		if node1.OuterEdge < min || node1.InnerEdge > max {
			node1 = node1.NextBand
		}
	}

	dustLeft := false
	node1 = accretionProcessor.DustBand
	for node1 != nil {
		if node1.DustPresent && node1.OuterEdge >= bodyInnerBound && node1.InnerEdge <= bodyOuterBound {
			dustLeft = true
		}
		node2 := node1.NextBand
		if node2 != nil {
			if node1.DustPresent == node2.DustPresent && node1.GasPresent == node2.GasPresent {
				node1.OuterEdge = node2.OuterEdge
				node1.NextBand = node2.NextBand
			}

		}
		node1 = node1.NextBand
	}

	return dustLeft
}

func (accretionProcessor *AccretionProcessor) accreteDust(seedMass, a, e, critMass, bodyInnerBound, bodyOuterBound, dustDensity float64) (float64, float64, float64, float64, bool) {
	tempMass := 0.0
	newMass := seedMass
	newDust, newGas, rInner, rOuter, reducedMass := 0.0, 0.0, 0.0, 0.0, 0.0

	for {
		tempMass = newMass
		newMass, newDust, newGas, rInner, rOuter, reducedMass = accretionProcessor.collectDust(newMass, a, e, critMass, dustDensity, accretionProcessor.DustBand)
		if (newMass - tempMass) < (0.0001 * tempMass) {
			break
		}
	}

	seedMass += newMass

	dustLeft := accretionProcessor.updateDustLanes(rInner, rOuter, seedMass, critMass, bodyInnerBound, bodyOuterBound)

	return seedMass, newDust, newGas, reducedMass, dustLeft

}

func (accretionProcessor *AccretionProcessor) coalescePlanetesimals(a, e, mass, critMass, dustMass, gasMass, stellLuminosityRatio, bodyInnerBound, bodyOuterBound, reducedMass, dustDensity float64, doMoons, dustLeft bool) bool {
	//the_planet :=

	var prevPlanet *Planet

	finished := false
	for thePlanet := accretionProcessor.Planet; thePlanet != nil; thePlanet = thePlanet.NextPlanet {
		diff := thePlanet.A - a

		var dist1, dist2 float64
		if diff > 0.0 {
			dist1 = (a * (1.0 + e) * (1.0 + reducedMass)) - a
			/* x aphelion	 */
			reducedMass = math.Pow((thePlanet.Mass / (1.0 + thePlanet.Mass)), (1.0 / 4.0))
			dist2 = thePlanet.A - (thePlanet.A * (1.0 - thePlanet.E) * (1.0 - reducedMass))
		} else {
			dist1 = a - (a * (1.0 - e) * (1.0 - reducedMass))
			/* x perihelion */
			reducedMass = math.Pow((thePlanet.Mass / (1.0 + thePlanet.Mass)), (1.0 / 4.0))
			dist2 = (thePlanet.A * (1.0 + thePlanet.E) * (1.0 + reducedMass)) - thePlanet.A
		}

		if (math.Abs(diff) <= math.Abs(dist1)) || (math.Abs(diff) <= math.Abs(dist2)) {
			newA := (thePlanet.Mass + mass) / ((thePlanet.Mass / thePlanet.A) + (mass / a))

			temp := thePlanet.Mass * math.Sqrt(thePlanet.A) * math.Sqrt(1.0-math.Pow(thePlanet.E, 2.0))
			temp = temp + (mass * math.Sqrt(a) * math.Sqrt(math.Sqrt(1.0-math.Pow(e, 2.0))))
			temp = temp / ((thePlanet.Mass + mass) * math.Sqrt(newA))
			temp = 1.0 - math.Pow(temp, 2.0)

			if temp < 0.0 || temp >= 1.0 {
				temp = 0.0
			}

			e = math.Sqrt(temp)

			if doMoons {
				existingMass := 0.0

				for _, moon := range thePlanet.Moons {
					existingMass += moon.Mass
				}

				if mass < critMass {
					if (mass*SunMassInEarthMasses) < 2.5 && (mass*SunMassInEarthMasses) > 0.0001 && existingMass < (thePlanet.Mass*0.05) {
						moon := &Planet{
							Mass:     mass,
							DustMass: dustMass,
							GasMass:  gasMass,
							GasGiant: false,
						}

						// check if the moon more mass then the planet
						if (moon.DustMass + moon.GasMass) > (thePlanet.DustMass + thePlanet.GasMass) {
							tmpDustMass := thePlanet.DustMass
							tmpGasMass := thePlanet.GasMass
							tmpMass := thePlanet.Mass

							thePlanet.DustMass = moon.DustMass
							thePlanet.GasMass = moon.GasMass
							thePlanet.Mass = moon.Mass

							moon.DustMass = tmpDustMass
							moon.GasMass = tmpGasMass
							moon.Mass = tmpMass
						}

						thePlanet.Moons = append(thePlanet.Moons, moon)
						finished = true

					}
				}
			}

			if !finished {
				temp = thePlanet.Mass + mass

				seedMass, newDust, newGas := 0.0, 0.0, 0.0
				seedMass, newDust, newGas, _, dustLeft = accretionProcessor.accreteDust(temp, newA, e, stellLuminosityRatio, bodyInnerBound, bodyOuterBound, dustDensity)

				thePlanet.A = newA
				thePlanet.E = e
				thePlanet.Mass = seedMass
				thePlanet.DustMass += dustMass + newDust
				thePlanet.GasMass += gasMass + newGas
				if temp >= critMass {
					thePlanet.GasGiant = true
				}

				for thePlanet.NextPlanet != nil && thePlanet.NextPlanet.A < newA {
					nextPalnet := thePlanet.NextPlanet

					if thePlanet == accretionProcessor.Planet {
						accretionProcessor.Planet = nextPalnet
					} else {
						prevPlanet.NextPlanet = nextPalnet
					}

					thePlanet.NextPlanet = nextPalnet.NextPlanet
					nextPalnet.NextPlanet = thePlanet
					prevPlanet = nextPalnet
				}

			}

			finished = true
			break
		} else {
			prevPlanet = thePlanet
		}

	}

	// Planetesimals didn't collide. Make it a planet.
	if !finished {

		// create a new planet
		planet := &Planet{
			A:        a,
			E:        e,
			Mass:     mass,
			DustMass: dustMass,
			GasMass:  gasMass,
			GasGiant: mass >= critMass,
			Moons:    make([]*Planet, 0),
		}

		// check if the first planet
		if accretionProcessor.Planet == nil {
			accretionProcessor.Planet = planet
			return dustLeft
		}

		// check if the planet before the first planet
		if a < accretionProcessor.Planet.A {
			planet.NextPlanet = accretionProcessor.Planet
			accretionProcessor.Planet = planet
			return dustLeft
		}

		// check if next planet from the first exists
		if accretionProcessor.Planet.NextPlanet == nil {
			accretionProcessor.Planet.NextPlanet = planet
			return dustLeft
		}

		// find the place to place the planet
		nextPlanet := accretionProcessor.Planet
		for nextPlanet != nil && nextPlanet.A < a {
			prevPlanet = nextPlanet
			nextPlanet = nextPlanet.NextPlanet
		}

		planet.NextPlanet = nextPlanet
		prevPlanet.NextPlanet = planet
	}

	return dustLeft

}

func (accretionProcessor *AccretionProcessor) DistPlanetaryMasses(stellMassRatio, stellLuminosityRatio, innerDust, outerDust, outerPlanetLimit, dustDensityCoeff float64, doMoons bool) {
	// create the initial dust band
	accretionProcessor.InitialDustBand(innerDust, outerDust)

	planetInnerBound := accretionProcessor.nearestPlanet(stellMassRatio)
	planetOuterBound := accretionProcessor.farthestPlanet(stellMassRatio)

	dustLeft := true
	for dustLeft {
		a := utils.RandFloat64(planetInnerBound, planetOuterBound)
		e := accretionProcessor.randEccentricity()

		mass := ProtoPlanetMass

		innerEffectLimit := accretionProcessor.innerEffectLimit(a, e, mass)
		outerEffectLimit := accretionProcessor.outerEffectLimit(a, e, mass)

		if accretionProcessor.isDustAvailable(innerEffectLimit, outerEffectLimit) {
			dustDensity := dustDensityCoeff * math.Sqrt(stellMassRatio) * math.Exp(-Alpha*math.Pow(a, (1.0/N)))
			critMass := accretionProcessor.criticalLimit(a, e, stellLuminosityRatio)

			dustMass, gasMass, reducedMass := 0.0, 0.0, 0.0
			mass, dustMass, gasMass, reducedMass, dustLeft = accretionProcessor.accreteDust(mass, a, e, critMass, planetInnerBound, planetOuterBound, dustDensity)

			dustMass += ProtoPlanetMass

			if mass > ProtoPlanetMass {
				dustLeft = accretionProcessor.coalescePlanetesimals(a, e, mass, critMass, dustMass, gasMass, stellLuminosityRatio, planetInnerBound, planetOuterBound, reducedMass, dustDensity, doMoons, dustLeft)
			}
		}

	}
}
