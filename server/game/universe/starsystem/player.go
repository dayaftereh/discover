package starsystem

import (
	"log"

	"github.com/dayaftereh/discover/server/game/player"
	"github.com/dayaftereh/discover/server/game/universe/starsystem/objects"
	"github.com/dayaftereh/discover/server/mathf"
	"github.com/pkg/errors"
)

func (starSystem *StarSystem) JoinPlayer(player *player.Player) {
	starSystem.lock.Lock()
	defer starSystem.lock.Unlock()

	// check if player already in the star system
	_, ok := starSystem.players[player.ID]
	if ok {
		log.Printf("player [ %s ] already joined star-system [ %s ]", player.Name, starSystem.Name)
		return
	}

	log.Printf("player [ %s ] joined star-system [ %s ]\n", player.Name, starSystem.Name)

	// store the player
	starSystem.players[player.ID] = player

	// get the object id
	id := starSystem.nextID()

	// create player game object
	gameObject := objects.NewPlayer(id, mathf.NewZeroVec3())

	// add the game object to world
	starSystem.world.AddObject(gameObject)

	// map player to game object
	starSystem.playersObject[player.ID] = id

}

func (starSystem *StarSystem) UpdatePlayer(player *player.Player, move *mathf.Vec3, rotation *mathf.Vec3) error {
	// lock the star system
	starSystem.lock.Lock()
	defer starSystem.lock.Unlock()

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

	// finally get the object from the world
	object := starSystem.world.GetObject(gameObjectID)

	if object == nil {
		return errors.Errorf("unable to find game-object with id [ %d ] for player [ %s ]", gameObjectID, player.Name)
	}

	// cast to game object
	gameObject := object.(objects.GameObject)

	// check if game object a player
	if gameObject.Type() != objects.GameObjectPlayer {
		return errors.Errorf("game-object with id [ %d ] is not a player object for player [ %s ]", gameObjectID, player.Name)
	}

	// convert game object to player
	playerObject := gameObject.(*objects.Player)

	// update the player
	playerObject.UpdateMovement(move, rotation)

	return nil
}

func (starSystem *StarSystem) DropPlayer(player *player.Player) error {
	// lock the star system
	starSystem.lock.Lock()
	defer starSystem.lock.Unlock()

	log.Printf("dropping player [ %s ] from star-system [ %s ]\n", player.Name, starSystem.Name)

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
		return errors.Errorf("unable to find game-object [ %d ] for player [ %s ] in star-system [ %s ]", gameObjectID, player.Name, starSystem.Name)
	}
	// remove the mapping
	delete(starSystem.playersObject, player.ID)
	// remove the game object from world
	starSystem.world.RemoveObject(gameObjectID)

	return nil

}
