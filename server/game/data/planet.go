package data

import (
	"github.com/dayaftereh/discover/server/mathf"
)

type Planet struct {
	Color    int64
	Mass     float64
	Radius   float64
	Position *mathf.Vec3
	Force    *mathf.Vec3
}
