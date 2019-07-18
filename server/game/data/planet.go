package data

import (
	"github.com/dayaftereh/discover/server/mathf"
)

type Planet struct {
	Color    *uint64
	Location *mathf.Vec3
	Rotation *mathf.Vec3
}
