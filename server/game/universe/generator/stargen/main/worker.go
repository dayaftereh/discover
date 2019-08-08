package main

import (
	"fmt"
	"log"
	"sync"

	persistence "github.com/dayaftereh/discover/server/game/persistence/types"
	"github.com/dayaftereh/discover/server/game/universe/generator/stargen"
	"github.com/dayaftereh/discover/server/game/universe/generator/stargen/types"
	"github.com/dayaftereh/discover/server/utils/atomic"
)

type Worker struct {
	worker     int
	executions int64
	running    *atomic.AtomicBool
	lock       sync.Mutex
	statistics *Statistics
	display    *Display
}

func NewWorker(worker int, executions int64, display *Display) *Worker {
	return &Worker{
		worker:     worker,
		executions: executions,
		running:    atomic.NewAtomicBool(false),
		statistics: NewStatistics(),
		display:    display,
	}
}

func (worker *Worker) Start() {
	alreadyRunning := worker.running.GetAndSet(true)
	if alreadyRunning {
		return
	}
	go worker.Run()
}

func (worker *Worker) inspectPlanet(planet *persistence.Planet) {
	count, ok := worker.statistics.PlanetTypes[planet.Type]
	if ok {
		worker.statistics.PlanetTypes[planet.Type] = count + 1
	} else {
		worker.statistics.PlanetTypes[planet.Type] = 1
	}

	count, ok = worker.statistics.AtmosphereTypes[planet.AtmosphereType]
	if ok {
		worker.statistics.AtmosphereTypes[planet.AtmosphereType] = count + 1
	} else {
		worker.statistics.AtmosphereTypes[planet.AtmosphereType] = 1
	}

	for _, gas := range planet.Atmosphere {
		count, ok = worker.statistics.AtmosphereGases[gas.Num]
		if ok {
			worker.statistics.AtmosphereGases[gas.Num] = count + 1
		} else {
			worker.statistics.AtmosphereGases[gas.Num] = 1
		}
	}
}

func (worker *Worker) appendPlanet(planet *persistence.Planet) {
	worker.inspectPlanet(planet)

	if planet.Type == types.PlanetUnknown {
		worker.statistics.UnknownPlanets++
	}

	// count each moon
	worker.statistics.Moons += int64(len(planet.Moons))

	// inspect each moon
	for _, moon := range planet.Moons {
		worker.inspectPlanet(moon)

		if moon.Type == types.PlanetUnknown {
			worker.statistics.UnknownMoons++
		}
	}
}

func (worker *Worker) updateStatistics(planets []*persistence.Planet) {
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
	executions := int64(0)
	for worker.running.Get() {
		// check if worker done
		if worker.executions > 0 && executions >= worker.executions {
			return
		}

		// generate stellar system
		sun, planets := stargen.GenerateStellarSystem(true, true, true)

		name := fmt.Sprintf("worker-%d-execution-%d", worker.worker, executions)
		err := worker.display.export(sun, planets, name)
		if err != nil {
			log.Println(err)
			return
		}

		worker.updateStatistics(planets)

		executions++
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

	for typ, count := range worker.statistics.PlanetTypes {
		c, ok := statistics.PlanetTypes[typ]
		if !ok {
			statistics.PlanetTypes[typ] = count
		} else {
			statistics.PlanetTypes[typ] = c + count
		}
	}

	for oxygen, count := range worker.statistics.AtmosphereTypes {
		c, ok := statistics.AtmosphereTypes[oxygen]
		if !ok {
			statistics.AtmosphereTypes[oxygen] = count
		} else {
			statistics.AtmosphereTypes[oxygen] = c + count
		}
	}

	for gas, count := range worker.statistics.AtmosphereGases {
		c, ok := statistics.AtmosphereGases[gas]
		if !ok {
			statistics.AtmosphereGases[gas] = count
		} else {
			statistics.AtmosphereGases[gas] = c + count
		}
	}
}

func (worker *Worker) Stop() {
	worker.running.Set(false)
}
