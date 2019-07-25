package starsystem

import (
	"github.com/dayaftereh/discover/server/game/universe/starsystem/objects"
)

func (starSystem *StarSystem) findGameObjectsByType(gameObjectTpye objects.GameObjectType) map[int64]objects.GameObject {
	worldObjects := starSystem.world.GetObjects()

	founds := make(map[int64]objects.GameObject)
	for id, object := range worldObjects {
		gameObject := object.(objects.GameObject)
		if gameObject.Type() == gameObjectTpye {
			founds[id] = gameObject
		}
	}

	return founds
}
