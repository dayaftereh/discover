package starsystem

import (
	types "github.com/dayaftereh/discover/server/api/types/connection"
	"github.com/dayaftereh/discover/server/game/engine/world"
	"github.com/dayaftereh/discover/server/game/universe/starsystem/objects"
)

func objectsToOutbound(gameObjects map[int64]world.Object) map[int64]*types.GameObject {
	outbound := make(map[int64]*types.GameObject)
	for id, object := range gameObjects {
		gameObject := object.(objects.GameObject)
		outbound[id] = gameObjectToOutbound(gameObject)
	}
	return outbound
}

func gameObjectToOutbound(gameObject objects.GameObject) *types.GameObject {
	// get the body of the game object
	rigidbody := gameObject.RigidBody()

	// get location and roation
	position := rigidbody.Position.Clone()
	rotation := rigidbody.Rotation.ToEuler()

	removeable := false

	// create the outgoing gameobject
	return &types.GameObject{
		Position:   position,
		Rotation:   rotation,
		Removeable: &removeable,
	}
}
