package object

import (
	"github.com/dayaftereh/discover/server/game/data"
	"github.com/dayaftereh/discover/server/mathf"
	"github.com/dayaftereh/discover/server/mathf/orbit"
)

// GravitationalConstant is an empirical physical constant
var GravitationalConstant float64 = 1

type Planet struct {
	id        int64
	color     int64
	radius    float64
	orbit     *orbit.Orbit
	epoch     float64
	rigidbody *RigidBody
}

func NewPlanet(id int64) *Planet {
	rigidbody := NewRigidBody(1)

	return &Planet{
		id:        id,
		rigidbody: rigidbody,
	}
}

func (planet *Planet) Load(orbit *orbit.Orbit, data *data.Planet) {
	// setup the orbit
	planet.orbit = orbit
	planet.epoch = 0.0

	// setup planet
	planet.color = data.Color
	planet.radius = data.Radius

	// setup rigidbody
	planet.rigidbody.Mass = data.Mass
	planet.rigidbody.Position = orbit.Position()
	planet.rigidbody.LinearFactor = mathf.NewZeroVec3()

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
	// update the rigidbody
	planet.rigidbody.Update(delta)
	// move the planet
	planet.epoch += delta
	// get the current orbit
	currentOrbit := planet.orbit.Update(planet.epoch)
	// update planet to current orbit
	planet.rigidbody.Position = currentOrbit.Position()
}

func (planet *Planet) Color() int64 {
	return planet.color
}

func (planet *Planet) Type() GameObjectType {
	return PlanetObject
}

func (planet *Planet) Write() *data.Planet {
	return &data.Planet{
		Color:  planet.color,
		Mass:   planet.rigidbody.Mass,
		Radius: planet.radius,
	}
}
