package objects

import (
	"github.com/dayaftereh/discover/server/game/engine/physics"
	"github.com/dayaftereh/discover/server/mathf"
)

var GameObjectPlayer GameObjectType = "planet"

type Player struct {
	id        int64
	radius    float64
	rigidbody *physics.RigidBody
	// movement
	move     *mathf.Vec3
	rotation *mathf.Vec3
	// force
	linearForce  *mathf.Vec3
	angularForce *mathf.Vec3
}

func NewPlayer(id int64, position *mathf.Vec3) *Player {
	rigidbody := physics.NewRigidBody()
	rigidbody.Position = position

	rigidbody.LinearDamping = 0.5
	rigidbody.AngularDamping = 0.5

	radius := 1.0
	rigidbody.Inertia = physics.CalculateSphereInertia(radius, rigidbody.Mass)

	rigidbody.UpdateInertiaWorld(true)

	return &Player{
		id:           id,
		radius:       radius,
		rigidbody:    rigidbody,
		move:         nil,
		rotation:     nil,
		linearForce:  mathf.NewVec3(100.0, 100.0, 100.0),
		angularForce: mathf.NewVec3(1.0, 1.0, 1.0),
	}
}

func (player *Player) ID() int64 {
	return player.id
}

func (player *Player) Radius() float64 {
	return player.radius
}

func (player *Player) RigidBody() *physics.RigidBody {
	return player.rigidbody
}

func (player *Player) Update(delta float64) {
	if player.move != nil {
		move := mathf.ClampVec3(player.move, -1.0, 1.0)
		// calculate the linear force
		force := player.linearForce.MultiplyVec(move)
		// apply move force
		player.rigidbody.ApplyLocalForce(force, mathf.NewZeroVec3())
	}

	if player.rotation != nil {
		torque := mathf.ClampVec3(player.rotation, -1.0, 1.0)
		// calculate the angular force
		force := player.angularForce.MultiplyVec(torque)
		// apply the rotaion
		player.rigidbody.AddLocalTorque(force)
	}

	// update the rigidbody
	player.rigidbody.Update(delta)
}

func (player *Player) UpdateMovement(move *mathf.Vec3, rotation *mathf.Vec3) {
	player.move = move
	player.rotation = rotation
}

func (player *Player) Color() int64 {
	return int64(0x34ff81)
}

func (player *Player) Type() GameObjectType {
	return GameObjectPlayer
}

func (player *Player) Destroy() {

}
