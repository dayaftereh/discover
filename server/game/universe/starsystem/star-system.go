package starsystem

import (
	"time"

	"github.com/dayaftereh/discover/server/game/player"
	"github.com/dayaftereh/discover/server/mathf"

	"github.com/dayaftereh/discover/server/game/engine/world"
)

type StarSystemFunction func(starSystem *StarSystem) error

type StarSystem struct {
	ID int64
	//private
	players map[string]int64
	world   *world.World
	close   chan bool
	queue   chan StarSystemFunction
}

func NewStarSystem(id int64) *StarSystem {
	starSystem := &StarSystem{
		ID:      id,
		players: make(map[string]int64),
		world:   world.NewWorld(),
		close:   make(chan bool),
		queue:   make(chan StarSystemFunction),
	}

	go starSystem.loop()

	return starSystem
}

func (starSystem *StarSystem) loop() {
	for {
		select {
		case <-time.After(10 * time.Millisecond):
			starSystem.update()
		case f := <-starSystem.queue:
			starSystem.execute(f)
		case <-starSystem.close:
			return
		}
	}
}

func (starSystem *StarSystem) execute(function StarSystemFunction) {
	err := function(starSystem)
	if err != nil {
	}
}

func (starSystem *StarSystem) update() {
	starSystem.world.Update(10)
}

func (starSystem *StarSystem) JoinPlayer(player *player.Player) {
	starSystem.queue <- func(starSystem *StarSystem) error {

		return nil
	}
}

func (starSystem *StarSystem) UpdatePlayer(player *player.Player, lookAt *mathf.Vec3) {
	starSystem.queue <- func(starSystem *StarSystem) error {

		return nil
	}
}

func (starSystem *StarSystem) DropPlayer(player *player.Player) {
	starSystem.queue <- func(starSystem *StarSystem) error {

		return nil
	}
}
