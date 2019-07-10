package object

import (
	"github.com/dayaftereh/discover/server/game/engine/shape"
	"github.com/dayaftereh/discover/server/mathf"
)

type Body struct {
	Mass float64

	// Whether to produce contact forces when in contact with other bodies
	CollisionResponse bool

	// Set to true if you don't want the body to rotate
	FixedRotation bool

	// World space position of the body.
	Position         *mathf.Vec3
	PreviousPosition *mathf.Vec3

	// World space velocity of the body.
	Velocity *mathf.Vec3
	// Linear force on the body in world space.
	Force *mathf.Vec3

	// World space rotational force on the body, around center of mass.
	Torque *mathf.Vec3

	Inertia *mathf.Vec3

	// World space rotational force on the body, around center of mass.
	Quaternion *mathf.Quaternion

	shapes []shape.Shape

	// Total bounding radius of the Body including its shapes, relative to body.position.
	BoundingRadius float64
}

func NewBody(position *mathf.Vec3, mass float64) *Body {
	return &Body{
		Mass:     mass,
		Position: position,
	}
}

func (body *Body) PointToLocalFrame(worldPoint *mathf.Vec3) *mathf.Vec3 {
	p := worldPoint.Subtract(body.Position)
	r := body.Quaternion.Conjugate().MultiplyVec(p)
	return r
}

func (body *Body) VectorToLocalFrame(worldVector *mathf.Vec3) *mathf.Vec3 {
	r := body.Quaternion.Conjugate().MultiplyVec(worldVector)
	return r
}

func (body *Body) PointToWorldFrame(localPoint *mathf.Vec3) *mathf.Vec3 {
	p := body.Quaternion.MultiplyVec(localPoint)
	r := p.Add(body.Position)
	return r
}

func (body *Body) VectorToWorldFrame(localVector *mathf.Vec3) *mathf.Vec3 {
	r := body.Quaternion.MultiplyVec(localVector)
	return r
}

func (body *Body) AddShape(shape *shape.Shape) {
	body.shapes = append(body.shapes, *shape)

	body.UpdateMassProperties()
	body.UpdateBoundingRadius()
}

func (body *Body) UpdateBoundingRadius() {
	radius := float64(0)
	for _, shape := range body.shapes {
		shape.UpdateBoundingSphereRadius()
		r := shape.BoundingSphereRadius()
		if r > radius {
			radius = r
		}
	}
	body.BoundingRadius = radius
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
