package world

import "github.com/dayaftereh/discover/server/game/engine/object"

type World struct {
	ID   int64
	Tick int64

	Bodies map[int64]*object.Body
}

func NewWorld() *World {
	return &World{}
}

func (world *World) AddBody(body *object.Body) {
	world.Bodies[body.ID] = body
}

func (world *World) RemoveBody(id int64) *object.Body {
	body, ok := world.Bodies[id]
	if !ok {
		return nil
	}

	delete(world.Bodies, id)

	return body
}

func (world *World) Update(dt float64) {
	world.Tick++

}

func (world *World) internalUpdate(dt float64) {

}
