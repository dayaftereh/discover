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
		playerObject := starSystem.world.GetObject(playerObjectID)

		// get all objects in player range
		worldObjects := starSystem.world.GetGameObjectsInSphere(playerObject, 10000.0)

		// convert objects to outbound
		outboundObjects := objectsToOutbound(worldObjects)

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
			Objects: outboundObjects,
		}

		// push the update for the player
		player.Push(worldUpdate)
	}
}
