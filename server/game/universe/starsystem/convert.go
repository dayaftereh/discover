package starsystem

import (
	types "github.com/dayaftereh/discover/server/api/types/connection"
	"github.com/dayaftereh/discover/server/game/engine/object"
)

func gameObjectsToOutbound(gameObjects map[int64]object.GameObject) map[int64]*types.GameObject {
	outbound := make(map[int64]*types.GameObject)
	for id, gameObject := range gameObjects {
		outbound[id] = gameObjectToOutbound(gameObject)
	}
	return outbound
}

func gameObjectToOutbound(gameObject object.GameObject) *types.GameObject {
	// get the body of the game object
	body := gameObject.Body()

	// get the radius of the game object
	radius := body.BoundingRadius

	// get location and roation
	position := body.Position.Clone()
	rotation := body.Quaternion.ToEuler()

	removeable := false

	// create the outgoing gameobject
	return &types.GameObject{
		Radius:     &radius,
		Position:   position,
		Rotation:   rotation,
		Removeable: &removeable,
	}
}
