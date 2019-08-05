package objects

import "github.com/dayaftereh/discover/server/game/engine/physics"

type GameObject interface {
	ID() int64
	Type() GameObjectType
	RigidBody() *physics.RigidBody
	Update(delta float64)
	Destroy()
}
