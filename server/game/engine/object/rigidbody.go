package object

import "github.com/dayaftereh/discover/server/mathf"

type RigidBody struct {
	Mass float64

	Position *mathf.Vec3

	// World space rotational force on the body, around center of mass.
	Rotation *mathf.Quaternion

	// World space velocity of the body.
	Velocity *mathf.Vec3
	// Angular velocity of the body, in world space. Think of the angular velocity as a vector, which the body rotates around. The length of this vector determines how fast (in radians per second) the body rotates.
	AngularVelocity *mathf.Vec3

	// Linear force on the body in world space.
	Force *mathf.Vec3
	// World space rotational force on the body, around center of mass.
	Torque *mathf.Vec3

	// LinearFactor use to limit the motion along any world axis. (1,1,1) will allow motion along all axes while (0,0,0) allows none.
	LinearFactor *mathf.Vec3

	// AngularFactor use to limit the rotational motion along any world axis. (1,1,1) will allow rotation along all axes while (0,0,0) allows none.
	AngularFactor *mathf.Vec3

	Inertia      *mathf.Vec3
	InertiaWorld *mathf.Mat3
}

func NewRigidBody(mass float64) *RigidBody {
	return &RigidBody{
		Position: mathf.NewZeroVec3(),
		Rotation: mathf.NewZeroQuaternion(),

		Velocity:        mathf.NewZeroVec3(),
		AngularVelocity: mathf.NewZeroVec3(),

		Force:  mathf.NewZeroVec3(),
		Torque: mathf.NewZeroVec3(),
	}
}

func (rigidbody *RigidBody) InverseMass() float64 {
	if rigidbody.Mass > 0 {
		return 1.0 / rigidbody.Mass
	}
	return 0.0
}

func (rigidbody *RigidBody) InverseInertia() *mathf.Vec3 {
	return mathf.NewVec3(
		1.0/rigidbody.Inertia.X,
		1.0/rigidbody.Inertia.Y,
		1.0/rigidbody.Inertia.Z,
	)
}

func (rigidbody *RigidBody) InverseInertiaWorld() *mathf.Mat3 {
	return rigidbody.InertiaWorld
}

func (rigidbody *RigidBody) PointToLocalFrame(worldPoint *mathf.Vec3) *mathf.Vec3 {
	p := worldPoint.Subtract(rigidbody.Position)
	r := rigidbody.Rotation.Conjugate().MultiplyVec(p)
	return r
}

func (rigidbody *RigidBody) VectorToLocalFrame(worldVector *mathf.Vec3) *mathf.Vec3 {
	r := rigidbody.Rotation.Conjugate().MultiplyVec(worldVector)
	return r
}

func (rigidbody *RigidBody) PointToWorldFrame(localPoint *mathf.Vec3) *mathf.Vec3 {
	p := rigidbody.Rotation.MultiplyVec(localPoint)
	r := p.Add(rigidbody.Position)
	return r
}

func (rigidbody *RigidBody) VectorToWorldFrame(localVector *mathf.Vec3) *mathf.Vec3 {
	r := rigidbody.Rotation.MultiplyVec(localVector)
	return r
}

func (rigidbody *RigidBody) ApplyForce(force *mathf.Vec3, relativePoint *mathf.Vec3) {
	// Add linear force
	rigidbody.Force = rigidbody.Force.Add(force)

	// Compute produced rotational force
	rotForce := relativePoint.Cross(force)

	// Add rotational force
	rigidbody.Torque = rigidbody.Torque.Add(rotForce)

}

func (rigidbody *RigidBody) ApplyLocalForce(localForce *mathf.Vec3, localPoint *mathf.Vec3) {
	// Transform the force vector to world space
	worldForce := rigidbody.VectorToWorldFrame(localForce)
	relativePointWorld := rigidbody.VectorToWorldFrame(localPoint)

	rigidbody.ApplyForce(worldForce, relativePointWorld)
}

func (rigidbody *RigidBody) ApplyImpulse(impulse *mathf.Vec3, relativePoint *mathf.Vec3) {
	// Compute produced central impulse velocity
	velo := impulse.Multiply(rigidbody.InverseMass())

	// Add linear impulse
	rigidbody.Velocity = rigidbody.Velocity.Add(velo)

	// Compute produced rotational impulse velocity
	rotVelo := relativePoint.Cross(impulse)

	invInertia := rigidbody.InverseInertiaWorld()
	/*
	   rotVelo.x *= this.invInertia.x;
	   rotVelo.y *= this.invInertia.y;
	   rotVelo.z *= this.invInertia.z;
	*/
	rotVeloInertia := invInertia.MultiplyVec(rotVelo)

	// Add rotational Impulse
	rigidbody.AngularVelocity = rigidbody.AngularVelocity.Add(rotVeloInertia)
}

func (rigidbody *RigidBody) ApplyLocalImpulse(localImpulse *mathf.Vec3, localPoint *mathf.Vec3) {
	// Transform the force vector to world space
	worldImpulse := rigidbody.VectorToWorldFrame(localImpulse)
	relativePointWorld := rigidbody.VectorToWorldFrame(localPoint)

	rigidbody.ApplyImpulse(worldImpulse, relativePointWorld)

}

func (rigidbody *RigidBody) Update(delta float64) {

	invMassDelta := rigidbody.InverseMass() * delta

	velo := mathf.NewVec3(
		rigidbody.Velocity.X*invMassDelta*rigidbody.LinearFactor.X,
		rigidbody.Velocity.Y*invMassDelta*rigidbody.LinearFactor.Y,
		rigidbody.Velocity.Z*invMassDelta*rigidbody.LinearFactor.Z,
	)

	tx := rigidbody.Torque.X * rigidbody.AngularFactor.X
	ty := rigidbody.Torque.Y * rigidbody.AngularFactor.Y
	tz := rigidbody.Torque.Z * rigidbody.AngularFactor.Z

	invInertia := rigidbody.InverseInertiaWorld()
	e := invInertia.Elements()

	angularVelo := mathf.NewVec3(
		delta*(e[0]*tx+e[1]*ty+e[2]*tz),
		delta*(e[3]*tx+e[4]*ty+e[5]*tz),
		delta*(e[6]*tx+e[7]*ty+e[8]*tz),
	)

	// Use new velocity  - leap frog
	// update position
	veloDelta := velo.Multiply(delta)
	rigidbody.Position = rigidbody.Position.Add(veloDelta)

	// update rotation
	rigidbody.Rotation = rigidbody.Rotation.Integrate(angularVelo, delta, rigidbody.AngularFactor)
}