package objects

import (
	"github.com/dayaftereh/discover/server/game/data"
	"github.com/dayaftereh/discover/server/game/engine/physics"
	"github.com/dayaftereh/discover/server/mathf"
)

var GameObjectSun GameObjectType = "sun"

type Sun struct {
	id        int64
	color     int64
	radius    float64
	data      *data.Sun
	rigidbody *physics.RigidBody
}

func NewSun(id int64) *Sun {
	rigidbody := physics.NewRigidBody()
	return &Sun{
		id:        id,
		rigidbody: rigidbody,
	}
}

func (sun *Sun) Load(sunData *data.Sun) {
	sun.data = sunData
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

func (sun *Sun) RigidBody() *physics.RigidBody {
	return sun.rigidbody
}

func (sun *Sun) Update(delta float64) {
	// update the rigidbody
	sun.rigidbody.Update(delta)
}

func (sun *Sun) Color() int64 {
	return sun.color
}

func (sun *Sun) Type() GameObjectType {
	return GameObjectSun
}

func (sun *Sun) Write() *data.Sun {
	return sun.data
}

func (sun *Sun) Destroy() {

}
