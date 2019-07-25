package objects

import (
	"github.com/dayaftereh/discover/server/game/data"
	"github.com/dayaftereh/discover/server/game/engine/physics"
	"github.com/dayaftereh/discover/server/mathf"
	"github.com/dayaftereh/discover/server/mathf/orbit"
)

var GameObjectPlanet GameObjectType = "planet"

type Planet struct {
	id        int64
	color     int64
	radius    float64
	orbit     *orbit.Orbit
	rigidbody *physics.RigidBody
	// custom
	epoch float64
	data  *data.Planet
}

func NewPlanet(id int64) *Planet {
	rigidbody := physics.NewRigidBody()

	return &Planet{
		id:        id,
		rigidbody: rigidbody,
	}
}

func (planet *Planet) Load(sun *Sun, data *data.Planet) {
	// store the data for the planet
	planet.data = data
	// setup the orbit
	planet.orbit = planet.createOrbit(sun, data.Orbit)
	planet.epoch = 0.0

	// setup planet
	planet.color = data.Color
	planet.radius = data.Radius

	// setup rigidbody
	planet.rigidbody.Mass = data.Mass
	planet.rigidbody.Position = planet.orbit.Position()
	planet.rigidbody.LinearFactor = mathf.NewZeroVec3()

	// load the Inertia
	planet.rigidbody.Inertia = physics.CalculateSphereInertia(planet.radius, planet.rigidbody.Mass)
	planet.rigidbody.UpdateInertiaWorld(true)
}

func (planet *Planet) createOrbit(sun *Sun, orbitData *data.Orbit) *orbit.Orbit {
	// get the RigidBody of the sun
	sunRigidBody := sun.RigidBody()
	// get th radius of the sun
	centralBodyRadius := sun.Radius()
	// calucate mu
	mu := mathf.GravitationalConstant * sunRigidBody.Mass

	// create the orbit
	orbit := orbit.OrbitFromParams(&orbit.OrbitParameter{
		MU:                  &mu,
		CentralBodyRadius:   &centralBodyRadius,
		Apogee:              &orbitData.Apogee,
		Perigee:             &orbitData.Perigee,
		Eccentricity:        &orbitData.Eccentricity,
		Inclination:         &orbitData.Inclination,
		RightAscension:      &orbitData.RightAscension,
		ArgumentOfPeriapsis: &orbitData.ArgumentOfPeriapsis,
	})

	return orbit
}

func (planet *Planet) ID() int64 {
	return planet.id
}

func (planet *Planet) Radius() float64 {
	return planet.radius
}

func (planet *Planet) RigidBody() *physics.RigidBody {
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
	return GameObjectPlanet
}

func (planet *Planet) Write() *data.Planet {
	return planet.data
}

func (planet *Planet) Destroy() {

}
