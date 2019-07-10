package object

import (
	"github.com/dayaftereh/discover/server/mathf"
)

type Planet struct {
	id   int64
	body *Body
}

func NewPlanet(id int64, position *mathf.Vec3) *Planet {
	return &Planet{
		id:   id,
		body: NewBody(position, 1e6),
	}
}

func (planet *Planet) ID() int64 {
	return planet.id
}

func (planet *Planet) Body() *Body {
	return planet.body
}
