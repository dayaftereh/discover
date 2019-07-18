package object

import (
	"github.com/dayaftereh/discover/server/mathf"
)

type Planet struct {
	id        int64
	color     int64
	radius    float64
	rigidbody *RigidBody
}

func NewPlanet(id int64, color int64) *Planet {
	rigidbody := NewRigidBody(1e6)

	radius := 10.0
	I := 2.0 * rigidbody.Mass * radius * radius / 5.0
	rigidbody.Inertia = mathf.NewVec3(I, I, I)

	rigidbody.UpdateInertiaWorld(true)

	return &Planet{
		id:        id,
		color:     color,
		radius:    radius,
		rigidbody: rigidbody,
	}
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

}

func (planet *Planet) Color() int64 {
	return planet.color
}

func (planet *Planet) Type() GameObjectType {
	return PlanetObject
}
