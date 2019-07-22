package object

import (
	"log"

	"github.com/dayaftereh/discover/server/game/data"
	"github.com/dayaftereh/discover/server/mathf"
)

var GravitationalConstant float64 = 1

type Planet struct {
	id              int64
	color           int64
	radius          float64
	sunMass         float64
	rigidbody       *RigidBody
	initialPosition *mathf.Vec3
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
	planet.initialPosition = data.Position.Clone()

	planet.rigidbody.LinearDamping = 0.0

	// set rigidbody
	planet.rigidbody.Mass = data.Mass
	planet.rigidbody.Position = data.Position.Clone()

	// add inital force
	//planet.rigidbody.ApplyLocalForce(mathf.NewVec3(0, 0, 1).Multiply(100000.0), mathf.NewZeroVec3())

	planet.rigidbody.Velocity = mathf.NewVec3(0, 1, 1).Multiply(1000)

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

	log.Printf("delta:%f\n", delta)
	center := mathf.NewZeroVec3()

	// calulate distance to center of star system
	r := planet.rigidbody.Position.DistanceTo(center)

	log.Printf("R:%f\n", r)
	log.Printf("Mass:%f\n", planet.rigidbody.Mass)
	log.Printf("sunMass:%f\n", planet.sunMass)

	if r < 10 {
		r = 10
	}

	// calculate the force
	force := (GravitationalConstant * planet.sunMass * planet.rigidbody.Mass) / (r * r)

	log.Printf("force:%f\n", force)

	// get vector to center of star syste,
	v := center.Subtract(planet.rigidbody.Position).Normalize()

	log.Printf("force:%f\n", v)

	forceV := v.Multiply(force)

	planet.rigidbody.ApplyForce(forceV, center)

	log.Printf("Velocity0:%v\n", planet.rigidbody.Velocity)

	// update the rigidbody
	planet.rigidbody.Update(delta)

	log.Printf("Position:%v\n", planet.rigidbody.Position)
	log.Printf("Velocity:%v\n", planet.rigidbody.Velocity)
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
	}
}
