package starsystem

import (
	"log"
	"time"

	"github.com/dayaftereh/discover/server/utils"

	"github.com/pkg/errors"

	types "github.com/dayaftereh/discover/server/api/types/connection"
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
	// objects
	objects map[int64]object.GameObject
	// Engine
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
		// objects
		objects: make(map[int64]object.GameObject),

		world: world.NewWorld(),

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
	}
}

func (starSystem *StarSystem) update() {

}

func (starSystem *StarSystem) pushWorldUpdate() {
	// current server time
	now := utils.SystemMillis()

	// map with all game objects
	gameObjects := make(map[int64]*types.GameObject)
	for gameObjectID, gameObject := range starSystem.objects {
		body := gameObject.Body()

		radius := body.BoundingRadius

		// get location and roation
		position := body.Position.Clone()
		rotation := body.Quaternion.ToEuler()

		// create the outgoing gameobject
		gameObjects[gameObjectID] = &types.GameObject{
			Radius:   &radius,
			Position: position,
			Rotation: rotation,
		}
	}

	// update each player
	for playerID, player := range starSystem.players {
		// get the player game object
		playerObjectID, ok := starSystem.playersObject[playerID]
		// check if player has game object
		if !ok {
			continue
		}
		// get the outgoing game object
		gameObject := gameObjects[playerObjectID]

		tick := int64(0)

		// push the update
		update := &types.WorldUpdate{
			Type:    types.Update,
			Tick:    &tick,
			Time:    &now,
			Player:  gameObject,
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
		starSystem.objects[id] = object.NewPlayer(id, mathf.NewZeroVec3())

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
		_, ok = starSystem.objects[gameObjectID]

		// check if game object exists
		if !ok {
			return errors.Errorf("unable to find game-object [ %d ] for player [ %s ]", gameObjectID, player.Name)
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
		// remove the game object
		delete(starSystem.objects, gameObjectID)

		return nil
	}
}

func (starSystem *StarSystem) Shutdown() {
	starSystem.close <- true
}
