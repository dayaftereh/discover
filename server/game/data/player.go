package data

import (
	"github.com/dayaftereh/discover/server/mathf"
)

type Player struct {
	Name       string
	StarSystem *int64
	Location   *mathf.Vec3
}
