package world

import (
	"github.com/dayaftereh/discover/server/game/engine/object"
)

type World struct {
	tick    int64
	objects map[int64]object.GameObject
}

func NewWorld() *World {
	return &World{}
}

func (world *World) GetGameObject(id int64) object.GameObject {
	gameObject, ok := world.objects[id]
	if !ok {
		return nil
	}
	return gameObject
}

func (world *World) AddGameObject(gameObject object.GameObject) {
	id := gameObject.ID()
	world.objects[id] = gameObject
}

func (world *World) RemoveGameObject(id int64) object.GameObject {
	gameObject, ok := world.objects[id]
	if !ok {
		return nil
	}

	delete(world.objects, id)

	return gameObject
}

func (world *World) GetGameObjectsInSphere(target object.GameObject, radius float64) map[int64]object.GameObject {
	// get the body
	body := target.Body()

	contains := make(map[int64]object.GameObject)

	for id, gameObject := range world.objects {
		// get the body of the game object
		gameObjectBody := gameObject.Body()

		// clauclate the distance
		distance := body.Position.DistanceTo(gameObjectBody.Position)

		if distance <= radius {
			contains[id] = gameObject
		}
	}

	return contains
}

func (world *World) GetTick() int64 {
	return world.tick
}

func (world *World) Update(delta float64) {
	world.tick++

}
