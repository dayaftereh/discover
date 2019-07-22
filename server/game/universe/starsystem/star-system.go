package starsystem

import (
	"log"
	"sync"
	"time"

	"github.com/dayaftereh/discover/server/game/data"

	"github.com/dayaftereh/discover/server/game/engine"
	"github.com/dayaftereh/discover/server/game/player"

	"github.com/dayaftereh/discover/server/game/engine/world"
)

type StarSystem struct {
	ID      int64
	Name    string
	counter int64
	lock    sync.Mutex
	// starSystem
	sunMass  float64
	sunColor int64
	// players
	players       map[string]*player.Player
	playersObject map[string]int64
	// Engine
	clock *engine.Clock
	world *world.World
	// events
	close chan bool
	//
}

func NewStarSystem(starSystemData *data.StarSystem) *StarSystem {
	// create the star system
	starSystem := &StarSystem{
		ID:      starSystemData.ID,
		Name:    starSystemData.Name,
		counter: 0,

		// players
		playersObject: make(map[string]int64),
		players:       make(map[string]*player.Player),

		// engine
		world: world.NewWorld(),
		clock: engine.NewClock(),

		// close
		close: make(chan bool),
	}

	// load the star system from data
	starSystem.loadData(starSystemData)

	go starSystem.loop()

	return starSystem
}

func (starSystem *StarSystem) loop() {
	// update by 30 fps
	timer := (1000 / 30) * time.Millisecond

	// start the clock
	starSystem.clock.Start()

	log.Printf("starting star-system [ %d ]\n", starSystem.ID)
	for {
		select {
		case <-time.After(timer):
			starSystem.update()
		case <-starSystem.close:
			log.Printf("closing star-system [ %d ] thread\n", starSystem.ID)
			return
		}
	}
}

func (starSystem *StarSystem) update() {
	// lock the star system
	starSystem.lock.Lock()
	defer starSystem.lock.Unlock()

	// get the delta for the update
	delta := starSystem.clock.Delta()

	if delta > 0.035 {
		log.Printf("Update: delat [ %f ]\n", delta)
	}

	// update the world
	starSystem.world.Update(delta)

	// push the world update for the players
	starSystem.pushWorldUpdates()
}

func (starSystem *StarSystem) nextID() int64 {
	// increment counter
	starSystem.counter++
	// get next id
	return starSystem.counter
}

func (starSystem *StarSystem) Shutdown() {
	log.Printf("shutdown star-system with id [ %d ]\n", starSystem.ID)
	starSystem.close <- true
}
