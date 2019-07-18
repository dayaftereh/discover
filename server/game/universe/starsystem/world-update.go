package starsystem

import (
	types "github.com/dayaftereh/discover/server/api/types/connection"
	"github.com/dayaftereh/discover/server/utils"
)

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

		// get all objects in player range
		playerObjects := starSystem.world.GetGameObjectsInSphere(gameObject, 100.0)

		// convert game objects for outbound
		gameObjects := gameObjectsToOutbound(playerObjects)

		// get the world update tick
		tick := starSystem.world.GetTick()

		// create the player update
		playerUpdate := &types.PlayerUpdate{
			GameObjectId: &playerObjectID,
		}

		// create the world update
		worldUpdate := &types.WorldUpdateMessage{
			Type:    types.WorldUpdate,
			Tick:    &tick,
			Time:    &now,
			Player:  playerUpdate,
			Objects: gameObjects,
		}

		// push the update for the player
		player.Push(worldUpdate)
	}
}
