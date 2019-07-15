package object

import (
	"github.com/dayaftereh/discover/server/mathf"
)

type Player struct {
	id        int64
	radius    float64
	rigidbody *RigidBody
}

func NewPlayer(id int64, position *mathf.Vec3) *Player {
	rigidbody := NewRigidBody(1.0)
	rigidbody.Position = position

	radius := 1.0
	I := 2.0 * rigidbody.Mass * radius * radius / 5.0
	rigidbody.Inertia = mathf.NewVec3(I, I, I)

	rigidbody.UpdateInertiaWorld(true)

	return &Player{
		id:        id,
		radius:    radius,
		rigidbody: rigidbody,
	}
}

func (player *Player) ID() int64 {
	return player.id
}

func (player *Player) Radius() float64 {
	return player.radius
}

func (player *Player) RigidBody() *RigidBody {
	return player.rigidbody
}

func (player *Player) Update(move *mathf.Vec3, rotation *mathf.Vec3) {
	// apply move force
	player.rigidbody.ApplyForce(move, mathf.NewZeroVec3())

	// apply the rotaion
	player.rigidbody.AddTorque(rotation)
}
