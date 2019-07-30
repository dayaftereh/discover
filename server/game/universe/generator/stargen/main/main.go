package main

import (
	"fmt"

	"github.com/dayaftereh/discover/server/game/universe/generator/stargen"
	"github.com/dayaftereh/discover/server/utils"
)

type Statistics struct {
	Executions     int64
	Moons          int64
	Planets        int64
	UnknownPlanets int64
	UnknownMoons   int64
	Oxygen         map[stargen.Oxygen]int64
	Types          map[stargen.PlanetType]int64
}

func NewStatistics() *Statistics {
	statistics := &Statistics{
		Types:  make(map[stargen.PlanetType]int64),
		Oxygen: make(map[stargen.Oxygen]int64),
	}

	// set all planet types to zero
	statistics.Types[stargen.PlanetRock] = 0
	statistics.Types[stargen.PlanetVenusian] = 0
	statistics.Types[stargen.PlanetTerrestrial] = 0
	statistics.Types[stargen.PlanetGasGiant] = 0
	statistics.Types[stargen.PlanetMartian] = 0
	statistics.Types[stargen.PlanetWater] = 0
	statistics.Types[stargen.PlanetIce] = 0
	statistics.Types[stargen.PlanetSubGasGiant] = 0
	statistics.Types[stargen.PlanetSubSubGasGiant] = 0
	statistics.Types[stargen.PlanetAsteroids] = 0
	statistics.Types[stargen.Planet1Face] = 0
	statistics.Types[stargen.PlanetUnknown] = 0

	// set all Oxygen types to zero
	statistics.Oxygen[stargen.None] = 0
	statistics.Oxygen[stargen.Toxic] = 0
	statistics.Oxygen[stargen.Unbreathable] = 0
	statistics.Oxygen[stargen.Breathable] = 0

	return statistics
}

func (statistics *Statistics) String() string {
	s := fmt.Sprintf("Executions: %d\n\n", statistics.Executions)

	s = fmt.Sprintf("%s Planets: %d, %1.3f\n", s, statistics.Planets, float64(statistics.Planets)/float64(statistics.Executions))
	s = fmt.Sprintf("%s UnknownPlanets: %d, %1.3f\n", s, statistics.UnknownPlanets, float64(statistics.UnknownPlanets)/float64(statistics.Executions))

	s = fmt.Sprintf("%s Moons: %d, %1.3f\n", s, statistics.Moons, float64(statistics.Moons)/float64(statistics.Executions))
	s = fmt.Sprintf("%s UnknownMoons: %d, %1.3f\n\n", s, statistics.UnknownMoons, float64(statistics.UnknownMoons)/float64(statistics.Executions))

	s = fmt.Sprintf("%s Types:\n", s)

	// find found types
	for typ, count := range statistics.Types {
		s = fmt.Sprintf("%s %v: %d, %1.3f\n", s, typ, count, float64(count)/float64(statistics.Executions))
	}

	s = fmt.Sprintf("%s Oxygen:\n", s)

	// find found types
	for oxygen, count := range statistics.Oxygen {
		s = fmt.Sprintf("%s %v: %d, %1.3f\n", s, oxygen, count, float64(count)/float64(statistics.Executions))
	}

	s = fmt.Sprintf("%s\n\n", s)

	return s
}

func InspectPlanet(statistics *Statistics, planet *stargen.Planet) {
	count, ok := statistics.Types[planet.Type]
	if ok {
		statistics.Types[planet.Type] = count + 1
	} else {
		statistics.Types[planet.Type] = 1
	}

	count, ok = statistics.Oxygen[planet.Breathability]
	if ok {
		statistics.Oxygen[planet.Breathability] = count + 1
	} else {
		statistics.Oxygen[planet.Breathability] = 1
	}
}

func AppendPlanet(statistics *Statistics, planet *stargen.Planet) {
	InspectPlanet(statistics, planet)

	if planet.Type == stargen.PlanetUnknown {
		statistics.UnknownPlanets++
	}

	// count each moon
	statistics.Moons += int64(len(planet.Moons))

	// inspect each moon
	for _, moon := range planet.Moons {
		InspectPlanet(statistics, moon)

		if moon.Type == stargen.PlanetUnknown {
			statistics.UnknownMoons++
		}
	}
}

func main() {
	lastPrint := 0.0
	statistics := NewStatistics()

	for {
		statistics.Executions++

		// generate stellar system
		_, planets := stargen.GenerateStellarSystem(true, true, true)
		// add found planets
		statistics.Planets += int64(len(planets))

		// look at each planet
		for _, planet := range planets {
			AppendPlanet(statistics, planet)
		}

		// calculate last print delta of statistics
		delta := utils.SystemSeconds() - lastPrint

		// above 10 s
		if delta > 10.0 {
			lastPrint = utils.SystemSeconds()
			// show statistics
			fmt.Println(statistics)
		}
	}

}
