package object

import (
	"github.com/dayaftereh/discover/server/game/data"
	"github.com/dayaftereh/discover/server/mathf"
)

type Sun struct {
	id        int64
	color     int64
	radius    float64
	rigidbody *RigidBody
}

func NewSun(id int64) *Sun {
	rigidbody := NewRigidBody(1)
	return &Sun{
		id:        id,
		rigidbody: rigidbody,
	}
}

func (sun *Sun) Load(sunData *data.Sun) {
	// setup sun
	sun.color = sunData.Color
	sun.radius = sunData.Radius

	// setup rigidbody
	sun.rigidbody.Mass = sunData.Mass
	sun.rigidbody.LinearFactor = mathf.NewZeroVec3()
}

func (sun *Sun) ID() int64 {
	return sun.id
}

func (sun *Sun) Radius() float64 {
	return sun.radius
}

func (sun *Sun) RigidBody() *RigidBody {
	return sun.rigidbody
}

func (sun *Sun) Update(delta float64) {
	// update the rigidbody
	sun.rigidbody.Update(delta)
}

func (sun *Sun) Color() int64 {
	return sun.color
}

func (planet *Sun) Type() GameObjectType {
	return SunObject
}
