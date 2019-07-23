package object

import (
	"log"

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
	orbit           *mathf.Orbit2
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

	planet.orbit = mathf.NewOrbitFromVectors(data.Mass, data.Position, mathf.NewVec3(0, -7000, 0))

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

	log.Println("before")
	planet.orbit.Update(delta)
	log.Println("after")

	planet.rigidbody.Position = planet.orbit.Position()
	log.Printf("Position: %v\n", planet.rigidbody.Position)
	// update the rigidbody
	//planet.rigidbody.Update(delta)
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
