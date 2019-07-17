package object

import (
	"log"

	"github.com/dayaftereh/discover/server/mathf"
)

type Player struct {
	id        int64
	radius    float64
	rigidbody *RigidBody
	//
	move     *mathf.Vec3
	rotation *mathf.Vec3
}

func NewPlayer(id int64, position *mathf.Vec3) *Player {
	rigidbody := NewRigidBody(5.0)
	rigidbody.Position = position

	radius := 1.0
	I := 2.0 * rigidbody.Mass * radius * radius / 5.0
	rigidbody.Inertia = mathf.NewVec3(I, I, I)

	//rigidbody.UpdateInertiaWorld(true)

	return &Player{
		id:        id,
		radius:    radius,
		rigidbody: rigidbody,
		move:      nil,
		rotation:  nil,
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

func (player *Player) Update(delta float64) {
	if player.move != nil {
		move := mathf.ClampVec3(player.move, -1.0, 1.0)
		log.Printf("move: %v", move)
		// calculate the force
		force := move.Multiply(350.0)
		// apply move force
		player.rigidbody.ApplyLocalForce(force, mathf.NewZeroVec3())
	}

	if player.rotation != nil {
		torque := mathf.ClampVec3(player.rotation, -1.0, 1.0)
		// apply the rotaion
		player.rigidbody.AddLocalTorque(torque)
	}

	player.rigidbody.Update(delta)
}

func (player *Player) UpdateMovement(move *mathf.Vec3, rotation *mathf.Vec3) {
	player.move = move
	player.rotation = rotation
}
