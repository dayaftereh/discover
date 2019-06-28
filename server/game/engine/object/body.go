package object

import (
	"github.com/dayaftereh/discover/server/game/engine/shape"
	"github.com/dayaftereh/discover/server/math"
)

type Body struct {
	// ID is the unique identifier for the body
	ID string

	Mass float64

	// Whether to produce contact forces when in contact with other bodies
	CollisionResponse bool

	// Set to true if you don't want the body to rotate
	FixedRotation bool

	// World space position of the body.
	Position         math.Vec3
	PreviousPosition math.Vec3

	// World space velocity of the body.
	Velocity math.Vec3
	// Linear force on the body in world space.
	Force math.Vec3

	// World space rotational force on the body, around center of mass.
	Torque math.Vec3

	// World space rotational force on the body, around center of mass.
	Quaternion math.Quaternion

	shapes []shape.Shape

	// Total bounding radius of the Body including its shapes, relative to body.position.
	BoundingRadius float64
}
