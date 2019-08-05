package objects

import (
	"github.com/dayaftereh/discover/server/game/engine/physics"
	"github.com/dayaftereh/discover/server/game/persistence/types"
	"github.com/dayaftereh/discover/server/mathf"
)

var GameObjectSun GameObjectType = "sun"

type Sun struct {
	id        int64
	rigidbody *physics.RigidBody
	// public
	Data *types.Sun
}

func NewSun(id int64, data *types.Sun) *Sun {
	return &Sun{
		id:        id,
		Data:      data,
		rigidbody: physics.NewRigidBody(),
	}
}

func (sun *Sun) Init() error {
	// setup rigidbody
	sun.rigidbody.Mass = sun.Data.Mass
	sun.rigidbody.LinearFactor = mathf.NewZeroVec3()

	return nil
}

func (sun *Sun) ID() int64 {
	return sun.id
}
func (sun *Sun) RigidBody() *physics.RigidBody {
	return sun.rigidbody
}

func (sun *Sun) Update(delta float64) {
	// update the rigidbody
	sun.rigidbody.Update(delta)
}

func (sun *Sun) Type() GameObjectType {
	return GameObjectSun
}

func (sun *Sun) Radius() float64 {
	return sun.Data.EcosphereRadius
}

func (sun *Sun) Destroy() {

}
