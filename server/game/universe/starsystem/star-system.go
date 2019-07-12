package starsystem

import (
	"log"
	"time"

	"github.com/dayaftereh/discover/server/utils"

	"github.com/pkg/errors"

	types "github.com/dayaftereh/discover/server/api/types/connection"
	"github.com/dayaftereh/discover/server/game/engine"
	"github.com/dayaftereh/discover/server/game/player"
	"github.com/dayaftereh/discover/server/mathf"

	"github.com/dayaftereh/discover/server/game/engine/object"
	"github.com/dayaftereh/discover/server/game/engine/world"
)

type StarSystemFunction func(starSystem *StarSystem) error

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
	log.Printf("starting star-system [ %d ]\n", starSystem.ID)
	for {
		select {
		case <-time.After(10 * time.Millisecond):
			starSystem.update()
		case f := <-starSystem.queue:
			starSystem.execute(f)
		case <-starSystem.close:
			log.Printf("closing star-system [ %d ]\n", starSystem.ID)
			return
		}
	}
}

func (starSystem *StarSystem) execute(function StarSystemFunction) {
	err := function(starSystem)
	if err != nil {
		log.Printf("fail to execute function in star-system [ %d ], because %v", starSystem.ID, err)
	}
}

func (starSystem *StarSystem) update() {
	// get the delta for the update
	delta := starSystem.clock.Delta()

	// update the world
	starSystem.world.Update(delta)

	// push the world update for the players
	starSystem.pushWorldUpdates()
}

func (starSystem *StarSystem) pushWorldUpdates() {
	// current server time
	now := utils.SystemMillis()

	// update each player
	for playerID, player := range starSystem.players {
		// get the player game object
		playerObjectID, ok := starSystem.playersObject[playerID]
		// check if player has game object
		if !ok {
			continue
		}
		// get the player game object
		gameObject := starSystem.world.GetGameObject(playerObjectID)

		// convert the player gameobject to outbound object
		playerGameObject := gameObjectToOutbound(gameObject)

		// get all objects in player range
		playerObjects := starSystem.world.GetGameObjectsInSphere(gameObject, 100.0)

		// convert game objects for outbound
		gameObjects := gameObjectsToOutbound(playerObjects)

		// get the world update tick
		tick := starSystem.world.GetTick()

		// push the update
		update := &types.WorldUpdate{
			Type:    types.Update,
			Tick:    &tick,
			Time:    &now,
			Player:  playerGameObject,
			Objects: gameObjects,
		}

		// push the update for the player
		player.Push(update)
	}
}

func (starSystem *StarSystem) nextID() int64 {
	// increment counter
	starSystem.counter++
	// get next id
	return starSystem.counter
}

func (starSystem *StarSystem) JoinPlayer(player *player.Player) {
	starSystem.queue <- func(starSystem *StarSystem) error {
		// store the player
		starSystem.players[player.ID] = player

		// get the object id
		id := starSystem.nextID()

		// create player game object
		gameObject := object.NewPlayer(id, mathf.NewZeroVec3())

		// add the game object to world
		starSystem.world.AddGameObject(gameObject)

		// map player to game object
		starSystem.playersObject[player.ID] = id

		return nil
	}
}

func (starSystem *StarSystem) UpdatePlayer(player *player.Player, move *mathf.Vec3, rotation *mathf.Vec3) {
	starSystem.queue <- func(starSystem *StarSystem) error {
		_, ok := starSystem.players[player.ID]
		// check if the player is joined to this star system
		if !ok {
			return nil
		}

		// get game object id for player
		gameObjectID, ok := starSystem.playersObject[player.ID]
		// check if game object id ist mapped
		if !ok {
			return errors.Errorf("unable to find game-object id for player [ %s ]", player.Name)
		}

		// finally get the game object
		gameObject := starSystem.world.GetGameObject(gameObjectID)

		if gameObject == nil {
			return errors.Errorf("unable to find game-object with id [ %d ] for player [ %s ]", gameObjectID, player.Name)
		}

		return nil
	}
}

func (starSystem *StarSystem) DropPlayer(player *player.Player) {
	starSystem.queue <- func(starSystem *StarSystem) error {
		_, ok := starSystem.players[player.ID]
		// check if the player is joined to this star system
		if !ok {
			return nil
		}
		// remove the player from star syste,
		delete(starSystem.players, player.ID)

		// get the gameObjectID for the player
		gameObjectID, ok := starSystem.playersObject[player.ID]
		// check if game object found
		if !ok {
			return errors.Errorf("unable to find game-object [ %d ] for player [ %s ]", gameObjectID, player.Name)
		}
		// remove the mapping
		delete(starSystem.playersObject, player.ID)
		// remove the game object from world
		starSystem.world.RemoveGameObject(gameObjectID)

		return nil
	}
}

func (starSystem *StarSystem) Shutdown() {
	starSystem.close <- true
}
