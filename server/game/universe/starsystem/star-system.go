package starsystem

import (
	"log"
	"sync"
	"time"

	"github.com/dayaftereh/discover/server/game/persistence"
	"github.com/dayaftereh/discover/server/game/persistence/types"
	"github.com/dayaftereh/discover/server/game/universe/starsystem/objects"
	"github.com/dayaftereh/discover/server/utils/atomic"

	"github.com/dayaftereh/discover/server/game/engine"
	"github.com/dayaftereh/discover/server/game/player"

	"github.com/dayaftereh/discover/server/game/engine/world"
)

type StarSystem struct {
	Name    string
	counter int64
	lock    sync.Mutex
	close   chan bool
	running *atomic.AtomicBool
	// the star system data
	Data *types.StarSystem
	// starSystem
	sun *objects.Sun
	// players
	players       map[string]*player.Player
	playersObject map[string]int64
	// Engine
	clock *engine.Clock
	world *world.World
	// persistence
	persistence *persistence.PersistenceManager
}

func NewStarSystem(name string, persistence *persistence.PersistenceManager) *StarSystem {
	// create the star system
	starSystem := &StarSystem{
		Name:    name,
		counter: 0,
		// players
		playersObject: make(map[string]int64),
		players:       make(map[string]*player.Player),

		// engine
		world: world.NewWorld(),
		clock: engine.NewClock(),

		// close
		close:   make(chan bool),
		running: atomic.NewAtomicBool(false),
		// set persistence manager
		persistence: persistence,
	}

	return starSystem
}

func (starSystem *StarSystem) loop() {
	defer func() {
		log.Printf("closing star-system [ %s ] thread\n", starSystem.Name)
	}()

	// update by 30 fps
	timer := (1000 / 30) * time.Millisecond

	// start the clock
	starSystem.clock.Start()

	log.Printf("starting star-system [ %s ]\n", starSystem.Name)
	for starSystem.running.Get() {
		select {
		case <-time.After(timer):
			starSystem.update()
		case <-starSystem.close:
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
	// try to shutdown the star system
	running := starSystem.running.GetAndSet(false)
	// stop if the star system already shutdowned
	if !running {
		return
	}

	log.Printf("shutdown star-system with id [ %d ]\n", starSystem.Name)
	starSystem.close <- true
}
