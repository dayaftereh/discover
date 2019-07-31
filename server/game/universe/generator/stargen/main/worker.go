package main

import (
	"sync"

	"github.com/dayaftereh/discover/server/game/universe/generator/stargen"
	"github.com/dayaftereh/discover/server/utils/atomic"
)

type Worker struct {
	running    *atomic.AtomicBool
	lock       sync.Mutex
	statistics *Statistics
}

func NewWorker() *Worker {
	return &Worker{
		running:    atomic.NewAtomicBool(false),
		statistics: NewStatistics(),
	}
}

func (worker *Worker) Start() {
	alreadyRunning := worker.running.GetAndSet(true)
	if alreadyRunning {
		return
	}
	go worker.Run()
}

func (worker *Worker) inspectPlanet(planet *stargen.Planet) {
	count, ok := worker.statistics.Types[planet.Type]
	if ok {
		worker.statistics.Types[planet.Type] = count + 1
	} else {
		worker.statistics.Types[planet.Type] = 1
	}

	count, ok = worker.statistics.Oxygen[planet.Breathability]
	if ok {
		worker.statistics.Oxygen[planet.Breathability] = count + 1
	} else {
		worker.statistics.Oxygen[planet.Breathability] = 1
	}

	for _, gas := range planet.Atmosphere {
		count, ok = worker.statistics.Gases[gas.Num]
		if ok {
			worker.statistics.Gases[gas.Num] = count + 1
		} else {
			worker.statistics.Gases[gas.Num] = 1
		}
	}
}

func (worker *Worker) appendPlanet(planet *stargen.Planet) {
	worker.inspectPlanet(planet)

	if planet.Type == stargen.PlanetUnknown {
		worker.statistics.UnknownPlanets++
	}

	// count each moon
	worker.statistics.Moons += int64(len(planet.Moons))

	// inspect each moon
	for _, moon := range planet.Moons {
		worker.inspectPlanet(moon)

		if moon.Type == stargen.PlanetUnknown {
			worker.statistics.UnknownMoons++
		}
	}
}

func (worker *Worker) updateStatistics(planets []*stargen.Planet) {
	worker.lock.Lock()
	defer worker.lock.Unlock()

	worker.statistics.Executions++

	// add found planets
	worker.statistics.Planets += int64(len(planets))

	// look at each planet
	for _, planet := range planets {
		worker.appendPlanet(planet)
	}
}

func (worker *Worker) Run() {
	for worker.running.Get() {

		// generate stellar system
		_, planets := stargen.GenerateStellarSystem(true, true, true)

		worker.updateStatistics(planets)

	}
}

func (worker *Worker) AddToStatistics(statistics *Statistics) {
	worker.lock.Lock()
	defer worker.lock.Unlock()

	statistics.Executions += worker.statistics.Executions
	statistics.Planets += worker.statistics.Planets
	statistics.UnknownPlanets += worker.statistics.UnknownPlanets
	statistics.Moons += worker.statistics.Moons
	statistics.UnknownMoons += worker.statistics.UnknownMoons

	for typ, count := range worker.statistics.Types {
		c, ok := statistics.Types[typ]
		if !ok {
			statistics.Types[typ] = count
		} else {
			statistics.Types[typ] = c + count
		}
	}

	for oxygen, count := range worker.statistics.Oxygen {
		c, ok := statistics.Oxygen[oxygen]
		if !ok {
			statistics.Oxygen[oxygen] = count
		} else {
			statistics.Oxygen[oxygen] = c + count
		}
	}

	for gas, count := range worker.statistics.Gases {
		c, ok := statistics.Gases[gas]
		if !ok {
			statistics.Gases[gas] = count
		} else {
			statistics.Gases[gas] = c + count
		}
	}
}

func (worker *Worker) Stop() {
	worker.running.Set(false)
}
