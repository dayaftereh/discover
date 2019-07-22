package object

import (
	"log"
	"math"

	"github.com/dayaftereh/discover/server/game/data"
	"github.com/dayaftereh/discover/server/mathf"
)

// GravitationalConstant is an empirical physical constant
var GravitationalConstant float64 = 1

type Planet struct {
	id              int64
	color           int64
	radius          float64
	sunMass         float64
	rigidbody       *RigidBody
	initialPosition *mathf.Vec3
	initialForce    *mathf.Vec3
}

func NewPlanet(id int64) *Planet {
	rigidbody := NewRigidBody(1)

	return &Planet{
		id:        id,
		rigidbody: rigidbody,
	}
}

func (planet *Planet) Load(sunMass float64, data *data.Planet) {
	planet.sunMass = sunMass

	// set planet
	planet.color = data.Color
	planet.radius = data.Radius
	planet.initialForce = data.Force
	planet.initialPosition = data.Position.Clone()

	// set rigidbody
	planet.rigidbody.Mass = data.Mass
	planet.rigidbody.Position = data.Position.Clone()

	// load the Inertia
	planet.rigidbody.Inertia = CalculateSphereInertia(planet.radius, planet.rigidbody.Mass)
	planet.rigidbody.UpdateInertiaWorld(true)
}

func (planet *Planet) ID() int64 {
	return planet.id
}

func (planet *Planet) Radius() float64 {
	return planet.radius
}

func (planet *Planet) RigidBody() *RigidBody {
	return planet.rigidbody
}

func (planet *Planet) Update(delta float64) {
	// center / sun
	center := mathf.NewZeroVec3()

	// calulate distance to center of star system
	r := planet.rigidbody.Position.DistanceTo(center)

	// Newton's law of universal gravitation
	// f = g * (M1 * M2) / r^2
	force := GravitationalConstant * ((planet.sunMass * planet.rigidbody.Mass) / (r * r))

	// get vector to center of star syste,
	direction := center.Subtract(planet.rigidbody.Position).Normalize()

	// calculate force vector
	forceVector := direction.Multiply(force)

	// apply the force to the planet
	planet.rigidbody.ApplyForce(forceVector, center)

	// apply the inital force of the planet

	v := math.Sqrt((GravitationalConstant * planet.sunMass) / 200.0)
	initialForce := planet.rigidbody.Mass * v / delta

	planet.rigidbody.ApplyLocalForce(mathf.NewVec3(0, 0, initialForce), center)

	// update the rigidbody
	planet.rigidbody.Update(delta)

	log.Printf("Position: %v", planet.rigidbody.Position)
}

func (planet *Planet) Color() int64 {
	return planet.color
}

func (planet *Planet) Type() GameObjectType {
	return PlanetObject
}

func (planet *Planet) Write() *data.Planet {
	return &data.Planet{
		Color:    planet.color,
		Mass:     planet.rigidbody.Mass,
		Position: planet.initialPosition.Clone(),
		Radius:   planet.radius,
		Force:    planet.initialForce,
	}
}
