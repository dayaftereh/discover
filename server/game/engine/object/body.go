package object

import (
	"github.com/dayaftereh/discover/server/game/engine/shape"
	"github.com/dayaftereh/discover/server/mathf"
)

type Body struct {
	// ID is the unique identifier for the body
	ID int64

	Mass float64

	// Whether to produce contact forces when in contact with other bodies
	CollisionResponse bool

	// Set to true if you don't want the body to rotate
	FixedRotation bool

	// World space position of the body.
	Position         mathf.Vec3
	PreviousPosition mathf.Vec3

	// World space velocity of the body.
	Velocity mathf.Vec3
	// Linear force on the body in world space.
	Force mathf.Vec3

	// World space rotational force on the body, around center of mass.
	Torque mathf.Vec3

	// World space rotational force on the body, around center of mass.
	Quaternion mathf.Quaternion

	shapes []shape.Shape

	// Total bounding radius of the Body including its shapes, relative to body.position.
	BoundingRadius float64
}

func (body *Body) PointToLocalFrame(vec *mathf.Vec3) *mathf.Vec3 {
	return nil
}

func (body *Body) VectorToLocalFrame(vec *mathf.Vec3) *mathf.Vec3 {
	return nil
}

func (body *Body) PointToWorldFrame(vec *mathf.Vec3) *mathf.Vec3 {
	return nil
}

func (body *Body) VectorToWorldFrame(vec *mathf.Vec3) *mathf.Vec3 {
	return nil
}

func (body *Body) AddShape(shape *shape.Shape) {

}

func (body *Body) UpdateBoundingRadius() {

}

func (body *Body) UpdateMassProperties() {

}

func (body *Body) ApplyForce(force *mathf.Vec3, relativePoint *mathf.Vec3) {

}

func (body *Body) ApplyLocalForce(force *mathf.Vec3, localPoint *mathf.Vec3) {

}

func (body *Body) ApplyImpulse(impulse *mathf.Vec3, relativePoint *mathf.Vec3) {

}

func (body *Body) ApplyLocalImpulse(impulse *mathf.Vec3, localPoint *mathf.Vec3) {

}

func (body *Body) VelocityAtWorldPoint(worldPoint *mathf.Vec3) *mathf.Vec3 {
	return nil
}

func (body *Body) Integrate(dt float64, quaternionNormalize bool) *mathf.Vec3 {
	return nil
}
