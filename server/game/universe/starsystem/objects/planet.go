package objects

import (
	"github.com/dayaftereh/discover/server/game/engine/physics"
	"github.com/dayaftereh/discover/server/game/persistence/types"
	"github.com/dayaftereh/discover/server/mathf"
	"github.com/dayaftereh/discover/server/mathf/orbit"
)

var GameObjectPlanet GameObjectType = "planet"

type Planet struct {
	id        int64
	sun       *Sun
	orbit     *orbit.Orbit
	rigidbody *physics.RigidBody
	// custom
	epoch float64
	// public
	Data *types.Planet
}

func NewPlanet(id int64, data *types.Planet, sun *Sun) *Planet {
	return &Planet{
		id:        id,
		sun:       sun,
		Data:      data,
		rigidbody: physics.NewRigidBody(),
	}
}

func (planet *Planet) Init() error {
	// setup the orbit
	//planet.orbit = planet.createOrbit()
	planet.epoch = 0.0

	// setup rigidbody
	planet.rigidbody.Mass = planet.Data.Mass
	//planet.rigidbody.Position = planet.orbit.Position()
	planet.rigidbody.LinearFactor = mathf.NewZeroVec3()

	// load the Inertia
	planet.rigidbody.Inertia = physics.CalculateSphereInertia(planet.Data.Radius, planet.rigidbody.Mass)
	planet.rigidbody.UpdateInertiaWorld(true)

	return nil
}

func (planet *Planet) createOrbit() *orbit.Orbit {
	// calucate mu
	mu := mathf.GravitationalConstant * planet.sun.Data.Mass

	// create the orbit
	orbit := orbit.OrbitFromParams(&orbit.OrbitParameter{
		MU: &mu,
	})

	return orbit
}

func (planet *Planet) ID() int64 {
	return planet.id
}

func (planet *Planet) RigidBody() *physics.RigidBody {
	return planet.rigidbody
}

func (planet *Planet) Update(delta float64) {
	/*// update the rigidbody
	planet.rigidbody.Update(delta)
	// move the planet
	planet.epoch += delta
	// get the current orbit
	currentOrbit := planet.orbit.Update(planet.epoch)
	// update planet to current orbit
	planet.rigidbody.Position = currentOrbit.Position()*/
}

func (planet *Planet) Type() GameObjectType {
	return GameObjectPlanet
}

func (planet *Planet) Radius() float64 {
	return planet.Data.Radius
}

func (planet *Planet) Destroy() {

}
