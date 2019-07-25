package world

import "github.com/dayaftereh/discover/server/game/engine/physics"

type Object interface {
	ID() int64
	Radius() float64
	RigidBody() *physics.RigidBody
	Update(delta float64)
}
