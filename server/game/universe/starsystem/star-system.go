package starsystem

import (
	"log"
	"sync"
	"time"

	"github.com/dayaftereh/discover/server/game/engine"
	"github.com/dayaftereh/discover/server/game/player"

	"github.com/dayaftereh/discover/server/game/engine/world"
)

type StarSystem struct {
	ID      int64
	counter int64
	lock    sync.Mutex
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

func NewStarSystem(id int64) *StarSystem {
	starSystem := &StarSystem{
		ID:      id,
		counter: 0,

		// players
		playersObject: make(map[string]int64),
		players:       make(map[string]*player.Player),

		world: world.NewWorld(),
		clock: engine.NewClock(),

		close: make(chan bool),
	}

	go starSystem.loop()

	return starSystem
}

func (starSystem *StarSystem) loop() {
	// update by 30 fps
	timer := (1000 / 30) * time.Millisecond
	log.Printf("starting star-system [ %d ]\n", starSystem.ID)
	for {
		select {
		case <-time.After(timer):
			starSystem.update()
		case <-starSystem.close:
			log.Printf("closing star-system [ %d ]\n", starSystem.ID)
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

	if delta > 35.0 {
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
	starSystem.close <- true
}
