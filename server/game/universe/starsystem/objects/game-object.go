package objects

import "github.com/dayaftereh/discover/server/game/engine/physics"

type GameObject interface {
	ID() int64
	Color() int64
	Type() GameObjectType
	Radius() float64
	RigidBody() *physics.RigidBody
	Update(delta float64)
	Destroy()
}
