package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dayaftereh/discover/server/game/universe/generator/stargen"
)

type Statistics struct {
	Executions     int64
	Moons          int64
	Planets        int64
	UnknownPlanets int64
	UnknownMoons   int64
	Oxygen         map[stargen.Oxygen]int64
	Types          map[stargen.PlanetType]int64
	Gases          map[int64]int64
}

func NewStatistics() *Statistics {
	statistics := &Statistics{
		Types:  make(map[stargen.PlanetType]int64),
		Oxygen: make(map[stargen.Oxygen]int64),
		Gases:  make(map[int64]int64),
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
		s = fmt.Sprintf("%s\t%v: %d, %1.3f\n", s, typ, count, float64(count)/float64(statistics.Executions))
	}

	s = fmt.Sprintf("\n%s Oxygen:\n", s)

	// find found types
	for oxygen, count := range statistics.Oxygen {
		s = fmt.Sprintf("%s\t%v: %d, %1.3f\n", s, oxygen, count, float64(count)/float64(statistics.Executions))
	}

	s = fmt.Sprintf("\n%s Gases:\n", s)

	// print the found gases
	for gasNum, count := range statistics.Gases {
		gas, ok := stargen.GasesTable[gasNum]
		if ok {
			s = fmt.Sprintf("%s\t%v: %d, %1.3f\n", s, gas.Symbol, count, float64(count)/float64(statistics.Executions))
		}
	}

	s = fmt.Sprintf("%s\n\n", s)

	return s
}

func CollectAndPrint(workers []*Worker, statistics *Statistics) {
	for _, worker := range workers {
		worker.AddToStatistics(statistics)
	}

	fmt.Println(statistics)
}

func Loop(workers []*Worker, statistics *Statistics) {
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case <-time.After(10 * time.Second):
			CollectAndPrint(workers, statistics)
		case <-signalChannel:
			return
		}
	}
}

func main() {

	concurrently := 3

	workers := make([]*Worker, 0)
	for i := 0; i < concurrently; i++ {
		worker := NewWorker()
		workers = append(workers, worker)
		worker.Start()
	}

	statistics := NewStatistics()

	Loop(workers, statistics)

	for _, worker := range workers {
		worker.Stop()
	}
}
