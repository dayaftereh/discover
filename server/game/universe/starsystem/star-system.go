package starsystem

import (
	"log"
	"time"

	"github.com/dayaftereh/discover/server/game/engine"
	"github.com/dayaftereh/discover/server/game/player"

	"github.com/dayaftereh/discover/server/game/engine/world"
)

type StarSystem struct {
	ID      int64
	counter int64
	// players
	players       map[string]*player.Player
	playersObject map[string]int64
	// Engine
	clock *engine.Clock
	world *world.World
	// events
	close chan bool
	queue chan StarSystemFunction
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
		queue: make(chan StarSystemFunction),
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
		case f := <-starSystem.queue:
			starSystem.execute(f)
		case <-starSystem.close:
			log.Printf("closing star-system [ %d ]\n", starSystem.ID)
			return
		}
	}
}

func (starSystem *StarSystem) update() {
	// get the delta for the update
	delta := starSystem.clock.Delta()

	log.Printf("Update: delat [ %f ]\n", delta)

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
