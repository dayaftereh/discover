package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dayaftereh/discover/server/game/universe/generator/stargen/chemical"
	"github.com/dayaftereh/discover/server/game/universe/generator/stargen/types"
)

type Statistics struct {
	Executions      int64
	Moons           int64
	Planets         int64
	UnknownPlanets  int64
	UnknownMoons    int64
	AtmosphereGases map[int64]int64
	AtmosphereTypes map[types.AtmosphereType]int64
	PlanetTypes     map[types.PlanetType]int64
}

func NewStatistics() *Statistics {
	statistics := &Statistics{
		PlanetTypes:     make(map[types.PlanetType]int64),
		AtmosphereGases: make(map[int64]int64),
		AtmosphereTypes: make(map[types.AtmosphereType]int64),
	}

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
	for typ, count := range statistics.PlanetTypes {
		s = fmt.Sprintf("%s\t%v: %d, %1.3f\n", s, typ, count, float64(count)/float64(statistics.Executions))
	}

	s = fmt.Sprintf("\n%s Oxygen:\n", s)

	// find found types
	for oxygen, count := range statistics.AtmosphereTypes {
		s = fmt.Sprintf("%s\t%v: %d, %1.3f\n", s, oxygen, count, float64(count)/float64(statistics.Executions))
	}

	s = fmt.Sprintf("\n%s Gases:\n", s)

	// print the found gases
	for gasNum, count := range statistics.AtmosphereGases {
		gas, ok := chemical.PeriodicTable[gasNum]
		if ok {
			s = fmt.Sprintf("%s\t%v: %d, %1.3f\n", s, gas.Name, count, float64(count)/float64(statistics.Executions))
		}
	}

	s = fmt.Sprintf("%s\n\n", s)

	return s
}

func CollectAndPrint(workers []*Worker, statistics *Statistics, executions int64) bool {
	for _, worker := range workers {
		worker.AddToStatistics(statistics)
	}

	fmt.Println(statistics)

	return !(statistics.Executions < executions)
}

func Loop(workers []*Worker, statistics *Statistics, executions int64) {
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case <-time.After(5 * time.Second):
			done := CollectAndPrint(workers, statistics, executions)
			if done {
				return
			}
		case <-signalChannel:
			return
		}
	}
}

func main() {
	display, err := NewDisplay("./dist/star-systems", "./server/game/universe/generator/stargen/main/templates/*.html")
	if err != nil {
		log.Panicln(err)
	}

	executions := int64(1)
	concurrently := 1

	workers := make([]*Worker, 0)
	for i := 0; i < concurrently; i++ {
		worker := NewWorker(i+1, executions, display)
		workers = append(workers, worker)
		worker.Start()
	}

	statistics := NewStatistics()

	Loop(workers, statistics, executions)

	for _, worker := range workers {
		worker.Stop()
	}
}
